package mma2composer

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func nativeEnabledMemory() Memory {
	return Memory{
		UnitID:      1,
		Coils:       &Area{Start: 0, Count: 128},
		HoldingRegs: &Area{Start: 0, Count: 256},
		Persistence: &Persistence{Enabled: boolPtr(true), Directory: "/var/lib/mma2/unit1"},
		Extra:       map[string]interface{}{"state_sealing": map[string]interface{}{"enabled": true, "area": "coil", "address": 0}},
	}
}

// NP-01 self-check: the native per-memory persistence block round-trips through
// YAML with enabled/directory/optional ranges and no range identity leakage.
func TestNativePersistenceRoundTrip(t *testing.T) {
	memory := nativeEnabledMemory()
	memory.Persistence.Ranges = &PersistenceRanges{
		HoldingRegs: []PersistenceArea{{Start: 20, Count: 10}, {Start: 60, Count: 10}},
	}
	data, err := yaml.Marshal(EffectiveConfig{Listeners: []Listener{{
		ID: "main", Listen: ":502", Memory: []Memory{memory},
	}}})
	if err != nil {
		t.Fatal(err)
	}
	raw := string(data)
	for _, want := range []string{"persistence:", "enabled: true", "directory: /var/lib/mma2/unit1", "holding_registers:"} {
		if !strings.Contains(raw, want) {
			t.Fatalf("emitted YAML missing %q:\n%s", want, raw)
		}
	}
	var round EffectiveConfig
	if err := yaml.Unmarshal(data, &round); err != nil {
		t.Fatal(err)
	}
	if len(round.Listeners) != 1 || len(round.Listeners[0].Memory) != 1 {
		t.Fatalf("round trip lost structure: %+v", round)
	}
	got := round.Listeners[0].Memory[0].Persistence
	if got == nil || got.Enabled == nil || !*got.Enabled {
		t.Fatalf("enabled lost: %+v", got)
	}
	if got.Directory != "/var/lib/mma2/unit1" {
		t.Fatalf("directory lost: %q", got.Directory)
	}
	if !reflect.DeepEqual(got.Ranges, memory.Persistence.Ranges) {
		t.Fatalf("ranges not preserved:\n want=%+v\n got =%+v", memory.Persistence.Ranges, got.Ranges)
	}
	if err := ValidateMemory(round.Listeners[0].Memory[0]); err != nil {
		t.Fatalf("valid native persistence rejected: %v", err)
	}
}

// enabled false, or an absent block, disables persistence for that memory only.
func TestNativePersistenceDisabledAccepted(t *testing.T) {
	off := nativeEnabledMemory()
	off.Persistence = &Persistence{Enabled: boolPtr(false)}
	if err := ValidateMemoryPersistence(off); err != nil {
		t.Fatalf("explicit disabled should be accepted: %v", err)
	}
	absent := nativeEnabledMemory()
	absent.Persistence = nil
	if err := ValidateMemoryPersistence(absent); err != nil {
		t.Fatalf("absent persistence should be accepted: %v", err)
	}
}

// enabled true without a directory fails closed.
func TestNativePersistenceRequiresDirectory(t *testing.T) {
	memory := nativeEnabledMemory()
	memory.Persistence.Directory = "   "
	err := ValidateMemoryPersistence(memory)
	if err == nil || !strings.Contains(err.Error(), "directory") {
		t.Fatalf("expected directory-required error, got %v", err)
	}
}

// A candidate with an omitted Ranges block persists all allocated areas: valid.
func TestNativePersistenceRangesOmittedValid(t *testing.T) {
	if err := ValidateMemoryPersistence(nativeEnabledMemory()); err != nil {
		t.Fatalf("omitted ranges should be accepted: %v", err)
	}
}

