# SIM-018 — Decide the Local UI-to-Simulator Runtime Boundary

Status: COMPLETED 2026-09-08 — accepted local integration architecture recorded.

## Primary outcome

Record one implementable local appliance mechanism by which the OS.js Modbus Simulator window can invoke the existing Go `ApplyRouter` and read `RuntimeStatus`, without restoring the deleted `simbridge` service or exposing a Simulator configuration/control HTTP API.

## Scope

- Identify which independently started local process owns the long-lived `ApplyRouter`, schedulers, and Raw Ingest clients.
- Define the browser-to-OS.js-server and OS.js-server-to-Go call path.
- Define request/response schemas for load, Save & Apply, and selected-device runtime status.
- Define authentication/authorization, local-only exposure, startup ordering, disconnect behavior, and ownership of the canonical Simulator document.
- Record the decision in repository architecture documentation and prepare exact implementation boundaries for SIM-019 through SIM-021.

## Non-scope

- Do not implement the transport, runtime host, routes, UI, or status polling.
- Do not introduce START, STOP, SPAWN, KILL, or REPLACE control for MMA2.
- Do not treat OS.js settings as proof that MMA2 applied a document.
- Do not restore `simbridge`, `/api/devices`, or `/api/devices/status`.

## Acceptance criteria

1. One local mechanism is selected, with alternatives and rejection reasons recorded.
2. The design names the process that survives while schedules run and defines restart/recovery behavior.
3. The design preserves the corrected contract: MMA2 auto-starts independently; Simulator may only request RESTART after a committed valid structural change; generated values use Raw Ingest.
4. The design provides a truthful path for apply results and runtime status without a network-exposed Simulator config API.

## Verification

Review the decision against SIM-010 through SIM-017 and trace these two sequences without an undefined hop:

- `Save & Apply -> ApplyRouter -> shared config commit -> MMA2 RESTART -> ready -> scheduler -> Raw Ingest`.
- `RuntimeStatus -> local integration -> OS.js window`.

## Completion evidence

- `docs/SIMULATOR_RUNTIME_INTEGRATION.md` selects the authenticated existing OS.js session WebSocket -> OS.js relay provider -> Unix-domain socket -> long-lived Go runtime path.
- The runtime process, canonical document owner, startup/reconnect behavior, versioned load/apply/status contract, mutation idempotency, and security/exposure rules are explicit.
- The design retains independent MMA2 boot ownership and restart-only control, Raw Ingest as the sole generated-value path, and adds no Simulator HTTP route or TCP listener.
- Both required sequences are fully traced, and rejected alternatives explain why browser-only settings and a recreated bridge cannot meet the operator-status requirement.

## Dependencies

- SIM-010 through SIM-017 (completed and archived).

## Sizing

Implementation 0, environment 1, behavioral 0, verification 1, decision/recovery 2 = 4. Tightly bounded architecture decision; it must complete before implementation tasks are promoted.
