# PERSIST-004 — Locked System RBE Behavior

Status: ACTIVE — HUMAN ASSIGNED 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-003
Next: PERSIST-005

## Primary outcome
Represent persistence RBE rules as system-owned and non-editable/non-deletable through supported configuration/UI paths.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add new persistence runtime behavior.

## Acceptance
1. System persistence RBE is visibly identified as system-owned/locked.
2. User edit/delete attempts cannot mutate the system-derived range or ownership.
3. Ordinary user RBE behavior is unchanged.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-003. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/1=6; ownership/lock behavior only.

## CWAL
PERSIST-003 is delivered on GitHub main at `4e8f774eb946cc9a946938e4f3a28eaa13de5c9d`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
