# UMIG-EM-003-V — VERIFY: Cross-process writer serialization

Status: ACTIVE — independent OpenHands JR, sandbox preparation authorized; live concurrency execution pending exact packet.
Stage / owner: VERIFY / OpenHands JR. Previous: UMIG-EM-003-R-T archived COMPLETE/PASS at 358a52d4f0f3b2aa5c924d97d3613a8d032c8ba0. Next: UMIG-EM-004 PLANNED; separate promotion required.

Primary outcome: independently demonstrate cross-process serialization for simultaneous synthetic Simulator and Replicator config applies: no lost update, preserved foreign owner, coherent config/restart acknowledgement.

Target: existing disposable OpenHands environment; no separate sandbox or new clone requirement. Prepare only through `workflow/cwal/umig-em-003-v-sandbox-preflight.md`: a single clean fast-forward when necessary; one sandbox-local `sudo -n dockerd` startup if absent; then confirm local daemon, verify-only Compose topology, read-only inventories and exact local apply/restart interfaces. Docker startup is sandbox tooling preparation, NOT product execution. If prep fails, report one blocker and STOP. No operator/production host, services, ports, data, Docker contexts or volumes.

Live VERIFY remains gated: coding agent must issue a separate source-pinned exact packet defining a new uniquely named Docker project and data volume, synthetic-only request bodies/endpoints, at most two applies each, timing and bounded timeout, before/after config and owner hashes, restart request/ack evidence, and label-scoped cleanup `down --remove-orphans` without `-v` while retaining only its named volume. No Compose startup, synthetic apply, live test, report push or product/source/workflow edits until that packet is activated. Go race TEST PASS is not live VERIFY. JR returns raw evidence; coding agent reviews/archives; no automatic successor.

Acceptance: (1) isolation confirmed with exact resources and bounded timestamped requests; (2) no lost update or foreign-owner mutation, truthful serialized restart chain; (3) document config/owner hashes and label-safe project cleanup, no unrelated changes. Contradiction FAIL; environment unavailable BLOCKED. No production action, RBE output or public Modbus bind.

Current handoff authority: one-pass PREPARATION ONLY. OpenCode `UMIG-EM-004` remains PLANNED.
