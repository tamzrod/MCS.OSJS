# PERSIST-015 — Startup Snapshot Loader

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
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

## CWAL
PERSIST-014 is delivered on GitHub main at `98a952a89b087f12a9b99b6bd648ce97ac44e6f2`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. If a reused workspace still has local-only PERSIST-013/PERSIST-014 transport-blocked commits, treat them as superseded duplicate history and reconcile them with bounded safe Git mechanics; do not replay them as new product work. Execute exactly PERSIST-015, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
