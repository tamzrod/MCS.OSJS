# SIM-013 — Compose Simulator Entry Into Shared MMA2 Config

Status: ACTIVE — promoted for sequential execution after SIM-012.

## Primary outcome

Safely create or modify Simulator-owned entries in the shared MMA2 configuration from a valid Simulator definition without damaging entries owned by other programs.

## Scope

- Load current shared MMA2 configuration.
- Translate enabled Simulator definitions into Simulator-owned MMA2 entries.
- Apply existing ownership/collision rules.
- Build and validate the complete candidate shared MMA2 configuration before commit.
- Commit only a valid candidate, using an atomic/safe write path.
- Preserve foreign/non-Simulator entries.

## Failure contract

If Simulator validation, ownership/collision validation, or complete MMA2 candidate validation fails: shared MMA2 config remains unchanged and no later restart action is authorized.

## Non-scope

- Do not send MMA2 restart in this task.
- Do not own/start/stop MMA2.
- Do not start scheduler execution.

## Acceptance criteria

1. A valid Simulator definition produces the expected Simulator-owned MMA2 entry.
2. Updating the Simulator changes only its owned entry/reservation.
3. Foreign MMA2 entries survive unchanged.
4. Ownership collision is rejected without modifying shared config.
5. Invalid complete candidate config is rejected without modifying shared config.

## Verification

Use fixture/shared configs containing both Simulator and foreign entries. Compare pre/post config and prove validation/collision failures leave the prior file intact.

## Dependencies

SIM-012.

## Sizing

Implementation 1, environment 0, behavioral 2, verification 1, decision/recovery 1 = 5. Tightly coupled around one deterministic compose/validate/commit transaction.