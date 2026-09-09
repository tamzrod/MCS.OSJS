# SIM-021A — Define Operator Runtime Status Semantics

Status: PROMOTED — 2026-09-08 by human; ordered SIM-021A → SIM-021E.

## Primary outcome

Make `SchedulerApplier.RuntimeStatus` produce truthful evidence for exactly two operator-facing states: MMA2 and the selected Simulator device.

## Scope

- Keep Raw Ingest as an internal protocol/error source, not a separately displayed operator status.
- Define MMA2 state from actual listener readiness evidence.
- Define Simulator state from device enablement, scheduler availability/readiness, MMA2 readiness, and Raw Ingest success/failure evidence.
- Track enough successful Raw Ingest evidence so Simulator cannot claim `RUNNING` before at least one generated batch has been accepted.
- Preserve the last Raw Ingest error string for diagnostics until a later successful ingest clears it.
- Use the compact state vocabulary below:
  - MMA2: `RUNNING`, `WAITING`, `STOPPED`, `ERROR` only when those states can be distinguished truthfully from existing runtime evidence.
  - Simulator: `RUNNING`, `WAITING`, `STOPPED`, `ERROR`.

## Required semantics

- Disabled device -> Simulator `STOPPED`.
- Enabled device with scheduler not yet armed / MMA2 not yet ready -> Simulator `WAITING`, never `RUNNING`.
- Enabled device with MMA2 ready but no successful Raw Ingest yet -> Simulator `WAITING`.
- Enabled device after at least one successful Raw Ingest -> Simulator `RUNNING`.
- Raw Ingest failure -> Simulator `ERROR` and retain the underlying error text.
- Later successful Raw Ingest -> clear the error and return to `RUNNING` when all other readiness conditions are satisfied.
- MMA2 probe failure must never leave Simulator shown as `RUNNING`.

## Non-scope

- No OS.js UI changes.
- No polling.
- No new transport or API.
- No FC Last/Next display work.
- No register-value viewer, charts, or history.

## Acceptance criteria

1. Runtime tests prove an enabled device cannot become `RUNNING` before successful Raw Ingest evidence exists.
2. Runtime tests prove disabled, waiting, running, and Raw-Ingest-error Simulator states.
3. Runtime tests prove MMA2 unavailability prevents a `RUNNING` Simulator result.
4. Raw Ingest failure text survives until a later successful ingest clears it.

## Verification

Run the focused Go tests for `SchedulerApplier.RuntimeStatus`, then the Simulator Go test suite.

## Dependencies

- SIM-015 runtime arming gate.
- SIM-019 version-1 runtime `status` operation.
- SIM-020 canonical runtime/apply path.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. One bounded Go runtime-status behavior change.
