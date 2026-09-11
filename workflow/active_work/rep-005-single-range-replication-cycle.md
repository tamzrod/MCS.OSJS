# REP-005 — Single-Range Replication Cycle

Status: ACTIVE
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

## Non-Scope

- No scheduler or recurring poll loop.
- No transformations, scaling, offsets, endian conversion, or cross-area mapping.
- No multiple source ranges/devices.
- No UI.

## Acceptance Criteria

1. A Replicator destination reservation is created/updated only under owner `replicator`, preserving foreign reservations.
2. One invocation reads the configured source values and writes the same values to the configured destination addresses.
3. Any source-read, ownership, validation, or raw-ingest failure returns an error without falsely reporting a completed replication cycle.

## Verification

Focused tests using controlled source values and a destination raw-ingest endpoint, including a foreign-owner collision case and a successful 1:1 value copy.

## Dependencies

- REP-001 shared MMA2 reservation composer.
- REP-002 shared MMA2 raw-ingest client.
- REP-003 Replicator configuration.
- REP-004 Modbus source reader.

## Sizing

Implementation surface 1; environment uncertainty 0; behavioral surface 2; verification surface 1; decision/recovery surface 0. Total: 4 — acceptable because the read→write cycle is one atomic primary behavior with one bounded verification workflow.
