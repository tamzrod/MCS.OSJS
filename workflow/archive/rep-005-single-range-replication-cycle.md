# REP-005 — Single-Range Replication Cycle

Status: COMPLETED 2026-09-11 — JR verification PASS accepted.
Previous: REP-004
Next: REP-006

## Primary Outcome

Implement one Replicator cycle that reads a configured register range from the source Modbus endpoint and writes the unchanged values into the configured Replicator-owned MMA2 destination.

## Scope

- Compose/reserve the Replicator destination through the shared MMA2 reservation composer with producer identity `replicator`.
- Read one configured source range through the REP-004 source reader.
- Require source count and destination count to match for the first 1:1 mapping milestone.
- Write the acquired values through the shared MMA2 raw-ingest client.
- Return clear success/error status for one invocation of the replication cycle.
- Never overwrite a Simulator-owned or other foreign `(port, unit_id)` reservation.

## Acceptance Criteria

1. A Replicator destination reservation is created/updated only under owner `replicator`, preserving foreign reservations.
2. One invocation reads the configured source values and writes the same values to the configured destination addresses.
3. Any source-read, ownership, validation, or raw-ingest failure returns an error without falsely reporting a completed replication cycle.

## Verification Evidence

JR tested commit `7eda5d964c795a0f2576e163ab5a599f7359fee7` with a clean working tree.

- `gofmt -l .` produced no output.
- `TestRunOnceCopiesConfiguredRegisters` passed.
- `TestRunOnceRejectsForeignOwnedDestination` passed.
- `TestValidateCycleMappingRejectsCrossArea` passed.
- Full `go test -count=1 ./...` passed.
- `go vet ./...` passed with no diagnostics.

The trailing missing parenthesis in JR's prose-only toolchain note was corrected during the next coding-agent handoff update; it did not affect test evidence.
