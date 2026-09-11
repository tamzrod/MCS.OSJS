# REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

## Primary Outcome
Make Save & Apply the authoritative shared-MMA2 transaction boundary for Replicator destination ownership and activation.

## Required Save & Apply Sequence
When the operator clicks **Save & Apply**, execute these steps in this order:

1. **Check Port + Unit ID ownership**
   - Read the latest shared MMA2 ownership state for the requested destination `(port, unit_id)`.
   - Do not trust an earlier UI Owner/Status snapshot as the safety check.

2. **Claim only if free; otherwise return error**
   - If `(port, unit_id)` is free, claim it for `replicator`.
   - If it is already owned by another producer, stop immediately and return a visible conflict error containing the actual owner and reservation.
   - Do not overwrite, delete, or mutate the foreign reservation.
   - Do not change the previously persisted Replicator configuration on conflict.

3. **Save into shared MMA2 settings**
   - After ownership validation succeeds, compose/persist the Replicator destination into the shared MMA2 effective configuration and ownership document.
   - Persist the Replicator device document only as part of the successful apply path.
   - Preserve all foreign producer configuration/reservations.

4. **Restart MMA2 and verify activation**
   - Request MMA2 restart through the established restart-request/ack lifecycle.
   - Wait for the matching restart acknowledgement and destination readiness before reporting success.
   - If restart/apply fails, report failure truthfully rather than claiming Save & Apply succeeded.

## Scope
- Treat shared MMA2 ownership as authoritative at Save & Apply time, regardless of any earlier UI suggestion/inspection state.
- Check the requested resolved `(port, unit_id)` against the latest ownership document before persistence/composition.
- A reservation owned by another producer (Simulator or any future producer) returns a conflict error containing the actual owner and reservation.
- Do not persist the edited Replicator document when the apply conflicts.
- Do not drop/rebuild Replicator reservations until the conflict check for the candidate document succeeds.
- UI must surface Save & Apply ownership/restart failures as errors.
- Pre-save Owner/Status display is advisory only and must not be the safety boundary.
- Add a regression test where Simulator owns a reservation and Replicator manually attempts to save the same reservation.

## Acceptance Criteria
1. Simulator owns `(5020,1)`; Replicator Save & Apply requesting `(5020,1)` fails before shared MMA2 config mutation.
2. Error identifies `(5020,1)` and owner `simulator`.
3. Simulator ownership/config remains unchanged after the rejected save.
4. Previously persisted Replicator document/reservations remain unchanged after the rejected save.
5. A genuinely free destination is claimed by `replicator`, written into shared MMA2 settings, then activated through MMA2 restart.
6. Save & Apply reports success only after the MMA2 restart acknowledgement/readiness path completes.
7. Restart failure produces a visible apply error rather than false success.

## Verification
Backend regression tests plus browser test using Simulator and Replicator side by side:

- conflict path: Simulator owns `(5020,1)` → Replicator manually selects `(5020,1)` → Save & Apply returns conflict and changes nothing;
- free path: choose a free `(port, unit_id)` → Save & Apply claims it → shared MMA2 config/owners contain it → MMA2 restart completes → destination becomes available.

## Sizing
Implementation 1; behavior 1; verification 1; recovery/atomicity 1; lifecycle 1. Total 5 — tightly bounded transaction/lifecycle task.
