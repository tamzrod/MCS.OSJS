# SIM-019 — Host the Simulator Runtime Locally

Status: COMPLETED 2026-09-09 — local runtime host and contract verified.

## Primary outcome

Run one long-lived local Simulator runtime owner that boot-restores the persisted document and exposes the SIM-018-approved local command/status contract to the OS.js shell integration.

## Scope

- Instantiate and retain `NewRuntimeApplyRouter` and its `SchedulerApplier` for the process lifetime.
- Implement only the local transport and lifecycle selected by SIM-018.
- Support canonical document load, `ApplyRouter.Apply`, and `RuntimeStatus(name)` operations.
- Preserve apply errors, restart/readiness errors, and Raw Ingest errors verbatim enough for truthful UI presentation.
- Recover predictably when MMA2 is unavailable at startup and when it later becomes ready through a valid Save & Apply.

## Non-scope

- No browser UI changes.
- No externally reachable Simulator configuration API.
- No direct Modbus writes and no direct MMA2 memory mutation outside Raw Ingest.
- No MMA2 lifecycle operation other than the existing restart request.

## Acceptance criteria

1. Exactly one runtime owner retains schedulers after an initiating UI request ends.
2. An accepted structural apply follows commit, one restart request, readiness, scheduler arming, and Raw Ingest ordering.
3. Timing-only apply updates schedules without an MMA2 config write or restart request.
4. Rejected apply and unavailable-MMA2 cases return truthful errors and never claim RUNNING.

## Verification

Exercise the approved local contract against the real Go runtime with accepted structural, timing-only, rejected, and MMA2-unavailable cases; prove no TCP listener or HTTP route is added for Simulator configuration/status.

## Completion evidence

- `simulator/cmd/modbus-simulator-runtime` owns one long-lived `NewRuntimeApplyRouter`, scheduler set, and Raw Ingest path; graceful shutdown stops Simulator schedulers only.
- `simulator/runtime_server.go` implements bounded length-prefixed version-1 `load`, `apply`, and `status` messages on `$OSJS_DATA_DIR/run/modbus-simulator.sock`, serialized mutations, stable errors, request-id replay protection, stale-socket recovery, and live-owner rejection.
- The ModbusSimulator OS.js package server relays only those allowlisted operations from authenticated application WebSocket messages to the Unix socket. It registers no HTTP route.
- Deployment adds a restart-supervised, host-networked runtime sidecar sharing only the existing data volume; the built image exposes no port.
- Runtime tests cover load/apply/status, request idempotency, Unix framing, single ownership, real `ApplyRouter` structural and timing paths, restart-request clearing after readiness, duplicate rejection, and MMA2-unavailable preservation/reporting.
- `go test -count=1 ./...`, `go vet ./...`, server syntax check, OS.js local-package/full builds, Compose config rendering, and runtime image build pass on 2026-09-09.

## Dependencies

- SIM-018.

## Sizing

Implementation 2, environment 1, behavioral 2, verification 1, decision/recovery 0 = 6. The architectural choice is removed into SIM-018; this task remains the bounded long-lived runtime-host implementation.
