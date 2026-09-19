# Handoff

## Current human direction — 2026-09-19

Human approved staged Docker-hosted OS.js Toolkit Memory CODE after independently reviewed UMIG-003-T build PASS and UMIG-003-V real-browser fixture VERIFY PASS. ChatGPT owns CODE, source checkpoints, independent JR evidence review and task advancement; OpenHands/JR owns TEST/VERIFY only via `operation cwal.md`. Operator reports working Docker Compose: do not touch deployment, volumes or user configuration. Keep legacy apps, MMA2, Replicator, Diagnostics fixtures and separate Windows Electron. Only BLACK SHEEP WALL edits ICC; stale ICC is not a test prerequisite.

## Task selection and checkpoint

**SOLE ACTIVE: UMIG-004-T — independent Memory adapter/relay contract TEST. PENDING, NOT RUN, NOT PASS.** UMIG-004 CODE archived at `workflow/archive/umig-004-connect-memory-tab.md`, source checkpoint `508c6b031e675c696de3ff7bdb4291dd005d54ae`. UMIG-004-V and later sequence remain QUEUED with original gates. DOCKER-001 remains QUEUED/paused, not independently certified here. UMIG-008/009 launcher cutover and legacy UI removal remain Planning and require separate human approval. Do not deploy this code or perform real configuration writes in this TEST.

UMIG-003-T independent TEST PASS is preserved in `handoff.md` at immutable commit `b002c9f852456b3447ae0bef41ae4f69eb3073e1`; UMIG-003-V rendered VERIFY PASS at immutable `6f82196700b1c312652fdd7584a508e5452c4822`. Both archived after ChatGPT reviewed JR evidence. Source-only Memory code is NOT functionally verified: the previous fixture-only browser PASS says nothing about this replacement.

## UMIG-004 code scope and protocol

Source diff `f596670527956302031e1ba4aa804386b0cfc133..508c6b031e675c696de3ff7bdb4291dd005d54ae` reviewed: seven files only, under `OSJS/src/packages/MCSModbusToolkit/` and `OSJS/tests/`. New provider `server.js` and Toolkit `metadata.json` `server` entry use authenticated OS.js application WebSocket, strict `load/apply/status` allowlist and existing `$OSJS_DATA_DIR/run/modbus-simulator.sock`; 4-byte big-endian, max-1-MiB JSON framed v1 requests. Browser `index.js` uses new `memory-transport.js` retaining entire `{version,request_id,ok,result,error}` response, reused `memory-contract.js`, and Toolkit-owned `memory-editor.js`. Memory pane replaces display fixture BEFORE window mount. Load is automatic but never writes; Save & Apply is explicit, and failed apply retains unsaved draft. Canonical server document is the only Memory data source. Device status polls a selected persisted name and shows UNKNOWN/UNAVAILABLE for missing/mismatched/failing observations. None=0, Random=positive; empty-document Save & Apply remains available after deleting last device. Replicator and Diagnostics are still fixtures, header global runtime status is UNKNOWN, no legacy UI import. No Docker/Go/Electron/legacy/ICC changes, no deployment and no CODE-claimed test pass.

## JR TEST TASK — CURRENT: UMIG-004-T

GOAL / TARGET: Independently verify Toolkit-owned Memory v1 request transport, prepared Memory contract, read-only fake Simulator Unix relay, model/None-Random/error behavior and OS.js build/discovery from the committed source. This is a deterministic UNIT/BUILD stage, NOT a real backend apply, rendered UI test or Docker acceptance.

SAFE SETUP / PRECONDITIONS: Use a FRESH disposable Linux checkout of `tamzrod/MCS.OSJS` at current `origin/main`, NOT the operator's deployment. Record full `git rev-parse HEAD`, `git status --porcelain` (must be tracked-clean before setup), `node --version`, `npm --version`; compatible Node 16 (`>=10 <17`) required, sandbox/user-local installation permitted without repo changes. Verify exact ancestor `git merge-base --is-ancestor 508c6b031e675c696de3ff7bdb4291dd005d54ae HEAD` exits 0. Confirm `workflow/active_work/umig-004-t-memory-adapter.md` is SOLE ACTIVE, archived UMIG-004 says COMPLETE CODE/source-only, and UMIG-004-V is QUEUED. If checkout dirty, state mismatch, safe Linux/Unix-domain socket or Node 16 unavailable after reasonable safe sandbox setup, BLOCKED; never reset/clean/restore product, workflow or ICC. If dependencies absent, `cd OSJS && npm install --no-audit --no-fund` is allowed ONLY in disposable checkout, and only if it causes no tracked modifications; record setup and check tracked status. Relay test creates an independent fake local Unix server solely in an OS tmp `mkdtemp` directory, points `OSJS_DATA_DIR` there, and removes ONLY that directory; never reference real Docker/shared runtime socket, start MMA2, operate Modbus, or read/write real configuration. No operator Docker commands, production services or persistent data.

