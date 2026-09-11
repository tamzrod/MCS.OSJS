# REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

Status: COMPLETED / VERIFIED

## Ownership Rule
Shared MMA2 ownership is keyed by the exact `(port, unit_id)` pair. Same port/different unit and same unit/different port are valid when their exact pairs are free.

## Primary Outcome
Make Save & Apply the authoritative shared-MMA2 transaction boundary for Replicator destination ownership and activation.

## Verified Behavior
- Exact foreign pair rejects before shared mutation.
- Same Unit ID on a different port succeeds when free.
- Same port with a different Unit ID succeeds when free.
- Successful apply composes shared MMA2 settings, requests restart, waits for acknowledgement/readiness, and only then reports success.
- Rejected apply preserves foreign ownership and the previously persisted Replicator configuration.

## Verification Result
Backend gates and rendered UI Save & Apply tests passed in the corrected pair-scoped ownership test packet recorded in `handoff.md`.
