# FMT approved execution queue

Human approved promotion and continuous dependency-aware processing of FMT-001–FMT-008 on 2026-09-22. Queue is authorized; at most one packet ACTIVE at a time. Select an eligible packet, checkpoint ACTIVE in packet/queue/handoff, execute within its scope, record outcome, rescan. Never run obsolete OTR packets or infer completion from their evidence.

| ID | Packet | Blocker | Status |
| --- | --- | --- | --- |
| FMT-001 | fmt-001-toolkit-source-reconciliation.md | NONE | PENDING |
| FMT-002 | fmt-002-electron-ui-inventory.md | NONE | PENDING |
| FMT-003 | fmt-003-osjs-toolkit-ui-inventory.md | NONE | PENDING |
| FMT-004 | fmt-004-backend-boundary-inventory.md | NONE | PENDING |
| FMT-005 | fmt-005-build-baseline.md | FMT-001, FMT-003 | PENDING |
| FMT-006 | fmt-006-ui-gap-map.md | FMT-001, FMT-002, FMT-003 | PENDING |
| FMT-007 | fmt-007-backend-gap-map.md | FMT-001, FMT-004 | PENDING |
| FMT-008 | fmt-008-gap-task-planning.md | FMT-005, FMT-006, FMT-007 | PENDING |

Packet blockers and required independently verified evidence are authoritative; if any metadata disagrees, checkpoint BLOCKED for the inconsistent task, never guess or bypass prerequisites. These promoted packets retain their original permitted write scope; promotion grants no product change, independent test PASS, arbitrary push, merge or main push. FMT-005's exact build command, environment, expected result and disposable output still require an explicit separately authorized test instruction before it may run. FMT tasks requiring an independent agent must be handed off to that agent, not impersonated. Old OTR queue and packets are RETIRED/NON-EXECUTABLE even if still present pending safe archival. Stop when no eligible tasks remain and report all unresolved tasks and reasons.
