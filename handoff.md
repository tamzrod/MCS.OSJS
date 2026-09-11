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

Verdict: PASS

### 1. Sync & backend gate — PASS
- `git pull --ff-only origin main`: fast-forward `969e305..5aaa0f6` (pair-ownership reversion(; `git status --short` clean; HEAD `5aaa0f62c35bf7d68a86112e8846d21ae1be8c99`. PASS.
- `gofmt -l .`: no output. (Go 1.22.12 sandbox-local `~/.local/go`, PATH exported.)
- Focused tests: `ok github.com/tamzrod/MCS.OSJS/replicator 0.008s` — PASS.
- `go test -count=1 ./...`: `ok ... 12.143s` (+ `? .../modbus-replicator-runtime [no test files]`( → `FULL_SUITE_OK` — PASS.
- `go vet ./...`: exit  ́0 → `VET_OK` — PASS.



###2. Deploy current build — PASS
- `docker compose build osjs-shell modbus-replicator-runtime mma2`: all three images Built (EXIT=0(.
- `docker compose up -d osjs-shell modbus-replicator-runtime mma2 modbus-simulator-runtime`: UP_EXIT=0; replicator-runtime recreated/started; osjs-shell, simulator-runtime, MMA2 remained up.

- `docker compose ps`: all four services up; `mcs-osjs-shell` Up (healthy(; simulator/replicator/MMA2 running. Replicator runtime log: `Replicator runtime listening on /data/run/modbus-replicator.sock` (stable, no crash-loop(. PASS..



###3. Prepare Simulator ownership — PASS
Simulator owned `(5020,1)` exactly (Sim-PLC-1, Port 5020 / Unit ID  ́1, FC3 0–16(; `owners.yaml` baseline: reservation `(5020,1)/simulator`. PASS..



###4. Same Unit ID on different Port — MUST SUCCEED — PASS
In the rendered Replicator UI (OS.js window, CDP-driven real DOM input/change events + real rendered button clicks(:
1-2. Rep-PLC-1 valid; disabled Auto Port and Auto Unit ID.

3. Set destination Port `5022`, Unit ID `1`.
4. Ownership inspection (rendered UI( reported `Owner: replicator / Status: AVAILABLE` — not Simulator-owned. PASS..
5. Clicked the rendered **Save & Apply** button:
- Save & Apply SUCCEEDED; UI status: `Replicator settings saved to shared MMA2 configuration and activated through the MMA2 restart path. Applied at 9/11/2026,  ́11:48:06 AM.; Owner→replicator / OWNED. PASS..
- `owners.yaml`: `(5020,1)/simulator` + `(5022,1)/replicator` — both present. PASS.
.
- MMA2 restart/readiness completed: log `11:48:06 restart request consumed → shutdown requested → v2.0.2 starting → ingress replicator-5022-1 listening on  ́0.0.0.0:5022`. PASS..
- Replicator runtime returned RUNNING / Source OK (status probe: `"running":true,"source_status":"OK"`(. PASS.



###5. Same Port with different Unit ID — MUST SUCCEED — PASS
Edited Replicator destination to Port `5020`, Unit ID `2` (same Port as Simulator, different Unit(:
- Ownership inspection (rendered UI( reported `Owner: replicator / Status: AVAILABLE`. PASS..
- Clicked the rendered **Save & Apply** button: SUCCEEDED (UI `Applied atath 11:48:33 AM.`); `owners.yaml`: Simulator `(5020,1)` untouched + Replicator now `(5020,2)`; MMA2 restart-request/ack completed (`11:48:33 restart request consumed … ingress sim-5020-1 listening` covering unit 2(; runtime RUNNING / Source OK (status probe `"running":true`(. PASS.



###6. Exact pair conflict — MUST FAIL — PASS
Edited Replicator destination to exact pair Port `5020`, Unit ID `1`:
- ownership inspection (rendered UI( showed `Owner: simulator / Status: IN USE`. PASS..
- Clicked the rendered **Save & Apply** button:
- Save & Apply FAILED; visible UI error: `Save & Apply failed: device "Rep-PLC-1": mma2 reservation owned by another producer: destination (5020,1) owned by "simulator"`. PASS..
- Error identifies exact destination `(5020,1)` and owner `simulator`. PASS..
- No MMA2 restart occurred for the rejected apply: MMA2 log shows no `restart request consumed` after the step-5 success at **11:48:33**; MMA2 stayed up (`Up` in compose ps;). PASS..
- Previously persisted valid Replicator destination remains unchanged: `devices.yaml` still `(5020,2)` (auto(false(false(. PASS..
- Simulator reservation remains unchanged: `(5020,1)/simulator` in owners.yaml. PASS..
All step  6 checks were performed with the rendered UI button — no direct socket/API apply substituted. PASS.



###7. Final state — PASS
- `/data/config/mma2/owners.yaml`: Simulator `(5020,1)` + Replicator `(5020,2)` — coherent pair-scoped ownership, both reservations intact.
.
- `/data/config/replicator/devices.yaml`: Rep-PLC-1 persisted at `(5020,2)` (auto(false(false(; canonical `pull_block:` (function 3, start0, count16(,scan_rate_ms1000(.
- `git status --short` (repo root(: empty — no unexpected tracked changes before report edit. PASS.



Caveat (environment, not product(:
- Sandbox lacks Go/Node; used prior local installs (`~/.local/go`, `~/.local/node-v16.20.2-linux-x64`) per JR sandbox authority; docker daemon sandbox-local; OSJS/`node_modules` remains from earlier builds (gitignored; tree clean(.
- This packet began with pair-scoped persisted state from the prior packet runs (Simulator `(5020,1)` + Replicator `(5021,1)`(. That is a legitimate valid pair-scoped baseline; no volume destruction or product change was needed; Step4+5 demonstrated transitions to `(5022,1)` theen `(5020,2)` and Step6 rejection preserved `(5020,2)` co. The exact-pair conflict evidence came entirely from rendered-UI button clicks and observations.,as the packet required.

Verdict: PASS
