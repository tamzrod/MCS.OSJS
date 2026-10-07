package mma2composer

import "testing"

// PERSIST-006 self-check: removing an authoritative area removes exactly its
// own system persistence RBE, keeps other persisted-area projections, and
// leaves the caller's slice untouched.
func TestPruneRemovedPersistenceRBERemovesOnlyRemovedArea(t *testing.T) {
	// Layout with coils, holding_registers and input_registers persisted.
	memory := Memory{
		UnitID:      1,
		Coils:       &Area{Start: 0, Count: 4},
		HoldingRegs: &Area{Start: 0, Count: 8},
		InputRegs:   &Area{Start: 0, Count: 2},
	}
	projection := DerivePersistenceRBE(memory)
	if len(projection) != 3 {
		t.Fatalf("want 3 projections, got %+v", projection)
	}
	original := append([]PersistenceRBERule(nil), projection...)

	// Remove the holding_registers area.
	memory.HoldingRegs = nil
	pruned := PruneRemovedPersistenceRBE(projection, memory)

	if len(pruned) != 2 {
		t.Fatalf("want 2 remaining projections, got %+v", pruned)
	}
	for _, rule := range pruned {
		if rule.Area == "holding_registers" {
			t.Fatalf("removed area projection survived: %+v", pruned)
		}
	}
	if !equalRules(pruned, []PersistenceRBERule{
		{Area: "coils", Start: 0, Count: 4, SystemOwned: true},
		{Area: "input_registers", Start: 0, Count: 2, SystemOwned: true},
	}) {
		t.Fatalf("other projections changed: %+v", pruned)
	}
	// The input slice must not be mutated.
	if !equalRules(projection, original) {
		t.Fatalf("input projection mutated: %+v vs %+v", projection, original)
	}
}

// A zero-count area is treated as removed, and a layout that still contains the
// area keeps its rule.
func TestPruneRemovedPersistenceRBECountZeroIsRemoved(t *testing.T) {
	memory := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, HoldingRegs: &Area{Start: 0, Count: 8}}
	projection := DerivePersistenceRBE(memory)

	memory.Coils = &Area{Start: 0, Count: 0}
	pruned := PruneRemovedPersistenceRBE(projection, memory)
	if len(pruned) != 1 || pruned[0].Area != "holding_registers" {
		t.Fatalf("zero-count area not treated as removed: %+v", pruned)
	}

	// Removing nothing keeps everything.
	prunedNone := PruneRemovedPersistenceRBE(projection, Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, HoldingRegs: &Area{Start: 0, Count: 8}})
	if !equalRules(prunedNone, projection) {
		t.Fatalf("no-op prune changed projection: %+v", prunedNone)
	}
}
