# Handoff

## Current

ACTIVE SEQUENCE:
- REP-BLOCK-001 — Explicit Source Pull Block Model
- REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

State: HUMAN OWNERSHIP FAILURE RECTIFIED — READY FOR TARGETED RETEST

Human testing found the remaining ownership defect directly in the UI:

- Simulator already owned a destination Unit ID;
- Replicator automatic allocation reused that same Unit ID on another port;
- Save & Apply reported success and restarted MMA2.

That behavior exposed an incorrect ownership assumption in the implementation: it treated only the exact `(port, unit_id)` pair as exclusive. The required workflow treats **Port and Unit ID as independently exclusive shared MMA2 resources**.

## Rectification

REP-OWN-001 and the implementation now enforce:

- a foreign-owned Port cannot be reused by Replicator even with a different Unit ID;
- a foreign-owned Unit ID cannot be reused by Replicator even on a different Port;
- automatic allocation skips every occupied Port and every occupied Unit ID;
- pre-save inspection reports `IN USE` when either requested resource belongs to another producer;
- Save & Apply re-reads latest ownership and rejects either foreign Port or foreign Unit ID before shared MMA2 mutation, restart, or Replicator document persistence;
- conflicts identify the actual foreign owner and conflicting resource;
- Replicator devices in one edited document also receive distinct destination Ports and Unit IDs.

Regression coverage now includes:
- same pair conflict;
- same Port / different Unit ID conflict;
- different Port / same Unit ID conflict (the human-found failure);
- automatic allocator skipping occupied Ports and Unit IDs;
- rejected Save & Apply preserving both foreign ownership and the previously persisted Replicator document.

No completion/archive claim is made until this rectification is verified.

## JR TEST TASK — Human ownership defect retest

JR role: TEST AND REPORT ONLY. Do not fix source, modify Active Work/planning/ICC, or expand scope. Do not destroy Docker volumes.

JR may edit only the `## JR TEST REPORT` section and may commit/push only `handoff.md` after testing.

### 1. Sync and focused backend gate

From repository root:

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
cd replicator
gofmt -l .
go test -count=1 -run '^(TestSuggestDestinationSkipsOccupiedPortsAndUnits|TestResolveManualForeignCollisionReportsOwner|TestResolveRejectsForeignPortWithDifferentUnit|TestResolveRejectsForeignUnitWithDifferentPort|TestSaveTimeOwnershipGuardRejectsLatestForeignOwner|TestSaveApplyRejectsForeignReservationBeforeMutation)$' .
go test -count=1 ./...
go vet ./...
```

Expected:
- clean tree before testing;
- `gofmt -l .` prints nothing;
- focused tests PASS;
- full Replicator tests PASS;
- vet exits 0.

### 2. Deploy current build

```bash
cd ../deploy
docker compose build osjs-shell modbus-replicator-runtime mma2
docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime
docker compose ps
```

Expected: all four services running; osjs-shell healthy.

### 3. Prepare Simulator ownership

Using the rendered Simulator UI, create/save an enabled source owning:

```text
Port: 5020
Unit ID: 1
FC3 Start: 0
FC3 Count: 16
```

Capture:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
```

Required baseline: Simulator owns port `5020` and Unit ID `1`.

### 4. Automatic allocation — required UI check

Open Modbus Replicator and add a new enabled device with automatic destination allocation.

Required result before Save & Apply:
- automatically selected destination Port is **not 5020**;
- automatically selected destination Unit ID is **not 1**;
- allocation does not reuse either foreign-owned resource.

Then click the rendered **Save & Apply** button.

Required result:
- apply succeeds only on a Port and Unit ID both free from Simulator ownership;
- `owners.yaml` contains distinct Simulator and Replicator Port values and distinct Unit ID values;
- Simulator reservation remains unchanged.

### 5. Same Unit ID / different Port — human-regression UI check

With Simulator still owning `5020 / Unit 1` and Replicator persisted on another valid destination:

1. disable Auto Port and Auto Unit ID;
2. choose a different otherwise-free Port, for example `5022`;
3. manually set Unit ID `1`;
4. wait for ownership inspection;
5. click the rendered **Save & Apply** button.

Required result:
- UI shows `IN USE` / owner `simulator` for the Unit ID conflict;
- actual Save & Apply fails;
- visible error identifies Unit ID `1` and owner `simulator`;
- MMA2 is **not** restarted for the rejected apply;
- persisted Replicator document remains on its previous valid destination;
- Simulator ownership remains unchanged.

Do not substitute a direct socket/API apply for this UI-button test. If the rendered button cannot be clicked/observed reliably, report BLOCKED.

### 6. Same Port / different Unit ID — required UI check

Discard the rejected edit. Then:

1. set manual Port `5020`;
2. set a different otherwise-free Unit ID, for example `2` or another currently unused value;
3. click the rendered **Save & Apply** button.

Required result:
- UI shows `IN USE` / owner `simulator` for the Port conflict;
- actual Save & Apply fails;
- visible error identifies port `5020` and owner `simulator`;
- no MMA2 restart occurs for the rejected apply;
- prior Replicator and Simulator state remains unchanged.

### 7. Final state

Capture:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
docker exec mcs-modbus-replicator-runtime cat /data/config/replicator/devices.yaml
git status --short
```

Expected: ownership remains coherent and no unexpected tracked changes exist before report edit.

## JR TEST REPORT

Pending targeted retest of the human-found ownership defect.
