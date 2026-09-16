package simulator

import (
	"fmt"
	"net"
	"reflect"
	"sync"
	"time"
)

type ApplyPath string

const (
	ApplyNoChange   ApplyPath = "no-change"
	ApplyStructural ApplyPath = "mma2-structural"
	ApplyTiming     ApplyPath = "random-runtime"
)

type StructuralApplier interface {
	ApplyStructural(previous, edited Document) error
}

type TimingApplier interface {
	ApplyTiming(previous, edited Document) error
}

// SchedulerApplier is the live SIM-003 timing consumer. It lazily creates one
// scheduler per persisted device and applies timing-only edits without touching
// MMA2 configuration. Generated values continue through SIM-004 raw ingest.
type SchedulerApplier struct {
	mu             sync.Mutex
	store          Store
	ready          bool
	restartTimeout time.Duration
	schedulers     map[string]*Scheduler
	devices        map[string]DeviceDefinition
	rawErrors      map[string]string
	ingestOK       map[string]bool
}

func NewSchedulerApplier(store Store, initial Document) *SchedulerApplier {
	return newSchedulerApplier(store, initial, true)
}

func newSchedulerApplier(store Store, initial Document, ready bool) *SchedulerApplier {
	a := &SchedulerApplier{store: store, ready: ready, restartTimeout: DefaultRestartReadyTimeout, schedulers: make(map[string]*Scheduler), devices: make(map[string]DeviceDefinition), rawErrors: make(map[string]string), ingestOK: make(map[string]bool)}
	a.replaceSchedulers(initial)
	return a
}

func (a *SchedulerApplier) replaceSchedulers(doc Document) {
	a.mu.Lock()
	old := a.schedulers
	a.schedulers = make(map[string]*Scheduler)
	a.devices = make(map[string]DeviceDefinition)
	a.rawErrors = make(map[string]string)
	a.ingestOK = make(map[string]bool)
	a.mu.Unlock()
	for _, scheduler := range old {
		scheduler.Stop()
	}
	for _, def := range doc.Devices {
		device := def
		a.mu.Lock()
		a.devices[device.Name] = device
		a.mu.Unlock()
		if !a.ready || !device.Enabled {
			continue
		}
		scheduler := NewScheduler(device, func(values Values) {
			err := NewRawIngestClient(device).Send(values)
			a.mu.Lock()
			if err == nil {
				delete(a.rawErrors, device.Name)
				a.ingestOK[device.Name] = true
			} else {
				a.rawErrors[device.Name] = err.Error()
			}
			a.mu.Unlock()
		})
		a.mu.Lock()
		a.schedulers[device.Name] = scheduler
		a.mu.Unlock()
	}
}

func (a *SchedulerApplier) ApplyStructural(_, edited Document) error {
	if err := a.store.ComposeDocument(edited); err != nil {
		return err
	}
	// MMA2 is independently managed. Committed shared-config structural changes
	// request one restart and wait for acknowledgment/readiness before arming.
	cfg, err := a.store.loadEffective()
	if err != nil {
		return fmt.Errorf("load composed MMA2 config: %w", err)
	}
	ports := composedSimulatorPorts(edited)
	restartRequest := RestartRequestForConfig(time.Now(), cfg, ports)
	if err := a.store.WriteRestartRequest(restartRequest); err != nil {
		return fmt.Errorf("mma2 restart request failed: %w", err)
	}
	if err := a.store.WaitRestartAcknowledged(restartRequest.ConfigSHA256, a.restartTimeout); err != nil {
		return err
	}
	if err := WaitMMA2Ready(ports, a.restartTimeout); err != nil {
		return err
	}
	if err := a.store.ClearRestartRequest(); err != nil {
		return err
	}
	return nil
}

// ArmSchedules arms schedules after the committed shared-config apply/restart
// readiness chain has succeeded. MMA2 itself is never owned by this applier.
func (a *SchedulerApplier) ArmSchedules(doc Document) {
	a.mu.Lock()
	a.ready = true
	a.mu.Unlock()
	a.replaceSchedulers(doc)
}

