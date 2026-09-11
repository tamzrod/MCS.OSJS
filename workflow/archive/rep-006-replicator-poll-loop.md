# REP-006 — Replicator Poll Loop

Status: COMPLETED 2026-09-11 — JR verification PASS accepted.
Previous: REP-005
Next: REP-007

## Primary Outcome

Run the proven single-range replication cycle repeatedly at the configured polling interval with truthful running/error state and clean stop behavior.

## Scope

- Add a bounded runtime loop around REP-005 using the persisted polling interval.
- Start only after configuration is valid and the Replicator-owned MMA2 destination is ready.
- Execute one replication cycle per tick without overlapping cycles.
- Preserve truthful error state when a cycle fails and continue on the next deterministic tick.
- Support clean stop/cancel without leaving another cycle running.

## Non-Scope

- No advanced retry/backoff policy.
- No multiple workers/devices/ranges.
- No UI controls or status display.
- No metrics/history database.

## Acceptance Criteria

1. The runtime repeatedly invokes the single-range cycle at the configured interval without overlapping executions.
2. Stop/cancel terminates the loop cleanly.
3. Runtime state distinguishes running success from source/destination error rather than reporting success after a failed cycle.

## Verification

JR verified commit `bffee36e5660e15875dbe912be0e8d56d0c93d0c`:

- clean working tree before test;
- `gofmt -l .` produced no output;
- repeated serial execution / no-overlap test passed;
- error-state / continuation-policy test passed;
- clean cancellation test passed;
- full `go test -count=1 ./...` passed;
- `go vet ./...` passed with no diagnostics.

The trailing missing `)` in JR's sandbox-tooling prose note was corrected by the coding agent during advancement and did not affect test evidence.

## Dependencies

- REP-003 Replicator configuration.
- REP-005 single-range replication cycle.
