# Handoff

## Sole ACTIVE and evidence boundary

Human explicitly approved promotion of `UMIG-EM-002-B` on 2026-09-19. Exactly ONE ACTIVE task: `workflow/active_work/umig-em-002-b-toolkit-build.md` (TEST / OpenHands JR). The approved Electron-to-OS.js design remains in `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md`. Predecessor `UMIG-EM-002-T` is archived COMPLETE/PASS only for Go UNIT regression, backed by immutable JR reports `38857e6` and `dc0a11c` and review `425efaac4b06289723dd0d84bee7a0e873b34a88`. Go shared manager/lock/CAS remains unimplemented design; RBE and access-event output are gated off. `UMIG-EM-002-V` disposable live VERIFY remains PLANNED and not authorized; UMIG-007 visual parity deferred. Only BLACK SHEEP WALL edits ICC; ICC workflow context is stale, repository task and this handoff are execution authority.

The Node/OS.js source checkpoint is `8b5541fcefeba8df82bea675b15bd0d03c7c6c20`. A concurrent local commit between `425efaac` and this checkpoint touched ONLY `electron/SETTINGS_PERMISSIONS.md`, `electron/build/installer.nsh`, `electron/build/settings-access/main_windows.go`, `electron/build/settings-access/main_windows_test.go`, `electron/test/installer-permissions.test.js` and `ICC/context/electron-settings-owner.md`; it did NOT modify OS.js or Go. The human-promoted task is separate from those Windows changes. No product code or tests were run by ChatGPT. This packet replaces the prior NO ACTIVE stop gate and authorizes ONLY the Node/build TEST below.

## JR TEST TASK — CURRENT: UMIG-EM-002-B Toolkit Node/build regression

GOAL: independently execute the existing twelve Toolkit Node unit/fixture tests, build all local OS.js packages, discover the package manifest, and build the desktop shell on unchanged OS.js source. No Go test, simulator/replicator socket, GUI, Docker, live service, production, customer configuration, RBE/access-event listener, Windows installer test, source repair, or other task is authorized. Old Go PASS cannot substitute for these results.

TARGET / PRECONDITIONS: JR's OWN disposable `tamzrod/MCS.OSJS` checkout; NEVER an operator/production checkout or mounted data volume. From repository root, prove clean tracked status and that `origin/main` descends from source checkpoint `8b5541fcefeba8df82bea675b15bd0d03c7c6c20`. The NET changed paths from that checkpoint to fetched `origin/main` may ONLY be `handoff.md`, `planning/microtask/umig-em-002-b-toolkit-build.md`, `workflow/active_work/umig-em-002-b-toolkit-build.md`. Any different path means BLOCKED/STOP: do not retarget to a new source revision. Run, recording actual output and exit for each:

```sh
git fetch origin main
git status --porcelain
git rev-parse HEAD
git rev-parse origin/main
git merge-base --is-ancestor 8b5541fcefeba8df82bea675b15bd0d03c7c6c20 origin/main
git diff --name-only 8b5541fcefeba8df82bea675b15bd0d03c7c6c20 origin/main
git merge-base --is-ancestor HEAD origin/main
```

All these commands must exit 0, pre-test status must be empty, only three allowed paths may appear in diff, and HEAD must be ancestor of origin/main. If an ancestry check fails solely because `git rev-parse --is-shallow-repository` proves a shallow checkout, ONE `git fetch --unshallow origin` and recheck ONLY the failed ancestry commands are authorized; record both results. Otherwise BLOCKED; no reset/clean/rebase/cherry-pick, no product changes. If HEAD differs from origin/main after all checks PASS, run exactly ONE `git merge --ff-only origin/main`, then repeat `git rev-parse HEAD`, `git rev-parse origin/main`, `git status --porcelain`; require equal SHA and empty status before tests. If already equal, skip merge and record equal SHAs. If fast-forward fails, BLOCKED/STOP. This is a pre-test source-safe sync, not report-transport permission to rebase later.

