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

Verdict: **PASS**
Reason: Replicator race suite and all mandatory named tests passed.
UTC: 2026-09-20T03:49:19.236422+00:00
Source checkpoint: `538324a472eb15ff8ef97ee66f826fa02ba9462f`; tested activation HEAD/live main: `8ca9a4b6a6b7e28bd419ed710fbb4825da7e6b5f`; worktree: `/home/sysadmin/apps/MCS.OSJS-jr`.
Runner: `python3 workflow/cwal/umig-em-003-r-t.py` (one human-approved invocation).
Test scope: Replicator offline Go race suite ONLY; earlier Composer/Simulator results are historical, not rerun.
E2E disabled; test-owned ephemeral loopback fixtures only; no Docker, service or operator device action.
Complete command transcripts (merged original stdout/stderr) and exit codes:

### pwd — exit 0

Command: `pwd -P`

```text
/home/sysadmin/apps/MCS.OSJS-jr
```

### initial status — exit 0

Command: `git status --porcelain --untracked-files=all`

```text
(no output)
```

### HEAD — exit 0

Command: `git rev-parse HEAD`

```text
8ca9a4b6a6b7e28bd419ed710fbb4825da7e6b5f
```

### tracking — exit 0

Command: `git rev-parse origin/main`

```text
8ca9a4b6a6b7e28bd419ed710fbb4825da7e6b5f
```

### worktrees — exit 0

Command: `git worktree list --porcelain`

```text
worktree /home/sysadmin/apps/MCS.OSJS
HEAD 1a04664e5fdc21e3bd323a97a4f12f3f9d39abdd
branch refs/heads/main

worktree /home/sysadmin/apps/MCS.OSJS-jr
HEAD 8ca9a4b6a6b7e28bd419ed710fbb4825da7e6b5f
detached
```

### live remote — exit 0

Command: `git ls-remote --exit-code origin refs/heads/main`

```text
8ca9a4b6a6b7e28bd419ed710fbb4825da7e6b5f	refs/heads/main
```

### source ancestor — exit 0

Command: `git merge-base --is-ancestor 538324a472eb15ff8ef97ee66f826fa02ba9462f HEAD`

```text
(no output)
```

### source diff — exit 0

Command: `git diff --name-only 538324a472eb15ff8ef97ee66f826fa02ba9462f HEAD`

```text
handoff.md
workflow/active_work/umig-em-003-r-recovery-regression.md
workflow/active_work/umig-em-003-r-t-recovery-test.md
workflow/archive/umig-em-003-r-recovery-regression.md
workflow/cwal/umig-em-003-r-t.py
```

### platform — exit 0

Command: `uname -s`

```text
Linux
```

### Go version — exit 0

Command: `go version`

```text
go version go1.26.0 linux/amd64
```

### disk — exit 0

Command: `df -Pk . /home/sysadmin /tmp`

```text
Filesystem     1024-blocks      Used Available Capacity Mounted on
/dev/nvme0n1p2   490048472 394732112  70349756      85% /
/dev/nvme0n1p2   490048472 394732112  70349756      85% /
tmpfs             23817752     70784  23746968       1% /tmp
```

### TEST replicator — exit 0

Command: `timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...`

```text
=== RUN   TestAdvancedSettingsPersistComposeAndInherit
--- PASS: TestAdvancedSettingsPersistComposeAndInherit (0.01s)
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
--- PASS: TestSimulatorToReplicatorE2E (4.69s)
=== RUN   TestDestinationMemorySupportsFC1AndFC2
--- PASS: TestDestinationMemorySupportsFC1AndFC2 (0.00s)
=== RUN   TestValidateDeviceAllowsAllReadFunctions
--- PASS: TestValidateDeviceAllowsAllReadFunctions (0.00s)
=== RUN   TestRuntimeManagerApplyLifecycleAndStatus
--- PASS: TestRuntimeManagerApplyLifecycleAndStatus (0.19s)
=== RUN   TestValidateRejectsDisjointSameFCPullBlocks
--- PASS: TestValidateRejectsDisjointSameFCPullBlocks (0.00s)
=== RUN   TestValidateAllowsContiguousSameFCPullBlocks
--- PASS: TestValidateAllowsContiguousSameFCPullBlocks (0.00s)
=== RUN   TestApplyPreCommitFailureRestoresPreviousPollers
--- PASS: TestApplyPreCommitFailureRestoresPreviousPollers (0.01s)
=== RUN   TestApplyPostCommitRestartFailureStopsPreviousPollers
--- PASS: TestApplyPostCommitRestartFailureStopsPreviousPollers (0.08s)
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
--- PASS: TestRuntimeRepeatsWithoutOverlap (0.10s)
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
=== RUN   TestReplicatorComposeRejectsBusySharedWriterLock
--- PASS: TestReplicatorComposeRejectsBusySharedWriterLock (5.00s)
=== RUN   TestManagerCommittedUnacknowledgedRestartFailsClosed
--- PASS: TestManagerCommittedUnacknowledgedRestartFailsClosed (0.11s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/replicator	11.296s
=== RUN   TestUnixRuntimeListenerMatchesRelayAndProtectsLiveOwner
--- PASS: TestUnixRuntimeListenerMatchesRelayAndProtectsLiveOwner (0.00s)
=== RUN   TestUnixRuntimeListenerReplacesStaleSocket
--- PASS: TestUnixRuntimeListenerReplacesStaleSocket (0.00s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/replicator/cmd/modbus-replicator-runtime	1.013s
```

### post status — exit 0

Command: `git status --porcelain --untracked-files=all`

```text
(no output)
```

### post HEAD — exit 0

Command: `git rev-parse HEAD`

```text
8ca9a4b6a6b7e28bd419ed710fbb4825da7e6b5f
```

### pre-report remote — exit 0

Command: `git ls-remote --exit-code origin refs/heads/main`

```text
8ca9a4b6a6b7e28bd419ed710fbb4825da7e6b5f	refs/heads/main
```

Scope: Go TEST only; does not establish OS.js/Electron live VERIFY or production readiness.
JR STOP: no code or task changes and no additional tests authorized.
