package replicator

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

// BlockRuntimeStatus is the truthful runtime state for one Pull Block/poller.
type BlockRuntimeStatus struct {
	Index     int    `json:"index"`
	Running   bool   `json:"running"`
	Cycles    uint64 `json:"cycles"`
	Source    string `json:"source_status"`
	LastPoll  string `json:"last_poll,omitempty"`
	LastError string `json:"last_error,omitempty"`
}

// DeviceRuntimeStatus preserves the previous aggregate fields while exposing
// the ordered per-block states required by the Pull Blocks editor.
type DeviceRuntimeStatus struct {
	Name      string               `json:"name"`
	Enabled   bool                 `json:"enabled"`
	Running   bool                 `json:"running"`
	Cycles    uint64               `json:"cycles"`
	Source    string               `json:"source_status"`
	LastPoll  string               `json:"last_poll,omitempty"`
	LastError string               `json:"last_error,omitempty"`
	Blocks    []BlockRuntimeStatus `json:"blocks"`
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

// RuntimeManager owns all UI-configured Replicator poll loops. Each enabled
// Pull Block gets one managedRuntime while destination ownership remains device-level.
type RuntimeManager struct {
	store   Store
	timeout time.Duration

	mu       sync.RWMutex
	document Document
	runtimes map[string][]*managedRuntime
}

func NewRuntimeManager(store Store) *RuntimeManager {
	return &RuntimeManager{store: store, timeout: DefaultApplyTimeout, runtimes: make(map[string][]*managedRuntime)}
}

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

func (m *RuntimeManager) Apply(edited Document) (Document, bool, error) {
	resolved, err := m.store.ResolveDocumentDestinations(edited)
	if err != nil {
		return Document{}, false, err
	}
	if err := m.store.CheckDocumentOwnership(resolved); err != nil {
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
	restorePrevious := func(applyErr error) (Document, bool, error) {
		if restoreErr := m.replaceRuntimes(previousResolved); restoreErr != nil {
			return Document{}, structural, fmt.Errorf("%w; restore previous Replicator runtimes: %v", applyErr, restoreErr)
		}
		return Document{}, structural, applyErr
	}

	resolved, effective, err := m.store.ComposeDocumentDestinations(resolved)
	if err != nil {
		return restorePrevious(err)
	}
	if err := m.store.requestMMA2Restart(effective, documentDestinationPorts(resolved), m.timeout); err != nil {
		return restorePrevious(err)
	}
	if err := m.store.SaveDocument(resolved); err != nil {
		// MMA2 has already accepted the new structure at this point. Do not start
		// old pollers against a potentially different destination map; fail closed.
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

func blockStatus(index int, runtime *managedRuntime) BlockRuntimeStatus {
	status := BlockRuntimeStatus{Index: index, Source: "WAITING"}
	if runtime == nil {
		return status
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
	return status
}

func (m *RuntimeManager) Status(name string) (DeviceRuntimeStatus, error) {
	m.mu.RLock()
	doc := cloneDocument(m.document)
	runtimes := append([]*managedRuntime(nil), m.runtimes[name]...)
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
	blocks := device.blocks()
	status := DeviceRuntimeStatus{Name: name, Enabled: device.Enabled, Source: "WAITING", Blocks: make([]BlockRuntimeStatus, len(blocks))}
	if !device.Enabled {
		status.Source = "DISABLED"
		for i := range status.Blocks {
			status.Blocks[i] = BlockRuntimeStatus{Index: i, Source: "DISABLED"}
		}
		return status, nil
	}
	for i := range blocks {
		var runtime *managedRuntime
		if i < len(runtimes) {
			runtime = runtimes[i]
		}
		bs := blockStatus(i, runtime)
		status.Blocks[i] = bs
		status.Cycles += bs.Cycles
		status.Running = status.Running || bs.Running
		if bs.LastPoll > status.LastPoll {
			status.LastPoll = bs.LastPoll
		}
		if status.LastError == "" && bs.LastError != "" {
			status.LastError = bs.LastError
		}
	}
	switch {
	case len(status.Blocks) == 0:
		status.Source = "WAITING"
	case anyBlockSource(status.Blocks, "ERROR"):
		status.Source = "ERROR"
	case allBlockSource(status.Blocks, "OK"):
		status.Source = "OK"
	default:
		status.Source = "WAITING"
	}
	return status, nil
}

func anyBlockSource(blocks []BlockRuntimeStatus, value string) bool {
	for _, block := range blocks {
		if block.Source == value {
			return true
		}
	}
	return false
}

func allBlockSource(blocks []BlockRuntimeStatus, value string) bool {
	if len(blocks) == 0 {
		return false
	}
	for _, block := range blocks {
		if block.Source != value {
			return false
		}
	}
	return true
}

func (m *RuntimeManager) Stop() {
	m.stopAll()
}

func (m *RuntimeManager) replaceRuntimes(doc Document) error {
	newRuntimes := make(map[string][]*managedRuntime)
	for _, device := range doc.Devices {
		if !device.Enabled {
			continue
		}
		blocks := device.blocks()
		deviceRuntimes := make([]*managedRuntime, 0, len(blocks))
		for i, block := range blocks {
			cfg, err := device.runtimeConfigForBlock(block)
			if err != nil {
				return fmt.Errorf("device %q pull block %d: %w", device.Name, i, err)
			}
			if err := validateCycleMapping(cfg); err != nil {
				return fmt.Errorf("device %q pull block %d: %w", device.Name, i, err)
			}
			deviceRuntimes = append(deviceRuntimes, newManagedRuntime(cfg))
		}
		newRuntimes[device.Name] = deviceRuntimes
	}
	m.mu.Lock()
	m.document = cloneDocument(doc)
	m.runtimes = newRuntimes
	m.mu.Unlock()
	for _, group := range newRuntimes {
		for _, runtime := range group {
			runtime.start()
		}
	}
	return nil
}

func (m *RuntimeManager) stopAll() {
	m.mu.Lock()
	old := m.runtimes
	m.runtimes = make(map[string][]*managedRuntime)
	m.mu.Unlock()
	for _, group := range old {
		for _, runtime := range group {
			runtime.stop()
		}
	}
}

func sameMMA2Structure(a, b Document) bool {
	return equalStrings(structureKeys(a), structureKeys(b))
}

func structureKeys(doc Document) []string {
	keys := make([]string, 0)
	for _, device := range doc.Devices {
		if !device.Enabled {
			continue
		}
		for _, block := range device.blocks() {
			keys = append(keys, fmt.Sprintf("%d/%d/fc%d/%d/%d", device.Destination.Port, device.Destination.UnitID, block.Function, block.Start, block.Count))
		}
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
	for i := range out.Devices {
		out.Devices[i].PullBlocks = append([]PullBlock(nil), doc.Devices[i].PullBlocks...)
	}
	return out
}
