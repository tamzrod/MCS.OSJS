# REP-002 — Share MMA2 Raw-Ingest Client

Status: COMPLETED 2026-09-10 — shared Raw Ingest client and Simulator migration verified.
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

## Completion Evidence

- `mma2raw` is a Simulator-independent Go module exposing Raw Ingest areas, configured ranges, a generic client, and the v1 encoder/send/ack implementation.
- Exact packet fixtures cover FC1 coils, FC2 discrete inputs, FC3 holding registers, and FC4 input registers, including LSB-first bits, big-endian registers, and header fields.
- Simulator now maps its device parameters and scheduler values onto `mma2raw.Client`; its private encoder and TCP send implementation were removed.
- `GOCACHE=/tmp/mcs-osjs-rep002-go-cache go test -count=1 ./...` and `go vet ./...` passed in `mma2raw/`.
- The same full test and vet gates passed in `simulator/`; the suite reported `ok github.com/tamzrod/MCS.OSJS/simulator 1.417s`.
