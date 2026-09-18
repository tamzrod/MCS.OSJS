# UMIG-005-V — VERIFY: Live Replicator Behavior

Status: QUEUED — promoted 2026-09-18; wait for UMIG-005-T PASS and safe runtime.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-005-T
Next: UMIG-006

## Primary outcome
Observe Toolkit Replicator load/apply/status/suggest using an actual test runtime.

## Instruction / expected / evidence
Use only disposable device endpoints and backed-up test configs. Load existing test definitions, inspect destination suggestion/ownership, Save & Apply a safe test change, observe real polling/status/error recovery and restore test state as instructed. Expect a single authoritative transaction, preserved foreign reservations, actual responses and no false health. Record test topology, commands/actions, before/after config, backend messages, UI/logs, cleanup and HEAD. Missing safe runtime or unsatisfied backend prerequisite = BLOCKED.

## Non-scope
No production config changes, protocol fixes, Windows COMMS acceptance or workflow advancement.

## Dependencies
UMIG-005-T PASS, safe runtime and current JR packet. Never use the operator's persistent Docker data for destructive testing.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
