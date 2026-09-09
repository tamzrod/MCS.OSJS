# Handoff

## Current

ACTIVE: SIM-021E — Surface Runtime Status Errors

## Authorized Sequence

SIM-021A → SIM-021B → SIM-021C → SIM-021D → SIM-021E

SIM-021A through SIM-021D are completed and archived. SIM-021E is `ACTIVE` and already human-authorized.

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.
