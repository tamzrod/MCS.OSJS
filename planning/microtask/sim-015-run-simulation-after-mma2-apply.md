# SIM-015 — Run Simulation After Successful MMA2 Apply

Status: planning material only. Human promotion is required before execution.

## Primary outcome

Start or update Simulator schedules only after the shared MMA2 config is committed, MMA2 restart succeeds, and MMA2 is ready; feed generated values through existing Raw Ingest into MMA2 memory.

## Required sequence

`Save & Apply -> valid shared config commit -> MMA2 RESTART -> MMA2 ready -> scheduler start/update -> Raw Ingest -> MMA2 memory`.

## Scope

- Reuse the existing scheduler/random generation machinery.
- Reuse the existing MMA2 Raw Ingest client/data path.
- Arm/replace schedules only after successful MMA2 apply/readiness.
- Surface scheduler/raw-ingest failures truthfully.

## Non-scope

- Do not own MMA2 lifecycle.
- Do not add another config/control API.
- Do not implement system-boot restoration yet.

## Acceptance criteria

1. Valid Save & Apply followed by successful MMA2 restart/readiness starts enabled schedules.
2. FC1-FC4 configured schedules advance their Last/Next timing.
3. Generated values are accepted by MMA2 Raw Ingest.
4. Failed MMA2 restart/readiness does not claim the simulation is RUNNING.
5. Raw-ingest failure produces truthful error state.

## Verification

Exercise one enabled four-area Simulator definition and observe scheduler timing plus MMA2 memory updates after successful apply; inject/observe readiness or ingest failure and verify state.

## Dependencies

SIM-014.

## Sizing

Implementation 1, environment 0, behavioral 2, verification 1, decision/recovery 0 = 4. Tightly coupled runtime path with one end-to-end verification workflow.