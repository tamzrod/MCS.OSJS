# Handoff

## Human direction — 2026-09-19

Human approved prioritizing the pending independent fixture/build and rendered Toolkit verification before connecting prepared components to live backends. OpenHands/JR remains TEST/VERIFY-only; ChatGPT owns CODE, evidence review and workflow advancement. Existing Docker Compose is operator-reported working, not independently certified here, and must not be disturbed. Frozen one-time Electron UI donor: `1c971b9a6e00bafadf329df8821421a40cfc079c`. Portable Electron and Modpoll are separate.

## Task selection and authority

**Sole ACTIVE: UMIG-003-T — independent Toolkit renderer build/discovery and fixture TEST.** Human explicitly resumed this queued gate after isolated code preparation. Previous UMIG-003 CODE is COMPLETE in `workflow/archive/umig-003-copy-electron-renderer.md` (renderer source-only checkpoint `fede1715fadd5900da12fd9630793e3514117caf`); UMIG-003-V remains QUEUED, not passed. UMIG-004 through UMIG-007A remain QUEUED with original predecessor/PASS gates. DOCKER-001 is QUEUED/paused; UMIG-008/009 launcher cutover and legacy UI removal are Planning and require separate authorization. Do not mutate Docker volumes, user data, backends, Electron, or legacy OS.js apps.

OpenHands/JR follows `operation cwal.md` and only the exact CURRENT JR packet below; may edit only its authorized report section, push only `handoff.md`, and STOP. JR never fixes product, edits Active Work or ICC, or selects a successor. Only BLACK SHEEP WALL updates ICC. Active Work and this handoff, not the stale ICC cache, select the task.

## Source provenance and test boundary

The frozen donor was adapted into the Toolkit-owned three-tab fixture renderer at `fede1715fadd5900da12fd9630793e3514117caf`: `OSJS/src/packages/MCSModbusToolkit/{index.js,index.scss,renderer.css,css-text-loader.js,fixtures.js,toolkit-renderer.js}` and focused `OSJS/tests/toolkit-fixtures.test.js`. It mounts in one OS.js window with ShadowRoot-scoped CSS and local tab listeners; statuses are UNKNOWN/UNAVAILABLE, example fields are read-only, and action buttons are disabled. This source checkpoint was inspected but its build and rendered behavior were NOT verified.

Afterward three separate SOURCE-ONLY prep tasks added **dormant, unimported** `memory-contract.js` (`1159d690ddfe85d7ff17fe0e556eb94899cd84c6`), `replicator-contract.js` (`b00bb29dfae4597016820483d00209383a91cf8b`), and `diagnostics-model.js` (`1123ad0d14551d6b628104158b1419363cca43b2`) plus corresponding focused tests. Those tests were AUTHORED but NOT RUN; these modules are not imported into Toolkit UI or connected to OS.js transport, have no live status or writes, and are NOT product-test targets in UMIG-003-T. No real backend adapters, freshness mechanism, UI bindings, live actions or global service-health proof exists. The prior base-to-current diff `5a8ec7d119a3b27ef97d81c20d26a9a28023becb..4a036a9525f58f8a8404ddb10ebfae21dd678b38` was reviewed: only Toolkit source/tests and workflow/handoff/archive paths changed. Historical UMIG-002-T/V PASS established only the earlier single-window placeholder.

## JR TEST TASK — CURRENT: UMIG-003-T

GOAL: independently build/discover the committed OS.js Toolkit fixture renderer; run its existing focused fixture contract; inspect for prohibited Electron/legacy runtime dependencies. This is a build/static test, NOT a rendered GUI, Docker, live backend, or dormant Memory/Replicator/Diagnostics unit test.

SAFE SETUP / PRECONDITIONS: Use a disposable checkout of `tamzrod/MCS.OSJS` at current `origin/main`, not the operator's live deployment. Record full `git rev-parse HEAD`, `git status --porcelain` (must be tracked-clean before dependency setup), `node --version` and `npm --version`. Confirm Node 16 (`>=10 <17` for this OS.js project), and check the exact commands `git merge-base --is-ancestor fede1715fadd5900da12fd9630793e3514117caf HEAD` and `git merge-base --is-ancestor 4a036a9525f58f8a8404ddb10ebfae21dd678b38 HEAD` both exit 0. Read task files to confirm archived UMIG-003 is COMPLETE; `workflow/active_work/umig-003-t-renderer-build.md` is the sole ACTIVE task; UMIG-003-V is QUEUED; dormant UMIG-CF-001/002/003 are archived SOURCE-ONLY, not PASS. If required prerequisites are absent, task selection conflicts, or checkout was already dirty, report BLOCKED; do not reset/clean/restore/edit product, workflow or ICC. Safe local Node/tool installation is allowed outside the repo. If npm dependencies are missing, `cd OSJS && npm install --no-audit --no-fund` is permitted solely in the disposable checkout, record it and check no tracked files changed; otherwise BLOCKED if setup would require tracked modifications. Never touch Docker containers, volumes, real services, production config or user data.

