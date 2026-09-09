package simulator

import (
	"net"
	"testing"
	"time"
)

func newStatusDevice(name string, enabled bool, port uint16, intervalMS uint32) DeviceDefinition {
	d := validDevice()
	d.Name = name
	d.Enabled = enabled
	d.MMA2.Port = port
	d.MMA2.FC1 = Area{}
	d.MMA2.FC2 = Area{}
	d.MMA2.FC4 = Area{}
	d.MMA2.FC3 = Area{Start: 0, Count: 1}
	d.RandomRuntime = RandomRuntimeParams{FC3IntervalMS: intervalMS}
	return d
}

func rawFixturePort(t *testing.T, fixture *rawFixture) uint16 {
	t.Helper()
	_, portText, err := net.SplitHostPort(fixture.ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	return uint16(parsePort(portText, t))
}

func runtimeStatusOf(t *testing.T, a *SchedulerApplier, name string) DeviceRuntimeStatus {
	t.Helper()
	s, err := a.RuntimeStatus(name)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func waitForDeviceState(t *testing.T, a *SchedulerApplier, name, want string) DeviceRuntimeStatus {
	t.Helper()
	deadline := time.Now().Add(3 * time.Second)
	var last DeviceRuntimeStatus
	for time.Now().Before(deadline) {
		last = runtimeStatusOf(t, a, name)
		if last.Device == want {
			return last
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("device %q never reached %q: %+v", name, want, last)
	return DeviceRuntimeStatus{}
}

func statusScheduler(t *testing.T, a *SchedulerApplier, name string) *Scheduler {
	t.Helper()
	a.mu.Lock()
	defer a.mu.Unlock()
	sched := a.schedulers[name]
	if sched == nil {
		t.Fatalf("enabled device %q must have an armed schedule", name)
	}
	return sched
}

func TestRuntimeStatusDisabledAndUnarmedDevicesAreNotRunning(t *testing.T) {
	disabled := newStatusDevice("disabled", false, 1, 1000)
	disabledApplier := newSchedulerApplier(Store{}, Document{Devices: []DeviceDefinition{disabled}}, true)
	t.Cleanup(disabledApplier.Stop)
	if got := runtimeStatusOf(t, disabledApplier, disabled.Name).Device; got != "STOPPED" {
		t.Fatalf("disabled device status = %q, want STOPPED", got)
	}

	unarmed := newStatusDevice("unarmed", true, 1, 1000)
	unarmedApplier := newSchedulerApplier(Store{}, Document{Devices: []DeviceDefinition{unarmed}}, false)
	t.Cleanup(unarmedApplier.Stop)
	status := runtimeStatusOf(t, unarmedApplier, unarmed.Name)
	if status.Device != "WAITING" || status.RawIngest != "STOPPED" {
		t.Fatalf("unarmed device status = %+v, want WAITING with stopped ingest", status)
	}
}

func TestRuntimeStatusRequiresAcceptedRawIngestAndRecoversFromFailure(t *testing.T) {
	fixture := newRawFixture(t, 0x21, rawRespOK)
	device := newStatusDevice("runtime-evidence", true, rawFixturePort(t, fixture), 100)
	applier := NewSchedulerApplier(Store{}, Document{Devices: []DeviceDefinition{device}})
	t.Cleanup(applier.Stop)

	initial := runtimeStatusOf(t, applier, device.Name)
	if initial.MMA2 != "RUNNING" || initial.Device != "WAITING" {
		t.Fatalf("before accepted ingest status = %+v, want MMA2 RUNNING and Simulator WAITING", initial)
	}

	failed := waitForDeviceState(t, applier, device.Name, "ERROR")
	if failed.RawIngest != "ERROR" || failed.RawError == "" {
		t.Fatalf("failed ingest status = %+v, want retained diagnostic", failed)
	}

	recovered := waitForDeviceState(t, applier, device.Name, "RUNNING")
	if recovered.MMA2 != "RUNNING" || recovered.RawIngest != "OK" || recovered.RawError != "" {
		t.Fatalf("recovered ingest status = %+v, want clean RUNNING evidence", recovered)
	}
}

func TestRuntimeStatusMMA2UnavailabilityPreventsRunning(t *testing.T) {
	fixture := newRawFixture(t, rawRespOK)
	device := newStatusDevice("mma2-loss", true, rawFixturePort(t, fixture), 25)
	applier := NewSchedulerApplier(Store{}, Document{Devices: []DeviceDefinition{device}})
	t.Cleanup(applier.Stop)

	_ = waitForDeviceState(t, applier, device.Name, "RUNNING")
	if err := fixture.ln.Close(); err != nil {
		t.Fatal(err)
	}
	status := runtimeStatusOf(t, applier, device.Name)
	if status.MMA2 == "RUNNING" || status.Device == "RUNNING" {
		t.Fatalf("status after MMA2 listener loss = %+v, must not remain RUNNING", status)
	}
}
