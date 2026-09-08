# Simulator Device Config + MMA2 Compose/Ownership + Scheduler + Raw Ingest + SaveDocument + Loopback Bridge + Managed MMA2 Lifecycle (SIM-001 → SIM-008)

Baseline commit: 7e63460
Working tree: clean
Audited Overlay:(none(; files once audited as overlay (`simulator/device.go`,`simulator/store.go`,`simulator/store_document_test.go`( are now committed in HEAD, alongside `simulator/bridge.go`+`simulator/bridge_test.go`.
Source dependencies: simulator/device.go, simulator/validate.go, simulator/store.go, simulator/store_test.go, simulator/store_document_test.go, simulator/mma2_config.go, simulator/compose_test.go, simulator/scheduler.go, simulator/scheduler_test.go, simulator/raw_ingest.go`, simulator/raw_ingest_test.go`, simulator/bridge.go`, simulator/bridge_test.go`, simulator/apply.go`, simulator/apply_test.go`, simulator/lifecycle.go`, simulator/lifecycle_test.go`, simulator/cmd/simbridge/main.go`, simulator/README.md`, deploy/docker-compose.yml`, OSJS/src/server/config.js`, OSJS/src/server/providers/simulator-bridge.js`, MMA2/internal/config/validate.go`, planning/Brainstorm/osjs-modbus-simulator.md`, workflow/active_work/sim-001-simulator-device-config.md`, workflow/active_work/sim-002a-mma2-config-ownership.md`, workflow/active_work/sim-003-random-runtime.md`, workflow/active_work/sim-004-raw-ingest.md`, workflow/active_work/sim-005-osjs-window.md`, workflow/active_work/sim-006-save-apply-routing.md`, workflow/active_work/sim-007-runtime-status.md`, workflow/active_work/sim-008-managed-mma2-lifecycle.md`, workflow/active_work/sim-009-serve-simulator-through-mma2.md`
Zoom In:(none; leaf node)
Zoom Out: active-work

## Established truth (SIM-001)

Repository-root `simulator/` owns persistent device definitions. It is not MMA2 effective configand is not written into `MMA2/` or `OSJS/` packaged source.

One device has two persisted domains:
- identity: name, enabled
- mma2: port, unit_id, fc1-fc4 start/count
- random_runtime: fc1-fc4 interval_ms

Host-mounted root is `OSJS_DATA_DIR` only (`osjs-data` volume -> `/data`). Persist path: `$OSJS_DATA_DIR/config/simulator/devices.yaml`. Unset env is an error; no invented host path.

Validation (from MMA2 validate.go + brainstorm): port > 0; unit_id <= 255; unused FC count  ﻿0; start+count within 16-bit space; configured FC requires interval_ms > 0. Invalid SaveOne leaves prior file bytes unchanged.

## Established truth (SIM-002A)

- Compose artifacts(both atomic temp+rename): `$OSJS_DATA_DIR/config/mma2/config.yaml` (effective MMA2 runtime config)and `$OSJS_DATA_DIR/config/mma2/owners.yaml` (ownership registry; `OSJS_DATA_DIR` unset → error;no invented host path.
- `(port,unit_id)`is the unique reservation key shared by Simulator and Replicator; ownership first-come-first-save;every effective reservation carries a machine-readable YAML `owner` entry in the registry.
- Enforcement:free key → simulator save/delete becomes owner;existing key + same owner → update/delete allowed;existing key + different owner → `ErrReservationOwnedByOther` rejected before any write, leaving prior effective config + registry byte-unchanged.
- Compose/drop operates reservation-locally:merge into an existing listener on the same TCP port,never touching other listeners'memory or ownership entries;FC mapping:FC1=coils,FC2=discrete_inputs,FC3=holding_registers,FC4=input_registers;areas with count  ﻿0 omitted;per-memory allow-all policy mirrors MMA2 smoke-test config.
- SIM-002A is DONE(2026-09-07). Six tests in `simulator/compose_test.go` cover free-save, foreign-collision reject, own-update, delete-own-only, foreign-delete reject,and invalid-save-no-persist. SIM-002B lifecycle activation is DONE(2026-09-07(,evidenced in its own task file;it is outside this node's source deps but the live Modbus surface it proved is the ingestion target for SIM-004(.

## Established truth (SIM-003 random-runtime scheduler)

- `simulator/scheduler.go` implements per-FC independent randomization schedules;consumes only simulator-owned `random_runtime` interval domain(`FC1IntervalMS`..`FC4IntervalMS`(,loaded from `DeviceDefinition.RandomRuntime`;FC area counts are read once from the device's MMA2 params merely to size generated batches;unconfigured count 0 areas get no schedule entry and never fire(.
- Value generation: FC1/FC2 booleans(coils/discrete-inputs mnemonic mapping per SIM-002A(;FC3/FC4 uint16 within 16-bit space(holding/input registers(. `math/rand` seeded per scheduler instance;`onUpdate` callback fired from scheduler goroutine per fire(.
- Timing state: per-FC `last`/`next` tracked in scheduler;`Timing()` returns snapshot map for SIM-007 runtime-status surface;;`Stop()` closes stop channel,waits goroutine exit(.
- Timing-only change: `UpdateTiming(RandomRuntimeParams)` recomputes only affected FC's next deadline from now,preserves existing last-update state,and never touches MMA2(config/owners/memory/process(` — no restart,no config writes (per SIM-003 Acceptance Criteria 3 and SIM-006 routing boundary(.
- SIM-003 is DONE(2026-09-07(. `simulator/scheduler_test.go` (139 lines( covers cadence concurrency(FC1 5ms vs FC2 10ms>8 vs >=4 fires in ~105ms;; FC3/FC4 count 0 never fired( and interval-change-without-restart(FC3 9ms→2ms next deadline recomputed,FC1/FC2/FC4 retained(;`go vet ./...` clean;`go test -count=1 ./...` → `ok github.com/tamzrod/MCS.OSJS/simulator 0.162s`(.

## Established truth (SIM-004 raw ingest)

- `simulator/raw_ingest.go` implements the MMA2 raw-ingest v1 client:`encodeRawPacket(unitID, addr, count uint16, area uint8, values Values)` frames the MMA2 raw-ingest byte contract introduced by SIM-002B(header RI+version+area+BE uint16 unitID/addr/count+payload;FC1/FC2 coil/discrete values pack LSB-first bit regions,FC3/FC4 register values are big-endian uint16;`sendRawPacket(addr string, pkt []byte)` dials TCP, sends one frame, reads one response byte,,response mismatch -> error.
- `(*RawIngestClient.Send(v Values)` is the sole simulator-to-MMA2 memory path:it maps FC1-FC4 mnemonics onto the device simulator-owned MMA2 ranges,rejects unconfigured count-0 areas,and targets exclusively raw ingest—no direct-memory and no Modbus-write population.
- SIM-004 is DONE(2026-09-07(. `simulator/raw_ingest_test.go` (4 tests): bit-area LSB packing bytes,,reg-area big-endian bytes,,real-TCP Send round-trip against a fixture server honoring frames with OK(0x06(/reject(0x21( responses,,and unconfigured-FC rejection(;`gofmt -l .` clean;`go vet ./...` clean;`go test -count=1 ./...` -> `ok github.com/tamzrod/MCS.OSJS/simulator`.`
## Established truth (SIM-005 backend persistence layer)

- `Store.SaveDocument(doc Document)` (`simulator/store.go`) is the multi-device atomic save path the approved OS.js window (SIM-005)uses for Add/Duplicate/Delete/edit:it validates every device in doc first;any invalid device aborts the whole save leavingthe previously persisted file bytes unchanged;then one atomic temp+rename replace( persists the entire simulator-owned document. `SaveOne` now delegates to `SaveDocument(Document{Devices: []DeviceDefinition{def}}}`preserving SIM-001 single-device semantics.
- All simulator model structs in `simulator/device.go` (DeviceDefinition, MMA2Params, Area, RandomRuntimeParams, Document(now carry both `yaml` and `json` tags;the JSON wire format bridges browser form values to the SIM-001 Go model via the loopback simulator-server bridge (SIM-005; no backend MMA2 activation in scope(.
- `simulator/store_document_test.go` proves exact multi-device round-trip through Store, atomic rejection on invalid device without replacing prior bytes,and the Add/Duplicate/Delete load-modify-save reshape persisting the whole document. `gofmt -l .` clean;`go vet ./...` clean;`go test -count=1 ./...` -> `ok github.com/tamzrod/MCS.OSJS/simulator`.`

## Established truth (SIM-005 OS.js window)

- `OSJS/src/packages/ModbusSimulator/index.js` and `index.scss` implement the approved two-pane searchable device-list/editor and bind all simulator-owned fields/actions to the SIM-001 JSON document through same-origin `GET/PUT /api/devices`.
- Add, Duplicate, Delete, and edits remain local until Save & Apply; Discard restores the last persisted snapshot. SIM-005 does not activate MMA2; the UI explicitly leaves live application to SIM-006.
- Real-browser verification proved open/add/edit/save/duplicate/delete/discard and exact JSON/YAML round-trip in an isolated data root containing no effective MMA2 config. Production package/full builds, server syntax check, gofmt, vet, and Go tests pass. SIM-005 completed at pushed commit `28d27f9`.

## Established truth (SIM-006 Save & Apply routing)

- `simulator/apply.go` validates and classifies full-document edits as `mma2-structural`, `random-runtime`, or `no-change`; only the selected downstream consumer runs, and the simulator document is persisted only after downstream success.
- Structural edits compose the complete simulator reservation set through SIM-002 ownership rules and replace schedulers; timing-only edits call `Scheduler.UpdateTiming` without MMA2 configuration writes. Applying bridge responses expose the path/message to the OS.js status bar.
- Unit/HTTP tests prove all routes and prior-state preservation. Three browser saves proved structural port change, timing-only interval change, and rejected duplicate reservation with the last valid document/config intact. SIM-006 completed at pushed commit `9532c31`.

## Established truth (SIM-007 runtime status)

- The applying bridge exposes selected-device runtime status: device RUNNING/STOPPED/ERROR, MMA2 TCP reachability, raw-ingest health/error, exact total configured points, and FC1-FC4 scheduler Last/Next timestamps.
- The OS.js window polls once per second and renders only compact current status—no register table, history, chart, or new control behavior.
- A live-fixture Go test proves RUNNING state, all four timing streams, and exact point total; real-browser verification proves truthful unavailable/error display when MMA2 is absent. SIM-007 completed at pushed commit `ceee393`.
## Established truth (SIM-008 managed MMA2 lifecycle)

- `simulator/lifecycle.go`: one mutex-serialized MMA2 child-process owner. `MMA2Activator` (interface: `Activate(configPath, cfg(` + `Stop((`; `MMA2Lifecycle` starts a child from `$MMA2_BINARY` (env; fallback `mma2` on PATH( via `mma2 <config.yaml>`, waits for all effective listeners plus child-stability brief confirmation, gracefully stops (SIGINT→1s kill fallback(,, replaces the process only after the replacement is proven ready,and restores the prior process + config on replacement failure via a `.rollback.yaml` temp artifact. `effectiveListeners(cfg)` maps the composed `EffectiveMMA2Config.Listeners`; bare hosts (`""`,`0.0.0.0`,`::(`) are dialed via `127.0.0.1`.
- `simulator/apply.go` (post-SIM-008(: `SchedulerApplier` carries a `lifecycle MMA2Activator`; structural apply snapshots the effective config + ownership files, composes through `Store.ComposeDocument` (ownership-safe(,, then `lifecycle.Activate(effectiveConfigPath,cfg(` and on any failure restores both artifact files byte-for-byte before returning the error (simulator document persistence remains after downstream success(. `Stop()` also stops the managed MMA2 child.
- `NewRuntimeApplyRouter(store(` bootstraps `NewMMA2Lifecycle()`, activates an existing effective config at startup,and falls the whole router on activation failure (simbridge dies with the error(. `SIMULATOR_BRIDGE_ADDR` env permits isolated loopback verification; `127.0.0.1:18211` remains the default bridge bind。 `SchedulerApplier.replaceSchedulers` starts one `Scheduler` per enabled device whose `onUpdate` sends generated values through `NewRawIngestClient(device(` (SIM-004 raw ingest((and records `rawErrors[device]` (cleared on success(.
- `simulator/lifecycle_test.go`: `listenerProcess` fake MMA2 child binds fixed listeners;`TestMMA2LifecycleActivationReplacementAndRollback` proves initial activation, replacement (second listener up, first released(,,and forced-failure rollback restoringthe prior listener. SIM-008 completed+verified at pushed commit `7e63460` (2026-09-08(;its task file records the real MMA2 workflow (build from `MMA2/cmd/mma2`; isolated data root; bridge `127.0.0.1:18212`; accepted structural PUT moved service `15031`→`15032`; foreign collision returned 422 with exactly one managed child restored(.
