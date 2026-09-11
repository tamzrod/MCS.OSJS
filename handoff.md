# Handoff

## Current

ACTIVE SEQUENCE:
- REP-BLOCK-002 — Independent Pull Block Pollers
- REP-BLOCK-003 — Device / Pull Blocks Folder Tabs

State: IMPLEMENTED — READY FOR JR VERIFICATION

REP-BLOCK-001 and REP-OWN-001 were verified by the prior rendered-UI pair-ownership packet and archived. The authoritative MMA2 ownership rule remains the exact `(port, unit_id)` pair.

## Implemented behavior

### REP-BLOCK-002
- One Replicator device now carries an ordered `pull_blocks` collection.
- Legacy flat source fields and the prior single `pull_block` shape migrate to one collection entry on load.
- Each enabled Pull Block receives its own runtime poll loop/ticker and scan rate.
- Source Endpoint, Source Unit ID, Enabled, Name, and destination reservation remain device-level.
- Per-block runtime status exposes index, running, cycles, source status, last poll, and last error.
- MMA2 destination memory spans all configured FC3/FC4 block ranges while each poller writes only its own configured range.
- Shared destination ownership remains one exact `(port, unit_id)` reservation per device.

### REP-BLOCK-003
- Existing left device tree/list remains the device selector and keeps Add/Duplicate/Delete behavior.
- Right-side selected-device configuration is separated into exactly two classic folder-style tabs:
  - `Device`
  - `Pull Blocks`
- `Device` contains Name, Enabled, Endpoint, Source Unit ID, Destination, Owner, and ownership Status.
- `Pull Blocks` contains block cards/rows; blocks are not another tab layer.
- Add Block / Duplicate Block / Delete Block operate only on the Pull Block collection.
- Save & Apply / Discard remain device-level actions.
- Per-block runtime status is rendered in the Pull Blocks tab.

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
- legacy single-block migration to one `pull_blocks` entry;
- multi-block persistence/order;
- independent per-block cadence/status;
- exact pair ownership regressions still pass.

### 2. Deploy current build

```bash
cd ../deploy
docker compose build osjs-shell modbus-replicator-runtime mma2
docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime
docker compose ps
```

Expected: all services running and osjs-shell healthy.

### 3. Verify unchanged left device selector

Open the rendered Modbus Replicator UI.

Required:
- left device list/search remains on the left;
- Add / Duplicate / Delete device controls remain there and still operate on devices;
- selecting a device changes only the right-side editor.

### 4. Verify the two right-side folder tabs

For one selected device, confirm the right side has exactly two top-level folder-style tabs:

```text
Device | Pull Blocks
```

Required:
- active tab visually joins the content panel like a classic file folder;
- no per-block tab strip exists;
- `Device` contains only device-level source/destination settings;
- `Pull Blocks` contains the block collection.

### 5. Multi-block Save & Apply

Use a Simulator source and configure one Replicator device with at least two Pull Blocks, for example:

```text
Block 1: FC3 Start 0   Count 8  Scan 100 ms
Block 2: FC4 Start 100 Count 4  Scan 500 ms
```

Click the rendered **Save & Apply** button.

Required:
- Save & Apply succeeds through the existing MMA2 restart/readiness path;
- `devices.yaml` persists both blocks in `pull_blocks:` in the same order;
- one destination `(port, unit_id)` reservation exists for the device;
- per-block status is independently visible;
- the faster block accumulates cycles faster than the slower block when source conditions allow observation.

### 6. Block editing behavior

In the rendered Pull Blocks tab:
- Add Block;
- select a block card;
- Duplicate Block;
- edit FC/start/count/scan rate independently;
- Delete one block;
- Save & Apply;
- reopen/reload the UI.

Required:
- only block collection changes;
- device-level Endpoint/Unit/Destination remain unchanged;
- remaining blocks retain their values/order;
- deleting one block removes/stops only that poller after apply;
- at least one block must remain.

### 7. Ownership regression

With Simulator owning `(5020,1)`, confirm again:
- Replicator `(5022,1)` succeeds when free;
- Replicator `(5020,2)` succeeds when free;
- Replicator `(5020,1)` rejects as exact-pair conflict.

Do not substitute direct socket/API apply calls for required rendered UI button checks. If the rendered UI cannot be reliably observed, report BLOCKED.

## JR TEST REPORT

Verdict: PENDING
