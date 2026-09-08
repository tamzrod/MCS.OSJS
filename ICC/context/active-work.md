# Active Work — Empty; Completed Program Archived

Baseline commit: 6069fef
Working tree: clean
Audited Overlay:(none(; files once audited as overlay (`simulator/device.go`,`simulator/store.go`,`simulator/store_document_test.go`( are now committed in HEADand the SIM-010 boundary files (`simulator/bridge.go`,`simulator/bridge_test.go`,`simulator/cmd/simbridge/main.go`,`OSJS/src/server/providers/simulator-bridge.js`( were deleted at `0b356da`,;none remain in the working tree.
Source dependencies: handoff.md, workflow/active_work/README.md, workflow/archive/*, MMA2/testdata/smoke-test.yaml, simulator/*, OSJS/src/packages/ModbusSimulator/
Parent: L0-project
Zoom In: simulator-device-config
Zoom Out: L0-project

## Execution Order (handoff)

No current Active Work microtask remains. The records below are archived completion evidence and do not authorize execution.

1. `mma2-basic-install-test.md` — **COMPLETED** 2026-09-06 (MMA2-001 import+build; MMA2-002 smoke test).
2. `sim-001-simulator-device-config.md` — **COMPLETED** 2026-09-06. Simulator-owned two-domain persist under `$OSJS_DATA_DIR/config/simulator/devices.yaml`. See Zoom In `simulator-device-config`.
3. `sim-002a-mma2-config-ownership.md` — **COMPLETED 2026-09-07**. Compose+ownership protection (`(port,unit_id)` reservation keys, YAML `owner` registry at `$OSJS_DATA_DIR/config/mma2/owners.yaml`, `ErrReservationOwnedByOther` on foreign collision)。 Six tests in `simulator/compose_test.go` verified. See Zoom In `simulator-device-config`.
4. `sim-002b-mma2-lifecycle-activation.md` — **COMPLETED 2026-09-07**. Live MMA2 proven through the composed effective config (process healthy; Modbus-TCP (5020, 1) coils/discrete 64 + holding/input 100 read/write; raw PDU hreg0=4321). See Zoom In `simulator-device-config`.
5. `sim-003-random-runtime.md` — per-FC random runtime scheduler.**COMPLETED 2026-09-07** (scheduler.go+scheduler_test.go;both tests pass;timing-only UpdateTiming never restarts MMA2(. See Zoom In `simulator-device-config`.
6. `sim-004-raw-ingest.md` — raw ingest into MMA2 memory.**COMPLETED 2026-09-07** (`simulator/raw_ingest.go` + `simulator/raw_ingest_test.go`; MMA2 raw-ingest v1 client; bits-LSB-first/regs big-endian frames; Send( exclusively via raw ingest(; 4 tests incl real-TCP round-trip;gofmt/vet clean; `go test -count=1 ./...` ok. See Zoom In `simulator-device-config`.
7. `sim-005-osjs-window.md` — OS.js Modbus Simulator window. **COMPLETED 2026-09-08** at pushed commit `28d27f9` (two-pane editor, approved fields/actions/ranges, same-origin persistence, builds, Go suite, and real-browser round-trip verified without effective MMA2 writes).
8. `sim-006-save-apply-routing.md` — Save & Apply routing. **COMPLETED 2026-09-08** at pushed commit `9532c31` (classification, correct consumers, rollback-on-rejection, bridge/UI status, tests, and three browser saves verified).
9. `sim-007-runtime-status.md` — runtime status surface. **COMPLETED 2026-09-08** at pushed commit `ceee393` (device/MMA2/raw-ingest status, exact point total, FC1-FC4 timing, tests, builds, and browser verification complete).
10. `sim-008-managed-mma2-lifecycle.md` — **COMPLETED 2026-09-08** at pushed commit `7e63460`. Managed child-process activation, readiness, replacement, rollback,and real MMA2 bind-failure recovery verified.

11. `sim-009-serve-simulator-through-mma2.md` — **SUPERSEDED 2026-09-08** by SIM-010…SIM-017; its premise (Simulator owns+activates a managed MMA2 child( was an incorrect architecture assumption. Current authoritative boundary: MMA2 auto-starts on boot; Simulator must not START/STOP/SPAWN/KILL/REPLACE it; only RESTART (SIM-014( after a committed valid shared-config change; values enter via Raw Ingest; no `simbridge` service and no Simulator config API. SIM-009's committed apply.go ready-gate delta (pre-ready schedulers unarmed; truthful STOPPED/ERROR status( remains in HEAD at `83a5b69`.

SIM-001…SIM-008 are complete+verified; SIM-009 was superseded. SIM-010 through SIM-016 are complete. SIM-017 — End-to-End Simulator+MMA2 Verification — **COMPLETED 2026-09-08** at pushed commit `57c8714`: a verification-only real-MMA2 harness proved independent boot, restore, changing schedules through Raw Ingest, FC1–FC4 reads, one restart after valid apply, no restart/config damage after rejection, foreign preservation, and reboot resume. The promoted sequence is complete; no current Active Work microtask remains.
