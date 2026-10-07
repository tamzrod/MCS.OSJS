# PERSIST-005 — Memory Range Synchronization

Status: ACTIVE — HUMAN ASSIGNED 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-004
Next: PERSIST-006

## Primary outcome
Keep each persistence-owned RBE projection synchronized when its authoritative memory area's start/count changes.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not alter user-owned RBE rules or snapshot files.

## Acceptance
1. Changing area start updates derived persistence RBE start.
2. Changing area count updates derived persistence RBE count.
3. No second user edit is required.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-004. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5; one synchronization path.

## CWAL
PERSIST-004 is delivered on GitHub main at `936ab24560d57cbf8aacc8180bc00a2527dcfad5`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
