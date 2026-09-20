# UMIG-EM-003 — CODE: Complete shared Go configuration writer serialization

Status: QUEUED — human-approved 2026-09-20; original sizing 6 split into UMIG-EM-003-L primitive, UMIG-EM-003-S Simulator wiring and this final Replicator/integration task.
Stage / owner: CODE / ChatGPT.
Previous: UMIG-EM-003-S.
Next: UMIG-EM-003-T (QUEUED TEST; experimental OpenCode JR trial only when new exact packet is activated).

## Primary outcome
Wire the Replicator's outermost config transactions to the SAME shared Linux file lock as Simulator and audit all current Go configuration mutation entry points for one lock across the complete load/compare/compose/validate/write/restart-request/ack boundary. The future shared MMA manager must use the same API and is not implemented here. Preserve all foreign reservations and standalone Windows Electron behavior.

## Acceptance (CODE source checkpoint; not a TEST PASS)
1. Protect Replicator manager boot/apply and legacy compose/RunOnce writers without nested lock reacquisition, retain bounded timeout/recovery semantics and never perform runtime polling/network I/O under the lock except necessary restart acknowledgment/readiness. Concurrent conflicting applies must serialize and fail closed; no stale ownership may silently succeed.
2. Cross-read Simulator and Replicator call paths: each current authoritative config write participates in ONE shared transaction; validate before mutation, preserve foreign reservations and detect committed-but-unacknowledged restart truthfully. No changes to installed/operator data or production.
3. Add focused Linux concurrent/conflict/timeout/crash regression test code, read back the exact product diff, and commit a source-only checkpoint. Do not claim independent Go test or live verification before JR. If broader contract requires architectural changes, STOP and rescope before unsafe partial wiring.

Non-scope: new MMA manager/UI/endpoint, live/production Docker, RBE or access-events exposure, legacy launcher cutover, Electron changes. Dependencies: complete archived L and S; `replicator/manager.go`, `replicator/compose_document.go`, `replicator/cycle.go`, existing request/ack path and shared `mma2composer`. Sizing 2/0/1/1/1=5, tightly coupled one Replicator transaction plus final shared-path audit. The only independent test authority is separately ACTIVE UMIG-EM-003-T. ICC only edited by BLACK SHEEP WALL.
