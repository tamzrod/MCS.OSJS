# FMT approved execution queue

Approval: human explicitly authorized promotion and activation of all FMT-001–FMT-008 on 2026-09-22. This is a continuous queue authorization, NOT eight simultaneous ACTIVE tasks. Operation CWAL selects and checkpoints exactly one eligible packet at a time, executes it, records terminal outcome and rescans. Historical OTR packets and OTR_PROMOTION_QUEUE.md are retired, non-executable and excluded even if physically present. Never use historical OTR reports as verified completion.

| ID | Packet | Blocker | Initial status |
| --- | --- | --- | --- |
| FMT-001 | fmt-001-toolkit-source-reconciliation.md | NONE | PENDING |
| FMT-002 | fmt-002-electron-ui-inventory.md | NONE | PENDING |
| FMT-003 | fmt-003-osjs-toolkit-ui-inventory.md | NONE | PENDING |
| FMT-004 | fmt-004-backend-boundary-inventory.md | NONE | PENDING |
| FMT-005 | fmt-005-build-baseline.md | FMT-001 | PENDING |
| FMT-006 | fmt-006-ui-gap-map.md | FMT-001, FMT-002, FMT-003 | PENDING |
| FMT-007 | fmt-007-backend-gap-map.md | FMT-001, FMT-004 | PENDING |
| FMT-008 | fmt-008-gap-task-planning.md | FMT-005, FMT-006, FMT-007 | PENDING |

Before running, verify blocker fields against each packet; packet fields govern if this initial summary differs, and record discrepancy rather than guessing. Source planning copies are historical planning references once promoted. Each task's scope and delivery gates remain binding. No product changes, test PASS, archive completion or main push is implied by promotion. A runner with insufficient role/permissions records obstacle and proceeds only with other eligible tasks. Do not mark queue COMPLETE while unresolved work remains.
