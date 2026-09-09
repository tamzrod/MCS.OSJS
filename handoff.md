# Handoff

## Status

ACTIVE -- 2026-09-08. SIM-021 was promoted into `workflow/active_work/` by human authorization as five ordered microtasks: SIM-021A → SIM-021E. SIM-018 through SIM-020 remain complete+archived; SIM-022 remains retired (not completed) into archive by human decision. JR resumes authorized execution per Operation CWAL until Active Work drains.

## Promoted Active Sequence (execution authority;( ordered by dependencies(

1. SIM-021A — Define Operator Runtime Status Semantics.
2. SIM-021B — Lock the Runtime Status RPC Contract。(dep: SIM-021A、 SIM-019(
3. SIM-021C — Poll Selected-Device Runtime Status。(dep: SIM-021B(
4. SIM-021D — Render the Compact MMA2 + Simulator Status Row。(dep: SIM-021C(
5. SIM-021E — Surface Runtime Status Errors Without Polling Noise。(dep: SIM-021C、 SIM-021D(

Placement context (human brainstorm, non-authoritative(: display **MMA2 status** and **Simulator status** in-window, minimalist preferred, recorded in `planning/Brainstorm/modbus-simulator-status-display.md`; SIM-021D chooses one compact row below the `Device Definition` heading.

## Completed Predecessor Work

Prior MMA2 work: SIM-001 through SIM-008 are complete and verified. SIM-009 is superseded by the corrected architecture sequence. SIM-010 through SIM-017 are complete, archived,and end with SIM-017 end-to-end verification. SIM-018 through SIM-020 are complete, archived:(SIM-018 decided the local UI-to-runtime boundary; SIM-019 hosted the long-lived local runtime; SIM-020 connected Save & Apply to real runtime.) SIM-021 are now promoted as five ordered microtasks (SIM-021A→SIM-021E(, restoring the runtime-status work to authorized Active Work; SIM-022 (visible end-to-end verification( remains retired-uncompleted by the earlier clearance..

Current archive inventory: SIM-001 through SIM-020 + SIM-022 archived (SIM-009 retained as superseded record; the SIM-021 planning files were promoted out of planning into active work(;`workflow/active_work/` contains the five SIM-021A…E active task files.

## Architecture Boundary

MMA2 owns its own boot/start lifecycle and auto-starts on system boot. The Simulator does not own MMA2 and may not START, STOP, SPAWN, KILL, or REPLACE it. The only MMA2 lifecycle/control action the Simulator may request is RESTART after successfully committing a valid shared MMA2 configuration change. Simulator-generated values enter MMA2 through Raw Ingest. There is no `simbridge` serviceand no Simulator configuration API in the target architecture.

## Execution Rule

JR executes only work present in `workflow/active_work/` and reflected here. Execute the promoted sequence strictly in dependency order. Do not inspect Planning to choose or widen work.

## Promotion Rule

Wheneverthe Active Workload changes, update this file in the same promotion/change so execution context cannot drift.
