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

Verdict: **FAIL**
Reason: replicator exact suite returned 1; stopped without retry.
UTC: 2026-09-20T03:14:36.598820+00:00
Worktree: `/home/sysadmin/apps/MCS.OSJS-jr`; source checkpoint: `949d7a8eb328dfa66f8359ffe81cfe359b8db2b7`; test HEAD / original remote: `1f5c30856d546274e9fea7e1c9a1ef43297fdbfc`.
Runner: `python3 workflow/cwal/umig-em-003-t.py` (one approved invocation).
Safety: offline Go modules, MCS_RUN_E2E != 1, no Docker, sudo or operator endpoints.
Test modules executed in order: mma2composer, simulator, replicator; unrun: none.
Preflight and post-check stdout/stderr plus exit codes (exact commands are recorded below):

### pwd: exit 0

Command: `pwd -P`

```text
/home/sysadmin/apps/MCS.OSJS-jr
```

### status: exit 0

Command: `git status --porcelain --untracked-files=all`

```text
(no output)
```

### HEAD: exit 0

Command: `git rev-parse HEAD`

```text
1f5c30856d546274e9fea7e1c9a1ef43297fdbfc
```

### origin/main: exit 0

Command: `git rev-parse origin/main`

```text
1f5c30856d546274e9fea7e1c9a1ef43297fdbfc
```

### worktrees: exit 0

Command: `git worktree list --porcelain`

```text
worktree /home/sysadmin/apps/MCS.OSJS
HEAD 1a04664e5fdc21e3bd323a97a4f12f3f9d39abdd
branch refs/heads/main

worktree /home/sysadmin/apps/MCS.OSJS-jr
HEAD 1f5c30856d546274e9fea7e1c9a1ef43297fdbfc
detached
```

### remote: exit 0

Command: `git ls-remote --exit-code origin refs/heads/main`

```text
1f5c30856d546274e9fea7e1c9a1ef43297fdbfc	refs/heads/main
```

### source ancestor: exit 0

Command: `git merge-base --is-ancestor 949d7a8eb328dfa66f8359ffe81cfe359b8db2b7 HEAD`

```text
(no output)
```

### source diff: exit 0

Command: `git diff --name-only 949d7a8eb328dfa66f8359ffe81cfe359b8db2b7 HEAD`

```text
handoff.md
operation cwal.md
workflow/active_work/umig-em-003-shared-lock.md
workflow/active_work/umig-em-003-t-shared-lock.md
workflow/archive/umig-em-003-shared-lock.md
workflow/cwal/umig-em-003-t.py
```

### system: exit 0

Command: `uname -s`

```text
Linux
```

### Go: exit 0

Command: `go version`

```text
go version go1.26.0 linux/amd64
```

### disk: exit 0

Command: `df -Pk . /home/sysadmin /tmp`

```text
Filesystem     1024-blocks      Used Available Capacity Mounted on
/dev/nvme0n1p2   490048472 394824604  70257264      85% /
/dev/nvme0n1p2   490048472 394824604  70257264      85% /
tmpfs             23817752     68628  23749124       1% /tmp
```

### TEST mma2composer: exit 0

Command: `timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...`

```text
=== RUN   TestRebuildRetainsListenerAndForeignAdvancedSettings
--- PASS: TestRebuildRetainsListenerAndForeignAdvancedSettings (0.00s)
=== RUN   TestProducerIdentityAndForeignCollision
--- PASS: TestProducerIdentityAndForeignCollision (0.00s)
=== RUN   TestFirstComeFirstSavePersistsProducer
--- PASS: TestFirstComeFirstSavePersistsProducer (0.00s)
=== RUN   TestDropOneReservationUsesRequestedPort
--- PASS: TestDropOneReservationUsesRequestedPort (0.00s)
=== RUN   TestCommitRestoresConfigWhenOwnersReplaceFails
--- PASS: TestCommitRestoresConfigWhenOwnersReplaceFails (0.00s)
=== RUN   TestWriterLockConcurrentForeignOwnershipConflict
--- PASS: TestWriterLockConcurrentForeignOwnershipConflict (0.02s)
=== RUN   TestWriterLockTimeoutAndRelease
--- PASS: TestWriterLockTimeoutAndRelease (0.06s)
=== RUN   TestWriterLockCrashRelease
--- PASS: TestWriterLockCrashRelease (1.02s)
=== RUN   TestWriterLockRejectsAmbiguousPath
--- PASS: TestWriterLockRejectsAmbiguousPath (0.00s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/mma2composer	2.124s
```

### TEST simulator: exit 0

Command: `timeout 360s env GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly go test -race -count=1 -timeout=300s -v ./...`

