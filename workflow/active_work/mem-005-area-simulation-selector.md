# MEM-005 — Area Simulation Selector

Status: QUEUED
Previous: MEM-004
Next: MEM-006

## Primary outcome
Add per-area None/Random dropdown to the existing Device Definition table.

## Scope
Retain FC1–FC4, Start, Count, Address Range and existing actions. None maps to interval 0 and disables interval input; Random uses/restores positive interval (default 1000 ms). Preserve per-area interval when switching within an editing session.

## Non-scope
No schema change, Time/Sine Wave generator or architecture overhaul.

## Acceptance
1. Each area independently selects None/Random.
2. Existing save/load uses numeric interval values.
3. Controls remain usable while status updates occur.

## Verification
`node --check electron/renderer/app.js` and static DOM mapping inspection; Windows interaction in MEM-008.

## Dependencies
MEM-004.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
