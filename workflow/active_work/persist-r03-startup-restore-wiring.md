# PERSIST-R03 — Startup Restore Wiring

Status: ACTIVE — ARMED 2026-10-08 by PERSIST-R02 delivery (OpenHands JR DEV)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-R02
Next: PERSIST-UI01

## Primary outcome
Wire the existing startup persistence loader/restore/verification/final-unseal contracts into the real runtime while preserving fail-closed State Sealing behavior.

## Scope
- Add real non-test startup call sites for `LoadPersistenceSnapshots` and `RestorePersistencePlan`.
- Persistence-enabled memory starts sealed, loads validated snapshots, restores through existing Raw Ingest v1, verifies the full required set, and performs the existing final explicit unseal only after success.
- Any loader, Raw Ingest, verification or commit failure remains sealed and preserves the existing deterministic failure classification.

## Non-scope
No new seal flag, no alternate unseal path, no automatic retry/repair, no filesystem-format redesign, no UI, no historian behavior, no ICC edits.

## Acceptance
1. Real startup path stays sealed through all restore writes and unseals only after existing verification succeeds.
2. Missing/corrupt/incompatible snapshot or Raw Ingest/verification/commit failure leaves the runtime sealed with the existing classified reason.
3. Focused runtime tests prove success and at least one fail-closed startup case using the real runtime orchestration path.

## Evidence / handoff
Record exact startup call sites, changed paths, focused tests, bounded regression and delivered source SHA. After genuine delivery, automatically arm PERSIST-UI01 for the next invocation and STOP.

## Dependencies
Requires genuine PERSIST-R02 delivery.

## Sizing
2/1/2/1/1=7.

## CWAL
Already human-promoted. Execute only when predecessor evidence is genuine and selected by handoff. After delivery, automatically activate PERSIST-UI01 and STOP.
