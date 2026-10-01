package simulator

import (
	"encoding/json"
	"net"
	"os"
	"reflect"
	"strings"
	"testing"
)

func identityDevice(fields map[string]interface{}) DeviceDefinition {
	device := validDevice()
	if fields != nil {
		device.MMA2.FC43 = &fields
	}
	return device
}

func TestFC43ProjectionUsesPortAndUnitAndPreservesReset(t *testing.T) {
	cfg := EffectiveMMA2Config{Listeners: []MMA2Listener{
		{Listen: "0.0.0.0:502", Memory: []MMA2Memory{
			{UnitID: 1, Extra: map[string]interface{}{"fc43": map[string]interface{}{"product_code": "first"}}},
			{UnitID: 2, Extra: map[string]interface{}{"fc43": map[string]interface{}{"product_code": "second"}}},
		}},
		{Listen: "0.0.0.0:503", Memory: []MMA2Memory{{UnitID: 2, Extra: map[string]interface{}{"fc43": map[string]interface{}{"product_code": "other-port"}}}}},
	}}
	params := MMA2Params{Port: 502, UnitID: 2}
	if err := projectAdvancedSettings(&params, cfg); err != nil {
		t.Fatal(err)
	}
	if params.FC43 == nil || (*params.FC43)["product_code"] != "second" {
		t.Fatal("wrong logical device identity projected")
	}
	reset := map[string]interface{}{}
	params.FC43 = &reset
	if err := projectAdvancedSettings(&params, cfg); err != nil {
		t.Fatal(err)
	}
	inheritMemorySettings(&params, cfg)
	if len(*params.FC43) != 0 {
		t.Fatal("explicit defaults re-inherited")
	}
}

func TestFC43SaveOmissionOverridesAndDefaults(t *testing.T) {
	store := Store{Root: t.TempDir()}
	for _, fields := range []map[string]interface{}{nil,
		{"vendor_name": "github.com/tamzrod", "product_code": "MMA2", "major_minor_revision": "2.x"},
		{"product_code": "MyDevice"}, {},
	} {
		device := identityDevice(fields)
		wire, err := json.Marshal(device)
		if err != nil {
			t.Fatal(err)
		}
		var decoded DeviceDefinition
		if err := json.Unmarshal(wire, &decoded); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(device.MMA2.FC43, decoded.MMA2.FC43) {
			t.Fatal("JSON lost override presence")
		}
		if err := store.SaveAndCompose(decoded); err != nil {
			t.Fatal(err)
		}
		cfg, err := store.loadEffective()
		if err != nil {
			t.Fatal(err)
		}
		actual, exists := cfg.Listeners[0].Memory[0].Extra["fc43"]
		if len(fields) == 0 {
			if exists {
				t.Fatalf("default identity serialized: %v", actual)
			}
		} else if !reflect.DeepEqual(actual, fields) {
			t.Fatalf("identity: got %v want %v", actual, fields)
		}
		loaded, err := store.Load()
		if err != nil {
			t.Fatal(err)
		}
		if err := store.ComposeDocument(loaded); err != nil {
			t.Fatal(err)
		}
		cfg, err = store.loadEffective()
		if err != nil {
			t.Fatal(err)
		}
		if len(fields) == 0 && cfg.Listeners[0].Memory[0].Extra["fc43"] != nil {
			t.Fatal("defaults re-inherited on reload")
		}
	}
}

func TestInvalidAdvancedSettingsNeverPersistOrRestart(t *testing.T) {
	for _, section := range []string{"empty", "non-ascii", "oversized", "rbe", "sealing", "policy"} {
		for _, enabled := range []bool{true, false} {
			t.Run(section+map[bool]string{true: "/enabled", false: "/disabled"}[enabled], func(t *testing.T) {
				store := Store{Root: t.TempDir()}
				if err := store.SaveAndCompose(validDevice()); err != nil {
					t.Fatal(err)
				}
				paths := []string{store.EffectiveConfigPath(), store.OwnershipPath(), store.DevicesPath()}
				before := make([][]byte, len(paths))
				for i, path := range paths {
					var err error
					before[i], err = os.ReadFile(path)
					if err != nil {
						t.Fatal(err)
					}
				}
				device := identityDevice(map[string]interface{}{"product_code": "valid"})
				device.Enabled = enabled
				switch section {
				case "empty":
					(*device.MMA2.FC43)["product_code"] = ""
				case "non-ascii":
					(*device.MMA2.FC43)["product_code"] = "é"
				case "oversized":
					(*device.MMA2.FC43)["product_code"] = strings.Repeat("a", 245)
				case "rbe":
					device.MMA2.RBE = map[string]interface{}{"coils": []interface{}{map[string]interface{}{"id": 1, "start": 65535, "count": 2}}}
				case "sealing":
					device.MMA2.StateSealing = map[string]interface{}{"enabled": true, "area": "coil", "address": 65535}
				case "policy":
					device.MMA2.Policy = &MMA2Policy{Rules: []MMA2PolicyRule{{ID: "invalid", SourceIP: []string{"invalid-ip"}, AllowFC: []uint8{43}}}}
				}
				recorder := &applyRecorder{}
				router := NewApplyRouter(store, recorder, recorder)
				validChange := identityDevice(map[string]interface{}{"product_code": "EarlierValidChange"})
				validChange.Name = "earlier-valid-device"
				validChange.MMA2.Port++
				if _, err := router.Apply(Document{Devices: []DeviceDefinition{validChange, device}}); err == nil {
					t.Fatal("invalid apply accepted")
				}
				if recorder.structural != 0 || recorder.timing != 0 {
					t.Fatal("invalid input reached runtime apply")
				}
				if err := store.SaveAndCompose(device); err == nil {
					t.Fatal("invalid direct save accepted")
				}
				for i, path := range paths {
					after, err := os.ReadFile(path)
					if err != nil {
						t.Fatal(err)
					}
					if string(before[i]) != string(after) {
						t.Fatalf("invalid save changed %s", path)
					}
				}
			})
		}
	}
}

func TestFC43ValidApplyUsesRestartPath(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	store := Store{Root: t.TempDir()}
	device := identityDevice(map[string]interface{}{"product_code": strings.Repeat("a", 244)})
	device.MMA2.Port = uint16(listener.Addr().(*net.TCPAddr).Port)
	applier := newSchedulerApplier(store, Document{}, false)
	defer applier.Stop()
	acknowledgeNextRestart(store)
	result, err := NewApplyRouter(store, applier, applier).Apply(Document{Devices: []DeviceDefinition{device}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Path != ApplyStructural {
		t.Fatalf("wrong path: %s", result.Path)
	}
	cfg, err := store.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(cfg.Listeners[0].Memory[0].Extra["fc43"], *device.MMA2.FC43) {
		t.Fatal("FC43 not committed")
	}
}
