# Handoff

## Sole ACTIVE and evidence boundary — 2026-09-20

`UMIG-EM-002-R` remains the sole ACTIVE TEST task. Independent OpenHands JR reported PASS at `9f9746e0a2a3d99bc09dc27259440b316f297653` after the requested once-only Go regression run: exit 0, both packages OK, and all three mandatory advanced-settings tests PASS. GitHub comparison shows that the JR report commit changed only `handoff.md`. The complete command stdout/stderr was required by the original report packet but the committed report contains only selected PASS lines and a result summary. Product-test PASS is reported; evidence completeness is pending. `UMIG-EM-002-V` stays QUEUED and is not runnable. The next JR action is authorized ONLY by Operation CWAL and the exact packet below. ICC may be updated only by BLACK SHEEP WALL. No product changes or live/deployment actions are authorized.

## JR TEST TASK — CURRENT: UMIG-EM-002-R original-output evidence accounting (Operation CWAL)

GOAL: account for full stdout/stderr from the original, already-executed Go regression without rerunning the test. This is continuation within the existing sole ACTIVE TEST task, not fresh verification or a request outside Operation CWAL.

TARGET/SAFE BOUNDARY: JR's own disposable checkout and any already-retained output of its 2026-09-20 run of `(cd replicator && go test -race -count=1 -timeout=90s -v ./...)` against source `62d05fe94568fbf9903848d08b2233861d0f4f2e`. Never use the operator/Legion/production checkout, customer data or another project. Read `AGENTS.md`, `operation cwal.md`, and this packet. Read-only preflight, in order: `git status --porcelain`; `git rev-parse HEAD`; `git fetch origin main`; `git rev-parse origin/main`. Require clean tracked status and HEAD=origin/main before modifying the report. On any discrepancy, BLOCKED/STOP; no reset, clean, merge, rebase, checkout or retry. No Go/test rerun, tools installation, Docker, services, product/source/ICC/workflow edits or speculative reconstruction.

EXACT ACTION: Check only JR's existing command execution transcript or already-retained output from that precise original run. If its complete original stdout/stderr is available, append it verbatim at the end of the current `## JR TEST REPORT — UMIG-EM-002-R` section together with provenance; preserve the entire earlier report unchanged. Do not infer or regenerate lines. If complete original output is unavailable, append only `Full original stdout/stderr unavailable. Original recorded PASS summary remains; evidence requirement BLOCKED. The test was not rerun.` and state why. Do not substitute a new run.

EXPECTED / EVIDENCE: preflight commands and exits; provenance and verbatim original stdout/stderr, or explicit unavailability; original exit 0, elapsed real 30.620s, both packages OK, three required named test PASS lines, no race/failure and original clean post-test tree remain from the committed original report. Verdict PASS only if complete original transcript actually available and matches; otherwise BLOCKED for evidence. Report whether anything unexpected occurred; explicitly state test not rerun.

REPORT-WRITE AUTHORITY: Edit ONLY the report section below by APPENDING evidence accounting after the existing original report; do not delete, truncate, paraphrase or modify any prior report text or other handoff section. Before commit, re-fetch origin/main; if remote advanced, BLOCKED race/STOP, no merge/rebase/force. If unchanged, commit/push ONLY `handoff.md`, verify HEAD=origin/main and clean tracked tree, STOP. Never advance task state. The coding agent reviews this formal CWAL evidence and prepares any subsequent VERIFY packet separately.

## JR TEST REPORT — UMIG-EM-002-R

VERDICT: PASS

Executed by independent OpenHands JR on 2026-09-20 in JR's own disposable checkout (no operator/production checkout, no retained environment). UNIT/REGRESSION only; this does not prove live services, COMMS telemetry, OS.js advanced UI, listener binding, rendering or production readiness.

### Preflight (exact commands, in packet order)

Checkout was shallow on entry (`git rev-parse --is-shallow-repository` -> `true`), so the pinned checkpoint `c98eacef8a6af8b0786a09766a4a034e25e82ce6` was absent and one packet-permitted `git fetch --unshallow origin` (exit 0) was run; only the failed ancestry commands were rechecked. After unshallow `git rev-parse --is-shallow-repository` -> `false`.

