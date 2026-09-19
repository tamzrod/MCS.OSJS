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

**VERDICT: PASS** — all 9 Node test files and all 3 build/discovery commands exited 0; all required source/artifact inspections met expectation. UNIT/BUILD only. No live GUI, deployed Replicator, COMMS probe, Docker health or production safety is inferred.

### Preconditions / repository state
- Checkout: `origin/main` of `https://github.com/tamzrod/MCS.OSJS.git`, HEAD `f28d5f6b9d46d280ce6486f72de0afaec927ffe4` = `origin/main` = `origin/HEAD` (`UMIG-005: archive source-only Replicator integration and activate independent adapter UNIT/BUILD test`).
- `git status --porcelain` at start: empty (tracked-clean). Final `git status --porcelain` after all 12 commands: empty (tracked-clean). No tracked mutation from test/build; `.gitignore` covers `node_modules/`, `dist/`, `packages.json`, `*.log`.
- Initial clone was shallow (`git rev-parse --is-shallow-repository` → `true`, ancestor unavailable). Allowed read-only `git fetch --unshallow origin` performed; no reset/clean.
- `git merge-base --is-ancestor 11bea98391b4953ca2356385bf2ffd7164bb657c HEAD` → exit 0.
- Task state: `workflow/active_work/umig-005-t-replicator-adapter.md` = exactly one `ACTIVE` task (UMIG-005-T, TEST NOT RUN/NO PASS); `umig-005-connect-replicator-tab.md` absent from active_work and present as `workflow/archive/umig-005-connect-replicator-tab.md` (COMPLETE); `umig-005-v-replicator-runtime.md` = `QUEUED`, `Previous: UMIG-005-T`. Handoff and active_work agree.
- Environment: sandbox default was Node v22.23.2 / npm 10.9.8, outside project `engines` `>=10 <17`. Per CWAL §3, sandbox-local Node v16.20.2 / npm 8.19.4 installed at `~/.local/node-v16.20.2-linux-x64` (prepended to PATH for the test session); no repo files touched. `OSJS/node_modules` absent → `cd OSJS && npm install --no-audit --no-fund` in the disposable checkout: exit 0, "added 954 packages in 18s", no tracked mutation. No Go service, socket, Docker, production data or external endpoint contacted.

### Commands — raw exit results (run from `OSJS/`, in packet order, Node v16.20.2)
| # | Command | Exit | Output (fixture case lines) |
|---|---------|------|------------------------------|
| 1 | `node tests/toolkit-replicator-contract.test.js` | 0 | protocol and injected transport: checked / load, direct status, suggest/inspect, apply snapshot, envelopes: checked / invalid requests cannot reach transport: checked / runtime/transport failures and malformed replies fail closed: checked / `UMIG-CF-002 Replicator contract cases complete` |
| 2 | `node tests/toolkit-replicator-errors.test.js` | 0 | single apply, ownership collision, missing device and Go destination inspection: checked / `UMIG-005 Toolkit Replicator typed-error cases complete` |
| 3 | `node tests/toolkit-replicator-adapter.test.js` | 0 | canonical defaults, legacy pull block normalization and copy isolation / source, FC1–FC4 block/range/gap/cadence and case-insensitive duplicate validation / direct runtime status, per-block errors and fail-closed unavailable/unknown / complete |
| 4 | `node tests/toolkit-replicator-transport.test.js` | 0 | correlated replies, duplicate rejection, full envelope and namespace / timeout, disposal and send-after-close / complete |
| 5 | `node tests/toolkit-replicator-relay.test.js` | 0 | authentication, request ID and per-service allowlists / separate Toolkit Unix sockets, v1 envelopes, one apply and direct per-block status / correlated Replicator-only runtime unavailable without Memory cross-routing / complete |
| 6 | `node tests/toolkit-memory-contract.test.js` | 0 | version+injected transport; load/status/snapshot/unique IDs/envelope; invalid not reaching transport; runtime+transport failures; wrong correlation/missing data fail closed / `UMIG-CF-001 … complete` |
| 7 | `node tests/toolkit-memory-adapter.test.js` | 0 | provider envelope + stale reply ignored; explicit apply snapshot/canonical result/IDLE; typed failures/unavailable/teardown; None/Random round-trip; unknown/wrong-device/unavailable mapping / `UMIG-004 … complete` |
| 8 | `node tests/toolkit-memory-relay.test.js` | 0 | authenticated session + v1 ID + allowlist; preserved v1 framing/envelope on isolated Unix socket; correlated explicit unavailable / `UMIG-004 … complete` |
| 9 | `node tests/toolkit-fixtures.test.js` | 0 | fixture UNKNOWN constant; runtime/COMMS/diagnostics fail closed; Memory and Replicator example shapes; snapshots not mutable across windows / `UMIG-003 … complete` |
| 10 | `npm run build:local-packages` | 0 | built 5 local packages exactly once: MCSModbusToolkit, ModbusReplicator, ModbusSimulator, NamelessClassicIcons, NamelessWorkstationTheme; Toolkit child emitted `main.js` 46 KiB, `main.css` 121 bytes, webpack 4.47.0 |
| 11 | `npm run package:discover` | 0 | `✔ 7 package(s) discovered`; includes `- mcs-modbus-toolkit as MCSModbusToolkit [symlink, local]` |
| 12 | `npm run build` | 0 | webpack osjs bundle built (`osjs.js` 118 KiB, `vendors~osjs.js` 488 KiB, `index.html` etc.); no errors |

