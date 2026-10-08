# PERSIST-R02 — Runtime Save Wiring

Status: QUEUED — HUMAN PROMOTED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-R01
Next: PERSIST-R03

## Primary outcome
Wire persistence-owned RBE events in the real runtime to the existing `PersistenceSnapshotWriter` and the PERSIST-R01 filesystem adapter.

## Scope
- Create the non-test runtime call site for `NewPersistenceSnapshotWriter`.
- Subscribe only the system-derived persistence RBE projection and route its one-byte rule events to snapshot writes.
- Preserve change-only behavior: read the authoritative configured area on event and persist only changed bytes/words through existing writer semantics.

## Non-scope
No startup restore, no polling loop, no user-RBE mutation, no alternate ranges, no UI, no historian behavior, no ICC edits.

## Acceptance
1. A real runtime persistence RBE event reaches `PersistenceSnapshotWriter` for the correct memory/area identity.
2. User-owned RBE remains independent and persistence-owned ranges remain derived/system-owned.
3. Focused runtime tests prove changed state saves and unchanged state causes no unnecessary snapshot rewrite.

## Evidence / handoff
Record exact runtime call sites, changed paths, focused tests, bounded regression and delivered source SHA. After genuine delivery, automatically arm PERSIST-R03 for the next invocation and STOP.

## Dependencies
Requires genuine PERSIST-R01 delivery.

## Sizing
2/1/1/1/0=5.

## CWAL
Already human-promoted. Do not request another approval. When predecessor evidence is genuine and this packet is selected by handoff, execute exactly this task, then automatically activate PERSIST-R03 for the next invocation and STOP.
