package replicator

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"gopkg.in/yaml.v3"
)

func advancedDevice(t *testing.T) DeviceDefinition {
	t.Helper()
	device := validDeviceDefinition("advanced")
	device.PullBlocks = []PullBlock{{Function: 1, Start: 0, Count: 16, ScanRateMS: 1000}}
	device.PullBlock = device.PullBlocks[0]
	device.Destination = DestinationSelection{Port: 5502, UnitID: 1}
	if err := json.Unmarshal([]byte(`{"policy":{"rules":[{"id":"read","source_ip":["192.168.1.1","2001:db8::/32"],"allow_fc":[1]}]},"state_sealing":{"enabled":true,"area":"coil","address":0,"exception":6},"rbe":{"coils":[{"id":1,"name":"change","start":1,"count":1}]}}`), &device.MMA2Advanced); err != nil {
		t.Fatal(err)
	}
	return device
}

func advancedStore(t *testing.T) Store {
	t.Helper()
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, ProducerReplicator)
	cfg := mma2composer.EffectiveConfig{Extra: map[string]interface{}{"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:19001"}}}}
	owners, err := composer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if err := composer.Commit(cfg, owners); err != nil {
		t.Fatal(err)
	}
	return store
}

func TestAdvancedSettingsPersistComposeAndInherit(t *testing.T) {
	store := advancedStore(t)
	device := advancedDevice(t)
	resolved, effective, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if err != nil {
		t.Fatal(err)
	}
	memory := effective.Listeners[0].Memory[0]
	if memory.Policy.Rules[0].SourceIP[1] != "2001:db8::/32" || memory.Extra["rbe"] == nil || memory.Extra["state_sealing"] == nil {
		t.Fatalf("lost advanced settings: %+v", memory)
	}
	if err := store.SaveDocument(resolved); err != nil {
		t.Fatal(err)
	}
	loaded, err := store.LoadDocument()
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Devices[0].MMA2Advanced["rbe"] == nil {
		t.Fatal("reload lost RBE")
	}
	wire, err := json.Marshal(loaded)
	if err != nil {
		t.Fatal(err)
	}
	var roundTrip Document
	if err := json.Unmarshal(wire, &roundTrip); err != nil {
		t.Fatal(err)
	}
	if roundTrip.Devices[0].MMA2Advanced["policy"] == nil {
		t.Fatal("IPC lost policy")
	}
	roundTrip.Devices[0].MMA2Advanced = nil
	inherited, _, err := store.ComposeDocumentDestinations(roundTrip)
	if err != nil {
		t.Fatal(err)
	}
	if inherited.Devices[0].MMA2Advanced["rbe"] == nil {
		t.Fatal("legacy save lost RBE")
	}
	inherited.Devices[0].MMA2Advanced["rbe"] = nil
	_, cleared, err := store.ComposeDocumentDestinations(inherited)
	if err != nil {
		t.Fatal(err)
	}
	if cleared.Listeners[0].Memory[0].Extra["rbe"] != nil {
		t.Fatal("explicit removal was ignored")
	}
}

func TestInvalidAdvancedSettingsDoNotReplaceEffectiveConfig(t *testing.T) {
	store := advancedStore(t)
	device := advancedDevice(t)
	if _, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	composer := mma2composer.New(store.Root, ProducerReplicator)
	before, err := composer.LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	expected, err := yaml.Marshal(before)
	if err != nil {
		t.Fatal(err)
	}
	device.MMA2Advanced["state_sealing"].(map[string]interface{})["address"] = 999
	if _, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}}); err == nil {
		t.Fatal("accepted out-of-range sealing")
	}
	after, err := composer.LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	actual, err := yaml.Marshal(after)
	if err != nil {
		t.Fatal(err)
	}
	if string(actual) != string(expected) {
		t.Fatal("invalid settings changed effective config")
	}
	if _, err := os.Stat(store.DocumentPath()); !os.IsNotExist(err) {
		t.Fatal("composition unexpectedly wrote document")
	}
}

func TestAdvancedSettingsCloneIsIndependent(t *testing.T) {
	original := Document{Devices: []DeviceDefinition{advancedDevice(t)}}
	copied := cloneDocument(original)
	copied.Devices[0].MMA2Advanced["state_sealing"].(map[string]interface{})["enabled"] = false
	if original.Devices[0].MMA2Advanced["state_sealing"].(map[string]interface{})["enabled"] != true {
		t.Fatal("shared mutable settings")
	}
}
