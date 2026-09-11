# REP-UI-004 — Replicator Ownership Guard

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Enforce first-come, first-served MMA2 ownership so Replicator can mutate only Replicator-owned destination reservations and cannot overwrite or delete foreign-owned memory.

## Scope
- Ownership key remains `(port, unit_id)`.
- Free reservation may be claimed by Replicator.
- Existing Replicator-owned reservation may be updated by Replicator.
- Simulator/other producer ownership causes a truthful collision rejection.
- Replicator deletion/rebuild paths may remove only Replicator-owned reservations.
- Surface owner and collision state to the Replicator UI.

## Non-Scope
- No change to Simulator ownership authority.
- No forced takeover mechanism.
- No ownership priority beyond first-come, first-served.
- No runtime restart/apply orchestration.

## Acceptance Criteria
1. Replicator cannot overwrite a foreign-owned `(port, unit_id)`.
2. Replicator cannot delete a foreign-owned reservation/memory while rebuilding its own state.
3. Replicator can update/release its own reservation.
4. UI reports the actual foreign owner on collision.

## Verification
Focused ownership tests with free, Replicator-owned, and Simulator-owned reservations; verify foreign entries remain byte/logically intact after Replicator update/delete attempts.

## Dependencies
REP-UI-003; existing `mma2composer` collision and producer reservation behavior.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 1. Total 3 — good bounded task.
