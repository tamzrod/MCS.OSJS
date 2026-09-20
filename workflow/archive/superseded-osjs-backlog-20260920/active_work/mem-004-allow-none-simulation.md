> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# MEM-004 — Allow None Simulation

Status: ACTIVE
Previous: MEM-003
Next: MEM-005

Paused: Human approved switching the sole ACTIVE task to RLED-001. This is a pause, not a completion; verification remains pending.

## Primary outcome
Allow configured memory area with count > 0 and zero simulation interval.

## Scope
Relax Go validation for zero intervals. Reuse scheduler's existing positive-interval-only behavior. Add focused None/no-generation tests.

## Non-scope
No MMA2 area changes, ingest protocol changes, new generators or schema rename.

## Acceptance
1. FC1–FC4 count > 0 with interval 0 validates and persists.
2. No random writes are scheduled for None.
3. Positive interval remains functional.

## Verification
`cd simulator && go test -count=1 ./... && go vet ./...`; record unavailable gates rather than claim success. Windows end-to-end in MEM-008.

## Dependencies
MEM-003.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
