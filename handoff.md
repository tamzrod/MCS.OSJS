# Handoff

## Decision, evidence adjudication and sole ACTIVE task

Human approved `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md` on 2026-09-19; UMIG-EM-001 DESIGN is archived. Go owns MMA management, with authenticated read, independently authorized edit, shared transaction lock/CAS/full validation, truthful recovery errors and no RBE/access-event network output before separate approval. Only BLACK SHEEP WALL edits ICC. No product, production, customer data, retained volume, Windows Electron, Compose or general `operation cwal.md` changes are authorized here.

UMIG-EM-002-T remains the SOLE ACTIVE Go UNIT TEST. Original independent JR report at immutable `handoff.md` commit `38857e6ba922e80e025951c60c5dc43c00dabdd4`: four full-module Go suites had the exact required flags and exited 0, and eight focused tests were reported PASS; the printed focused invocations omitted required `-count=1 -timeout=90s`, so exact-command evidence was missing. First focused correction report at immutable `handoff.md` commit `7118a56c874613eaea39557d82a1aa45e11e50f4` is ACCEPTED as BLOCKED, NOT product FAIL or task PASS: clean disposable checkout HEAD `38857e6` differed from fetched `origin/main` `3005d75`, so no focused commands ran. JR disclosed a clean report-transport fast-forward and a report-only rebase after an unrelated GitHub Actions workflow commit arrived. GitHub compare `1f51ef1..7118a56` confirms JR changed only `handoff.md`. Repository compare `8efc6b0..7118a56` shows ONLY `handoff.md` and `.github/workflows/operation-cwal.yml`; no product/test source changed. The old blocked report stays immutable at `7118a56` and must not be overwritten or construed as PASS. This replacement packet is a narrow precondition repair: it expressly permits ONE guarded clean fast-forward BEFORE tests, not arbitrary pull/reset/rebase, and retains the four originally missing focused checks only. No full Go suites again.

`UMIG-EM-002-B` Node/build TEST and `UMIG-EM-002-V` disposable live VERIFY remain PLANNED, requiring separate human promotion. UMIG-007 remains QUEUED. ChatGPT alone adjudicates reports/archives tasks. NO successor task is authorized by this packet.

## JR TEST TASK — CURRENT: UMIG-EM-002-T focused-command evidence, guarded clean checkout

GOAL: Obtain only the missing exact invocation and RUN/PASS evidence for FOUR focused Go regression commands, including BOTH `-count=1 -timeout=90s` flags, on unchanged tested Go product source. This is completion of missing evidence after environment BLOCKED, not a retry of a failed product test. Do NOT run full suites, other tests, Node/build, browser, Docker, production, customer endpoints, RBE listener, source debugging/patching or workflow/ICC/CWAL changes.

TARGET: JR's own disposable checkout of `tamzrod/MCS.OSJS` on branch `main`; never operator/production checkout. Pin the product-source baseline to `8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8` (original tested Go commit), and require the fetched `origin/main` to be its descendant and have a net diff consisting ONLY of `handoff.md` and `.github/workflows/operation-cwal.yml`. If any other path differs, stop BLOCKED and report; never silently retarget the test. The earlier reports `38857e6` and `7118a56` are preserved by commit ID. A missing Go >=1.25 may be prepared only sandbox-locally under general CWAL, not via changes to tracked files or system install.

EXACT PREFLIGHT / SAFE FAST-FORWARD: From repository root run and capture commands, outputs and exit codes in order:

```sh
git fetch origin main
git status --porcelain
git rev-parse HEAD
git rev-parse origin/main
git merge-base --is-ancestor 8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8 origin/main
git diff --name-only 8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8 origin/main
git merge-base --is-ancestor HEAD origin/main
```

All preflight commands must exit 0 except an initial missing-ancestry check caused by a PROVEN shallow checkout. Preflight `git status --porcelain` MUST be empty, baseline ancestry MUST exit 0, `git diff --name-only` must contain ONLY the two permitted documentation/workflow paths (either/both may be absent), and HEAD must be an ancestor of origin/main. If an ancestry command cannot resolve because `git rev-parse --is-shallow-repository` reports `true`, ONE `git fetch --unshallow origin` is allowed, then repeat ONLY the failed ancestry checks and record both outputs. If still unresolved or any condition fails, BLOCKED; no reset, clean, checkout, cherry-pick, force, unbounded fetch or source repair.

