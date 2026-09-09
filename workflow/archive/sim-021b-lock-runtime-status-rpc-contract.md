# SIM-021B — Lock the Runtime Status RPC Contract

Status: COMPLETED 2026-09-09 — version-1 status response contract verified.
Previous: SIM-021A
Next: SIM-021C

## Primary outcome

Make the existing version-1 `status` RPC carry the SIM-021A MMA2 + Simulator truth unchanged through the existing Unix-socket/runtime-service boundary.

## Scope

- Keep the existing authenticated OS.js WebSocket -> allowlisted relay -> Unix-domain socket path.
- Keep operation name `status` and selected-device request payload by device name.
- Ensure the returned status object exposes the operator-facing MMA2 state, Simulator state, and diagnostic error text required by SIM-021A.
- Remove any need for the UI to infer runtime health from `Enabled`, saved configuration, FC timing, or transport success alone.
- Add/adjust focused runtime-server contract tests for the returned status payload.

## Non-scope

- No new HTTP route.
- No new TCP listener.
- No new WebSocket message family.
- No MMA2 lifecycle operation.
- No UI polling or rendering.
- No FC Last/Next presentation.

## Acceptance criteria

1. A version-1 `status` request for a valid selected device returns the SIM-021A MMA2 + Simulator states without client-side health inference.
2. Diagnostic Raw Ingest failure text remains available in the response for error messaging but is not a third operator-facing status.
3. Invalid/missing device-name requests retain stable runtime errors.
4. Existing `load` and `apply` protocol behavior is unchanged.

## Verification

Run focused `RuntimeService` status contract tests, then the Simulator Go test suite.

## Completion evidence

- Focused contract tests serialize the version-1 response and prove `mma2_status`, `device_status`, and `raw_ingest_error` carry SIM-021A truth unchanged.
- Missing names and malformed payloads retain the stable `INVALID_REQUEST` response; an unknown device retains its runtime error text.
- Existing combined load/apply/status and idempotency coverage still passes unchanged.
- Focused RuntimeService tests, `go test -count=1 ./...`, and `go vet ./...` pass in `simulator/`.

## Dependencies

- SIM-021A.
- SIM-019 local runtime boundary.

## Sizing

Implementation 1, environment 0, behavioral 0, verification 1, decision/recovery 0 = 2. One bounded protocol-contract surface using the already-established transport.
