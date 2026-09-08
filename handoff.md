# Handoff

## Status

ACTIVE — SIM-009 current.

## Current Active Work

Active Work is the ordered set of task files under `workflow/active_work/` listed below.

Execution order:
1. `workflow/active_work/mma2-basic-install-test.md` — existing active MMA2 work; preserve its internal execution order and completion state as repository truth. **COMPLETED 2026-09-06** (both sub-tasks done+verified; evidence in task file).
2. `workflow/active_work/sim-001-simulator-device-config.md` — SIM-001: establish simulator device configuration model. **COMPLETED 2026-09-06** (round-trip + invalid-reject verified; evidence in task file).
3. `workflow/active_work/sim-002a-mma2-config-ownership.md` — SIM-002A: compose simulator MMA2 configuration with persistent YAML ownership protection. **COMPLETED 2026-09-07** (six ownership-enforcement tests verified; evidence in task file).
4. `workflow/active_work/sim-002b-mma2-lifecycle-activation.md` — SIM-002B: activate the already-valid effective MMA2 configuration and verify live Modbus exposure. **COMPLETED 2026-09-07** (live Modbus-TCP (5020,1) proven: coils →  64, discrete  64, holding+input  100, writes/read-backs round-trip, raw FC3 PDU hreg0=4321; MMA2 process healthy; evidence in task file).
5. `workflow/active_work/sim-003-random-runtime.md` — SIM-003: implement per-FC random runtime scheduler. **COMPLETED 2026-09-07** (concurrent FC1-FC4 cadence + timing-only interval change without restart verified; evidence in task file).
6. `workflow/active_work/sim-004-raw-ingest.md` — SIM-004: send simulator values through MMA2 raw ingest. **COMPLETED 2026-09-07** (simulator/raw_ingest.go + simulator/raw_ingest_test.go; raw-ingest v1 client drives FC1-FC4 values into simulator-owned MMA2 memory exclusively via raw ingest; 4 tests cover bit/reg frame encoding, real-TCP send round-trip with OK/reject response handling, and unconfigured-FC rejection;gofmt/vet clean; go test -count=1 ./... -> ok; evidence in task file).
7. `workflow/active_work/sim-005-osjs-window.md` — SIM-005: build the approved OS.js Modbus Simulator window. **COMPLETED 2026-09-08** (corrupt frontend replaced; two-pane device list/editor, all approved fields/actions/ranges, same-origin persistence, production builds, Go suite, and real-browser add/edit/save/duplicate/delete/discard round-trip verified without touching effective MMA2 configuration; evidence in task file).
8. `workflow/active_work/sim-006-save-apply-routing.md` — SIM-006: route Save & Apply to the correct parameter consumer. **COMPLETED 2026-09-08** (structural/timing/no-change classifier, SIM-002/SIM-003 consumers, bridge/UI result surface, rollback-on-rejection tests, and exactly three browser saves verified; evidence in task file).
9. `workflow/active_work/sim-007-runtime-status.md` — SIM-007: show simulator runtime status. **COMPLETED 2026-09-08** (selected-device/MMA2/raw-ingest state, exact point total, FC1-FC4 last/next timing, backend tests, builds, and browser verification complete; evidence in task file).
10. `workflow/active_work/sim-008-managed-mma2-lifecycle.md` — SIM-008: activate ownership-validated effective MMA2 configuration through one managed lifecycle. **COMPLETED 2026-09-08** at pushed commit `7e63460`.
11. `workflow/active_work/sim-009-serve-simulator-through-mma2.md` — SIM-009: serve generated simulator values through managed MMA2 and restore enabled devices at startup. **CURRENT**.

JR must execute and verify simulator tasks in SIM-001 → SIM-002A → SIM-002B → SIM-003 → SIM-004 → SIM-005 → SIM-006 → SIM-007 → SIM-008 → SIM-009 dependency order. A later task does not authorize skipping an incomplete dependency.

SIM-002A establishes the shared MMA2 ownership boundary before lifecycle activation proceeds. `(port, unit_id)` is the unique reservation key shared by Simulator and Replicator. Each persisted effective MMA2 reservation carries a machine-readable YAML `owner` entry. Ownership is first-come-first-save; a program may modify or delete only reservations it owns, and attempts to overwrite or remove another program's reservation must be rejected.

## Execution Rule

JR executes only work present in `workflow/active_work/` and reflected here.

Do not inspect Planning to choose or widen work.

## Promotion Rule

Whenever the Active Workload changes, update this file in the same promotion/change so execution context cannot drift.