EXACT COMMANDS / ACTIONS, in order; record each command, meaningful stdout/stderr, exit code; STOP at first real product failure:
1. `cd OSJS && node tests/toolkit-memory-contract.test.js`
2. `cd OSJS && node tests/toolkit-memory-adapter.test.js`
3. `cd OSJS && node tests/toolkit-memory-relay.test.js`
4. `cd OSJS && npm run build:local-packages`
5. `cd OSJS && npm run package:discover`
6. `cd OSJS && npm run build`
7. From `OSJS/`, run `ls -l src/packages/MCSModbusToolkit/dist/main.js src/packages/MCSModbusToolkit/dist/main.css` and inspect `src/packages/MCSModbusToolkit/metadata.json` and generated `dist/metadata.json` for the exact `"server": "server.js"` provider and correct `MCSModbusToolkit` discovery. Inspect `index.js`, `memory-transport.js`, `memory-editor.js`, `server.js` and emitted Toolkit `dist/main.js` import references to confirm no runtime imports of legacy `ModbusSimulator`/`ModbusReplicator` UI, Electron host or the dormant `replicator-contract.js`/`diagnostics-model.js`. Confirm only Toolkit's Memory tab is bound; Replicator/Diagnostics remain fixture-only. Check server accepts only authenticated v1 `load/apply/status` and uses `$OSJS_DATA_DIR/run/modbus-simulator.sock` without new HTTP endpoint. These are inspections, not authorization to connect to the actual backend.
8. From repo root, record `git diff --name-only f596670527956302031e1ba4aa804386b0cfc133..HEAD` and final `git status --porcelain` BEFORE report edit. Expected migration delta: exactly seven Toolkit source/test paths from UMIG-004 CODE plus authorized `workflow/active_work/`, `workflow/archive/` and `handoff.md`; NO product changes under Docker, Go/simulator, MMA2, Electron, or legacy OS.js packages. Checkout must remain tracked-clean.

EXPECTED PASS: all six test/build/discovery commands exit 0; three focused test outputs demonstrate correlated load/apply/status and explicit errors, session-only None/Random interval round-trip, deleted-last empty-document validation, authenticated allowlisted fake Unix relay and unavailable error; Toolkit discovered and JS/CSS artifacts exist; Toolkit `server.js` registered and source/bundle contain no prohibited runtime imports; inspected transport uses shared socket path but tests use ONLY their disposable fake; tracked status clean. An executed test/build/artifact/dependency contradiction = FAIL. Missing indispensable safe execution environment, ambiguous configuration or unobservable mandatory evidence = BLOCKED. Passing this unit/build stage does not establish rendered Memory UI, working live backend, safe production writes, Docker health or legacy cutover.

EVIDENCE TO RETURN: full HEAD and ancestor/sole-ACTIVE checks, Node/npm and safe setup, exact commands/outputs/exits, all focused test-case summaries, fake Unix server path and cleanup/no live runtime, build/discovery/metadata and JS/CSS artifacts, source and bundle import inspection, baseline diff, pre/post tracked status, unexpected outputs and classification. Do not run additional tests or modify product.

REPORT-WRITE AUTHORITY: JR replaces ONLY `## JR TEST REPORT — UMIG-004-T` immediately below with factual PASS/FAIL/BLOCKED and direct evidence, then commits/pushes ONLY `handoff.md` and STOPS. If unexpected tracked changes occur, do not clean/restore: record exact paths and stop. JR cannot modify CODE, task files, other handoff sections, ICC, or advance stages. ChatGPT independently reviews evidence and handles later archive/UMIG-004-V promotion. Stale ICC baseline is NOT a blocker.

## JR TEST REPORT — UMIG-004-T

PENDING — exact independent TEST packet published. JR has not run the three focused tests or build/discovery for this stage; no PASS/FAIL is predetermined.

## Next action and recommendation

Next action (OpenHands/JR): invoke `OPERATION CWAL` against latest `main`, execute current UMIG-004-T packet, report and stop. Recommendation: use only the disposable local fake runtime; preserve working Docker stack, production configuration and legacy apps. Review true TEST evidence before activating UMIG-004-V live/editor VERIFY.
