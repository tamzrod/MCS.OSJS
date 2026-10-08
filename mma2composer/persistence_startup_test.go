package mma2composer

import "testing"

// PERSIST-R03 self-check: the startup orchestrator loads validated snapshots
// while sealed, attaches the authoritative sealing flag from config, restores
// via Raw Ingest, and commits only after verification.
func TestPersistenceRestoreAtStartupCommits(t *testing.T) {
	key, areas, source := loaderFixture(t)
	rules := make([]PersistenceRBERule, 0, len(areas))
	for i, a := range areas {
		rules = append(rules, PersistenceRBERule{ID: uint8(i + 1), Area: a.Area, Start: a.Start, Count: a.Count, SystemOwned: true})
	}
	mem := PersistenceStartupMemory{
		Key:     key,
		Enabled: true,
		Sealed:  true,
		Rules:   rules,
		Extra:   map[string]interface{}{"state_sealing": map[string]interface{}{"enabled": true, "area": "coil", "address": 0}},
	}
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistenceAtStartup(mem, source, writer)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Result.Completed || !result.Result.Committed {
		t.Fatalf("startup restore must commit: %+v", result.Result)
	}
	if result.Plan.SealingFlag == nil || result.Plan.SealingFlag.Address != 0 {
		t.Fatalf("sealing flag must come from config: %+v", result.Plan.SealingFlag)
	}
}

// A missing snapshot at startup leaves the memory sealed with the explicit
// missing classification and writes nothing.
func TestPersistenceRestoreAtStartupFailClosed(t *testing.T) {
	key, areas, source := loaderFixture(t)
	delete(source.snapshots, "coils")
	rules := make([]PersistenceRBERule, 0, len(areas))
	for i, a := range areas {
		rules = append(rules, PersistenceRBERule{ID: uint8(i + 1), Area: a.Area, Start: a.Start, Count: a.Count, SystemOwned: true})
	}
	mem := PersistenceStartupMemory{
		Key:     key,
		Enabled: true,
		Sealed:  true,
		Rules:   rules,
		Extra:   map[string]interface{}{"state_sealing": map[string]interface{}{"area": "coil", "address": 0}},
	}
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistenceAtStartup(mem, source, writer)
	if err != nil {
		t.Fatal(err)
	}
	if result.Result.Committed || !result.Result.Sealed {
		t.Fatalf("missing snapshot must stay sealed: %+v", result.Result)
	}
	if result.Result.Failure != PersistenceRestoreFailureMissingSnapshot {
		t.Fatalf("classification wrong: %+v", result.Result)
	}
	if len(writer.calls) != 0 {
		t.Fatal("no write may occur for a missing snapshot")
	}
}

// A non-sealed startup precondition blocks restore with the unsealed reason.
func TestPersistenceRestoreAtStartupUnsealedRefused(t *testing.T) {
	key, areas, source := loaderFixture(t)
	rules := make([]PersistenceRBERule, 0, len(areas))
	for i, a := range areas {
		rules = append(rules, PersistenceRBERule{ID: uint8(i + 1), Area: a.Area, Start: a.Start, Count: a.Count, SystemOwned: true})
	}
	mem := PersistenceStartupMemory{Key: key, Enabled: true, Sealed: false, Rules: rules}
	result, err := RestorePersistenceAtStartup(mem, source, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Result.Committed || !result.Result.Sealed {
		t.Fatalf("an unsealed precondition must not commit: %+v", result.Result)
	}
}
