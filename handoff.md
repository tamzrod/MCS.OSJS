# Handoff

## Current

ACTIVE SEQUENCE:
- REP-BLOCK-001 — Explicit Source Pull Block Model
- REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

State: IMPLEMENTED — JR REPORT REJECTED FOR EVIDENCE SUBSTITUTION — READY FOR TARGETED UI RETEST

The implementation remains unchanged. JR's latest report is not accepted as a complete PASS because required Save & Apply UI actions were replaced with direct runtime API/socket `apply` requests when browser targeting became difficult. That substitution is explicitly prohibited by `operation cwal.md` and by the test packet's requirement to click **Save & Apply** itself.

Accepted evidence from the latest run does not need to be repeated:
- clean repository state;
- Replicator gofmt/focused/full tests/vet;
- OS.js builds/discovery;
- deployment health;
- Pull Block UI presence;
- canonical `pull_block:` persistence;
- RUNNING / Source OK behavior;
- backend ownership rejection behavior;
- ownership/config preservation after rejected backend apply;
- backend free-destination apply and MMA2 restart lifecycle.

Still unverified at the required surface:
1. clicking the actual Replicator **Save & Apply** button on a foreign-owned `(5020,1)` must visibly fail with owner `simulator`;
2. clicking the actual **Save & Apply** button on a free destination must succeed only after MMA2 restart/readiness;
3. clicking the actual **Save & Apply** button after a scan-rate-only edit must again drive the MMA2 restart lifecycle and return success only after readiness.

No completion/archive claim is made until those UI-button paths are directly verified.

## JR TEST TASK — Direct Save & Apply UI closure retest

JR role: TEST AND REPORT ONLY. Do not fix source, modify Active Work/planning/ICC, or expand scope. Do not destroy Docker volumes.

JR may edit only the `## JR TEST REPORT` section and may commit/push only `handoff.md` after testing.

### Hard evidence rule for this retest

The required action is an actual user-surface click on the rendered **Save & Apply** button in Modbus Replicator.

The following are NOT valid substitutes:
- direct Unix socket calls;
- direct runtime API requests;
- calling `server.js`/`callRuntime` manually;
- backend-only apply commands;
- unit tests;
- source inspection;
- inferring that the button would call the same code path.

If browser automation cannot reliably click the rendered button or observe its result, report `BLOCKED`. Do not replace the UI test with another surface and do not report PASS.

### 1. Sync and deploy current main

From repository root:

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
```

Expected: clean tree before testing.

Ensure current services are up. Rebuild OS.js/Replicator only if the sandbox is not already on current main:

```bash
cd deploy
docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime
docker compose ps
```

Expected: osjs-shell healthy; Simulator runtime, Replicator runtime, and MMA2 running.

### 2. Prepare known ownership state

Using the running Simulator UI, ensure an enabled Simulator device owns:

```text
Port: 5020
Unit ID: 1
FC3 Start: 0
FC3 Count: 16
```

Using the Replicator UI, ensure one enabled Replicator device is persisted on a different free destination and reaches RUNNING / Source OK.

Capture baseline:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
docker exec mcs-modbus-replicator-runtime cat /data/config/replicator/devices.yaml
```

### 3. Required UI conflict click

In the rendered Modbus Replicator window:

1. select the persisted Replicator device;
2. disable Auto Port and Auto Unit ID;
3. set destination Port `5020`, Unit ID `1`;
4. wait for ownership inspection to show `simulator / IN USE`;
5. **click the rendered Save & Apply button**.

Required result from the UI click:
- Save & Apply visibly fails;
- visible error identifies destination `5020/1` and owner `simulator`;
- the previous persisted Replicator document remains unchanged;
- Simulator `(5020,1)` remains owned by `simulator`;
- no foreign MMA2 entry is removed or overwritten.

Capture after the UI click:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
docker exec mcs-modbus-replicator-runtime cat /data/config/replicator/devices.yaml
```

If the button cannot be reliably clicked/observed, verdict is BLOCKED.

### 4. Required UI free-destination click

Click Discard in the Replicator UI. Choose a genuinely free destination manually or restore automatic allocation.

Then **click the rendered Save & Apply button**.

Required result:
- UI reports success;
- destination becomes owned by `replicator` in `owners.yaml`;
- shared MMA2 configuration contains the destination;
- MMA2 restart-request/ack/readiness cycle occurs;
- Replicator returns RUNNING / Source OK;
- success appears only after MMA2 activation/readiness.

Capture relevant MMA2 logs and ownership state.

### 5. Required UI scan-rate-only click

In the same rendered Replicator UI, change only Pull Block Scan Rate.

Then **click the rendered Save & Apply button** again.

Required result:
- UI reports success;
- MMA2 restart-request/ack/readiness cycle occurs again despite unchanged destination structure;
- Replicator returns RUNNING / Source OK after apply.

Capture MMA2 logs showing the second restart lifecycle.

### 6. Final repository state

From repository root:

```bash
git status --short
```

Expected: no unexpected tracked changes before report edit.

## JR TEST REPORT

Pending targeted direct-UI retest.
