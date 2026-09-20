# UMIG-EM-003-S — CODE: Simulator shared-lock transaction boundary

Status: ACTIVE — authorized successor after `UMIG-EM-003-L` source-only completion and archive on 2026-09-20. Source implementation underway; NO independent test PASS.
Stage / owner: CODE / ChatGPT.
Previous: UMIG-EM-003-L (archived COMPLETE / SOURCE ONLY, no TEST PASS).
Next: UMIG-EM-003 (QUEUED CODE: Replicator integration and audit).

## Primary outcome
Ensure all relevant Linux Simulator configuration writers acquire the `mma2composer.WithWriterLock` shared filesystem lock ONCE at an outermost transaction boundary spanning load/compare, compose/validate, disk writes and restart request/ack/readiness, not only `Composer.Commit`. Never nest the lock through `ComposeDocument`, `SaveOne` or composer helpers. Protect boot restore and independently accessible legacy writer entry points; preserve current standalone Windows behavior.

## Acceptance (source-only CODE)
1. Inventory Simulator writer entry points before changing code and route them through one non-nested Linux transaction boundary. Lock timeout/unavailable => no mutation; lock release on failure. A structural apply holds through ack/readiness; timing-only path stays outside when it does not edit MMA2.
2. Maintain effective config/owner, Simulator device document and restart semantics; distinguish committed but unacknowledged restart from success (no invented rollback/green). No host/production operations, new network endpoints, or broad composer redesign.
3. Read back precise changed paths, create focused regression cases for independent JR's future TEST, record source checkpoint and scope limits without claiming TEST or VERIFY PASS.

Non-scope: Replicator wiring, shared manager/UI, Electron, installed configuration and Docker. If direct legacy callers cannot be protected without nested locking or a new API contract, stop and rescope rather than claiming complete coverage. Dependencies: archived L lock primitive; `simulator/mma2_config.go`, `simulator/apply.go`, `simulator/store.go`, `simulator/restart.go` and applicable runtime callers. Sizing 2/0/1/1/1=5, tightly coupled one Simulator transaction. ICC is stale and only BLACK SHEEP WALL may edit it.
