> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# UMIG-EM-002-V — VERIFY: New Go baseline in disposable stack
Status: PLANNED; separate human promotion and exact live packet required. Stage/owner VERIFY / OpenHands JR. Previous: UMIG-EM-002-B. Next: UMIG-EM-003 (PLANNED).

Outcome: Confirm changed Go/MMA2/Replicator backend compatibility in a NEW uniquely owned empty disposable project; not retained earlier verify volumes or production. Preconditions: completed reviewed EM-002-T Go and EM-002-B Toolkit tests, safe daemon/Compose provenance, unique project/volume ownership, loopback UI and no host Modbus exposure. Proposed steps: inspect `deploy/verify/compose.yaml`, prove isolated daemon and free port, `docker compose --project-name <uniquely verified project> -f deploy/verify/compose.yaml config`, ONE `up -d --build`, healthz + two Unix sockets, RBE output disabled/unpublished, synthetic Memory -> Replicator source/destination and direct v1 per-block COMMS evidence, ownership/config hashes and scoped teardown `down --remove-orphans` WITHOUT `-v` only after label confirmation. On precondition failure BLOCKED; product contradiction FAIL. No product fixes, restart workaround, enabled RBE listener, customer data, old volume reuse or visual/production acceptance. Exact commands, maximum applies, safe cleanup and evidence must be written when separately promoted.

Size 0/2/0/1/1=4; split environment preflight if unbounded before promotion.
