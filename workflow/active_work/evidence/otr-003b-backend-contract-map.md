# OTR-003B: Backend Contract Operation Matrix

**Status:** `COMPLETED`  
**Evidence Files:** [`electron/preload.js`](electron/preload.js), [`server.js`](../src/packages/MCSModbusToolkit/server.js), [`memory-contract.js`](../src/packages/MCSModbusToolkit/memory-contract.js), [`replicator-contract.js`](../src/packages/MCSModbusToolkit/replicator-contract.js)  
**Date:** 2026-09-22

---

## Socket Boundary Map

| Component | Type       | Protocol | Endpoint                      | Max Payload    |
|-----------|------------|----------|-------------------------------|----------------|
| MMA      | IPC Boundry   | Stdio     | N/A (renderer-only)           | 1MB            |
| Simulator | Runtime Sock| Net Stream| `{OSDATA}/run/modbus-simulator.sock` | 1MB          |
| Replicator| Runtime Sock| Net Stream| `{OSDATA}/run/modbus-replicator.sock` | 1MB        |

---

## Operation Matrix

### Memory (Simulator Component)

| Operation | Description                          | Contract Spec    | Enforced?  |
|-----------|--------------------------------------|------------------|------------|
| `load`    | Load memory state                    | UMIG-CF-001      | ✅ Runtime enforced by Toolkit socket layer |
| `apply`   | Apply memory changes                 | UMIG-CF-001      | ✅ Runtime enforced by Toolkit socket layer |
| `status`  | Query runtime status                | UMIG-CF-001      | ✅ Runtime enforced by Toolkit socket layer |
| N/A       | All other operations                 | N/A              | Rejected    |

### Replicator Component

| Operation | Description                          | Contract Spec    | Enforced?  |
|-----------|--------------------------------------|------------------|------------|
| `load`    | Load replicator state                | UMIG-CF-002      | ✅ Runtime enforced by Toolkit socket layer |
| `apply`   | Apply replicator changes             | UMIG-CF-002      | ✅ Runtime enforced by Toolkit socket layer |
| `status`  | Query runtime status                | UMIG-CF-002      | ✅ Runtime enforced by Toolkit socket layer |
| `suggest` | Request configuration suggestions    | UMIG-CF-002      | ✅ Runtime enforced by Toolkit socket layer |
| N/A       | All other operations                 | N/A              | Rejected    |

### MMA Component (Renderer Interface)

| Method | Function                       | Contract Spec      | Enforced?  |
|--------|--------------------------------|--------------------|------------|
| `init` | Initialize toolkit instance   | UMIG-CF-003        | ✅ Runtime enforced via preload IPC boundary |
| `onmessage` | Forward incoming message      | UMIG-CF-003       | ✅ Runtime enforced via preload IPC boundary |

---

## Summary Table

| Service Name | Socket Path                                      | Protocol | Operations                              | Max Payload | Status    |
|--------------|--------------------------------------------------|----------|---------------------------------------|-------------|------------|
| Simulator    | `{OSDATA}/run/modbus-simulator.sock`             | Net      | `load`, `apply`, `status`            | 1MB         | Enforced   |
| Replicator   | `{OSDATA}/run/modbus-replicator.sock`            | Net      | `load`, `apply`, `status`, `suggest` | 1MB         | Enforced   |

---

## Notes

- MMA renderer communication remains **IPC-bound only**; no direct runtime socket access.
- Simulator/Replicator runtime calls must originate from Toolkit entry points via their authenticated IPC interface.
- All requests are validated against contract version `1` and proper request/response envelopes.
- Request timeout gate at **25 seconds**.
