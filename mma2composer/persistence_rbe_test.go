package mma2composer

import "testing"

// PERSIST-003 self-check: one derived persistence RBE rule per present area,
// with start/count taken from the authoritative layout and no independent
// range fields introduced.
func TestDerivePersistenceRBEOneRulePerPresentArea(t *testing.T) {
	memory := Memory{
		UnitID:         1,
		Coils:          &Area{Start: 4, Count: 8},
		HoldingRegs:    &Area{Start: 100, Count: 50},
		DiscreteInputs: nil,
		InputRegs:      &Area{Start: 0, Count: 0},
	}
	rules := DerivePersistenceRBE(memory)

	want := []PersistenceRBERule{
		{Area: "coils", Start: 4, Count: 8},
		{Area: "holding_registers", Start: 100, Count: 50},
	}
	if len(rules) != len(want) {
		t.Fatalf("want %d derived rules, got %d: %+v", len(want), len(rules), rules)
	}
	for i := range want {
		if rules[i] != want[i] {
			t.Fatalf("rule %d: want %+v, got %+v", i, want[i], rules[i])
		}
	}
}

// Derived start/count must track the area exactly, and no range field may be
// settable independently of the layout.
func TestDerivePersistenceRBEDerivedRangeTracksLayout(t *testing.T) {
	memory := Memory{UnitID: 2, Coils: &Area{Start: 0, Count: 4}}
	before := DerivePersistenceRBE(memory)
	if len(before) != 1 || before[0].Start != 0 || before[0].Count != 4 {
		t.Fatalf("initial derivation wrong: %+v", before)
	}
	// Mutating the authoritative area is the only way the derived range changes.
	memory.Coils = &Area{Start: 16, Count: 32}
	after := DerivePersistenceRBE(memory)
	if len(after) != 1 || after[0].Start != 16 || after[0].Count != 32 {
		t.Fatalf("derived range did not follow layout change: %+v", after)
	}
}

// No area => no persistence RBE.
func TestDerivePersistenceRBEEmptyWhenNoAreas(t *testing.T) {
	if rules := DerivePersistenceRBE(Memory{UnitID: 1}); len(rules) != 0 {
		t.Fatalf("expected no derived rules, got %+v", rules)
	}
}
