package mma2composer

// PersistenceStartupMemory describes one persistence-enabled memory at runtime
// startup. Rules are the authoritative system-owned persistence RBE rules;
// Extra is the configured memory block (used only to locate the authoritative
// State Sealing flag). Sealed is the startup precondition: persistence-enabled
// memory starts sealed, so restore may proceed.
type PersistenceStartupMemory struct {
	Key     PersistenceMemoryKey
	Enabled bool
	Sealed  bool
	Rules   []PersistenceRBERule
	Extra   map[string]interface{}
}

// PersistenceStartupResult is the observed outcome of one memory's startup
// restore: the loaded plan and the restore result (including the deterministic
// failure classification and sealed state).
type PersistenceStartupResult struct {
	Key    PersistenceMemoryKey
	Plan   PersistenceRestorePlan
	Result PersistenceRestoreResult
}

// RestorePersistenceAtStartup is the real startup orchestration for one
// persistence-enabled memory: it derives the configured load set from the
// authoritative rules, loads and validates the durable snapshots while the
// memory is sealed (LoadPersistenceSnapshots), attaches the authoritative State
// Sealing flag from configuration, then restores through the existing Raw Ingest
// v1 contract (RestorePersistencePlan) — verifying the full required-area set
// and performing the existing final explicit unseal only after success.
//
// It adds no new seal flag, no alternate unseal path and no retry. This pass
// keeps the memory sealed throughout: the loader reads nothing when unsealed,
// the restore forces the sealing bit sealed during area writes, and the final
// unseal runs only after verification succeeds. Any failure returns the plan and
// a sealed, deterministically-classified result rather than a fabricated default.
func RestorePersistenceAtStartup(mem PersistenceStartupMemory, source PersistenceSnapshotSource, writer PersistenceRawIngestWriter) (PersistenceStartupResult, error) {
	areas, err := PersistenceSnapshotAreas(mem.Rules)
	if err != nil {
		return PersistenceStartupResult{}, err
	}
	plan, err := LoadPersistenceSnapshots(mem.Key, mem.Enabled, mem.Sealed, areas, source)
	if err != nil {
		return PersistenceStartupResult{}, err
	}
	if flag, ok := PersistenceSealingFlagFromExtra(mem.Extra); ok {
		plan.SealingFlag = &flag
	}
	result, err := RestorePersistencePlan(plan, writer)
	return PersistenceStartupResult{Key: mem.Key, Plan: plan, Result: result}, err
}
