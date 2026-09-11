# Handoff

## Current

ACTIVE SEQUENCE: REP-UI-001 through REP-UI-007
State: JR PASS WITH TWO UI GAPS RECTIFIED — READY FOR TARGETED RETEST

JR's latest consolidated retest passed the MMA2 compile/deployment fix and the Replicator end-to-end runtime path. Accepted evidence already covers:

- MMA2 build/tests/vet and stable supervised deployment;
- Simulator source creation and simulator ownership `(5020,1)`;
- Replicator automatic destination allocation to a separate reservation;
- Save & Apply, persistence, reopen, Discard, Duplicate;
- Replicator RUNNING / Source OK runtime status;
- truthful runtime ERROR on unreachable source and recovery to OK;
- foreign-owned destination rejection without corrupting persisted state;
- desktop and Start-menu launch;
- earlier Replicator focused/full tests, vet, Simulator-to-Replicator E2E, and OS.js build/discovery.

Two UI-level observations remained incomplete even though backend/unit evidence was safe:

1. deleting the final Replicator device removed the selected editor, so the browser UI no longer exposed a usable Save & Apply control to commit the empty document and release the last Replicator reservation;
2. the foreign-owner collision was rejected correctly, but the UI did not present a sufficiently explicit visible `IN USE` ownership warning in JR's browser snapshot.

## Rectification

`OSJS/src/packages/ModbusReplicator/index.js` now:

- keeps Save & Apply and Discard available after the final device is locally deleted;
- explicitly tells the operator that all devices are marked for deletion and Save & Apply will release Replicator-owned destinations;
- re-inspects destination ownership as soon as both automatic destination controls are disabled;
- renders an explicit warning when a manual destination is `IN USE`, including the actual owner and `(port, unit_id)`.

No backend ownership, allocation, runtime, MMA2 lifecycle, or persistence architecture changed.

All seven Active Work tasks remain active until this targeted JR retest is reviewed by the coding agent.

## JR TEST TASK — Targeted Replicator UI closure retest

JR role: TEST AND REPORT ONLY. Do not fix product failures, modify Active Work, planning, brainstorm, ICC, or product source. Runtime UI configuration changes are authorized. Do not destroy the Docker volume.

JR may modify only `handoff.md` below `## JR TEST REPORT`, then commit/push that report only.

### 1. Sync and clean repository state

From repository root:

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
```

Expected: clean tree before testing.

### 2. OS.js changed-surface build gate

```bash
cd OSJS
npm run build:local-packages
npm run package:discover
npm run build
```

Use the already prepared compatible Node environment from prior JR runs. Expected: all commands exit 0 and ModbusReplicator builds/discovers successfully.

Rebuild/restart only the OS.js shell so the updated package is deployed:

```bash
cd ../deploy
docker compose build osjs-shell
docker compose up -d osjs-shell
docker compose ps
```

Expected: osjs-shell healthy; simulator, replicator runtime, and MMA2 remain running.

### 3. Foreign ownership warning UI

Use/recreate the known valid setup:

- Simulator source owns `(5020,1)`;
- one persisted Replicator device owns a different destination such as `(5021,1)`.

In Modbus Replicator:

1. select the persisted Replicator device;
2. disable Auto Port and Auto Unit ID;
3. set destination Port `5020`, Unit ID `1`;
4. allow the ownership inspection to complete.

Expected UI before Save & Apply:

- Owner visibly reads `simulator`;
- Status visibly reads `IN USE`;
- a clear warning says destination `5020/1` is owned by `simulator` and cannot be claimed by Replicator.

Attempt Save & Apply.

Expected: apply remains rejected and persisted owners remain unchanged. Capture:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
```

Then click Discard to restore the valid persisted Replicator destination.

### 4. Final-device delete/release UI

With exactly one persisted Replicator device present:

1. select it;
2. click Delete;
3. verify the editor becomes an empty/pending-delete state but still shows Save & Apply and Discard;
4. verify the UI text says all devices are marked for deletion and Save & Apply will release Replicator-owned destinations;
5. click Save & Apply.

Expected:

- apply succeeds with an empty Replicator document;
- reopening/reloading Modbus Replicator shows no Replicator devices;
- the deleted Replicator reservation is gone from `owners.yaml`;
- simulator-owned `(5020,1)` remains untouched.

Capture:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/replicator/devices.yaml
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
```

Expected final ownership: simulator `(5020,1)` remains; no reservation for the deleted Replicator device remains.

### 5. Final repository state

From repository root:

```bash
git status --short
```

Expected: no unexpected tracked changes except `handoff.md` after JR writes the report.

## JR TEST REPORT

Pending targeted retest.