```text
=== RUN   TestAdvancedSettingsRoundTripAndCompose
--- PASS: TestAdvancedSettingsRoundTripAndCompose (0.01s)
=== RUN   TestSchedulerApplierReadyGatePreventsArmingAndFalseHealth
--- PASS: TestSchedulerApplierReadyGatePreventsArmingAndFalseHealth (0.03s)
=== RUN   TestStructuralApplyAndStopDoNotControlIndependentMMA2
--- PASS: TestStructuralApplyAndStopDoNotControlIndependentMMA2 (0.03s)
=== RUN   TestSchedulerApplierRuntimeStatusMatchesTimingAndPoints
--- PASS: TestSchedulerApplierRuntimeStatusMatchesTimingAndPoints (0.03s)
=== RUN   TestApplyRouterStructuralTimingAndRejectedPaths
--- PASS: TestApplyRouterStructuralTimingAndRejectedPaths (0.01s)
=== RUN   TestApplyRouterRejectsInvalidBeforeConsumers
--- PASS: TestApplyRouterRejectsInvalidBeforeConsumers (0.00s)
=== RUN   TestBootRestoreArmsEnabledSchedulesWithoutRestartRequest
--- PASS: TestBootRestoreArmsEnabledSchedulesWithoutRestartRequest (0.04s)
=== RUN   TestBootRestoreSurfacesUnavailableMMA2WithoutRestart
--- PASS: TestBootRestoreSurfacesUnavailableMMA2WithoutRestart (0.16s)
=== RUN   TestSaveAndComposeFreeReservationPersistsOwner
--- PASS: TestSaveAndComposeFreeReservationPersistsOwner (0.01s)
=== RUN   TestComposeRejectsForeignCollisionUnchanged
--- PASS: TestComposeRejectsForeignCollisionUnchanged (0.00s)
=== RUN   TestSaveAndComposeUpdatesOwnReservationPreservesForeign
--- PASS: TestSaveAndComposeUpdatesOwnReservationPreservesForeign (0.01s)
=== RUN   TestDeleteAndComposeRemovesOnlyOwn
--- PASS: TestDeleteAndComposeRemovesOnlyOwn (0.01s)
=== RUN   TestDeleteAndComposeRejectsForeignOwned
--- PASS: TestDeleteAndComposeRejectsForeignOwned (0.00s)
=== RUN   TestSaveAndComposeRejectsInvalidBeforeAnyPersist
--- PASS: TestSaveAndComposeRejectsInvalidBeforeAnyPersist (0.00s)
=== RUN   TestComposeDocumentIncludesOnlyEnabledAndPreservesForeign
--- PASS: TestComposeDocumentIncludesOnlyEnabledAndPreservesForeign (0.00s)
=== RUN   TestComposeDocumentCollisionAndInvalidCandidateLeaveFilesUnchanged
=== RUN   TestComposeDocumentCollisionAndInvalidCandidateLeaveFilesUnchanged/foreign_collision
=== RUN   TestComposeDocumentCollisionAndInvalidCandidateLeaveFilesUnchanged/invalid_complete_candidate
--- PASS: TestComposeDocumentCollisionAndInvalidCandidateLeaveFilesUnchanged (0.00s)
    --- PASS: TestComposeDocumentCollisionAndInvalidCandidateLeaveFilesUnchanged/foreign_collision (0.00s)
    --- PASS: TestComposeDocumentCollisionAndInvalidCandidateLeaveFilesUnchanged/invalid_complete_candidate (0.00s)
=== RUN   TestEndToEndSimulatorMMA2Architecture
    e2e_test.go:84: set MCS_RUN_E2E=1 for real MMA2 verification
--- SKIP: TestEndToEndSimulatorMMA2Architecture (0.00s)
=== RUN   TestNoneSimulationKeepsAllAreasAllocatedWithoutGeneration
--- PASS: TestNoneSimulationKeepsAllAreasAllocatedWithoutGeneration (0.03s)
=== RUN   TestMixedNoneAndRandomOnlySchedulesRandomArea
--- PASS: TestMixedNoneAndRandomOnlySchedulesRandomArea (0.01s)
=== RUN   TestNoneRuntimeStatusSeparatesMemoryAndGeneratorHealth
--- PASS: TestNoneRuntimeStatusSeparatesMemoryAndGeneratorHealth (0.00s)
=== RUN   TestDisabledNoneAndMixedWaitingStatus
--- PASS: TestDisabledNoneAndMixedWaitingStatus (0.00s)
=== RUN   TestRawIngestClientSendOK
--- PASS: TestRawIngestClientSendOK (0.00s)
=== RUN   TestRawIngestClientSendRejectsUnconfiguredFC
--- PASS: TestRawIngestClientSendRejectsUnconfiguredFC (0.00s)
=== RUN   TestRawIngestClientSendResponseError
--- PASS: TestRawIngestClientSendResponseError (0.00s)
=== RUN   TestStructuralApplyRequestsRestartAndClearsAfterReadiness
--- PASS: TestStructuralApplyRequestsRestartAndClearsAfterReadiness (0.03s)
=== RUN   TestStructuralApplyRejectedCommitWritesNoRestartRequest
--- PASS: TestStructuralApplyRejectedCommitWritesNoRestartRequest (0.00s)
=== RUN   TestStructuralApplyRestartTimeoutReportsFailureAndLeavesPendingRequest
--- PASS: TestStructuralApplyRestartTimeoutReportsFailureAndLeavesPendingRequest (0.34s)
=== RUN   TestWaitMMA2Ready
--- PASS: TestWaitMMA2Ready (0.21s)
=== RUN   TestRuntimeServiceStatusContractCarriesOperatorStatesAndDiagnostic
--- PASS: TestRuntimeServiceStatusContractCarriesOperatorStatesAndDiagnostic (0.00s)
=== RUN   TestRuntimeServiceStatusRequestErrorsRemainStable
=== RUN   TestRuntimeServiceStatusRequestErrorsRemainStable/missing_name
=== RUN   TestRuntimeServiceStatusRequestErrorsRemainStable/invalid_payload
=== RUN   TestRuntimeServiceStatusRequestErrorsRemainStable/unknown_device
--- PASS: TestRuntimeServiceStatusRequestErrorsRemainStable (0.00s)
    --- PASS: TestRuntimeServiceStatusRequestErrorsRemainStable/missing_name (0.00s)
    --- PASS: TestRuntimeServiceStatusRequestErrorsRemainStable/invalid_payload (0.00s)
    --- PASS: TestRuntimeServiceStatusRequestErrorsRemainStable/unknown_device (0.00s)
=== RUN   TestRuntimeServiceLoadApplyStatusAndIdempotency
--- PASS: TestRuntimeServiceLoadApplyStatusAndIdempotency (0.00s)
=== RUN   TestRuntimeTransportFraming
--- PASS: TestRuntimeTransportFraming (0.01s)
=== RUN   TestLiveRuntimeContractAppliesTimingRejectsConflictAndReportsUnavailable
--- PASS: TestLiveRuntimeContractAppliesTimingRejectsConflictAndReportsUnavailable (0.18s)
=== RUN   TestRuntimeStatusDisabledAndUnarmedDevicesAreNotRunning
--- PASS: TestRuntimeStatusDisabledAndUnarmedDevicesAreNotRunning (0.00s)
=== RUN   TestRuntimeStatusRequiresAcceptedRawIngestAndRecoversFromFailure
--- PASS: TestRuntimeStatusRequiresAcceptedRawIngestAndRecoversFromFailure (0.20s)
=== RUN   TestRuntimeStatusMMA2UnavailabilityPreventsRunning
--- PASS: TestRuntimeStatusMMA2UnavailabilityPreventsRunning (0.03s)
=== RUN   TestSchedulerRunsAllFCsConcurrentlyAtOwnCadence
--- PASS: TestSchedulerRunsAllFCsConcurrentlyAtOwnCadence (0.13s)
=== RUN   TestSchedulerIntervalChangeUpdatesScheduleWithoutRestart
--- PASS: TestSchedulerIntervalChangeUpdatesScheduleWithoutRestart (0.03s)
=== RUN   TestSaveDocumentMultiDeviceRoundTrip
--- PASS: TestSaveDocumentMultiDeviceRoundTrip (0.00s)
=== RUN   TestSaveDocumentRejectsInvalidDeviceWithoutReplacing
--- PASS: TestSaveDocumentRejectsInvalidDeviceWithoutReplacing (0.00s)
=== RUN   TestSaveDocumentAddDuplicateDeleteReshape
--- PASS: TestSaveDocumentAddDuplicateDeleteReshape (0.00s)
=== RUN   TestConfigRootFromEnvRefusesInventedPath
--- PASS: TestConfigRootFromEnvRefusesInventedPath (0.00s)
=== RUN   TestRoundTripExactValues
--- PASS: TestRoundTripExactValues (0.00s)
=== RUN   TestInvalidInputsDoNotReplaceLastValid
=== RUN   TestInvalidInputsDoNotReplaceLastValid/port_0
=== RUN   TestInvalidInputsDoNotReplaceLastValid/unit_id_256
=== RUN   TestInvalidInputsDoNotReplaceLastValid/fc3_start+count_overflow
=== RUN   TestInvalidInputsDoNotReplaceLastValid/empty_name
--- PASS: TestInvalidInputsDoNotReplaceLastValid (0.00s)
    --- PASS: TestInvalidInputsDoNotReplaceLastValid/port_0 (0.00s)
    --- PASS: TestInvalidInputsDoNotReplaceLastValid/unit_id_256 (0.00s)
    --- PASS: TestInvalidInputsDoNotReplaceLastValid/fc3_start+count_overflow (0.00s)
    --- PASS: TestInvalidInputsDoNotReplaceLastValid/empty_name (0.00s)
=== RUN   TestSimulatorComposeRejectsBusySharedWriterLock
--- PASS: TestSimulatorComposeRejectsBusySharedWriterLock (5.01s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/simulator	7.585s
=== RUN   TestWatchDocumentReloadsNoneSchedule
--- PASS: TestWatchDocumentReloadsNoneSchedule (0.26s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/simulator/cmd/modbus-simulator-runtime	1.272s
```

