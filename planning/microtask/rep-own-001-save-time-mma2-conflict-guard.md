# REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

## Primary Outcome
Make Save & Apply authoritative for shared MMA2 ownership: if the requested destination `(port, unit_id)` is already owned by another MMA2 producer, the save must fail visibly and must not alter either producer's reservation.

## Scope
- Treat shared MMA2 ownership as authoritative at Save & Apply time, regardless of any earlier UI suggestion/inspection state.
- Before persistence/composition, check the requested resolved `(port, unit_id)` against the latest ownership document.
- A reservation owned by another producer (Simulator or any future producer) returns a conflict error containing the actual owner and reservation.
- Do not persist the edited Replicator document when the apply conflicts.
- Do not drop/rebuild Replicator reservations until the conflict check for the candidate document succeeds.
- UI must surface the Save & Apply conflict as an error; pre-save Owner/Status display is advisory only and must not be the safety boundary.
- Add a regression test where Simulator owns a reservation and Replicator manually attempts to save the same reservation.

## Acceptance Criteria
1. Simulator owns `(5020,1)`; Replicator Save & Apply requesting `(5020,1)` fails.
2. Error identifies `(5020,1)` and owner `simulator`.
3. Simulator ownership/config remains unchanged after the rejected save.
4. Previously persisted Replicator document/reservations remain unchanged after the rejected save.
5. A genuinely free destination still applies normally.

## Verification
Backend regression test plus browser test using the Simulator and Replicator side by side; manually request the Simulator-owned destination and click Save & Apply.

## Sizing
Implementation 1; behavior 1; verification 1; recovery/atomicity 1. Total 4 — bounded.
