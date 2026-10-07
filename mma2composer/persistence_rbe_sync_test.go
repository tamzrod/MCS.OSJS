package mma2composer

import "testing"

// PERSIST-005 self-check: the persistence RBE projection stays synchronized
// with the authoritative memory layout when an area's start/count changes, with
// no second user edit, and added/removed areas follow the layout.
func TestSynchronizePersistenceRBEFollowsAreaRange(t *testing.T) {
	memory := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, HoldingRegs: &Area{Start: 0, Count: 8}}

	before := SynchronizePersistenceRBE(memory)
	if len(before) != 2 || before[0] != (PersistenceRBERule{Area: "coils", Start: 0, Count: 4, SystemOwned: true}) {
		t.Fatalf("initial projection wrong: %+v", before)
	}

	// Change only the authoritative area; the derived start/count must follow.
	memory.Coils = &Area{Start: 10, Count: 20}
	afterStart := SynchronizePersistenceRBE(memory)
	if afterStart[0].Start != 10 || afterStart[0].Count != 20 {
		t.Fatalf("start/count did not follow area change: %+v", afterStart[0])
	}
	if !afterStart[0].SystemOwned {
		t.Fatalf("synchronized rule lost system ownership: %+v", afterStart[0])
	}

	// Count-only change.
	memory.Coils = &Area{Start: 10, Count: 77}
	afterCount := SynchronizePersistenceRBE(memory)
	if afterCount[0].Count != 77 || afterCount[0].Start != 10 {
		t.Fatalf("count did not follow area change: %+v", afterCount[0])
	}

	// Newly allocated area appears; removed area drops out.
	memory.InputRegs = &Area{Start: 0, Count: 5}
	memory.HoldingRegs = nil
	finalSet := SynchronizePersistenceRBE(memory)
	if len(finalSet) != 2 {
		t.Fatalf("want 2 rules after layout change, got %+v", finalSet)
	}
	if finalSet[0].Area != "coils" || finalSet[1].Area != "input_registers" {
		t.Fatalf("areas did not follow layout: %+v", finalSet)
	}

	// Synchronization requires no user edit: reconcile is idempotent once the
	// projection already matches the layout.
	if again := SynchronizePersistenceRBE(memory); !equalRules(again, finalSet) {
		t.Fatalf("synchronization not idempotent: %+v vs %+v", again, finalSet)
	}
}

func equalRules(a, b []PersistenceRBERule) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