func (a *SchedulerApplier) ApplyTiming(previous, edited Document) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(previous.Devices) != len(edited.Devices) {
		return fmt.Errorf("timing apply cannot change device inventory")
	}
	for i, def := range edited.Devices {
		scheduler := a.schedulers[previous.Devices[i].Name]
		if scheduler == nil {
			return fmt.Errorf("scheduler for %q is unavailable", previous.Devices[i].Name)
		}
		scheduler.UpdateTiming(def.RandomRuntime)
		a.devices[def.Name] = def
		if def.Name != previous.Devices[i].Name {
			delete(a.schedulers, previous.Devices[i].Name)
			a.schedulers[def.Name] = scheduler
		}
	}
	return nil
}

func (a *SchedulerApplier) Stop() {
	a.mu.Lock()
	schedulers := a.schedulers
	a.schedulers = make(map[string]*Scheduler)
	a.mu.Unlock()
	for _, scheduler := range schedulers {
		scheduler.Stop()
	}
}

type FCRuntimeStatus struct {
	Last string `json:"last"`
	Next string `json:"next"`
}

type DeviceRuntimeStatus struct {
	Name        string                     `json:"name"`
	Device      string                     `json:"device_status"`
	MMA2        string                     `json:"mma2_status"`
	RawIngest   string                     `json:"raw_ingest_status"`
	RawError    string                     `json:"raw_ingest_error,omitempty"`
	TotalPoints uint32                     `json:"total_points"`
	FC          map[string]FCRuntimeStatus `json:"fc"`
}

// needsRandomIngest is true only if an allocated area has positive timing.
// A zero interval allocates/serves the same MMA2 area, but has no generator.
func needsRandomIngest(device DeviceDefinition) bool {
	return (device.MMA2.FC1.Count > 0 && device.RandomRuntime.FC1IntervalMS > 0) ||
		(device.MMA2.FC2.Count > 0 && device.RandomRuntime.FC2IntervalMS > 0) ||
		(device.MMA2.FC3.Count > 0 && device.RandomRuntime.FC3IntervalMS > 0) ||
		(device.MMA2.FC4.Count > 0 && device.RandomRuntime.FC4IntervalMS > 0)
}

func (a *SchedulerApplier) RuntimeStatus(name string) (DeviceRuntimeStatus, error) {
	a.mu.Lock()
	device, ok := a.devices[name]
	scheduler := a.schedulers[name]
	rawError := a.rawErrors[name]
	ready := a.ready
	ingestOK := a.ingestOK[name]
	a.mu.Unlock()
	if !ok {
		return DeviceRuntimeStatus{}, fmt.Errorf("device %q not found", name)
	}
	status := DeviceRuntimeStatus{Name: name, Device: "STOPPED", MMA2: "STOPPED", RawIngest: "WAITING", RawError: rawError, FC: make(map[string]FCRuntimeStatus)}
	status.TotalPoints = uint32(device.MMA2.FC1.Count) + uint32(device.MMA2.FC2.Count) + uint32(device.MMA2.FC3.Count) + uint32(device.MMA2.FC4.Count)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", fmt.Sprint(device.MMA2.Port)), 150*time.Millisecond)
	if err == nil {
		status.MMA2 = "RUNNING"
		_ = conn.Close()
	}
	if !ready {
		status.RawIngest = "STOPPED"
		if device.Enabled {
			status.Device = "WAITING"
		}
		return status, nil
	}
	if !device.Enabled {
		return status, nil
	}
	if !needsRandomIngest(device) {
		status.RawIngest = "NOT REQUIRED"
		status.RawError = ""
		if status.MMA2 == "RUNNING" {
			status.Device = "IDLE"
		} else {
			status.Device = "WAITING"
		}
		return status, nil
	}
	switch {
	case rawError != "":
		status.RawIngest = "ERROR"
		status.Device = "ERROR"
	case status.MMA2 != "RUNNING":
		status.Device = "WAITING"
	case ingestOK:
		status.RawIngest = "OK"
		status.Device = "RUNNING"
	default:
		status.Device = "WAITING"
	}
	if scheduler != nil {
		for fc, timing := range scheduler.Timing() {
			last := "Never"
			if !timing.Last.IsZero() {
				last = timing.Last.Format(time.RFC3339Nano)
			}
			status.FC[fmt.Sprintf("fc%d", fc)] = FCRuntimeStatus{Last: last, Next: timing.Next.Format(time.RFC3339Nano)}
		}
	}
	return status, nil
}

