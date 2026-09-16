# MEM-004 — Allow None Simulation in Backend

Status: QUEUED
Previous: MEM-003
Next: MEM-005

## Primary Outcome
Allow an allocated memory area to exist without an active random generator by accepting a zero random interval.

## Scope
- Remove the validation rule that requires a positive random interval whenever area Count is nonzero.
- Preserve all address/port/unit validation.
- Preserve scheduler behavior where interval zero means no schedule is armed.
- Add or update focused Go test coverage for configured memory with interval zero.

## Non-Scope
- No new generator types.
- No data model migration.
- No UI selector yet.

## Acceptance Criteria
1. A memory area with Count > 0 and interval 0 validates.
2. The area remains part of MMA2 structural configuration.
3. The random scheduler does not generate values for interval 0.
4. Existing positive intervals still randomize normally.

## Verification
Focused simulator Go tests plus source re-read. Windows rendering is not required for this backend task.

## Dependencies
MEM-003.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 0. Total 3.