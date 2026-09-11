# REP-BLOCK-003 — Device / Pull Blocks Folder Tabs

## Primary Outcome
Separate the selected Replicator device's right-side configuration into exactly two classic folder-shaped top-level tabs: **Device** and **Pull Blocks**, while leaving the existing left device tree/list unchanged.

## Required Layout

```text
┌───────────────────┬───────────────────────────────────────────────┐
│ DEVICES           │      ___________       ______________        │
│                   │     /  Device   \_____/  Pull Blocks \       │
│ Search...         │    /_____________\   /________________\      │
│                   │    │                                         │
│ ▸ Rep-PLC-1       │    │   selected main-tab content             │
│ ▸ Rep-PLC-2       │    │                                         │
│ ▸ Rep-PLC-3       │    │                                         │
│                   │    │                                         │
│                   │    │                                         │
│                   │    └─────────────────────────────────────────│
│                   │                                              │
│ Add Duplicate Del │                 Save & Apply   Discard        │
└───────────────────┴───────────────────────────────────────────────┘
```

Hierarchy:

```text
Replicator window
├── LEFT: existing device tree/list — unchanged
└── RIGHT: selected device configuration
    ├── TAB 1: Device
    │   ├── Name / Enabled
    │   ├── Endpoint / Source Unit ID
    │   └── Destination / Ownership
    └── TAB 2: Pull Blocks
        ├── Block 1 — FC / Start / Count / Scan Rate / Status
        ├── Block 2 — FC / Start / Count / Scan Rate / Status
        └── ...
```

## Scope
- Keep the complete existing left device tree/list, search, selection, Add, Duplicate, and Delete behavior as-is.
- Add exactly two top-level folder-shaped tabs to the **right-side selected-device configuration area**: `Device` and `Pull Blocks`.
- Tabs visually resemble classic physical file-folder tabs, not modern flat/browser/pill tabs.
- The active tab visually joins its content panel; inactive tab remains visibly behind/unselected.
- `Device` tab contains device-level fields: Name, Enabled, Endpoint, Source Unit ID, Destination Port, Destination Unit ID, Auto controls, Owner, and ownership Status.
- `Pull Blocks` tab contains the device's Pull Block collection.
- Pull Blocks themselves are **not tabs**. Render them as rows/cards/list entries within the Pull Blocks tab.
- Each block exposes FC, Start, Count, Scan Rate, and that block's runtime status.
- Add Block creates a new block with safe defaults.
- Duplicate Block copies only the selected block.
- Delete Block removes only the selected block.
- Save & Apply and Discard remain device-level actions and operate on the complete selected device, including the full Pull Block collection.
- Keep styling consistent with the existing classic MCS.OSJS workstation/window appearance.

## Non-Scope
- No change to the left device tree/list layout or behavior.
- No tab per Pull Block.
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
4. Multiple Pull Blocks render inside the Pull Blocks tab as entries/rows/cards, not as another tab layer.
5. Add/Duplicate/Delete Block affect only the Pull Block collection and do not alter the left device tree or device-level fields.
6. Different blocks retain independent FC/start/count/scan-rate values after Save & Apply and reopen.
7. Runtime status for each block corresponds to that block's poller.
8. Save & Apply persists the complete device and Pull Block collection through the existing authoritative apply path.

## Verification
Rendered OS.js UI test: first compare the left device tree/list before and after the change and verify it is unchanged. Select one device and verify the right side exposes exactly two classic folder tabs, `Device` and `Pull Blocks`. Verify device fields under Device; create at least two block entries under Pull Blocks with different scan rates; Save & Apply; reopen; verify both blocks and their per-block statuses; delete one block and verify the remaining poller is unaffected.

## Dependencies
REP-BLOCK-002 — Independent Pull Block Pollers.

## Sizing
Implementation surface 1; environment 0; behavior 2; verification 1; decision/recovery 1. Total 5 — tightly coupled right-side UI workflow, deliberately separated from backend multi-block runtime work.
