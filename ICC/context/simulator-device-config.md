# Simulator Device Configuration (SIM-001)

Baseline commit: 8cb9f0c08ae4e5717ce22d4423fd297eba03d170
Working tree: clean
Source dependencies: simulator/device.go, simulator/validate.go, simulator/store.go, simulator/store_test.go, simulator/README.md, deploy/docker-compose.yml, OSJS/src/server/config.js, MMA2/internal/config/validate.go, planning/Brainstorm/osjs-modbus-simulator.md, workflow/active_work/sim-001-simulator-device-config.md
Parent: active-work
Zoom In: (none; leaf node)
Zoom Out: active-work

## Established truth

Repository-root `simulator/` owns persistent device definitions. It is not MMA2 effective config and is not written into `MMA2/` or `OSJS/` packaged source.

One device has two persisted domains:
- identity: name, enabled
- mma2: port, unit_id, fc1-fc4 start/count
- random_runtime: fc1-fc4 interval_ms

Host-mounted root is `OSJS_DATA_DIR` only (`osjs-data` volume -> `/data`). Persist path: `$OSJS_DATA_DIR/config/simulator/devices.yaml`. Unset env is an error; no invented host path.

Validation (from MMA2 validate.go + brainstorm): port > 0; unit_id <= 255; unused FC count 0; start+count within 16-bit space; configured FC requires interval_ms > 0. Invalid SaveOne leaves prior file bytes unchanged.

SIM-001 is DONE. SIM-002 activation/lifecycle is a separate Active Work item and is outside this node.
