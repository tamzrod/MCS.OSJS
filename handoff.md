# Handoff

## Current

SIM-023C — Render Neutral Status for Unapplied Selection — ACTIVE.

## Authorized Sequence

SIM-023A → SIM-023B → SIM-023C

- SIM-023A COMPLETED 2026-09-09 (archived) — poll target resolved from persisted/applied identity; new/renamed-unapplied rows expose no poll target.
- SIM-023B COMPLETED 2026-09-09 (archived) — polling paused during Save & Apply; in-flight responses dropped; forced refresh resumes after the apply settles.
- SIM-023C is ACTIVE.
- Sequence complete after SIM-023C (no successor).

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.
