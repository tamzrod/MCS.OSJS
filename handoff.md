# Handoff

## Current

ACTIVE: REP-001 — Share MMA2 Reservation Composer

## Authorized Sequence

REP-001 → REP-002 → REP-003 → REP-004 → REP-005 → REP-006 → REP-007

## Continuation

REP-001 is `ACTIVE`. REP-002 through REP-007 are `QUEUED` and already human-authorized; they require no additional approval when advanced through their explicit `Previous` / `Next` links.

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.