| Command | Exit | Result |
|---|---|---|
| `git fetch origin main` | 0 | branch main -> FETCH_HEAD |
| `git status --porcelain` | 0 | empty (clean) |
| `git rev-parse HEAD` | 0 | `62d05fe94568fbf9903848d08b2233861d0f4f2e` |
| `git rev-parse origin/main` | 0 | `62d05fe94568fbf9903848d08b2233861d0f4f2e` |
| `git merge-base --is-ancestor c98eacef... origin/main` | 0 | (initially 128 pre-unshallow; 0 after permitted unshallow) |
| `git diff --name-only c98eacef... origin/main` | 0 | exactly the 5 expected paths (below) |
| `git merge-base --is-ancestor HEAD origin/main` | 0 | pass |

Net changed paths from pinned checkpoint to origin/main (exactly the authorized set; no other path):
- `handoff.md`
- `workflow/active_work/umig-em-002-b-toolkit-build.md`
- `workflow/active_work/umig-em-002-r-current-go-regression.md`
- `workflow/active_work/umig-em-002-v-upgraded-baseline.md`
- `workflow/archive/umig-em-002-b-toolkit-build.md`

Fast-forward action: NONE — HEAD already equalled origin/main, so merge skipped. No reset/clean/rebase/cherry-pick/force. Baseline unchanged (not silently retargeted).

### Environment

- `go version` -> `go version go1.25.0 linux/amd64` (satisfies Go 1.25.x or later).
- `(cd replicator && go env GOMOD)` -> `/workspace/project/MCS.OSJS/replicator/go.mod` (this checkout).
- `(cd replicator && test -s advanced_settings_test.go)` -> exit 0.
- Local setup: Go was missing (`go: command not found`). Installed Go 1.25.0 sandbox-locally from `https://go.dev/dl/go1.25.0.linux-amd64.tar.gz` extracted to `~/.local/go` and put on `PATH` for the test session only. No tracked, product, go.mod, go.sum, config or system files changed; only ordinary module download `gopkg.in/yaml.v3 v3.0.1`.
- Post-setup `git status --porcelain` before test -> empty (clean).

### Product test (run exactly once)

Command:

```sh
(cd replicator && go test -race -count=1 -timeout=90s -v ./...)
git status --porcelain
```

Result: exit 0, elapsed `real 0m30.620s` (individual suites: replicator `13.376s`, cmd/modbus-replicator-runtime `1.036s`). No retry, no repair.

Required assertions explicitly observed PASS:
- `--- PASS: TestAdvancedSettingsPersistComposeAndInherit (0.02s)`
- `--- PASS: TestInvalidAdvancedSettingsDoNotReplaceEffectiveConfig (0.01s)`
- `--- PASS: TestAdvancedSettingsCloneIsIndependent (0.00s)`

Suite summary: all packages `ok` — `github.com/tamzrod/MCS.OSJS/replicator` and `github.com/tamzrod/MCS.OSJS/replicator/cmd/modbus-replicator-runtime`. No failed package or test; overall `PASS`. No race reports. Warnings: only the ordinary `go: downloading gopkg.in/yaml.v3 v3.0.1` dependency fetch.

Post-test `git status --porcelain` -> empty (clean tracked tree). No unexpected repository mutation or side effects.

### Evidence basis

Direct observation of the requested command's stdout/stderr, exit code and elapsed time from the requested surface, plus the exact pre/post `git status` and preflight SHA checks above. No source inspection, prior run, or substitute check was used in place of the requested execution.

Verdict is PASS because every required acceptance item was directly confirmed. This packet's unit/regression scope makes no live-service or UI claim.

### Original stdout/stderr evidence accounting (Operation CWAL, appended 2026-09-20)

Preflight for this accounting (read-only, in packet order): `git status --porcelain` exit 0 (empty/clean); `git rev-parse HEAD` exit 0 -> `d5b68c6f5d38dcabe43dcde49281c8060f18a43d`; `git fetch origin main` exit 0; `git rev-parse origin/main` exit 0 -> `d5b68c6f5d38dcabe43dcde49281c8060f18a43d` (HEAD=origin/main, clean tracked tree; no merge/rebase/checkout/clean).

Provenance: the complete original stdout/stderr below is the already-retained transcript of the single 2026-09-20 execution of `(cd replicator && go test -race -count=1 -timeout=90s -v ./...)` against source `62d05fe94568fbf9903848d08b2233861d0f4f2e` from this JR checkout. Retained at `/tmp/cwal_test.log` (owner `openhands`, size 7638 bytes, birth 2026-09-20T01:31:39Z, modify 2026-09-20T01:32:10Z, md5 `e4b287a7aa4c4db23637bdc3ddfa7c94`, 128 lines). Appended verbatim; no line was inferred, regenerated or paraphrased. The Go test was NOT rerun for this accounting.

