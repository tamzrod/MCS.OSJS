# Handoff

## Current

ACTIVE SEQUENCE:
- REP-BLOCK-002 — Independent Pull Block Pollers
- REP-BLOCK-003 — Device / Pull Blocks Folder Tabs

State: IMPLEMENTED — READY FOR JR/HUMAN RETEST

REP-BLOCK-001 and REP-OWN-001 remain archived. Shared MMA2 ownership remains exact `(port, unit_id)` pair ownership.

## Current implementation

### REP-BLOCK-002
- One Replicator device carries an ordered `pull_blocks` collection.
- Each Pull Block has independent FC / Start / Count / Scan Rate and its own poll loop/status.
- Pull Blocks now support all Modbus read areas:
  - FC1 Coils
  - FC2 Discrete Inputs
  - FC3 Holding Registers
  - FC4 Input Registers
- FC1/FC2 source bits are decoded from Modbus bit-packed responses and written through MMA2 Raw Ingest as bit values.
- FC3/FC4 continue to replicate uint16 register values.
- Destination MMA2 memory composes coils/discrete_inputs/holding_registers/input_registers from the configured blocks.
- The destination memory keeps the external Modbus access policy, so third-party clients can read the served FC1-FC4 memory.
- Same-FC Pull Blocks must overlap or touch; gapped same-FC blocks are rejected instead of exposing unpolled memory gaps.
- Failed pre-activation Save & Apply restores the previously persisted pollers.
- Human test already confirmed the Replicator destination is externally serving changing FC3 register values after the access-policy fix.

### Operational Status semantics
- `Owner: replicator` is ownership metadata only.
- `OWNED` is not an operational success status.
- Device `Status = OK` is driven by successful runtime replication cycles: destination memory is active, the source read succeeds, and Raw Ingest write succeeds.
- A failed source read or destination write reports non-OK/error runtime state.

### REP-BLOCK-003
- Left device tree/list remains unchanged.
- Right-side selected-device editor keeps exactly two folder tabs: `Device` and `Pull Blocks`.
- `Pull Blocks` is now a compact spreadsheet/table rather than large cards.
- One block = one row with columns:
  - `#`
  - `FC`
  - `Start`
  - `Count`
  - `Scan Rate (ms)`
  - `Status`
  - `Last Poll`
- FC selector contains FC1, FC2, FC3, and FC4.
- Add / Duplicate / Delete Block operate on the selected table row.

## JR TEST TASK — REP-BLOCK-002 + REP-BLOCK-003

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
- legacy migration and ordered multi-block persistence;
- independent block cadence/status;
- exact pair ownership regressions;
- failed-restart runtime restoration;
- same-FC contiguous/gap guard;
- `TestReadSourceRangeFC1AndFC2` PASS;
- `TestDestinationMemorySupportsFC1AndFC2` PASS;
- `TestValidateDeviceAllowsAllReadFunctions` PASS.

### 2. Deploy current build

```bash
cd ../deploy
docker compose build osjs-shell modbus-replicator-runtime mma2
docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime
docker compose ps
```

Expected: all services running and osjs-shell healthy.

### 3. Rendered UI layout

Open Modbus Replicator.

Required:
- left device tree/list remains unchanged;
- right side has exactly `Device` and `Pull Blocks` folder tabs;
- Pull Blocks renders as a compact spreadsheet/table, not block cards;
- table FC selector exposes FC1, FC2, FC3, FC4;
- several block rows fit vertically without card-sized wasted space.

### 4. FC1-FC4 end-to-end replication

Using the Simulator as source, create representative blocks for FC1-FC4 within configured Simulator ranges and Save & Apply.

Required for each FC:
1. source poll succeeds;
2. Replicator writes the unchanged value/bit set into MMA2 destination memory;
3. a real external Modbus client reads the Replicator destination using the corresponding FC;
4. returned values match the source after replication.

For FC1/FC2 verify boolean bit patterns, including a count that is not a multiple of 8 so bit packing/unpacking is exercised.

### 5. Operational Status

Required:
- Device tab shows `Owner: replicator` separately from operational Status;
- healthy end-to-end replication shows `Status: OK`, not `OWNED`;
- source failure or destination-write failure must not show `OK`;
- Pull Blocks table shows per-row runtime Status and Last Poll.

### 6. Same-FC guard and ownership regression

Verify a gapped same-FC pair is rejected, contiguous same-FC pair succeeds, and exact `(port, unit_id)` ownership behavior remains unchanged:
- `(5022,1)` allowed if free;
- `(5020,2)` allowed if free;
- `(5020,1)` rejected when Simulator owns that exact pair.

Do not substitute direct socket/API apply calls for required rendered UI checks.

## JR TEST REPORT

Verdict: PENDING
