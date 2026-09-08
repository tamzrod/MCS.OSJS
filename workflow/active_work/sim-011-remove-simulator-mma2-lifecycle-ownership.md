# SIM-011 — Remove Simulator Ownership of MMA2 Lifecycle

Status: ACTIVE — promoted for sequential execution after SIM-010.

## Primary outcome

Make MMA2 lifecycle independent from the Simulator: MMA2 auto-starts on system boot and the Simulator does not own its process.

## Architecture contract

MMA2 owns its own boot/start lifecycle. The Simulator may never START, STOP, SPAWN, KILL, or REPLACE MMA2. The only lifecycle/control operation the Simulator will eventually be allowed to request is RESTART; restart implementation belongs to SIM-014.

## Scope

- Remove/refactor Simulator `MMA2Lifecycle` behavior that creates or owns an MMA2 child process.
- Remove structural-apply behavior that starts/stops/replaces MMA2.
- Preserve shared MMA2 config composition and ownership validation for later tasks.
- Preserve scheduler/raw-ingest implementation without arming it through MMA2 process ownership.

## Non-scope

- Do not implement restart signaling yet.
- Do not redesign shared MMA2 config.
- Do not implement boot restoration or end-to-end simulation.

## Acceptance criteria

1. Starting the Simulator cannot create an MMA2 process.
2. Stopping the Simulator cannot terminate MMA2.
3. Simulator structural changes cannot replace MMA2.
4. MMA2 is treated as an independently managed appliance runtime.

## Verification

Unit/static verification proving no Simulator runtime path invokes MMA2 start/stop/spawn/replace behavior; affected tests pass.

## Dependencies

SIM-010.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. Good JR-sized boundary correction.