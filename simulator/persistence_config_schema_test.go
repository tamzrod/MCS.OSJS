package simulator

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// PERSIST-001 self-check: persistence.enabled is represented per existing
// Port -> Unit ID -> Memory identity, round-trips, and introduces no
// persistence-owned range fields.
func TestPersistenceConfigurationSchemaRoundTrip(t *testing.T) {
	store := Store{Root: t.TempDir()}
	if err := store.composer().Commit(EffectiveMMA2Config{Extra: map[string]interface{}{
		"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:9001"}},
	}}, OwnershipDoc{}); err != nil {
		t.Fatal(err)
	}

	enabled := true
	disabled := false
	withEnabled := validDevice()
	withEnabled.MMA2.Persistence = &mma2composer.Persistence{Enabled: &enabled, Directory: "/var/lib/mma2/unit1"}
	// PERSIST-002 invariant: persistence ON requires state sealing present+enabled.
	withEnabled.MMA2.StateSealing = map[string]interface{}{"enabled": true, "area": "coil", "address": 0}
	withDisabled := validDevice()
	withDisabled.MMA2.Persistence = &mma2composer.Persistence{Enabled: &disabled}

	encoded, err := json.Marshal(withEnabled)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), `"persistence":{`) || !strings.Contains(string(encoded), `"enabled":true`) || !strings.Contains(string(encoded), `"directory":"/var/lib/mma2/unit1"`) {
		t.Fatalf("persistence not encoded per memory: %s", encoded)
	}
	var round DeviceDefinition
	if err := json.Unmarshal(encoded, &round); err != nil {
		t.Fatal(err)
	}
	if round.MMA2.Persistence == nil || round.MMA2.Persistence.Enabled == nil || *round.MMA2.Persistence.Enabled != true {
		t.Fatalf("persistence lost on JSON round trip: %s", encoded)
	}
	if round.MMA2.Persistence.Directory != "/var/lib/mma2/unit1" {
		t.Fatalf("persistence directory lost on JSON round trip: %s", encoded)
	}

	for _, tc := range []struct {
		name string
		def  DeviceDefinition
		want bool
	}{
		{"enabled", withEnabled, true},
		{"disabled", withDisabled, false},
	} {
		if err := store.SaveAndCompose(tc.def); err != nil {
			t.Fatalf("%s: compose failed: %v", tc.name, err)
		}
		doc, err := store.Load()
		if err != nil {
			t.Fatal(err)
		}
		got := doc.Devices[0].MMA2.Persistence
		if got == nil || got.Enabled == nil || *got.Enabled != tc.want {
			t.Fatalf("%s: persistence lost in device document: %+v", tc.name, got)
		}
		cfg, err := store.composer().LoadEffective()
		if err != nil {
			t.Fatal(err)
		}
		var mem *mma2composer.Memory
		for li := range cfg.Listeners {
			for mi := range cfg.Listeners[li].Memory {
				if cfg.Listeners[li].Memory[mi].UnitID == 1 {
					mem = &cfg.Listeners[li].Memory[mi]
				}
			}
		}
		if mem == nil || mem.Persistence == nil || mem.Persistence.Enabled == nil || *mem.Persistence.Enabled != tc.want {
			t.Fatalf("%s: persistence lost in effective config: %+v", tc.name, mem)
		}
		if _, ok := mem.Extra["persistence"]; ok {
			t.Fatalf("%s: persistence leaked into generic Extra map", tc.name)
		}
		raw, err := os.ReadFile(store.EffectiveConfigPath())
		if err != nil {
			t.Fatal(err)
		}
		for _, banned := range []string{"persistence_start", "persistence_count", "persistence_area"} {
			if strings.Contains(string(raw), banned) {
				t.Fatalf("%s: persistence introduced range identity %q: %s", tc.name, banned, raw)
			}
		}
	}

	if err := store.SaveAndCompose(validDevice()); err != nil {
		t.Fatalf("baseline compose: %v", err)
	}
	cfg, err := store.composer().LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range cfg.Listeners {
		for _, m := range l.Memory {
			if m.Persistence == nil || m.Persistence.Enabled == nil || *m.Persistence.Enabled != false {
				t.Fatalf("explicit-but-absent legacy save should inherit the existing memory flag: %+v", m)
			}
		}
	}

	// A configuration that never had persistence stays byte-identical apart
	// from the newly composed listener: no persistence key is introduced.
	fresh := Store{Root: t.TempDir()}
	if err := fresh.composer().Commit(EffectiveMMA2Config{Extra: map[string]interface{}{
		"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:9001"}},
	}}, OwnershipDoc{}); err != nil {
		t.Fatal(err)
	}
	if err := fresh.SaveAndCompose(validDevice()); err != nil {
		t.Fatalf("fresh baseline compose: %v", err)
	}
	freshCfg, err := fresh.composer().LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range freshCfg.Listeners {
		for _, m := range l.Memory {
			if m.Persistence != nil {
				t.Fatalf("baseline without persistence gained it: %+v", m)
			}
		}
	}
	freshRaw, err := os.ReadFile(fresh.EffectiveConfigPath())
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(freshRaw), "persistence") {
		t.Fatalf("baseline effective config mentions persistence: %s", freshRaw)
	}
}
