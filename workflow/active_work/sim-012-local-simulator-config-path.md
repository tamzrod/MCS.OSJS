# SIM-012 — Establish Local Simulator Configuration Path

Status: ACTIVE — promoted for sequential execution after SIM-011.

## Primary outcome

Allow the OS.js Modbus Simulator to create, load, edit, and save Simulator-owned configuration locally without `simbridge` and without a Simulator configuration HTTP API.

## Scope

- Rewire the existing Simulator UI to the appropriate local MCS.OSJS persistence/runtime mechanism.
- Preserve the existing Simulator device/config domain model and exact field round-trip semantics.
- Keep Simulator config separate from the shared MMA2 config.

## Non-scope

- Do not modify shared MMA2 config.
- Do not restart MMA2.
- Do not start simulation schedules as part of Save & Apply.
- Do not introduce a replacement network bridge/API.

## Acceptance criteria

1. Open Simulator and load persisted Simulator definitions.
2. Create/edit/duplicate/delete/discard/save definitions through the UI.
3. Reopen/reload and prove exact persisted field round-trip.
4. No Simulator config HTTP service is required.
5. Saving Simulator config alone does not modify MMA2 config or lifecycle.

## Verification

Exercise the local UI/persistence workflow and affected automated tests. Confirm persisted Simulator document content matches the edited definitions.

## Dependencies

SIM-011.

## Sizing

Implementation 1, environment 1, behavioral 1, verification 1, decision/recovery 1 = 5. Keep bounded to one persistence integration path; split further if local OS.js integration requires independent environment discovery.