# PERSIST-008 — User RBE Compatibility

Status: ACTIVE — HUMAN ASSIGNED 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-007
Next: PERSIST-009

## Primary outcome
Preserve independent user-created RBE rules, including overlap with persistence-owned ranges.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not redefine RBE wire semantics.

## Acceptance
1. User rules may coexist with system persistence rules.
2. Overlapping ranges are accepted when otherwise valid.
3. Persistence lifecycle operations do not mutate user rules.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-007. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5.

## CWAL
PERSIST-007 is delivered on GitHub main at `feb1e622d38cf8107a1df9ee182e5df17f95790f`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
