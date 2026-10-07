# PERSIST-015 — Startup Snapshot Loader

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-014
Next: PERSIST-016

## Primary outcome
Load and validate persistence snapshots during startup while the target MMA2 memory remains sealed.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not perform Raw Ingest writes or unseal.

## Acceptance
1. Loader only acts for persistence-enabled sealed memories.
2. Missing/invalid/incompatible snapshots produce explicit restore state.
3. Loader never exposes partially loaded state through Modbus.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-014. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/1/2/1/2=8; startup loader only.

## Coding evidence (PERSIST-015, OpenHands JR DEV)

- Source checkpoint base: `45ff248b1e35607cb984231ce5bf507096eae174` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed/added paths (product):
  - A `mma2composer/persistence_snapshot_loader.go` (blob `29cf1950cbec796508f1ab3c7c214c6ed5f30208`) — `LoadPersistenceSnapshots` loads and validates every configured persisted area at startup, acting only for persistence-enabled **sealed** memories (otherwise `disabled`/`unsealed`/`empty` with no snapshot read). Each area is checked for compatibility (format version, Port→Unit ID identity, area, start/count) and integrity (length + SHA-256) via the PERSIST-014 manifest before it is marked `ready`. Aggregate state is `ready` only when every configured area is ready, so a partially valid set can never be exposed. No write, Raw Ingest call or unseal occurs. Helpers: `PersistenceSnapshotAreas` (derives the load set from PERSIST-003/009 rules), `snapshotIncompatibility`, `outcomePriority`; interfaces `PersistenceSnapshotSource`, `PersistenceSnapshotArea`, and the `PersistenceRestoreOutcome`/`PersistenceRestorePlan` types.
  - A `mma2composer/persistence_snapshot_loader_test.go` (blob `8ff2147ef88ca29b499e3c07d897e54fc07a47f0`) — loader self-check.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceSnapshot .` → `ok`, exit 0 (all `TestPersistenceSnapshotLoader*`/`TestPersistenceSnapshotAreas*` cases pass).
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) loader acts only for persistence-enabled sealed memories — non-enabled returns `disabled` and unsealed returns `unsealed`, both without reading any snapshot; (2) missing/invalid/incompatible snapshots produce distinct explicit restore states (`missing`/`invalid`/`incompatible`) and never `ready`; (3) the loader performs no Raw Ingest write and no unseal, and the plan is `ready` only when every configured area is ready, so no partially loaded state can reach Modbus.
- Design boundary: Raw Ingest restore transport is PERSIST-016, seal-flag protection PERSIST-017, restore verification PERSIST-018 and unseal PERSIST-019, so this task defines load/validate only and performs no memory mutation. Placement is the shared schema owner `mma2composer` (consistent with PERSIST-012..014).
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: recorded in the repository commit for this task.

## CWAL
PERSIST-014 is delivered on GitHub main at `98a952a89b087f12a9b99b6bd648ce97ac44e6f2`. PERSIST-015 CODE is now complete and no longer ACTIVE; PERSIST-016 is the next eligible CODE packet. This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
