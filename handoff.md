# Handoff

## Status

ACTIVE — SIM-014 current. SIM-010 through SIM-017 are human-promoted for sequential execution.

## Current Active Work

Prior MMA2 work and SIM-001 through SIM-008 are completed and verified. SIM-009 is superseded by the corrected architecture sequence below.

Execution order:
1. `workflow/active_work/sim-010-remove-simulator-bridge-api.md` — SIM-010: remove Simulator bridgeand config API. **COMPLETED 2026-09-08** (removed `simulator/cmd/simbridge`, `simulator/bridge.go`, `simulator/bridge_test.go`, OS.js `simulator-bridge` provider + `/api/devices` + `/api/devices/status` routes, `SIMBRIDGE_ADDR`/`SIMULATOR_BRIDGE_ADDR`; removed bridge-served runtime-status UI polling/panel and bridge-reachability failure strings; preserved model/persistence/composition/ownership/scheduler/raw-ingest/UI. `gofmt -l .`, `go vet ./...`, `go test -count=1 ./...`, `node --check`, `npm run build:local-packages`, `npm run build` all pass.
2. `workflow/active_work/sim-011-remove-simulator-mma2-lifecycle-ownership.md` — SIM-011: remove Simulator ownership of MMA2 lifecycle. **COMPLETED 2026-09-08** (deleted child-process lifecycle/start-stop-replace/rollback code and tests; structural/boot paths now compose only; Simulator stop leaves independent MMA2 untouched; schedulers remain unarmed pending SIM-015; affected formatting, vet, tests, and static checks pass).
3. `workflow/active_work/sim-012-local-simulator-config-path.md` — SIM-012: establish local Simulator configuration path without a config HTTP API. **COMPLETED 2026-09-08** (Simulator definitions persist through the existing server-backed OS.js settings service under a dedicated namespace; exact all-field reload round-trip plus add/duplicate/delete/discard/save verified in the rebuilt UI; no MMA2 config/lifecycle artifact created; builds and affected tests pass).
4. `workflow/active_work/sim-013-compose-shared-mma2-config.md` — SIM-013: safely compose Simulator-owned entries into shared MMA2 config. **COMPLETED 2026-09-08** (enabled-only translation; foreign fields/reservations preserved; complete candidate validated through MMA2's authoritative validator before safe commit; collision/invalid candidate leave config and owners byte-unchanged; tests pass).
5. `workflow/active_work/sim-014-mma2-restart-only-control.md` — SIM-014: add MMA2 RESTART-only control. **CURRENT**.
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
