# Handoff

## Current

SIM-023B — Pause Status Polling During Save & Apply — ACTIVE.

## Authorized Sequence

SIM-023A → SIM-023B → SIM-023C

- SIM-023A COMPLETED 2026-09-09 (archived) — poll target resolved from persisted/applied identity; new/renamed-unapplied rows expose no poll target.
- SIM-023B is ACTIVE.
- SIM-023C is QUEUED and already human-authorized.
- On verified SIM-023B completion, advance only its explicit `Next` task (SIM-023C).

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.
