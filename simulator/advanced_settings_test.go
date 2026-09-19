package simulator

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func TestAdvancedSettingsRoundTripAndCompose(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDevice()
	device.MMA2.Policy = &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{ID: "read", SourceIP: []string{"0.0.0.0/0", "::/0"}, AllowFC: []uint8{1, 3}}}}
	device.MMA2.Policy.Extra = map[string]interface{}{"custom_policy": "keep"}
	device.MMA2.Extra = map[string]interface{}{"custom_memory": "keep"}
	device.MMA2.StateSealing = map[string]interface{}{"enabled": false, "area": "coil", "address": 0}
	device.MMA2.RBE = map[string]interface{}{"coils": []interface{}{map[string]interface{}{"id": 1, "name": "watched", "start": 0, "count": 1}}}
	encoded, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	var roundTrip DeviceDefinition
	if err := json.Unmarshal(encoded, &roundTrip); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(roundTrip, device) {
		t.Fatalf("JSON round trip changed advanced settings: %s", encoded)
	}
	composer := store.composer()
	cfg := EffectiveMMA2Config{Extra: map[string]interface{}{"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:9001"}}, "custom_root": "keep"}}
	if err := composer.Commit(cfg, OwnershipDoc{}); err != nil {
		t.Fatal(err)
	}
	if err := store.SaveAndCompose(device); err != nil {
		t.Fatal(err)
	}
	doc, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(doc.Devices[0].MMA2.Policy, device.MMA2.Policy) {
		t.Fatal("policy lost on reload")
	}
	if doc.Devices[0].MMA2.StateSealing["enabled"] != false || doc.Devices[0].MMA2.RBE == nil {
		t.Fatal("advanced fields lost")
	}
	legacy := validDevice()
	legacy.MMA2.FC3.Count++
	if err := store.SaveAndCompose(legacy); err != nil {
		t.Fatal(err)
	}
	loaded, err := composer.LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Extra["custom_root"] != "keep" {
		t.Fatal("shared settings lost")
	}
	memory := loaded.Listeners[0].Memory[0]
	if !reflect.DeepEqual(memory.Policy, device.MMA2.Policy) || memory.Extra["rbe"] == nil || memory.Extra["state_sealing"] == nil {
		t.Fatal("legacy save lost advanced settings")
	}
	before, _ := os.ReadFile(store.EffectiveConfigPath())
	deviceBefore, _ := os.ReadFile(store.DevicesPath())
	device.MMA2.StateSealing = map[string]interface{}{"enabled": true, "area": "coil", "address": 65535}
	if err := store.SaveAndCompose(device); err == nil {
		t.Fatal("invalid sealing accepted")
	}
	after, _ := os.ReadFile(store.EffectiveConfigPath())
	if string(before) != string(after) {
		t.Fatal("invalid save changed effective config")
	}
	deviceAfter, _ := os.ReadFile(store.DevicesPath())
	if string(deviceBefore) != string(deviceAfter) {
		t.Fatal("invalid save changed device document")
	}
}
