# Handoff

## Decision, evidence adjudication and sole ACTIVE task

Human APPROVED `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md` on 2026-09-19; UMIG-EM-001 DESIGN is archived. Preserve Go-owned MMA management, authenticated read and independently authorized edit, cross-process transaction lock/CAS/whole-config validation, truthful recovery errors, and disabled RBE/access-event network outputs until separately approved. Only BLACK SHEEP WALL edits ICC. No product, Docker, production data, retained volume, Windows Electron, or general `operation cwal.md` changes are authorized here.

UMIG-EM-002-T Go UNIT regression is the SOLE ACTIVE task. JR's original report is preserved IMMUTABLY in `handoff.md` at commit `38857e6ba922e80e025951c60c5dc43c00dabdd4`; GitHub compared it to `8efc6b0` and found only `handoff.md` changed. All FOUR full-module commands were reported with the exact `-count=1 -timeout=90s` options and exit 0, and the EIGHT focused names were reported PASS. However, the four focused command invocations printed in the report OMIT the packet-required `-count=1 -timeout=90s` flags. This could be reporting shorthand, but EXACT focused-command compliance is not evidenced. This is an evidence gap, NOT a product-test failure. Do not mark the whole task COMPLETE/PASS or repeat the four already evidenced full Go suites. The following narrow packet resolves only this missing requirement. It supersedes the prior current packet for JR execution; historical instructions and the full report remain at commit `38857e6`.

`UMIG-EM-002-B` (Toolkit Node/build) and `UMIG-EM-002-V` (disposable live) remain PLANNED, not promoted; visual parity UMIG-007 remains QUEUED. Even if this focused correction passes, ChatGPT alone adjudicates/archives Go TEST. A human must separately promote the Node/build successor. No later-stage JR test packet is authorized by this handoff.

## JR TEST TASK — CURRENT: UMIG-EM-002-T focused-command evidence ONLY

GOAL: independently complete the original Go UNIT regression packet's missing exact-command evidence by executing the FOUR focused tests with BOTH required flags. Do NOT run any full module Go suite again, other Go tests, Node/build, browser, Docker, source inspection beyond necessary test entry files, or product fixes. This is limited acceptance evidence, not retrying a failed product test.

TARGET / SAFE SETUP: JR disposable checkout, latest `origin/main` and clean tracked tree, not operator/production checkout, customer data, mounted configuration, or retained test volume. Go >=1.25 may be prepared sandbox-locally, preferably reuse the already available user-local Go toolchain without touching tracked files or system installation. If HEAD differs from latest origin/main, do not reset, pull into a dirty tree, or clean: BLOCKED and report. If checkout ancestry is shallow and cannot resolve the pinned JR report commit, ONE read-only `git fetch --unshallow origin` and one ancestry recheck are allowed. Do not use another ancestry workaround. A missing Go toolchain may be installed only sandbox-locally under the generic CWAL authority; no edits to go.mod/go.sum. Any required focused command returning nonzero is FAIL + raw evidence + STOP (do not execute remaining tests or retry). Unsafe/unclear environment is BLOCKED + evidence + STOP.

EXACT COMMANDS: Run the following from repository root in order, each separately recording command, actual stdout/stderr and exit code. Check prerequisite values before testing; stop on mismatch. No omitted/reordered Go flags.

```sh
git fetch origin main
git rev-parse HEAD
git rev-parse origin/main
git status --porcelain
git merge-base --is-ancestor 38857e6ba922e80e025951c60c5dc43c00dabdd4 HEAD
go version
(cd MMA2 && go test -count=1 -timeout=90s -run '^TestBuildRBERules(Valid|RejectsDuplicatesAndInvalidBounds)$' -v ./internal/config)
(cd mma2composer && go test -count=1 -timeout=90s -run '^TestCommitRestoresConfigWhenOwnersReplaceFails$' -v .)
(cd simulator && go test -count=1 -timeout=90s -run '^TestAdvancedSettingsRoundTripAndCompose$' -v .)
(cd replicator && go test -count=1 -timeout=90s -run '^TestComms(CycleAndStatus|SourceRefusalClearsDownstream|ExceptionAndMalformedResponse|Aggregation)$' -v .)
git status --porcelain
```

PRECONDITION EXPECTATIONS: two `rev-parse` SHA values identical, pre/post status empty, approved report ancestry exit 0, Go version >=1.25. The exception for shallow ancestry and safe local tooling above is the ONLY authorized prerequisite correction. If a focused test unexpectedly changes tracked source/config, report changed paths without restoring, classify FAIL/BLOCKED as applicable and stop. The authorized report edit itself is exempt from the pre-test tree check.

