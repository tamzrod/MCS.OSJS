# SIM-010 — Remove Simulator Bridge and Config API

Status: ACTIVE — promoted for sequential execution.

## Primary outcome

Remove the incorrectly introduced HTTP bridge/config-API boundary from the local Modbus Simulator architecture.

## Architecture contract

The Simulator is a local MCS.OSJS application. There is no `simbridge` service and no Simulator configuration API. MMA2 is an independent appliance component that auto-starts on system boot.

## Scope

- Remove `simulator/cmd/simbridge`.
- Remove the OS.js `simulator-bridge` provider and Simulator `/api/devices` and `/api/devices/status` routes.
- Remove `SIMBRIDGE_ADDR`, `SIMULATOR_BRIDGE_ADDR`, and bridge-only assumptions/tests/docs.
- Remove UI failure handling whose only purpose is reporting bridge reachability.
- Preserve the Simulator domain model, persistence capability, MMA2 config composition/ownership logic, scheduler, raw-ingest capability, and UI for later rewiring.

## Non-scope

- Do not redesign local Simulator persistence in this task.
- Do not change MMA2 shared-config semantics.
- Do not implement MMA2 restart control.
- Do not start/stop/spawn MMA2.

## Acceptance criteria

1. Repository builds/tests without the simulator bridge executable/provider.
2. No Simulator-specific configuration HTTP endpoint remains.
3. No runtime error can report `simulator bridge not running` because that architectural component no longer exists.
4. Reusable Simulator model/config/scheduler/raw-ingest code remains available for subsequent tasks.

## Verification

Run affected Go and OS.js tests/builds and search the repository for bridge routes/environment variables to prove the removed boundary is no longer referenced by runtime code.

## Dependencies

None. This is the first architecture-correction task and supersedes the bridge assumption introduced by prior SIM work.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. Good JR-sized corrective task.