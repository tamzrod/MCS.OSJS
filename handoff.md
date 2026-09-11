# Handoff

## Current

ACTIVE SEQUENCE:
- REP-BLOCK-002 — Independent Pull Block Pollers
- REP-BLOCK-003 — Device / Pull Blocks Folder Tabs

State: RECTIFIED — READY FOR JR RETEST

REP-BLOCK-001 and REP-OWN-001 remain archived. Shared MMA2 ownership remains exact `(port, unit_id)` pair ownership.

## Rectifications after first JR pass

The first JR pass verified the rendered UI, independent block cadence, persistence, and pair ownership, but failed the source-format gate because `document_test.go` was not gofmt-aligned. That formatting defect is corrected.

Two implementation defects found during review were also corrected:

1. Failed Save & Apply before successful MMA2 activation now restores the previously persisted Replicator pollers instead of leaving the device stopped.
2. Multiple Pull Blocks of the same FC may overlap or touch, but may not contain a gap. MMA2 exposes one contiguous area per FC for one `(port, unit_id)` memory; rejecting a disjoint same-FC shape prevents silently exposing destination registers that no Pull Block polls.

Examples:

```text
Allowed:
FC3 0..9
FC3 10..19

Rejected:
FC3 0..9
FC3 100..109
```

Different FC areas remain independent, for example FC3 `0..15` plus FC4 `100..103` is valid.

## Implemented behavior

### REP-BLOCK-002
- One Replicator device carries an ordered `pull_blocks` collection.
- Legacy flat source fields and prior single `pull_block` migrate into one collection entry.
- Each enabled Pull Block has its own runtime poll loop/ticker and scan rate.
- Endpoint, Source Unit ID, Enabled, Name, and destination reservation remain device-level.
- Per-block runtime status exposes index, running, cycles, source status, last poll, and last error.
- Same-FC Pull Blocks must form an overlap/contiguous destination area; gapped same-FC blocks are rejected before apply.
- Shared destination ownership remains one exact `(port, unit_id)` reservation per device.
- A failed pre-activation apply restores the previously persisted pollers.

### REP-BLOCK-003
- Existing left device tree/list remains unchanged as the device selector.
- Right-side configuration has exactly two classic folder-style tabs: `Device` and `Pull Blocks`.
- Blocks are cards/rows inside `Pull Blocks`; blocks are not tabs.
- Add Block / Duplicate Block / Delete Block modify only the Pull Block collection.
- Save & Apply / Discard remain device-level actions.
- Per-block runtime status is rendered in the Pull Blocks tab.

## JR TEST TASK — REP-BLOCK-002 + REP-BLOCK-003 RETEST

JR role: TEST AND REPORT ONLY. Do not fix source, modify Active Work/planning/ICC, or expand scope. Do not destroy Docker volumes.

JR may edit only the `## JR TEST REPORT` section and may commit/push only `handoff.md` after testing.

### 1. Sync and backend gate

From repository root:

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
cd replicator
gofmt -l .
go test -count=1 ./...
go vet ./...
```

Expected:
- clean tree before testing;
- `gofmt -l .` prints nothing;
- full Replicator tests PASS;
- vet exits 0.

Focus evidence must include:
- legacy single-block migration;
- ordered multi-block persistence;
- independent per-block cadence/status;
- `TestValidateRejectsDisjointSameFCPullBlocks` PASS;
- `TestValidateAllowsContiguousSameFCPullBlocks` PASS;
- `TestApplyRestartFailureRestoresPreviousPollers` PASS;
- exact pair ownership regressions still PASS.

### 2. Deploy current build

```bash
cd ../deploy
docker compose build osjs-shell modbus-replicator-runtime mma2
docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime
docker compose ps
```

Expected: all services running and osjs-shell healthy.

### 3. Rendered layout check

Open Modbus Replicator.

Required:
- left device tree/list/search/Add/Duplicate/Delete remains unchanged;
- right side has exactly two top-level folder tabs: `Device` and `Pull Blocks`;
- no per-block tab strip exists;
- block entries appear as cards/rows inside `Pull Blocks`.

### 4. Valid multi-block Save & Apply

Use one Replicator device with a representable block set, for example:

```text
Block 1: FC3 Start 0   Count 8  Scan 100 ms
Block 2: FC4 Start 100 Count 4  Scan 500 ms
```

Click rendered **Save & Apply**.

Required:
- apply succeeds through MMA2 restart/readiness;
- `devices.yaml` persists both blocks in order;
- one destination `(port, unit_id)` reservation exists;
- per-block status is independently visible;
- faster scan accumulates cycles faster when observable.

### 5. Same-FC representability guard

In the rendered Pull Blocks tab, configure two FC3 blocks with a gap, for example:

```text
FC3 Start 0   Count 10
FC3 Start 100 Count 10
```

Click rendered **Save & Apply**.

Required:
- apply is rejected;
- visible error states that same-FC blocks must overlap or be contiguous;
- no new Replicator document is persisted;
- previously valid runtime remains active.

Then change the second block to:

```text
FC3 Start 10 Count 10
```

Required: Save & Apply succeeds.

### 6. Failed restart recovery

Using a controlled test path from the backend regression, verify a missing/failed MMA2 restart acknowledgement causes Apply to fail while the previously persisted poller set is restored and running again. This backend behavior is covered by `TestApplyRestartFailureRestoresPreviousPollers`; do not deliberately break the live UI environment just to reproduce it manually.

### 7. Ownership regression

With Simulator owning `(5020,1)` confirm using rendered Save & Apply:
- Replicator `(5022,1)` succeeds when free;
- Replicator `(5020,2)` succeeds when free;
- Replicator `(5020,1)` rejects as exact-pair conflict.

Do not substitute direct socket/API apply calls for required rendered UI checks.

## JR TEST REPORT

Verdict: PENDING
