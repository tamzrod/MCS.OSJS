# UMIG-002-T — TEST: Toolkit Build and Discovery

Status: ACTIVE — retest pending after source-only discovery-manifest repair; no PASS yet.
Stage / owner: TEST / OpenHands (JR via Operation CWAL)
Previous: UMIG-002
Next: UMIG-002-V

## Primary outcome
Demonstrate that the standalone Toolkit package builds and is discovered by OS.js.

## Scope / safe setup
Use committed UMIG-002 source including repair `cd67e150139b60f3914db8c6f1998c96c5b8da07` in a disposable OS.js checkout. Prepare sandbox-local Node dependencies safely if needed, without changing tracked product source, manifest or configuration. Leave Windows Electron and user data untouched. Check HEAD and checkout state before running. Fulfill the bounded ICC prerequisite in `handoff.md` before JR execution.

## Test instruction
The current `JR TEST TASK — CURRENT: UMIG-002-T (RETEST)` in `handoff.md` is the exact packet. Execute `cd OSJS && npm run build:local-packages` and `cd OSJS && npm run package:discover` in order, recording each actual exit code and raw output. Inspect Toolkit `dist/main.js` and `dist/main.css` plus discovery output, `OSJS/packages.json`, `OSJS/dist/metadata.json` and `OSJS/dist/apps/MCSModbusToolkit/`.

Expected: both commands exit 0; Toolkit assets exist; Toolkit appears in the discovered package output and all named discovery destinations. An exit-zero-only reading is not a PASS.
Evidence: HEAD SHA with repair ancestor, pre/post `git status --short`, ICC-prerequisite status, toolchain and setup, separate raw command output and exit codes, discovered entries and artifact existence. A genuine product-test contradiction is FAIL; an unsafe, unavailable or stale-prerequisite test is BLOCKED. Source inspection alone is not PASS.

## Recorded failure / repair interruption
First JR run reported FAIL at `752a541`: both commands exited 0 and assets existed but Toolkit was absent from discovery manifests and `dist/apps/`. Human authorized bounded CODE task UMIG-002-R; it added only the missing OS.js local-package manifest at `cd67e15` and was archived at `workflow/archive/umig-002-r-toolkit-discovery-manifest.md`. Original raw FAIL evidence remains accessible via the immutable `handoff.md` version at commit `752a541`. This is a retry of the originally authorized TEST, not a change to the UMIG-002 → UMIG-002-T → UMIG-002-V feature sequence.

## Non-scope
No rendered UI launch, backend/runtime tests, source fixes, workflow advancement or ICC writes.

## Dependencies
UMIG-002 and UMIG-002-R CODE checkpoints archived. A current JR test packet is present in `handoff.md`; BLACK SHEEP WALL must refresh affected stale ICC context before JR executes. UMIG-002-V stays QUEUED until an actual TEST PASS is reviewed and the coding agent advances it.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
