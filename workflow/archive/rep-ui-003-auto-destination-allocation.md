# REP-UI-003 — Replicator Auto Destination Allocation

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Provide automatic Replicator destination `(port, unit_id)` selection from the shared MMA2 ownership state, with explicit manual override.

## Scope
- Default destination mode is automatic.
- Auto-fill the next available destination port/reservation.
- Auto-fill the next available Unit ID for the selected port.
- Display Port, Unit ID, Owner, and availability status.
- Allow the user to explicitly override Port and/or Unit ID.
- Allocation lookup must respect the shared ownership registry rather than hardcoded assumptions.

## Non-Scope
- No reservation commit yet.
- No deletion/rebuild behavior yet.
- No Replicator runtime restart/apply behavior.
- No multiple destinations.

## Acceptance Criteria
1. A new device receives an available destination suggestion automatically.
2. Suggested `(port, unit_id)` does not collide with an existing shared ownership entry.
3. User can manually specify a destination instead of the suggestion.

## Verification
Use a controlled ownership document containing occupied and free reservations; verify auto-selection skips occupied entries and manual override remains editable.

## Dependencies
REP-UI-002; existing `mma2composer` ownership registry.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 1. Total 3 — good bounded task.
