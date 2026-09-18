# UMIG-005-T — TEST: Replicator Adapter Contract

Status: QUEUED — promoted 2026-09-18; wait for UMIG-005 source checkpoint.
Stage / owner: TEST / OpenHands (JR)
Previous: UMIG-005
Next: UMIG-005-V

## Primary outcome
Prove Replicator load/apply/status/suggest adapter mappings and fail-closed behavior via fixtures.

## Instruction / expected / evidence
Run the exact focused adapter tests supplied in the active JR packet. Include success, backend ownership collision, apply failure, missing device, per-block status and unavailable/unknown COMMS cases. Expected: only one backend apply transaction, genuine error propagation and no fabricated green. Return command/output/exit, cases, HEAD and changed paths; BLOCKED if tests unavailable, FAIL if behavior contradicts expected.

## Non-scope
No real device writes, live acceptance, source fixes or UI redesign.

## Dependencies
UMIG-005 code checkpoint and current JR packet.

## Sizing
Surface 0, environment 0, behavior 0, verification 1, recovery 0 = 1.
