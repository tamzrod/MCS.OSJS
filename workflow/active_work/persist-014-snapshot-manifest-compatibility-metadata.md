# PERSIST-014 — Snapshot Manifest / Compatibility Metadata

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: PERSIST-013
Next: PERSIST-015

## Primary outcome
Add snapshot metadata sufficient to reject incompatible or corrupted snapshots before restore.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not restore or unseal memory.

## Acceptance
1. Manifest binds snapshot to Port, Unit ID, area, start, count and format version.
2. Integrity metadata detects corrupted/incomplete area payloads.
3. Changed incompatible memory layout causes explicit rejection rather than silent truncation/remap.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-013. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/2=7; metadata/validation only.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