No required contradiction occurred, so execution did not stop early; no retries were used to force a pass.

### Required assertions observed
- Authenticated relay: relay test line 30–35 asserts unauthenticated `{_osjs_client:false}` → `INVALID_REQUEST`; non-allowlisted Replicator op `restart` and empty `request_id` → `INVALID_REQUEST`; Memory-namespaced ID `memory-test` requesting `suggest` → `INVALID_REQUEST`.
- Allowlisted separate sockets: with `OSJS_DATA_DIR` set to a `fs.mkdtempSync` root, fake servers bound to `run/modbus-replicator.sock` and `run/modbus-simulator.sock`; assertions confirm Replicator requests reached only the Replicator socket (`repReceived === requests`) and the Memory request reached only the Simulator socket (`memoryReceived === [memory]`). Both sockets removed server-side; whole temp root removed in `finally`; `OSJS_DATA_DIR` restored. No deployed Docker volume contacted.
- Correlated v1 envelopes: all four Replicator replies asserted `version===1 && ok===true` with `request_id` echoing `mcs-replicator-*`; error path returns `ok:false`, echoed `request_id: 'mcs-replicator-down'`, `error.code === 'RUNTIME_UNAVAILABLE'` after the Replicator server closed.
- One apply request: errors test asserts `calls.filter(c=>c.operation==='apply').length === 1` and call order `['apply','status','suggest']`.
- Typed errors: `APPLY_FAILED` ("owned by another producer") and `STATUS_FAILED` ("not found"), asserted by `code` and message regex; `suggest({inspect:true,…})` maps `owner: 'simulator'`, `status: 'IN USE'`.
- Source/Pull Block validation: adapter test covers endpoint/unit, FC1–FC4, uint16 start/count and range `start+count<=65536`, positive uint32 scan rate, gap rejection for same-FC blocks and case-insensitive unique names; also legacy `pull_block` → `pull_blocks` migration and copy isolation.
- Truthful per-block status: adapter `displayStatus` maps `status.source_status` + `blocks[]` with `last_error`; invalid/stale/error → `UNKNOWN`/`UNAVAILABLE`, never a fabricated healthy value.
- No fabricated COMMS health: editor sets the four Network/TCP/Modbus/MMA2 LEDs `data-state="UNKNOWN"`, `disabled`, `aria-label`/`title` stating no direct probe, with comment "COMMS layers not directly measured; no green inferred from source polling".
- Memory regression: memory contract/adapter/relay tests all exit 0 with the same case lines as previously accepted.

