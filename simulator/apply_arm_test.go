package simulator

import (
	"net"
	"testing"
	"time"
)

func TestApplyArmsEnabledSchedulesAfterStructuralApplyReadiness(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	_, portText, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	store := Store{Root: t.TempDir()}
	enabled := validDevice()
	enabled.Name = "armed-device"
	enabled.MMA2.Port = uint16(parsePort(portText, t))
	enabled.RandomRuntime = RandomRuntimeParams{FC1IntervalMS: 5, FC2IntervalMS: 7, FC3IntervalMS: 9, FC4IntervalMS: 11}
	disabled := validDevice()
	disabled.Name = "disabled-device"
	disabled.Enabled = false
	disabled.MMA2.Port = 15041

	applier := newSchedulerApplier(store, Document{}, false)
	defer applier.Stop()
	router := NewApplyRouter(store, applier, applier)

	edited := Document{Devices: []DeviceDefinition{enabled, disabled}}
	if _, err := router.Apply(edited); err != nil {
		t.Fatalf("apply after ready listener should succeed: %v", err)
	}

	applier.mu.Lock()
	_, armed := applier.schedulers["armed-device"]
	_, disabledArmed := applier.schedulers["disabled-device"]
	applier.mu.Unlock()
	if !armed {
		t.Fatal("enabled schedule must be armed after successful apply/readiness")
	}
	if disabledArmed {
		t.Fatal("disabled device must not arm a schedule")
	}

	time.Sleep(30 * time.Millisecond)
	status, err := applier.RuntimeStatus("armed-device")
	if err != nil {
		t.Fatal(err)
	}
	if status.MMA2 != "RUNNING" || status.Device != "RUNNING" || status.RawIngest != "OK" {
		t.Fatalf("armed schedule must report running: %+v", status)
	}
	for _, fc := range []string{"fc1", "fc2", "fc3", "fc4"} {
		if status.FC[fc].Last == "" || status.FC[fc].Last == "Never" || status.FC[fc].Next == "" {
			t.Fatalf("%s timing must advance after arming: %+v", fc, status.FC[fc])
		}
	}
}

func TestApplyStructuralRestartFailureDoesNotArmSchedules(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDevice()
	device.Name = "never-armed"
	device.MMA2.Port = freePort(t)
	device.RandomRuntime = RandomRuntimeParams{FC1IntervalMS: 5, FC2IntervalMS: 5, FC3IntervalMS: 9, FC4IntervalMS: 11}

	applier := newSchedulerApplier(store, Document{}, false)
	applier.restartTimeout = 200 * time.Millisecond
	defer applier.Stop()
	router := NewApplyRouter(store, applier, applier)

	edited := Document{Devices: []DeviceDefinition{device}}
	if _, err := router.Apply(edited); err == nil {
		t.Fatal("expected restart/readiness failure")
	}

	applier.mu.Lock()
	n := len(applier.schedulers)
	applier.mu.Unlock()
	if n != 0 {
		t.Fatalf("failed apply must not arm schedules: %d armed", n)
	}
}
