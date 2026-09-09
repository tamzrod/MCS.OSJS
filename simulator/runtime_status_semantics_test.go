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
