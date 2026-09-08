package simulator

import (
	"fmt"
	"reflect"
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
	store      Store
	schedulers map[string]*Scheduler
}

func NewSchedulerApplier(store Store, initial Document) *SchedulerApplier {
	a := &SchedulerApplier{store: store, schedulers: make(map[string]*Scheduler)}
	a.replaceSchedulers(initial)
	return a
}

func (a *SchedulerApplier) replaceSchedulers(doc Document) {
	for _, scheduler := range a.schedulers {
		scheduler.Stop()
	}
	a.schedulers = make(map[string]*Scheduler)
	for _, def := range doc.Devices {
		device := def
		a.schedulers[device.Name] = NewScheduler(device, func(values Values) {
			_ = NewRawIngestClient(device).Send(values)
		})
	}
}

func (a *SchedulerApplier) ApplyStructural(_, edited Document) error {
	if err := a.store.ComposeDocument(edited); err != nil {
		return err
	}
	a.replaceSchedulers(edited)
	return nil
}

func (a *SchedulerApplier) ApplyTiming(previous, edited Document) error {
	if len(previous.Devices) != len(edited.Devices) {
		return fmt.Errorf("timing apply cannot change device inventory")
	}
	for i, def := range edited.Devices {
		scheduler := a.schedulers[previous.Devices[i].Name]
		if scheduler == nil {
			return fmt.Errorf("scheduler for %q is unavailable", previous.Devices[i].Name)
		}
		scheduler.UpdateTiming(def.RandomRuntime)
		if def.Name != previous.Devices[i].Name {
			delete(a.schedulers, previous.Devices[i].Name)
			a.schedulers[def.Name] = scheduler
		}
	}
	return nil
}

func (a *SchedulerApplier) Stop() {
	for _, scheduler := range a.schedulers {
		scheduler.Stop()
	}
}

func NewRuntimeApplyRouter(store Store) (*ApplyRouter, *SchedulerApplier, error) {
	initial, err := store.Load()
	if err != nil {
		return nil, nil, err
	}
	timing := NewSchedulerApplier(store, initial)
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
