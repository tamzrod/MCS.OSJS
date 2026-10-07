# PERSIST-006 — Memory Area Removal Synchronization

Status: ACTIVE — HUMAN ASSIGNED 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-005
Next: PERSIST-007

## Primary outcome
Remove the corresponding derived persistence RBE when an authoritative memory area is removed.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement persistence-disable behavior.

## Acceptance
1. Removing one area removes only its system persistence RBE.
2. Other persisted-area RBE projections remain.
3. User RBE rules remain untouched.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-005. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/1/1/1=4.

## CWAL
PERSIST-005 is delivered on GitHub main at `c743174204455afa83bc92b859fd47582a7afdbc`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
