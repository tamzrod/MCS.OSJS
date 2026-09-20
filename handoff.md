# Handoff

## Current authority — 2026-09-20

The human approved the split `UMIG-EM-003` CODE stages and **one OpenCode V2 JR test trial** on the Legion. CODE stages `UMIG-EM-003-L`, `UMIG-EM-003-S`, `UMIG-EM-003` are archived **SOURCE ONLY**, and the exact final PRODUCT SOURCE checkpoint is `949d7a8eb328dfa66f8359ffe81cfe359b8db2b7`. No coding-agent Go TEST or live VERIFY PASS is claimed. Prior `UMIG-EM-002-V` OpenHands synthetic live VERIFY PASS remains separate and does not prove this lock.

SOLE ACTIVE: `workflow/active_work/umig-em-003-t-shared-lock.md` (TEST / OpenCode V2 JR, specifically human-approved). `UMIG-EM-003-V` is PLANNED, requires independent human promotion. OpenHands is fallback, NOT runner for this task. ICC is stale and only BLACK SHEEP WALL edits ICC. No production, operator service, Docker, installed data, host Modbus, RBE access or Electron changes are authorized.

The previous OpenCode attempt began on `c2ba2b37336ff4d314452b608979744c65bad76c` and displayed three claimed suite successes but stopped without a complete independent report/post-check or report push. **Its verdict is INCOMPLETE, not accepted PASS.** The human expressly requested correcting ACTIVE task + `operation cwal.md` + this packet, and authorized ONE full fresh execution rather than piecemeal continuation. The current exact runner `workflow/cwal/umig-em-003-t.py` was authored by the coding agent as test infrastructure and is untested source until JR runs it. This activation includes ONLY six allowed workflow/directive/runner files changed from the product checkpoint, no product code.

## JR TEST TASK — CURRENT: UMIG-EM-003-T, one complete independent OpenCode CWAL run and report push

GOAL: Execute the complete locked-source Go TEST from A to report delivery without stopping after a partial preflight or three short PASS summaries. This packet supersedes the prior incomplete chat-only packet. Do not reuse the prior unadjudicated run as evidence. The current `operation cwal.md` is binding. JR may read `AGENTS.md`, `operation cwal.md`, this `handoff.md`, sole ACTIVE task, and the specifically named runner/test files needed for this task. Do not infer any other test authority.

TARGET / HUMAN SETUP: ONLY the detached `$HOME/apps/MCS.OSJS-jr` worktree on the Legion; **never** the main development checkout. Before launching a NEW OpenCode session, HUMAN (not JR) checks clean worktree, runs `git fetch origin main` and `git switch --detach origin/main`, and verifies the new activation HEAD. The runner verifies LIVE GitHub `main` with read-only `git ls-remote origin refs/heads/main` in addition to `HEAD` and local `origin/main` (remote query is the ONLY permitted non-test network action, apart from the authorized final `git push`). If freshness cannot be established, BLOCKED without tests or report push. JR may NOT fetch, switch, merge, clean, restore, reset or modify source to repair stale state. Human reviews `workflow/cwal/umig-em-003-t.py` on this commit before approving its EXACT command. Worktree is not a sandbox; retain OpenCode `shell: ask`, file-edit DENY, denied subagents, NO `--auto`, NO allow-always. One explicit approval of the named, reviewed, fixed script covers ONLY its enumerated preflight, offline Go suites, post-check, handoff-only report and non-force report push; no open-ended shell authorization. No sudo, install, download, Docker/Compose, persistent host service, external test endpoint, customer data, operator config or other repository writes. Go's existing own ephemeral loopback test fixtures and normal external Go caches/`t.TempDir` are allowed; `MCS_RUN_E2E=1` is forbidden.

**EXACT SINGLE JR COMMAND — from `$HOME/apps/MCS.OSJS-jr`, execute ONCE, after human shell approval:**
```sh
python3 workflow/cwal/umig-em-003-t.py
```
Do NOT separately re-execute A/B/C or launch background sessions. The fixed script prints intermediate exits, executes all gates and either prints a final confirmed report commit/verdict or an exact BLOCKED/incomplete/FAIL+transport error. If the invocation itself is interrupted or the model loses tools, do not claim completion or rerun tests: record the last confirmed stage and ask for coding-agent adjudication. If it returns nonzero with `CWAL REPORT PUSHED`, the report may validly be FAIL/BLOCKED; inspect the actual verdict rather than treating script exit 1 as a missing report.

