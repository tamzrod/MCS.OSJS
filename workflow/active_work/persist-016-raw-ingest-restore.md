# PERSIST-016 — Raw Ingest Restore

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: PERSIST-015
Next: PERSIST-017

## Primary outcome
Restore validated snapshot area payloads into the matching MMA2 memory through the existing Raw Ingest v1 contract.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not unseal memory or change Raw Ingest protocol.

## Acceptance
1. Each snapshot area restores to the same area/start/count identity.
2. Every Raw Ingest response is checked; non-zero response aborts restore.
3. No cross-area mirroring is added by persistence.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-015. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/1/2/1/2=8; restore transport only.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
