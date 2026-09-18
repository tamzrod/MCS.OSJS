# UMIG-006-T — TEST: Diagnostics Mapping Fixtures

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: TEST / OpenHands (JR)
Previous: UMIG-006
Next: UMIG-006-V

## Primary outcome
Verify diagnostics mapping of healthy, stopped, unavailable and unsupported states.

## Instruction / expected / evidence
Run the exact focused diagnostics test command in the current JR packet using mocked OS.js responses. Expect missing states UNKNOWN/UNAVAILABLE, errors shown and Windows-only controls inert/disabled. Return command/output/exit, case names, HEAD and changed paths; no source edits. Missing tester = BLOCKED, actual assertion failure = FAIL.

## Non-scope
No live service manipulation, rendered UI acceptance or new endpoints.

## Dependencies
UMIG-006 code checkpoint, promotion, current JR packet.

## Sizing
Surface 0, environment 0, behavior 0, verification 1, recovery 0 = 1.
