# Handoff

## Current

ACTIVE SEQUENCE:
- REP-BLOCK-001 — Explicit Source Pull Block Model
- REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

State: HUMAN CORRECTION APPLIED — READY FOR PAIR-OWNERSHIP RETEST

Human testing corrected the ownership semantics:

- Simulator owns `(5020,1)`.
- Replicator using `(5022,1)` is valid because the port is different.
- Replicator using `(5020,2)` is also valid because the Unit ID is different.
- Only the exact `(port, unit_id)` pair is an ownership collision.

The previous rectification incorrectly made Port and Unit ID independently exclusive. That change has been reverted.

## Correct Ownership Rule

Shared MMA2 ownership key is exactly:

```text
(port, unit_id)
```

Therefore:

- `(5020,1)` owned by Simulator blocks only `(5020,1)`;
- `(5022,1)` is free unless that exact pair is reserved;
- `(5020,2)` is free unless that exact pair is reserved;
- automatic allocation skips occupied pairs, not globally reused ports or Unit IDs;
- Save & Apply re-reads latest ownership and rejects only an exact foreign pair before shared MMA2 mutation/restart/persistence.

No completion/archive claim is made until this corrected rule is verified.

## JR TEST TASK — Pair-scoped MMA2 ownership retest

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
go test -count=1 -run '^(TestSuggestDestinationSkipsOccupiedPairOnly|TestResolveManualForeignCollisionReportsOwner|TestResolveAllowsSameUnitOnDifferentPort|TestResolveAllowsSamePortWithDifferentUnit|TestSaveTimeOwnershipGuardRejectsExactForeignPairOnly|TestSaveApplyRejectsForeignReservationBeforeMutation)$' .
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

Expected: all services running and osjs-shell healthy.

### 3. Prepare Simulator ownership

Using the rendered Simulator UI, ensure Simulator owns:

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

### 4. Same Unit ID on different Port — MUST SUCCEED

In the rendered Replicator UI:

1. configure a valid Replicator device;
2. disable Auto Port and Auto Unit ID;
3. set destination Port `5022`, Unit ID `1`;
4. wait for ownership inspection;
5. click the rendered **Save & Apply** button.

Required result:
- ownership inspection reports AVAILABLE, not Simulator-owned;
- Save & Apply succeeds;
- owners contain both `(5020,1)/simulator` and `(5022,1)/replicator`;
- MMA2 restart/readiness completes;
- Replicator returns RUNNING / Source OK.

### 5. Same Port with different Unit ID — MUST SUCCEED

Edit the Replicator destination to:

```text
Port: 5020
Unit ID: 2
```

Click the rendered **Save & Apply** button.

Required result:
- ownership inspection reports AVAILABLE;
- Save & Apply succeeds;
- Simulator `(5020,1)` remains untouched;
- Replicator owns `(5020,2)`;
- MMA2 restart/readiness completes.

### 6. Exact pair conflict — MUST FAIL

Edit the Replicator destination to:

```text
Port: 5020
Unit ID: 1
```

Click the rendered **Save & Apply** button.

Required result:
- ownership inspection shows owner `simulator` / `IN USE`;
- Save & Apply fails;
- visible error identifies exact destination `(5020,1)` and owner `simulator`;
- no MMA2 restart occurs for the rejected apply;
- previously persisted valid Replicator destination remains unchanged;
- Simulator reservation remains unchanged.

Do not substitute direct socket/API apply calls for these UI-button checks. If the rendered button cannot be reliably clicked/observed, report BLOCKED.

### 7. Final state

Capture:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
docker exec mcs-modbus-replicator-runtime cat /data/config/replicator/devices.yaml
git status --short
```

Expected: coherent pair-scoped ownership and no unexpected tracked changes before report edit.

## JR TEST REPORT

Pending corrected pair-ownership retest.
