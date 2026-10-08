# PERSIST-019 — Atomic Unseal / Commit Step

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-018
Next: PERSIST-020

## Primary outcome
Make the final successful restore action an explicit write of the existing State Sealing flag to 1.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add runtime reseal or alternate commit flags.

## Acceptance
1. Unseal occurs only after restore verification success.
2. The flag location comes only from current State Sealing configuration.
3. No earlier restore step can expose memory to Modbus.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-018. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/2=6.

## CWAL
PERSIST-018 is delivered on GitHub main at `3eb118f507688d483c5d4985f74fe6a98b531bc3`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task: unseal only after `VerifyPersistenceRestore` success, write the existing authoritative State Sealing flag to 1 as the final commit action, derive its location only from current State Sealing configuration, add no runtime reseal or alternate commit flag, prepare PERSIST-020 for the next invocation only after genuine completion, and STOP.
