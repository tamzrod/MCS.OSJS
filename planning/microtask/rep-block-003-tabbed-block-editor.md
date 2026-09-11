# REP-BLOCK-003 — Folder-Tabbed Pull Block Editor

## Primary Outcome
Present each Replicator Pull Block as its own **folder-shaped tab** so the operator can add, select, edit, duplicate, and delete independent block pollers inside one Replicator device.

## Scope
- Replace the single visible Pull Block editor with a tab strip bound to the device's Pull Block collection.
- One folder-shaped tab = one Pull Block/poller.
- Tabs must visually resemble classic physical file-folder tabs rather than modern flat browser tabs: a raised tab with a folder-tab silhouette connected to the content panel.
- The selected tab visually joins the Pull Block content area; inactive tabs remain visibly behind/unselected.
- Keep the styling consistent with the existing classic MCS.OSJS workstation/window appearance; do not introduce a modern rounded/pill tab style.
- Add Block creates a new block with safe defaults.
- Duplicate Block copies the selected block into a new independent block.
- Delete Block removes only the selected block.
- Selecting a tab edits only that block's FC, Start, Count, and Scan Rate.
- Display the selected block's runtime status: RUNNING/STOPPED, Source OK/ERROR/WAITING, Last Poll, Last Error.
- Keep Name, Enabled, Endpoint, Source Unit ID, and Destination outside the tabs because they remain device-level settings.
- Save & Apply submits the full block collection through the existing authoritative ownership/apply path.

## Non-Scope
- No change to MMA2 ownership semantics.
- No per-block destination port/unit.
- No drag/drop tab reordering in this task.
- No per-block source endpoint or Unit ID.
- No charts/history/metrics.
- No modern pill/rounded browser-style tabs.

## Acceptance Criteria
1. A device with multiple Pull Blocks renders one **folder-shaped tab** per block and switching tabs shows the correct block values.
2. The active tab is visually connected to its content panel like a classic file-folder tab; inactive tabs are visually distinct.
3. Add/Duplicate/Delete affect only the block collection and do not alter device-level source/destination fields.
4. Different blocks may retain different FC/start/count/scan-rate values after Save & Apply and reopen.
5. Runtime status shown for the selected tab corresponds to that block rather than a device-wide aggregate.
6. Deleting one block does not stop or remove the remaining block pollers after Save & Apply.

## Verification
Rendered OS.js UI test with one Replicator device containing at least two blocks at different scan rates: first verify the tab strip has the required classic folder-tab silhouette and selected-tab/content-panel connection; then switch tabs, edit each independently, Save & Apply, reopen, verify values persist, verify per-tab runtime status, and delete one block without affecting the other.

## Dependencies
REP-BLOCK-002 — Independent Pull Block Pollers.

## Sizing
Implementation surface 1; environment 0; behavior 2; verification 1; decision/recovery 1. Total 5 — tightly coupled UI workflow, deliberately separated from backend multi-block runtime work.
