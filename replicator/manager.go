package replicator

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

// DeviceRuntimeStatus is the compact truthful state consumed by the OS.js UI.
type DeviceRuntimeStatus struct {
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	Running   bool   `json:"running"`
	Cycles    uint64 `json:"cycles"`
	Source    string `json:"source_status"`
	LastPoll  string `json:"last_poll,omitempty"`
	LastError string `json:"last_error,omitempty"`
}

type managedRuntime struct {
	cfg    Config
	cancel context.CancelFunc
	done   chan struct{}

	mu    sync.RWMutex
	state RuntimeState
}

func newManagedRuntime(cfg Config) *managedRuntime {
	return &managedRuntime{cfg: cfg, done: make(chan struct{})}
}

func (r *managedRuntime) start() {
	ctx, cancel := context.WithCancel(context.Background())
	r.cancel = cancel
	go func() {
		defer close(r.done)
		r.setRunning(true)
		defer r.setRunning(false)
		r.runCycle()
		ticker := time.NewTicker(time.Duration(r.cfg.Source.PollIntervalMS) * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				r.runCycle()
			}
		}
	}()
}

func (r *managedRuntime) stop() {
	if r.cancel == nil {
		return
	}
	r.cancel()
	<-r.done
}

func (r *managedRuntime) runCycle() {
	start := time.Now()
	_, err := runConfigCycle(r.cfg)
	end := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.state.Cycles++
	r.state.LastCycleStart = start
	r.state.LastCycleEnd = end
	if err != nil {
		r.state.LastSuccess = false
		r.state.LastError = err.Error()
		return
	}
	r.state.LastSuccess = true
	r.state.LastError = ""
}

func (r *managedRuntime) setRunning(running bool) {
	r.mu.Lock()
	r.state.Running = running
	r.mu.Unlock()
}

func (r *managedRuntime) snapshot() RuntimeState {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.state
}

// RuntimeManager owns all UI-configured Replicator poll loops. Shared MMA2
// structure is composed at apply/boot boundaries, never once per poll cycle.
type RuntimeManager struct {
	store   Store
	timeout time.Duration

	mu       sync.RWMutex
	document Document
	runtimes map[string]*managedRuntime
}

func NewRuntimeManager(store Store) *RuntimeManager {
	return &RuntimeManager{store: store, timeout: DefaultApplyTimeout, runtimes: make(map[string]*managedRuntime)}
}

// Boot restores the persisted document. It never requests an MMA2 restart on
// ordinary process boot; the independently managed MMA2 appliance is expected
// to auto-start from the already-persisted shared config.
func (m *RuntimeManager) Boot() error {
	doc, err := m.store.LoadDocument()
	if err != nil {
		return err
	}
	resolved, _, err := m.store.ComposeDocumentDestinations(doc)
	if err != nil {
		return err
	}
	ports := documentDestinationPorts(resolved)
	if len(ports) > 0 {
		if err := waitPortsReady(ports, m.timeout); err != nil {
			return err
		}
	}
	if err := m.store.SaveDocument(resolved); err != nil {
		return err
	}
	return m.replaceRuntimes(resolved)
}

// Apply validates, resolves ownership-aware automatic destinations, refreshes
// the Replicator-owned shared-config slice, requests an MMA2 restart only when
// the served destination structure actually changed, then persists and starts
// poll loops from the committed document.
func (m *RuntimeManager) Apply(edited Document) (Document, bool, error) {
	resolved, err := m.store.ResolveDocumentDestinations(edited)
	if err != nil {
		return Document{}, false, err
	}
	previous, err := m.store.LoadDocument()
	if err != nil {
		return Document{}, false, err
	}
	previousResolved, err := m.store.ResolveDocumentDestinations(previous)
	if err != nil && len(previous.Devices) > 0 {
		return Document{}, false, fmt.Errorf("resolve persisted Replicator document: %w", err)
	}
	structural := !sameMMA2Structure(previousResolved, resolved)

	m.stopAll()
	resolved, effective, err := m.store.ComposeDocumentDestinations(resolved)
	if err != nil {
		return Document{}, structural, err
	}
	if structural {
		if err := m.store.requestMMA2Restart(effective, documentDestinationPorts(resolved), m.timeout); err != nil {
			return Document{}, true, err
		}
	}
	if err := m.store.SaveDocument(resolved); err != nil {
		return Document{}, structural, err
	}
	if err := m.replaceRuntimes(resolved); err != nil {
		return Document{}, structural, err
	}
	return resolved, structural, nil
}

