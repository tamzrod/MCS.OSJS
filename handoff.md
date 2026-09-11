# Handoff

## Current

ACTIVE: REP-007 — Simulator-to-Replicator End-to-End Verification
State: READY FOR JR E2E TEST

REP-006 is complete and archived. REP-007 is the final task in the currently authorized Replicator sequence (`Next: none`).

## REP-006 Completion

JR verification PASS was accepted for REP-006:

- tested commit: `bffee36e5660e15875dbe912be0e8d56d0c93d0c`;
- clean working tree before test;
- `gofmt -l .` produced no output;
- repeated serial execution / no-overlap test passed;
- error-state / continuation-policy test passed;
- clean cancellation test passed;
- full `go test -count=1 ./...` passed;
- `go vet ./...` passed with no diagnostics;
- Go 1.27.1 was used from user-local `~/.local/go` only (no global install; sandbox-local test tooling per Operation CWAL; manually pinned by the previous JR sync step).

The prior JR report's missing closing parenthesis was corrected here automatically. It was prose-only and did not affect the PASS evidence.

## REP-007 Verification Harness

REP-007 is verification-only. No new Replicator product behavior is introduced.

Added test infrastructure:

- `replicator/e2e_test.go`;
- Replicator test dependency on the existing Simulator module.

The E2E test uses:

- the actual `simulator.Store` and `ComposeDocument()` path to create the source reservation under owner `simulator`;
- a distinct Replicator-owned destination reservation;
- the real MMA2 binary built from `MMA2/cmd/mma2`;
- the same shared Raw Ingest v1 path used by Simulator runtime values to seed/change the Simulator-owned source registers;
- the REP-006 Replicator poll loop;
- normal Modbus FC3 reads from the Replicator destination;
- explicit ownership checks before and after replication.

The test verifies initial source values copy unchanged, a later source value change propagates within the poll window, and Simulator/Replicator ownership remain distinct.

Workflow state:

- REP-006 archived under `workflow/archive/rep-006-replicator-poll-loop.md`;
- REP-007 promoted from `QUEUED` to `ACTIVE`;
- REP-007 has `Next: none` and JR must not invent further work.

## JR TEST TASK — REP-007

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

### 3. Simulator → Replicator end-to-end verification

```bash
go test -count=1 -run '^TestSimulatorToReplicatorE2E$' .
```

Expected: exit code 0 and PASS.

This focused test must prove all of the following in one workflow:

- Simulator creates/owns the source MMA2 reservation;
- Replicator creates/owns a distinct destination reservation;
- initial FC3 source register values are copied unchanged to the Replicator destination;
- destination is read through normal Modbus FC3;
- source values are changed again through Simulator's Raw Ingest transport path;
- the Replicator poll loop propagates the changed values within the polling window;
- both ownership entries remain present and distinct after the copy.

### 4. Full Replicator regression

```bash
go test -count=1 ./...
```

Expected: exit code 0 and package PASS.

### 5. Native Go verification

```bash
go vet ./...
```

Expected: exit code 0 with no diagnostics.

### 6. JR report

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

Simulator-to-Replicator E2E:
Command: go test -count=1 -run '^TestSimulatorToReplicatorE2E$' .
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
git commit -m "JR report REP-007 verification"
git push origin main
```

JR must stop after the report push. Do not fix failures. Do not archive REP-007. Do not invent or advance any next task because REP-007 has `Next: none`.

## JR TEST REPORT

Verdict: PENDING
