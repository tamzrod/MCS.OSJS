# MEM-007 — Make None-Mode Status Truthful

Status: QUEUED
Previous: MEM-006
Next: MEM-008

## Primary Outcome
Prevent a Memory device with no active simulation schedules from being shown as waiting for random generation.

## Scope
- Distinguish memory availability from simulator-generation activity in runtime status.
- When all configured areas use None and MMA2 is available, do not report a misleading simulation WAITING state merely because Raw Ingest has never fired.
- Preserve existing ERROR/WAITING/RUNNING semantics for devices that actually have Random schedules.

## Non-Scope
- No header Simulator indicator.
- No new diagnostics protocol.
- No new generator types.

## Acceptance Criteria
1. All-None enabled device can report memory/MMA2 availability without pretending random simulation is pending.
2. Random-enabled devices retain existing generation/error semantics.
3. Raw Ingest errors remain visible when generation is active.

## Verification
Focused simulator Go tests/source review; rendered status check is MEM-008.

## Dependencies
MEM-006.

## Sizing
Implementation 1; environment 0; behavior 2; verification 1; decision/recovery 0. Total 4 — tightly coupled runtime-status adjustment.