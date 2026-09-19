# Handoff

## Direction and ownership — 2026-09-19

ChatGPT owns CODE, exact task-specific JR TEST packets, evidence review and workflow advancement; OpenHands/JR executes ONLY the current TEST/VERIFY packet via general `operation cwal.md`, writes ONLY the authorized JR report section of `handoff.md`, commits/pushes ONLY `handoff.md`, then STOPS. General `operation cwal.md` stays unchanged; only BLACK SHEEP WALL edits ICC. Preserve operator production Docker/`osjs-data`, user configurations, legacy applications, Electron and retained verification volumes. No production daemon access or `docker compose down -v`.

## Task selection / immutable evidence

**SOLE ACTIVE: UMIG-006-T — independent Diagnostics UNIT/BUILD TEST; NOT RUN / NO PASS.** UMIG-006 CODE archived COMPLETE SOURCE-ONLY at `workflow/archive/umig-006-adapt-diagnostics-tab.md`, checkpoint `c996794ff14dd65479e84123724f6b9fd9235ee6`. GitHub compare `a327766..c996794` confirms exactly six Toolkit source/test paths and NO product runtime, Go, MMA2, production Compose, ICC or workflow changes in the coding diff. Source was read back; no tests/build/browser executed by ChatGPT. UMIG-006-V QUEUED pending independently reviewed UMIG-006-T PASS; UMIG-007+ QUEUED, UMIG-008/009 cutover/removal Planning only, DOCKER-001 paused. No current live runtime test authorization.

Prior UMIG-005-V real-browser live core behavior was accepted with ONE UNVERIFIED check: JR switched tabs/re-selected rather than closing the actual Toolkit window and relaunching from the OS.js Start menu. Immutable JR evidence `handoff.md` at `1a04664e5fdc21e3bd323a97a4f12f3f9d39abdd`, review/archive commit `a32776608317669d75bea2b87369f69dfb518b55`. The exact window close/reopen canonical reload and no-auto-apply check remains mandatory in FUTURE UMIG-006-V browser packet. Do not claim it passed in UMIG-005-V or perform it during this UNIT/BUILD task. Retained volumes `mcsverify-1789784184-232_verify-data` and `mcsverify-rep-005-20260919_verify-data` must not be reused or removed without separate authority.

## UMIG-006 source-only checkpoint / unsupported capabilities

Toolkit `index.js` now clears Diagnostics fixture DOM before mount and wires `diagnostics-editor.js` to the SAME pre-existing Memory and Replicator contract objects, reusing their independent transports. `diagnostics-observer.js` loads each actual canonical document and calls read-only status for its explicitly labelled FIRST configured device; no devices means no status call and UNKNOWN. Independent load/status errors map UNAVAILABLE without hiding the other device, and generation tokens discard stale replies. The formerly dormant `diagnostics-model.js` now maps real status and Simulator None/IDLE. Global service/supervisor health remains UNKNOWN; runtime mode, binary/data paths and Windows service controls UNAVAILABLE/disabled; the text pane is labelled read-only observations, NOT genuine runtime logs. Periodic five-second reads stop on window destroy. No apply/suggest, service control, Windows IPC, additional HTTP/socket endpoint, direct file write, COMMS probe or fixture fallback is part of Diagnostics. Memory/Replicator editor implementation and legacy packages unchanged.

Six source/test paths in checkpoint: modified `OSJS/src/packages/MCSModbusToolkit/index.js`, `diagnostics-model.js`; added `diagnostics-observer.js`, `diagnostics-editor.js`, `OSJS/tests/toolkit-diagnostics-observer.test.js`, `OSJS/tests/toolkit-diagnostics-editor.test.js`. Existing `OSJS/tests/toolkit-diagnostics-model.test.js` remains and is mandatory. CODE has only authored and reviewed source/tests; independent commands below are NOT yet executed.

## JR TEST TASK — CURRENT: UMIG-006-T

GOAL: Independently verify Diagnostics pure mapping, canonical read-only selection, partial failure, DOM disabled controls, stale-response/disposal, Memory and Replicator regressions, toolkit import isolation and build/discovery. UNIT/BUILD ONLY. Do NOT test a real browser, backend or Docker here.

SAFE SETUP / STOP GATES: In a fresh disposable tracked-clean checkout of `tamzrod/MCS.OSJS` at latest `origin/main`, record full HEAD and `git status --porcelain` (empty); `git merge-base --is-ancestor c996794ff14dd65479e84123724f6b9fd9235ee6 HEAD` must exit 0. If shallow, read-only `git fetch --unshallow origin` allowed. Confirm `workflow/active_work/umig-006-t-diagnostics-fixtures.md` SOLE ACTIVE; UMIG-006 CODE archived COMPLETE SOURCE-ONLY; UMIG-006-V QUEUED/Previous UMIG-006-T. Read `operation cwal.md`, THIS packet and archive. Use Node 16 / npm 8 because `OSJS/package.json` supports `>=10 <17`; if sandbox defaults differ, install Node16 sandbox-local outside repo. If needed, `cd OSJS && npm install --no-audit --no-fund` allowed in the disposable checkout ONLY, expected exit 0 and tracked-clean. Never reset/clean/change tracked sources or lockfiles, access operator Docker/data or run tests against actual Go sockets/devices. If mandatory preparation is unavailable after safe local attempt, BLOCKED and STOP.

