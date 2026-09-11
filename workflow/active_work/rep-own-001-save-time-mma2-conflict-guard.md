# REP-OWN-001 — Save-Time Shared MMA2 Conflict Guard

## Primary Outcome
Make Save & Apply the authoritative shared-MMA2 transaction boundary for Replicator destination ownership and activation.

## Ownership Rule
Port and Unit ID are independently exclusive shared MMA2 resources for this workflow.

A destination is considered free only when:
- its **Port** is not owned by another producer; and
- its **Unit ID** is not owned by another producer.

Replicator must not treat a different `(port, unit_id)` pair as free when either the port or the unit ID is already owned by another producer. Automatic allocation must skip both foreign-owned ports and foreign-owned Unit IDs.

## Required Save & Apply Sequence
When the operator clicks **Save & Apply**, execute these steps in this order:

1. **Check Port + Unit ID ownership independently**
   - Read the latest shared MMA2 ownership state.
   - Check the requested destination port against all ownership entries.
   - Check the requested destination Unit ID against all ownership entries.
   - Do not trust an earlier UI Owner/Status snapshot as the safety check.

2. **Claim only if both are free; otherwise return error**
   - If both Port and Unit ID are free, claim them for `replicator`.
   - If either resource is owned by another producer, stop immediately and return a visible conflict error containing the actual owner and conflicting port or Unit ID.
   - Do not overwrite, delete, merge into, or mutate the foreign reservation.
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
- Automatic destination selection must choose a port not already owned and a Unit ID not already owned.
- Save-time validation must re-read the latest ownership document and independently reject foreign port or Unit ID ownership before persistence/composition.
- Do not persist the edited Replicator document when the apply conflicts.
- Do not drop/rebuild Replicator reservations until the ownership check for the candidate document succeeds.
- UI must surface Save & Apply ownership/restart failures as errors.
- Pre-save Owner/Status display is advisory only and must not be the safety boundary.
- Add regression tests for same-pair, same-port/different-unit, and different-port/same-unit foreign conflicts.

## Acceptance Criteria
1. Simulator owns `(5020,1)`; Replicator Save & Apply requesting `(5020,1)` fails.
2. Simulator owns port `5020`; Replicator requesting `(5020,2)` also fails because the port is foreign-owned.
3. Simulator owns Unit ID `1`; Replicator requesting `(5021,1)` also fails because the Unit ID is foreign-owned.
4. Automatic allocation never returns foreign-owned port `5020` or foreign-owned Unit ID `1` in that state.
5. Error identifies the conflicting resource and owner `simulator`.
6. Simulator ownership/config remains unchanged after rejected saves.
7. Previously persisted Replicator document/reservations remain unchanged after rejected saves.
8. A destination whose Port and Unit ID are both free is claimed by `replicator`, written into shared MMA2 settings, then activated through MMA2 restart.
9. Save & Apply reports success only after the MMA2 restart acknowledgement/readiness path completes.
10. Restart failure produces a visible apply error rather than false success.

## Verification
Backend regression tests plus browser test using Simulator and Replicator side by side:

- same pair: Simulator `(5020,1)` → Replicator `(5020,1)` → reject;
- same port: Simulator `(5020,1)` → Replicator `(5020,2)` → reject;
- same unit: Simulator `(5020,1)` → Replicator `(5021,1)` → reject;
- auto allocation: must choose neither port `5020` nor Unit ID `1`;
- free path: choose a Port and Unit ID unused by any producer → Save & Apply claims them → shared MMA2 config/owners contain them → MMA2 restart completes → destination becomes available.

## Sizing
Implementation 1; behavior 1; verification 1; recovery/atomicity 1; lifecycle 1. Total 5 — tightly bounded transaction/lifecycle task.
