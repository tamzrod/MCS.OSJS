# SIM-021 — Restore Truthful Runtime Status in the OS.js Window

Status: RETIRED — 2026-09-08 by human decision. Not completed;simulator window status work moves into the new brainstorm topic (`modbus-simulator-status-display.md`(, where placement remains undecided. This record stays for preservation.

## Primary outcome

Show operators whether MMA2 is reachable and whether the selected Simulator device is actively producing values accepted through Raw Ingest.

## Scope

- Read `SchedulerApplier.RuntimeStatus` through the SIM-018-approved local boundary.
- Show selected-device state (`RUNNING`, `STOPPED`, or `ERROR`).
- Show MMA2 state (`RUNNING`, `RESTARTING`/`WAITING`, `STOPPED`, or `ERROR`) using actual readiness evidence.
- Show Raw Ingest state (`OK`, `WAITING`, or `ERROR`) and the last error when present.
- Show FC1–FC4 configured point counts plus last successful and next scheduled update timestamps.
- Show last successful apply outcome/time and refresh while the window is open.

## Non-scope

- No register-value viewer, charts, history, or waveform editor.
- No health inferred from saved settings, an enabled checkbox, or merely having a schedule configured.
- No MMA2 start/stop controls.

## Acceptance criteria

1. An enabled but unready device never displays RUNNING.
2. MMA2 state changes visibly across ready, restart/wait, and unavailable conditions.
3. Raw Ingest displays OK only after a successful ingest; failures remain visible until a later success and include actionable error text.
4. For each configured FC, Last advances after successful sends and Next reflects scheduler timing; unused FCs display inactive.

## Verification

Use one enabled four-FC device, observe multiple successful cycles, interrupt MMA2 or Raw Ingest, and restore it; compare every displayed transition and FC timestamp with runtime evidence.

## Dependencies

- SIM-019.
- SIM-020 for presenting the last apply result in the same window.

## Sizing

Implementation 1, environment 0, behavioral 2, verification 1, decision/recovery 0 = 4. One bounded observability surface backed by the established runtime status model.
