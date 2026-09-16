# MEM-001 — Rename Memory UI

Status: COMPLETED — static verification passed 2026-09-17
Previous: none
Next: MEM-002

## Primary outcome
Rename existing Electron Simulator presentation to Memory without changing internal services or persisted documents.

## Scope
Rename Simulator navigation label to Memory, Simulator Devices heading to Devices, Function table header to Area. Preserve `data-tab="simulator"`, panel IDs, existing controls/actions and backend wiring.

## Non-scope
No API, config, service, ownership or OS.js application renames.

## Acceptance
1. Memory tab, Devices heading and Area column display.
2. Existing FC fields and actions remain.
3. Internal selectors/storage unchanged.

## Verification
Static source verification passed: `electron/renderer/index.html` shows visible Memory while retaining `data-tab="simulator"` and `panel-simulator`; `electron/renderer/app.js` renders `Devices`, `Device Definition`, and `Area` while retaining `simulatorCall`, existing FC fields/actions, and simulator internal identifiers. Windows rendered acceptance remains MEM-008.

## Dependencies
None.

## Sizing
Surface 1, environment 0, behavior 0, verification 1, recovery 0 = 2.
