# UMIG-005 — CODE: Connect Toolkit Replicator Adapter

Status: COMPLETE — source-only coding checkpoint, 2026-09-19; independent TEST/VERIFY NOT RUN.
Stage / owner: CODE / ChatGPT
Previous: UMIG-004-V (archived live Memory behavior accepted with procedural deviation)
Next: UMIG-005-T (now sole ACTIVE TEST)

## Outcome and source evidence
Toolkit's own `index.js` imports `replicator-contract.js` and new `replicator-transport.js` + `replicator-editor.js`, replaces Replicator fixture DOM before attaching to OS.js, and leaves existing Memory editor and fixture Diagnostics separate. Toolkit `server.js` remains one authenticated application-WebSocket provider; namespaced `mcs-replicator-*` requests are allowlisted (`load/apply/status/suggest`) and sent only through `$OSJS_DATA_DIR/run/modbus-replicator.sock` with existing 4-byte big-endian framing, size bound, timeout, full v1 envelope and correlated unavailable errors; other request IDs preserve Simulator relay and its `load/apply/status` allowlist. No legacy package dependency or HTTP route.

`replicator-adapter.js` validates source endpoint, unit, pull block FC1–FC4, uint16 range, scan cadence, contiguous/overlapping same-FC blocks and case-insensitive unique device names. It uses Go's one apply transaction as the sole write and Go destination suggestion/explicit inspection; advisory suggestions are not called reservations, backend ownership checks and typed errors remain authoritative. Direct Go status result is mapped with per-block source/status/errors, rejects stale/unmatched responses through contract, and COMMS Network/TCP/Modbus/MMA2 LEDs remain UNKNOWN without corresponding probes. Canonical load failure never falls back to Replicator fixtures. Editor supports Add/Duplicate/Delete (including final delete), Pull Blocks edits, Save & Apply, Discard, status polling, manual destination ownership inspection, unavailable states and closure.

Source checkpoint `11bea98391b4953ca2356385bf2ffd7164bb657c` changed ONLY eight Toolkit source/test files (GitHub compare `8d5926b..11bea9` confirmed). Nine total scoped source/test files after adding `OSJS/tests/toolkit-replicator-errors.test.js` in the workflow-closure commit. Authored focused future JR tests: existing prepared contract plus adapter, transport, fake-Unix relay, typed ownership/missing-status error tests, Memory regression and bundle/build checks. Source readback of editor, index, relay, adapter and tests performed. NO tests/build/browser/live runtime verification executed by CODE. No Go, MMA2, production Docker, legacy packages, Memory implementation, Diagnostics or ICC modifications.

## Limitations / next stage
UMIG-005-T must independently run exact commands in current `handoff.md` and check Go v1 result mapping, single apply, ownership/error propagation, missing status, per-block truth, no green COMMS invention and Memory regression. UMIG-005-V remains QUEUED until TEST PASS is reviewed, then separate disposable real-browser/backend packet required. Existing isolated verify Compose from Memory lacks a Replicator service and is NOT approved as a Replicator live-test target. No production writes, live cutover or legacy retirement.

## Sizing
Surface 2, environment 0, behavior 1, verification 0, recovery 0 = 3.
