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

Pending independent JR execution. JR may replace this section only; no prior PASS is implied.
