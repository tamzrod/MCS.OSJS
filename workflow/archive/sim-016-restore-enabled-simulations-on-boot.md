# SIM-016 — Restore Enabled Simulations on System Boot

Status: COMPLETED 2026-09-08.

## Completion evidence

- The boot router `NewRuntimeApplyRouter` now loads persisted definitions and performs
  the ownership-safe compose of Simulator reservations into the already-persisted
  shared MMA2 config, preserving SIM-011 independent-lifecycle behavior.
  It waits for the independently auto-started MMA2 to accept every composed
  listener port via `WaitMMA2Ready`, then arms enabled schedules per the latest
  persisted state through the SIM-015 arming gate. Ordinary boot neither
  restarts MMA2 nor writes a restart request; if MMA2 never becomes ready,
  the router survives alive with schedules unarmed so runtime status surfaces
  truthful STOPPED/ERROR state and no restart artifact is created.
- `gofmt -l`, `go vet ./...`, `go test -count=1 ./...` all pass in `simulator/`.
- Two new boot tests in `simulator/boot_restore_test.go`: armed enabled schedules
  resume with advancing FC1-FC4 timing while disabled devices stay unarmed and no
  restart request is written; unavailable MMA2 leaves the router alive with zero
  armed schedulers, no restart request,and truthful STOPPED/ERROR runtime status.
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