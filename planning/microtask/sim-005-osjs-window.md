# SIM-005 — Build the Approved OS.js Modbus Simulator Window

Status: planning material only. Human promotion is required before execution.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Primary Outcome

An OS.js Modbus Simulator application opens with the approved device-list/editor layout and edits simulator-owned definitions without exposing effective MMA2 configuration.

## Scope

- Create/open the Modbus Simulator using established MCS.OSJS OS.js application conventions.
- Implement the device list and search surface.
- Implement Add, Duplicate, Delete, Save & Apply, and Discard controls.
- Implement Name, Listen Port, Unit ID, and Enabled fields.
- Implement FC1-FC4 rows with Start, Count, and `Randomize Every (ms)`.
- Show calculated address ranges.
- Bind form values to simulator-owned definitions.

## Non-Scope

- No backend MMA2 activation logic.
- No scheduler implementation.
- No raw-ingest implementation.
- No live register table.
- No charts, waveform controls, scripts, or raw YAML editor.
- No Replicator configuration.

## Acceptance Criteria

1. The OS.js simulator application opens and presents the approved two-pane device-list/editor layout.
2. All approved device and FC fields can be edited and round-trip through the simulator configuration model.
3. Add, Duplicate, Delete, and Discard affect simulator-owned definitions only and do not directly modify effective MMA2 configuration.

## Verification

Open the application, create/select/edit/duplicate/delete/discard simulator definitions, and verify exact field round-trip through SIM-001 persistence without touching effective MMA2 config.

## Dependencies

- SIM-001.
- Existing MCS.OSJS OS.js application conventions.

## Sizing

**4 / 10 — Medium.** One UI outcome across several related controls. Backend behavior is deliberately excluded.