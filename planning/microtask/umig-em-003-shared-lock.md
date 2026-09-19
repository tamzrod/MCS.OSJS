# UMIG-EM-003 — CODE: Shared Go configuration writer lock
Status: PLANNED; human contract/sequence approval required. Stage/owner CODE / ChatGPT. Previous: UMIG-EM-002-V. Next: UMIG-EM-003-T.

Outcome: one crash-released, bounded Linux lock shared by Simulator, Replicator and future global manager for the WHOLE load/compose/validate/write/restart-request/ack transaction, not just `mma2composer.Commit`. Use agreed host-mounted path and explicit timeout; reject ambiguous lock ownership. Existing standalone Windows Electron process remains untouched.

Acceptance: (1) all Go writers use the same lock and no nested-lock deadlock; (2) concurrent conflicting applies serialize and preserve foreign reservation/config with deterministic failure; (3) crash/timeout never implies success. Scope Go composer/writer callsites and focused tests only, no new UI/protocol, RBE exposure or installed data. Record source checkpoint, diff/readback, no test PASS. Sizing 2/1/1/1/1=6: mandatory implementation split/re-review before promotion if the shared-lock surface cannot be kept one atomic behavior.