TEST ENVIRONMENT: From `OSJS`, `node --version` must be 16.x (the established compatible Node within OSJS engine >=10 <17); `npm --version` must work. Sandbox-local/user-local Node 16/npm preparation is allowed under general CWAL if absent/incorrect, never system-global/product modification; record setup/version. Build requires local `OSJS/node_modules/.bin/webpack`. If missing, prepare dependencies in the disposable checkout with exactly ONE `(cd OSJS && npm install --no-save --package-lock=false --legacy-peer-deps)` after Node/npm readiness; use only normal dependency fetch and record exit/warnings. This may create ignored node_modules/dist, but MUST NOT edit tracked manifest, test, source, workflow, config or lock files; run `git status --porcelain` after setup and require empty. If installation cannot safely succeed, BLOCKED/STOP; do not patch dependencies, switch toolchain versions to force PASS, or invoke production. If deps already present, skip install and record it. The prior test gate's compatible Node 16.20.2/npm 8.19.4 is precedent, not a substitute for this environment evidence.

EXACT PRODUCT COMMANDS: From repository root, run these separately in this order. Record EACH exact invocation, stdout/stderr, exit code. On first executed contradiction/nonzero STOP FAIL; do not continue/retry or investigate/fix. These files actually live in `OSJS/tests/` (NOT the obsolete Planning placeholder). Each must genuinely execute its assertions and exit 0; do not accept skipped or missing scripts.

```sh
(cd OSJS && node tests/toolkit-fixtures.test.js)
(cd OSJS && node tests/toolkit-memory-contract.test.js)
(cd OSJS && node tests/toolkit-memory-adapter.test.js)
(cd OSJS && node tests/toolkit-memory-relay.test.js)
(cd OSJS && node tests/toolkit-replicator-contract.test.js)
(cd OSJS && node tests/toolkit-replicator-adapter.test.js)
(cd OSJS && node tests/toolkit-replicator-errors.test.js)
(cd OSJS && node tests/toolkit-replicator-relay.test.js)
(cd OSJS && node tests/toolkit-replicator-transport.test.js)
(cd OSJS && node tests/toolkit-diagnostics-model.test.js)
(cd OSJS && node tests/toolkit-diagnostics-observer.test.js)
(cd OSJS && node tests/toolkit-diagnostics-editor.test.js)
(cd OSJS && npm run build:local-packages)
(cd OSJS && npm run package:discover)
(cd OSJS && npm run build)
(cd OSJS && test -s src/packages/MCSModbusToolkit/dist/main.js)
(cd OSJS && test -s src/packages/MCSModbusToolkit/dist/main.css)
(cd OSJS && test -s packages.json)
(cd OSJS && grep -E 'mcs-modbus-toolkit|MCSModbusToolkit' packages.json)
(cd OSJS && test -s dist/index.html)
(cd OSJS && wc -c src/packages/MCSModbusToolkit/dist/main.js src/packages/MCSModbusToolkit/dist/main.css packages.json dist/index.html)
git status --porcelain
```

EXPECTED / ACCEPTANCE: All twelve named Node scripts, three repo-native build/discovery commands and six output checks (nonempty Toolkit JS, CSS, manifest, discoverable Toolkit name, desktop HTML and actual byte counts) return exit 0. Manifest/discovery output must identify the Toolkit. Tests must affirm their real fixture/contract assertions rather than a substitute syntax check. `npm` or Sass deprecation warnings alone are not failures if exact required commands exit 0; report warnings. After test, tracked tree remains clean. A failing executed assertion/build/artifact => FAIL; missing test evidence or unsafe/unavailable environment => BLOCKED. This is UNIT/BUILD only, not rendered UI, Go backend integration, deployed service, COMMS green LEDs, network exposure or production certification.

EVIDENCE / REPORT-WRITE AUTHORITY: Return HEAD/origin/main/preflight flags/diff/guarded ff results, Node/npm versions and dependency setup, verbatim per-command invocation/exit and relevant output for twelve scripts, three build phases and six artifact commands, warning text, pre/post tracked status, changed paths, unexpected side effects, and truthful PASS/FAIL/BLOCKED. Replace ONLY `## JR TEST REPORT — UMIG-EM-002-B` below; preserve all other sections and files. When report ready, re-fetch origin/main. If remote advanced, DO NOT rebase, merge, force-push or edit a different file for transport: report BLOCKED race in chat and retain local evidence for ChatGPT. If remote unchanged, commit/push ONLY `handoff.md`, verify HEAD=origin/main and tracked clean, and STOP. No autonomous workflow promotion, ICC/CWAL/product edits, or fabricated test completion.

## JR TEST REPORT — UMIG-EM-002-B

Pending independent JR execution. No Node/build PASS is implied.
