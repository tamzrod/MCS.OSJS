# Handoff

## Current

SIM-023A — Poll Only the Applied Device Identity — ACTIVE.

## Authorized Sequence

SIM-023A → SIM-023B → SIM-023C

- SIM-023A is ACTIVE.
- SIM-023B and SIM-023C are QUEUED and already human-authorized.
- On verified SIM-023A completion, advance only its explicit `Next` task.

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.
