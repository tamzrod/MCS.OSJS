package mma2composer

import "testing"

func boolEnabled(v bool) *bool { return &v }

// PERSIST-007 self-check: disabling persistence removes all persistence-owned
// RBE for that memory, ordinary user RBE is unaffected, and re-enabling
// regenerates projections from the current memory ranges.
func TestPersistenceDisableCleanupAndReEnable(t *testing.T) {
	memory := Memory{
		UnitID:      1,
		Coils:       &Area{Start: 0, Count: 4},
		HoldingRegs: &Area{Start: 0, Count: 8},
		Persistence: &Persistence{Enabled: boolEnabled(true)},
	}

	// Enabled: projection is present.
	enabledProjection := PersistenceRBEProjection(memory)
	if len(enabledProjection) != 2 {
		t.Fatalf("enabled projection should have 2 rules, got %+v", enabledProjection)
	}
	for _, rule := range enabledProjection {
		if !rule.SystemOwned {
			t.Fatalf("projection rule not system-owned: %+v", rule)
		}
	}

	// Disable: all persistence-owned RBE for the memory is removed.
	memory.Persistence.Enabled = boolEnabled(false)
	if cleared := PersistenceRBEProjection(memory); len(cleared) != 0 {
		t.Fatalf("disabling persistence left %+v", cleared)
	}

	// Persistence absent also means no persistence-owned RBE.
	noPersistence := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}}
	if cleared := PersistenceRBEProjection(noPersistence); len(cleared) != 0 {
		t.Fatalf("absent persistence produced %+v", cleared)
	}

	// Re-enable after a range change: projections regenerate from current ranges.
	memory.Coils = &Area{Start: 10, Count: 20}
	memory.Persistence.Enabled = boolEnabled(true)
	regenerated := PersistenceRBEProjection(memory)
	if !equalRules(regenerated, []PersistenceRBERule{
		{Area: "coils", Start: 10, Count: 20, SystemOwned: true},
		{Area: "holding_registers", Start: 0, Count: 8, SystemOwned: true},
	}) {
		t.Fatalf("re-enable did not regenerate from current ranges: %+v", regenerated)
	}
}

// User-owned RBE is a separate set; the persistence projection never carries or
// touches it.
func TestPersistenceProjectionDoesNotRepresentUserRBE(t *testing.T) {
	memory := Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, Persistence: &Persistence{Enabled: boolEnabled(true)}}
	projection := PersistenceRBEProjection(memory)
	for _, rule := range projection {
		if !rule.SystemOwned {
			t.Fatalf("persistence projection must be system-owned only: %+v", rule)
		}
	}
}
