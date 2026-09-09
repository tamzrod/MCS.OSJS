# Active Work — Simulator Runtime Integration

Baseline commit: a17191c; working tree: dirty (planning/microtask/mma2-basic-install-test.md and planning/microtask/osjs-base-webapp-init.md deleted uncommitted by human, outside this branch.
Audited Overlay: handoff.md; five sim-021 planning microtasks promoted into workflow/active_work as SIM-021A..SIM-021E (ordered; ICC nodes refreshed herein.
Source dependencies: handoff.md, workflow/active_work/*, workflow/archive/*, docs/SIMULATOR_RUNTIME_INTEGRATION.md, simulator/*, OSJS/src/packages/ModbusSimulator/, planning/Brainstorm/modbus-simulator-status-display.md
Parent: L0-project
Zoom In: simulator-device-config
Zoom Out: L0-project

## Current Execution Order
Current ordered execution sequence (promoted into `workflow/active_work/` as SIM-021A..SIM-021E on 2026-09-08;next up: SIM-021A(; JR executes per Operation CWAL until Active Work drains.

1. SIM-018 — **COMPLETED 2026-09-08**. Chose authenticated OS.js session WebSocket -> allowlisted OS.js relay -> Unix-domain socket -> independently supervised Go runtime. The Go Store is canonical;there is no Simulator HTTP/TCP API. Archived.
2. SIM-019 — **COMPLETED 2026-09-09**. Long-lived Go runtime, Unix RPC, OS.js WebSocket relay,and deployment sidecar verified without a Simulator HTTP/TCP endpoint. Archived.
3. SIM-020 — **COMPLETED 2026-09-09**. UI canonical loadand real apply transaction verified for structural, timing-only, rejected, Discard,and reload paths. Archived.

4. SIM-021 — **PROMOTED 2026-09-08**. Human-authorized as five ordered active microtasks: SIM-021A define operator status semantics; SIM-021B lock status RPC contract; SIM-021C poll selected-device status; SIM-021D render compact status row; SIM-021E surface status errors. Records now live in `workflow/active_work/`.
5. SIM-022 — **RETIRED 2026-09-08**（not completed; human clearance; remains un-promoted. Record: `workflow/archive/sim-022-visible-end-to-end-verification.md`.

The records below are archived predecessor evidence and do not authorize execution.

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

SIM-001…SIM-008 are complete+verified; SIM-009 was superseded. SIM-010 through SIM-016 are complete. SIM-017 — End-to-End Simulator+MMA2 Verification — **COMPLETED 2026-09-08** at pushed commit `57c8714`: a verification-only real-MMA2 harness proved independent boot, restore, changing schedules through Raw Ingest, FC1–FC4 reads, one restart after valid apply, no restart/config damage after rejection, foreign preservation, and reboot resume.