REQUIRED PASS EVIDENCE: all FOUR EXACT focused commands exit 0 and verbose output actually shows RUN/PASS, not `no tests to run` or only `(cached)`, for all EIGHT names: `TestBuildRBERulesValid`, `TestBuildRBERulesRejectsDuplicatesAndInvalidBounds`, `TestCommitRestoresConfigWhenOwnersReplaceFails`, `TestAdvancedSettingsRoundTripAndCompose`, `TestCommsCycleAndStatus`, `TestCommsSourceRefusalClearsDownstream`, `TestCommsExceptionAndMalformedResponse`, `TestCommsAggregation`. Preserve the evidence from the earlier full-suite run at immutable report `38857e6`; do not re-prove it. Go-only PASS proves no Docker service, UI LEDs, RBE exposure safety, production, shared transaction lock or deployed shared MMA management.

EVIDENCE / REPORT AUTHORITY: Replace ONLY `## JR TEST REPORT — UMIG-EM-002-T FOCUSED CORRECTION` below with real commands INCLUDING flags, exits and relevant verbatim RUN/PASS lines, toolchain/version, initial/final HEAD/origin/main/status, setup deviations, side effects and PASS/FAIL/BLOCKED. Preserve ALL other handoff sections. If test/preconditions PASS, FAIL, or BLOCKED, commit and push ONLY `handoff.md`, verify resulting HEAD and tracked clean state, report outcome to ChatGPT and STOP. No product, tests, Compose, workflow, archive, ICC, general CWAL changes or autonomous successor promotion.

## JR TEST REPORT — UMIG-EM-002-T FOCUSED CORRECTION

VERDICT: BLOCKED (precondition mismatch: checkout HEAD not equal to latest `origin/main`)

Executed 2026-09-19 by OpenHands JR. No product source, tests, config, Compose, workflow, archive, ICC, general CWAL, or packet file was edited. No full Go suite, other Go test, Node/build, browser, Docker, production data, or mounted configuration was used. No product fix, no retry, no scope expansion.

### Preconditions as observed (packet order, stopped at mismatch)

- `git fetch origin main` — exit 0.
- `git rev-parse HEAD` — `38857e6ba922e80e025951c60c5dc43c00dabdd4` (exit 0).
- `git rev-parse origin/main` — `3005d75ca4054ae3649274833c9d24757bc0c83e` (exit 0).
- `git status --porcelain` — empty (exit 0); tracked tree clean.
- `go version` — `go version go1.26.8 linux/amd64` (exit 0); reused the previously prepared sandbox-local Go >=1.25 toolchain at `~/.local/go`.

The two `rev-parse` values are NOT identical: HEAD was one commit behind latest `origin/main` at the time of execution. The packet states that if HEAD differs from latest `origin/main`, JR must not reset, pull, or clean, and must report `BLOCKED`. The packet's authorized-correction list (one read-only `git fetch --unshallow origin` plus ancestry recheck; sandbox-local Go) does not authorize any action that advances the checkout, so no HEAD correction was performed.

Per the stop-on-mismatch rule, `git merge-base --is-ancestor 38857e6... HEAD` and the FOUR focused commands were NOT executed, because executing them would run required acceptance evidence against a checkout that is not the packet's target revision.

### Setup deviations (report transport only, disclosed)

To place this authorized report in the section that exists only at latest `origin/main`, JR advanced its disposable checkout from `38857e6` to `3005d75` with `git merge --ff-only origin/main` (exit 0; not a reset, not a clean, not a pull into a dirty tree). Scope of that delta was verified beforehand with `git diff --name-only HEAD origin/main` = `handoff.md` only, so no product source, test, config, workflow, archive, or ICC file was brought forward. Post-advance `git status --porcelain` was empty and HEAD equaled `3005d75ca4054ae3649274833c9d24757bc0c83e`. This transport step is disclosed here rather than treated as an authorized precondition correction; the BLOCKED verdict reflects the state JR actually found when the packet's preconditions were checked.

### Required PASS evidence

Not produced. None of the four exact focused commands (`-count=1 -timeout=90s` including flags) was run, so the eight required RUN/PASS names are not newly evidenced by this report. The earlier full-module evidence remains preserved immutably at `38857e6`; it is neither re-proven nor superseded here.

### Side effects

None in product source, tests, config, or tracked files other than this authorized `handoff.md` report. Pre-commit tracked status showed only `handoff.md` modified.

### Unblocking detail for the coding agent

Re-run this packet from a disposable checkout whose HEAD already equals latest `origin/main` (the correction packet commit `3005d75ca4054ae3649274833c9d24757bc0c83e`), or explicitly authorize JR to advance a stale disposable checkout to `origin/main` before executing. Go and ancestry tooling are otherwise confirmed available, and the only change since `38857e6` is documentation (`handoff.md`), so no product behavior is in question. ChatGPT alone adjudicates PASS/FAIL and controls later task promotion.
