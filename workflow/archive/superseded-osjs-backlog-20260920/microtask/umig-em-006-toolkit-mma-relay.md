> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# UMIG-EM-006 — CODE: Authenticated Toolkit shared-MMA relay
Status: PLANNED; explicit human architecture/sequence approval required. Stage/owner CODE / ChatGPT. Previous: UMIG-EM-005-V. Next: UMIG-EM-006-T.

Outcome: expose the existing Go v1 `mma-load`/`mma-apply` ONLY through Toolkit's authenticated local WebSocket/Unix routing and a separate scoped client contract. Identify supported OS.js session/capability API first; fail closed if manage-MMA role cannot be established, never treat `_osjs_client` as edit authorization.

Acceptance: (1) load requires authenticated session, apply requires explicit server-side manage-MMA capability; (2) strict service operation whitelist and correlated v1 envelope/revision/typed errors, no cross-route; (3) one apply dispatch and no file/HTTP/browser direct config writes. Do not add UI or expose RBE/listeners. Limit Toolkit relay/contract/tests, checkpoint diff/readback only, no test PASS. Size 2/1/1/1/0=5; split authorization discovery if source capability unavailable.
