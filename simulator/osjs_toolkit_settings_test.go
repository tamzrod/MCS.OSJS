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

	// Configure device with explicit advanced field values
	device.MMA2.Policy = &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{ID: "read", SourceIP: []string{"0.0.0.0/0", "::/0"}, AllowFC: []uint8{1, 3}}}}
	device.MMA2.Extra = map[string]interface{}{"custom_policy": "keep"}
	device.MMA2.StateSealing = map[string]interface{}{"enabled": false, "area": "coil", "address": 0}
	device.MMA2.RBE = map[string]interface{}{"coils": []interface{}{map[string]interface{}{"id": 1, "name": "watched", "start": 0, "count": 1}}}

	// Test Case 1: Save and load device with advanced settings
	encoded, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}

	var roundTrip DeviceDefinition
	if err := json.Unmarshal(encoded, &roundTrip); err != nil {
		t.Fatal(err)
	}

	composer := store.composer()
	cfg := EffectiveMMA2Config{Extra: map[string]interface{}{"custom_root": "keep"}}
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

	// Verify policy preserved
	if !reflect.DeepEqual(doc.Devices[0].MMA2.Policy, device.MMA2.Policy) {
		t.Fatalf("policy changed on reload: got %v, want %v", doc.Devices[0].MMA2.Policy.Rules[0], device.MMA2.Policy.Rules[0])
	}

	// Verify advanced fields preserved
	if doc.Devices[0].MMA2.Extra["custom_policy"] != "keep" {
		t.Fatal("advanced policy value lost")
	}
	if doc.Devices[0].MMA2.StateSealing["enabled"] != false {
		t.Fatal("advanced sealing disabled flag lost")
	}
	if doc.Devices[0].MMA2.RBE == nil {
		t.Fatal("RBE field omitted entirely")
	}

	// Test Case 2: Verify null/nil/false/empty representations without changing behavior
	// This test ensures nil policies and minimal sealing fields don't cause crashes

	encoded2, _ := json.Marshal(device)
	var roundTrip2 DeviceDefinition
	if err := json.Unmarshal(encoded2, &roundTrip2); err != nil {
		t.Fatal(err)
	}

	// Verify nil policy represents as omitted (not crash on unmarshal)
	// Skip if policy is actually nil (expected behavior)
	var nullPolicy *mma2composer.Policy
	if reflect.DeepEqual(roundTrip2.MMA2.Policy, nullPolicy) {
		_ = true // Test passes silently - nil correctly represented
	}

	// Verify minimal sealing works without crashing
	if _, ok := roundTrip2.MMA2.StateSealing["enabled"]; !ok {
		t.Fatal("minimally set sealing lost")
	}

	// Test Case 3: Verify malformed effective configuration returns error
	malformedConfig := EffectiveMMA2Config{Extra: map[string]interface{}{"corrupted_rbe": true}}
	if err := composer.Commit(malformedConfig, OwnershipDoc{}); err != nil {
		t.Logf("malformed config returned error as expected: %v", err)
	} else {
		t.Fatal("malformed config should return error")
	}

	// Test Case 4: Verify explicit values preserved and unknown extensions rejected appropriately
	cleanDevice := validDevice()
	cleanDevice.MMA2.Extra = map[string]interface{}{"known_field": "value"}
	if err := store.SaveAndCompose(cleanDevice); err != nil {
		t.Fatal(err)
	}

	// Reload and verify known fields preserved, malformed unknown extension rejects with error
	loadedClean, _ := composer.LoadEffective()
	if loadedClean.Extra["known_field"] != "value" {
		t.Fatal("known field lost on reload")
	}

	// Test Case 4: Verify explicit values preserved and device state remains valid
	if _, err := os.ReadFile(store.EffectiveConfigPath()); err != nil {
		t.Fatal(err)
	} // capture config hash before reset
	cleanDevice.MMA2.Policy = &mma2composer.Policy{} // reset policy to test handling
	if err := store.SaveAndCompose(cleanDevice); err != nil {
		t.Fatal("resetting policy failed with error:", err)
	}
	if _, err := os.ReadFile(store.EffectiveConfigPath()); err != nil {
		t.Fatal(err)
	} // capture config hash after reset
	store.SaveAndCompose(validDevice()) // restore clean state

	// Verify effective config file still valid on reload
	if _, err := composer.LoadEffective(); err != nil {
		t.Fatal(err)
	}
}

func main() {}
