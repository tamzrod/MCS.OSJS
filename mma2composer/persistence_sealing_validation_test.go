package mma2composer

import (
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// PERSIST-002 self-check: persistence-enabled memory must be rejected unless
// State Sealing is present and enabled, and validation must not mutate sealing.
func TestPersistenceRequiresEnabledStateSealing(t *testing.T) {
	sealing := func(enabled interface{}) map[string]interface{} {
		base := map[string]interface{}{"area": "coil", "address": 0}
		if enabled != nil {
			base["enabled"] = enabled
		}
		return base
	}
	memoryWith := func(persistence *Persistence, sealingBlock map[string]interface{}) Memory {
		m := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, Extra: map[string]interface{}{}}
		m.Persistence = persistence
		if sealingBlock != nil {
			m.Extra["state_sealing"] = sealingBlock
		}
		return m
	}

	on := &Persistence{Enabled: boolPtr(true), Directory: "/var/lib/mma2/unit1"}
	off := &Persistence{Enabled: boolPtr(false)}

	for _, tc := range []struct {
		name        string
		memory      Memory
		wantErr     bool
		errContains string
	}{
		{"persistence ON, sealing absent", memoryWith(on, nil), true, "state sealing"},
		{"persistence ON, sealing explicit disabled", memoryWith(on, sealing(false)), true, "state sealing"},
		{"persistence ON, sealing enabled omitted (defaults on)", memoryWith(on, sealing(nil)), false, ""},
		{"persistence ON, sealing explicit enabled", memoryWith(on, sealing(true)), false, ""},
		{"persistence OFF, sealing absent", memoryWith(off, nil), false, ""},
		{"persistence OFF, sealing disabled", memoryWith(off, sealing(false)), false, ""},
		{"persistence nil, sealing absent", memoryWith(nil, nil), false, ""},
		{"persistence nil, sealing disabled", memoryWith(nil, sealing(false)), false, ""},
	} {
		err := ValidateMemory(tc.memory)
		if tc.wantErr {
			if err == nil {
				t.Fatalf("%s: expected error, got nil", tc.name)
			}
			if !strings.Contains(err.Error(), tc.errContains) {
				t.Fatalf("%s: error %q does not mention %q", tc.name, err.Error(), tc.errContains)
			}
		} else if err != nil {
			t.Fatalf("%s: expected no error, got %v", tc.name, err)
		}
	}
}

// Validation must never silently enable or mutate State Sealing, even on the
// rejection path.
func TestPersistenceValidationDoesNotMutateSealing(t *testing.T) {
	on := &Persistence{Enabled: boolPtr(true), Directory: "/var/lib/mma2/unit1"}
	memory := Memory{
		UnitID:      1,
		Coils:       &Area{Start: 0, Count: 4},
		Persistence: on,
		Extra:       map[string]interface{}{"state_sealing": map[string]interface{}{"enabled": false, "area": "coil", "address": 0}},
	}
	before := Memory{}
	bytes, _ := yaml.Marshal(memory)
	if err := yaml.Unmarshal(bytes, &before); err != nil {
		t.Fatal(err)
	}
	if err := ValidateMemory(memory); err == nil {
		t.Fatal("expected sealed-disabled persistence to be rejected")
	}
	if !reflect.DeepEqual(memory, before) {
		t.Fatalf("validation mutated memory:\n before=%+v\n after =%+v", before, memory)
	}
}

func boolPtr(v bool) *bool { return &v }
