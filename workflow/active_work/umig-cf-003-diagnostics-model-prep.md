# UMIG-CF-003 — CODE: Prepare Dormant Diagnostics Observation Model

Status: ACTIVE — human approved the recommended isolated Diagnostics code preparation on 2026-09-19.
Stage / owner: CODE / ChatGPT
Previous: none — separately human-approved preparation, NOT advancement past UMIG-003-T/V or UMIG-006.
Next: none — STOP after source-only checkpoint; explicit next selection required.

## Outcome
Add a standalone Toolkit-owned pure read-only mapping of *per-device* Simulator and Replicator status contract results into truthful Diagnostics observations. Add future OpenHands focused tests. Preserve UNKNOWN on absent, stale/unverified, mismatched or malformed data and UNAVAILABLE for explicit failures. Do not imply any per-device observation proves global Docker service health or MMA2 supervisor health.

## Exact scope
Add only `OSJS/src/packages/MCSModbusToolkit/diagnostics-model.js` and `OSJS/tests/toolkit-diagnostics-model.test.js`. Use existing Simulator `status` result `{status:{name,device_status,mma2_status,...}}` and Replicator direct `status` result `{name,enabled,running,source_status,blocks}` as sources. Require expected device identity; derive only from supplied, fresh, trusted-integration observations. Keep global runtime/supervisor state UNKNOWN without a separate authoritative source, native paths/runtime mode UNAVAILABLE, Windows Start/Stop always disabled. Preserve explicit error text without pretending it is a successful status. No timer, network, filesystem, Electron or legacy package import.

## Gate
Read back both source files, compare bounded source diff and record exact SHA; no tests run or PASS asserted. Do not import model into renderer/entrypoint, wire adapters, enable writes, create endpoints or change existing deployment/UI. Archive source-only and return no ACTIVE. UMIG-003-T/V and UMIG-006 remain QUEUED/NOT RUN with all predecessors intact. Only BLACK SHEEP WALL changes ICC.
