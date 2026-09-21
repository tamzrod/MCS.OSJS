package simulator

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// TestAdvancedProjectionPresence verifies that omitted versus explicit null, false, and empty
// advanced values are represented without changing apply behavior.
func TestAdvancedProjectionPresence(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDevice()
	device.MMA2.Policy = &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{ID: "read", SourceIP: []string{"0.0.0.0/0", "::/0"}, AllowFC: []uint8{1, 3}}}}
	device.MMA2.Extra = map[string]interface{}{"custom_policy": "keep"}
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

	var projected map[string]interface{}
	if err := json.Unmarshal(encoded, &projected); err != nil {
		t.Fatal(err)
	}
	projectedMMA2, ok := projected["mma2"].(map[string]interface{})
	if !ok {
		t.Fatalf("mma2 projection has type %T", projected["mma2"])
	}
	sealing, ok := projectedMMA2["state_sealing"].(map[string]interface{})
	if !ok || sealing["enabled"] != false {
		t.Fatalf("explicit false sealing was not projected: %#v", projectedMMA2["state_sealing"])
	}

	omitted := validDevice()
	omittedJSON, err := json.Marshal(omitted)
	if err != nil {
		t.Fatal(err)
	}
	var omittedProjection map[string]interface{}
	if err := json.Unmarshal(omittedJSON, &omittedProjection); err != nil {
		t.Fatal(err)
	}
	omittedMMA2, ok := omittedProjection["mma2"].(map[string]interface{})
	if !ok {
		t.Fatalf("omitted mma2 projection has type %T", omittedProjection["mma2"])
	}
	for _, key := range []string{"policy", "state_sealing", "rbe"} {
		if _, exists := omittedMMA2[key]; exists {
			t.Fatalf("omitted %s was projected: %#v", key, omittedMMA2[key])
		}
	}

	nullProjection := omittedProjection
	nullMMA2 := nullProjection["mma2"].(map[string]interface{})
	nullMMA2["policy"] = nil
	nullMMA2["state_sealing"] = nil
	nullMMA2["rbe"] = nil
	nullJSON, err := json.Marshal(nullProjection)
	if err != nil {
		t.Fatal(err)
	}
	var fromNull DeviceDefinition
	if err := json.Unmarshal(nullJSON, &fromNull); err != nil {
		t.Fatal(err)
	}
	if fromNull.MMA2.Policy != nil || fromNull.MMA2.StateSealing != nil || fromNull.MMA2.RBE != nil {
		t.Fatalf("explicit null advanced values did not normalize to nil: %#v", fromNull.MMA2)
	}

	empty := validDevice()
	empty.MMA2.Policy = &mma2composer.Policy{}
	empty.MMA2.StateSealing = map[string]interface{}{}
	empty.MMA2.RBE = map[string]interface{}{}
	emptyJSON, err := json.Marshal(empty)
	if err != nil {
		t.Fatal(err)
	}
	var emptyProjection map[string]interface{}
	if err := json.Unmarshal(emptyJSON, &emptyProjection); err != nil {
		t.Fatal(err)
	}
	emptyMMA2 := emptyProjection["mma2"].(map[string]interface{})
	if _, exists := emptyMMA2["policy"]; !exists {
		t.Fatal("explicit empty policy was omitted")
	}
	for _, key := range []string{"state_sealing", "rbe"} {
		if _, exists := emptyMMA2[key]; exists {
			t.Fatalf("empty %s was not normalized to omission", key)
		}
	}

	composer := store.composer()
	cfg := EffectiveMMA2Config{Extra: map[string]interface{}{
		"rbe":         map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:9001"}},
		"custom_root": "keep",
	}}
	if err := composer.Commit(cfg, OwnershipDoc{}); err != nil {
		t.Fatal(err)
	}

	if err := store.SaveAndCompose(device); err != nil {
		t.Fatalf("failed to save device with advanced settings: %v", err)
	}

	doc, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Devices) != 1 {
		t.Fatalf("loaded %d devices, want 1", len(doc.Devices))
	}

	if !reflect.DeepEqual(doc.Devices[0].MMA2.Policy, device.MMA2.Policy) {
		t.Fatalf("policy changed on reload: got %v, want %v", doc.Devices[0].MMA2.Policy.Rules[0], device.MMA2.Policy.Rules[0])
	}

	if doc.Devices[0].MMA2.Extra["custom_policy"] != "keep" {
		t.Fatal("advanced policy value lost")
	}
	if doc.Devices[0].MMA2.StateSealing["enabled"] != false {
		t.Fatal("advanced sealing disabled flag lost")
	}
	if doc.Devices[0].MMA2.RBE == nil {
		t.Fatal("RBE field omitted entirely")
	}

	loaded, err := composer.LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Extra["custom_root"] != "keep" {
		t.Fatal("shared root extension lost")
	}
	var memory *MMA2Memory
	for listenerIndex := range loaded.Listeners {
		if listenPort(loaded.Listeners[listenerIndex].Listen) != device.MMA2.Port {
			continue
		}
		for memoryIndex := range loaded.Listeners[listenerIndex].Memory {
			candidate := &loaded.Listeners[listenerIndex].Memory[memoryIndex]
			if candidate.UnitID == device.MMA2.UnitID {
				memory = candidate
				break
			}
		}
	}
	if memory == nil {
		t.Fatalf("effective memory (%d,%d) not found", device.MMA2.Port, device.MMA2.UnitID)
	}
	if memory.Extra["custom_policy"] != "keep" || memory.Extra["rbe"] == nil || memory.Extra["state_sealing"] == nil {
		t.Fatalf("advanced memory projection incomplete: %#v", memory.Extra)
	}

	configBefore, err := os.ReadFile(store.EffectiveConfigPath())
	if err != nil {
		t.Fatal(err)
	}
	deviceBefore, err := os.ReadFile(store.DevicesPath())
	if err != nil {
		t.Fatal(err)
	}
	invalid := device
	invalid.MMA2.StateSealing = map[string]interface{}{"enabled": true, "area": "coil", "address": 65535}
	if err := store.SaveAndCompose(invalid); err == nil {
		t.Fatal("invalid recognized state sealing was accepted")
	}
	configAfter, err := os.ReadFile(store.EffectiveConfigPath())
	if err != nil {
		t.Fatal(err)
	}
	deviceAfter, err := os.ReadFile(store.DevicesPath())
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(configAfter, configBefore) {
		t.Fatal("invalid save changed effective config")
	}
	if !reflect.DeepEqual(deviceAfter, deviceBefore) {
		t.Fatal("invalid save changed device document")
	}
}
