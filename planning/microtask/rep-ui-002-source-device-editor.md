# REP-UI-002 — Replicator Source Device Editor

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Add the first-version source device editor to the selected-device pane.

## Scope
- Fields: Name, Enabled, Endpoint, Source Unit ID, FC, Start, Count, Scan Rate.
- Keep one source range per device for this version.
- Keep Scan Rate device-wide.
- Use controls and validation presentation consistent with the existing OS.js application style.

## Non-Scope
- No multiple read blocks.
- No per-block scan rates.
- No destination allocation.
- No persistence/apply behavior.
- No runtime polling changes.

## Acceptance Criteria
1. Selecting/creating a device exposes all approved source fields.
2. FC, numeric fields, enabled state, and endpoint can be edited in the UI.
3. UI does not expose unsupported multiple-range or per-block scan-rate behavior.

## Verification
Open Replicator, create/select a device, edit every source field, and verify the UI retains the in-window edited state without requiring backend persistence.

## Dependencies
REP-UI-001.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 0. Total 3 — good bounded task.
