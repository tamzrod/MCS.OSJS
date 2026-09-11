# Handoff

## Current

ACTIVE: REP-006 — Replicator Poll Loop
State: IMPLEMENTED — AWAITING JR TEST

REP-005 is complete and archived. REP-006 remains ACTIVE until JR returns verification evidence and the coding agent reviews it.

## REP-005 Completion

JR verification PASS was accepted for REP-005:

- tested commit: `7eda5d964c795a0f2576e163ab5a599f7359fee7`;
- clean working tree before test;
- `gofmt -l .` produced no output;
- successful exact-value 1:1 replication-cycle test passed;
- foreign-owner collision protection passed;
- cross-area rejection passed;
- full `go test -count=1 ./...` passed;
- `go vet ./...` passed with no diagnostics;
- Go 1.27.1 was used from user-local `~/.local/go` only (no global install; sandbox-local test tooling per Operation CWAL; manually pinned by the previous JR sync step).

The prior JR report's missing closing parenthesis was corrected here automatically. It was prose-only and did not affect the PASS evidence.

## REP-006 Implementation Summary

REP-006 implementation is now on `main` under `replicator/`.

Implemented:

- `Runtime` poll-loop wrapper around the proven REP-005 single-cycle behavior;
- persisted configuration validation before runtime start;
- destination ownership/composition readiness check before the runtime is marked running;
- immediate first cycle followed by one cycle per configured `poll_interval_ms` tick;
- strictly serial cycle execution, so slow cycles cannot overlap;
- deterministic runtime policy: record a failed cycle truthfully, then continue on the next normal tick with no retry/backoff burst;
- race-safe `Snapshot()` runtime state with running flag, cycle count, last success/error, and last cycle timestamps;
- successful cycles clear stale error state;
- context cancellation stops the loop cleanly after any currently executing serial cycle returns;
- no advanced backoff, multiple workers/devices/ranges, UI, or metrics/history storage.

Files added for REP-006:

- `replicator/runtime.go`
- `replicator/runtime_test.go`

Workflow state:

- REP-005 archived under `workflow/archive/rep-005-single-range-replication-cycle.md`;
- REP-006 promoted from `QUEUED` to `ACTIVE`;
- REP-007 remains next and must not be advanced by JR.

## JR TEST TASK — REP-006

JR role: TEST AND REPORT ONLY.

Do not modify Go source, module files, Active Work, ICC, planning files, workflow status, or implementation files.

JR may prepare sandbox-local test tooling according to `operation cwal.md` if needed. Do not modify project manifests or source merely to install a test tool.

JR is authorized to modify **only `handoff.md`** to record the test report, then commit and push that report so the coding agent can inspect it.

### 1. Sync and repository state

From repository root:

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
```

Expected:

- pull succeeds;
- `git status --short` is empty.

If the working tree is not clean, do not clean or restore it. Record `BLOCKED` with exact status output.

### 2. Formatting check

```bash
cd replicator
gofmt -l .
```

Expected: no output.

If files are listed, record FAIL. Do not run `gofmt -w`.

### 3. Repeated serial execution / no overlap

```bash
go test -count=1 -run '^TestRuntimeRepeatsWithoutOverlap$' .
```

Expected: exit code 0 and PASS.

This verifies a cycle slower than the configured interval still runs serially with maximum concurrency of one, repeats multiple times, and leaves truthful success/stopped state.

### 4. Error state and continuation policy

```bash
go test -count=1 -run '^TestRuntimeRecordsCycleErrorAndContinues$' .
```

Expected: exit code 0 and PASS.

This verifies failed cycles are recorded as failures with the exact last error and the runtime continues only on subsequent normal ticks.

### 5. Clean cancellation

```bash
go test -count=1 -run '^TestRuntimeCancelStopsCleanly$' .
```

Expected: exit code 0 and PASS.

This verifies cancellation returns cleanly and the final runtime state is not falsely left RUNNING.

### 6. Full Replicator regression

```bash
go test -count=1 ./...
```

Expected: exit code 0 and package PASS.

### 7. Native Go verification

```bash
go vet ./...
```

Expected: exit code 0 with no diagnostics.

### 8. JR report

Replace everything below `## JR TEST REPORT` with the actual result. Preserve every other section exactly.

Use this structure:

```text
Verdict: PASS | FAIL | BLOCKED
Tested commit: <git rev-parse HEAD before report commit>

Git status before test:
<exact output>

Formatting check:
Command: gofmt -l .
Exit/result: <result>
Output:
<exact output>

Serial poll loop:
Command: go test -count=1 -run '^TestRuntimeRepeatsWithoutOverlap$' .
Exit/result: <result>
Output:
<exact output>

Error state / continuation:
Command: go test -count=1 -run '^TestRuntimeRecordsCycleErrorAndContinues$' .
Exit/result: <result>
Output:
<exact output>

Clean cancellation:
Command: go test -count=1 -run '^TestRuntimeCancelStopsCleanly$' .
Exit/result: <result>
Output:
<exact output>

Full Replicator tests:
Command: go test -count=1 ./...
Exit/result: <result>
Output:
<exact output>

Go vet:
Command: go vet ./...
Exit/result: <result>
Output:
<exact output>

Unexpected behavior:
<none or exact observation>
```

After writing the report, JR may commit and push **handoff.md only**:

```bash
cd ..
git add handoff.md
git diff --cached -- handoff.md
git commit -m "JR report REP-006 verification"
git push origin main
```

JR must stop after the report push. Do not fix failures. Do not archive REP-006. Do not advance REP-007.

## JR TEST REPORT

Verdict: PENDING
