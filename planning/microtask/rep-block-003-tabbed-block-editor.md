# REP-BLOCK-003 — Tabbed Pull Block Editor

## Primary Outcome
Present each Replicator Pull Block as its own tab so the operator can add, select, edit, duplicate, and delete independent block pollers inside one Replicator device.

## Scope
- Replace the single visible Pull Block editor with a tab strip bound to the device's Pull Block collection.
- One tab = one Pull Block/poller.
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

## Acceptance Criteria
1. A device with multiple Pull Blocks renders one tab per block and switching tabs shows the correct block values.
2. Add/Duplicate/Delete affect only the block collection and do not alter device-level source/destination fields.
3. Different blocks may retain different FC/start/count/scan-rate values after Save & Apply and reopen.
4. Runtime status shown for the selected tab corresponds to that block rather than a device-wide aggregate.
5. Deleting one block does not stop or remove the remaining block pollers after Save & Apply.

## Verification
Rendered OS.js UI test with one Replicator device containing at least two blocks at different scan rates: switch tabs, edit each independently, Save & Apply, reopen, verify values persist, then verify per-tab runtime status and delete one block without affecting the other.

## Dependencies
REP-BLOCK-002 — Independent Pull Block Pollers.

## Sizing
Implementation surface 1; environment 0; behavior 2; verification 1; decision/recovery 1. Total 5 — tightly coupled UI workflow, deliberately separated from backend multi-block runtime work.
