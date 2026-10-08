# PERSIST-RUNTIME — Persistence Runtime Integration

Status: SUPERSEDED — DECOMPOSED AND HUMAN-PROMOTED 2026-10-08
Stage: CODE
Owner: unassigned
Predecessor: PERSIST-021
Successor: PERSIST-022

## Why this exists

PERSIST-012..021 implemented and self-checked the persistence contracts, codecs, manifest validation, loader, restore gate, seal protection, final unseal, failure classification and observational status. Before independent PERSIST-022 verification, repository call-site inspection at source checkpoint `1310c1367c7f98f084187006084d624d8e20e07b` found that the lifecycle is not yet wired into a running appliance path:

- `NewPersistenceSnapshotWriter` has no non-test runtime caller.
- `LoadPersistenceSnapshots` has no non-test runtime caller.
- `RestorePersistencePlan` has no non-test runtime caller.
- `PersistenceRuntimeStatusFromPlan` has no non-test runtime caller.
- persistence snapshot store/source contracts have test fakes but no observed concrete appliance filesystem implementation connecting them to startup/RBE.

Therefore PERSIST-022 cannot honestly verify the promoted lifecycle "save state → restart sealed → restore → verify → unseal → Modbus sees restored state" through the actual runtime.

## Primary outcome

Wire the already-implemented persistence contracts into the authoritative simulator/orchestrator runtime without changing MMA2 memory authority or protocol semantics.

## Scope

1. Add the concrete appliance persistence store/source for raw snapshot payloads plus manifest metadata under the appliance data root.
2. Connect persistence-owned RBE events to `PersistenceSnapshotWriter` so runtime changes persist without continuous polling.
3. On startup, for persistence-enabled memories:
   - remain sealed;
   - load and validate configured snapshots;
   - restore through existing Raw Ingest v1;
   - verify the complete required-area set;
   - unseal only through the existing final commit step after verification success;
   - remain sealed with the existing deterministic failure classification on any failure.
4. Feed actual observed save/restore state and timestamps into `PersistenceRuntimeStatusFromPlan`; do not report configured-only placeholder health after a restore attempt exists.
5. Add a repository-native disposable integration test/runner suitable for independent PERSIST-022 that proves:
   - a value is saved;
   - the runtime restarts sealed;
   - Modbus is rejected while sealed;
   - snapshot restore completes;
   - final unseal occurs last;
   - Modbus reads the restored value afterward;
   - at least one failed restore remains sealed.

## Non-scope

- No new persistence authority or alternate seal flag.
- No Raw Ingest protocol change.
- No historian/telemetry database behavior.
- No production/customer paths or global service operations.
- No automatic repair/retry loop.
- No ICC edits.

## Acceptance

1. The runtime has real non-test call sites for save, startup load/restore, and runtime status projection.
2. Filesystem snapshot + manifest read/write is deterministic and rooted only in the appliance data root.
3. Startup restore remains sealed until the existing verified final commit/unseal step.
4. Failure stays sealed and reports the existing classified reason.
5. A committed disposable integration harness exists and can be executed once by independent JR without source modification.
6. PERSIST-022 can be pinned with exact commands against that harness.

## Evidence / handoff

Coding owner records exact changed paths, targeted tests, full bounded regression, resulting source checkpoint, and the exact independent-JR command that exercises the committed disposable integration harness. Do not claim PERSIST-022 PASS.

## Sizing

3/1/3/2/2=11; runtime integration + concrete filesystem adapter + committed disposable integration harness.

## Superseded by
Human decomposed and promoted this oversized proposal into:
- `PERSIST-R01 — Filesystem Snapshot Adapter`
- `PERSIST-R02 — Runtime Save Wiring`
- `PERSIST-R03 — Startup Restore Wiring`
- `PERSIST-UI01 — Electron Persistence Settings`
- `PERSIST-R04 — Runtime Status Wiring`
- `PERSIST-R05 — Disposable Persistence E2E Harness`

Do not execute this parent proposal. The active selector is `handoff.md`.
