# Simulator Device Config + MMA2 Compose/Ownership + Scheduler (SIM-001 + SIM-002A + SIM-003)

Baseline commit: 7cef8dd
Working tree: clean
Audited Overlay: (none (clean tree(;committed delta 21264ab..7cef8dd absorbed (handoff.md, workflow/active_work/sim-003-random-runtime.md, simulator/scheduler.go, simulator/scheduler_test.go((
Source dependencies: simulator/device.go, simulator/validate.go, simulator/store.go, simulator/store_test.go, simulator/mma2_config.go, simulator/compose_test.go, simulator/scheduler.go, simulator/scheduler_test.go, simulator/README.md, deploy/docker-compose.yml, OSJS/src/server/config.js, MMA2/internal/config/validate.go, planning/Brainstorm/osjs-modbus-simulator.md, workflow/active_work/sim-001-simulator-device-config.md, workflow/active_work/sim-002a-mma2-config-ownership.md, workflow/active_work/sim-003-random-runtime.md
Parent: active-work
Zoom In:(none; leaf node)
Zoom Out: active-work

## Established truth (SIM-001)

Repository-root `simulator/` owns persistent device definitions. It is not MMA2 effective configand is not written into `MMA2/` or `OSJS/` packaged source.

One device has two persisted domains:
- identity: name, enabled
- mma2: port, unit_id, fc1-fc4 start/count
- random_runtime: fc1-fc4 interval_ms

Host-mounted root is `OSJS_DATA_DIR` only (`osjs-data` volume -> `/data`). Persist path: `$OSJS_DATA_DIR/config/simulator/devices.yaml`. Unset env is an error; no invented host path.

Validation (from MMA2 validate.go + brainstorm): port > 0; unit_id <= 255; unused FC count  ﻿0; start+count within 16-bit space; configured FC requires interval_ms >​ 0. Invalid SaveOne leaves prior file bytes unchanged.

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
