# SIM-007 — Show Simulator Runtime Status

Status: COMPLETED — verified 2026-09-08.

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

**Evidence (2026-09-08):** Added selected-device runtime status to the applying bridge and OS.js window. `SchedulerApplier.RuntimeStatus` reports device `RUNNING`/`STOPPED`/`ERROR`, MMA2 TCP reachability, observed raw-ingest state/error, exact total FC point count, and per-FC scheduler `Last`/`Next` timestamps. The UI polls the selected device once per second and renders a compact status summary plus FC1-FC4 timing rows without exposing register values or history. `TestSchedulerApplierRuntimeStatusMatchesTimingAndPoints` ran all four FCs at distinct 5/7/9/11 ms intervals against a live raw-ingest fixture and proved RUNNING MMA2/device state, populated last/next times, and exact configured-point total. Real-browser verification displayed the isolated device's truthful ERROR/STOPPED/ERROR state with total 64 and advancing FC1-FC4 last/next timestamps while MMA2 was intentionally absent. `gofmt -l .`, `go vet ./...`, `go test -count=1 ./...`, `node --check`, `npm run build:local-packages`, and `npm run build` pass.

## Dependencies

- SIM-004.
- SIM-005.
- SIM-006.

## Sizing

**2 / 10 — Small.** One bounded observability outcome using already-established runtime state.
