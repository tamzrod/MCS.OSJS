# RLED-003 — Source TCP Health

Status: QUEUED
Previous: RLED-002
Next: RLED-004

## Primary outcome
Expose the observed TCP connection result for each real source poll.

## Scope
Record successful dial, refused/unreachable connection and dial timeout at `replicator/reader.go` and propagate to the existing per-block runtime status. A successful dial describes the last recent dial, not an always-open connection. Preserve errors and block identity.

## Non-scope
No ICMP check, UI change or persistent socket.

## Acceptance
1. Successful dial, refusal and timeout are distinguishable.
2. Unobserved, disabled and stale observations cannot be green.
3. Existing FC1–FC4 read behavior is unchanged.

## Verification
Focused Go tests with loopback listener and refused/timed-out dial fixtures; `go test ./replicator/...` and `go vet ./replicator/...`.

## Dependencies
RLED-002.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
