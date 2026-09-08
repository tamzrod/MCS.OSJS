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
}

func NewSchedulerApplier(store Store, initial Document) *SchedulerApplier {
	return newSchedulerApplier(store, initial, true)
}

func newSchedulerApplier(store Store, initial Document, ready bool) *SchedulerApplier {
	a := &SchedulerApplier{store: store, ready: ready, restartTimeout: DefaultRestartReadyTimeout, schedulers: make(map[string]*Scheduler), devices: make(map[string]DeviceDefinition), rawErrors: make(map[string]string)}
	a.replaceSchedulers(initial)
	return a
}

func (a *SchedulerApplier) replaceSchedulers(doc Document) {
	a.mu.Lock()
	old := a.schedulers
	a.schedulers = make(map[string]*Scheduler)
	a.devices = make(map[string]DeviceDefinition)
	a.rawErrors = make(map[string]string)
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
	// MMA2 is an independently managed appliance. After a successful shared-config
	// commit the Simulator requests exactly one MMA2 RESTART (SIM-014(, waits for
	// readiness,and reports failure truthfully. It must not start, stop, spawn, kill,
	// or replace MMA2; it never owns the appliance process. Scheduler arming after the
	// observed apply/readiness event belongs to SIM-015..
	cfg, err := a.store.loadEffective()
	if err != nil {
		return fmt.Errorf("load composed MMA2 config: %w", err)
	}
	ports := composedSimulatorPorts(edited)
	if err := a.store.WriteRestartRequest(RestartRequestForConfig(time.Now(), cfg, ports)); err != nil {
		return fmt.Errorf("mma2 restart request failed: %w", err)
	}
	if err := WaitMMA2Ready(ports, a.restartTimeout); err != nil {
		return err
	}
	// Readiness confirmed: the request has been honored; clear it so no stale
	// request survives into a later boot restore (SIM-016(..
	if err := a.store.ClearRestartRequest(); err != nil {
		return err
	}
	return nil
}

// ArmSchedules arms enabled schedules for theat committed document after the
// full Save & Apply chain (shared-config commit -> MMA2 RESTART -> readiness(
// has succeeded. It replaces prior schedules so each successful apply re-arms
// with the latest persisted state (SIM-015(。
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

func (a *SchedulerApplier) RuntimeStatus(name string) (DeviceRuntimeStatus, error) {
	a.mu.Lock()
	device, ok := a.devices[name]
	scheduler := a.schedulers[name]
	rawError := a.rawErrors[name]
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
	if a.ready {
		if !device.Enabled {
			return status, nil
		}
		if rawError != "" {
			status.RawIngest = "ERROR"
			status.Device = "ERROR"
		} else if status.MMA2 == "RUNNING" {
			status.RawIngest = "OK"
			status.Device = "RUNNING"
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
	} else {
		status.RawIngest = "STOPPED"
		if device.Enabled {
			status.Device = "ERROR"
		}
	}
	return status, nil
}

func NewRuntimeApplyRouter(store Store) (*ApplyRouter, *SchedulerApplier, error) {
	initial, err := store.Load()
	if err != nil {
		return nil, nil, err
	}
	// Compose persisted Simulator reservations into the shared MMA2 configuration,
	// but never manage the independently started MMA2 process. Schedulers remain
	// unarmed until the post-apply readiness work introduced by SIM-015.
	if err := store.ComposeDocument(initial); err != nil {
		return nil, nil, err
	}
	timing := newSchedulerApplier(store, initial, false)
	return NewApplyRouter(store, timing, timing), timing, nil
}

type ApplyResult struct {
	Document Document  `json:"document"`
	Path     ApplyPath `json:"apply_path"`
	Message  string    `json:"message"`
}

// ApplyRouter classifies a valid edit before invoking exactly one downstream
// consumer. Persistence happens only after the selected consumer succeeds, so
// validation, ownership, activation, or scheduler errors leave the previously
// active simulator document untouched.
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
	// SIM-015: schedules may arm only after the full chain (commit, restarted,
	// ready,persisted( succeeded. The live SchedulerApplier owns arming; plumbing
	// fakes (e.g. applyRecorder( don't arm schedules..
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