func TestNativePersistenceRangeValidation(t *testing.T) {
	base := func() Memory {
		memory := nativeEnabledMemory()
		memory.Persistence.Ranges = &PersistenceRanges{}
		return memory
	}
	cases := []struct {
		name        string
		mutate      func(*Memory)
		errContains string
	}{
		{"valid subset", func(m *Memory) {
			m.Persistence.Ranges.HoldingRegs = []PersistenceArea{{Start: 10, Count: 5}}
		}, ""},
		{"empty ranges object", func(m *Memory) {}, "cannot be empty"},
		{"zero count", func(m *Memory) {
			m.Persistence.Ranges.Coils = []PersistenceArea{{Start: 0, Count: 0}}
		}, "count must be > 0"},
		{"outside allocated area", func(m *Memory) {
			m.Persistence.Ranges.HoldingRegs = []PersistenceArea{{Start: 250, Count: 10}}
		}, "outside allocated memory"},
		{"unallocated area", func(m *Memory) {
			m.Persistence.Ranges.InputRegs = []PersistenceArea{{Start: 0, Count: 1}}
		}, "not allocated"},
		{"overlapping ranges", func(m *Memory) {
			m.Persistence.Ranges.Coils = []PersistenceArea{{Start: 0, Count: 10}, {Start: 5, Count: 10}}
		}, "overlapping"},
		{"adjacent non-overlapping ok", func(m *Memory) {
			m.Persistence.Ranges.Coils = []PersistenceArea{{Start: 0, Count: 10}, {Start: 10, Count: 10}}
		}, ""},
	}
	for _, tc := range cases {
		memory := base()
		tc.mutate(&memory)
		err := ValidateMemoryPersistence(memory)
		if tc.errContains == "" {
			if err != nil {
				t.Fatalf("%s: expected no error, got %v", tc.name, err)
			}
			continue
		}
		if err == nil || !strings.Contains(err.Error(), tc.errContains) {
			t.Fatalf("%s: expected error containing %q, got %v", tc.name, tc.errContains, err)
		}
	}
}

// Validation is read-only: it never mutates the caller's memory.
func TestNativePersistenceValidationDoesNotMutate(t *testing.T) {
	memory := nativeEnabledMemory()
	memory.Persistence.Ranges = &PersistenceRanges{Coils: []PersistenceArea{{Start: 0, Count: 1}}}
	encoded, _ := yaml.Marshal(memory)
	var before Memory
	if err := yaml.Unmarshal(encoded, &before); err != nil {
		t.Fatal(err)
	}
	if err := ValidateMemoryPersistence(memory); err != nil {
		t.Fatalf("valid memory rejected: %v", err)
	}
	if !reflect.DeepEqual(memory, before) {
		t.Fatalf("validation mutated memory:\n before=%+v\n after =%+v", before, memory)
	}
}

// A root-level persistence key must fail candidate validation.
func TestCandidateRejectsRootPersistence(t *testing.T) {
	cfg := EffectiveConfig{
		Listeners: []Listener{{ID: "main", Listen: ":502", Memory: []Memory{{UnitID: 1, Coils: &Area{Start: 0, Count: 8}}}}},
		Extra:     map[string]interface{}{"persistence": map[string]interface{}{"enabled": true}},
	}
	if err := ValidateCandidatePersistence(cfg); err == nil {
		t.Fatal("expected root persistence to be rejected")
	}
	ok := EffectiveConfig{
		Listeners: []Listener{{ID: "main", Listen: ":502", Memory: []Memory{{UnitID: 1, Coils: &Area{Start: 0, Count: 8}}}}},
	}
	if err := ValidateCandidatePersistence(ok); err != nil {
		t.Fatalf("no root persistence should be accepted: %v", err)
	}
}

// Commit rejects a persistence candidate with an invalid range and preserves the
// existing unrelated configuration byte-for-byte.
func TestCommitRejectsInvalidPersistence(t *testing.T) {
	c := New(t.TempDir(), "simulator")
	prior := validCandidate(61001, 1)
	if err := c.Commit(prior, OwnershipDoc{}); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(c.EffectiveConfigPath())
	if err != nil {
		t.Fatal(err)
	}
	bad := validCandidate(61001, 1)
	bad.Listeners[0].Memory[0].Persistence = &Persistence{
		Enabled:   boolPtr(true),
		Directory: "/var/lib/mma2/unit1",
		Ranges:    &PersistenceRanges{HoldingRegs: []PersistenceArea{{Start: 0, Count: 99}}},
	}
	if err := c.Commit(bad, OwnershipDoc{}); err == nil {
		t.Fatal("expected invalid persistence range to be rejected")
	}
	after, err := os.ReadFile(c.EffectiveConfigPath())
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != string(before) {
		t.Fatal("rejected commit changed the effective config")
	}
}
