# Simulator Device Config + MMA2 Compose/Ownership (SIM-001 + SIM-002A)

Baseline commit: 8cb9f0c08ae4e5717ce22d4423fd297eba03d170
Working tree: dirty
Audited Overlay: simulator/mma2_config.go, simulator/compose_test.go, workflow/active_work/sim-002a-mma2-config-ownership.md, handoff.md
Source dependencies: simulator/device.go, simulator/validate.go, simulator/store.go, simulator/store_test.go, simulator/mma2_config.go, simulator/compose_test.go, simulator/README.md, deploy/docker-compose.yml, OSJS/src/server/config.js, MMA2/internal/config/validate.go, planning/Brainstorm/osjs-modbus-simulator.md, workflow/active_work/sim-001-simulator-device-config.md, workflow/active_work/sim-002a-mma2-config-ownership.md
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
- SIM-002A is DONE(2026-09-07). Six tests in `simulator/compose_test.go` cover free-save, foreign-collision reject, own-update, delete-own-only, foreign-delete reject,and invalid-save-no-persist. SIM-002B lifecycle activation is the next Active Work item,outside this node.
