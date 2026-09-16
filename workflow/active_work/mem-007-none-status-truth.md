# MEM-007 — Truthful None-mode Status

Status: QUEUED
Previous: MEM-006
Next: MEM-008

## Primary outcome
All-None device is not incorrectly shown waiting for a simulation ingest cycle.

## Scope
Distinguish no enabled positive-interval area from active random generator awaiting ingest. Continue reporting MMA2 listener health separately. Add focused all-None, mixed, disabled and unavailable MMA2 status tests.

## Non-scope
No new protocol, ownership, lifecycle or external status API.

## Acceptance
1. All-None with reachable MMA2 reports memory available and Raw Ingest not required.
2. Mixed None/Random still waits for successful ingest when applicable.
3. Disabled or unavailable MMA2 never produces false RUNNING state.

## Verification
`cd simulator && go test -count=1 ./... && go vet ./...`; Windows live verification in MEM-008.

## Dependencies
MEM-006.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
