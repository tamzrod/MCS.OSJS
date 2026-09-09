# REP-002 — Share MMA2 Raw-Ingest Client

Status: QUEUED
Previous: REP-001
Next: REP-003

## Primary Outcome

Move/generalize the proven Simulator MMA2 raw-ingest packet/client implementation into a reusable internal package so Simulator and Replicator use one MMA2 write path.

## Scope

- Extract the producer-neutral packet encoding and TCP send/ack behavior currently implemented in `simulator/raw_ingest.go`.
- Preserve the established raw-ingest v1 wire contract and FC1-FC4 area encoding behavior.
- Update Simulator to use the shared client with no intended runtime behavior change.
- Expose a small API that lets Replicator later write mapped values into its own MMA2 destination memory.

## Non-Scope

- No Modbus source acquisition.
- No Replicator scheduling/poll loop.
- No MMA2 config ownership changes.
- No protocol redesign.
- No OS.js UI.

## Acceptance Criteria

1. Simulator no longer owns a duplicate producer-specific raw-ingest encoder/client implementation.
2. Existing Simulator raw-ingest packet bytes and response/error semantics remain unchanged.
3. The shared client can be constructed independently of Simulator-specific device/runtime types.

## Verification

Move or adapt the existing raw-ingest tests to the shared package and run the affected Simulator tests. Packet fixtures for coils, discrete inputs, holding registers, and input registers must remain byte-identical.

## Dependencies

- Existing completed Simulator raw-ingest implementation.

## Sizing

Implementation surface 1; environment uncertainty 0; behavioral surface 1; verification surface 1; decision/recovery surface 0. Total: 3 — good JR task.
