package simulator

import (
	"net"
	"testing"
	"time"
)

func TestBootRestoreArmsEnabledSchedulesWithoutRestartRequest(t *testing.T) {
	fixture := newRawFixture(t, rawRespOK)
	_, portText, err := net.SplitHostPort(fixture.ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	port := uint16(parsePort(portText, t))

	store := Store{Root: t.TempDir()}
	enabled := validDevice()
	enabled.Name = "boot-restored"
	enabled.MMA2.Port = port
	enabled.RandomRuntime = RandomRuntimeParams{FC1IntervalMS: 5, FC2IntervalMS: 7, FC3IntervalMS: 9, FC4IntervalMS: 11}
	disabled := validDevice()
	disabled.Name = "boot-disabled"
	disabled.Enabled = false
	disabled.MMA2.Port = 15041

	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{enabled, disabled}}); err != nil {
		t.Fatal(err)
	}

	router, applier, err := newRuntimeApplyRouter(store, 500*time.Millisecond)
	if err != nil {
		t.Fatalf("boot restore must not fail when MMA2 is ready: %v", err)
	}
	_ = router
	defer applier.Stop()

	applier.mu.Lock()
	_, armed := applier.schedulers["boot-restored"]
	_, disabledArmed := applier.schedulers["boot-disabled"]
	ready := applier.ready
	applier.mu.Unlock()
	if !ready || !armed {
		t.Fatal("boot must arm enabled schedules after MMA2 readiness")
	}
	if disabledArmed {
		t.Fatal("boot must not arm disabled schedules")
	}

	if _, pending, err := store.LoadRestartRequest(); err != nil {
		t.Fatal(err)
	} else if pending {
		t.Fatal("ordinary unchanged boot must not issue an MMA2 restart request")
	}

	time.Sleep(30 * time.Millisecond)
	status, err := applier.RuntimeStatus("boot-restored")
	if err != nil {
		t.Fatal(err)
	}
	if status.MMA2 != "RUNNING" || status.Device != "RUNNING" || status.RawIngest != "OK" {
		t.Fatalf("boot-restored schedule must report running: %+v", status)
	}
	for _, fc := range []string{"fc1", "fc2", "fc3", "fc4"} {
		if status.FC[fc].Last == "" || status.FC[fc].Last == "Never" || status.FC[fc].Next == "" {
			t.Fatalf("%s timing must advance after boot restore: %+v", fc, status.FC[fc])
		}
	}
}

func TestBootRestoreSurfacesUnavailableMMA2WithoutRestart(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDevice()
	device.Name = "unavailable-mma2"
	device.MMA2.Port = freePort(t)
	device.RandomRuntime = RandomRuntimeParams{FC1IntervalMS: 5, FC2IntervalMS: 5, FC3IntervalMS: 9, FC4IntervalMS: 11}
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}

	router, applier, err := newRuntimeApplyRouter(store, 150*time.Millisecond)
	if err != nil {
		t.Fatalf("unavailable MMA2 must not fail boot restore: %v", err)
	}
	_ = router
	defer applier.Stop()

	applier.mu.Lock()
	n := len(applier.schedulers)
	applier.mu.Unlock()
	if n != 0 {
		t.Fatalf("boot must not arm schedules when MMA2 is unavailable: %d armed", n)
	}
	if _, pending, err := store.LoadRestartRequest(); err != nil {
		t.Fatal(err)
	} else if pending {
		t.Fatal("unavailable boot must not write an MMA2 restart request")
	}

	status, err := applier.RuntimeStatus(device.Name)
	if err != nil {
		t.Fatal(err)
	}
	if status.MMA2 != "STOPPED" || status.RawIngest != "STOPPED" {
		t.Fatalf("unavailable boot must surface truthful state: %+v", status)
	}
	if status.Device != "ERROR" {
		t.Fatalf("enabled but unavailable device must report ERROR: %+v", status)
	}
}
