# Handoff

## Current direction — 2026-09-19

Human reports existing Docker Compose deployment working and approved Electron renderer snapshot `1c971b9a6e00bafadf329df8821421a40cfc079c` for a one-time visual donor. MCS Modbus Toolkit remains in the existing `osjs-shell` image; MMA2, Simulator and Replicator backends, Electron Windows product, persistent data and legacy OS.js UIs remain untouched. Portable Electron and Modpoll are outside this authorized sequence.

## Authority and current work

ChatGPT owns CODE, source checkpoints, independent report review and workflow advancement. OpenHands/JR is TEST/VERIFY only via `operation cwal.md`: run only the exact current packet, change no product/workflow/ICC files, and write only the authorized report section. Only BLACK SHEEP WALL edits ICC. Git Active Work and this handoff, not the stale ICC snapshot, select work.

**Sole ACTIVE: UMIG-003-T — independent Toolkit renderer build/discovery and fixture TEST.** UMIG-001 approved donor is archived; UMIG-003 CODE is archived with source-only checkpoint `fede1715fadd5900da12fd9630793e3514117caf`. Its seven changed files were read back and bounded diff inspected; no build, runtime, GUI or backend test was performed by CODE. UMIG-003-V and later previously promoted Toolkit stages through UMIG-007A remain QUEUED, never automatically PASS. UMIG-008 launcher cutover and UMIG-009 legacy UI removal remain Planning and require separate human approval. DOCKER-001 stays QUEUED/paused: working deployment was reported by human without independent platform-specific evidence, not marked PASS. Preserve user data and Docker volumes.

## UMIG-003 source checkpoint and limits

Frozen donor: `electron/renderer/index.html`, `style.css`, `app.js`, `comms-status.js` at `1c971b9`; source-only provenance is `workflow/archive/umig-001-freeze-electron-ui-donor.md`. Toolkit-owned implementation is in `OSJS/src/packages/MCSModbusToolkit/{index.js,index.scss,renderer.css,css-text-loader.js,fixtures.js,toolkit-renderer.js}` with a focused independent fixture test at `OSJS/tests/toolkit-fixtures.test.js`. The HTML layout and original CSS are adapted inside an OS.js single-window ShadowRoot, so donor global selectors do not style desktop chrome. Only window-local tab click listeners exist and are removed on destruction; no timers, Electron imports, legacy package imports, runtime bridge, network, persistence or real service writes are authored. Example Memory FC1–FC4 and Replicator block/destination layouts and Diagnostics render with read-only fields; Save & Apply, CRUD, Windows Start/Stop and other actions are disabled. MMA2/Simulator/Replicator, COMMS and destination ownership render UNKNOWN; diagnostics paths are UNAVAILABLE. No donor executable, icon, installer asset or additional Toolkit container is copied.

This is **fixture-only rendering**, not a full live copy of Electron editing behavior. Real Toolkit-owned forms and backend adapters must be implemented in UMIG-004/005; Diagnostics status mapping in UMIG-006. An independent build/test is now required, followed by separate rendered UMIG-003-V verification. No new functionality is assumed working based on source review or the human's existing Docker deployment report. Historical UMIG-002-T build/discovery PASS at `e9d25e3`, UMIG-002-V one-window placeholder PASS at `7349caa`, earlier icon/sound 404 observations and DOCKER-001 source commit `1c971b9` remain historical only.

## JR TEST TASK — CURRENT: UMIG-003-T

GOAL: independently build and discover the committed OS.js Toolkit renderer, execute its fixture contract and inspect for prohibited Electron/legacy runtime dependency. This stage does NOT authorize rendered GUI, Docker deployment or live backend calls.

