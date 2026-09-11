# REP-BLOCK-003 — Device / Pull Blocks Folder Tabs

## Primary Outcome
Separate the selected Replicator device's right-side configuration into exactly two classic folder-shaped top-level tabs: **Device** and **Pull Blocks**, while leaving the existing left device tree/list unchanged. The Pull Blocks tab must use a compact spreadsheet-style editor rather than large per-block cards.

## Required Layout

```text
┌───────────────────┬──────────────────────────────────────────────────────┐
│ DEVICES           │      ___________       ______________               │
│                   │     /  Device   \_____/  Pull Blocks \              │
│ Search...         │    /_____________\   /________________\             │
│                   │    │                                                │
│ ▸ Rep-PLC-1       │    │ [Add] [Duplicate] [Delete]                     │
│ ▸ Rep-PLC-2       │    │                                                │
│ ▸ Rep-PLC-3       │    │ # | FC | Start | Count | Scan | Status | Poll │
│                   │    │ 1 | 3  |   0   |  16   | 1000 | OK     | ...  │
│                   │    │ 2 | 4  |   0   |  16   | 1000 | OK     | ...  │
│                   │    │ 3 | 3  |  20   |   6   |  300 | ERROR  | ...  │
│                   │    │                                                │
│ Add Duplicate Del │    │                  Save & Apply   Discard         │
└───────────────────┴──────────────────────────────────────────────────────┘
```

Hierarchy:

```text
Replicator window
├── LEFT: existing device tree/list — unchanged
└── RIGHT: selected device configuration
    ├── TAB 1: Device
    │   ├── Name / Enabled
    │   ├── Endpoint / Source Unit ID
    │   └── Destination / Ownership / Operational Status
    └── TAB 2: Pull Blocks
        └── compact spreadsheet/grid of Pull Blocks
```

## Scope
- Keep the complete existing left device tree/list, search, selection, Add, Duplicate, and Delete behavior as-is.
- Add exactly two top-level folder-shaped tabs to the **right-side selected-device configuration area**: `Device` and `Pull Blocks`.
- Tabs visually resemble classic physical file-folder tabs, not modern flat/browser/pill tabs.
- The active tab visually joins its content panel; inactive tab remains visibly behind/unselected.
- `Device` tab contains device-level fields: Name, Enabled, Endpoint, Source Unit ID, Destination Port, Destination Unit ID, Auto controls, Owner, and operational Status.
- `Pull Blocks` tab contains the device's Pull Block collection as a **compact spreadsheet-style table/grid**.
- Do not render each Pull Block as a large card, panel, fieldset, or vertically stacked form section.
- One Pull Block = one row.
- Use fixed/compact columns for at least: Block/#, FC, Start, Count, Scan Rate, Source/Status, Last Poll. Last Error may be shown in-row when space permits or in a compact detail/status area for the selected row.
- FC, Start, Count, and Scan Rate must be directly editable from the row using compact controls.
- Row selection is the authoritative selected Pull Block for Duplicate/Delete actions.
- Add Block creates a new row with safe defaults.
- Duplicate Block copies only the selected row/block.
- Delete Block removes only the selected row/block.
- Per-block runtime status must remain independently visible without adding a second full-width status panel beneath every row.
- Save & Apply and Discard remain device-level actions and operate on the complete selected device, including the full Pull Block collection.
- Keep styling consistent with the existing classic MCS.OSJS workstation/window appearance.
- Prioritize information density so several Pull Blocks can be viewed simultaneously without unnecessary vertical scrolling.

## Non-Scope
- No change to the left device tree/list layout or behavior.
- No tab per Pull Block.
- No large block cards or vertically repeated status panels.
- No change to MMA2 ownership semantics.
- No per-block destination port/unit.
- No per-block source endpoint or Unit ID.
- No drag/drop block reordering.
- No charts/history/metrics.
- No modern pill/rounded browser-style tabs.

## Acceptance Criteria
1. Existing left device tree/list remains visually and behaviorally unchanged.
2. Right-side configuration has exactly two main folder-shaped tabs: `Device` and `Pull Blocks`.
3. Device-level fields appear only in the Device tab; Pull Block definitions appear in the Pull Blocks tab.
4. Pull Blocks render as a compact spreadsheet/table with one block per row, not as cards/panels and not as another tab layer.
5. At least several block rows are simultaneously visible at normal window size without each row consuming large vertical space.
6. FC/start/count/scan-rate are editable directly in each row.
7. Add/Duplicate/Delete Block affect only the Pull Block collection and act on the selected row.
8. Different blocks retain independent FC/start/count/scan-rate values after Save & Apply and reopen.
9. Runtime status and Last Poll for each row correspond to that block's poller while remaining compact.
10. Save & Apply persists the complete device and Pull Block collection through the existing authoritative apply path.

## Verification
Rendered OS.js UI test: first compare the left device tree/list before and after the change and verify it is unchanged. Select one device and verify the right side exposes exactly two classic folder tabs, `Device` and `Pull Blocks`. In Pull Blocks, verify blocks appear as compact spreadsheet rows rather than cards; create enough blocks to prove multiple rows remain visible simultaneously; edit FC/start/count/scan rate inline; select rows and exercise Add/Duplicate/Delete; Save & Apply; reopen; verify values/order/status persist and deleting one row removes only that poller.

## Dependencies
REP-BLOCK-002 — Independent Pull Block Pollers.

## Sizing
Implementation surface 1; environment 0; behavior 2; verification 1; decision/recovery 1. Total 5 — tightly coupled right-side UI workflow, deliberately separated from backend multi-block runtime work.
