# Handoff

## Human decision and current task selection

2026-09-19: Human APPROVED the architecture/security/failure decisions in `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md`. UMIG-EM-001 DESIGN archived COMPLETE (human decision, no product/test PASS). Go Simulator Unix socket is the intended shared MMA manager; authenticated load + separately verified manage-MMA authorization for apply; mandatory shared multi-writer transaction lock/CAS/full validation and truthful post-commit recovery errors; RBE and access-event network outputs remain off pending separate safety review. Implement only through separately authorized CODE/T/V stages. Source reference `449cfda` with feature commits `ff6846f`, `1ca1740`; later relevant commits require diff/re-pin.

SOLE ACTIVE `workflow/active_work/umig-em-002-t-upgraded-baseline.md` — TEST / OpenHands JR. Next `UMIG-EM-002-B` Toolkit Node/build TEST remains PLANNED and requires separate human promotion; `UMIG-EM-002-V` live VERIFY likewise PLANNED. Deferred UMIG-007 visual parity and UMIG-007A deployment are not current; cutover/legacy retirement remain human-gated. Earlier UMIG-006-V report `f6b7549` and review `ec390c0` are historical tested-SHA evidence only, not acceptance of changed Go source. ICC workflow nodes are known stale; repository source/handoff is authoritative, only BLACK SHEEP WALL may edit ICC. No change to general `operation cwal.md` or product code by this workflow advancement.

## JR TEST TASK — CURRENT: UMIG-EM-002-T (Go regression only)

GOAL: Independently establish Go UNIT regression evidence for the already merged upgraded MMA2 RBE, mma2composer, Simulator advanced persistence and Replicator per-cycle COMMS. This packet is complete authority. JR must not debug/patch code or guess another test. It does not request OS.js Node/build, GUI, Docker, production, Windows, exposure of network output or workflow changes.

TARGET/PRECONDITIONS: Use JR's own disposable checkout, not operator/production checkout or data root. Fetch latest `origin/main` read-only and verify checked-out HEAD equals origin/main, HEAD is descendant of approved workflow checkpoint `68c41227d841609e6d73d77e1f87eb325499773c`, and tracked tree clean. No Docker daemon, config volume, application socket or customer device is needed. Go modules `MMA2`, `mma2composer`, `simulator` and `replicator` require Go >=1.25. Install a missing toolchain/dependencies only sandbox-locally per general CWAL, without modifying go.mod/go.sum or tracked files. If checkout unsafe/dirty or local tool prep impossible, BLOCKED + evidence + STOP. Do not reset/clean/restore. If a required product test exits nonzero, FAIL + evidence + STOP immediately; no other suites, retries, source investigation, workaround or fix.

EXACT COMMANDS — run from repository root in this order, separately capture stdout/stderr and exit code of EACH, stop at first nonzero. Shell blocks below are commands, not suggestions:

```sh
git fetch origin main
git rev-parse HEAD
git rev-parse origin/main
git status --porcelain
git merge-base --is-ancestor 68c41227d841609e6d73d77e1f87eb325499773c HEAD
go version
(cd MMA2 && go test -count=1 -timeout=90s ./...)
(cd mma2composer && go test -count=1 -timeout=90s ./...)
(cd simulator && go test -count=1 -timeout=90s ./...)
(cd replicator && go test -count=1 -timeout=90s ./...)
(cd MMA2 && go test -count=1 -timeout=90s -run '^TestBuildRBERules(Valid|RejectsDuplicatesAndInvalidBounds)$' -v ./internal/config)
(cd mma2composer && go test -count=1 -timeout=90s -run '^TestCommitRestoresConfigWhenOwnersReplaceFails$' -v .)
(cd simulator && go test -count=1 -timeout=90s -run '^TestAdvancedSettingsRoundTripAndCompose$' -v .)
(cd replicator && go test -count=1 -timeout=90s -run '^TestComms(CycleAndStatus|SourceRefusalClearsDownstream|ExceptionAndMalformedResponse|Aggregation)$' -v .)
git status --porcelain
```