### Source / artifact inspection
- `src/packages/MCSModbusToolkit/metadata.json` (read): `type: application`, `name: MCSModbusToolkit`, `server: "server.js"`, `files: ["main.js","main.css"]`.
- Discovery metadata is emitted at `OSJS/dist/metadata.json` (a JSON array), not `src/packages/MCSModbusToolkit/dist/metadata.json` — packet wrote "(if emitted)"; the Toolkit entry there is identical to source metadata, `server: "server.js"`. `OSJS/packages.json` (array) lists all 7 packages including `src/packages/MCSModbusToolkit`.
- Built artifacts exist and are nonempty: `src/packages/MCSModbusToolkit/dist/main.js` 47134 bytes, `dist/main.css` 121 bytes (`main.js.map` 126495, `main.css.map` 532).
- Import graph: `index.js` imports `osjs`, `./metadata.json`, `./toolkit-renderer`, `./memory-contract`, `./memory-transport`, `./memory-editor`, `./replicator-contract`, `./replicator-transport`, `./replicator-editor`; `replicator-editor.js` requires `./replicator-adapter`; `server.js` requires only `net`/`path`; `toolkit-renderer.js` requires its own `renderer.css` and `./fixtures`. Memory's own editor/transport remain imported and shared with its unchanged contract. Diagnostics remains fixture-only (`diagnostics-model.js` exists but has no requirer in the Toolkit sources).
- Forbidden imports: `grep -rni "ModbusReplicator\|ModbusSimulator\|electron\|diagnostics-model"` over Toolkit `*.js` matches only two provenance comments in `toolkit-renderer.js` ("one-time visual adaptation of electron/renderer at 1c971b9", "No Electron host, OS.js backend, network, file, or persistence calls here"). Emitted `dist/main.js` contains zero occurrences of `ModbusReplicator`, `ModbusSimulator`, `electron` or `diagnostics-model`; the source map `sources` list is exactly the Toolkit's own modules + `external "OSjs"`, with no legacy package module. Comments/asset strings distinguished from executed imports.
- Relay surface: `server.js` exposes one authenticated OS.js `onmessage` provider (`ws._osjs_client`, `version===1`, non-empty string `request_id`, object payload); service selected by `request_id.startsWith('mcs-replicator-')`; `REPLICATOR_OPS = {load,apply,status,suggest}`, `MEMORY_OPS = {load,apply,status}`; socket basenames exactly `modbus-replicator.sock` / `modbus-simulator.sock` under `$OSJS_DATA_DIR/run`; `VERSION=1`, `MAX_MESSAGE=1024*1024`, 4-byte big-endian `writeUInt32BE`/`readUInt32BE` framing with size/count bounds and 25 s socket timeout; failures return correlated `RUNTIME_UNAVAILABLE`. No `http`/`express`/`createServer` route in the Toolkit server (only the string "No … HTTP endpoint" in a comment); no direct config write — apply is delegated to the Go `apply` operation.

### Diff bounds
`git diff --name-status 8d5926b511f7dab99e7c8639110851f8dea36a54..HEAD` → 13 paths, all authorized: 5 modified/added Toolkit sources (`index.js`, `server.js` modified; `replicator-adapter.js`, `replicator-editor.js`, `replicator-transport.js` added), 4 Toolkit Replicator tests added, plus `handoff.md`, the UMIG-005 archive add + active-work delete/move and `umig-005-t-replicator-adapter.md`. No Go, MMA2, production Compose, legacy-package or ICC changes.

### Unexpected behavior
- None blocking. Two benign notes: (1) discovery metadata lands at `OSJS/dist/metadata.json` rather than a package-local `dist/metadata.json`, consistent with the packet's "(if emitted)" wording; (2) sandbox default Node (v22) is outside the project engine range, so a sandbox-local Node 16 was used per CWAL §3 — the build/tests all ran cleanly on it.

### Scope honesty
This proves the CODE checkpoint's unit/build/relay/adapter assertions and static import/build surface only. It does NOT prove live Replicator GUI, a functioning deployed Replicator, real COMMS probes, Docker health, production safety or Go runtime behavior — those remain for UMIG-005-V with a separately accepted disposable target. No source fix, cleanup, archive, promotion or successor action was performed by JR.

## Next action and recommendation

Next action (OpenHands/JR): execute `OPERATION CWAL` using this exact UMIG-005-T UNIT/BUILD packet at latest `main`, report and STOP. Recommendation: adjudicate actual test evidence before UMIG-005-V; its Replicator live runtime needs a separately bounded disposable target, not production or the former Memory-only test stack.