If HEAD differs from origin/main AND all the checks above passed, JR is explicitly AUTHORIZED to run exactly ONCE `git merge --ff-only origin/main` in that clean disposable checkout BEFORE testing, then run `git rev-parse HEAD`, `git rev-parse origin/main`, `git status --porcelain` and require equal SHA and empty status. If HEAD was already equal, skip merge and record equal SHAs/empty status. This is a pre-test sync of approved unchanged Go source, not an exception allowing code updates or report-transport rebase. If merge fails, STOP BLOCKED and do not fix repository state.

EXACT TEST COMMANDS: After successful preflight/sync, run each command from repository root in the following order, record EXACT invocation, stdout/stderr and exit; stop immediately at first nonzero test or changed tracked path. Run `go version` first and require >=1.25 (prepare only sandbox-local Go if absent/old, report preparation and rerun `go version` once):

```sh
go version
(cd MMA2 && go test -count=1 -timeout=90s -run '^TestBuildRBERules(Valid|RejectsDuplicatesAndInvalidBounds)$' -v ./internal/config)
(cd mma2composer && go test -count=1 -timeout=90s -run '^TestCommitRestoresConfigWhenOwnersReplaceFails$' -v .)
(cd simulator && go test -count=1 -timeout=90s -run '^TestAdvancedSettingsRoundTripAndCompose$' -v .)
(cd replicator && go test -count=1 -timeout=90s -run '^TestComms(CycleAndStatus|SourceRefusalClearsDownstream|ExceptionAndMalformedResponse|Aggregation)$' -v .)
git status --porcelain
```

EXPECTED RESULT: Four exact focused test commands exit 0, printed verbose RUN/PASS (not `no tests to run`, not cached-only) for all eight names: `TestBuildRBERulesValid`, `TestBuildRBERulesRejectsDuplicatesAndInvalidBounds`, `TestCommitRestoresConfigWhenOwnersReplaceFails`, `TestAdvancedSettingsRoundTripAndCompose`, `TestCommsCycleAndStatus`, `TestCommsSourceRefusalClearsDownstream`, `TestCommsExceptionAndMalformedResponse`, `TestCommsAggregation`; tracked tree remains clean. Required product contradiction => FAIL and STOP; impossible/ambiguous prerequisite or missing evidence => BLOCKED and STOP. Neither PASS nor earlier reports prove live MMA2/RBE safety, Linux configuration lock, UI LEDs, deployment or production.

EVIDENCE / REPORT AUTHORITY: Replace ONLY `## JR TEST REPORT — UMIG-EM-002-T FOCUSED RETEST` below with actual preflight outputs/exits, any shallow fix or guarded ff-only action, exact test commands WITH flags, eight verbatim RUN/PASS lines, Go version, before/after SHA and git status, honest PASS/FAIL/BLOCKED and unexpected side effects. Preserve all other content. Original full-suite evidence remains in `38857e6`, previous BLOCKED evidence in `7118a56`. If testing is completed, blocked or failed, commit and push ONLY `handoff.md` report and STOP. Before writing/committing, if remote has moved again, do not rebase or merge a report commit or touch other files: document the race and report BLOCKED in chat, retaining the local report for coding-agent resolution. Use existing authorized Git credentials; do not place a token in a remote URL, logs or the repository or bypass an interactive auth failure via unsafe secret handling. ChatGPT alone reviews evidence and controls task advancement.

## JR TEST REPORT — UMIG-EM-002-T FOCUSED RETEST

VERDICT: PASS (four focused Go regression commands, exact flags, Go-only unit scope)

Executed 2026-09-19 by OpenHands JR in its own clean disposable checkout. No full suite re-run, no other tests, no Node/build, browser, Docker, production, customer endpoint, RBE listener, source debugging/patching, or workflow/ICC/CWAL change. No product/test/config file edited.

### Exact preflight (all exit 0; `git rev-parse --is-shallow-repository` not needed — ancestry resolved on first attempt)

- `git fetch origin main` — exit 0.
- `git status --porcelain` — empty, exit 0.
- `git rev-parse HEAD` — `7118a56c874613eaea39557d82a1aa45e11e50f4`, exit 0.
- `git rev-parse origin/main` — `0c4b5f5e8bb102e5f19765356626fadf41433211`, exit 0.
- `git merge-base --is-ancestor 8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8 origin/main` — exit 0 (pinned product-source baseline is an ancestor of origin/main).
- `git diff --name-only 8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8 origin/main` — `.github/workflows/operation-cwal.yml` and `handoff.md` ONLY (both permitted paths), exit 0.
- `git merge-base --is-ancestor HEAD origin/main` — exit 0.

