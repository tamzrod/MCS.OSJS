# REP-006 — Replicator Poll Loop

Status: ACTIVE
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

Focused runtime tests with a short test interval and controlled cycle success/failure, proving repeated invocation, no overlap, truthful error state, and clean cancellation.

## Dependencies

- REP-003 Replicator configuration.
- REP-005 single-range replication cycle.

## Sizing

Implementation surface 1; environment uncertainty 0; behavioral surface 1; verification surface 1; decision/recovery surface 1. Total: 3 — good JR task.
