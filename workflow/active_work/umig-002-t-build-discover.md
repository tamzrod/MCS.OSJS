# UMIG-002-T — TEST: Toolkit Build and Discovery

Status: QUEUED — human authorized independent OpenHands testing; not run.
Stage / owner: TEST / OpenHands (JR)
Previous: UMIG-002
Next: UMIG-002-V

## Primary outcome
Demonstrate that the standalone Toolkit package builds and is discovered by OS.js.

## Scope / safe setup
Use the committed UMIG-002 source in a disposable OS.js checkout. Prepare local Node dependencies if safely available; do not touch Windows Electron or existing user data. Do not alter product source or manifests.

## Test instruction
Command: `cd OSJS && npm run build:local-packages && npm run package:discover`.
Expected: both commands exit 0; the `MCSModbusToolkit` package has built output and appears in generated/discovered OS.js metadata, with no Electron build or installation required.
Evidence: HEAD, `git status --short` before/after, exact command output and exit codes, the discovered Toolkit entry and artifact paths. If environment/build prerequisites prevent execution, report BLOCKED; an actual build failure is FAIL. Do not substitute syntax/source inspection.

## Non-scope
No UI launch, backend/runtime tests, source fixes, task-state changes or ICC writes.

## Dependencies
Verified source-only UMIG-002 coding checkpoint and active `JR TEST TASK` in `handoff.md`.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
