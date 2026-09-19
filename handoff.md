# Handoff

## Direction and authority — 2026-09-19

Human approved staged MCS.OSJS Toolkit migration and UMIG-005 Replicator CODE. ChatGPT owns product CODE/source checkpoints, JR evidence adjudication and task advancement. OpenHands/JR owns independent TEST/VERIFY via `operation cwal.md`, may edit ONLY the currently authorized JR REPORT section, push ONLY `handoff.md`, then STOP. Only BLACK SHEEP WALL edits ICC. Preserve the working operator Docker deployment, `osjs-data`, user/production settings, standalone legacy apps and Windows Electron. Never touch the production Docker daemon or use `docker compose down -v`.

## Task selection and immutable predecessors

**SOLE ACTIVE: UMIG-005-T — independent Toolkit Replicator UNIT/BUILD TEST; NOT RUN / NO PASS.** Source-only UMIG-005 archived at `workflow/archive/umig-005-connect-replicator-tab.md`; source checkpoint `11bea98391b4953ca2356385bf2ffd7164bb657c`, plus one authored typed-error test in this handoff commit. UMIG-005-V QUEUED and requires this TEST PASS review and a separately accepted disposable Replicator runtime target. UMIG-006+ QUEUED; UMIG-008/009 cutover/removal Planning only; DOCKER-001 queued/paused. No new deployment or live Replicator acceptance.

UMIG-004-V Memory live BEHAVIOR was accepted and archived with an explicit JR restart-command deviation. Evidence `handoff.md` at `d230f60f231086bdba95b981555ad02bd015e677`, reviewer disposition/archive at `8d5926b511f7dab99e7c8639110851f8dea36a54`. The exact prescribed `docker compose start modbus-simulator-runtime` FAILED because a one-shot seed refused the already populated test volume; JR used `docker start` on the same project-owned container, then directly observed browser recovery. The recovery behavior passed the stage gate; exact-packet compliance did NOT pass. Corrected future ownership-checked restart method is documented in `deploy/verify/README.md`; no repeat is authorized under the present unit/build task. Earlier UMIG-004-T independent unit/build PASS and UMIG-003-T/V fixture/build/rendered PASS archived. The disposable Memory test volume `mcsverify-1789784184-232_verify-data` was retained; do not reuse or remove it without distinct authority.

## UMIG-005 source-only checkpoint — no test claim

