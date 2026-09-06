# SIM-007 — Show Simulator Runtime Status

Status: ACTIVE — human-promoted for JR execution.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Primary Outcome

The OS.js simulator window displays enough runtime state to prove the selected device, MMA2 path, raw-ingest path, and each FC scheduler are operating.

## Scope

- Show selected-device status as `RUNNING`, `STOPPED`, or `ERROR`.
- Show MMA2 runtime status.
- Show raw-ingest status.
- Show FC1-FC4 last random update times.
- Show FC1-FC4 next update timing/state.
- Show total configured points for the selected simulator device.
- Refresh status without turning the application into a live register viewer.

## Non-Scope

- No live memory/register table.
- No charts or historical trending.
- No waveform controls.
- No new simulator control behavior.

## Acceptance Criteria

1. The selected device shows current device, MMA2, and raw-ingest status.
2. FC1-FC4 last/next timing shown in the UI agrees with the active scheduler state.
3. Total points equals the selected device's configured FC counts.

## Verification

Run one simulator device with different FC intervals, observe multiple update cycles, and compare displayed runtime state against scheduler/MMA2 evidence.

## Dependencies

- SIM-004.
- SIM-005.
- SIM-006.

## Sizing

**2 / 10 — Small.** One bounded observability outcome using already-established runtime state.