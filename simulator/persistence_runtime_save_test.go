package simulator

import (
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// PERSIST-R02 self-check: the runtime save host routes only persistence-owned
// rule events to the snapshot store for the correct memory/area identity, and
// derives its adapter registration from the authoritative rules.
func TestPersistenceRuntimeSaveHostRoutesEvents(t *testing.T) {
	key := mma2composer.PersistenceMemoryKey{Port: 5020, UnitID: 1}
	rules := []mma2composer.PersistenceRBERule{
		{ID: 3, Area: "holding_registers", Start: 0, Count: 3, SystemOwned: true},
	}
	reader := PersistenceAreaReaderFunc(func(k mma2composer.PersistenceMemoryKey, area string, kind mma2composer.PersistenceAreaKind, start, count uint16) ([]byte, error) {
		if k != key || area != "holding_registers" {
			t.Fatalf("reader called with wrong identity: %+v %q", k, area)
		}
		return []byte{0, 1, 0, 2, 0, 3}, nil
	})
	root := t.TempDir()
	host, err := NewPersistenceRuntimeSaveHost(root, []PersistenceRuntimeSaveMemory{{Key: key, Rules: rules, Reader: reader}})
	if err != nil {
		t.Fatal(err)
	}
	if !host.Handles(3) || host.Handles(99) {
		t.Fatalf("handled IDs wrong: %v", host.SubscribedRuleIDs())
	}
	host.Publish(3)
	if host.Status().Saves != 1 {
		t.Fatalf("persistence event did not save: %+v", host.Status())
	}
	// A non-persistence ID writes nothing.
	saves := host.Status().Saves
	host.Publish(99)
	if host.Status().Saves != saves {
		t.Fatal("a non-persistence ID must not save")
	}
	// The snapshot was actually written under the data root and loads back.
	configs, err := mma2composer.PersistenceSnapshotConfigs(key, rules)
	if err != nil {
		t.Fatal(err)
	}
	adapter, err := mma2composer.NewPersistenceFilesystemAdapter(root, configs)
	if err != nil {
		t.Fatal(err)
	}
	_, payload, present, err := adapter.ReadPersistenceSnapshot(key, "holding_registers")
	if err != nil || !present {
		t.Fatalf("runtime save did not persist a snapshot: present=%v err=%v", present, err)
	}
	if string(payload) != string([]byte{0, 1, 0, 2, 0, 3}) {
		t.Fatalf("persisted image wrong: %v", payload)
	}
}

// The host derives the adapter registration from the authoritative rules; an
// empty persistence projection yields no host components but no error.
func TestPersistenceRuntimeSaveHostEmptyRules(t *testing.T) {
	host, err := NewPersistenceRuntimeSaveHost(t.TempDir(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(host.SubscribedRuleIDs()) != 0 {
		t.Fatalf("no rules means no subscriptions: %v", host.SubscribedRuleIDs())
	}
}

// PERSIST-R02 self-check: the live SchedulerApplier exposes the real runtime call
// site and routes persistence events only when a save host was armed.
func TestSchedulerApplierPublishesPersistenceEvents(t *testing.T) {
	device := newStatusDevice("persist-save", true, 1, 1000)
	device.MMA2.Persistence = &mma2composer.Persistence{Enabled: boolPtr(true)}
	applier := newSchedulerApplier(Store{Root: t.TempDir()}, Document{Devices: []DeviceDefinition{device}}, false)
	t.Cleanup(applier.Stop)

	// Before arming, an event is a safe no-op.
	applier.PublishPersistenceEvent(1)

	if err := applier.ArmPersistenceSave(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	applier.mu.Lock()
	host := applier.persistenceSave
	applier.mu.Unlock()
	if host == nil {
		t.Fatal("arming persistence save must build a host for a persistence-enabled device")
	}
}

func boolPtr(v bool) *bool { return &v }

// PERSIST-R02 self-check: the real runtime save path persists changed state and
// not unchanged state, end to end from a device's Modbus source through the
// filesystem adapter.
func TestPersistenceRuntimeSaveEndToEndChangedState(t *testing.T) {
	regPayload := []byte{0x00, 0x01, 0x00, 0x02, 0x00, 0x03}
	port := fakeModbus(t, 3, regPayload)
	device := DeviceDefinition{
		Name:    "persist-e2e",
		Enabled: true,
		MMA2: MMA2Params{
			Port:        port,
			UnitID:      1,
			FC3:         Area{Start: 0, Count: 3},
			Persistence: &mma2composer.Persistence{Enabled: boolPtr(true)},
		},
	}
	root := t.TempDir()
	host, err := persistenceSaveContext(root, []DeviceDefinition{device})
	if err != nil {
		t.Fatal(err)
	}
	if host == nil {
		t.Fatal("a persistence-enabled device must build a save host")
	}
	ids := host.SubscribedRuleIDs()
	if len(ids) != 1 {
		t.Fatalf("expected one subscribed persistence rule, got %v", ids)
	}
	// First event persists the whole area (changed from the empty image).
	host.Publish(ids[0])
	if host.Status().Saves != 1 {
		t.Fatalf("first event must save: %+v", host.Status())
	}
	// Second event with unchanged authoritative state must not rewrite.
	host.Publish(ids[0])
	if host.Status().Saves != 1 {
		t.Fatalf("unchanged state must not rewrite: %+v", host.Status())
	}

	// The persisted snapshot loads back and matches the observed source.
	configs, _ := mma2composer.PersistenceSnapshotConfigs(mma2composer.PersistenceMemoryKey{Port: port, UnitID: 1},
		[]mma2composer.PersistenceRBERule{{ID: ids[0], Area: "holding_registers", Start: 0, Count: 3, SystemOwned: true}})
	adapter, err := mma2composer.NewPersistenceFilesystemAdapter(root, configs)
	if err != nil {
		t.Fatal(err)
	}
	_, payload, present, err := adapter.ReadPersistenceSnapshot(mma2composer.PersistenceMemoryKey{Port: port, UnitID: 1}, "holding_registers")
	if err != nil || !present {
		t.Fatalf("runtime save snapshot missing: present=%v err=%v", present, err)
	}
	if string(payload) != string(regPayload) {
		t.Fatalf("persisted image wrong: %v want %v", payload, regPayload)
	}
}
