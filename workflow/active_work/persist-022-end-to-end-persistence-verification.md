# PERSIST-022 — End-to-End Persistence Verification

Status: BLOCKED — runtime lifecycle not wired; exact VERIFY packet cannot yet be pinned
Stage: VERIFY
Owner: OpenHands / independent JR
Previous: PERSIST-021
Next: none

## Primary outcome
Independently verify the complete safe persistence lifecycle on a disposable configuration: save state → restart sealed → restore → verify → unseal → Modbus sees restored state.

## Scope
The current handoff must pin the exact source checkpoint, disposable target, preflight, exact ordered actions/commands, expected observations, evidence destination and report permissions before this packet can be selected.

## Non-scope
No production/customer data, no source fixes, no invented commands, no global service actions.

## Acceptance
1. Execute only the exact safe disposable runtime/actions pinned in the current handoff and capture raw evidence.
2. Observe that Modbus access remains sealed until restoration completes and restored values are correct after unseal.
3. Exercise at least one authorized failure case showing failed restore remains sealed, if included in the promoted exact packet.

## Evidence / handoff
JR changes no product source, follows `operation cwal.md`, captures genuine evidence and STOPS after verdict.

## Dependencies
Requires genuine completion evidence for PERSIST-021, a complete current VERIFY packet in `handoff.md`, and an actual runtime path capable of executing the lifecycle.

Current blocker at product checkpoint `1310c1367c7f98f084187006084d624d8e20e07b`: repository call-site inspection found no non-test runtime caller of `NewPersistenceSnapshotWriter`, `LoadPersistenceSnapshots`, `RestorePersistencePlan`, or `PersistenceRuntimeStatusFromPlan`. The concrete appliance snapshot store/source and startup/RBE wiring required for true end-to-end execution are therefore not yet observable in source.

Review-stage proposal: `workflow/micro_task/persist-runtime-integration-gap.md`.

## Sizing
0/2/0/2/2=6.

## CWAL
This VERIFY packet remains human-promoted but is **BLOCKED**, not executable. Do not run independent JR yet: the runtime currently lacks observable non-test persistence lifecycle wiring, so the exact disposable end-to-end execution packet cannot be truthfully pinned. Do not substitute unit tests for the required runtime lifecycle and do not invent a test runner.
