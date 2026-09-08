# Simulator Device Config + Local Runtime Integration (SIM-001 → SIM-022)

Baseline commit: a1f03e6
Working tree: clean
Audited Overlay:(none(; files once audited as overlay (`simulator/device.go`,`simulator/store.go`,`simulator/store_document_test.go`( are now committed in HEADand the SIM-010 boundary files (`simulator/bridge.go`,`simulator/bridge_test.go`,`simulator/cmd/simbridge/main.go`,`OSJS/src/server/providers/simulator-bridge.js`( were deleted at `0b356da`,;none remain in the working tree.
Source dependencies: docs/SIMULATOR_RUNTIME_INTEGRATION.md, MMA2/pkg/configvalidate/validate.go, simulator/*, deploy/docker-compose.yml, OSJS/src/server/*, OSJS/src/packages/ModbusSimulator/, workflow/active_work/sim-018-decide-local-ui-runtime-boundary.md through workflow/active_work/sim-022-visible-end-to-end-verification.md, workflow/archive/sim-001-simulator-device-config.md through workflow/archive/sim-017-end-to-end-simulator-mma2-verification.md
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

- `OSJS/src/packages/ModbusSimulator/index.js` and `index.scss` implement the approved two-pane searchable device-list/editor bound to the SIM-001 domain model (two-domain MMA2+random-runtime fields, FC1-FC4 ranges, Add/Duplicate/Delete/Discard/Save(. SIM-010 removed the `/api/devices` transport and bridge-served runtime-status panel+status polling;the UI now edits local state and `save()` normalizes+mirrors to `persisted` locally without any HTTP boundary, pending SIM-012's local persistence rewiring. Save & Apply keeps its validation gating.
- Add, Duplicate, Delete, and edits remain local until Save & Apply; Discard restores the last persisted snapshot. SIM-005 does not activate MMA2; the UI explicitly leaves live application to SIM-006.
- Real-browser verification proved open/add/edit/save/duplicate/delete/discard and exact JSON/YAML round-trip in an isolated data root containing no effective MMA2 config. Production package/full builds, server syntax check, gofmt, vet, and Go tests pass. SIM-005 completed at pushed commit `28d27f9`.

## Established truth (SIM-006 Save & Apply routing)

- `simulator/apply.go` validates and classifies full-document edits as `mma2-structural`, `random-runtime`, or `no-change`; only the selected downstream consumer runs, and the simulator document is persisted only after downstream success.
- Structural edits compose the complete simulator reservation set through SIM-002 ownership rules and replace schedulers; timing-only edits call `Scheduler.UpdateTiming` without MMA2 configuration writes. (SIM-010 removed the HTTP bridge that exposed apply path/message to the status bar)
- SIM-010 removed the bridge HTTP tests along with the bridge; remaining unit tests prove all apply paths and prior-state preservation. Three browser saves proved structural port change, timing-only interval change, and rejected duplicate reservation with the last valid document/config intact. SIM-006 completed at pushed commit `9532c31`.

## Established truth (SIM-007 runtime status)

- SIM-007's runtime-status surface was delivered through the applying bridge (`/api/devices/status`( and the OS.js window's 1s polling+compact status panel (device/MMA2/raw-ingest state, exact total points, FC1-FC4 Last/Next(; with SIM-010 both the bridge transport and the bridge-served runtime-status UI were removed. The Go `SchedulerApplier.RuntimeStatus` API itself remains in `simulator` for downstream reuse(.)
- A live-fixture Go test proves RUNNING state, all four timing streams, and exact point total; real-browser verification proves truthful unavailable/error display when MMA2 is absent. SIM-007 completed at pushed commit `ceee393`.
## Established truth (SIM-008 managed MMA2 lifecycle)

- `simulator/lifecycle.go`: one mutex-serialized MMA2 child-process owner. `MMA2Activator` (interface: `Activate(configPath, cfg(` + `Stop((`; `MMA2Lifecycle` starts a child from `$MMA2_BINARY` (env; fallback `mma2` on PATH( via `mma2 <config.yaml>`, waits for all effective listeners plus child-stability brief confirmation, gracefully stops (SIGINT→1s kill fallback(,, replaces the process only after the replacement is proven ready,and restores the prior process + config on replacement failure via a `.rollback.yaml` temp artifact. `effectiveListeners(cfg)` maps the composed `EffectiveMMA2Config.Listeners`; bare hosts (`""`,`0.0.0.0`,`::(`) are dialed via `127.0.0.1`.
- `simulator/apply.go` (post-SIM-008(: `SchedulerApplier` carries a `lifecycle MMA2Activator`; structural apply snapshots the effective config + ownership files, composes through `Store.ComposeDocument` (ownership-safe(,, then `lifecycle.Activate(effectiveConfigPath,cfg(` and on any failure restores both artifact files byte-for-byte before returning the error (simulator document persistence remains after downstream success(. `Stop()` also stops the managed MMA2 child.
- `NewRuntimeApplyRouter(store(` bootstraps `NewMMA2Lifecycle()`, activates an existing effective config at startup,and falls the whole router on activation failure (simbridge dies with the error(. `SIMULATOR_BRIDGE_ADDR` env permits isolated loopback verification; `127.0.0.1:18211` remains the default bridge bind。 `SchedulerApplier.replaceSchedulers` starts one `Scheduler` per enabled device whose `onUpdate` sends generated values through `NewRawIngestClient(device(` (SIM-004 raw ingest((and records `rawErrors[device]` (cleared on success(.

## Established truth (SIM-009 committed ready-gate delta; superseded premise

- Commit `83a5b69` (SIM-009, visible in HEAD( changed `simulator/apply.go`+`apply_test.go`: `SchedulerApplier` gained a `ready` flag; `replaceSchedulers` arms schedulers only when `ready && device.Enabled`; `ApplyStructural` sets `ready=true` after a successful compose+activate so schedulers start only for the resulting config;`RuntimeStatus` gates on `ready`—pre-ready surfaces truthful `STOPPED`/`ERROR`/`WAITING` (no `RUNNING`/`OK`, no FC timing(; and when an ingest failure occurs after arming, Device/RawIngest report `ERROR` with the raw error string—never a false healthy claim. `NewRuntimeApplyRouter` now composes+activates at boot through the same ownership-safe path and on activation failure keeps the router alive (schedulers unarmed, UI truthful ERROR( until a later accepted structural apply recovers the service. `TestSchedulerApplierReadyGatePreventsArmingAndFalseHealth` proves schedulers stay unarmed pre-ready and status is truthful failure.
- The task's broader premise—Simulator owning/activating a managed MMA2 child—was **superseded** at `88a43f8`: authoritative boundary now is MMA2 auto-starts on boot; Simulator never START/STOP/SPAWN/KILL/REPLACE it; only RESTART (SIM-014( after a committed valid shared-config change;the `simbridge` service and Simulator config API are removed in SIM-010. Preserved reusable model/ownership/composition/scheduler/raw-ingest code lanes for SIM-010…SIM-017 in `workflow/active_work/`.
- `simulator/lifecycle_test.go`: `listenerProcess` fake MMA2 child binds fixed listeners;`TestMMA2LifecycleActivationReplacementAndRollback` proves initial activation, replacement (second listener up, first released(,,and forced-failure rollback restoringthe prior listener. SIM-008 completed+verified at pushed commit `7e63460` (2026-09-08(;its task file records the real MMA2 workflow (build from `MMA2/cmd/mma2`; isolated data root; bridge `127.0.0.1:18212`; accepted structural PUT moved service `15031`→`15032`; foreign collision returned 422 with exactly one managed child restored(.

## Established truth (SIM-011 independent MMA2 lifecycle)

- SIM-011 supersedes and removes SIM-008's process-ownership implementation. `simulator/lifecycle.go` and `lifecycle_test.go` no longer exist; Simulator code has no MMA2 spawn, signal, kill, stop, replace, `MMA2_BINARY`, `MMA2Lifecycle`, or `MMA2Activator` path.
- `SchedulerApplier.ApplyStructural` performs ownership-safe `ComposeDocument` only. `SchedulerApplier.Stop` stops Simulator schedulers only. `NewRuntimeApplyRouter` composes persisted reservations but never activates MMA2 and keeps schedulers unarmed pending SIM-015's independent apply/readiness work.
- `TestStructuralApplyAndStopDoNotControlIndependentMMA2` proves shared-config composition while an independent listener remains reachable after Simulator stop. Simulator formatting, vet, and full tests pass at pushed commit `82d6ac6`.

## Established truth (SIM-012 local Simulator persistence)

- The Modbus Simulator window stores its normalized document through the existing server-backed `osjs/settings` service under `mcs/modbus-simulator.document`. It loads that namespace on window render and advances its Discard snapshot only after `settings.save()` succeeds.
- This is the standard per-user OS.js settings path (`/data/vfs/<user>/.osjs/settings.json`), not a Simulator HTTP API. Saving does not call Simulator Go composition, create shared MMA2 config, restart MMA2, or arm schedules.
- Live rebuilt-container verification exercised add/edit/all fields/save/reload/duplicate/delete/discard and proved exact JSON field round-trip at pushed commit `e14ee96`; JS/package/full builds and Simulator regressions pass.

## Established truth (SIM-013 validated shared-config composition)

- `ComposeDocument` translates enabled definitions only, removes/replaces only Simulator-owned `(port, unit_id)` reservations, and retains foreign listeners, memories, ownership records, and inline YAML fields.
- `MMA2/pkg/configvalidate.YAML` exposes the authoritative MMA2 parser/validator without starting MMA2. The complete candidate is marshaled and validated before any replace; config is restored byte-for-byte if the owners replace fails.
- Fixtures prove expected FC mapping, disabled omission, own update, foreign preservation, collision rejection,and invalid-candidate rejection with both shared artifacts unchanged. No restart, lifecycle ownership, or schedule arming is part of SIM-013. Completed at pushed commit `24fd5b1`.

## Established truth (SIM-014 restart-only request)

- SIM-014 added the single MMA2 control op the Simulator may issue after a successful shared-config commit: exactly one RESTART. `ApplyStructural` writes a machine-readable `restart-request.yaml` artifact beside `config.yaml`/`owners.yaml` (fingerprint+requested_at+composed ports( only after `ComposeDocument` succeeds, then waits for every composed listener port to accept a TCP dial (`WaitMMA2Ready`(,and clears the request only when readiness is confirmed. On restart/readiness failure it returns error without claiming apply success,and the request remains pending as truthful evidence. `composedSimulatorPorts` dedups enabled+area devices. No start/stop/spawn/kill/replace op, process ownership, HTTP API, or general MMA2 control API was added. Completed at pushed commit `db2276c`; four restart tests and live MMA2 reload (port 5020→5021( verified.
## Established truth( SIM-015 scheduling arming after apply/readiness

- ApplyRouter arms enabled schedules only after the full Save & Apply chain succeeds( shared-config commit, MMA2 RESTART, ready,persisted(. `ArmSchedules` sets the ready gate and replaces prior schedules with the latest persisted state;disabled devices never arm;restart/readiness failure returns apply error before persistence with nothing armed. Raw-ingest faults surface truthfully via runtime status. Two arming tests+existing scheduler/raw-ingest suites pass. Completed at pushed commit `3652b98`.

## Established truth( SIM-016 boot restore without restart

- Simulator boot path `NewRuntimeApplyRouter` loads persisted definitions, ownership-safe composes Simulator reservations into the already-persisted shared MMA2 config ( SIM-011 behavior(,then waits for the independently auto-started MMA2 to accept every composed listener port via `WaitMMA2Ready`,then arms enabled schedules through the SIM-015 arming gate. Ordinary unchanged boot neither restarts MMA2 nor writes a restart request. If MMA2 never becomes ready,the router survives alive with zero armed schedulers, truthful STOPPED/ERROR runtime status,and no restart artifact is created;a later Save & Apply can recover. Two boot tests in `simulator/boot_restore_test.go` prove armed-enabled-resume-without-restart-request,and unavailable-MMA2-no-arm-no-restart. Completed at pushed commit `6d24f7a`.

## Established truth (SIM-017 end-to-end verification)

- `simulator/e2e_test.go` is a verification-only harness that independently builds/starts/restarts real MMA2. The Simulator remains limited to shared-config composition, restart request, readiness, scheduling, and Raw Ingest.
- The test proves boot restore, changing scheduled FC3 data, successful real Modbus reads for FC1/FC2/FC3/FC4, exactly one restart after a valid port edit, and new-listener readiness after reload.
- Duplicate reservation is rejected with shared config byte-unchanged and no restart request. A foreign-owned reservation remains readable before/after apply and after reboot. Full MMA2 plus Simulator-router restart resumes the enabled schedule without manual apply or restart artifact. Race-enabled capstone passed at `57c8714`.

## Established truth (SIM-018 local runtime boundary)

- The browser uses the existing authenticated OS.js session WebSocket; an allowlisted OS.js server provider relays versioned `load`, `apply`, and `status` messages to a Unix-domain socket at `$OSJS_DATA_DIR/run/modbus-simulator.sock`. No Simulator HTTP route or TCP listener is permitted.
- An independently supervised Go runtime owns the long-lived `ApplyRouter`, schedulers, and Raw Ingest clients. Closing the browser does not stop simulation.
- `$OSJS_DATA_DIR/config/simulator/devices.yaml` is the single canonical document. OS.js per-user settings are not a second Simulator config store.
- The runtime never owns MMA2 lifecycle. MMA2 auto-starts independently; Simulator retains only restart-request-plus-readiness behavior after committed structural apply, and generated values use Raw Ingest.
- SIM-019 through SIM-022 are promoted in dependency order.
