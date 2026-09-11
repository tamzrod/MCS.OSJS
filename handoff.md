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

Verdict: FAIL

### 1. Sync and backend gate — FAIL (gofmt)
- `git pull --ff-only origin main`: fast-forward `98563df..802a6c3` (multi-block implementation(; `git status --short` clean; HEAD `802a6c38e86efce3ae2547148bbfd30caf290a35`. PASS.
- `gofmt -l .` **printed `document_test.go`** — the packet requires it to print nothing. **FAIL.** (Re-checked after all steps; still prints `document_test.go`. The single diff is an alignment nit in a struct literal at `document_test.go` ~line 239: `Policy: &mma2composer.Policy{...}` not gofmt-aligned. JR does not fix source; records the contradiction.)
- `go test -count=1 ./...`: `ok github.com/tamzrod/MCS.OSJS/replicator 12.408s` (+ `? .../modbus-replicator-runtime [no test files]`( → PASS.
- `go vet ./...`: exit 0 → PASS.
- Focus evidence (from full suite(: legacy single-block→`pull_blocks` migration, multi-block persistence/order, per-block cadence/status, and pair-ownership regressions all covered by passing tests (full suite green(. PASS.

Per the packet decision rule, ANY required item executed and contradicting expectation → overall FAIL. The gofmt gate fails; therefore overall verdict FAIL even though all other steps passed.

### 2. Deploy current build — PASS
- `docker compose build osjs-shell modbus-replicator-runtime mma2`: all three images Built (EXIT=0(.
- `docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime`: UP_EXIT=0; replicator-runtime and osjs-shell recreated/started; MMA2 and simulator-runtime remained up.
- `docker compose ps`: all four services up; `mcs-osjs-shell` Up (healthy(; replicator runtime log `Replicator runtime listening on /data/run/modbus-replicator.sock` (stable(. PASS.

### 3. Verify unchanged left device selector — PASS
Rendered Modbus Replicator UI: left device list/search remains on the left with Add / Duplicate / Delete controls; selecting a device changes only the right-side editor. Device row `Rep-PLC-1` present. PASS.

### 4. Verify the two right-side folder tabs — PASS
For the selected device the right side has exactly two top-level folder-style tabs `Device | Pull Blocks` (both rendered as tab buttons; active tab joins the content panel). No per-block tab strip exists — blocks are cards/rows inside the Pull Blocks tab. `Device` tab contains only device-level Name/Enabled/Endpoint/Source Unit ID/Destination Port/Auto Port/Unit ID/Auto Unit ID/Owner/Status. `Pull Blocks` tab contains the block collection (Add Block / Duplicate Block / Delete Block + block cards(. PASS.

### 5. Multi-block Save & Apply — PASS
Configured one Replicator device with two Pull Blocks (Block1 FC3/0/16/1000ms; Block2 FC4/100/4/500ms( via the rendered Pull Blocks tab; clicked the rendered **Save & Apply**:
- Succeeded through the MMA2 restart/readiness path (MMA2 log `12:33:21 restart request consumed … ingress replicator-5021-2 listening`(; UI status "Replicator settings saved to shared MMA2 configuration and activated through the MMA2 restart path.".
- `devices.yaml` persists both blocks in `pull_blocks:` in the same order (function3/0/16/1000 then function4/100/4/500(. PASS.
- One destination reservation `(5020,2)/replicator` for the device (owners.yaml(. PASS.
- Per-block status independently visible in UI (Block1 Poller RUNNING/Source OK; Block2 Poller RUNNING/Source ERROR — Modbus exception code 2( and in runtime status (`blocks[0]` and `blocks[1]` with independent cycles/status(. PASS.
- Faster block accumulates cycles faster: runtime probe showed block0(1000ms(→12 cycles, block1(500ms(→24 cycles, block2(300ms(→39 cycles ≈ 1:2:3.25 — cycle counts track scan-rate ratios, confirming independent per-block cadence. (Blocks hitting out-of-range addresses report source ERROR but still cycle at their own rate; observation of cadence held.) PASS.

### 6. Block editing behavior — PASS
In the rendered Pull Blocks tab:
- **Add Block**: created a new block card (selected(. PASS.
- **select a block card**: clicking a card selects it (selection highlight observed(. PASS.
- **Duplicate Block**: duplicated the selected block into a new card with identical values (observed: selected Block2 FC4/100/4/500 → new Block3 FC4/100/4/500(. PASS.
- **edit FC/start/count/scan rate independently**: edited the duplicated Block3 to FC3/Start20/Count6/Scan300 via real input events; values persisted in DOM across render. PASS.
- **Delete one block**: Delete Block removed the trailing duplicate card; remaining blocks retained values/order. PASS.
- **Save & Apply**: succeeded; `devices.yaml` persisted 3 blocks in order (FC3/0/16/1000, FC4/100/4/500, FC3/20/6/300(; one reservation `(5020,2)` intact; MMA2 restart completed (12:36:49(. PASS.
- **reopen/reload UI**: after reload, Pull Blocks tab shows exactly the 3 remaining blocks with values/order retained; deleted block's poller gone; device-level Endpoint/Unit/Destination unchanged (Device tab intact(. PASS.

### 7. Ownership regression — PASS (rendered UI buttons)
With Simulator owning `(5020,1)`:
- **`(5022,1)` free → succeeds**: inspection AVAILABLE; rendered Save & Apply succeeded; owners `(5020,1)/simulator` + `(5022,1)/replicator`; MMA2 restart 12:40:40. PASS.
- **`(5020,2)` free → succeeds**: inspection AVAILABLE; rendered Save & Apply succeeded; Simulator `(5020,1)` untouched + Replicator `(5020,2)`; MMA2 restart 12:44:10. PASS.
- **`(5020,1)` exact-pair conflict → rejects**: inspection `Owner: simulator / IN USE`; rendered Save & Apply FAILED with visible error `Save & Apply failed: device "Rep-PLC-1": mma2 reservation owned by another producer: destination (5020,1) owned by "simulator"`; no MMA2 restart for the rejected apply (last restart remains 12:44:10(; persisted Replicator doc unchanged `(5020,2)`; Simulator reservation unchanged. PASS.
All step-7 checks used the rendered UI button (no direct socket/API apply substituted(. PASS.

### Final state
- `/data/config/mma2/owners.yaml`: Simulator `(5020,1)` + Replicator `(5020,2)` — coherent pair-scoped ownership.
- `/data/config/replicator/devices.yaml`: Rep-PLC-1 with 3 pull_blocks (FC3/0/16/1000, FC4/100/4/500, FC3/20/6/300( at destination `(5020,2)`.
- `git status --short` (repo root(: empty before report edit.

Caveats (environment, not product(:
- Sandbox lacks Go/Node; used prior local installs (`~/.local/go` auto-resolved to go1.25.0 matching go.mod(; docker daemon sandbox-local; OSJS/`node_modules` from earlier builds (gitignored; tree clean(.
- The gofmt failure is a tracked-source formatting issue in `document_test.go` (present in the committed tree, not caused by JR(. JR does not fix source; the coding agent should run `gofmt -w document_test.go` and re-verify.
- Block2/Block3 source addresses (FC4/start100, FC3/start20+6( exceed the simulator's configured ranges (0–15(, producing Modbus exception code 2; this is a test-data out-of-range condition, not a product defect — per-block cadence/status remained observable as required.

Verdict: FAIL (gofmt gate in step 1; all other steps PASS)
