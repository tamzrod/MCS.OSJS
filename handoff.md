# Handoff

## Current

ACTIVE SEQUENCE: REP-UI-001 through REP-UI-007
State: IMPLEMENTED — READY FOR JR VERIFICATION

The human explicitly authorized the coding agent to implement the complete promoted Replicator UI sequence before JR verification. All seven Active Work tasks remain active until the coding agent reviews JR evidence; JR must not archive or advance them.

## Implemented Scope

The implementation now covers the complete promoted sequence:

1. REP-UI-001 — Modbus Replicator OS.js package, Start/Application-menu registration through package discovery, desktop shortcut, and the approved two-pane device-list/editor shell.
2. REP-UI-002 — source editor fields: Name, Enabled, Endpoint, Unit ID, FC3/FC4, Start, Count, and device-wide Scan Rate.
3. REP-UI-003 — ownership-aware automatic destination allocation with manual Port/Unit-ID override.
4. REP-UI-004 — first-come `(port, unit_id)` ownership guard; Replicator rebuilds only Replicator-owned reservations and reports foreign owners truthfully.
5. REP-UI-005 — persisted multi-device Replicator document, Save & Apply, validation, reload, Discard, duplicate, and delete/release behavior.
6. REP-UI-006 — independently managed Replicator runtime service, shared MMA2 composition, structural restart-request/ack/readiness lifecycle, and no new config API or bridge architecture.
7. REP-UI-007 — per-device runtime state surfaced to the UI: Replicator RUNNING/STOPPED, Source OK/ERROR/WAITING, Last Poll, and Last Error.

The legacy single-range Replicator API remains in place for its existing tests. The UI/runtime path uses a separate document manager so multiple device poll loops do not rebuild MMA2 ownership on every poll cycle.

No completion/archival claim is made here. Verification is pending JR.

## JR TEST TASK — REP-UI-001 through REP-UI-007

JR role: TEST AND REPORT ONLY.

Do not edit product source, manifests, Docker files, Active Work, planning, brainstorm, or ICC. Do not fix failures. Sandbox/user-local test tooling is allowed under `operation cwal.md`. Runtime test configuration changes made through the running Simulator/Replicator UIs are explicitly authorized for this verification; do not delete the persistent Docker volume and do not use `docker compose down -v`.

JR may modify **only `handoff.md`** to replace the `JR TEST REPORT` section after testing, then commit/push that report only.

### 1. Sync and clean repository state

From repository root:

```bash
git pull --ff-only origin main
git status --short
git rev-parse HEAD
```

Expected:

- pull succeeds;
- `git status --short` is empty;
- record the tested commit before writing the JR report.

If the tree is unexpectedly dirty, do not clean/reset/restore it. Record `BLOCKED` with exact status output and stop.

### 2. Replicator formatting and focused new tests

```bash
cd replicator
gofmt -l .
go test -count=1 -run '^(TestDocumentSaveLoadRoundTrip|TestSuggestDestinationSkipsOccupiedReservations|TestResolveManualForeignCollisionReportsOwner|TestComposeDocumentPreservesForeignAndAllReplicatorReservations|TestDeviceRuntimeConfigMapsOneToOneRange|TestRuntimeManagerApplyLifecycleAndStatus)$' .
```

Expected:

- `gofmt -l .` prints nothing;
- focused tests exit 0;
- the tests prove document persistence, automatic allocation, truthful foreign-owner collision, preservation of foreign reservations, Replicator-only rebuild/delete behavior, 1:1 source/destination mapping, structural restart contract, non-structural no-restart behavior, and truthful runtime error status.

If `gofmt -l .` lists files, record FAIL and continue only far enough to capture the remaining requested evidence; do not run `gofmt -w`.

### 3. Existing Simulator → Replicator end-to-end regression

Still under `replicator/`:

```bash
go test -count=1 -run '^TestSimulatorToReplicatorE2E$' .
```

Expected: exit code 0 and PASS. This must continue proving the real Simulator-owned source reservation, distinct Replicator-owned destination, real MMA2 runtime, Raw Ingest destination write, normal Modbus destination read, propagated source change, and distinct ownership.

### 4. Full Replicator regression and vet

```bash
go test -count=1 ./...
go vet ./...
```

Expected:

- all Replicator tests pass;
- `go vet ./...` exits 0 with no diagnostics.

### 5. OS.js package/build gates

From repository root:

```bash
cd OSJS
npm ci
npm run build:local-packages
npm run package:discover
npm run build
```

Use a sandbox/user-local Node version compatible with the repository engine if necessary. Do not modify package manifests or lockfiles to make the build pass.