The sole CODE diff `8d5926b..11bea9` was independently compared through GitHub: exactly eight paths, all under `OSJS/src/packages/MCSModbusToolkit/` or `OSJS/tests/`; one additional authored `OSJS/tests/toolkit-replicator-errors.test.js` is in this workflow handoff. The Toolkit entry now uses prepared `replicator-contract.js`, its own `replicator-transport.js`, `replicator-adapter.js` and `replicator-editor.js`; fixture Replicator DOM is removed before mount, but Diagnostics remains fixture-only and Memory uses its unchanged separate transport/editor. Authenticated Toolkit `server.js` routes `mcs-replicator-*` requests to `$OSJS_DATA_DIR/run/modbus-replicator.sock` with v1 4-byte BE framing, size/timeout bounds, operation allowlist `load/apply/status/suggest` and correlated errors; other request IDs retain the Simulator socket and `load/apply/status`. Direct Go `status` result (not Memory's `{status}` wrapper), `load {document,suggestion}`, `apply {document,structural,message,completed_at}`, and `suggest {port,unit_id,owner,status}` are handled by the existing contract. Apply calls the Go backend ONCE, never writes a config directly. Manual ownership inspection is advisory; backend is final owner/validation authority. Per-block status is displayed; separate Network/TCP/Modbus/MMA2 COMMS LEDs remain UNKNOWN without probes. No legacy package imports. The source and future tests were authored/read back; ChatGPT has NOT RUN unit tests, build, browser, actual relay or live Replicator.

## JR TEST TASK — CURRENT: UMIG-005-T

GOAL: Independently verify Toolkit-owned Replicator v1 contract, validation/mapping, request transport and dual-socket OS.js relay, all Go-shaped ownership/error/status cases, Memory regression, and OS.js build/discovery. This is UNIT/BUILD only, NOT live browser/backend/Docker acceptance.

SAFE SETUP / EXACT PRECONDITIONS: Work in a fresh tracked-clean disposable checkout of `tamzrod/MCS.OSJS` at `origin/main`. Record full `git rev-parse HEAD`, `git status --porcelain` (must be empty) and `git merge-base --is-ancestor 11bea98391b4953ca2356385bf2ffd7164bb657c HEAD` (exit 0). Confirm `workflow/active_work/umig-005-t-replicator-adapter.md` is the SOLE ACTIVE task, UMIG-005 CODE archived COMPLETE, UMIG-005-V QUEUED (Previous UMIG-005-T). Read this packet, archived CODE task, and `operation cwal.md`. Ensure Node 16 (project supports `>=10 <17`), npm 8; if sandbox default differs, install compatible sandbox-local Node outside repo per `operation cwal.md` §3. In disposable checkout only run `cd OSJS && npm install --no-audit --no-fund` if dependencies missing; expect exit 0 and no tracked mutation. No production data/runtime/Docker/socket, no Go service or external endpoint, no changing source/lockfiles. If the checkout is shallow and ancestor unavailable, `git fetch --unshallow origin` is allowed read-only; never reset/clean tracked files. Record environment preparation.

EXACT TEST COMMANDS — from `OSJS/`, execute in this order, capture EACH exit and actual output, STOP on first required contradiction (FAIL), with no fixes/retries to force PASS:
1. `node tests/toolkit-replicator-contract.test.js`
2. `node tests/toolkit-replicator-errors.test.js`
3. `node tests/toolkit-replicator-adapter.test.js`
4. `node tests/toolkit-replicator-transport.test.js`
5. `node tests/toolkit-replicator-relay.test.js`
6. `node tests/toolkit-memory-contract.test.js`
7. `node tests/toolkit-memory-adapter.test.js`
8. `node tests/toolkit-memory-relay.test.js`
9. `node tests/toolkit-fixtures.test.js`
10. `npm run build:local-packages`
11. `npm run package:discover`
12. `npm run build`

EXACT SOURCE/ARTIFACT INSPECTION: Record `src/packages/MCSModbusToolkit/metadata.json` and built `src/packages/MCSModbusToolkit/dist/metadata.json` (if emitted); require the Toolkit's server metadata to be `server.js`, built `dist/main.js` and `dist/main.css` to exist and be nonempty, and discovery output to name `mcs-modbus-toolkit as MCSModbusToolkit`. Read `src/packages/MCSModbusToolkit/index.js`, `server.js`, `replicator-adapter.js`, `replicator-editor.js`, `replicator-transport.js` and the relevant emitted bundle/import graph: confirm no imports from legacy `ModbusReplicator`/`ModbusSimulator` UI/server, Electron or diagnostics-model; Memory's existing editor/transport remains imported, Diagnostics still fixture-only, Replicator contract+transport+editor imported. Verify source relay has authenticated WebSocket, Replicator request-ID namespace, per-service allowlists, exact two Unix socket basenames, MAX_MESSAGE=1MiB and four-byte BE framing; no new HTTP endpoint. Source and emitted bundle may contain provenance COMMENTS mentioning legacy/donor; distinguish comments/asset strings from executed imports. No live test by reading these files.

EXPECTED: All nine Node files and three build/discovery commands exit 0. Asserted cases must directly show authenticated/allowlisted separate sockets, complete correlated v1 envelopes, one apply request, typed APPLY_FAILED ownership collision and STATUS_FAILED missing device, suggest/inspect IN USE, source/Pull Block validation including FC/range/gaps/cadence, genuine per-block ERROR and UNKNOWN/UNAVAILABLE status, timeout and close handling, no fabricated COMMS health and Memory regression. The focused tests use only injected transport or a mkdtemp fake Unix server and remove it in `finally`; no actual Replicator/MMA2/service/config write. Build artifacts and discovery exist; no prohibited imports. From repository root compare `8d5926b511f7dab99e7c8639110851f8dea36a54..HEAD`, expect nine Toolkit source/test paths plus ONLY authorized UMIG-005 archive, UMIG-005-T state and handoff, no Go/MMA2/production Compose/legacy/ICC changes. Record final `git status --porcelain` (tracked-clean). Unexpected tracked change => record paths, do not restore/clean, report BLOCKED per CWAL.

EVIDENCE TO RETURN: HEAD/ancestor/sole ACTIVE, Node/npm versions and setup, all 12 commands with raw exit/output, relevant fixture case names/assertions, metadata/discovery/build artifact sizes, per-service socket and request/response/error mapping, source/bundle forbidden-import inspection, bounded file-diff list, final tracked status, unexpected behavior. Do NOT infer live GUI, functioning deployed Replicator, COMMS probe, Docker health or production safety from this UNIT/BUILD test. No source fixes by JR.

VERDICT: PASS only if EVERY required test/build/inspection result meets expectation; executed contradictory product behavior = FAIL; unavailable required evidence after reasonable safe sandbox prep = BLOCKED. State exactly which command failed. No source correction, live exploration or autonomous successor selection.

REPORT-WRITE AUTHORITY: JR may replace ONLY the following `## JR TEST REPORT — UMIG-005-T` section up to `## Next action and recommendation` with factual evidence. Commit and push ONLY `handoff.md`, then STOP. Never edit CODE, workflow, archive, ICC, other handoff sections or invoke a successor. ChatGPT alone adjudicates/archive/promotes.

## JR TEST REPORT — UMIG-005-T

PENDING — Replicator source-only checkpoint reviewed, focused test files authored but NOT RUN. No build, real browser, actual Go runtime or deployment verification executed for this checkpoint.

## Next action and recommendation

Next action (OpenHands/JR): execute `OPERATION CWAL` using this exact UMIG-005-T UNIT/BUILD packet at latest `main`, report and STOP. Recommendation: adjudicate actual test evidence before UMIG-005-V; its Replicator live runtime needs a separately bounded disposable target, not production or the former Memory-only test stack.
