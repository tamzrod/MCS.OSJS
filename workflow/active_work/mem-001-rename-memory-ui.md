# MEM-001 — Rename Memory UI

Status: ACTIVE
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
Re-read markup/renderer for static labels and selectors. Windows rendered acceptance is MEM-008.

## Dependencies
None.

## Sizing
Surface 1, environment 0, behavior 0, verification 1, recovery 0 = 2.