func (m *RuntimeManager) Document() Document {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return cloneDocument(m.document)
}

func (m *RuntimeManager) Status(name string) (DeviceRuntimeStatus, error) {
	m.mu.RLock()
	doc := cloneDocument(m.document)
	runtime := m.runtimes[name]
	m.mu.RUnlock()
	var device *DeviceDefinition
	for i := range doc.Devices {
		if doc.Devices[i].Name == name {
			device = &doc.Devices[i]
			break
		}
	}
	if device == nil {
		return DeviceRuntimeStatus{}, fmt.Errorf("device %q not found", name)
	}
	status := DeviceRuntimeStatus{Name: name, Enabled: device.Enabled, Source: "WAITING"}
	if !device.Enabled {
		status.Source = "DISABLED"
		return status, nil
	}
	if runtime == nil {
		return status, nil
	}
	snapshot := runtime.snapshot()
	status.Running = snapshot.Running
	status.Cycles = snapshot.Cycles
	status.LastError = snapshot.LastError
	if !snapshot.LastCycleEnd.IsZero() {
		status.LastPoll = snapshot.LastCycleEnd.UTC().Format(time.RFC3339Nano)
	}
	switch {
	case snapshot.Cycles == 0:
		status.Source = "WAITING"
	case snapshot.LastSuccess:
		status.Source = "OK"
	default:
		status.Source = "ERROR"
	}
	return status, nil
}

func (m *RuntimeManager) Stop() {
	m.stopAll()
}

func (m *RuntimeManager) replaceRuntimes(doc Document) error {
	newRuntimes := make(map[string]*managedRuntime)
	for _, device := range doc.Devices {
		if !device.Enabled {
			continue
		}
		cfg, err := device.runtimeConfig()
		if err != nil {
			return fmt.Errorf("device %q: %w", device.Name, err)
		}
		if err := validateCycleMapping(cfg); err != nil {
			return fmt.Errorf("device %q: %w", device.Name, err)
		}
		newRuntimes[device.Name] = newManagedRuntime(cfg)
	}
	m.mu.Lock()
	m.document = cloneDocument(doc)
	m.runtimes = newRuntimes
	m.mu.Unlock()
	for _, runtime := range newRuntimes {
		runtime.start()
	}
	return nil
}

func (m *RuntimeManager) stopAll() {
	m.mu.Lock()
	old := m.runtimes
	m.runtimes = make(map[string]*managedRuntime)
	m.mu.Unlock()
	for _, runtime := range old {
		runtime.stop()
	}
}

func sameMMA2Structure(a, b Document) bool {
	return equalStrings(structureKeys(a), structureKeys(b))
}

func structureKeys(doc Document) []string {
	keys := make([]string, 0, len(doc.Devices))
	for _, device := range doc.Devices {
		if !device.Enabled {
			continue
		}
		keys = append(keys, fmt.Sprintf("%d/%d/fc%d/%d/%d", device.Destination.Port, device.Destination.UnitID, device.Function, device.Start, device.Count))
	}
	sort.Strings(keys)
	return keys
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cloneDocument(doc Document) Document {
	out := Document{Devices: make([]DeviceDefinition, len(doc.Devices))}
	copy(out.Devices, doc.Devices)
	return out
}
