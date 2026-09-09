# REP-007 — Simulator-to-Replicator End-to-End Verification

Status: QUEUED
Previous: REP-006
Next: none

## Primary Outcome

Prove the first complete Replicator milestone by using the working Simulator as the controlled external Modbus source and verifying that changing source values are reproduced in a separate Replicator-owned MMA2 destination.

## Scope

- Run the existing working Simulator with a small deterministic register range.
- Configure Replicator to read that Simulator Modbus endpoint.
- Configure a distinct Replicator-owned MMA2 `(port, unit_id)` destination.
- Start the Replicator poll loop.
- Read the destination through normal Modbus and verify values match the Simulator source.
- Change source values and verify the destination follows on a later poll.
- Confirm the Simulator reservation remains intact and separately owned.

## Non-Scope

- No real field hardware.
- No multiple source devices or ranges.
- No transformations/scaling.
- No OS.js Replicator UI.
- No performance/stress testing.

## Acceptance Criteria

1. Initial source values exposed by Simulator appear unchanged at the Replicator MMA2 destination.
2. After Simulator source values change, the Replicator destination updates to the new values within the expected polling window.
3. Simulator and Replicator use distinct ownership entries/reservations and neither overwrites the other.

## Verification

One explicit end-to-end workflow: START SIMULATOR → READ SOURCE → START REPLICATOR → READ DESTINATION MATCH → CHANGE SOURCE → WAIT POLL → READ DESTINATION MATCH → CHECK OWNERSHIP. Record the exact endpoints, unit IDs, addresses, values, and observed result.

## Dependencies

- Working basic Simulator milestone.
- REP-001 through REP-006 completed and verified.

## Sizing

Implementation surface 0; environment uncertainty 1; behavioral surface 1; verification surface 1; decision/recovery surface 0. Total: 3 — good JR verification task.
