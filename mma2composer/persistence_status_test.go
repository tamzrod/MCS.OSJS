package mma2composer

import "testing"

// PERSIST-021 self-check: a committed restore reports healthy, unsealed and no
// failure, and never claims control authority.
func TestPersistenceRuntimeStatusCommitted(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	plan.SealingFlag = &PersistenceSealingFlag{Address: 4}
	result, err := RestorePersistencePlan(plan, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Committed {
		t.Fatalf("fixture must commit: %+v", result)
	}
	status := PersistenceRuntimeStatusFromPlan(plan, result, nil, nil)
	if !status.Configured || !status.Healthy || status.Sealed {
		t.Fatalf("committed status wrong: %+v", status)
	}
	if status.RestoreOutcome != "none" || status.SnapshotHealth != "ready" {
		t.Fatalf("committed status fields wrong: %+v", status)
	}
	if status.LastRestore != nil || status.LastSave != nil {
		t.Fatalf("absent observations must be omitted: %+v", status)
	}
}

// PERSIST-021 self-check: a restore failure is clearly visible while the sealed
// state remains authoritative.
func TestPersistenceRuntimeStatusFailureStaysSealed(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	plan.SealingFlag = &PersistenceSealingFlag{Address: 4}
	writer := &fakeRawIngestWriter{responses: map[PersistenceRawIngestArea]byte{RawIngestHoldingRegisters: 0x21}}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	status := PersistenceRuntimeStatusFromPlan(plan, result, nil, nil)
	if !status.Configured {
		t.Fatalf("persistence is configured: %+v", status)
	}
	if status.Healthy || !status.Sealed {
		t.Fatalf("failed restore must stay sealed and unhealthy: %+v", status)
	}
	if status.RestoreOutcome != "raw_ingest_response" {
		t.Fatalf("failure must be visible in the status: %+v", status)
	}
	if status.SnapshotHealth != "ready" {
		t.Fatalf("snapshot health is independent of the restore outcome: %+v", status)
	}
}

// A missing/incompatible snapshot is visible in both snapshot health and the
// restore outcome, and the memory stays sealed.
func TestPersistenceRuntimeStatusSnapshotFailureVisible(t *testing.T) {
	plan := brokenLoaderPlan(t, func(s *fakeSnapshotSource) { delete(s.snapshots, "coils") })
	result, err := RestorePersistencePlan(plan, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	status := PersistenceRuntimeStatusFromPlan(plan, result, nil, nil)
	if status.SnapshotHealth != "missing" || status.RestoreOutcome != "missing_snapshot" {
		t.Fatalf("snapshot failure must be visible: %+v", status)
	}
	if status.Healthy || !status.Sealed {
		t.Fatalf("snapshot failure must stay sealed: %+v", status)
	}
}

// A disabled device reports not-configured and is not presented as a health
// failure.
func TestPersistenceRuntimeStatusDisabled(t *testing.T) {
	status := PersistenceRuntimeStatusConfigured(false)
	if status.Configured || status.Healthy || !status.Sealed {
		t.Fatalf("disabled status wrong: %+v", status)
	}
	if status.SnapshotHealth != "" || status.RestoreOutcome != "" {
		t.Fatalf("disabled status must not fabricate health: %+v", status)
	}
}

// Enabled-but-unobserved persistence is conservative: configured, not healthy,
// still sealed, with no fabricated snapshot/restore health.
func TestPersistenceRuntimeStatusConfiguredButUnobserved(t *testing.T) {
	status := PersistenceRuntimeStatusConfigured(true)
	if !status.Configured || status.Healthy || !status.Sealed {
		t.Fatalf("unobserved configured status wrong: %+v", status)
	}
	if status.SnapshotHealth != "" || status.RestoreOutcome != "" {
		t.Fatalf("unobserved status must not fabricate health: %+v", status)
	}
}

// Last save/restore observations are reported when present and are copied, not
// aliased.
func TestPersistenceRuntimeStatusObservations(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	plan.SealingFlag = &PersistenceSealingFlag{Address: 4}
	result, err := RestorePersistencePlan(plan, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	save := &PersistenceTimestamp{At: "2026-10-08T00:00:00Z"}
	restore := PersistenceRestoreObservationFromResult(result)
	status := PersistenceRuntimeStatusFromPlan(plan, result, save, &restore)
	if status.LastSave == nil || status.LastSave.At != save.At {
		t.Fatalf("last save must be reported: %+v", status)
	}
	if status.LastRestore == nil || status.LastRestore.Failure != "none" || !status.LastRestore.Committed {
		t.Fatalf("last restore must be reported: %+v", status)
	}
	// Mutating the caller's observation must not change the reported copy.
	restore.Failure = "mutated"
	if status.LastRestore.Failure != "none" {
		t.Fatal("last restore must be a copy, not an alias")
	}
}

// The projection is observational: it returns the same values regardless of how
// often it is called and never mutates the plan or result.
func TestPersistenceRuntimeStatusIsPureProjection(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	plan.SealingFlag = &PersistenceSealingFlag{Address: 4}
	result, err := RestorePersistencePlan(plan, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	beforePlan, beforeResult := plan, result
	first := PersistenceRuntimeStatusFromPlan(plan, result, nil, nil)
	second := PersistenceRuntimeStatusFromPlan(plan, result, nil, nil)
	if first != second {
		t.Fatalf("projection must be deterministic: %+v vs %+v", first, second)
	}
	if plan.Enabled != beforePlan.Enabled || plan.State != beforePlan.State || result.Committed != beforeResult.Committed {
		t.Fatal("projection must not mutate its inputs")
	}
}
