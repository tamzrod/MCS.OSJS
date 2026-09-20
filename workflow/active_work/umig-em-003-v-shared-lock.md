# UMIG-EM-003-V — VERIFY: Cross-process writer serialization

Status: ACTIVE — independent OpenHands JR. CURRENT stage is CWAL preparation only; live concurrency execution awaits a separate exact packet.
Stage / owner: VERIFY / OpenHands JR. Previous: UMIG-EM-003-R-T archived COMPLETE/PASS at 358a52d4f0f3b2aa5c924d97d3613a8d032c8ba0. Next: UMIG-EM-004 PLANNED; separate human promotion required.

Primary outcome: independently demonstrate cross-process serialization for simultaneous synthetic Simulator and Replicator config applies: no lost update, preserved foreign owner, coherent config/restart acknowledgement.

**OpenHands invocation: `OPERATION CWAL` ONLY.** The sole CURRENT `JR TEST TASK` in `handoff.md` defines preparation authority and explicitly incorporates `workflow/cwal/umig-em-003-v-sandbox-preflight.md` as subordinate exact commands. That file is NOT a standalone directive. Use the existing disposable OpenHands environment; no new sandbox/clone requirement. Preparation permits one clean fast-forward when needed, one sandbox-local `sudo -n dockerd` startup if absent, then verify-only topology/inventory/interface reads and a strictly handoff-only report/push. BLOCKED on unsafe/missing prerequisites; no operator resources, source edits or independent runner discretion. Preparation COMPLETE is NOT VERIFY PASS.

Live VERIFY remains gated: coding agent must issue a separate source-pinned executable CWAL packet defining a new uniquely named Docker project and test-only volume, exact synthetic-only request bodies/endpoints, at most two applies per runtime, bounds, before/after config and owner hashes, restart request/ack evidence, and label-scoped `down --remove-orphans` without `-v` while retaining only its test volume. No Compose startup/build, synthetic applies, live tests, product edits, or task advancement until then. Go race TEST PASS does not substitute for live VERIFY. JR reports raw evidence; coding agent reviews and archives; no automatic successor.

Acceptance: (1) isolated resources and bounded timestamped concurrent applies; (2) no lost update/foreign-owner mutation and coherent serialized restart chain; (3) config/owner hashes and label-safe cleanup without unrelated changes. Contradiction FAIL; unavailable environment BLOCKED. No production action, RBE output or public Modbus bind.

OpenCode `UMIG-EM-004` stays PLANNED pending accepted live VERIFY and separate promotion.
