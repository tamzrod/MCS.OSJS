# REP-004 — Modbus Source Reader

Status: COMPLETED 2026-09-11 — JR verification PASS accepted.
Previous: REP-003
Next: REP-005

## Primary Outcome

Implement the smallest Replicator-side Modbus TCP source reader needed to read one configured register range from the working Simulator endpoint.

## Scope

- Add a Replicator Modbus TCP client/read function for the first supported register area.
- Use the persisted source host, port, unit ID, start address, and count from REP-003.
- Return the acquired values in a small producer-neutral value shape suitable for the later MMA2 write step.
- Handle connection/read failure as an error without modifying destination MMA2 memory.
- Keep connection lifecycle simple and deterministic for the first milestone.

## Non-Scope

- No continuous polling loop.
- No MMA2 destination write.
- No retries/backoff beyond the minimum needed by the chosen client library.
- No multiple source devices.
- No scaling, byte swapping, or transformations.
- No UI.

## Acceptance Criteria

1. Replicator can connect to the working Simulator Modbus TCP endpoint using config values.
2. A configured small register range is read with values matching the Simulator source.
3. Connection/read errors are returned cleanly and do not trigger any destination write behavior.

## Verification

JR verified commit `49417dffd3a07cedc2500d8ecbf318aa65ba36a9` with a clean working tree. `gofmt -l .` returned no files; persisted-config integration passed; FC3/FC4 protocol tests passed; error-path tests passed; full `go test -count=1 ./...` passed; and `go vet ./...` returned no diagnostics.

## Dependencies

- REP-003 Replicator configuration.
- Working Simulator endpoint for integration verification.

## Sizing

Implementation surface 1; environment uncertainty 1; behavioral surface 1; verification surface 1; decision/recovery surface 0. Total: 4 — acceptable only because the task is one tightly coupled read path with one deterministic protocol verification workflow.
