# Handoff

## Current

ACTIVE SEQUENCE:
- REP-BLOCK-001 — Explicit Source Pull Block Model
- REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

State: IMPLEMENTED — READY FOR JR VERIFICATION

The previous REP-UI sequence has been cleared from Active Work. Current implementation authority is limited to the two tasks above.

## Implemented Scope

### REP-BLOCK-001
- Replicator `DeviceDefinition` now contains one explicit `pull_block` with FC, Start, Count, and Scan Rate.
- Runtime source polling and destination range mapping are derived from that explicit block.
- Validation addresses the pull block directly.
- persisted legacy device YAML with device-level `function/start/count/scan_rate_ms` is migrated in memory and becomes canonical `pull_block` YAML on the next successful save.
- OS.js Replicator UI now renders a distinct **Pull Block** section.

### REP-OWN-001
Save & Apply now follows the required order:

1. resolve destination and read latest shared MMA2 ownership;
2. reject foreign `(port, unit_id)` ownership with owner/reservation in the error, before Replicator document persistence;
3. claim free/self-owned destinations while preserving all foreign MMA2 reservations and write shared MMA2 config/owners;
4. request MMA2 restart through restart-request/ack and wait for destination readiness;
5. persist the Replicator document and restart Replicator poll loops only after MMA2 activation succeeds.

Save & Apply now requests the MMA2 restart lifecycle on every successful apply, including source-only or scan-rate-only edits. Pre-save UI Owner/Status remains advisory; backend Save & Apply is the safety boundary.

No completion/archive claim is made until JR evidence is reviewed.

## JR TEST TASK — REP-BLOCK-001 + REP-OWN-001

JR role: TEST AND REPORT ONLY. Do not fix source, modify Active Work/planning/ICC, or expand scope. Runtime test configuration through Simulator/Replicator UI is authorized. Do not destroy Docker volumes.

JR may edit only the `## JR TEST REPORT` section and may commit/push only `handoff.md` after testing.

### 1. Sync and clean state

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
```

Expected: clean tree before testing. If unexpectedly dirty, report BLOCKED and stop.

### 2. Replicator formatting/tests/vet

```bash
cd replicator
gofmt -l .
go test -count=1 -run '^(TestDocumentSaveLoadRoundTrip|TestLoadDocumentMigratesLegacyRangeIntoPullBlock|TestSuggestDestinationSkipsOccupiedReservations|TestResolveManualForeignCollisionReportsOwner|TestSaveTimeOwnershipGuardRejectsLatestForeignOwner|TestComposeDocumentPreservesForeignAndAllReplicatorReservations|TestDeviceRuntimeConfigMapsPullBlockOneToOne|TestSaveApplyRejectsForeignReservationBeforeMutation|TestRuntimeManagerApplyLifecycleAndStatus)$' .
go test -count=1 ./...
go vet ./...
```

Expected:
- `gofmt -l .` prints nothing;
- all focused tests PASS;
- full Replicator test suite PASS;
- vet exits 0.

### 3. OS.js build

From repository root:

```bash
cd OSJS
npm run build:local-packages
npm run package:discover
npm run build
```

Expected: ModbusReplicator builds/discovers and full build exits 0.

### 4. Deploy changed services

```bash
cd ../deploy
docker compose build osjs-shell modbus-replicator-runtime mma2
docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime
docker compose ps
```

Expected: osjs-shell healthy; Simulator runtime, Replicator runtime, and MMA2 running.

### 5. Pull Block UI/runtime check

Use or create a Simulator source at:

```text
Port: 5020
Unit ID: 1
FC3 Start: 0
FC3 Count: 16
```

In Replicator create one enabled device with:

```text
Endpoint: 127.0.0.1:5020
Source Unit ID: 1
Pull Block: FC3 / Start 0 / Count 16 / Scan Rate 1000 ms
Destination: automatic
```

Expected:
- UI visibly has a **Pull Block** section;
- Save & Apply succeeds on a free destination;
- persisted `/data/config/replicator/devices.yaml` contains `pull_block:` and does not persist the old loose device-level FC/start/count/scan-rate representation;
- Replicator reaches RUNNING / Source OK.

Capture:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/replicator/devices.yaml
```

### 6. Authoritative ownership conflict check

Ensure Simulator owns `(5020,1)` and Replicator currently has a different valid destination.

In Replicator:
- disable Auto Port and Auto Unit ID;
- manually set destination Port `5020`, Unit ID `1`;
- click **Save & Apply**.

Required result:
- Save & Apply FAILS;
- error identifies destination `5020/1` and owner `simulator`;
- Simulator reservation remains owned by `simulator`;
- the previously persisted Replicator document remains unchanged;
- no foreign shared MMA2 entry is removed or overwritten.

Capture before/after as needed:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
docker exec mcs-modbus-replicator-runtime cat /data/config/replicator/devices.yaml
```

Do not accept a pre-save warning alone as proof. The required evidence is that clicking Save & Apply itself rejects the conflict.

### 7. Free destination claim + MMA2 restart check

Discard the rejected edit. Choose a genuinely free destination manually or use automatic allocation, then click Save & Apply.

Required result:
- destination is claimed by `replicator` in `owners.yaml`;
- shared MMA2 effective settings include the Replicator destination;
- MMA2 restart-request/ack path completes;
- Save & Apply reports success only after restart/readiness;
- Replicator returns RUNNING / Source OK.

Then make a **scan-rate-only** edit and Save & Apply again.

Required result:
- even though destination structure is unchanged, MMA2 restart lifecycle is requested/completed again;
- apply succeeds only after that restart path completes.

Capture relevant runtime/MMA2 logs and ownership evidence.

### 8. Final repository state

```bash
cd ..
git status --short
```

Expected: no unexpected tracked changes before the report edit.

## JR TEST REPORT

Verdict: PENDING
