# PERSIST-018 — Restore Verification

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
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
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