PRECONDITION CHECKS: The two `rev-parse` outputs must be identical, `git status --porcelain` before and after test commands must be empty, `merge-base` must exit 0, and `go version` must establish Go >=1.25 (if not, first prepare safe sandbox-local Go and re-run only `go version`, reporting setup). On a shallow ancestry error, ONE read-only `git fetch --unshallow origin` and rerun `merge-base` is authorized; no other ancestry workaround. If tests unexpectedly change tracked source/config, report FAIL/BLOCKED with changed paths without cleaning them. Read `AGENTS.md`, `workflow/active_work/README.md`, the ACTIVE task and current packet solely to confirm authority; source inspection is limited to the four named test files to attribute a failing assertion, not to debugging.

EXPECTED RESULTS / GATES: All four full module Go suites and four focused commands exit 0. Focused verbose output must confirm actually RUN/PASS (not 'no tests to run') for `TestBuildRBERulesValid`, `TestBuildRBERulesRejectsDuplicatesAndInvalidBounds`, `TestCommitRestoresConfigWhenOwnersReplaceFails`, `TestAdvancedSettingsRoundTripAndCompose`, `TestCommsCycleAndStatus`, `TestCommsSourceRefusalClearsDownstream`, `TestCommsExceptionAndMalformedResponse`, `TestCommsAggregation`. Unit tests are allowed to open temporary loopback fixtures under their own temp roots only. Do not infer live RBE safety, installed backend readiness, production, UI COMMS LEDs, complete cross-process locking or MMA management support from a Go PASS. If any required item cannot be observed, BLOCKED; an executed contradiction is FAIL.

EVIDENCE TO RETURN: exact HEAD and origin/main SHAs, Go version/tool setup, pre/post tracked status, raw per-command invocation/exits and relevant output including eight focused test names, changed path list if any, unexpected findings and honest verdict. No substitutions or retry-driven debugging.

REPORT-WRITE AUTHORITY: Replace ONLY `## JR TEST REPORT — UMIG-EM-002-T` below with the report and evidence. Preserve ALL other text and files; do NOT edit source, tests, Compose, workflow, archive, ICC, general CWAL, or the packet. If authorized preconditions/test actions complete or BLOCK/FAIL, commit and push ONLY `handoff.md` with the report to `origin/main`, verify resulting HEAD and clean tracked tree, then STOP. ChatGPT alone adjudicates PASS/FAIL and controls later task promotion.

## JR TEST REPORT — UMIG-EM-002-T

VERDICT: PASS (Go UNIT regression only)

Executed 2026-09-19 by OpenHands JR in its own disposable checkout. One checkout only; no product source, tests, config, Compose, workflow, archive, ICC, general CWAL, or packet file edited. No Docker, production, data root, GUI, or network output used. No retry-driven debugging; no suite rerun to force a pass.

### Preconditions

- `git fetch origin main` exit 0.
- `git rev-parse HEAD` = `8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8` (exit 0).
- `git rev-parse origin/main` = `8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8` (exit 0). Identical to HEAD.
- Pre-test `git status --porcelain` = empty (exit 0), tracked tree clean.
- `git merge-base --is-ancestor 68c41227d841609e6d73d77e1f87eb325499773c HEAD` initially returned exit 128 `fatal: Not a valid commit name 68c41227...` because the sandbox checkout was a shallow clone. Per packet, exactly one `git fetch --unshallow origin` was run (exit 0); the rerun returned exit 0, confirming HEAD is a descendant of the approved workflow checkpoint. No other ancestry workaround used.

### Sandbox toolchain setup (not a product change)

- `go` was absent (`go: command not found`, exit 127).
- The four modules declare `go 1.25.0`; sandbox Go was installed user-locally under `~/.local/go`.
- Installed `go1.26.8.linux-amd64.tar.gz` from the official go.dev distribution; SHA-256 verified OK against the published checksum `d0f743b33e8d8945e6b1f432edd15785c70507121d6e2a723b21285eddf8b57b`.
- `go version` = `go version go1.26.8 linux/amd64` (exit 0), satisfying Go >= 1.25.
- `PATH` exported only for test shell sessions. No `go.mod`, `go.sum`, tracked file, project config, or system-wide location was modified. `gopkg.in/yaml.v3 v3.0.1` was downloaded into the module cache as a normal dependency fetch.

