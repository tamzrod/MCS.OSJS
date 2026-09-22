# FMT-001 — Toolkit source reconciliation
Task ID: FMT-001
Task Name: Reconcile existing Toolkit source against disputed OTR claims
Blocker Task: NONE
Status: PENDING
Assigned Agent: OpenCode — read-only discovery
Stage: DISCOVERY

## Objective
Produce a bounded, source-anchored discrepancy report for the current Toolkit entry point and the two disputed OTR evidence reports. Do not inventory every editor deeply.
## Scope
Read: `OSJS/src/packages/MCSModbusToolkit/index.js`, `OSJS/src/packages/MCSModbusToolkit/memory-contract.js`, `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js`, `workflow/active_work/evidence/otr-002a-report.md`, `workflow/active_work/evidence/otr-003a-map.md`. Read `ICC/INDEX.md` for navigation; treat stale ICC as non-authoritative, do not edit it. Write ONLY `planning/microtask/evidence/fmt-001-source-reconciliation.md` when authorized. No product, historical report, workflow or ICC edits.
## Execution
1. Record exact HEAD, branch and clean/dirty status; inspect only listed source and reports.
2. Compare report claims to actual code and classify each examined claim supported/contradicted/unknown with path and line references.
3. Record the current entry imports, contract operations and unknown runtime behavior; stop without expanding into editor or backend implementation.
## Acceptance Criteria
1. Report identifies current Toolkit imports and both contract operation lists with source lines.
2. Disputed claims are individually classified with source evidence; no report is silently accepted.
3. Runtime/build behavior is marked unverified, not PASS.
## Evidence
`planning/microtask/evidence/fmt-001-source-reconciliation.md`: HEAD, branch/status, source line matrix, contradiction log, unknowns. No invented source paths or screenshots.
## Completion and delivery
COMPLETE only when all three criteria have documented evidence; FAIL for unmet acceptance, BLOCKED for unavailable required source/permission. Only report file may be added and task status updated after separate activation; no push, merge or main authority is granted by this planning packet. Human must specify exact commit/push ref on promotion. STOP after one task.
## Sizing
Implementation 0; environment 0; behavior 0; verification 1; decision/recovery 1 = 2/10. Read-only comparison of five pinned files.