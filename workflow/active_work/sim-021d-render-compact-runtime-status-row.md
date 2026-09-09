# SIM-021D — Render the Compact MMA2 + Simulator Status Row

Status: QUEUED
Previous: SIM-021C
Next: SIM-021E

## Primary outcome

Render one compact status row in the Device Definition pane showing only MMA2 and selected-device Simulator state.

## Scope

- Add a compact row directly below the `Device Definition` heading and above the editable fields.
- Render exactly two labeled state pairs:
  - `MMA2  ● <STATE>`
  - `Simulator  ● <STATE>`
- Bind the row only to status data produced by SIM-021C.
- Use both a visible state word and a small status indicator; color must not be the only status cue.
- Keep the row height and spacing minimal so the existing form/table layout remains effectively unchanged.
- When no device is selected, show Simulator as `—`/no-device rather than inventing a runtime state.
- When current status data is unavailable, show an unavailable/waiting presentation rather than stale `RUNNING`.
- Keep the existing bottom `.sim-status` bar for transient operation/error messages only.

## Non-scope

- No Raw Ingest status label.
- No FC Last/Next timestamps.
- No total-points display.
- No collapsible diagnostics panel.
- No device-list status chips.
- No new top/global header strip.

## Acceptance criteria

1. The Device Definition pane shows only MMA2 and Simulator runtime states in one compact row.
2. `RUNNING`, `WAITING`, `STOPPED`, and `ERROR` remain distinguishable without relying on color alone.
3. No-device and unavailable-status cases never display stale `RUNNING`.
4. Existing form controls, function rows, Save & Apply, Discard, and bottom message bar remain usable and visually intact.

## Verification

Run `node --check`, the ModbusSimulator/local-package build, and inspect the rebuilt window at its normal size for layout regression.

## Dependencies

- SIM-021C.

## Sizing

Implementation 1, environment 0, behavioral 0, verification 1, decision/recovery 0 = 2. One bounded UI presentation change.
