# MEM-005 — Add Simulation Mode Selector

Status: QUEUED
Previous: MEM-004
Next: MEM-006

## Primary Outcome
Expose per-area Simulation mode in the Memory Device Definition table using `None` and `Random`.

## Scope
- Replace the single `Randomize Every (ms)` presentation with separate `Simulation` and `Interval (ms)` columns.
- Add a per-area selector with `None` and `Random`.
- Map `None` to interval 0 and disable the interval input.
- Map `Random` to a positive interval and enable the interval input.
- Keep FC1-FC4 area/start/count/address-range editing intact.

## Non-Scope
- No Sine Wave, Time, or other generators.
- No persisted schema change beyond existing interval fields.
- No backend API rename.

## Acceptance Criteria
1. Each FC1-FC4 row exposes `Simulation: None | Random`.
2. None stores interval 0 and disables interval editing.
3. Random stores a positive interval and enables interval editing.
4. Save & Apply continues through the existing Simulator runtime call path.

## Verification
Static renderer source re-read plus backend compatibility checks from MEM-004; rendered Windows verification is MEM-008.

## Dependencies
MEM-004.

## Sizing
Implementation 2; environment 0; behavior 2; verification 1; decision/recovery 0. Total 5 — tightly coupled renderer-only interaction with one deterministic workflow.