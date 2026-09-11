# Handoff

## Current

ACTIVE: REP-004 — Modbus Source Reader
State: IMPLEMENTED — AWAITING JR TEST

REP-003 is complete and archived. REP-004 remains ACTIVE until JR returns verification evidence and the coding agent reviews it.

## REP-003 Completion

JR verification PASS was accepted for REP-003:

- tested commit: `c11ae364652e1f2b2e65b9d74397545050b25edb`;
- clean working tree before test;
- `gofmt -l .` produced no output;
- `go test -count=1 ./...` passed;
- `go vet ./...` passed with no diagnostics;
- Go 1.27.1 was used from user-local `~/.local/go` only (no global install; previously prepared in the sandbox per Operation CWAL guidelines).

The missing closing parenthesis in the prior JR prose note is corrected above. It was report-text only and did not affect the PASS evidence.

## REP-004 Implementation Summary

REP-004 implementation is now on `main` under `replicator/`.

Implemented:

- `RegisterValues`, a small producer-neutral register payload for later MMA2 writing;
- `Store.ReadConfiguredSource()`, which loads and validates the persisted REP-003 config before reading;
- deterministic one-shot Modbus TCP source reads;
- FC3 Holding Register reads;
- FC4 Input Register reads;
- Modbus TCP MBAP request/response validation;
- Unit ID, function, transaction ID, protocol ID, response-length, and byte-count validation;
- Modbus exception responses returned as errors;
- connection/read failures returned as errors;
- no MMA2 destination writes, retries/backoff loop, polling loop, scaling, byte swapping, or UI;
- controlled TCP protocol tests using Simulator-equivalent FC3/FC4 register semantics;
- persisted-config integration test proving source host/port/unit/function/start/count come through the REP-003 store.

Files added/changed for REP-004:

- `replicator/reader.go`
- `replicator/reader_test.go`

Workflow state:

- REP-003 archived under `workflow/archive/rep-003-replicator-device-config.md`;
- REP-004 promoted from `QUEUED` to `ACTIVE`;
- REP-005 remains next and must not be advanced by JR.

## JR TEST TASK — REP-004

JR role: TEST AND REPORT ONLY.

Do not modify Go source, Active Work, ICC, planning files, workflow status, or implementation files.

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

### 3. Persisted-config source-reader integration test

```bash
go test -count=1 -run '^TestReadConfiguredSourceUsesPersistedConfig$' .
```

Expected: exit code 0 and PASS.

This proves REP-004 reads its Modbus source endpoint/range from the persisted REP-003 configuration and obtains the controlled register values.

### 4. FC3 / FC4 protocol tests

```bash
go test -count=1 -run '^TestReadSourceRangeFC(3|4)$' .
```

Expected: exit code 0 and PASS.

This verifies the first supported register areas use correct Modbus TCP read semantics and return exact register values.

### 5. Error-path tests

```bash
go test -count=1 -run '^TestReadSourceRange(RejectsNonRegisterFunction|ConnectionFailure)$' .
```

Expected: exit code 0 and PASS.

This verifies unsupported bit functions and connection failure are returned as errors; REP-004 contains no destination-write behavior.

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

Persisted-config integration:
Command: go test -count=1 -run '^TestReadConfiguredSourceUsesPersistedConfig$' .
Exit/result: <result>
Output:
<exact output>

FC3 / FC4 protocol:
Command: go test -count=1 -run '^TestReadSourceRangeFC(3|4)$' .
Exit/result: <result>
Output:
<exact output>

Error paths:
Command: go test -count=1 -run '^TestReadSourceRange(RejectsNonRegisterFunction|ConnectionFailure)$' .
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
git commit -m "JR report REP-004 verification"
git push origin main
```

JR must stop after the report push. Do not fix failures. Do not archive REP-004. Do not advance REP-005.

## JR TEST REPORT

Verdict: PENDING