EXACT COMMANDS — execute in this exact order from `OSJS/`, record EACH command, full exit and actual case output, STOP product testing on first genuine required failure (FAIL); no source repair or retries to force PASS:
1. `node tests/toolkit-diagnostics-model.test.js`
2. `node tests/toolkit-diagnostics-observer.test.js`
3. `node tests/toolkit-diagnostics-editor.test.js`
4. `node tests/toolkit-memory-contract.test.js`
5. `node tests/toolkit-memory-adapter.test.js`
6. `node tests/toolkit-memory-relay.test.js`
7. `node tests/toolkit-replicator-contract.test.js`
8. `node tests/toolkit-replicator-errors.test.js`
9. `node tests/toolkit-replicator-adapter.test.js`
10. `node tests/toolkit-replicator-transport.test.js`
11. `node tests/toolkit-replicator-relay.test.js`
12. `node tests/toolkit-fixtures.test.js`
13. `npm run build:local-packages`
14. `npm run package:discover`
15. `npm run build`

EXACT SOURCE/ARTIFACT INSPECTION: Read Toolkit `index.js`, `diagnostics-model.js`, `diagnostics-observer.js`, `diagnostics-editor.js`, `server.js`, source `metadata.json`, `toolkit-renderer.js` and built import graph. Confirm `index.js` replaces `panel-diagnostics` children BEFORE `toolkit.element` is appended; imported new diagnostics editor and existing Memory/Replicator contracts/transports; no Diagnostics-specific provider, socket, privileged control, route, direct config writes, fixture fallback or native Windows calls. `toolkit-renderer.js` may still construct an initial Diagnostics fixture off-DOM: it must be cleared BEFORE mounting; distinguish that non-mounted donor fallback source from live Diagnostics. Source/bundle may include Electron provenance COMMENTS, not executed legacy imports. No import from legacy `ModbusSimulator`/`ModbusReplicator` packages, Electron host, or global stylesheet. Inspect the exact two preexisting Unix-socket basenames and per-service operation allowlists from source `server.js` (unchanged). Require discovered `mcs-modbus-toolkit as MCSModbusToolkit`, metadata `server: "server.js"` in source and OSJS `dist/metadata.json` Toolkit entry if emitted, nonempty Toolkit `src/packages/MCSModbusToolkit/dist/main.js` and `main.css`, OS.js bundle emitted; report sizes.

EXPECTED: All 12 Node tests and three build/discovery commands exit 0. Diagnostics model validates UNKNOWN global services, UNAVAILABLE paths/actions, actual Simulator IDLE, per-device status and independent failures. Observer cases prove first canonical device explicitly labelled, no status for empty docs, read-only load/status ONLY, malformed/outage and wrong-name fail closed. Fake-DOM cases prove Start/Stop disabled without click handlers, refresh strictly read-only, real error detail from named device only, stale value cleared immediately, pending response cannot repaint after destroy and timer canceled. Preexisting Memory/Replicator and fixtures regressions pass. No tests contact actual service: fake transport/fake mkdtemp Unix server is permitted ONLY where existing focused unit suites define it. Tracked diff from root `git diff --name-status a32776608317669d75bea2b87369f69dfb518b55..HEAD` must contain ONLY six Toolkit source/test paths plus authorized `workflow/archive/umig-006-adapt-diagnostics-tab.md`, removal of active UMIG-006, update of UMIG-006-T, and `handoff.md`; no Go/MMA2/production Compose/ICC/legacy changes. Final `git status --porcelain` empty before report. Any unexpected tracked mutation => record, do not reset/clean, BLOCKED.

EVIDENCE: full HEAD/ancestor/sole ACTIVE, Node/npm/setup and tracked status, 15 actual commands/exit/output, assertion groups, metadata/discovery/artifact size and import graph, read-only/disabled-control audit, exact bounded diff and any contradictions/deviations. No claim of rendered live Diagnostics, production acceptance, global service health, Windows COMMS, real window reopen, or cutover based on this stage.

VERDICT: PASS only if all required tests/build/inspections match; executed product contradiction = FAIL; unavailable mandatory evidence after safe preparation = BLOCKED. For a genuine FAIL, STOP product testing and report facts. No unrelated debug/code fixes, live exploration, autonomous successor selection, ICC/workflow edits or unscheduled commands.

REPORT-WRITE AUTHORITY: Replace ONLY `## JR TEST REPORT — UMIG-006-T` below, stopping before `## Next action and recommendation`. Commit/push ONLY `handoff.md`, then STOP. Leave every other file/section untouched. ChatGPT alone accepts evidence, archives and advances.

## JR TEST REPORT — UMIG-006-T

PENDING — source-only Diagnostics CODE checkpoint reviewed; NO independent tests, build, browser or backend verification have been run for UMIG-006.

## Next action and recommendation

Next action: OpenHands run `OPERATION CWAL` on latest `main` and the exact UMIG-006-T packet, report and STOP. Recommendation: keep this independent stage UNIT/BUILD only; the already queued UMIG-006-V later handles actual GUI safety and the explicitly unverified Replicator full-window close/relaunch check.
