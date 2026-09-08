# SIM-016 — Restore Enabled Simulations on System Boot

Status: planning material only. Human promotion is required before execution.

## Primary outcome

Restore enabled persisted simulations after system/appliance boot while preserving MMA2's independent auto-start lifecycle.

## Boot contract

System boot starts MMA2 independently. MMA2 reads the already-persisted shared config. The Simulator starts separately, loads persisted Simulator config, waits for MMA2 readiness, then runs enabled schedules.

Ordinary boot must not cause the Simulator to restart MMA2 merely because the Simulator started.

## Scope

- Load persisted Simulator definitions at Simulator runtime startup.
- Detect/wait for independently auto-started MMA2 readiness.
- Arm enabled schedules once MMA2 is ready.
- Resume Raw Ingest through the existing data path.
- Surface unavailable MMA2 without attempting to start it.

## Non-scope

- Do not send restart on ordinary boot when shared config has not changed.
- Do not start/stop/spawn MMA2.
- Do not modify shared MMA2 config merely to restore schedules.

## Acceptance criteria

1. Full appliance reboot auto-starts MMA2 independently.
2. Simulator restores enabled persisted definitions without opening the UI or pressing Save & Apply.
3. Enabled schedules resume after MMA2 becomes ready.
4. Simulator does not issue MMA2 restart during an unchanged ordinary boot.
5. If MMA2 is unavailable, Simulator reports/waits rather than starting MMA2.

## Verification

Persist an enabled simulation, reboot the appliance/runtime, and prove MMA2 starts independently and simulation values resume through Raw Ingest without manual UI action.

## Dependencies

SIM-015.

## Sizing

Implementation 1, environment 1, behavioral 1, verification 1, decision/recovery 0 = 4. One boot-restoration workflow.