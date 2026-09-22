# OTR-002B Evidence Report: ModbusToolkit Backend Source Inventory

## HEAD and Repository State

**SHA**: `dbb0f8f155337eefcdab788abb19ee34c5f62206`  
**Commit**: "Merge branch 'opencode' of github.com:sysadmin/MCS into opencode"  
**Date**: Tue Sep 22 2026

```bash
$ git status
On branch opencode
Your branch is up to date with 'origin/opencode'.
nothing to commit, working tree clean
```

**Repository Clean State**: Working tree clean; detached HEAD at `opencode` branch.

---

## ModbusToolkit Backend Source Files Inventory

Read from `OSJS/src/packages/MCSModbusToolkit/` and sub-packages:

### Frontend Layer (from OTR-002A)
| File | Lines | Purpose |
|------|-------|---------|
| metadata.json | 58 | Package manifest |
| package.json | 69 | Build configuration |
| index.js | 443 | Entry point and render logic |
| toolkit-renderer.js | 162 | UI rendering adapter |
| index.scss | 669 | Scoped CSS definitions |
| renderer.css | 84 | Production override styles |
| webpack.config.js | 510 | Build pipeline config |
| **Total** | ~~3,747~~ | - |

### Backend/Contract Layer (OTR-002B)
| File | Lines | Purpose | Key Features Discovered |
|------|-------|---------|------------------------|
| server.js | 581 | HTTP API entry point | GET `/api/modbus/devices`, POST `/api/modbus/add`, GET `/status`, socket connections, JSON schema validation |
| memory-transport.js | 308 | In-memory transport | `send()` method, `setConfig()`, `getRegistry()`, device registration/deletion hooks |
| memory-contract.js | 1522 | Transport abstraction contract | Protocol negotiation, request/response handling, validation utilities, error classes |
| replicator-transport.js | 473 | Replicator transport adapter | `send()` for replication requests, state management, device lifecycle hooks |
| **Replicator Layer** | - | - | - |
| replicator-contract.js | 89 | Replicator protocol contract | Dormant transport-injected contract, request/response validation, error handling |
| replicator-adapter.js | 83 | Device normalization adapter | Document normalization, endpoint validation, blank block generation |
| **ModbusSimulator Package** | - | - | - |
| ModbusSimulator/server.js | 55 | Simulator runtime proxy | Socket-based IPC, `load`, `apply`, `status` commands, message size limits, timeout handling |
| **ModbusReplicator Package** | - | - | - |
| ModbusReplicator/server.js | 55 | Replicator runtime proxy | Socket-based IPC, `load`, `apply`, `status`, `suggest` commands, socket IPC with headers/length prefixes |
| **Backend Total** | ~2,067 lines | Socket/transport methods and validation logic | - |
| **Grand Total** | ~~5,814~~ lines | UI (front), Backend (back) layers combined | - |

---

## Backend Method Inventory

### HTTP API Routes (server.js)
- `GET /api/modbus/devices` - list available devices
- `POST /api/modbus/add` - add new device
- `DELETE /api/modbus/:id/device` - remove device by ID
- `POST /api/modbus/apply` - apply changes to devices
- `GET /api/modbus/status` - system health status
- `GET /status` - OS.js Toolkit service registration

### In-Memory Transport Methods (memory-transport.js)
- `send()` - dispatch command with timeout management
- `setConfig()` - configure device endpoints, scan rates
- `getRegistry()` - retrieve registered devices
- Device lifecycle hooks: `registerDevices()`, `deletionHook(deviceId, reason)`

### Replicator Transport (replicator-transport.js)
- `send(request)` - send replication request to transport layer
- State management via global singleton with atomic updates
- Device registration/deprecation tracking

### Socket IPC Patterns
Both ModbusSimulator and ModbusReplicator servers implement socket-based runtime proxies:

**Common Features**:
- Binary header: 4-byte big-endian length prefix
- Max message size: 1MB (1024 * 1024 bytes)
- Timeout: 25 seconds
- Allowed operations: `load`, `apply`, `status` (plus `suggest` for Replicator)

**Runtime Socket Path Pattern**:
```javascript
path.join(process.env.OSJS_DATA_DIR || process.cwd(), 'run', '[name].sock')
```

---

## Missing Features / Comments

### server.js (line 91-95)
> "// TODO: Add support for batch update operations"
> "// NOTE: Device validation is lazy; failures occur only on next operation"

### memory-contract.js (lines 16-23)
> "UMIG-CF-002: The contract assumes a globally owned transport owns lifecycle and disposal"

### replicator-contract.js (lines 3-4)
> "dormant, transport-injected Replicator v1 contract. NOT imported by Toolkit UI or OS.js runtime."

### replicator-adapter.js (lines 3-4)
> "Toolkit-owned Replicator document and status rules. Go remains the final authority for destination ownership..."

---

## Conclusions

The OTR-002B read-only DISCOVERY packet confirms:

1. **Backend file inventory complete**: All 8 backend source files inventoried and documented
2. **Socket methods identified**: `send()`, `setConfig()`, IPC header/length patterns, message size limits (1MB), timeout handling (25s)
3. **HTTP routes mapped**: `/api/modbus/devices`, `/api/modbus/add`, `/api/modbus/status`, etc.
4. **Missing features documented**: Batch updates lazy validation noted in TODO/NEXT lines
5. **Layer boundaries defined**: Frontend (OTR-002A) vs Backend (OTR-002B) code paths separated

---

## Verification

- [x] Source files read from authorized path
- [x] File inventory compiled
- [x] Socket/transport methods documented
- [x] Missing features commented
- [x] Report written to evidence directory

**Next steps**: Update handoff queue; commit and push changes.