Expected:

- dependency install succeeds without tracked-file mutation;
- local package builder includes `src/packages/ModbusReplicator` and succeeds;
- package discovery succeeds and discovers `ModbusReplicator` as an application;
- full OS.js build exits 0.

After the build:

```bash
cd ..
git status --short
```

Expected: no unexpected tracked changes. Ignored/untracked build output generated by the repository-native build is acceptable only when it is normal build output; report it if visible.

### 6. Docker/deployment build and service startup

From repository root:

```bash
cd deploy
docker compose config
docker compose build osjs-shell modbus-simulator-runtime modbus-replicator-runtime mma2
docker compose up -d
docker compose ps
```

Expected:

- compose config validates;
- all four images build;
- `osjs-shell`, `modbus-simulator-runtime`, `modbus-replicator-runtime`, and `mma2` are running;
- `osjs-shell` becomes healthy;
- do not destroy the `osjs-data` volume.

Capture these logs after startup:

```bash
docker compose logs --no-color --tail=120 modbus-replicator-runtime mma2
```

Expected: Replicator runtime starts without a fatal restore/socket error; MMA2 starts under its supervisor.

### 7. Launcher and shell UI verification

Open the deployed OS.js desktop at its normal local endpoint (`http://127.0.0.1:18209` unless the test environment intentionally overrides `OSJS_PORT`).

Verify all of the following:

- a **Modbus Replicator** icon/shortcut is visible on the desktop;
- double-clicking the desktop shortcut opens Modbus Replicator;
- **Modbus Replicator** is also present in the Start/Application menu;
- launching it from Start opens the same Replicator application type;
- the window has a left device list and right device-definition editor;
- left controls include Search, Add, Duplicate, Delete;
- right source fields include Name, Enabled, Endpoint, Unit ID, FC, Start, Count, Scan Rate;
- destination section includes Port, Unit ID, Auto Port, Auto Unit ID, Owner, Status;
- footer/runtime area includes Replicator, Source, Last Poll, plus a Last Error display when applicable.

A generic workstation application icon is acceptable for this milestone; the requirement is a working desktop launcher and Start-menu launcher, not custom artwork.

### 8. Controlled Simulator source setup

Use the existing Modbus Simulator UI to create or reuse one enabled test source. Prefer these controlled values when the environment is fresh:

```text
Name: JR-Rep-Source
Port: 5020
Unit ID: 1
FC3 Start: 0
FC3 Count: 4
FC3 random interval: 60000 ms
```

Other FC ranges may remain unused. Save & Apply and wait until the Simulator shows MMA2/Simulator running normally.

Expected ownership after this step: `(5020,1)` belongs to `simulator`.

### 9. Replicator auto-allocation, Save & Apply, persistence, and runtime status

In Modbus Replicator, Add a device using:

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

Verify before/after Save & Apply:

- destination is auto-filled without using the Simulator-owned `(5020,1)` reservation;
- Save & Apply succeeds;
- Owner becomes `replicator` and Status becomes `OWNED` after apply;
- after a few poll intervals, Replicator shows `RUNNING`, Source shows `OK`, and Last Poll updates;
- close and reopen Modbus Replicator: `JR-Rep-Device` reloads from persisted configuration with the same resolved destination;
- make an unsaved field edit, click Discard, and verify the persisted value returns.

Also exercise Duplicate once:

- Duplicate `JR-Rep-Device`;
- the duplicate must receive a distinct automatic destination suggestion rather than silently sharing the first device's `(port, unit_id)`;
- delete the unsaved duplicate before continuing so the test remains simple.

### 10. Foreign-ownership collision guard

On `JR-Rep-Device`:

- disable Auto Port and Auto Unit ID;
- manually set Destination Port `5020`, Unit ID `1` (the Simulator-owned reservation);
- verify the UI reports `Owner: simulator` and `Status: IN USE` once the manual destination is inspected;
- attempt Save & Apply.

Expected:

- Save & Apply is rejected truthfully as a foreign ownership collision;
- Simulator ownership is not overwritten or deleted;
- Replicator must not claim `(5020,1)`.

