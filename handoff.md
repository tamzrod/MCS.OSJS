# Handoff

## Current authority — 2026-09-20

Human approved focused Replicator regression correction and a subsequent independently executed OpenCode JR TEST. Source-only CODE `UMIG-EM-003-R` is archived with exact new SOURCE checkpoint `538324a472eb15ff8ef97ee66f826fa02ba9462f`; prior failing JR report commit `2d61394adc671c7e6cfe728a012b047eec7a0c7d` is retained verbatim in `workflow/archive/umig-em-003-t-failed-report.md`. The code change touched ONLY `replicator/multiblock_regression_test.go`: pre-commit failure safely restores pollers; post-commit unacknowledged restart must NOT restore. Product runtime code unchanged. This new correction has NOT passed Go test/build or live VERIFY. The previous Composer/Simulator PASS results apply only to their earlier tested source; do not claim they were rerun here.

SOLE ACTIVE: `workflow/active_work/umig-em-003-r-t-recovery-test.md` (TEST / independent OpenCode V2 JR, one focused human-authorized run). CODE predecessor archived SOURCE ONLY; `UMIG-EM-003-V` remains PLANNED with separate human gate. OpenHands standby only; JR cannot edit task/ICC/source or select successor. Only BLACK SHEEP WALL edits ICC.

## JR TEST TASK — CURRENT: UMIG-EM-003-R-T focused Replicator recovery regression TEST

GOAL: Independently run one pinned offline Replicator Go race suite; capture actual evidence and push its report using ONLY the reviewed coding-agent-authored fixed runner. This packet supersedes closed `UMIG-EM-003-T` and does not authorize rerunning Composer/Simulator, installed MMA2, live E2E or production. OpenCode reads this `handoff.md`, `operation cwal.md`, sole ACTIVE task and exact named runner; `handoff.md` is complete authority.

TARGET/HUMAN PREPARATION: ONLY clean, detached `$HOME/apps/MCS.OSJS-jr` on Legion; not the main development checkout. HUMAN, outside OpenCode, checks clean checkout, runs `git fetch origin main` and `git switch --detach origin/main`, and verifies activation HEAD from GitHub. HUMAN reviews `workflow/cwal/umig-em-003-r-t.py` on that exact commit before approving its one named command. This worktree is NOT a sandbox. OpenCode agent remains shell ask, file editing denied, subagents denied, no `--auto`/allow-always. One explicit approval of the reviewed runner authorizes ONLY its enumerated read-only preflight/live GitHub ref query, ONE bounded offline Replicator test suite, read-only postcheck, replacing the named handoff JR report section, one-file Git commit, non-force push to main, remote confirmation. No arbitrary commands, independent file edits, installations/downloads, sudo, Docker/Compose, service restarts, operator devices/data, other network endpoints or other Git writes. Go tests may use only their own ephemeral loopback fixtures and external caches/temp dirs. The pre-existing local experimental profile's handoff-only exception continues to apply; do not weaken permissions.

**EXACT SINGLE JR COMMAND (ONCE ONLY, from the detached JR worktree, after human approval):**
```sh
python3 workflow/cwal/umig-em-003-r-t.py
```
Do not separately run steps A/B/C or manually rerun tests. If invocation interrupts, report INCOMPLETE with last confirmed step and do not rerun without newly authorized packet. Script exit 1 with confirmed `CWAL REPORT PUSHED` can mean real product FAIL; do not confuse with missing delivery.

A. PREFLIGHT: script checks exact worktree, initial clean, HEAD = origin/main = read-only LIVE `git ls-remote origin refs/heads/main`, product SOURCE ancestor `538324a472eb15ff8ef97ee66f826fa02ba9462f`, and exactly FIVE workflow-only changed paths from SOURCE to activated packet: `handoff.md`, deleted `workflow/active_work/umig-em-003-r-recovery-regression.md`, added `workflow/archive/umig-em-003-r-recovery-regression.md`, changed `workflow/active_work/umig-em-003-r-t-recovery-test.md`, added `workflow/cwal/umig-em-003-r-t.py`. Sole ACTIVE and CODE predecessor verified, Linux Go >=1.25, E2E !=1, >=3 GiB disk for workspace/HOME/tmp, offline deps/race compiler (never install/download). Stale/dirty/ambiguous => BLOCKED before product test; JR MUST NOT fetch/switch/merge/reset/clean/restore to repair. Do not treat local tracking ref alone as latest remote evidence.

B. EXACT TEST inside the runner, ONCE, stop on failure and do not rerun:
```sh
(cd replicator && timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...)
```
PASS requires exit 0, all package results OK, no panic/race/FAIL, actual `--- PASS` lines for `TestApplyPreCommitFailureRestoresPreviousPollers`, `TestApplyPostCommitRestartFailureStopsPreviousPollers`, `TestManagerCommittedUnacknowledgedRestartFailsClosed`, and `TestRuntimeManagerApplyLifecycleAndStatus`. Any other package/test FAIL means overall FAIL. Missing offline tooling/dependencies or required evidence => BLOCKED. Preserve full original stdout/stderr and exact exit.

C. POST-TEST: even on product FAIL if safe, clean `git status --porcelain --untracked-files=all`, same original test HEAD, and live remote still original test HEAD. Unexpected source mutation or remote race => STOP with underlying product verdict and transport BLOCKED, no cleanup.

REPORT / TRANSPORT AUTHORIZATION: The ONE human-approved fixed runner is the only file-write exception to OpenCode's edit-deny profile. After clean C it writes ONLY the final `## JR TEST REPORT — UMIG-EM-003-R-T` section of `handoff.md`; preserves all other handoff text; includes true PASS/FAIL/BLOCKED, timestamp, exact source & tested activation SHA, original full command/stdout/stderr/exit records, post-check, skipped context/scope. Valid FAIL also gets reported and pushed. Verify sole `handoff.md` changed/staged, diff check, commit only handoff, verify commit scope, recheck live remote, ONE non-force `git push origin HEAD:refs/heads/main`, verify remote equals new report commit and clean local tree. No workflow or code/ICC changes, no force/retry. Output `CWAL REPORT PUSHED: verdict=..., commit=...; JR STOP` ONLY after delivery confirmed. If blocked/interrupted/transport fails, disclose exact stage and product verdict; never invent PASS/push. Coding agent alone reviews report and decides TEST archive and any separately authorized next task. Go unit TEST is not live VERIFY.

## JR TEST REPORT — UMIG-EM-003-R-T

PENDING — source-only correction is awaiting independent OpenCode JR. No test or runtime PASS claimed.
