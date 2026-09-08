# Handoff

## Status

ACTIVE — SIM-018 is complete; continue Operation CWAL with SIM-019.

## Current Active Work

Execution order:

1. `workflow/active_work/sim-018-decide-local-ui-runtime-boundary.md` — **COMPLETED 2026-09-08**. Selected authenticated OS.js session WebSocket -> allowlisted OS.js relay -> Unix-domain socket -> independently supervised Go runtime; canonical document remains the Go Store; no Simulator HTTP/TCP API; MMA2 remains independently managed and restart-only.
2. `workflow/active_work/sim-019-host-simulator-runtime-locally.md` — implement the long-lived local runtime owner and approved command/status contract.
3. `workflow/active_work/sim-020-connect-save-apply-to-runtime.md` — connect the OS.js Save & Apply action to the real `ApplyRouter` transaction.
4. `workflow/active_work/sim-021-restore-truthful-runtime-status.md` — show MMA2, Simulator, Raw Ingest, apply, and per-FC timing status.
5. `workflow/active_work/sim-022-visible-end-to-end-verification.md` — verify the complete visible workflow against real MMA2 and Modbus reads.

All five tasks were explicitly promoted by the user on 2026-09-08. Complete and checkpoint each task through Operation CWAL before beginning the next.

## Completed Predecessor Work

Prior MMA2 work and SIM-001 through SIM-008 are completed and verified. SIM-009 is superseded by the corrected architecture sequence below.

Archived execution order:
1. `workflow/archive/sim-010-remove-simulator-bridge-api.md` — SIM-010: remove Simulator bridgeand config API. **COMPLETED 2026-09-08** (removed `simulator/cmd/simbridge`, `simulator/bridge.go`, `simulator/bridge_test.go`, OS.js `simulator-bridge` provider + `/api/devices` + `/api/devices/status` routes, `SIMBRIDGE_ADDR`/`SIMULATOR_BRIDGE_ADDR`; removed bridge-served runtime-status UI polling/panel and bridge-reachability failure strings; preserved model/persistence/composition/ownership/scheduler/raw-ingest/UI. `gofmt -l .`, `go vet ./...`, `go test -count=1 ./...`, `node --check`, `npm run build:local-packages`, `npm run build` all pass.
2. `workflow/archive/sim-011-remove-simulator-mma2-lifecycle-ownership.md` — SIM-011: remove Simulator ownership of MMA2 lifecycle. **COMPLETED 2026-09-08** (deleted child-process lifecycle/start-stop-replace/rollback code and tests; structural/boot paths now compose only; Simulator stop leaves independent MMA2 untouched; schedulers remain unarmed pending SIM-015; affected formatting, vet, tests, and static checks pass).
3. `workflow/archive/sim-012-local-simulator-config-path.md` — SIM-012: establish local Simulator configuration path without a config HTTP API. **COMPLETED 2026-09-08** (Simulator definitions persist through the existing server-backed OS.js settings service under a dedicated namespace; exact all-field reload round-trip plus add/duplicate/delete/discard/save verified in the rebuilt UI; no MMA2 config/lifecycle artifact created; builds and affected tests pass).
4. `workflow/archive/sim-013-compose-shared-mma2-config.md` — SIM-013: safely compose Simulator-owned entries into shared MMA2 config. **COMPLETED 2026-09-08** (enabled-only translation; foreign fields/reservations preserved; complete candidate validated through MMA2's authoritative validator before safe commit; collision/invalid candidate leave config and owners byte-unchanged; tests pass).
5. `workflow/archive/sim-014-mma2-restart-only-control.md` — SIM-014: add MMA2 RESTART-only control. **COMPLETED 2026-09-08** (new `simulator/restart.go` restart-request artifact beside the shared config(, written only after a successful `ComposeDocument` commit,, exactly one RESTART-plus-ready-wait per apply,,cleared after confirmed readiness,,left pending when restart unconfirmed;`WaitMMA2Ready` dials every composed listener port until accept or timeout; ApplyStructural reports restart/readiness failure truthfully without claiming apply success; four restart tests pass;live restart reload proven( MMA2 rebuilt with port 5020→5021 config swap;;new listener accepts after restart(; no other lifecycle operation added; no general control API.
6. `workflow/archive/sim-015-run-simulation-after-mma2-apply.md` — SIM-015: run schedulesand Raw Ingest after successful MMA2 apply/readiness. **COMPLETED 2026-09-08** (ApplyRouter arms enabled schedules only after the full Save & Apply chain( shared-config commit,, MMA2 RESTART,, readiness,,persistence( succeeded;`ArmSchedules` replaces prior scheduleswith the latest persisted state;disabled devices never arm;restart/readiness failure returns apply error before persistence,, arming nothing;,raw-ingest faults surfaced truthfully via runtime status;two new arming tests pass;existing scheduler+raw-ingest suites stay green(.
7. `workflow/archive/sim-016-restore-enabled-simulations-on-boot.md` — SIM-016: restore enabled simulations after boot while MMA2 auto-starts independently. **COMPLETED 2026-09-08** (boot router loads persisted definitions, ownership-safe compose, waits for independently auto-started MMA2 via `WaitMMA2Ready`, then arms enabled schedules through the SIM-015 arming gate; ordinary boot neither restarts MMA2 nor writes a restart request; if MMA2 never becomes ready, the router survives alive with schedules unarmed and truthful STOPPED/ERROR runtime status; two new boot-restore tests prove armed-enabled-resume-without-restart-request and unavailable-MMA2-no-arm-no-restart;`gofmt -l`,`go vet ./...`,`go test -count=1 ./...` all pass(.
8. `workflow/archive/sim-017-end-to-end-simulator-mma2-verification.md` — SIM-017: end-to-end architecture verification. **COMPLETED 2026-09-08** (verification-only real MMA2 harness proves independent boot, restore, changing Raw Ingest schedules, real FC1–FC4 reads, exactly one restart after valid apply, no restart/config damage after rejection, foreign preservation, and automatic reboot resume; race-enabled capstone and full Simulator suite pass).

The promoted SIM-010 → SIM-017 sequence is complete and archived.

## Architecture Boundary

MMA2 owns its own boot/start lifecycle and auto-starts on system boot. The Simulator does not own MMA2 and may not START, STOP, SPAWN, KILL, or REPLACE it. The only MMA2 lifecycle/control action the Simulator may request is RESTART after successfully committing a valid shared MMA2 configuration change. Simulator-generated values enter MMA2 through Raw Ingest.

There is no `simbridge` service and no Simulator configuration API in the target architecture.

## Execution Rule

JR executes only work present in `workflow/active_work/` and reflected here. Execute the promoted sequence strictly in dependency order.

Do not inspect Planning to choose or widen work.

## Promotion Rule

Whenever the Active Workload changes, update this file in the same promotion/change so execution context cannot drift.