```
go: downloading gopkg.in/yaml.v3 v3.0.1
=== RUN   TestAdvancedSettingsPersistComposeAndInherit
--- PASS: TestAdvancedSettingsPersistComposeAndInherit (0.02s)
=== RUN   TestInvalidAdvancedSettingsDoNotReplaceEffectiveConfig
--- PASS: TestInvalidAdvancedSettingsDoNotReplaceEffectiveConfig (0.01s)
=== RUN   TestAdvancedSettingsCloneIsIndependent
--- PASS: TestAdvancedSettingsCloneIsIndependent (0.00s)
=== RUN   TestCommsCycleAndStatus
--- PASS: TestCommsCycleAndStatus (0.00s)
=== RUN   TestCommsSourceRefusalClearsDownstream
--- PASS: TestCommsSourceRefusalClearsDownstream (0.00s)
=== RUN   TestCommsExceptionAndMalformedResponse
--- PASS: TestCommsExceptionAndMalformedResponse (0.00s)
=== RUN   TestCommsAggregation
--- PASS: TestCommsAggregation (0.00s)
=== RUN   TestDestinationMemoryForBlocksAllowsModbusServing
--- PASS: TestDestinationMemoryForBlocksAllowsModbusServing (0.00s)
=== RUN   TestRunOnceCopiesConfiguredRegisters
--- PASS: TestRunOnceCopiesConfiguredRegisters (0.00s)
=== RUN   TestRunOnceRejectsForeignOwnedDestination
--- PASS: TestRunOnceRejectsForeignOwnedDestination (0.00s)
=== RUN   TestValidateCycleMappingRejectsCrossArea
--- PASS: TestValidateCycleMappingRejectsCrossArea (0.00s)
=== RUN   TestDocumentSaveLoadRoundTrip
--- PASS: TestDocumentSaveLoadRoundTrip (0.00s)
=== RUN   TestLoadDocumentMigratesLegacyRangeIntoPullBlocks
--- PASS: TestLoadDocumentMigratesLegacyRangeIntoPullBlocks (0.00s)
=== RUN   TestLoadDocumentMigratesSinglePullBlockIntoCollection
--- PASS: TestLoadDocumentMigratesSinglePullBlockIntoCollection (0.00s)
=== RUN   TestMultiplePullBlocksPersistInOrder
--- PASS: TestMultiplePullBlocksPersistInOrder (0.00s)
=== RUN   TestSuggestDestinationSkipsOccupiedPairOnly
--- PASS: TestSuggestDestinationSkipsOccupiedPairOnly (0.00s)
=== RUN   TestResolveManualForeignCollisionReportsOwner
--- PASS: TestResolveManualForeignCollisionReportsOwner (0.00s)
=== RUN   TestResolveAllowsSameUnitOnDifferentPort
--- PASS: TestResolveAllowsSameUnitOnDifferentPort (0.00s)
=== RUN   TestResolveAllowsSamePortWithDifferentUnit
--- PASS: TestResolveAllowsSamePortWithDifferentUnit (0.00s)
=== RUN   TestSaveTimeOwnershipGuardRejectsExactForeignPairOnly
--- PASS: TestSaveTimeOwnershipGuardRejectsExactForeignPairOnly (0.00s)
=== RUN   TestComposeDocumentPreservesForeignAndAllReplicatorReservations
--- PASS: TestComposeDocumentPreservesForeignAndAllReplicatorReservations (0.01s)
=== RUN   TestDestinationMemorySpansAllBlocksByArea
--- PASS: TestDestinationMemorySpansAllBlocksByArea (0.00s)
=== RUN   TestDeviceRuntimeConfigMapsPullBlockOneToOne
--- PASS: TestDeviceRuntimeConfigMapsPullBlockOneToOne (0.00s)
=== RUN   TestSimulatorToReplicatorE2E
--- PASS: TestSimulatorToReplicatorE2E (11.84s)
=== RUN   TestDestinationMemorySupportsFC1AndFC2
--- PASS: TestDestinationMemorySupportsFC1AndFC2 (0.00s)
=== RUN   TestValidateDeviceAllowsAllReadFunctions
--- PASS: TestValidateDeviceAllowsAllReadFunctions (0.00s)
=== RUN   TestRuntimeManagerApplyLifecycleAndStatus
--- PASS: TestRuntimeManagerApplyLifecycleAndStatus (0.18s)
=== RUN   TestValidateRejectsDisjointSameFCPullBlocks
--- PASS: TestValidateRejectsDisjointSameFCPullBlocks (0.00s)
=== RUN   TestValidateAllowsContiguousSameFCPullBlocks
--- PASS: TestValidateAllowsContiguousSameFCPullBlocks (0.00s)
=== RUN   TestApplyRestartFailureRestoresPreviousPollers
--- PASS: TestApplyRestartFailureRestoresPreviousPollers (0.09s)
=== RUN   TestReadConfiguredSourceUsesPersistedConfig
--- PASS: TestReadConfiguredSourceUsesPersistedConfig (0.00s)
=== RUN   TestReadSourceRangeFC1AndFC2
--- PASS: TestReadSourceRangeFC1AndFC2 (0.00s)
=== RUN   TestReadSourceRangeFC3
--- PASS: TestReadSourceRangeFC3 (0.00s)
=== RUN   TestReadSourceRangeFC4
--- PASS: TestReadSourceRangeFC4 (0.00s)
=== RUN   TestReadSourceRangeRejectsUnsupportedFunction
--- PASS: TestReadSourceRangeRejectsUnsupportedFunction (0.00s)
=== RUN   TestReadSourceRangeConnectionFailure
--- PASS: TestReadSourceRangeConnectionFailure (0.00s)
=== RUN   TestRuntimeRepeatsWithoutOverlap
--- PASS: TestRuntimeRepeatsWithoutOverlap (0.13s)
=== RUN   TestRuntimeRecordsCycleErrorAndContinues
--- PASS: TestRuntimeRecordsCycleErrorAndContinues (0.04s)
=== RUN   TestRuntimeCancelStopsCleanly
--- PASS: TestRuntimeCancelStopsCleanly (0.00s)
=== RUN   TestSaveApplyRejectsForeignReservationBeforeMutation
--- PASS: TestSaveApplyRejectsForeignReservationBeforeMutation (0.00s)
=== RUN   TestConfigRootFromEnv
--- PASS: TestConfigRootFromEnv (0.00s)
=== RUN   TestSaveLoadRoundTrip
--- PASS: TestSaveLoadRoundTrip (0.00s)
=== RUN   TestInvalidConfigRejectedBeforePersistence
--- PASS: TestInvalidConfigRejectedBeforePersistence (0.00s)
=== RUN   TestInvalidReplacementLeavesPriorBytesUnchanged
--- PASS: TestInvalidReplacementLeavesPriorBytesUnchanged (0.00s)
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_port_zero
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_unit_too_high
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_function_invalid
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_count_zero
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_range_overflow
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/poll_interval_zero
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_port_zero
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_unit_too_high
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_area_invalid
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_count_zero
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_range_overflow
=== RUN   TestValidateConfigRejectsInvalidRangesAndRequiredValues/count_mismatch
--- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_port_zero (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_unit_too_high (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_function_invalid (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_count_zero (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/source_range_overflow (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/poll_interval_zero (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_port_zero (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_unit_too_high (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_area_invalid (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_count_zero (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/destination_range_overflow (0.00s)
    --- PASS: TestValidateConfigRejectsInvalidRangesAndRequiredValues/count_mismatch (0.00s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/replicator	13.376s
=== RUN   TestUnixRuntimeListenerMatchesRelayAndProtectsLiveOwner
--- PASS: TestUnixRuntimeListenerMatchesRelayAndProtectsLiveOwner (0.02s)
=== RUN   TestUnixRuntimeListenerReplacesStaleSocket
--- PASS: TestUnixRuntimeListenerReplacesStaleSocket (0.00s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/replicator/cmd/modbus-replicator-runtime	1.036s

real	0m30.620s
user	1m6.466s
sys	0m13.754s
TEST_EXIT=0
```

Matches the pre-existing report above: original exit `TEST_EXIT=0`, elapsed `real 0m30.620s`, both packages `ok` (`github.com/tamzrod/MCS.OSJS/replicator 13.376s`, `github.com/tamzrod/MCS.OSJS/replicator/cmd/modbus-replicator-runtime 1.036s`), 59 `=== RUN` and 59 `--- PASS` (including the three required named tests), no `FAIL`, no `WARNING: DATA RACE`, no panic, original post-test tree clean. Nothing unexpected occurred. Verdict PASS — complete original transcript is available and matches.
