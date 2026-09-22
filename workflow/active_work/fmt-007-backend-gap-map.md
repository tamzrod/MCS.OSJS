# FMT-007 — Backend gap map
Task ID: FMT-007
Task Name: Reconcile Toolkit client operations with server dispatch
Blocker Task: FMT-001, FMT-004
Status: PENDING
Assigned Agent: OpenCode — read-only design
Stage: DESIGN

## Objective
Map declared Memory/Replicator client operations to actual Toolkit server dispatch, distinguishing unsupported/unknown runtime operations; no full runtime audit.
## Scope
Read `ICC/INDEX.md`, accepted FMT-001/004 reports and their pinned Toolkit contract, transport and server files. Write only `planning/microtask/evidence/fmt-007-client-server-gap-map.md` on activation. No runtime/service/device, production, ICC or historical report changes.
## Execution
1. Confirm blocker evidence and relevant source SHA.
2. Compare each declared client operation with actual transport/server dispatch; classify present, absent or runtime-unknown with source lines.
3. Record which runtime-specific follow-up discovery is required for write/restart/status safety; do not infer implementation from contract names.
## Acceptance Criteria
1. Each declared client operation has source-linked dispatch classification.
2. Unsupported and runtime-unknown operations are distinct; no invented endpoint or status PASS.
3. Further runtime discovery is proposed as separate small tasks with explicit boundaries.
## Evidence
`planning/microtask/evidence/fmt-007-client-server-gap-map.md`: SHA, operation/dispatch/unknown matrix, follow-up discovery candidates.
## Completion and delivery
COMPLETE only with all outcomes evidenced; FAIL if incomplete, BLOCKED if prerequisites invalid. Only report and authorized status may change. No commit/push authority until exact ref/command on promotion. STOP.
## Sizing
Implementation 0; environment 0; behavior 1; verification 1; decision/recovery 1 = 3/10. Toolkit client/server boundary only.