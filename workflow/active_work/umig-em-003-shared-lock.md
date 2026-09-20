# UMIG-EM-003 — CODE: Complete shared Go configuration writer serialization

Status: ACTIVE — approved human sequence; Simulator source-only checkpoint `b9d674759b6f1f3ea7d8e58389efb1553075875e` archived 2026-09-20. No JR TEST PASS yet.
Stage / owner: CODE / ChatGPT.
Previous: UMIG-EM-003-S (archived COMPLETE / SOURCE ONLY).
Next: UMIG-EM-003-T (QUEUED TEST; one approved experimental OpenCode JR trial only after exact packet).

## Primary outcome
Wire Replicator outermost config transactions to the SAME shared Linux `mma2composer.WithWriterLock` as Simulator; audit every relevant writer's load/compare/compose/validate/write/restart-request/ack chain. Future MMA manager is not implemented. Foreign reservations and Windows Electron behavior unchanged.

## Acceptance (CODE source checkpoint; not a TEST PASS)
1. Protect Replicator manager boot/apply and legacy compose/RunOnce writers without nested lock reacquisition, bounded timeout and false success; do not run poll/network I/O under the lock except restart acknowledgment/readiness.
2. Cross-read Simulator and Replicator call paths: every current authoritative config write participates in the shared transaction; validate before mutation, preserve foreign reservations, detect committed-but-unacknowledged restart truthfully. No operator or installed data.
3. Add focused Linux concurrent/conflict/timeout/crash regression source, read back exact product diff and commit SOURCE-ONLY checkpoint. If broader contract needs architectural changes, STOP and rescope rather than claim coverage.

Non-scope: shared manager/UI/endpoint, Docker, RBE or access-event exposure, launcher cutover, Electron edits or production. Dependencies: archived L and S; `replicator/manager.go`, `replicator/compose_document.go`, `replicator/cycle.go`, lifecycle and shared composer. Sizing 2/0/1/1/1=5. ICC only BLACK SHEEP WALL.
