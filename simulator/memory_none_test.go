package simulator

import (
	"net"
	"testing"
	"time"
)

func TestNoneSimulationKeepsAllAreasAllocatedWithoutGeneration(t *testing.T) {
	device := validDevice()
	device.RandomRuntime = RandomRuntimeParams{}
	if err := ValidateDevice(device); err != nil {
		t.Fatalf("allocated areas with None should validate: %v", err)
	}
	store := Store{Root: t.TempDir()}
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatalf("saving None memory: %v", err)
	}
	loaded, err := store.Load()
	if err != nil || len(loaded.Devices) != 1 {
		t.Fatalf("reload None memory: document=%+v err=%v", loaded, err)
	}
	if loaded.Devices[0].MMA2 != device.MMA2 || loaded.Devices[0].RandomRuntime != (RandomRuntimeParams{}) {
		t.Fatalf("None changed memory structure or intervals: %+v", loaded.Devices[0])
	}

	updates := make(chan Values, 4)
	scheduler := NewScheduler(device, func(v Values) { updates <- v })
	defer scheduler.Stop()
	if len(scheduler.Timing()) != 0 {
		t.Fatalf("None must have no scheduled FC: %+v", scheduler.Timing())
	}
	select {
	case got := <-updates:
		t.Fatalf("None generated a value: %+v", got)
	case <-time.After(30 * time.Millisecond):
	}
}

func TestMixedNoneAndRandomOnlySchedulesRandomArea(t *testing.T) {
	device := validDevice()
	device.RandomRuntime = RandomRuntimeParams{FC3IntervalMS: 5}
	updates := make(chan Values, 16)
	scheduler := NewScheduler(device, func(v Values) { updates <- v })
	defer scheduler.Stop()
	if timing := scheduler.Timing(); len(timing) != 1 || timing[FC3].Next.IsZero() {
		t.Fatalf("mixed schedule must contain only FC3: %+v", timing)
	}
	select {
	case got := <-updates:
		if got.FC != FC3 {
			t.Fatalf("None area generated a value: %+v", got)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Random FC3 did not generate")
	}
}

func TestNoneRuntimeStatusSeparatesMemoryAndGeneratorHealth(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := uint16(listener.Addr().(*net.TCPAddr).Port)
	device := validDevice()
	device.Name = "none-status"
	device.MMA2.Port = port
	device.RandomRuntime = RandomRuntimeParams{}
	applier := NewSchedulerApplier(Store{Root: t.TempDir()}, Document{Devices: []DeviceDefinition{device}})
	defer applier.Stop()
	status, err := applier.RuntimeStatus(device.Name)
	if err != nil {
		t.Fatal(err)
	}
	if status.MMA2 != "RUNNING" || status.Device != "IDLE" || status.RawIngest != "NOT REQUIRED" || len(status.FC) != 0 {
		t.Fatalf("all-None status incorrectly awaits simulation: %+v", status)
	}

	_ = listener.Close()
	status, err = applier.RuntimeStatus(device.Name)
	if err != nil {
		t.Fatal(err)
	}
	if status.Device == "RUNNING" || status.Device == "IDLE" || status.MMA2 != "STOPPED" {
		t.Fatalf("unavailable MMA2 must not report available: %+v", status)
	}
}

func TestDisabledNoneAndMixedWaitingStatus(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	port := uint16(listener.Addr().(*net.TCPAddr).Port)
	device := validDevice()
	device.Name = "disabled-none"
	device.MMA2.Port = port
	device.RandomRuntime = RandomRuntimeParams{}
	device.Enabled = false
	applier := NewSchedulerApplier(Store{Root: t.TempDir()}, Document{Devices: []DeviceDefinition{device}})
	defer applier.Stop()
	status, err := applier.RuntimeStatus(device.Name)
	if err != nil {
		t.Fatal(err)
	}
	if status.Device != "STOPPED" {
		t.Fatalf("disabled device reported running: %+v", status)
	}

	mixed := device
	mixed.Name = "mixed-status"
	mixed.Enabled = true
	mixed.RandomRuntime.FC3IntervalMS = 100000
	mixedApplier := NewSchedulerApplier(Store{Root: t.TempDir()}, Document{Devices: []DeviceDefinition{mixed}})
	defer mixedApplier.Stop()
	mixedStatus, err := mixedApplier.RuntimeStatus(mixed.Name)
	if err != nil {
		t.Fatal(err)
	}
	if mixedStatus.Device != "WAITING" || mixedStatus.RawIngest != "WAITING" {
		t.Fatalf("active random area must wait for first ingest: %+v", mixedStatus)
	}
}
