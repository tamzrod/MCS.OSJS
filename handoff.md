# Handoff

## Current

ACTIVE SEQUENCE: REP-UI-001 through REP-UI-007
State: RECTIFIED AGAIN — READY FOR JR RETEST

JR's second retest found that the prior MMA2 idle-state correction accidentally replaced working startup API calls with non-existent/incompatible helpers (`config.BuildStore`, `authority.New(cfg)`, `notify.New`, `accessevents.New(cfg)`). The idle-wait idea was correct, but the edit changed unrelated startup code and therefore did not compile.

## Rectification

`MMA2/cmd/mma2/main.go` has now been restored to the previously compiling startup path used by commit `584fcbaa8762f948b066dd4d3c00aadd30a1cc96`:

- `config.BuildMemoryStore(cfg)`
- `authority.New()` plus `config.BuildAuthorityPolicies(cfg)`
- `config.BuildNotifyRegistry(cfg)` plus `notify.NewEngine(...)`
- `accessevents.New(cfg.AccessEvents)` plus `accessevents.NewHandler(ae)`

The only intended behavior change retained is the terminal idle wait:

- old: bare `select {}`
- new: wait for SIGINT/SIGTERM

This keeps empty valid MMA2 configurations alive without changing the established runtime APIs.

Passing evidence from earlier runs remains accepted: Replicator tests/vet, Simulator-to-Replicator E2E, OS.js build/discovery, package visibility, Start-menu launch, and Replicator editor layout.

All seven Active Work tasks remain active. Do not archive or advance them until the coding agent reviews this retest.

## JR TEST TASK — MMA2 compile/deployment and blocked Replicator runtime retest

JR role: TEST AND REPORT ONLY. Do not fix product failures or edit product source. Runtime UI configuration changes are allowed. Do not destroy the persistent Docker volume.

JR may modify only `handoff.md` for the report and may commit/push that report only.

### 1. Sync and clean state

From repository root:

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
```

Expected: clean tree.

### 2. MMA2 compile gate

```bash
cd MMA2
gofmt -l cmd/mma2/main.go
go test -count=1 ./...
go vet ./...
```

Expected:

- gofmt prints nothing;
- all MMA2 packages build/test successfully;
- vet exits 0.

### 3. MMA2 deployment gate

From repository root:

```bash
cd deploy
docker compose build mma2
docker compose up -d mma2
docker compose ps
docker compose logs --no-color --tail=120 mma2
```

Expected:

- image builds;
- `mcs-mma2` stays Up, not Restarting;
- empty valid config may log `mma2 ingress started` and remain alive;
- SIGTERM/restart behavior remains supervisor-compatible.

Confirm all services:

```bash
docker compose ps
```

Expected: osjs-shell healthy; modbus-simulator-runtime, modbus-replicator-runtime, mma2 running.

### 4. Previously blocked end-to-end checks

Use the deployed OS.js desktop.

#### Simulator source

Create/reuse:

```text
Name: JR-Rep-Source
Port: 5020
Unit ID: 1
FC3 Start: 0
FC3 Count: 4
FC3 random interval: 60000 ms
```

Save & Apply. Expected: Simulator and MMA2 running; `(5020,1)` owned by `simulator`.

#### Replicator auto allocation + persistence

Create:

```text
Name: JR-Rep-Device
Enabled: checked
Endpoint: 127.0.0.1:5020
Source Unit ID: 1
FC: FC3
Start: 0
Count: 4
Scan Rate: 1000 ms
Destination: Auto Port + Auto Unit ID enabled
```

Verify:

- automatic destination avoids `(5020,1)`;
- Save & Apply succeeds;
- Owner `replicator`, Status `OWNED`;
- Replicator RUNNING, Source OK, Last Poll advances;
- close/reopen preserves resolved destination;
- Discard restores an unsaved edit;
- Duplicate receives a distinct automatic destination, then delete the unsaved duplicate.

#### Foreign ownership guard

Set manual destination to `5020 / 1` with auto disabled.

Expected:

- Owner `simulator`, Status `IN USE`;
- Save & Apply rejected;
- simulator ownership preserved.

Capture:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
```

Then Discard.

#### Runtime error and recovery

Set source Endpoint to `127.0.0.1:1`, Save & Apply.

Expected: Replicator stays RUNNING; Source ERROR; Last Error shows connection failure; polling continues.

Restore `127.0.0.1:5020`, Save & Apply. Expected: Source returns OK and Last Poll advances.

#### Delete/release

Delete `JR-Rep-Device`, Save & Apply.

Expected: its Replicator reservation disappears while simulator `(5020,1)` remains.

Capture final owners file again.

### 5. Desktop launcher

If reliable pointer interaction is available, double-click the Modbus Replicator desktop shortcut.

Expected: launches Modbus Replicator. If automation cannot reliably double-click, report INCONCLUSIVE for this single item only.

### 6. Final git state

From repository root:

```bash
git status --short
```

Expected: empty except handoff.md after writing the report.

## JR TEST REPORT

Verdict: PENDING
Tested commit: PENDING

Repository state:
PENDING

MMA2 compile gate:
PENDING

MMA2 deployment gate:
PENDING

Simulator source:
PENDING

Auto allocation + persistence:
PENDING

Ownership collision:
PENDING

Runtime status/error recovery:
PENDING

Delete/release ownership:
PENDING

Desktop launcher:
PENDING

Final git status:
PENDING

Unexpected behavior:
PENDING
