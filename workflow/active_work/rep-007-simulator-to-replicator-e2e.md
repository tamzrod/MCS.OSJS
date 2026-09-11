# REP-007 — Simulator-to-Replicator End-to-End Verification

Status: ACTIVE
Previous: REP-006
Next: none

## Primary Outcome

Prove the first complete Replicator milestone by using the working Simulator as the controlled external Modbus source and verifying that changing source values are reproduced in a separate Replicator-owned MMA2 destination.

## Scope

- Run the existing working Simulator configuration path with a small deterministic FC3 register range.
- Configure Replicator to read that Simulator-owned Modbus endpoint.
- Configure a distinct Replicator-owned MMA2 `(port, unit_id)` destination.
- Start the Replicator poll loop.
- Read the destination through normal Modbus and verify values match the Simulator source.
- Change Simulator-owned source values through the same Raw Ingest path used by Simulator runtime generation and verify the destination follows on a later poll.
- Confirm the Simulator reservation remains intact and separately owned.

## Non-Scope

- No real field hardware.
- No multiple source devices or ranges.
- No transformations/scaling.
- No OS.js Replicator UI.
- No performance/stress testing.

## Acceptance Criteria

1. Initial source values exposed through the Simulator-owned MMA2 reservation appear unchanged at the Replicator MMA2 destination.
2. After Simulator-owned source values change, the Replicator destination updates to the new values within the expected polling window.
3. Simulator and Replicator use distinct ownership entries/reservations and neither overwrites the other.

## Verification

`replicator/e2e_test.go` provides one explicit automated end-to-end workflow using the actual Simulator store/composer for the source reservation, the shared MMA2 appliance, Simulator's Raw Ingest transport semantics, the Replicator poll loop, normal Modbus destination reads, a second source update, and final ownership verification.

JR runs the focused REP-007 E2E test plus full Replicator regression and `go vet`.

## Dependencies

- Working basic Simulator milestone.
- REP-001 through REP-006 completed and verified.

## Sizing

Implementation surface 0; environment uncertainty 1; behavioral surface 1; verification surface 1; decision/recovery surface 0. Total: 3 — verification-only milestone.