REPOSITORY STATE AND SAFETY: On a disposable clone/checkout of `tamzrod/MCS.OSJS` at current `origin/main`, record full HEAD and clean tracked `git status --porcelain` before setup. Check `git merge-base --is-ancestor fede1715fadd5900da12fd9630793e3514117caf HEAD` exits 0; confirm archived `workflow/archive/umig-003-copy-electron-renderer.md` is COMPLETE and `workflow/active_work/umig-003-t-renderer-build.md` is the sole ACTIVE task, with UMIG-003-V QUEUED. If predecessor or task selection conflicts, checkout is already dirty, or no safe disposable environment/compatible Node 16 is available, report BLOCKED rather than resetting, modifying project files, running BLACK SHEEP WALL or substituting another test. No user deployment, Docker volume, service, user configuration or production data may be touched. If dependencies are absent, `cd OSJS && npm install --no-audit --no-fund` is permitted only within the disposable checkout; record it and verify no tracked file changes. Keep all product and workflow paths unmodified.

EXACT COMMANDS, in order, capturing exit status and full relevant output of each, stopping on first failure:
1. `cd OSJS && npm run build:local-packages`
2. `cd OSJS && npm run package:discover`
3. `cd OSJS && npm run build`
4. `cd OSJS && node tests/toolkit-fixtures.test.js`
5. From `OSJS/`, inspect `src/packages/MCSModbusToolkit/dist/main.js` and `main.css` existence with `ls -l src/packages/MCSModbusToolkit/dist/main.js src/packages/MCSModbusToolkit/dist/main.css`; record Toolkit discovery entry in step 2 output or generated metadata. Run `grep -R -n -E 'mcsDesktop|ipcRenderer|contextBridge' src/packages/MCSModbusToolkit --include='*.js'`; **exit 1 with no matches is the expected result**, exit 0 with matches or exit 2/error is a failure. Read the Toolkit bundle/import references and confirm no Electron path, Electron host import, or dependency on legacy Simulator/Replicator UI package. Do not treat historical donor provenance in comments as a runtime import.
6. From repository root, report `git diff --name-only 5a8ec7d119a3b27ef97d81c20d26a9a28023becb..HEAD` and final `git status --porcelain`; no tracked changes from the tester. The diff may include authorized Toolkit source, fixture test and workflow/handoff/archive files, but no Electron, legacy UI, MMA2, Docker or Go product change from this CODE checkpoint.

EXPECTED PASS: each build and focused fixture test exit 0, `MCSModbusToolkit` discovered, its local JS/CSS artifacts exist, fixture UNKNOWN/UNAVAILABLE/clone checks pass, inspected Toolkit source and output do not require Electron or legacy UI packages, and checkout remains tracked-clean. A real product command failure, missing artifact or forbidden runtime dependency is FAIL with exact evidence; absence of indispensable safe environment is BLOCKED. Do not infer actual rendered tab behavior, visual parity or backend health from a build PASS.

EVIDENCE TO RETURN: HEAD and ancestor/task-state checks; Node/npm versions; safe dependency setup if any; exact command lines, stdout/stderr and exit codes; discovery entry, JS/CSS artifact paths/sizes, fixture case output, explicit grep exit and bundle/import inspection, pre/post tracked status, any unexpected observations. Keep output specific; do not run unrelated tests or fix code.

REPORT-WRITE AUTHORITY: JR replaces **only** `## JR TEST REPORT — UMIG-003-T` section below with a factual PASS/FAIL/BLOCKED report and evidence, then commits/pushes only `handoff.md` and stops. It may not edit CODE files, task statuses, ICC, other handoff sections, or invoke BLACK SHEEP WALL. ChatGPT independently reviews evidence and handles any later archival/advancement. A stale ICC baseline is NOT a test prerequisite.

## JR TEST REPORT — UMIG-003-T

PENDING — this TEST has not been run. JR replaces only this section after executing the exact authorized packet; no outcome is pre-classified.

## Next action and recommendation

Next action (human/OpenHands): invoke `OPERATION CWAL` against latest `main` to run the UMIG-003-T packet above and return independent evidence. Recommendation: do not rebuild or modify the live Docker deployment for this build/static gate; run the separate UMIG-003-V actual GUI stage only after reviewed TEST PASS.
