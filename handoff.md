# Handoff

## Current

ACTIVE: REP-UI-001 — Replicator App Shell and Launchers
State: READY FOR IMPLEMENTATION

All Replicator UI microtasks from `brainstorm/replicator-ui-layout.md` are promoted to `workflow/active_work/` in sequence.

## Authorized Sequence

1. REP-UI-001 — Replicator App Shell and Launchers
2. REP-UI-002 — Replicator Source Device Editor
3. REP-UI-003 — Replicator Auto Destination Allocation
4. REP-UI-004 — Replicator Ownership Guard
5. REP-UI-005 — Replicator Save and Apply Device
6. REP-UI-006 — Replicator Runtime Apply Control
7. REP-UI-007 — Replicator Runtime Status

## Current Task Boundary

Implement REP-UI-001 only until it is ready for verification.

Required outcome:

- Modbus Replicator appears in the OS.js Start/Application menu;
- Modbus Replicator appears as a desktop icon/shortcut;
- both launch the same Replicator application;
- application window uses the approved two-pane shell;
- left pane contains the device list plus Add, Duplicate, Delete controls;
- right pane is the selected-device definition container;
- no backend config persistence, source polling, destination allocation, ownership mutation, or runtime behavior is added in REP-UI-001.

## Workflow State

`planning/microtask/` contains only its workflow support files. REP-UI-001 through REP-UI-007 now live under `workflow/active_work/` and are implementation-authoritative.

No JR test packet is active yet. The coding agent should implement REP-UI-001 first, then write the exact verification packet here before JR execution.
