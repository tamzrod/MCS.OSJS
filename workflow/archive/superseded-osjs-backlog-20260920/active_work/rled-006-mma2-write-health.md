> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# RLED-006 — MMA2 Destination Write Health

Status: QUEUED
Previous: RLED-005
Next: RLED-007

## Primary outcome
Expose destination health from actual MMA2 Raw Ingest write acknowledgements.

## Scope
At the existing sendDestination path, record readiness after an observed successful connection/write, last confirmed successful write timestamp and failures (dial, write, response timeout or rejection). Reuse `mma2raw.Client.Send` acknowledgement; keep this state independent from source.

## Non-scope
No Raw Ingest protocol, MMA2 ownership, reservation or config change.

## Acceptance
1. Ack-confirmed write is success; rejected, timed-out or failed write is error.
2. Not-yet-tested and stopped are unknown, not green.
3. Source errors do not fabricate MMA2 write failures.

## Verification
Focused local Raw Ingest test fixture and `go test ./replicator/...`; `go vet ./replicator/...`.

## Dependencies
RLED-005.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