### TEST replicator: exit 1

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
--- PASS: TestSimulatorToReplicatorE2E (4.74s)
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
=== RUN   TestApplyRestartFailureRestoresPreviousPollers
    multiblock_regression_test.go:73: previous poller was not restored after failed apply: replicator.DeviceRuntimeStatus{Comms:map[string]string{"mma2":"UNKNOWN", "modbus":"UNKNOWN", "network":"UNKNOWN", "tcp":"UNKNOWN"}, Network:replicator.CommsObservation{State:"", Outcome:"", Endpoint:"", Error:"", ObservedAt:"", ActivityAt:"", LastSuccessAt:"", ExceptionCode:(*uint8)(nil)}, Name:"PLC-recovery", Enabled:true, Running:false, Cycles:0x0, Source:"WAITING", LastPoll:"", LastError:"", Blocks:[]replicator.BlockRuntimeStatus{replicator.BlockRuntimeStatus{CycleComms:replicator.CycleComms{Network:replicator.CommsObservation{State:"", Outcome:"", Endpoint:"", Error:"", ObservedAt:"", ActivityAt:"", LastSuccessAt:"", ExceptionCode:(*uint8)(nil)}, TCP:replicator.CommsObservation{State:"", Outcome:"", Endpoint:"", Error:"", ObservedAt:"", ActivityAt:"", LastSuccessAt:"", ExceptionCode:(*uint8)(nil)}, Modbus:replicator.CommsObservation{State:"", Outcome:"", Endpoint:"", Error:"", ObservedAt:"", ActivityAt:"", LastSuccessAt:"", ExceptionCode:(*uint8)(nil)}, MMA2:replicator.CommsObservation{State:"", Outcome:"", Endpoint:"", Error:"", ObservedAt:"", ActivityAt:"", LastSuccessAt:"", ExceptionCode:(*uint8)(nil)}}, Function:0x3, Start:0xa, Count:0x4, Index:0, Running:false, Cycles:0x0, Source:"WAITING", LastPoll:"", LastError:""}}}
