# PERSIST-018 — Restore Verification

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-017
Next: PERSIST-019

## Primary outcome
Add the restore-completion gate that requires all configured persisted areas and their Raw Ingest acknowledgements/integrity checks to succeed before commit.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not unseal directly in this task.

## Acceptance
1. Complete required-area set is tracked deterministically.
2. Any failed/missing area prevents restore completion.
3. Success is emitted only after every required area is validated and committed.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-017. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/2=6.

## CWAL
PERSIST-017 is delivered on GitHub main at `fa1276c26cbb639f061436c3f111351232037b57`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, require the complete configured persistence area set and successful integrity/Raw Ingest acknowledgement state before reporting restore completion, do not unseal directly, prepare PERSIST-019 for the next invocation only after genuine completion, and STOP.
