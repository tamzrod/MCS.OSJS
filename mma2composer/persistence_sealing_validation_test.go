package mma2composer

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

// NPE-01 acceptance 5: persistence and State Sealing validate independently.
// A persistence-enabled memory is valid whether sealing is absent, disabled or
// enabled, and a sealing decision never affects persistence validity.
func TestPersistenceIndependentOfStateSealing(t *testing.T) {
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

	on := &Persistence{Enabled: boolPtr(true)}
	off := &Persistence{Enabled: boolPtr(false)}

	for _, tc := range []struct {
		name   string
		memory Memory
	}{
		{"persistence ON, sealing absent", memoryWith(on, nil)},
		{"persistence ON, sealing explicit disabled", memoryWith(on, sealing(false))},
		{"persistence ON, sealing enabled omitted (defaults on)", memoryWith(on, sealing(nil))},
		{"persistence ON, sealing explicit enabled", memoryWith(on, sealing(true))},
		{"persistence OFF, sealing absent", memoryWith(off, nil)},
		{"persistence OFF, sealing disabled", memoryWith(off, sealing(false))},
		{"persistence nil, sealing absent", memoryWith(nil, nil)},
		{"persistence nil, sealing disabled", memoryWith(nil, sealing(false))},
	} {
		if err := ValidateMemory(tc.memory); err != nil {
			t.Fatalf("%s: persistence and State Sealing must validate independently: %v", tc.name, err)
		}
	}
}

// Validation must never silently enable or mutate State Sealing on persistence's
// behalf: an enabled persistence block with sealing explicitly disabled is valid
// and leaves the sealing configuration untouched.
func TestPersistenceValidationDoesNotMutateSealing(t *testing.T) {
	on := &Persistence{Enabled: boolPtr(true)}
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
	if err := ValidateMemory(memory); err != nil {
		t.Fatalf("persistence must not require sealing: %v", err)
	}
	if !reflect.DeepEqual(memory, before) {
		t.Fatalf("validation mutated memory:\n before=%+v\n after =%+v", before, memory)
	}
}

func boolPtr(v bool) *bool { return &v }
