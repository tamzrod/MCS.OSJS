# Handoff

## Current authority — 2026-09-20

Human approved `UMIG-EM-003` CODE split and ONE experimental OpenCode V2 JR TEST. ChatGPT completed SOURCE-ONLY checkpoints `UMIG-EM-003-L` (Linux shared lock), `UMIG-EM-003-S` (Simulator), `UMIG-EM-003` (Replicator/final audit); those three are archived. The final PRODUCT SOURCE checkpoint is exactly `949d7a8eb328dfa66f8359ffe81cfe359b8db2b7`; later commits until this TEST activation are workflow/handoff only. Source readback and changed-file inventories were reviewed, but no Go build, test, live verification or production action has been executed on this new code. The earlier `UMIG-EM-002-V` disposable synthetic live verification was reviewed/archived PASS in its separate historical scope and does not establish writer-lock safety.

SOLE ACTIVE: `workflow/active_work/umig-em-003-t-shared-lock.md` (TEST / independent OpenCode V2 JR, expressly ONE human-approved exception). `UMIG-EM-003-V` remains PLANNED, requires separate human approval and cannot auto-advance. OpenHands is fallback only and is NOT this task's runner. ICC remains stale; only BLACK SHEEP WALL edits ICC. Do not use operator/production Docker, services, devices, real Modbus endpoints, installed files or customer data. No network-output enablement, deployment, Electron changes or live VERIFY is authorized.

## JR TEST TASK — CURRENT: UMIG-EM-003-T shared writer lock unit/race TEST (Operation CWAL, one OpenCode trial)

GOAL: Independently execute ONLY the pinned Go module and race/fixture tests for the new shared Linux MMA2 writer-lock and the Simulator/Replicator integrations. Run Operation CWAL as independent JR, NOT as the coding agent; do not inspect or execute a previous closed JR packet. The human will prepare the separate worktree and adjust the local experimental agent system instruction BEFORE starting this packet. JR may read `AGENTS.md`, `operation cwal.md`, this `handoff.md`, the sole ACTIVE test file, and named source/test files strictly needed to interpret evidence. `handoff.md` is the sole test authority, not a request to invent work.

TARGET/SAFETY: ONLY the clean detached local JR Git worktree at `$HOME/apps/MCS.OSJS-jr` on the user's Legion, NOT `$HOME/apps/MCS.OSJS` main checkout. The worktree is NOT a security sandbox: for EVERY requested shell command the human must individually review and explicitly approve, never use `--auto`, 'allow always', unattended execution or an alternative agent with edits enabled. JR has NO edit, commit, push, git checkout/fetch/merge/reset/clean/restore, install/download, sudo, Docker/Compose, system service, host network probe, live endpoint or customer/production action authority. Do not change source, config, workflows, ICC, AGENTS, directives, go.mod/go.sum or tracked/untracked repository files; ordinary Go tool cache and t.TempDir synthetic test artifacts outside repo are allowed. The existing repository Go tests may open test-owned ephemeral **loopback** fixture listeners ONLY; do not enable `MCS_RUN_E2E`, use customer addresses, or start external MMA2. On any unsafe/ambiguous target, STOP BLOCKED; no fallback to main checkout/OpenHands/Docker or guessed remediation.

A. EXACT READ-ONLY PREFLIGHT (run from the JR worktree; report output and exit of each):
```sh
pwd -P
git status --porcelain --untracked-files=all
git rev-parse HEAD
git rev-parse origin/main
git worktree list --porcelain
git merge-base --is-ancestor 949d7a8eb328dfa66f8359ffe81cfe359b8db2b7 HEAD
git diff --name-only 949d7a8eb328dfa66f8359ffe81cfe359b8db2b7 HEAD
uname -s
go version
test "${MCS_RUN_E2E:-}" != "1"
df -Pk . "$HOME" /tmp
```
REQUIRE exact path `$HOME/apps/MCS.OSJS-jr`, `git status` EMPTY, HEAD = origin/main = the current test-activation commit (record full SHA), source checkpoint an ancestor, and the diff from `949d7a8e...` through HEAD consists of EXACTLY these four workflow-only paths, not product code: `handoff.md`, deleted `workflow/active_work/umig-em-003-shared-lock.md`, added `workflow/archive/umig-em-003-shared-lock.md`, changed `workflow/active_work/umig-em-003-t-shared-lock.md`. Require the single ACTIVE task to be `UMIG-EM-003-T`, CODE predecessors archived, `uname -s` = Linux, installed Go >=1.25, `MCS_RUN_E2E` not 1, and >=3 GiB disk available on filesystem(s) used for test cache/temp. Require existing offline dependency/cache and race compiler availability (do NOT download or install anything, edit Go dependencies or change the checkout). `git fetch` or `git switch` is NOT authorized for JR: the human updates the clean worktree before invocation. A failed gate, missing preloaded offline dependency or ambiguity => BLOCKED/STOP BEFORE test, record evidence; do not fix it.

