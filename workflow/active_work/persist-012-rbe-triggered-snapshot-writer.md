# PERSIST-012 — RBE-Triggered Snapshot Writer

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-011
Next: PERSIST-013

## Primary outcome
Add a persistence runtime that reacts to persistence-owned RBE events and updates the corresponding snapshot without continuous polling.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement startup restore, manifests, unseal, or UI status.

## Acceptance
1. Persistence writer subscribes to the system persistence RBE signal.
2. On an event it obtains the authoritative configured area state and compares against the snapshot image.
3. Only changed bytes/register words are written; unchanged state causes no disk write.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-011. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/1/2/1/1=7; writer only.

## CWAL
PERSIST-011 independently PASSed on tested HEAD `313e3be60a29fff6c967d78c0cf677b8b888ed56` against pinned product checkpoint `190e464f6106c3e640d21b1f9352328447464bef`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
