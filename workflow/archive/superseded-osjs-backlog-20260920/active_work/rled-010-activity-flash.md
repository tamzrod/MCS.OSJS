> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# RLED-010 — Event-Driven Communication Flashes

Status: QUEUED
Previous: RLED-009
Next: RLED-011

## Primary outcome
Flash healthy LEDs briefly on genuinely new source activity and confirmed MMA2 writes.

## Scope
Use monotonically advancing event counters or unique timestamps from real runtime observations. A source TCP LED may flash on actual I/O; MMA2 flashes only after Raw Ingest positive acknowledgement. A normal status poll must not retrigger activity; never blink indefinitely.

## Non-scope
No artificial animation, new poller or flash on failed writes.

## Acceptance
1. Exactly one finite visual pulse per newly observed activity event.
2. Same status polled repeatedly causes no repeated pulse.
3. Failure never produces a green successful-write flash.

## Verification
Focused event deduplication/animation checks and Windows visual verification in RLED-011.

## Dependencies
RLED-009.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
