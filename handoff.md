# Handoff

## Current

ACTIVE: REP-005 — Single-Range Replication Cycle
State: IMPLEMENTED — AWAITING JR TEST

REP-004 is complete and archived. REP-005 remains ACTIVE until JR returns verification evidence and the coding agent reviews it.

## REP-004 Completion

JR verification PASS was accepted for REP-004:

- tested commit: `49417dffd3a07cedc2500d8ecbf318aa65ba36a9`;
- clean working tree before test;
- `gofmt -l .` produced no output;
- persisted-config source-reader integration passed;
- FC3 / FC4 protocol tests passed;
- error-path tests passed;
- full `go test -count=1 ./...` passed;
- `go vet ./...` passed with no diagnostics;
- Go 1.27.1 was used from user-local `~/.local/go` only (no global install; sandbox-local test tooling per Operation CWAL; manually pinned by the previous JR sync step).

The missing closing parenthesis in the prior JR prose note is corrected above. Going forward, the coding agent will correct harmless JR report punctuation during the next handoff update instead of sending JR back for prose-only fixes.

## REP-005 Implementation Summary

REP-005 implementation is now on `main` under `replicator/`.

Implemented:

- `Store.RunOnce()` for one complete read → write replication invocation;
- persisted REP-003 config loading and validation before the cycle;
- REP-004 Modbus source read reuse;
- source/destination count equality enforcement for 1:1 mapping;
- no cross-area mapping: FC3 source writes FC3 destination and FC4 source writes FC4 destination;
- shared `mma2composer` use with producer identity `replicator`;
- destination reservation ownership collision protection before composition;
- preservation of foreign reservations while replacing Replicator-owned reservation state;
- destination MMA2 memory composition for the configured register area/start/count;
- shared `mma2raw` client use for unchanged register writes;
- explicit cycle result returned only after successful raw-ingest acknowledgement;
- source-read, ownership, validation, mapping, and raw-ingest errors return without falsely reporting completion;
- no scheduler, recurring polling, scaling, endian conversion, multiple ranges/devices, or UI.

Files added/changed for REP-005:

- `replicator/go.mod`
- `replicator/cycle.go`
- `replicator/cycle_test.go`

Workflow state:

- REP-004 archived under `workflow/archive/rep-004-modbus-source-reader.md`;
- REP-005 promoted from `QUEUED` to `ACTIVE`;
- REP-006 remains next and must not be advanced by JR.

## JR TEST TASK — REP-005

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

### 3. Successful 1:1 replication-cycle test

```bash
go test -count=1 -run '^TestRunOnceCopiesConfiguredRegisters$' .
```

Expected: exit code 0 and PASS.

This test must prove one invocation reads controlled FC3 source values, composes a Replicator-owned MMA2 destination reservation, sends Raw Ingest v1 to the configured destination start/count, and preserves the exact register values unchanged.

### 4. Foreign-owner collision protection

```bash
go test -count=1 -run '^TestRunOnceRejectsForeignOwnedDestination$' .
```

Expected: exit code 0 and PASS.

This verifies an existing foreign-owned `(port, unit_id)` reservation is rejected and remains unchanged.

### 5. No cross-area mapping

```bash
go test -count=1 -run '^TestValidateCycleMappingRejectsCrossArea$' .
```

Expected: exit code 0 and PASS.

This verifies the first milestone does not silently transform FC3 ↔ FC4 mappings.

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

Successful replication cycle:
Command: go test -count=1 -run '^TestRunOnceCopiesConfiguredRegisters$' .
Exit/result: <result>
Output:
<exact output>

Foreign-owner collision:
Command: go test -count=1 -run '^TestRunOnceRejectsForeignOwnedDestination$' .
Exit/result: <result>
Output:
<exact output>

Cross-area rejection:
Command: go test -count=1 -run '^TestValidateCycleMappingRejectsCrossArea$' .
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
git commit -m "JR report REP-005 verification"
git push origin main
```

JR must stop after the report push. Do not fix failures. Do not archive REP-005. Do not advance REP-006.

## JR TEST REPORT

Verdict: PASS
Tested commit: 7eda5d964c795a0f2576e163ab5a599f7359fee7

Git status before test:
(empty)

Formatting check:
Command: gofmt -l .
Exit/result: exit code 0
Output:
(no output — no files listed)

Successful replication cycle:
Command: go test -count=1 -run '^TestRunOnceCopiesConfiguredRegisters$' .
Exit/result: exit code 0 — ok github.com/tamzrod/MCS.OSJS/replicator  0.006s
Output:
ok      github.com/tamzrod/MCS.OSJS/replicator  0.006s

Foreign-owner collision:
Command: go test -count=1 -run '^TestRunOnceRejectsForeignOwnedDestination$' .
Exit/result: exit code 0 — ok github.com/tamzrod/MCS.OSJS/replicator  0.006s
Output:
ok      github.com/tamzrod/MCS.OSJS/replicator  0.006s

Cross-area rejection:
Command: go test -count=1 -run '^TestValidateCycleMappingRejectsCrossArea$' .
Exit/result: exit code 0 — ok github.com/tamzrod/MCS.OSJS/replicator  0.004s
Output:
ok      github.com/tamzrod/MCS.OSJS/replicator  0.004s

Full Replicator tests:
Command: go test -count=1 ./...
Exit/result: exit code 0 — ok github.com/tamzrod/MCS.OSJS/replicator  0.011s
Output:
ok      github.com/tamzrod/MCS.OSJS/replicator  0.011s

Go vet:
Command: go vet ./...
Exit/result: exit code 0
Output:
(no output — no diagnostics)

Unexpected behavior:
None. Note: Go 1.27.1 toolchain was used from user-local ~/.local/go (no global install; sandbox-local test tooling per Operation CWAL; manually pinned by the previous JR sync step.
