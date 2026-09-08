package simulator

import (
	"errors"
	"testing"
)

type applyRecorder struct {
	structural int
	timing     int
	err        error
}

func (r *applyRecorder) ApplyStructural(_, _ Document) error { r.structural++; return r.err }
func (r *applyRecorder) ApplyTiming(_, _ Document) error     { r.timing++; return r.err }

func TestApplyRouterStructuralTimingAndRejectedPaths(t *testing.T) {
	store := Store{Root: t.TempDir()}
	base := Document{Devices: []DeviceDefinition{validDevice()}}
	if err := store.SaveDocument(base); err != nil {
		t.Fatal(err)
	}

	recorder := &applyRecorder{}
	router := NewApplyRouter(store, recorder, recorder)
	result, err := router.Apply(cloneDocument(base))
	if err != nil {
		t.Fatal(err)
	}
	if result.Path != ApplyNoChange || recorder.structural != 0 || recorder.timing != 0 {
		t.Fatalf("no-change route: result=%+v recorder=%+v", result, recorder)
	}
	structural := cloneDocument(base)
	structural.Devices[0].MMA2.FC3.Count++
	result, err = router.Apply(structural)
	if err != nil {
		t.Fatal(err)
	}
	if result.Path != ApplyStructural || recorder.structural != 1 || recorder.timing != 0 {
		t.Fatalf("structural route: result=%+v recorder=%+v", result, recorder)
	}

	timing := cloneDocument(structural)
	timing.Devices[0].RandomRuntime.FC3IntervalMS++
	result, err = router.Apply(timing)
	if err != nil {
		t.Fatal(err)
	}
	if result.Path != ApplyTiming || recorder.structural != 1 || recorder.timing != 1 {
		t.Fatalf("timing route: result=%+v recorder=%+v", result, recorder)
	}

	recorder.err = errors.New("ownership conflict")
	rejected := cloneDocument(timing)
	rejected.Devices[0].MMA2.Port++
	if _, err := router.Apply(rejected); err == nil {
		t.Fatal("expected downstream rejection")
	}
	active, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if active.Devices[0] != timing.Devices[0] {
		t.Fatalf("rejected apply replaced active state: %+v", active.Devices[0])
	}
	if recorder.structural != 2 || recorder.timing != 1 {
		t.Fatalf("unexpected route counts: %+v", recorder)
	}
}

func TestApplyRouterRejectsInvalidBeforeConsumers(t *testing.T) {
	store := Store{Root: t.TempDir()}
	base := Document{Devices: []DeviceDefinition{validDevice()}}
	if err := store.SaveDocument(base); err != nil {
		t.Fatal(err)
	}
	recorder := &applyRecorder{}
	router := NewApplyRouter(store, recorder, recorder)
	bad := cloneDocument(base)
	bad.Devices[0].MMA2.FC3 = Area{Start: 65530, Count: 20}
	if _, err := router.Apply(bad); err == nil {
		t.Fatal("expected validation error")
	}
	if recorder.structural != 0 || recorder.timing != 0 {
		t.Fatalf("invalid input reached consumers: %+v", recorder)
	}
}

func cloneDocument(doc Document) Document {
	out := Document{Devices: make([]DeviceDefinition, len(doc.Devices))}
	copy(out.Devices, doc.Devices)
	return out
}