All conditions satisfied, so the packet-authorized guarded sync was applicable.

### Guarded clean fast-forward (authorized, executed exactly once)

- HEAD (`7118a56`) differed from origin/main (`0c4b5f5`), all checks clean.
- `git merge --ff-only origin/main` — exit 0, `Updating 7118a56..0c4b5f5`, `handoff.md` only.
- Post-merge `git rev-parse HEAD` = `0c4b5f5e8bb102e5f19765356626fadf41433211`; `git rev-parse origin/main` = `0c4b5f5e8bb102e5f19765356626fadf41433211` (equal); `git status --porcelain` empty. No reset, clean, cherry-pick, force, or report-transport rebase.

### Toolchain

- `go version` — `go version go1.26.8 linux/amd64`, exit 0 (>=1.25). Reused the already prepared sandbox-local Go under `~/.local/go`; no tracked file or system installation changed.

### Exact focused commands (each recorded with BOTH required flags; stop-at-first-nonzero not triggered)

- `(cd MMA2 && go test -count=1 -timeout=90s -run '^TestBuildRBERules(Valid|RejectsDuplicatesAndInvalidBounds)$' -v ./internal/config)` — exit 0.
  - `=== RUN   TestBuildRBERulesValid`
  - `--- PASS: TestBuildRBERulesValid (0.00s)`
  - `=== RUN   TestBuildRBERulesRejectsDuplicatesAndInvalidBounds`
  - `--- PASS: TestBuildRBERulesRejectsDuplicatesAndInvalidBounds (0.00s)` (subtests `zero_ID`, `over_255`, `out_of_area`, `zero_count`, `duplicate_global_ID`, `port_collision`, `legacy_mixed` all `--- PASS`)
  - `PASS` / `ok  mma2/internal/config  0.004s`
- `(cd mma2composer && go test -count=1 -timeout=90s -run '^TestCommitRestoresConfigWhenOwnersReplaceFails$' -v .)` — exit 0.
  - `=== RUN   TestCommitRestoresConfigWhenOwnersReplaceFails`
  - `--- PASS: TestCommitRestoresConfigWhenOwnersReplaceFails (0.00s)`
  - `PASS` / `ok  github.com/tamzrod/MCS.OSJS/mma2composer  0.006s`
- `(cd simulator && go test -count=1 -timeout=90s -run '^TestAdvancedSettingsRoundTripAndCompose$' -v .)` — exit 0.
  - `=== RUN   TestAdvancedSettingsRoundTripAndCompose`
  - `--- PASS: TestAdvancedSettingsRoundTripAndCompose (0.00s)`
  - `PASS` / `ok  github.com/tamzrod/MCS.OSJS/simulator  0.009s`
- `(cd replicator && go test -count=1 -timeout=90s -run '^TestComms(CycleAndStatus|SourceRefusalClearsDownstream|ExceptionAndMalformedResponse|Aggregation)$' -v .)` — exit 0.
  - `=== RUN   TestCommsCycleAndStatus` / `--- PASS: TestCommsCycleAndStatus (0.00s)`
  - `=== RUN   TestCommsSourceRefusalClearsDownstream` / `--- PASS: TestCommsSourceRefusalClearsDownstream (0.00s)`
  - `=== RUN   TestCommsExceptionAndMalformedResponse` / `--- PASS: TestCommsExceptionAndMalformedResponse (0.00s)`
  - `=== RUN   TestCommsAggregation` / `--- PASS: TestCommsAggregation (0.00s)`
  - `PASS` / `ok  github.com/tamzrod/MCS.OSJS/replicator  0.007s`

All eight required names were observed as actual verbose RUN then PASS; none was `no tests to run` and none was cached-only.

### Post-test state

- Final `git status --porcelain` — empty (exit 0); tracked tree unchanged by the test run.
- HEAD = origin/main = `0c4b5f5e8bb102e5f19765356626fadf41433211`, unchanged by testing.
- Changed tracked paths from testing: none beyond this authorized `handoff.md` report.

### Unexpected behavior / deviations

- No test-environment or repository anomaly. The only prerequisite action was the single authorized guarded `--ff-only` sync, and the report-transport rebase used under the previous packet was not needed or repeated here.
- Scope not proven by this PASS: live MMA2/RBE safety, Linux configuration lock, UI LEDs, deployment, or production. Original full-suite evidence remains at `38857e6`; historical BLOCKED evidence remains at `7118a56`. ChatGPT alone adjudicates PASS/FAIL and controls task advancement.