EXACT COMMANDS / ACTIONS (in order, capture command, full relevant stdout/stderr and exit status; stop at the first real command failure):
1. `cd OSJS && npm run build:local-packages`
2. `cd OSJS && npm run package:discover`
3. `cd OSJS && npm run build`
4. `cd OSJS && node tests/toolkit-fixtures.test.js`
5. From `OSJS/`, run `ls -l src/packages/MCSModbusToolkit/dist/main.js src/packages/MCSModbusToolkit/dist/main.css` and record paths/sizes. Record `MCSModbusToolkit` discovery from step 2 or generated package metadata. Run `grep -R -n -E 'mcsDesktop|ipcRenderer|contextBridge' src/packages/MCSModbusToolkit --include='*.js'`: exit **1 with no matches is expected**; exit 0 with matches or exit 2/error fails. Inspect Toolkit entry-point imports (`index.js`, `toolkit-renderer.js`), committed dormant modules, and the emitted bundle/import references: the live renderer must not import any of the three dormant prep modules, Electron runtime/host, or legacy ModbusSimulator/ModbusReplicator UI packages. Do not misclassify comments referring to donor provenance as an executed import; an actual forbidden runtime import is FAIL.
6. From repo root run `git diff --name-only 5a8ec7d119a3b27ef97d81c20d26a9a28023becb..HEAD` and `git status --porcelain`. Expected compare paths: Toolkit-owned source and tests, plus authorized handoff/active/archive workflow files. No Electron, Go, MMA2, Docker or legacy UI product changes in this migration delta. Final tracked status must be clean **before** the authorized report update.

EXPECTED PASS: all four build/discovery/fixture commands exit 0; Toolkit discovered; JS/CSS artifacts exist; existing fixture tests pass; grep exits 1 without matches; imported runtime/bundle dependencies are free of Electron/legacy UI and the dormant three modules are not wired; checkout remains clean. A real executed product failure, forbidden import, missing artifact or tracked mutation is FAIL, with direct evidence; lack of safe required environment, conflicting task state or unobservable mandatory result is BLOCKED. A PASS does NOT prove tab rendering, visual parity, Docker health, backend connectivity, or the unrun new contract/model tests.

EVIDENCE TO RETURN: full HEAD, both ancestor checks, task state, Node/npm versions, safe setup actions, exact commands/outputs/exits, discovered Toolkit entry, artifact paths/sizes, fixture case output, grep exit, entry point/bundle inspection, scoped baseline diff, before/after tracked status, errors and unexpected observations. Do not add unrelated tests, fix code, restart services or infer evidence.

REPORT-WRITE AUTHORITY: JR replaces **only** `## JR TEST REPORT — UMIG-003-T` below with factual PASS/FAIL/BLOCKED, expected vs observed and direct evidence. It then commits/pushes ONLY `handoff.md` and stops. If the test leaves unexpected tracked changes, do not clean them; report the exact paths and stop, preserving evidence. JR may not edit other handoff sections, task statuses, product, workflow, ICC, or invoke BLACK SHEEP WALL. ChatGPT reviews report and alone decides archival or UMIG-003-V promotion; stale ICC is NOT a test prerequisite.

## JR TEST REPORT — UMIG-003-T

PENDING — current packet published for resumed independent test; JR has not executed this test. This section alone is authorized for observed PASS/FAIL/BLOCKED evidence.

## Next action and recommendation

Next action (OpenHands when available): run `OPERATION CWAL` from latest `main` against this exact UMIG-003-T packet. Recommendation: leave the working deployment untouched; after ChatGPT reviews genuine TEST PASS, separately activate UMIG-003-V for real rendered one-window/three-tab verification. Only then consider promoting gated live integration. Do not equate dormant prep or this unrun test with functional acceptance.
