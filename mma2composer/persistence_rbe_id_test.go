package mma2composer

import "testing"

// PERSIST-009 self-check: persistence RBE IDs are allocated globally unique,
// avoiding user-owned IDs, dynamically within the one-byte 1..255 space.
func TestAllocatePersistenceRBEIDsUniqueAndAvoidsUserIDs(t *testing.T) {
	rules, err := AllocatePersistenceRBEIDs(DerivePersistenceRBE(Memory{
		UnitID:      1,
		Coils:       &Area{Start: 0, Count: 4},
		HoldingRegs: &Area{Start: 0, Count: 8},
		InputRegs:   &Area{Start: 0, Count: 8},
	}), 1, 2, 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(rules) != 3 {
		t.Fatalf("want 3 allocated rules, got %+v", rules)
	}
	// User IDs 1,2,5 are reserved; lowest free IDs 3,4,6 are used next.
	wantIDs := []uint8{3, 4, 6}
	seen := map[uint8]bool{}
	for i, rule := range rules {
		if rule.ID != wantIDs[i] {
			t.Fatalf("rule %d: want ID %d, got %d", i, wantIDs[i], rule.ID)
		}
		if rule.ID == 0 || seen[rule.ID] {
			t.Fatalf("ID %d is invalid or duplicated", rule.ID)
		}
		seen[rule.ID] = true
	}
}

// No permanent partition: persistence does not reserve a fixed range, so a
// different user-ID set yields a different (but still collision-free) mapping.
func TestAllocatePersistenceRBEIDsNoPermanentPartition(t *testing.T) {
	memory := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, HoldingRegs: &Area{Start: 0, Count: 8}}
	derived := DerivePersistenceRBE(memory)

	lowFree, err := AllocatePersistenceRBEIDs(derived)
	if err != nil {
		t.Fatal(err)
	}
	// Without user reservations the lowest IDs 1,2 are used.
	if lowFree[0].ID != 1 || lowFree[1].ID != 2 {
		t.Fatalf("expected IDs 1,2 with no users, got %+v", lowFree)
	}

	// A user holding ID 1 pushes persistence to 2,3 for the same rules.
	shifted, err := AllocatePersistenceRBEIDs(derived, 1)
	if err != nil {
		t.Fatal(err)
	}
	if shifted[0].ID != 2 || shifted[1].ID != 3 {
		t.Fatalf("expected IDs 2,3 with user ID 1, got %+v", shifted)
	}
	if lowFree[0].ID == shifted[0].ID {
		t.Fatal("allocation should be dynamic, not a fixed partition")
	}
}

// Exhaustion fails closed with no partial set, and inputs are never mutated.
func TestAllocatePersistenceRBEIDsExhaustionAndNoMutation(t *testing.T) {
	// Reserve all but one ID so two rules cannot both be placed.
	used := make([]uint8, 0, 255)
	for id := 2; id <= 255; id++ {
		used = append(used, uint8(id))
	}
	rules := []PersistenceRBERule{
		{Area: "coils", Start: 0, Count: 4, SystemOwned: true, ID: 200},
		{Area: "holding_registers", Start: 0, Count: 8, SystemOwned: true, ID: 201},
	}
	original := append([]PersistenceRBERule(nil), rules...)
	assigned, err := AllocatePersistenceRBEIDs(rules, used...)
	if err == nil {
		t.Fatalf("expected exhaustion error, got %+v", assigned)
	}
	if assigned != nil {
		t.Fatalf("exhaustion must not return a partial set: %+v", assigned)
	}
	for i := range rules {
		if rules[i] != original[i] {
			t.Fatalf("input mutated on failure: %+v vs %+v", rules, original)
		}
	}

	// Exactly one free ID (1) can still place a single rule without mutating it.
	one, err := AllocatePersistenceRBEIDs(rules[:1], used...)
	if err != nil || len(one) != 1 || one[0].ID != 1 {
		t.Fatalf("single placement failed: %+v err=%v", one, err)
	}
	if rules[0].ID != 200 {
		t.Fatalf("success path mutated the input: %+v", rules[0])
	}
}
