# PERSIST-009 — RBE ID Collision Handling

Status: ACTIVE — HUMAN ASSIGNED 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-008
Next: PERSIST-010

## Primary outcome
Allocate persistence-owned RBE IDs without collision while preserving the existing RBE v1 ID contract.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not redesign RBE v1 or widen its one-byte ID space.

## Acceptance
1. System and user RBE IDs are globally unique as required by the current contract.
2. Users cannot directly assign/change system-owned IDs.
3. No arbitrary permanent user/system ID partition is introduced unless current implementation requires and documents it.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-008. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/1/2/1/2=7; collision policy only.

## CWAL
PERSIST-008 is delivered on GitHub main at `4783e0d6e1d3dbd55b29b8647a115905eb5bc285`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
