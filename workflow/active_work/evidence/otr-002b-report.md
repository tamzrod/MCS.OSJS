# OTR-002B — OS.js Toolkit Backend Baseline Evidence

**Status**: DISCOVERY COMPLETE  
**Gate Passed**: Yes (OTR-002A evidence report exists)  
**Branch Verified**: `opencode` at SHA `6bbe2a7d400e82d01c75a7c407204b2ef1e8f419`  
**Repository State**: Clean (modified handoff.md and evidence dir expected)

---

## Methods & Sockets Inventory

### MCSModbusToolkit/server.js
- **Entry Point**: WebSocket provider for authenticated OS.js connections
- **Sockets**: 
  - `modbus-simulator.sock` (Go runtime transport)
  - `modbus-replicator.sock` (Go runtime transport)
- **Runtime Timeout**: 25000ms (25s)
- **Allowed Operations** (`MEMORY_OPS`): `load`, `apply`, `status`
- **Allowed Operations** (`REPLICATOR_OPS`): `load`, `apply`, `status`, `suggest`
- **Message Handler**: Validates message size ≤ 1MB, serializes requests to JSON
- **Protocol**: Custom length-prefixed binary protocol (4-byte header + body)

### memory-transport.js
- **Transport Factory**: `createMemoryTransport(proc, {setTimer, clearTimer, timeoutMs})`
- **Pending Requests**: Per-request-id Map with Promise-based lifecycle
- **Request Validation**: Requires valid `request_id` string; rejects duplicates
- **Timeout Handling**: 26000ms default; auto-cleanup via timer callback
- **Error Types**: `Toolkit Memory transport closed`, `Invalid or duplicate Toolkit Memory request ID`, `Simulator runtime request timed out`

### memory-contract.js
- **Version**: 1 (UMIG-CF-001)
- **Client-Side Sequence**: Instance-based sequence tracking per client session
- **Validation Rules**:
  - Response must be record with matching version, request_id
  - `ok` boolean required for success/fail branching
  - Error objects may carry `code` and `message` fields
- **Error Class**: `MemoryContractError(code: string, message: string)`
- **Status**: Staged but NOT wired to any Toolkit entry point yet

### replicator-transport.js
- **Transport Factory**: `createReplicatorTransport(proc, {setTimer, clearTimer})`
- **Request ID Format Prefix**: `mcs-replicator-` (required validation)
- **Message Callback**: Emits via `proc.on('ws:message', response)` when closed or invalid request_id
- **Timeout Handling**: 26000ms default; pending requests rejected if window closes

### replicator-contract.js
- **Version**: 1 (UMIG-CF-002)
- **Client-Side Sequence**: Instance-based per client session
- **Validation Helpers**:
  - `validDocument(value)`: Ensures devices array present in record
  - `validSuggestion(value)`: Validates port (1–65535), unit_id (0–65535), owner string, status string
- **Invalid Payload Checks**:
  - `invalidRequest`: Request malformed or version mismatch
  - `invalidResponse`: Response envelope invalid
  - Missing/unexpected `version` or `request_id` rejected
- **Status**: Transport-injected, dormant (not imported by Toolkit UI or runtime)

### replicator-adapter.js
- **Document Transformers**:
  - `normalizeDocument(value)`: Converts `pull_block` → `pull_blocks[]` array for persistence
  - `blankDevice(sequence, suggestion)`: Factory producing template device with sample pull block
- **Validation Rules**:
  - Device name required and trimmed
  - Endpoint format: `[host]:port` or `ip:port`; port must be 1–65535
  - Unit ID: source (0–255), destination implied in device record
  - Destination object required
- **Document Cloning**: `copy()` via `JSON.parse(JSON.stringify())` for safe deep clone
- **Blank Block Factory**: `{function: 3, start: 0, count: 16, scan_rate_ms: 1000}`

---

### ModbusSimulator/server.js
- **Socket Path**: `{OSJS_DATA_DIR||process.cwd()}/run/modbus-simulator.sock`
- **Allowed Operations** (`ALLOWED`): `load`, `apply`, `status`
- **Message Handler**: Single Promise-based callRuntime awaiting Go backend response
- **Timeout**: 25000ms

### ModbusReplicator/server.js
- **Socket Path**: `{OSJS_DATA_DIR||process.cwd()}/run/modbus-replicator.sock`
- **Allowed Operations** (`ALLOWED`): `load`, `apply`, `status`, `suggest`
- **Message Handler**: Similar Promise-based callRuntime awaiting Go backend response
- **Timeout**: 25000ms

---

## Summary of Patterns
- **No OS.js messaging** in contract or transport layers (per file comments)
- **No config write** anywhere in these files
- **Transport lifecycle ownership** stays in the future Toolkit-owned module
- **Contract validation** separates from transport concerns
- **Go backend remains authority** for ownership, source validation, and writes

---

## File Inventory Confirmation
All 8 required backend files confirmed present:

1. `OSJS/src/packages/MCSModbusToolkit/server.js` ✅
2. `OSJS/src/packages/MCSModbusToolkit/memory-transport.js` ✅
3. `OSJS/src/packages/MCSModbusToolkit/memory-contract.js` ✅
4. `OSJS/src/packages/MCSModbusToolkit/replicator-transport.js` ✅
5. `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js` ✅
6. `OSJS/src/packages/MCSModbusToolkit/replicator-adapter.js` ✅
7. `OSJS/src/packages/ModbusSimulator/server.js` ✅
8. `OSJS/src/packages/ModbusReplicator/server.js` ✅

No additional files modified beyond evidence directory creation per baseline instructions.

---

**Verification**:  
Remote SHA matches local HEAD (6bbe2a7d400e82d01c75a7c407204b2ef1e8f419)  
Branch verified as `opencode`; repository clean aside from expected evidence artifacts.

**Evidence written to**: `workflow/active_work/evidence/otr-002b-report.md` ✅
