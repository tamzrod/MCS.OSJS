# Handoff

## Status

ACTIVE

## Current Active Work

Active Work is the ordered set of task files under `workflow/active_work/` listed below.

Execution order:
1. `workflow/active_work/mma2-basic-install-test.md` — existing active MMA2 work; preserve its internal execution order and completion state as repository truth. **COMPLETED 2026-09-06** (both sub-tasks done+verified; evidence in task file).
2. `workflow/active_work/sim-001-simulator-device-config.md` — SIM-001: establish simulator device configuration model. **COMPLETED 2026-09-06** (round-trip + invalid-reject verified; evidence in task file). **Current work is now SIM-002** (item 3 below).
3. `workflow/active_work/sim-002-mma2-activation.md` — SIM-002: validate and activate simulator MMA2 parameters.
4. `workflow/active_work/sim-003-random-runtime.md` — SIM-003: implement per-FC random runtime scheduler.
5. `workflow/active_work/sim-004-raw-ingest.md` — SIM-004: send simulator values through MMA2 raw ingest.
6. `workflow/active_work/sim-005-osjs-window.md` — SIM-005: build the approved OS.js Modbus Simulator window.
7. `workflow/active_work/sim-006-save-apply-routing.md` — SIM-006: route Save & Apply to the correct parameter consumer.
8. `workflow/active_work/sim-007-runtime-status.md` — SIM-007: show simulator runtime status.

JR must execute and verify simulator tasks in SIM-001 → SIM-007 dependency order. A later task does not authorize skipping an incomplete dependency.

## Execution Rule

JR executes only work present in `workflow/active_work/` and reflected here.

Do not inspect Planning to choose or widen work.

## Promotion Rule

Whenever the Active Workload changes, update this file in the same promotion/change so execution context cannot drift.
