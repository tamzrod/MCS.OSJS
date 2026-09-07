# SIM-005 — Build the Approved OS.js Modbus Simulator Window

Status: COMPLETED — verified 2026-09-08.

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

**Evidence (2026-09-08):** Replaced the corrupt `OSJS/src/packages/ModbusSimulator/index.js` with an ASCII-clean OS.js application and added `index.scss`. The application provides the approved two-pane searchable device list/editor, Add/Duplicate/Delete, Save & Apply/Discard, Name/Enabled/Listen Port/Unit ID, FC1-FC4 Start/Count/Randomize Every fields, calculated address ranges, inline validation, and same-origin JSON persistence. Corrected the blocking missing parenthesis in `OSJS/src/server/providers/simulator-bridge.js`. In an isolated runtime (`OSJS_DATA_DIR=/tmp/mcs-osjs-cwal-data`), a real in-app browser opened the desktop application, added and edited `CWAL Test PLC` (port 15020), saved it, duplicated/deleted it locally, and verified Discard restored the persisted definition. `GET /api/devices` and the resulting `config/simulator/devices.yaml` matched exactly; the isolated data tree contained no MMA2 effective-config artifact. `node --check src/server/providers/simulator-bridge.js`, `npm run build:local-packages`, `npm run build`, `gofmt -l .`, `go vet ./...`, and `go test -count=1 ./...` all passed (Go tests used loopback socket permission).

## Dependencies

- SIM-001.
- Existing MCS.OSJS OS.js application conventions.

## Sizing

**4 / 10 — Medium.** One UI outcome across several related controls. Backend behavior is deliberately excluded.
