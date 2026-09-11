# REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

## Primary Outcome
Make Save & Apply the authoritative shared-MMA2 transaction boundary for Replicator destination ownership and activation.

## Ownership Rule
The shared MMA2 ownership key is the exact **`(port, unit_id)` pair**.

Examples:
- Simulator `(5020,1)` blocks Replicator `(5020,1)`.
- Simulator `(5020,1)` does **not** block Replicator `(5022,1)`.
- Simulator `(5020,1)` does **not** block Replicator `(5020,2)`.

Automatic allocation skips occupied pairs. Port numbers and Unit IDs are not independently/global-exclusive resources.

## Required Save & Apply Sequence
1. Re-read the latest shared MMA2 ownership state for the requested destination pair.
2. Reject only an exact foreign-owned `(port, unit_id)` pair before shared mutation.
3. Preserve all foreign configuration and ownership entries.
4. Compose/persist the Replicator slice into shared MMA2 settings.
5. Request MMA2 restart through the established request/ack path and wait for readiness.
6. Persist the Replicator document and report success only after activation succeeds.

## Scope
- Save & Apply is the authoritative ownership check; earlier UI Owner/Status is advisory.
- Exact foreign pair conflicts must not mutate MMA2 or the previously persisted Replicator document.
- Same port/different unit and same unit/different port remain valid when those exact pairs are free.
- UI surfaces ownership/restart failures visibly.

## Acceptance Criteria
1. Simulator `(5020,1)` + Replicator `(5020,1)` rejects with owner `simulator`.
2. Simulator `(5020,1)` + Replicator `(5022,1)` succeeds when `(5022,1)` is free.
3. Simulator `(5020,1)` + Replicator `(5020,2)` succeeds when `(5020,2)` is free.
4. Automatic allocation skips occupied exact pairs.
5. Rejected apply preserves foreign ownership/config and previous Replicator persistence.
6. Successful Save & Apply writes shared MMA2 settings and completes restart acknowledgement/readiness before reporting success.

## Verification
Backend pair-scoped regression tests plus rendered UI Save & Apply tests for same-unit/different-port success, same-port/different-unit success, and exact-pair rejection.

## Sizing
Implementation 1; behavior 1; verification 1; recovery/atomicity 1; lifecycle 1. Total 5 — tightly bounded transaction/lifecycle task.
