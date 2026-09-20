> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# RLED-007 — Aggregate Device Health Truthfully

Status: QUEUED
Previous: RLED-006
Next: RLED-008

## Primary outcome
Expose one selected-device LED state per layer while keeping block-level reasons accessible.

## Scope
Aggregate existing block observations with deterministic severity and freshness derived from scan rates/timeouts; preserve FC/block error details and source/destination independence. Reset to unknown on disabled/unobserved or stale data.

## Non-scope
No new polling threads or change to device selection/config.

## Acceptance
1. Mixed successful/failed blocks retain the failed block and accurate LED severity.
2. Outdated successes do not stay green indefinitely.
3. Disabled and unobserved states are gray, never green.

## Verification
Table-driven Go tests for mixed blocks, aging, disabled and recovery; `go test ./replicator/...`.

## Dependencies
RLED-006.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