--- FAIL: TestApplyRestartFailureRestoresPreviousPollers (0.59s)
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
--- PASS: TestRuntimeRepeatsWithoutOverlap (0.12s)
=== RUN   TestRuntimeRecordsCycleErrorAndContinues
--- PASS: TestRuntimeRecordsCycleErrorAndContinues (0.04s)
=== RUN   TestRuntimeCancelStopsCleanly
--- PASS: TestRuntimeCancelStopsCleanly (0.01s)
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
FAIL
FAIL	github.com/tamzrod/MCS.OSJS/replicator	10.863s
=== RUN   TestUnixRuntimeListenerMatchesRelayAndProtectsLiveOwner
--- PASS: TestUnixRuntimeListenerMatchesRelayAndProtectsLiveOwner (0.00s)
=== RUN   TestUnixRuntimeListenerReplacesStaleSocket
--- PASS: TestUnixRuntimeListenerReplacesStaleSocket (0.00s)
PASS
ok  	github.com/tamzrod/MCS.OSJS/replicator/cmd/modbus-replicator-runtime	1.011s
FAIL
```

### post status: exit 0

Command: `git status --porcelain --untracked-files=all`

```text
(no output)
```

### post HEAD: exit 0

Command: `git rev-parse HEAD`

```text
1f5c30856d546274e9fea7e1c9a1ef43297fdbfc
```

### pre-report remote: exit 0

Command: `git ls-remote --exit-code origin refs/heads/main`

```text
1f5c30856d546274e9fea7e1c9a1ef43297fdbfc	refs/heads/main
```

Scope: Go TEST only; no live VERIFY, Electron or production claim.
JR STOP: no source edits, task advancement or further tests authorized.
