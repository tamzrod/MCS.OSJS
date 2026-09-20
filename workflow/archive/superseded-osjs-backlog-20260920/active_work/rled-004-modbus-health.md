> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# RLED-004 — Modbus Application Health

Status: QUEUED
Previous: RLED-003
Next: RLED-005

## Primary outcome
Expose a truthful Modbus response outcome separately from TCP connectivity.

## Scope
Classify successful FC1–FC4 responses, Modbus exception (FC and exception code), read timeout, malformed response and not-tested state using structured per-block observations. Retain current legacy status/error fields for compatibility.

## Non-scope
No change to Modbus request semantics or MMA2 writes.

## Acceptance
1. Valid response is OK; exception is WARNING; Modbus response timeout is ERROR.
2. TCP failure does not falsely claim a Modbus response failure.
3. Exception FC/code and offending pull block are available for tooltips.

## Verification
Focused Go unit fixtures covering success, exception and timeout; `go test ./replicator/...` and `go vet ./replicator/...`.

## Dependencies
RLED-003.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
