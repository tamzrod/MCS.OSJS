# PERSIST-007 — Persistence Disable Cleanup

Status: ACTIVE — HUMAN ASSIGNED 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-006
Next: PERSIST-008

## Primary outcome
Remove persistence-owned RBE projections when persistence is disabled for a memory.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not delete or rewrite user-owned RBE entries.

## Acceptance
1. Disabling persistence removes all persistence-owned RBE for that memory.
2. User RBE rules survive unchanged.
3. Re-enabling persistence regenerates projections from current memory ranges.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-006. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5.

## CWAL
PERSIST-006 is delivered on GitHub main at `9ac48cb23e369fb0384b3d8f00811435d5bf7b69`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
