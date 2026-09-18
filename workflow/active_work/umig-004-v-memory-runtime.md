# UMIG-004-V — VERIFY: Live Memory Behavior

Status: QUEUED — promoted 2026-09-18; wait for UMIG-004-T PASS and safe test target.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-004-T
Next: UMIG-005

## Primary outcome
Observe Memory load/apply/selected-device status against the existing Simulator runtime.

## Instruction / expected / evidence
In an isolated test instance with explicit disposable config or a verified backup/restore plan, launch OS.js Toolkit Memory, load a known device, modify one safe test value/config, Save & Apply and observe selected status, None/Random persistence and error/unavailable handling. Never modify production/customer configuration. Expected: actual backend acknowledgement, retained configuration and truthful status. Return test target, pre/post config evidence, UI screenshots/logs, actual responses, safety cleanup and HEAD. No safe live target = BLOCKED.

## Non-scope
No production resets, Replicator tests, backend fixes or final overall acceptance.

## Dependencies
UMIG-004-T PASS, verified safe test target and explicit current JR packet; the operator's working Docker deployment is not permission to write its persistent data.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
