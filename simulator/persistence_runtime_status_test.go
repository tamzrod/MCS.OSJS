package simulator

import (
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func enabledBool(value bool) *bool { return &value }

// PERSIST-021 self-check: the simulator runtime status carries observational
// persistence health for the operator surface, and it is purely informational.
func TestRuntimeStatusCarriesPersistenceHealth(t *testing.T) {
	device := newStatusDevice("persist-status", true, 1, 1000)
	device.MMA2.Persistence = &mma2composer.Persistence{Enabled: enabledBool(true)}
	applier := newSchedulerApplier(Store{}, Document{Devices: []DeviceDefinition{device}}, false)
	t.Cleanup(applier.Stop)

	status := runtimeStatusOf(t, applier, device.Name)
	if status.Persistence == nil {
		t.Fatal("runtime status must carry persistence health")
	}
	if !status.Persistence.Configured || status.Persistence.Healthy || !status.Persistence.Sealed {
		t.Fatalf("configured-but-unobserved persistence must be sealed and not healthy: %+v", status.Persistence)
	}
}

// A device without persistence reports not-configured persistence health.
func TestRuntimeStatusPersistenceDisabled(t *testing.T) {
	device := newStatusDevice("persist-off", true, 1, 1000)
	applier := newSchedulerApplier(Store{}, Document{Devices: []DeviceDefinition{device}}, false)
	t.Cleanup(applier.Stop)

	status := runtimeStatusOf(t, applier, device.Name)
	if status.Persistence == nil || status.Persistence.Configured {
		t.Fatalf("device without persistence must report not-configured: %+v", status.Persistence)
	}
	if !status.Persistence.Sealed || status.Persistence.Healthy {
		t.Fatalf("not-configured persistence must be sealed and not healthy: %+v", status.Persistence)
	}
}

// The persistence field is observational: reading it never changes device or
// ingest state, and repeated reads are stable.
func TestRuntimeStatusPersistenceIsObservational(t *testing.T) {
	device := newStatusDevice("persist-obs", true, 1, 1000)
	device.MMA2.Persistence = &mma2composer.Persistence{Enabled: enabledBool(true)}
	applier := newSchedulerApplier(Store{}, Document{Devices: []DeviceDefinition{device}}, false)
	t.Cleanup(applier.Stop)

	first := runtimeStatusOf(t, applier, device.Name)
	second := runtimeStatusOf(t, applier, device.Name)
	if first.Persistence == nil || second.Persistence == nil || *first.Persistence != *second.Persistence {
		t.Fatalf("persistence observation must be stable: %+v vs %+v", first.Persistence, second.Persistence)
	}
	if first.Device != second.Device || first.RawIngest != second.RawIngest {
		t.Fatalf("observing persistence must not change device/ingest state: %+v vs %+v", first, second)
	}
}