Capture ownership evidence:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
```

Expected: the Simulator `(5020,1)` entry remains owned by `simulator`; any successfully applied Replicator destination remains separately owned by `replicator`.

Then use Discard so the valid persisted Replicator destination is restored.

### 11. Runtime error truthfulness and recovery

Change only the Replicator source Endpoint to an intentionally unreachable local endpoint such as:

```text
127.0.0.1:1
```

Save & Apply.

Expected within a few scan intervals:

- Replicator remains `RUNNING` (poll loop alive);
- Source becomes `ERROR`;
- Last Error contains the connection failure;
- the poll loop continues trying on later intervals rather than exiting or fabricating success.

Restore Endpoint to `127.0.0.1:5020`, Save & Apply again, and verify Source returns to `OK` and Last Poll advances.

### 12. Delete/release ownership guard

Delete `JR-Rep-Device` and Save & Apply.

Expected:

- the Replicator-owned reservation for that device is released;
- the Simulator-owned `(5020,1)` reservation remains untouched;
- no foreign memory/reservation is removed.

Capture final ownership evidence again:

```bash
docker exec mcs-modbus-replicator-runtime cat /data/config/mma2/owners.yaml
```

Do not delete the Simulator test source merely to make the test pass.

### 13. Final repository-state check

From repository root:

```bash
git status --short
```

Expected: no unexpected tracked changes from testing. Do not clean/reset/restore unexpected changes; report them.

### 14. JR report

Replace only the content below `## JR TEST REPORT` with the actual evidence. Preserve every other section exactly.

Use this structure:

```text
Verdict: PASS | FAIL | BLOCKED
Tested commit: <git rev-parse HEAD before report edit>

Repository state:
<pull/status/commit evidence>

Replicator formatting + focused tests:
<commands, exit codes, exact meaningful output>

Simulator-to-Replicator E2E:
<command, exit code, output>

Full Replicator tests + vet:
<commands, exit codes, output>

OS.js build:
<commands, exit codes, package discovery/build evidence>

Docker deployment:
<compose config/build/up/ps/log evidence>

Launcher/UI shell:
<desktop launcher, Start launcher, fields/layout observations>

Auto allocation + persistence:
<resolved destination, Save & Apply, reopen, Discard, Duplicate observations>

Ownership collision:
<UI owner/status, rejected apply, owners.yaml evidence>

Runtime status/error recovery:
<RUNNING/OK, ERROR + last error, recovered OK observations>

Delete/release ownership:
<final owners.yaml evidence>

Final git status:
<exact output>

Unexpected behavior:
<none or exact observations>
```

After writing the report, JR may commit and push **handoff.md only**:

```bash
git add handoff.md
git diff --cached -- handoff.md
git commit -m "JR report Replicator UI stack verification"
git push origin main
```

JR must stop after the report push. Do not fix failures, edit source, archive tasks, modify Active Work, update ICC, or select further work.

## JR TEST REPORT

Verdict: FAIL
Tested commit: 584fcbaa8762f948b066dd4d3c00aadd30a1cc96

Repository state:
(pull fast-forward 13e87ec..584fcba succeeded; git status --short before test: empty; HEAD 584fcbaa8762f948b066dd4d3c00aadd30a1cc96)

Replicator formatting + focused tests:
Command: gofmt -l .
Exit/result: exit code 0 (no output)
Command: go test -count=1 -run '^(TestDocumentSaveLoadRoundTrip|TestSuggestDestinationSkipsOccupiedReservations|TestResolveManualForeignCollisionReportsOwner|TestComposeDocumentPreservesForeignAndAllReplicatorReservations|TestDeviceRuntimeConfigMapsOneToOneRange|TestRuntimeManagerApplyLifecycleAndStatus)$' .
Exit/result: exit code 0 — ok github.com/tamzrod/MCS.OSJS/replicator  0.049s
Output: go: downloading gopkg.in/yaml.v3 v3.0.1; ok github.com/tamzrod/MCS.OSJS/replicator  0.049s

Simulator-to-Replicator E2E:
Command: go test -count=1 -run '^TestSimulatorToReplicatorE2E$' .
Exit/result: exit code 0 — ok github.com/tamzrod/MCS.OSJS/replicator  13.417s

Full Replicator tests + vet:
Command: go test -count=1 ./...
Exit/result: exit code 0 — ok github.com/tamzrod/MCS.OSJS/replicator  12.958s; ? github.com/tamzrod/MCS.OSJS/replicator/cmd/modbus-replicator-runtime [no test files]
Command: go vet ./...
Exit/result: exit code 0 (no output, no diagnostics)

