# PERSIST-003 — Derived Persistence RBE Generation

Status: ACTIVE — HUMAN ASSIGNED 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-002
Next: PERSIST-004

## Primary outcome
Generate one system-owned persistence RBE rule per configured memory area from the authoritative memory layout.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement UI locking, user-rule collision policy, or disk persistence.

## Acceptance
1. Each present area gets exactly one derived persistence RBE.
2. Derived start/count equal the area's current start/count.
3. Persistence does not introduce independent editable range fields.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-002. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/1=6; bounded to derivation only.

## CWAL
PERSIST-002 is delivered on GitHub main at `fe3bda9e5c863c959500514e8c6083d2514526d1`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
