> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# RLED-005 — Optional Source Network Reachability

Status: QUEUED
Previous: RLED-004
Next: RLED-006

## Primary outcome
Expose an independent bounded ICMP reachability observation for the source endpoint.

## Scope
Check only the selected configured source host via a non-blocking, bounded platform-appropriate ping; rate-limit checks and avoid blocking Modbus pollers. Return unknown when ping is blocked, unsupported or indeterminate. Ensure DNS hostname handling does not become shell injection.

## Non-scope
No requirement that ICMP succeed for TCP or Modbus to be healthy; no external ping dependency bundled without verification.

## Acceptance
1. ICMP success is green; not-tested/blocked is gray; confirmed network failure has truthful detail.
2. Polling proceeds independently when ICMP fails.
3. Probe is bounded and safely handles the configured host.

## Verification
Focused platform helper tests and `go test ./replicator/...`; Windows behavior included in RLED-011.

## Dependencies
RLED-004.

## Sizing
Surface 1, environment 1, behavior 1, verification 1, recovery 0 = 4; split further if platform behavior proves uncertain.
