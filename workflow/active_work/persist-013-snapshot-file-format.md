# PERSIST-013 — Snapshot File Format

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: PERSIST-012
Next: PERSIST-014

## Primary outcome
Define and implement deterministic raw snapshot encoding for all supported Modbus memory areas.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add compatibility manifest or restore sequencing.

## Acceptance
1. Coils/discrete inputs use LSB-first packed bits.
2. Holding/input registers use big-endian uint16 words.
3. Encoding/decoding is deterministic for configured start/count.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-012. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
