# FMT-004 — Backend boundary inventory
Task ID: FMT-004
Task Name: Inventory Toolkit server-to-runtime operation boundaries
Blocker Task: NONE
Status: PENDING
Assigned Agent: OpenCode — read-only discovery
Stage: DISCOVERY

## Objective
Map only the existing Toolkit server and client transport entry points to their named runtime destinations; defer runtime internals to later source-pinned packets.
## Scope
Read `ICC/INDEX.md`, `OSJS/src/packages/MCSModbusToolkit/server.js`, `memory-transport.js`, `replicator-transport.js`, `memory-contract.js`, `replicator-contract.js` and direct server imports if present. Verify files exist; missing paths are findings, not permission to invent replacements. Write only `planning/microtask/evidence/fmt-004-backend-boundaries.md` after activation. No backend/runtime/device/service access, no product/ICC/workflow edits.
## Execution
1. Record HEAD/branch/status and inspect only Toolkit transport/server/contract entry points.
2. Map declared operation, actual dispatch, destination and known side-effect boundary with source lines.
3. List unknown runtime internals and propose separate per-runtime discovery, without asserting endpoint support.
## Acceptance Criteria
1. Contract declarations versus server dispatch are distinguished with exact source references.
2. Destination paths and side-effect/authorization unknowns are explicitly recorded.
3. Runtime-specific deeper work is identified as separate candidate tasks, not performed here.
## Evidence
`planning/microtask/evidence/fmt-004-backend-boundaries.md`: pinned SHA, operation/dispatch table, exact lines and unknowns.
## Completion and delivery
COMPLETE only with all outcomes evidenced; FAIL for missing outcomes; BLOCKED for inaccessible source. Report and authorized status only; no commit/push authority until exact ref/command supplied on promotion. STOP.
## Sizing
Implementation 0; environment 0; behavior 0; verification 1; decision/recovery 1 = 2/10. Explicitly excludes full MMA/simulator/replicator runtime audit.