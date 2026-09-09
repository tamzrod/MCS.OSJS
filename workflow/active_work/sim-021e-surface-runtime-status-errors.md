# SIM-021E — Surface Runtime Status Errors Without Polling Noise

Status: PROMOTED — 2026-09-08 by human; ordered SIM-021A → SIM-021E.

## Primary outcome

Use the existing bottom `.sim-status` message area to explain runtime-status failures without turning the 1-second status refresh into repeated message spam.

## Scope

- When Simulator state enters `ERROR`, show the diagnostic error text returned by the runtime status contract in the existing bottom message bar.
- When the status request itself fails because the local runtime/relay is unavailable, show one truthful unavailable/error message.
- Emit a new bottom-bar message only when the effective error condition changes; repeated identical 1-second poll results must not rewrite/spam the message.
- A later successful runtime state may replace the error with a concise recovery message, but steady-state `RUNNING` remains represented by the compact row rather than persistent bottom-bar text.
- Keep Save & Apply / Discard messages using the same existing transient message area.

## Non-scope

- No separate Raw Ingest status label.
- No toast system.
- No history/log viewer.
- No new runtime error transport.
- No changes to MMA2 lifecycle behavior.

## Acceptance criteria

1. A Raw Ingest failure makes Simulator `ERROR` and the bottom bar shows the returned diagnostic reason.
2. Repeated identical status polls do not repeatedly replace the same bottom-bar message.
3. Runtime/relay unavailability is shown truthfully and does not leave a stale `RUNNING` row.
4. Recovery restores the current compact status row without adding persistent third-state UI.

## Verification

Use a stubbed/erroring status response to verify message deduplication and recovery, then run `node --check` and the ModbusSimulator/local-package build.

## Dependencies

- SIM-021C.
- SIM-021D.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. One bounded client error-presentation behavior.
