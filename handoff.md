# Handoff

## Sole ACTIVE and evidence boundary — 2026-09-20

`UMIG-EM-002-R` is the only ACTIVE task: `workflow/active_work/umig-em-002-r-current-go-regression.md`, TEST / OpenHands JR. Its predecessor `UMIG-EM-002-B` is archived COMPLETE/PASS based on the independent 2026-09-19 JR report at commit `6612784f458339521e86ec822c4516c9af15df2b`; the complete old packet/report is preserved in that commit's `handoff.md`, not silently reclassified. OS.js source has not changed since that test, but newer Go Replicator advanced-settings source landed afterward. `UMIG-EM-002-V` is QUEUED, NOT runnable until this Go regression report is reviewed and a separate exact live-VERIFY packet replaces this one. `UMIG-EM-003` onward are PLANNED CODE; no UI cutover, shared lock, privileged MMA writes or network listener is authorized. `ICC/` workflow context is stale; only BLACK SHEEP WALL edits it. Git source, active task and this handoff are current execution authority.

Product checkpoint: `c98eacef8a6af8b0786a09766a4a034e25e82ce6`. The only expected net changes from that commit to this activation are `handoff.md`, the archived predecessor, deletion of its old ACTIVE file, and addition of the `UMIG-EM-002-R` ACTIVE and `UMIG-EM-002-V` QUEUED task files. No product source was changed by the activation.

## JR TEST TASK — CURRENT: UMIG-EM-002-R current Replicator Go regression

GOAL: independently test the exact current Go Replicator baseline, including persistence/composition/legacy inheritance/invalid-change and cloning tests introduced after prior Go test evidence. UNIT/REGRESSION only, not live verification.

TARGET: OpenHands JR's OWN disposable `tamzrod/MCS.OSJS` checkout only. Do not use an operator/production checkout, mounted service/customer volume or previously retained test environment. From repository root execute the following preflight commands separately in this order; record exact output and exit code:

```sh
git fetch origin main
git status --porcelain
git rev-parse HEAD
git rev-parse origin/main
git merge-base --is-ancestor c98eacef8a6af8b0786a09766a4a034e25e82ce6 origin/main
git diff --name-only c98eacef8a6af8b0786a09766a4a034e25e82ce6 origin/main
git merge-base --is-ancestor HEAD origin/main
```

Require clean status; every check exit 0; origin/main descends from pinned checkpoint, HEAD is ancestor. The diff must contain ONLY these exact paths: `handoff.md`, `workflow/archive/umig-em-002-b-toolkit-build.md`, `workflow/active_work/umig-em-002-b-toolkit-build.md`, `workflow/active_work/umig-em-002-r-current-go-regression.md`, `workflow/active_work/umig-em-002-v-upgraded-baseline.md` (deletion of the old ACTIVE path is expected). Any other path or failed non-shallow ancestry check means BLOCKED/STOP. If ancestry fails solely because `git rev-parse --is-shallow-repository` confirms a shallow checkout, one `git fetch --unshallow origin` and recheck only failed ancestry commands is permitted. If safe and HEAD differs from origin/main, run exactly ONE `git merge --ff-only origin/main`; then repeat `git rev-parse HEAD`, `git rev-parse origin/main`, `git status --porcelain`. Require equal SHAs and clean status before test. If already equal, skip merge. No reset, clean, rebase, cherry-pick or force. A changed source baseline is BLOCKED, not silently retargeted.

ENVIRONMENT: execute `go version` and `(cd replicator && go env GOMOD)`; require actual Go 1.25.x or later and GOMOD pointing to this checkout's `replicator/go.mod`. Sandbox/user-local installation of Go and ordinary module dependency downloads are permitted only without altering tracked files or production/system configuration. Record setup and exit. If safe setup unavailable, BLOCKED. Verify `(cd replicator && test -s advanced_settings_test.go)` exit 0. After setup repeat `git status --porcelain`; must be clean. No product/test/go.mod/go.sum edits to force passing.

EXACT PRODUCT TEST: from repository root run exactly once, capture full stdout/stderr, exit and elapsed time; on nonzero/contradiction STOP FAIL, no retry or repair:

```sh
(cd replicator && go test -race -count=1 -timeout=90s -v ./...)
git status --porcelain
```

EXPECTED: Go suite genuinely runs and exits 0. Verbose output must show PASS for `TestAdvancedSettingsPersistComposeAndInherit`, `TestInvalidAdvancedSettingsDoNotReplaceEffectiveConfig` and `TestAdvancedSettingsCloneIsIndependent`, plus no failed package/test. Those specific assertions cover preservation/inheritance, rejected invalid advanced config, and independent deep cloning. Post-test tracked status empty. Any test failure, timeout or nonzero => FAIL. Missing independent evidence or unsafe/unavailable environment => BLOCKED. No success claim from source inspection, prior Go PASS, Electron tests or an arbitrary syntax check. This does NOT prove real services, COMMS telemetry, OS.js advanced UI, actual listener binding, rendering or production readiness.

EVIDENCE / REPORT-WRITE AUTHORITY: Report all preflight SHAs/flags/net changed paths/fast-forward action, Go version/GOMOD and local setup, command stdout/stderr and exit, explicit three advanced test lines, total suite summary, warnings, post-test status, unexpected side effects and truthful PASS/FAIL/BLOCKED. Replace ONLY the `## JR TEST REPORT — UMIG-EM-002-R` section below; preserve the rest of this handoff and all other files. When report is ready, re-fetch origin/main; if it advanced, do NOT merge/rebase/force for report transport—report BLOCKED race in chat and retain evidence. If remote unchanged, commit/push ONLY `handoff.md`, verify HEAD=origin/main and clean tracked tree, STOP. Do not change ICC, product, workflow or another test packet. No autonomous task advancement.

## JR TEST REPORT — UMIG-EM-002-R

PENDING independent OpenHands JR execution. No current-Go regression PASS or live verification has been observed.
