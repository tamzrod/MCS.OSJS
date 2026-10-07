package mma2composer

import "testing"

// PERSIST-008 self-check: user-created RBE rules coexist with system persistence
// rules, overlap is permitted, and persistence lifecycle operations never mutate
// user rules.
func TestUserRBECompatibilityCoexistenceAndOverlap(t *testing.T) {
	memory := Memory{
		UnitID:      1,
		Coils:       &Area{Start: 0, Count: 8},
		HoldingRegs: &Area{Start: 0, Count: 8},
		Persistence: &Persistence{Enabled: boolEnabled(true)},
	}
	// A user rule fully inside the persistence-owned coils range, plus an
	// unrelated user rule on holding registers.
	userRules := []UserRBERule{
		{Area: "coils", Start: 2, Count: 4},
		{Area: "holding_registers", Start: 0, Count: 3},
	}
	userCopy := append([]UserRBERule(nil), userRules...)

	view := RBECoexist(memory, userRules)
	if len(view.Persistence) != 2 || len(view.UserRules) != 2 {
		t.Fatalf("expected coexistence of 2+2 rules, got %+v", view)
	}
	// User rules carried through unchanged.
	if !UserRBEUnchanged(view, userCopy) {
		t.Fatalf("user rules changed: %+v vs %+v", view.UserRules, userCopy)
	}
	// Overlap is reported and allowed.
	if !PersistenceOverlapsUser(view, userRules[0]) {
		t.Fatal("expected overlap with persistence coils range")
	}
	if !PersistenceOverlapsUser(view, userRules[1]) {
		t.Fatal("expected overlap with persistence holding range")
	}

	// A persistence lifecycle change (disable/enable/range change) must not
	// mutate user rules.
	memory.Persistence.Enabled = boolEnabled(false)
	viewDisabled := RBECoexist(memory, userRules)
	if len(viewDisabled.Persistence) != 0 {
		t.Fatalf("disable should clear persistence rules: %+v", viewDisabled)
	}
	if !UserRBEUnchanged(viewDisabled, userCopy) {
		t.Fatalf("disable mutated user rules: %+v", viewDisabled.UserRules)
	}
	memory.Coils = &Area{Start: 100, Count: 8}
	memory.Persistence.Enabled = boolEnabled(true)
	viewRegenerated := RBECoexist(memory, userRules)
	if !UserRBEUnchanged(viewRegenerated, userCopy) {
		t.Fatalf("regeneration mutated user rules: %+v", viewRegenerated.UserRules)
	}
	// The user's coils rule no longer overlaps after the persistence range moved.
	if PersistenceOverlapsUser(viewRegenerated, userRules[0]) {
		t.Fatal("stale overlap expected to clear after range change")
	}
}

// A user rule that does not overlap any persistence rule is simply not flagged.
func TestPersistenceOverlapNonOverlapping(t *testing.T) {
	memory := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, Persistence: &Persistence{Enabled: boolEnabled(true)}}
	view := RBECoexist(memory, nil)
	if PersistenceOverlapsUser(view, UserRBERule{Area: "coils", Start: 4, Count: 2}) {
		t.Fatal("adjacent range must not overlap")
	}
	if PersistenceOverlapsUser(view, UserRBERule{Area: "input_registers", Start: 0, Count: 2}) {
		t.Fatal("different area must not overlap")
	}
}

// The caller's user-rules slice must not be mutated by building the view.
func TestRBECoexistDoesNotMutateUserRules(t *testing.T) {
	memory := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, Persistence: &Persistence{Enabled: boolEnabled(true)}}
	userRules := []UserRBERule{{Area: "coils", Start: 1, Count: 1}}
	view := RBECoexist(memory, userRules)
	view.UserRules[0].Start = 999
	if userRules[0].Start != 1 {
		t.Fatalf("underlying user slice mutated: %+v", userRules)
	}
}
