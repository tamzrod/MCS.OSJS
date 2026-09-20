# UMIG-EM-002-V — VERIFY: upgraded Go/Toolkit baseline in disposable stack

Status: QUEUED — human authorized iterative migration 2026-09-20; NOT executable until predecessor PASS is reviewed and this task is the sole ACTIVE with an exact current JR packet.
Stage / owner: VERIFY / OpenHands JR; packet, safety approval review and task advancement / ChatGPT.
Previous: UMIG-EM-002-R.
Next: UMIG-EM-003 (PLANNED CODE; no automatic promotion).

## Primary outcome
Independently observe current Go/MMA2/Replicator and existing OS.js Toolkit compatibility in a newly owned, empty, isolated disposable Linux verification stack. No earlier retained volume, operator environment, live customer target or Windows installer.

## Required separate packet and acceptance

1. Preflight first: identify isolated disposable Docker daemon/context, fresh uniquely named Compose project and empty new volume, ownership labels, safe free loopback UI port, no published host Modbus port or RBE/access-event listener. Review `deploy/verify/compose.yaml` and rendered Compose config. If any proof is unavailable STOP BLOCKED; do not guess project names or reuse volumes.
2. Only after safe provenance, start ONE disposable stack and inspect `/healthz`, both Unix sockets, disabled RBE output, canonical synthetic Memory source and Replicator destination. Verify correlated per-block runtime telemetry only from direct actual responses. Compare config fingerprints/ownership before/after operations; do not assume UI render, green COMMS or restart ACK.
3. Evidence must include verbatim commands/observations, bounded apply counts, failures, and explicitly ownership-verified scoped cleanup (`down --remove-orphans` without `-v`); no other projects or volumes touched. Product contradiction FAIL, unsafe preflight BLOCKED.

Do not run this queued task or write its test packet until the preceding current-Go regression is adjudicated. No source fixes, RBE exposure, arbitrary network probing, production/customer data, old volume cleanup or final visual-parity claim. If setup and product behavior prove independently verifiable workflows, split and separately promote an environment preflight task first under `planning/microtask/rules.md`. Sizing 0/2/0/1/1=4; mandatory split before promotion if environment safety cannot be bounded.
