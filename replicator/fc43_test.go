package replicator

import (
	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"os"
	"reflect"
	"testing"
)

func TestFC43DestinationOverridesAndDefaults(t *testing.T) {
	store := advancedStore(t)
	device := advancedDevice(t)
	fields := map[string]interface{}{"product_code": "Replica"}
	device.MMA2Advanced["fc43"] = fields
	resolved, cfg, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(cfg.Listeners[0].Memory[0].Extra["fc43"], fields) {
		t.Fatal("FC43 lost")
	}
	if err := store.SaveDocument(resolved); err != nil {
		t.Fatal(err)
	}
	loaded, err := store.LoadDocument()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(loaded.Devices[0].MMA2Advanced["fc43"], fields) {
		t.Fatal("reload lost identity")
	}
	device.MMA2Advanced["fc43"] = map[string]interface{}{}
	_, cfg, err = store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if err != nil {
		t.Fatal(err)
	}
	if _, exists := cfg.Listeners[0].Memory[0].Extra["fc43"]; exists {
		t.Fatal("defaults re-inherited")
	}
}

func TestInvalidFC43DestinationPreservesYAML(t *testing.T) {
	for _, enabled := range []bool{true, false} {
		store := advancedStore(t)
		device := advancedDevice(t)
		if _, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}}); err != nil {
			t.Fatal(err)
		}
		composer := mma2composer.New(store.Root, ProducerReplicator)
		before, err := os.ReadFile(composer.EffectiveConfigPath())
		if err != nil {
			t.Fatal(err)
		}
		device.Enabled = enabled
		device.MMA2Advanced["fc43"] = map[string]interface{}{"product_code": ""}
		if _, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}}); err == nil {
			t.Fatal("invalid FC43 accepted")
		}
		after, err := os.ReadFile(composer.EffectiveConfigPath())
		if err != nil {
			t.Fatal(err)
		}
		if string(before) != string(after) {
			t.Fatal("invalid FC43 modified YAML")
		}
	}
}

func TestDefaultDestinationPolicyAllowsFC43(t *testing.T) {
	memory, err := destinationMemoryForBlocks(1, []PullBlock{{Function: 3, Start: 0, Count: 1}})
	if err != nil {
		t.Fatal(err)
	}
	if memory.Policy == nil || len(memory.Policy.Rules) != 1 {
		t.Fatal("default policy missing")
	}
	if !reflect.DeepEqual(memory.Policy.Rules[0].AllowFC, []uint8{1, 2, 3, 4, 5, 6, 15, 16, 43}) {
		t.Fatalf("default policy does not allow FC43: %v", memory.Policy.Rules[0].AllowFC)
	}
}
