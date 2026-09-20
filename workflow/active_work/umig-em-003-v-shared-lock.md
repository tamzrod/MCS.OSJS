# UMIG-EM-003-V — VERIFY: Cross-process writer serialization

Status: ACTIVE — HUMAN-PROMOTED 2026-09-20; EXECUTION BLOCKED awaiting exact safe disposable runtime target and complete CWAL packet.
Stage / owner: VERIFY / independent JR (OpenHands by default; OpenCode only with separate task-specific human authorization).
Previous: UMIG-EM-003-R-T archived COMPLETE/PASS at 358a52d4f0f3b2aa5c924d97d3613a8d032c8ba0. Next: UMIG-EM-004, currently PLANNED; human promotion separately required.

Primary outcome: Demonstrate actual cross-process serialization of synthetic Simulator and Replicator config transactions, with no lost update, correct foreign-owner preservation and coherent restart acknowledgment under bounded concurrent applies.

Scope: NEW independent disposable Docker project and uniquely named volumes/networks; local sandbox daemon provenance; no reuse of running deployment or prior verification project. Confirm exact container labels, volume names, port mappings, isolated data root, synthetic-only addresses and cleanup before any start. Read exact product/source checkpoint and compare live GitHub main before execution. Bounded simultaneous synthetic apply requests with exact request bodies, attempt count, maximum time and expected results to be defined by coding agent after target discovery; capture before/after config and ownership hashes, replies, restart requests/acks and logs. After checks, perform label-scoped cleanup only of this test's resources and retain its named data volume; no `down -v`.

Safety: No production/customer/operator devices or data, no existing Docker project/volume, no public Modbus/RBE/access-events bind, no sudo or unattended permission. Do not run tests until safe target is confirmed and handoff contains executable exact steps, endpoints, timeout, expected outcomes, evidence and cleanup. If sandbox or permissions unavailable, BLOCKED. Do not substitute a unit-test PASS for live VERIFY.

Acceptance (at most 3): (1) verified isolated topology and bounded concurrent applications with timestamped raw evidence; (2) no lost update, foreign reservations unchanged, serialized config/restart-ack chain per apply, with contradictory observations FAIL; (3) original synthetic state hashes and label-verified cleanup/retained volume documented with no unrelated resource changes.

Evidence: independent raw terminal commands, exit codes, relevant runtime/config/ownership hashes, bounded timing, service replies and before/after inventories, actual resource-cleanup record and honest PASS/FAIL/BLOCKED. JR does not edit product, tasks or ICC; report/push only when packet specifically authorizes it. Coding agent alone reviews and archives. No task successor automatically promoted.

Sizing: implementation 0 / environment 2 / behavior 1 / verification 1 / recovery 1 = 5. Environment discovery and execution separated by hard safety gate.

CURRENT BLOCKER: Which exact disposable Docker daemon/host is authorized, with no collision with a running operator deployment? The earlier OpenHands sandbox-local Docker setup is historical, not proof of availability or approval for a new run. Coding agent must pin new safe target and exact packet before invoking Operation CWAL.
