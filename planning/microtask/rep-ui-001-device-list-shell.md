# REP-UI-001 — Replicator App Shell and Launchers

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Create the Modbus Replicator OS.js application shell and make it directly launchable from both the OS.js desktop and Start/Application menu.

## Scope
- Add/open the Replicator OS.js application window using existing MCS.OSJS/Simulator UI patterns where practical.
- Register Modbus Replicator in the OS.js Start/Application menu.
- Add a Modbus Replicator desktop icon/shortcut.
- Reuse the Replicator application identity/icon/metadata for both launch surfaces where practical.
- Clicking either launcher opens the same existing Replicator application.
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
1. Modbus Replicator appears in the OS.js Start/Application menu and launches the Replicator window.
2. Modbus Replicator appears as a desktop icon/shortcut and launches the same Replicator window.
3. The Replicator window has the approved two-pane shell: device list on the left and selected device definition area on the right.
4. Add, Duplicate, and Delete controls are present without pretending backend actions are already implemented.

## Verification
Launch OS.js, open Modbus Replicator once from the Start/Application menu and once from the desktop icon, then visually verify both launch the same two-pane Replicator application shell.

## Dependencies
Existing MCS.OSJS OS.js shell and Simulator application/launcher patterns.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 0. Total 3 — good bounded task.
