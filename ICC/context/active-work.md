# Active Work Program (Simulator + MMA2)

Baseline commit: 28d27f9f74c097e03723ff8e731e898e28b57ea4
Working tree: clean
Audited Overlay:(none(; files once audited as overlay (`simulator/device.go`,`simulator/store.go`,`simulator/store_document_test.go`( are now committed in HEAD, alongside `simulator/bridge.go`+`simulator/bridge_test.go`.
Source dependencies: handoff.md, workflow/active_work/README.md, workflow/active_work/mma2-basic-install-test.md, workflow/active_work/sim-001-simulator-device-config.md, workflow/active_work/sim-002a-mma2-config-ownership.md, workflow/active_work/sim-002b-mma2-lifecycle-activation.md, workflow/active_work/sim-003-random-runtime.md, workflow/active_work/sim-004-raw-ingest.md, workflow/active_work/sim-005-osjs-window.md, workflow/active_work/sim-006-save-apply-routing.md, workflow/active_work/sim-007-runtime-status.md, MMA2/testdata/smoke-test.yaml, simulator/device.go, simulator/store.go, simulator/store_document_test.go, simulator/validate.go, simulator/mma2_config.go, simulator/compose_test.go, simulator/scheduler.go, simulator/scheduler_test.go, simulator/raw_ingest.go`, simulator/raw_ingest_test.go`, simulator/bridge.go`, simulator/bridge_test.go`, OSJS/src/server/providers/simulator-bridge.js`, OSJS/src/packages/ModbusSimulator/`
Parent: L0-project
Zoom In: simulator-device-config
Zoom Out: L0-project

## Execution Order (handoff)

1. `mma2-basic-install-test.md` — **COMPLETED** 2026-09-06 (MMA2-001 import+build; MMA2-002 smoke test).
2. `sim-001-simulator-device-config.md` — **COMPLETED** 2026-09-06. Simulator-owned two-domain persist under `$OSJS_DATA_DIR/config/simulator/devices.yaml`. See Zoom In `simulator-device-config`.
3. `sim-002a-mma2-config-ownership.md` — **COMPLETED 2026-09-07**. Compose+ownership protection (`(port,unit_id)` reservation keys, YAML `owner` registry at `$OSJS_DATA_DIR/config/mma2/owners.yaml`, `ErrReservationOwnedByOther` on foreign collision)。 Six tests in `simulator/compose_test.go` verified. See Zoom In `simulator-device-config`.
4. `sim-002b-mma2-lifecycle-activation.md` — **COMPLETED 2026-09-07**. Live MMA2 proven through the composed effective config (process healthy; Modbus-TCP (5020, 1) coils/discrete 64 + holding/input 100 read/write; raw PDU hreg0=4321). See Zoom In `simulator-device-config`.
5. `sim-003-random-runtime.md` — per-FC random runtime scheduler.**COMPLETED 2026-09-07** (scheduler.go+scheduler_test.go;both tests pass;timing-only UpdateTiming never restarts MMA2(. See Zoom In `simulator-device-config`.
6. `sim-004-raw-ingest.md` — raw ingest into MMA2 memory.**COMPLETED 2026-09-07** (`simulator/raw_ingest.go` + `simulator/raw_ingest_test.go`; MMA2 raw-ingest v1 client; bits-LSB-first/regs big-endian frames; Send( exclusively via raw ingest(; 4 tests incl real-TCP round-trip;gofmt/vet clean; `go test -count=1 ./...` ok. See Zoom In `simulator-device-config`.
7. `sim-005-osjs-window.md` — OS.js Modbus Simulator window. **COMPLETED 2026-09-08** at pushed commit `28d27f9` (two-pane editor, approved fields/actions/ranges, same-origin persistence, builds, Go suite, and real-browser round-trip verified without effective MMA2 writes).
8. `sim-006-save-apply-routing.md` — Save & Apply routing. **CURRENT — execute first.**
9. `sim-007-runtime-status.md` — runtime status surface.

A later task does not authorize skipping an incomplete dependency.
