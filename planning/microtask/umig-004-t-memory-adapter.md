# UMIG-004-T — TEST: Memory Adapter Contract

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: TEST / OpenHands (JR)
Previous: UMIG-004
Next: UMIG-004-V

## Primary outcome
Validate Simulator/Memory load, apply, status and error response mapping using deterministic fixtures.

## Instruction / expected / evidence
Run the coding agent's exact focused Memory adapter test command from `JR TEST TASK` in a disposable checkout. Exercise load, successful/failed apply, status unavailable and None/Random fixture cases. Expected: matching backend shapes, retained values, real errors, no import of the retired UI. Return exact command, output, exit, test-case names, HEAD and changed paths. Missing test definition or environment = BLOCKED; observed contradiction = FAIL.

## Non-scope
No live configuration writes, code fixes, other tabs or manual UI acceptance.

## Dependencies
UMIG-004 code checkpoint, human promotion, complete JR packet.

## Sizing
Surface 0, environment 0, behavior 0, verification 1, recovery 0 = 1.
