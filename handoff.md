# Handoff

## Current

SIM-024 — Wire MMA2 Restart Request to the Deployed Runtime (`ACTIVE`).

Observed failure: Simulator structural Save & Apply commits the shared MMA2 config and writes `restart-request.yaml`, then times out waiting for the new listener (for example `127.0.0.1:5020`) because the deployed stack has no production restart-request consumer and currently defines no MMA2 runtime service.

## Authorized Sequence

SIM-024

## Continuation

Operation CWAL should execute only `workflow/active_work/sim-024-wire-mma2-restart-request-to-runtime.md`.

The required fix belongs at the independent MMA2 appliance/deployment boundary: deploy MMA2 against `/data/config/mma2/config.yaml` and consume the existing restart request exactly once per request. Do not restore Simulator-owned MMA2 process lifecycle control.

Completion requires repository-native Go and Docker/deployment gates plus a real deployed structural Save & Apply proving MMA2 restarts/reloads once and the new listener becomes reachable. If Docker/deployment verification is unavailable, SIM-024 remains ACTIVE.

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.
