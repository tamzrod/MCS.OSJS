# UMIG-002-T — TEST: Toolkit Build and Discovery

Status: QUEUED — prior run FAIL; awaiting authorized CODE repair and a fresh JR packet.
Stage / owner: TEST / OpenHands (JR via Operation CWAL)
Previous: UMIG-002
Next: UMIG-002-V

## Primary outcome
Demonstrate that the standalone Toolkit package builds and is discovered by OS.js.

## Scope / safe setup
Use committed UMIG-002 source in a disposable OS.js checkout. Prepare sandbox-local Node dependencies safely if needed, without changing tracked product source, manifest or configuration. Leave Windows Electron and user data untouched. Check the HEAD and checkout state before running.

## Test instruction
Exact product test command: `cd OSJS && npm run build:local-packages && npm run package:discover`.
Expected: both commands exit 0; `MCSModbusToolkit` has `dist/main.js` and `dist/main.css` and appears in discovered OS.js metadata; Electron is not required.
Evidence: HEAD, `git status --short` before and after, raw command output and exit codes, discovered Toolkit metadata entry and artifact paths. A genuine build failure is FAIL; an unsafe or unavailable test environment is BLOCKED. Source inspection alone is not PASS.

## Recorded failure / repair interruption
JR executed the original packet and reported FAIL in `handoff.md` at commit `752a541`: build and discovery commands exited 0 and Toolkit assets existed, but Toolkit was absent from discovery manifests and `dist/apps/`. Toolkit lacks the `package.json` discovery marker required by OS.js. The human authorized a bounded CODE repair (`UMIG-002-R`) on 2026-09-18. This TEST is temporarily QUEUED, not passed; resume only after repair checkpoint, ICC review, and replacement JR packet. The repair is an interruption/retry, not a change to the original UMIG-002 → UMIG-002-T → UMIG-002-V feature sequence.

## Non-scope
No rendered UI launch, backend/runtime tests, source fixes, workflow advancement or ICC writes.

## Dependencies
UMIG-002 source checkpoint archived; repair UMIG-002-R must be completed before this TEST becomes ACTIVE again. Required ICC refresh is performed by BLACK SHEEP WALL, not JR, before execution.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
