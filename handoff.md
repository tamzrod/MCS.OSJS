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

VERDICT: **PASS** — independent UNIT/BUILD stage executed 2026-09-19 by OpenHands/JR under `operation cwal.md`, in a fresh disposable checkout. This PASS covers the Toolkit Memory transport/contract/relay unit behavior and OS.js build/discovery ONLY. It does NOT establish rendered Memory UI, a working live backend, safe production writes, Docker health, or legacy cutover.

### Preconditions / task state
- Fresh disposable checkout `/tmp/jr-umig-004-t` (clone of `tamzrod/MCS.OSJS`), NOT the operator deployment.
- `git rev-parse HEAD` = `54896ff1fbdbfa4c052c993d5b9b3d6e83759ec8` (current `origin/main`).
- `git status --porcelain` before setup = empty (tracked-clean); final before report = empty.
- `node --version` = `v16.20.2`, `npm --version` = `8.19.4` (Node 16 `>=10 <17` satisfied; Node 16 sandbox-local at `/tmp/jr-node16`, outside repo).
- `git merge-base --is-ancestor 508c6b031e675c696de3ff7bdb4291dd005d54ae HEAD` → exit 0.
- Task state: `workflow/active_work/umig-004-t-memory-adapter.md` = SOLE ACTIVE (PENDING/NOT RUN); archived `umig-004-connect-memory-tab.md` = COMPLETE source-only; UMIG-004-V QUEUED. No conflict.
- Safe setup: `npm install --no-audit --no-fund` in the disposable checkout only (954 packages, exit 0, no tracked file changes). The relay test created its own `mkdtemp` fake Unix server under `/tmp`; no Docker/shared runtime socket, MMA2, Modbus operation or real configuration was referenced.

### Commands and observed results
1. `cd OSJS && node tests/toolkit-memory-contract.test.js` → exit **0**. Cases: `version and injected transport`, `load, status, snapshot apply, unique IDs and protocol envelope`, `invalid requests do not reach transport`, `runtime errors and transport failures propagated`, `wrong correlation, missing data and wrong protocol fail closed` — all `checked`. Also ran before dependency install (pure Node, no framework).
2. `cd OSJS && node tests/toolkit-memory-adapter.test.js` → exit **0**. Cases: `Toolkit provider envelope, correlated load, stale reply ignored`; `explicit apply snapshot, canonical result, status payload and IDLE`; `typed runtime failures, unavailable transport and teardown`; `None/Random round-trip, fresh defaults, empty document and validation`; `unknown, wrong-device, unavailable and observed status mapping` — all `checked`.
3. `cd OSJS && node tests/toolkit-memory-relay.test.js` → exit **0**. Cases: `provider requires authenticated OS.js session, v1 ID and allowlisted operation` (unauthenticated → `INVALID_REQUEST`, `restart` not allowlisted, empty `request_id` rejected); `Toolkit-owned provider preserves v1 framing and response envelope on isolated Unix socket` (4-byte BE framing verified round-trip); `unavailable runtime returns correlated explicit error` (`RUNTIME_UNAVAILABLE` with matching `request_id` = `no-runtime`). Fake server bound only to its own `mkdtemp` path and removed in `finally`.
4. `cd OSJS && npm run build:local-packages` → exit **0** — `built 5 local packages exactly once: MCSModbusToolkit, ModbusReplicator, ModbusSimulator, NamelessClassicIcons, NamelessWorkstationTheme`.
5. `cd OSJS && npm run package:discover` → exit **0** — `mcs-modbus-toolkit as MCSModbusToolkit [symlink, local]`, 7 discovered.
6. `cd OSJS && npm run build` → exit **0** — `osjs.js` 118 KiB, `vendors~osjs.js` 488 KiB, `osjs.css` 2.17 KiB, `index.html` emitted.

### Artifacts, metadata, inspection (command 7)
- Artifacts exist: `src/packages/MCSModbusToolkit/dist/main.js` = 30255 bytes, `dist/main.css` = 121 bytes (`ls` exit 0).
- `metadata.json` contains `"server": "server.js"` in BOTH source and generated `dist/metadata.json`; `MCSModbusToolkit` discovery entry correct (`type: application`, files `main.js`,`main.css`).
- Source runtime imports (`index.js`): `./index.scss`, `osjs`, `./metadata.json`, `./toolkit-renderer`, `./memory-contract`, `./memory-transport`, `./memory-editor`. `server.js` imports only `net`, `path`. No import of legacy `ModbusSimulator`/`ModbusReplicator`, Electron host, or dormant `replicator-contract.js`/`diagnostics-model.js`.
- Emitted `dist/main.js` references: dormant `replicator-contract`/`diagnostics-model` = 0 matches; Electron tokens (`mcsDesktop|ipcRenderer|contextBridge`) = 0; legacy `ModbusSimulator`/`ModbusReplicator` = 0 (each grep exit 1).
- Only Toolkit's Memory tab is live-bound: `index.js` creates transport/contract/editor for `#simulator-root` only; `toolkit-renderer.js` still renders Replicator and Diagnostics from the static fixture snapshot. `createReplicatorContract` exists in the dormant module but is not imported.
- `server.js`: accepts only authenticated (`ws._osjs_client`) v1 `load/apply/status` (allowlist `Set`), rejects bad version/ID/operation/payload with request-correlated `INVALID_REQUEST`; maps socket errors to `RUNTIME_UNAVAILABLE`; uses `path.join(OSJS_DATA_DIR || cwd, 'run', 'modbus-simulator.sock')`; max 1 MiB, 4-byte big-endian framing. No `route`/`express`/`healthz`/`listen(` — no new HTTP endpoint (grep exit 1).
- `memory-transport.js` retains the complete `{version,request_id,ok,result,error}` envelope (contract validates it), rejects duplicate/empty IDs, times out pending requests, and rejects on close.

### Baseline diff and final status (command 8)
- `git diff --name-only f596670527956302031e1ba4aa804386b0cfc133..HEAD` → exactly the seven expected UMIG-004 CODE toolkit paths (`index.js`, `memory-editor.js`, `memory-transport.js`, `metadata.json`, `server.js`, `tests/toolkit-memory-adapter.test.js`, `tests/toolkit-memory-relay.test.js`) plus authorized `handoff.md` and `workflow/active_work/|archive/` files. No Docker, Go/simulator, MMA2, Electron, or legacy OS.js product changes.
- Final `git status --porcelain` = empty (tracked-clean). No leftover `/tmp/toolkit-memory-relay-*` directories (test cleaned its own `mkdtemp`).

### Expected vs observed
All gating expectations met: six commands exit 0; focused tests demonstrate correlated load/apply/status and explicit errors, session-only None/Random round-trip, deleted-last empty-document validation, authenticated allowlisted fake Unix relay and correlated unavailable error; Toolkit discovered with JS/CSS artifacts; `server.js` registered; no prohibited runtime imports in source or bundle; transport uses the shared socket path while tests use only their disposable fake; tracked status clean.

### Unexpected behavior
None. Only non-fatal npm deprecation notices. No product-level failure observed in scope. JR did not connect to any real backend, did not run additional tests, and made no product/workflow/ICC edits; no successor selected.

## Next action and recommendation

Next action (OpenHands/JR): invoke `OPERATION CWAL` against latest `main`, execute current UMIG-004-T packet, report and stop. Recommendation: use only the disposable local fake runtime; preserve working Docker stack, production configuration and legacy apps. Review true TEST evidence before activating UMIG-004-V live/editor VERIFY.
