package mma2composer

import "testing"

// PERSIST-004 self-check: persistence RBE rules are system-owned/locked and a
// supported configuration path cannot edit or delete them, while ordinary user
// RBE behavior is unaffected.
func TestPersistenceRulesAreSystemOwned(t *testing.T) {
	derived := DerivePersistenceRBE(Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, HoldingRegs: &Area{Start: 0, Count: 8}})
	if len(derived) != 2 {
		t.Fatalf("want 2 derived rules, got %d", len(derived))
	}
	for _, rule := range derived {
		if !rule.SystemOwned {
			t.Fatalf("derived rule not marked system-owned: %+v", rule)
		}
	}
}

func TestPersistenceRuleMutationRejected(t *testing.T) {
	derived := DerivePersistenceRBE(Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}, HoldingRegs: &Area{Start: 0, Count: 8}})
	coils := derived[0]
	holding := derived[1]

	unchanged := func() []PersistenceRBERule { return []PersistenceRBERule{coils, holding} }

	cases := []struct {
		name      string
		requested []PersistenceRBERule
	}{
		{"range edit", []PersistenceRBERule{{Area: "coils", Start: 0, Count: 99, SystemOwned: true}, holding}},
		{"start edit", []PersistenceRBERule{{Area: "coils", Start: 5, Count: 4, SystemOwned: true}, holding}},
		{"delete", []PersistenceRBERule{holding}},
		{"add", append(unchanged(), PersistenceRBERule{Area: "input_registers", Start: 0, Count: 1, SystemOwned: true})},
		{"ownership strip", []PersistenceRBERule{{Area: "coils", Start: 0, Count: 4, SystemOwned: false}, holding}},
	}
	for _, tc := range cases {
		if err := ValidatePersistenceRuleMutation(tc.requested, derived); err == nil {
			t.Fatalf("%s: expected rejection, got nil", tc.name)
		}
	}

	// The exact layout-derived set (identical value, any order) is accepted.
	reordered := []PersistenceRBERule{holding, coils}
	if err := ValidatePersistenceRuleMutation(reordered, derived); err != nil {
		t.Fatalf("unchanged set rejected: %v", err)
	}
	// Removing persistence from an area removes its rule via the layout, not via
	// a user edit: the derived set simply no longer contains that area.
	reduced := DerivePersistenceRBE(Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}})
	if err := ValidatePersistenceRuleMutation([]PersistenceRBERule{coils}, reduced); err != nil {
		t.Fatalf("layout-derived reduced set rejected: %v", err)
	}
}

// The guard must not mutate the caller's slices.
func TestPersistenceRuleMutationDoesNotMutateInputs(t *testing.T) {
	derived := DerivePersistenceRBE(Memory{UnitID: 1, Coils: &Area{Start: 0, Count: 4}})
	requested := []PersistenceRBERule{{Area: "coils", Start: 9, Count: 9, SystemOwned: true}}
	reqCopy := append([]PersistenceRBERule(nil), requested...)
	derCopy := append([]PersistenceRBERule(nil), derived...)
	if err := ValidatePersistenceRuleMutation(requested, derived); err == nil {
		t.Fatal("expected rejection")
	}
	for i := range requested {
		if requested[i] != reqCopy[i] {
			t.Fatalf("requested mutated: %+v vs %+v", requested, reqCopy)
		}
	}
	for i := range derived {
		if derived[i] != derCopy[i] {
			t.Fatalf("derived mutated: %+v vs %+v", derived, derCopy)
		}
	}
}
