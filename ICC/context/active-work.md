# Active Work Program (Simulator + MMA2)

Baseline commit: 8cb9f0c08ae4e5717ce22d4423fd297eba03d170
Working tree: clean
Source dependencies: handoff.md, workflow/active_work/README.md, workflow/active_work/mma2-basic-install-test.md, workflow/active_work/sim-001-simulator-device-config.md, workflow/active_work/sim-002-mma2-activation.md, workflow/active_work/sim-003-random-runtime.md, workflow/active_work/sim-004-raw-ingest.md, workflow/active_work/sim-005-osjs-window.md, workflow/active_work/sim-006-save-apply-routing.md, workflow/active_work/sim-007-runtime-status.md, MMA2/testdata/smoke-test.yaml, simulator/device.go, simulator/store.go, simulator/validate.go
Parent: L0-project
Zoom In: simulator-device-config
Zoom Out: L0-project

## Execution Order (handoff)

1. `mma2-basic-install-test.md` — **COMPLETED** 2026-09-06 (MMA2-001 import+build; MMA2-002 smoke test).
2. `sim-001-simulator-device-config.md` — **COMPLETED** 2026-09-06. Simulator-owned two-domain persist under `$OSJS_DATA_DIR/config/simulator/devices.yaml`. See Zoom In `simulator-device-config`.
3. `sim-002-mma2-activation.md` — **CURRENT**. Validate+activate simulator MMA2 params against shared MMA2 namespace. Task sizing 4/10; split if composition and lifecycle are independently unresolved.
4. `sim-003-random-runtime.md` — per-FC random runtime scheduler.
5. `sim-004-raw-ingest.md` — raw ingest into MMA2 memory.
6. `sim-005-osjs-window.md` — OS.js Modbus Simulator window.
7. `sim-006-save-apply-routing.md` — Save & Apply routing.
8. `sim-007-runtime-status.md` — runtime status surface.

A later task does not authorize skipping an incomplete dependency.
