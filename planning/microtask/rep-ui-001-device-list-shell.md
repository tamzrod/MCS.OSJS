# REP-UI-001 — Replicator Device List Shell

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Create the Modbus Replicator OS.js application shell with the approved two-column layout: device list on the left and selected device definition area on the right.

## Scope
- Add/open the Replicator OS.js application window using existing MCS.OSJS/Simulator UI patterns where practical.
- Left pane contains device list presentation and Add, Duplicate, Delete controls.
- Right pane provides the structural container for the selected device definition.
- Preserve the current brainstorm layout direction without wiring Replicator backend behavior yet.

## Non-Scope
- No config persistence.
- No source polling.
- No destination allocation.
- No ownership mutation.
- No runtime start/stop behavior.

## Acceptance Criteria
1. Replicator app opens as an OS.js window.
2. Left device pane and right definition pane are visibly distinct and match the approved layout direction.
3. Device controls are present without pretending backend actions are already implemented.

## Verification
Launch the OS.js desktop, open Modbus Replicator, and visually verify the two-pane shell and controls.

## Dependencies
Existing MCS.OSJS OS.js shell and Simulator application patterns.

## Sizing
Implementation 1; environment 0; behavior 0; verification 1; decision/recovery 0. Total 2 — good bounded task.