// NewRuntimeApplyRouter restores persisted Simulator reservations on boot;
// it never starts/stops MMA2 or requests a restart merely for ordinary boot.
func NewRuntimeApplyRouter(store Store) (*ApplyRouter, *SchedulerApplier, error) {
	return newRuntimeApplyRouter(store, DefaultRestartReadyTimeout)
}

func newRuntimeApplyRouter(store Store, bootTimeout time.Duration) (*ApplyRouter, *SchedulerApplier, error) {
	initial, err := store.Load()
	if err != nil {
		return nil, nil, err
	}
	if err := store.ComposeDocument(initial); err != nil {
		return nil, nil, err
	}
	timing := newSchedulerApplier(store, initial, false)
	ports := composedSimulatorPorts(initial)
	if len(ports) > 0 {
		if err := WaitMMA2Ready(ports, bootTimeout); err == nil {
			timing.ArmSchedules(initial)
		}
	}
	return NewApplyRouter(store, timing, timing), timing, nil
}

type ApplyResult struct {
	Document Document  `json:"document"`
	Path     ApplyPath `json:"apply_path"`
	Message  string    `json:"message"`
}

// ApplyRouter validates, classifies, applies then persists an edited document.
// Failed validation or downstream apply leaves the prior document unchanged.
type ApplyRouter struct {
	store      Store
	structural StructuralApplier
	timing     TimingApplier
}

func NewApplyRouter(store Store, structural StructuralApplier, timing TimingApplier) *ApplyRouter {
	return &ApplyRouter{store: store, structural: structural, timing: timing}
}

func (r *ApplyRouter) Apply(edited Document) (ApplyResult, error) {
	for i := range edited.Devices {
		if err := ValidateDevice(edited.Devices[i]); err != nil {
			return ApplyResult{}, fmt.Errorf("device %d: %w", i, err)
		}
	}
	previous, err := r.store.Load()
	if err != nil {
		return ApplyResult{}, err
	}
	path := classifyApply(previous, edited)
	switch path {
	case ApplyStructural:
		if r.structural == nil {
			return ApplyResult{}, fmt.Errorf("MMA2 structural apply path is unavailable")
		}
		if err := r.structural.ApplyStructural(previous, edited); err != nil {
			return ApplyResult{}, fmt.Errorf("MMA2 structural apply failed: %w", err)
		}
	case ApplyTiming:
		if r.timing == nil {
			return ApplyResult{}, fmt.Errorf("random-runtime apply path is unavailable")
		}
		if err := r.timing.ApplyTiming(previous, edited); err != nil {
			return ApplyResult{}, fmt.Errorf("random-runtime apply failed: %w", err)
		}
	}
	if err := r.store.SaveDocument(edited); err != nil {
		return ApplyResult{}, err
	}
	if path == ApplyStructural {
		if armer, ok := r.structural.(*SchedulerApplier); ok {
			armer.ArmSchedules(edited)
		}
	}
	return ApplyResult{Document: edited, Path: path, Message: applyMessage(path)}, nil
}

func classifyApply(previous, edited Document) ApplyPath {
	if len(previous.Devices) != len(edited.Devices) {
		return ApplyStructural
	}
	timingChanged := false
	for i := range edited.Devices {
		before, after := previous.Devices[i], edited.Devices[i]
		if before.Name != after.Name || before.Enabled != after.Enabled || !reflect.DeepEqual(before.MMA2, after.MMA2) {
			return ApplyStructural
		}
		if !reflect.DeepEqual(before.RandomRuntime, after.RandomRuntime) {
			timingChanged = true
		}
	}
	if timingChanged {
		return ApplyTiming
	}
	return ApplyNoChange
}

func applyMessage(path ApplyPath) string {
	switch path {
	case ApplyStructural:
		return "MMA2 structural changes applied through the restart/reload path."
	case ApplyTiming:
		return "Random-runtime timing updated without restarting MMA2."
	default:
		return "No runtime changes were required."
	}
}
