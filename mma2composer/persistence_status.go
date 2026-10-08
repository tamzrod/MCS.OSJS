package mma2composer

// PersistenceRuntimeStatus is the read-only, observational persistence health of
// one memory for the operator surface. It is derived entirely from the
// authoritative persistence enablement and restore/sealing state; it never
// controls, retries or bypasses restore or sealing.
//
// Health is reported as two independent facts so a failure is always visible
// while sealing stays authoritative:
//   - Configured: persistence is enabled for this memory. When false, the
//     restore/health fields are not applicable and are empty.
//   - Healthy: true only when a restore actually committed (validated, restored
//     and unsealed by the explicit commit step). It is false whenever restore
//     has not run or did not commit.
//
// Sealed is the authoritative current sealing state: a memory that is not
// successfully committed remains sealed.
type PersistenceRuntimeStatus struct {
	Configured bool `json:"configured"`

	// Sealed reports the authoritative current sealing state. It stays
	// authoritative even when restore fails: a failed restore leaves the memory
	// sealed.
	Sealed bool `json:"sealed"`

	// Healthy is true only when the restore committed (Completed and Committed).
	Healthy bool `json:"healthy"`

	// SnapshotHealth is the loader's aggregate snapshot state ("ready",
	// "missing", "invalid", "incompatible", "disabled", "unsealed", "empty").
	SnapshotHealth string `json:"snapshot_health"`

	// RestoreOutcome is the classified restore failure ("none" when committed),
	// so a failure is clearly visible while the sealed state stays authoritative.
	RestoreOutcome string `json:"restore_outcome"`

	// LastSave is the most recent RBE-triggered snapshot save, when observed.
	LastSave *PersistenceTimestamp `json:"last_save,omitempty"`

	// LastRestore is the most recent restore attempt, when observed.
	LastRestore *PersistenceRestoreObservation `json:"last_restore,omitempty"`
}

// PersistenceTimestamp is a recorded instant in RFC3339Nano form.
type PersistenceTimestamp struct {
	At string `json:"at"`
}

// PersistenceRestoreObservation is the observed result of a restore attempt: the
// classification and whether it committed. It is observational only and never
// grants control authority.
type PersistenceRestoreObservation struct {
	Failure   string `json:"failure"`
	Committed bool   `json:"committed"`
}

// PersistenceRestoreObservationFromResult derives the observed restore outcome
// from a restore result. It is a pure projection.
func PersistenceRestoreObservationFromResult(result PersistenceRestoreResult) PersistenceRestoreObservation {
	return PersistenceRestoreObservation{Failure: result.Failure.String(), Committed: result.Committed}
}

// PersistenceRuntimeStatusConfigured reports persistence for a memory that has
// not produced a restore plan/result in this observation. It reports only the
// configuration fact and the conservative sealed default; snapshot/restore
// fields are not applicable and are left empty. It is a pure projection.
func PersistenceRuntimeStatusConfigured(configured bool) PersistenceRuntimeStatus {
	// Snapshot/restore health is unknown (empty) because no restore plan/result
	// has been observed yet; the memory is conservatively reported sealed.
	return PersistenceRuntimeStatus{Configured: configured, Sealed: true}
}

// PersistenceRuntimeStatusFromPlan projects the observational persistence status
// for one memory from its authoritative restore plan/result. It is a pure
// read-only projection: it never writes, retries, unseals or changes sealing.
// Snapshot health and the restore outcome come from the plan/result; the sealed
// flag is authoritative and a failed restore keeps the memory sealed. lastSave
// and lastRestore are optional observed facts and are reported only when present.
func PersistenceRuntimeStatusFromPlan(plan PersistenceRestorePlan, result PersistenceRestoreResult, lastSave *PersistenceTimestamp, lastRestore *PersistenceRestoreObservation) PersistenceRuntimeStatus {
	status := PersistenceRuntimeStatus{
		Configured:     plan.Enabled,
		Sealed:         plan.Sealed && !result.Committed,
		Healthy:        result.Completed && result.Committed,
		SnapshotHealth: plan.State.String(),
		RestoreOutcome: result.Failure.String(),
		LastSave:       lastSave,
	}
	if lastRestore != nil {
		copy := *lastRestore
		status.LastRestore = &copy
	}
	return status
}
