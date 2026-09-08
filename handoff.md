# Handoff

## Status

ACTIVE — SIM-010 current. SIM-010 through SIM-017 are human-promoted for sequential execution.

## Current Active Work

Prior MMA2 work and SIM-001 through SIM-008 are completed and verified. SIM-009 is superseded by the corrected architecture sequence below.

Execution order:
1. `workflow/active_work/sim-010-remove-simulator-bridge-api.md` — SIM-010: remove Simulator bridge and config API. **CURRENT**.
2. `workflow/active_work/sim-011-remove-simulator-mma2-lifecycle-ownership.md` — SIM-011: remove Simulator ownership of MMA2 lifecycle.
3. `workflow/active_work/sim-012-local-simulator-config-path.md` — SIM-012: establish local Simulator configuration path without a config HTTP API.
4. `workflow/active_work/sim-013-compose-shared-mma2-config.md` — SIM-013: safely compose Simulator-owned entries into shared MMA2 config.
5. `workflow/active_work/sim-014-mma2-restart-only-control.md` — SIM-014: add MMA2 RESTART-only control.
6. `workflow/active_work/sim-015-run-simulation-after-mma2-apply.md` — SIM-015: run schedules and Raw Ingest after successful MMA2 apply/readiness.
7. `workflow/active_work/sim-016-restore-enabled-simulations-on-boot.md` — SIM-016: restore enabled simulations after boot while MMA2 auto-starts independently.
8. `workflow/active_work/sim-017-end-to-end-simulator-mma2-verification.md` — SIM-017: end-to-end architecture verification.

JR must execute and verify SIM-010 → SIM-011 → SIM-012 → SIM-013 → SIM-014 → SIM-015 → SIM-016 → SIM-017 in order. A later task does not authorize skipping an incomplete dependency.

## Architecture Boundary

MMA2 owns its own boot/start lifecycle and auto-starts on system boot. The Simulator does not own MMA2 and may not START, STOP, SPAWN, KILL, or REPLACE it. The only MMA2 lifecycle/control action the Simulator may request is RESTART after successfully committing a valid shared MMA2 configuration change. Simulator-generated values enter MMA2 through Raw Ingest.

There is no `simbridge` service and no Simulator configuration API in the target architecture.

## Execution Rule

JR executes only work present in `workflow/active_work/` and reflected here. Execute the promoted sequence strictly in dependency order.

Do not inspect Planning to choose or widen work.

## Promotion Rule

Whenever the Active Workload changes, update this file in the same promotion/change so execution context cannot drift.
