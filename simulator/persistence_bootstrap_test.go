package simulator

import (
	"bytes"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func TestBootstrapPersistenceTransitionCapturesCurrentMemory(t *testing.T) {
	regs := []byte{0x12, 0x34}
	_, port := newFakeMMA2Endpoint(t, regs, nil)

	before := validDevice()
	before.Name = "bootstrap"
	before.MMA2.Port = port
	before.MMA2.UnitID = 1
	before.MMA2.FC1 = Area{}
	before.MMA2.FC2 = Area{}
	before.MMA2.FC3 = Area{Start: 0, Count: 1}
	before.MMA2.FC4 = Area{}
	before.MMA2.Persistence = nil
	before.MMA2.StateSealing = map[string]interface{}{"enabled": false}

	after := before
	after.MMA2.Persistence = &mma2composer.Persistence{Enabled: boolPtr(true)}
	after.MMA2.StateSealing = map[string]interface{}{"enabled": true, "area": "coil", "address": 0, "exception": 6}

	root := t.TempDir()
	if err := bootstrapPersistenceTransitions(root, Document{Devices: []DeviceDefinition{before}}, Document{Devices: []DeviceDefinition{after}}); err != nil {
		t.Fatal(err)
	}

	key := mma2composer.PersistenceMemoryKey{Port: port, UnitID: 1}
	rules, err := persistenceOwnedRulesForMemory(memoryFromMMA2Params(after.MMA2))
	if err != nil { t.Fatal(err) }
	configs, err := mma2composer.PersistenceSnapshotConfigs(key, rules)
	if err != nil { t.Fatal(err) }
	adapter, err := mma2composer.NewPersistenceFilesystemAdapter(root, configs)
	if err != nil { t.Fatal(err) }
	manifest, got, present, err := adapter.ReadPersistenceSnapshot(key, "holding_registers")
	if err != nil { t.Fatal(err) }
	if !present || !bytes.Equal(got, regs) {
		t.Fatalf("initial snapshot wrong: present=%v got=%v", present, got)
	}
	if err := mma2composer.ValidatePersistenceSnapshotManifest(manifest, key, "holding_registers", mma2composer.PersistenceRegisters, 0, 1, got); err != nil {
		t.Fatalf("initial snapshot manifest invalid: %v", err)
	}
}

func TestBootstrapPersistenceTransitionSkipsAlreadyEnabled(t *testing.T) {
	device := validDevice()
	device.MMA2.Persistence = &mma2composer.Persistence{Enabled: boolPtr(true)}
	if err := bootstrapPersistenceTransitions(t.TempDir(), Document{Devices: []DeviceDefinition{device}}, Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
}

func TestBootstrapPersistenceTransitionRejectsUnsafeCombinedChanges(t *testing.T) {
	before := validDevice()
	before.Name = "unsafe"
	after := before
	after.MMA2.Persistence = &mma2composer.Persistence{Enabled: boolPtr(true)}
	after.MMA2.StateSealing = map[string]interface{}{"enabled": true, "area": "coil", "address": 0, "exception": 6}
	after.MMA2.FC3.Count++

	err := bootstrapPersistenceTransitions(t.TempDir(), Document{Devices: []DeviceDefinition{before}}, Document{Devices: []DeviceDefinition{after}})
	if err == nil {
		t.Fatal("enabling persistence while changing memory layout must fail before restart")
	}
}