B. EXACT TESTS (only if ALL A gates pass). From worktree root, execute in this order, ONCE EACH, stop on FIRST nonzero exit. These are the ONLY authorized product-test commands; environment flags prohibit dependency downloads and go.mod rewrites. The `-race` suites include the explicitly named new tests; do NOT rerun them separately or run live E2E:
```sh
(cd mma2composer && timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...)
(cd simulator && timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...)
(cd replicator && timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...)
```
REQUIRE all three commands exit 0; no FAIL/race/panic/hang/timeout; full package results OK. Require actual `-v` PASS evidence for composer `TestWriterLockTimeoutAndRelease`, `TestWriterLockCrashRelease`, `TestWriterLockRejectsAmbiguousPath`, `TestWriterLockConcurrentForeignOwnershipConflict`; Simulator `TestSimulatorComposeRejectsBusySharedWriterLock`; Replicator `TestReplicatorComposeRejectsBusySharedWriterLock`, `TestManagerCommittedUnacknowledgedRestartFailsClosed`, `TestRuntimeManagerApplyLifecycleAndStatus`; and existing foreign-reservation collision tests. A product test/compile failure => FAIL/STOP, not a fix task. Missing offline dependency/compiler => BLOCKED/STOP. If a command exits 124, report TIMEOUT rather than PASS. No retries or substituted tests.

C. POST-TEST READ-ONLY CHECK (even after FAIL if safe):
```sh
git status --porcelain --untracked-files=all
git rev-parse HEAD
```
REQUIRE clean repository and unchanged pinned HEAD. If tests modified files, record exact paths and STOP (no cleanup/reset). Nothing in a Go TEST proves live OS.js, external/installed network listeners, Electron, production or `UMIG-EM-003-V`.

EXPECTED/VERDICT: PASS ONLY for all A gates + all three exact Go race module suites with required actual named-test PASS evidence + clean C. Executed product compile/test contradiction => FAIL. Unsafe/ambiguous environment, missing offline dependencies/tooling, or required evidence unobservable => BLOCKED. Report any product FAIL separately even if another gate blocks further execution. Do not infer from static source inspection, old CI or prior OpenHands result.

EVIDENCE TO RETURN: Give `## JR TEST REPORT — UMIG-EM-003-T` IN THE OPENCODE CHAT, including verdict, timestamp, exact worktree path/HEAD/origin/source ancestor and four-path diff, sole ACTIVE check, Go version, E2E flag check, disk, EACH preflight command/exit, EACH executed exact test command/exit and verbatim stdout/stderr (or full captured transcript, with required named test lines and package statuses), any skipped/unrun command and reason, post-status, warnings and side effects. If output exceeds chat budget, preserve the original transcript within the current OpenCode session and identify where it can be copied from; never replace required evidence with a fabricated PASS. STOP immediately after report; do not promote or select another task.

REPORT-WRITE/TRANSPORT AUTHORITY: NONE. Experimental OpenCode has edit DENIED. It must NOT change `handoff.md` or any repository file, commit, push, create files or run shell redirection to bypass edit permission. Return the report in chat ONLY. ChatGPT coding agent will independently review the actual evidence and, only then, own any authorized report write/state advancement. OpenHands must not execute this one-task OpenCode trial.

## JR TEST REPORT — UMIG-EM-003-T

PENDING. No JR command/test/report has been observed for this activated packet. Source-only code is NOT a test PASS.
