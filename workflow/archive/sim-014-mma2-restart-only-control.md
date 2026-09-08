# SIM-014 — Add MMA2 Restart-Only Control

Status: COMPLETED 2026-09-08 — archived after verification.

## Primary outcome

After a successful shared MMA2 configuration commit, allow the Simulator to request exactly one MMA2 lifecycle/control operation: RESTART.

## Control boundary

Allowed: `RESTART`.

Forbidden: START, STOP, SPAWN, KILL, REPLACE, or process ownership.

MMA2 remains an independently managed appliance component that auto-starts on system boot.

## Scope

- Establish the local appliance mechanism by which the Simulator requests MMA2 restart.
- Issue restart only after SIM-013 has successfully committed a valid shared config.
- Wait/check for MMA2 to return ready after restart.
- Surface restart/readiness failure truthfully to the caller.

## Non-scope

- Do not start MMA2 if it is absent.
- Do not stop MMA2 as an independent Simulator operation.
- Do not start simulation schedules yet.
- Do not add a general MMA2 control API.

## Acceptance criteria

1. Successful config commit causes one MMA2 restart request.
2. Rejected/failed config commit causes no restart request.
3. MMA2 reloads the committed shared configuration and returns ready.
4. Simulator has no other MMA2 lifecycle operation.
5. Restart/readiness failure is reported without claiming successful apply.

## Verification

Run MMA2 independently, apply a valid listener change, observe restart/reload and readiness, then prove invalid/rejected config does not trigger restart.

## Completion evidence

- `simulator/restart.go` defines the sole control request: a fingerprinted RESTART artifact written only after a successful shared-config commit, followed by readiness polling.
- Successful readiness clears the request; timeout is returned truthfully and leaves the request pending. Rejected composition writes no request.
- Restart-focused tests pass. The real SIM-017 harness observes exactly one request after a valid port edit, independently restarts MMA2, and confirms the reloaded listener is ready.
- Static checks find no Simulator start, stop, spawn, kill, replace, or MMA2 process-ownership implementation.

## Dependencies

SIM-011 and SIM-013.

## Sizing

Implementation 1, environment 1, behavioral 1, verification 1, decision/recovery 1 = 5. Tightly bounded to one restart-only control contract.