OS.js build:
(No package-lock.json is committed;ware npm ci was unavoidable. Installed deps sandbox-locally with npm install --no-package-lock (956 packages, exit 0,no tracked mutation; Node v16.20.2 user-local per repo engines <17).
Command: npm run build:local-packages
Exit/result: exit code 0 — built 4 local packages exactly once: ModbusReplicator, ModbusSimulator, NamelessClassicIcons, NamelessWorkstationTheme
Command: npm run package:discover
Exit/result: exit code 0 — discovered 6 package(s): 4 local incl. modbus-replicator as ModbusReplicator [symlink, local],and 2 npm
Command: npm run build
Exit/result: exit code 0 — webpack build completed
Post-build git status: empty; dist/ and packages.json are git-ignored normal build output.

Docker deployment:
Command: docker compose config
Exit/result: exit code 0 (config valid)
Command: docker compose build osjs-shell modbus-simulator-runtime modbus-replicator-runtime mma2
Exit/result: exit code 0 — all four images built: mcs-osjs-shell:latest, mcs-modbus-simulator-runtime:latest, mcs-modbus-replicator-runtime:latest, mcs-mma2:latest
Command: docker compose up -d; docker compose ps
Result: mcs-osjs-shell Up (healthy) port 18209; modbus-simulator-runtime Up; modbus-replicator-runtime Up; **mma2 Restarting (1) 18s — NOT running**
Command: docker compose logs --no-color --tail=120 modbus-replicator-runtime mma2
Replicator runtime: "Replicator runtime listening on /data/run/modbus-replicator.sock" (starts fine)
mma2: "[mma2 v2.0.2] config loaded and validated successfully; authority policies loaded; notify engine enabled; mma2 ingress started; fatal error: all goroutines are asleep - deadlock!", goroutine 1 [select (no cases)] main.main() /build/cmd/mma2/main.go:162 +0x109c; exited unexpectedly: exit status 2; supervisor restart loop repeats.)

Launcher/UI shell:
osjs-shell loads both package servers: "Loading /src/packages/ModbusSimulator/server.js", "Loading /src/packages/ModbusReplicator/server.js"; server listening http://0.0.0.0:18209; /healthz => {"status":"ok","shell":"neutral"}
Desktop: "Modbus Replicator" launcher iconvisible (generic application-x-executable.svg icon); Start menu > Development > "Modbus Replicator" and "Modbus Simulator" both present; launching Modbus Replicator from Start opens the window with left device list (Search devices... input, Add, Duplicate, Delete) and right Device Definition editor.
Editor fields observed: Name, Enabled, Endpoint, Unit ID, FC (FC3/FC4 select), Start, Count, Scan Rate (ms); Destination: Port, Auto Port, Unit ID, Auto Unit ID, Owner, Status; footer Replicator/Source/Last Poll/Last Error area; Save & Apply/ Discard buttons.
Persisted pre-existing devices on volume: Replicator "Rep-PLC-1" 127.0.0.1:5020 FC3 Enabled (Owner replicator, Status AVAILABLE); Simulator "Sim-PLC-1" Port 5020 - Unit 1.
Desktop-icon double-click launch via browser automation was inconclusive (single-click does not launch in automation; Start-menu launch verified instead)

Auto allocation + persistence:
(Blocked by MMA2 deadlock. Sim-PLC-1 source exists on volume; Replicator Rep-PLC-1 editor present with Owner replicator Static AVAILABLE. Save & Apply could not complete ownership allocation end-to-end because the MMA2 ownership/raw-ingest backend is serviceless (mma2 crash-loop). No clean persisted-state close/reopen, Discard, or Duplicate flow could be completed.)

Ownership collision:
(Blocked by MMA2 deadlock. owners.yaml in mcs-modbus-replicator-runtime at /data/config/mma2/owners.yaml reads: reservations: [] — no ownership entries materialized; UI Owner/Status fields rendered but no backend reservation could be created or inspected end-to-end.)

Runtime status/error recovery:
(Blocked by MMA2 deadlock. Replicator runtime service itself started fine (listening on /data/run/modbus-replicator.sock); Source OK/ERROR transitions and Last Poll update could not be verified end-to-end because MMA2 backend is deadlocked.)

Delete/release ownership:
(Blocked by MMA2 deadlock. No ownership entries existed to release; final owners.yaml: reservations: [] — Simulator (5020,1) reservation never materialized because MMA2 cannot run.)

Final git status:
(empty — no unexpected tracked changes from testing)

Unexpected behavior:
mma2 container crash-loops with a fatal Go deadlock: "fatal error: all goroutines are asleep - deadlock!" at main.main() main.go:162 (empty config: listeners: []/owners: [] on fresh volume). This blocks Docker gate (all four running expected) and all downstream interactive MMA2-backed verifications (steps 8-12). Simulator/Replicator runtimes and osjs-shell work. Also: npm ci cannot run (no committed package-lock.json; condemned by packet rule against manifest alteration; used sandbox-local npm install --no-package-lock instead; all build gates passed.)
