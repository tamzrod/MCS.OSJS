# PERSIST-013 — Snapshot File Format

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
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
PERSIST-012 is delivered on GitHub main at `2f18f8196f5497c5435b697b8b1570f59cab2642`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