A. EXACT PREFLIGHT implemented inside runner: `pwd -P`, `git status --porcelain --untracked-files=all`, `git rev-parse HEAD`, `git rev-parse origin/main`, `git worktree list --porcelain`, read-only `git ls-remote --exit-code origin refs/heads/main`, `git merge-base --is-ancestor 949d7a8eb328dfa66f8359ffe81cfe359b8db2b7 HEAD`, `git diff --name-only 949d7a8eb328dfa66f8359ffe81cfe359b8db2b7 HEAD`, `uname -s`, `go version`, `df -Pk . "$HOME" /tmp`, current `MCS_RUN_E2E` flag and disk availability. REQUIRE the exact separate path, initially clean detached checkout, HEAD = origin/main = live GitHub main returned by `ls-remote`, source checkpoint ancestor and **EXACTLY THESE SIX NON-PRODUCT changed paths** from source checkpoint to test activation:
```text
handoff.md
operation cwal.md
workflow/active_work/umig-em-003-shared-lock.md  (deleted)
workflow/active_work/umig-em-003-t-shared-lock.md
workflow/archive/umig-em-003-shared-lock.md  (added)
workflow/cwal/umig-em-003-t.py  (added)
```
Require sole ACTIVE `UMIG-EM-003-T`, CODE predecessor archive, Linux, Go >=1.25, E2E not 1, >=3 GiB free for workspace/HOME/tmp, offline dependencies and race compiler available without install. Any ambiguity => BLOCKED/STOP BEFORE tests; script must never fetch/switch/fix.

B. EXACT PRODUCT TESTS inside runner, in this order, ONCE EACH, stop at first nonzero or missing mandatory evidence; do NOT manually rerun. Offline Go module read-only flags and `timeout` bound each suite:
```sh
(cd mma2composer && timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...)
(cd simulator && timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...)
(cd replicator && timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...)
```
REQUIRE all three exits 0, real Go package OK, no race/panic/FAIL/timeout, and original `--- PASS` lines for `TestWriterLockTimeoutAndRelease`, `TestWriterLockCrashRelease`, `TestWriterLockRejectsAmbiguousPath`, `TestWriterLockConcurrentForeignOwnershipConflict`, `TestProducerIdentityAndForeignCollision`; Simulator `TestSimulatorComposeRejectsBusySharedWriterLock`, `TestComposeRejectsForeignCollisionUnchanged`; Replicator `TestReplicatorComposeRejectsBusySharedWriterLock`, `TestManagerCommittedUnacknowledgedRestartFailsClosed`, `TestRuntimeManagerApplyLifecycleAndStatus`, `TestRunOnceRejectsForeignOwnedDestination`. Source/old screenshots are NOT evidence. Executed product compile/test contradiction => FAIL; missing offline compiler/dependency or required observation => BLOCKED. Do not fix, substitute or retry.

C. POST-TEST, implemented even after a failed suite when safe: `git status --porcelain --untracked-files=all` must be empty and `git rev-parse HEAD` must still equal PRE-test HEAD. Before writing report, live remote must still equal PRE-test HEAD. Any unexpected mutation or remote race => STOP, disclose true product verdict and blocked report transport; no reset/cleanup.

EXPECTED REPORT / TRANSPORT (EXPLICIT EXCEPTION FOR THIS SINGLE TASK): When A passed and C is clean, the human-approved runner replaces ONLY the final `## JR TEST REPORT — UMIG-EM-003-T` section of `handoff.md`, preserving all other content, with timestamp, real verdict, source/activation SHA, preflight/remote, command+exit and FULL original stdout/stderr of each executed command, named tests, skipped tests and reason, postcheck and scope. This deliberately scoped script-mediated write is the SOLE exception to OpenCode's edit DENY; no arbitrary `edit` or shell-write permission is granted. Runner checks that `handoff.md` alone changed/staged, `git diff --check` passes, commits ONLY `handoff.md`, verifies report commit contains ONLY `handoff.md`, rechecks live remote is original PRE-test SHA, then `git push origin HEAD:refs/heads/main` (non-force) and independently checks remote equals new report SHA and worktree clean. HUMAN explicitly authorizes this fixed report-only commit/push with the single script approval. Push a real FAIL report too when the environment/postcheck/remote remains safe; NEVER label a failed test PASS. If report write/commit/push cannot complete, disclose the underlying actual verdict AND exact delivery failure without inventing a pushed commit or retrying. JR outputs `CWAL REPORT PUSHED: verdict=..., commit=...; JR STOP` ONLY after confirmed delivery. No other files, task state, ICC, source or agent config changes authorized. A Go unit TEST cannot prove live VERIFY, Electron or production readiness.

CODING AGENT NEXT: Fetch the actual pushed report commit, inspect full transcript and changed paths, adjudicate independently and only then own TEST archive/successor decision. No auto-advance by JR. OpenHands remains standby. Do not use original closed `UMIG-EM-002-V` packet.

## JR TEST REPORT — UMIG-EM-003-T

PENDING — initial OpenCode attempt was incomplete and had no accepted final evidence or pushed report. The corrected full run has not yet been observed. SOURCE ONLY is not TEST PASS.
