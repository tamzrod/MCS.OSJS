# Handoff

## Current

No current ACTIVE task.

SIM-024 completed: the deployed stack independently supervises MMA2, acknowledges each exact restart request, reloads the committed config once, and returns the selected Simulator to `RUNNING` after accepted Raw Ingest.

## Authorized Sequence

SIM-024

## Continuation

The authorized SIM-024 sequence is completed and archived.

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.