### Full module suites (all exit 0, stop-at-first-nonzero not triggered)

- `(cd MMA2 && go test -count=1 -timeout=90s ./...)` exit 0. Packages: `ok mma2/internal/config 0.014s`, `ok mma2/internal/ingress 0.032s`, `ok mma2/internal/memorycore 0.006s`, `ok mma2/internal/notify 0.042s`, `ok mma2/internal/rbe 0.090s`, `ok mma2/internal/restartwatch 0.003s`, `ok mma2/internal/transport/modbus 0.067s`, `ok mma2/internal/transport/rawingest 0.005s`, `ok mma2/pkg/configvalidate 0.004s`; remaining packages `[no test files]`.
- `(cd mma2composer && go test -count=1 -timeout=90s ./...)` exit 0. `ok github.com/tamzrod/MCS.OSJS/mma2composer 0.008s`.
- `(cd simulator && go test -count=1 -timeout=90s ./...)` exit 0. `ok github.com/tamzrod/MCS.OSJS/simulator 1.447s`, `ok github.com/tamzrod/MCS.OSJS/simulator/cmd/modbus-simulator-runtime 0.258s`.
- `(cd replicator && go test -count=1 -timeout=90s ./...)` exit 0. `ok github.com/tamzrod/MCS.OSJS/replicator 12.659s`, `ok github.com/tamzrod/MCS.OSJS/replicator/cmd/modbus-replicator-runtime 0.006s`.

### Focused verbose commands (all exit 0; each required test observed RUN then PASS, no "no tests to run")

- `(cd MMA2 && go test -run '^TestBuildRBERules(Valid|RejectsDuplicatesAndInvalidBounds)$' -v ./internal/config)` exit 0: `--- PASS: TestBuildRBERulesValid (0.00s)`; `--- PASS: TestBuildRBERulesRejectsDuplicatesAndInvalidBounds (0.00s)` with subtests `zero_ID`, `over_255`, `out_of_area`, `zero_count`, `duplicate_global_ID`, `port_collision`, `legacy_mixed` all PASS. `ok mma2/internal/config 0.004s`.
- `(cd mma2composer && go test -run '^TestCommitRestoresConfigWhenOwnersReplaceFails$' -v .)` exit 0: `--- PASS: TestCommitRestoresConfigWhenOwnersReplaceFails (0.00s)`. `ok github.com/tamzrod/MCS.OSJS/mma2composer 0.006s`.
- `(cd simulator && go test -run '^TestAdvancedSettingsRoundTripAndCompose$' -v .)` exit 0: `--- PASS: TestAdvancedSettingsRoundTripAndCompose (0.00s)`. `ok github.com/tamzrod/MCS.OSJS/simulator 0.008s`.
- `(cd replicator && go test -run '^TestComms(CycleAndStatus|SourceRefusalClearsDownstream|ExceptionAndMalformedResponse|Aggregation)$' -v .)` exit 0: `--- PASS: TestCommsCycleAndStatus`, `--- PASS: TestCommsSourceRefusalClearsDownstream`, `--- PASS: TestCommsExceptionAndMalformedResponse`, `--- PASS: TestCommsAggregation` (all 0.00s). `ok github.com/tamzrod/MCS.OSJS/replicator 0.007s`.

All eight required focused test names were observed as actually RUN/PASS.

### Post-test repository state

- Post-test `git status --porcelain` = empty (exit 0). Tracked tree unchanged by the test run.
- Changed tracked path list: none (other than this authorized `handoff.md` report). HEAD unchanged at `8efc6b00ab5c4f5acdffbf5a6e55c922246a99b8`.

### Unexpected behavior

- Only the initial shallow-clone ancestry failure, resolved by the single packet-authorized `git fetch --unshallow origin`. No product or test side effects observed.

### Scope limitations (unchanged conclusions intentionally not drawn)

A Go UNIT PASS does not establish live RBE safety, installed backend readiness, production behavior, UI COMMS LEDs, complete cross-process locking, or MMA management support. ChatGPT alone adjudicates PASS/FAIL and controls later task promotion.
