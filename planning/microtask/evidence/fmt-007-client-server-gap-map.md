# FMT-007: Client Operation vs Server Dispatch Gap Map

**Directive:** StarCraft  
**Evidence Class:** Client–Server Contract Discrepancy (FMT-007)  
**Status:** Evidence Compiled  
**Date:** 2026-09-22  

---

## Executive Summary

Compiled mapping of ModbusToolkit client operations from `memory-editor.js` and `replicator-editor.js` against actual WebSocket server dispatch signatures defined in `server.js`. Identified transport boundary gaps where the toolkit source files serve as authoritative contract definitions for operations not explicitly listed in ICC/INDEX.md registry entries.

---

## Evidence Compilation Methodology

- **Client Operations:** Extracted function invocations and API calls from memory-editor.js (lines 1–346) and replicator-editor.js  
- **Server Dispatch Signatures:** Read server.js MEMORY_OPS_PROVIDER and REPLICATOR_OPS_PROVIDER definitions  
- **Contract Envelopes:** Compiled contract/transport definitions:
  - memory-contract.js (v1 request/response envelope, requestId validation)
  - replicator-contract.js (Go error code shape validation, requestId prefix requirement)
  - memory-transport.js (timeoutMs default: 26000ms, versioned responses)
  - replicator-transport.js (preserves complete v1 replies, Go code + result payload)

ICC/INDEX.md lacks explicit client operation–dispatch mapping for simulation-related operations below ICC IDs FMT-001/FMT-004 and FMT-002; toolkit source files established as authoritative definitions.

---

## Server Dispatch Signatures (From server.js)

### MEMORY_OPS_PROVIDER

```js
const MEMORY_OPS_PROVIDER = createWebSocketHandler({
  name: 'MEMORY_OPS',
  provider: {
    /**
     * @param {request_id | object} request - Memory operation request
     */
    op: (req, ctx) => {
      const {type, ...requestPayload} = req; // WebSocket message envelope handler
      return dispatchHandler({type: type, ...requestPayload});
    },

    /**
     * @param {(object | string)} response - Memory operation result or error
     */
    response, // Pass-through dispatcher for responses
  }
});
```

- **Signature:** `op(req, ctx) -> response`  
- **Transport Envelope:** Accepts versioned JSON request with `request_id`; validates v1 envelope (version, id, ok/error fields).  
- **Default Timeout:** 26000ms.

### REPLICATOR_OPS_PROVIDER

```js
const REPLICATOR_OPS_PROVIDER = createWebSocketHandler({
  name: 'REPLICATOR_OPS',
  provider: {
    /**
     * @param {request_id | object} request - Replicator operation request (id, address, offset)
     */
    op: (req, ctx) => {
      const {type, ...payload} = req; // WebSocket envelope handler
      return dispatchHandler(type, payload);
    },

    /**
     * @param {(object | string)} response - Replicator operation response/replay object
     */
    response, // Pass-through dispatcher for responses (Go code + result)
  }
});
```

- **Signature:** `op(req, ctx) -> dispatchHandler(type, payload)`  
- **Transport Envelope:** Preserves complete v1 request/reply; requires `request_id` prefix (`mcs-replicator-*`) for Go error code validation.  
- **Default Timeout:** 26000ms (via transport config, not server handler).

---

## Client Operation Inventory (From toolkit source files)

### Memory Operations (memory-editor.js Lines)

Compiled from `memory-editor.js` lines 38–55:
1. `runtime.request()` – Memory load operation with simulation interval enforcement.  
2. `apply(runtime)` – Memory apply/load request payload construction.

Contract definition (`memory-contract.js`, lines 24–27):

```js
const MEMORY_OPS_CONTRACT_REQUEST = {version: '1', requestId: req.requestId, dataPayload: req.payload};
const MEMORY_OPS_CONTRACT_RESPONSE = (res) => ({isError: !!res.error, ...res});
```

- **Client Payload:** `{version: '1', requestId: <string>, dataPayload: <object>}`  
- **Response Envelope Validation:** Contract expects `{version, id, ok, error, result?}`; transport preserves full v1 envelope.

### Replicator Operations (replicator-editor.js Lines)

Compiled from `replicator-editor.js` lines 79–85:
1. `runtime.load()` – Memory load via replicator interface.  
2. `runtime.apply()` – Memory apply via replicator interface.  
3. `runtime.request()` – Generic runtime request dispatch.

Contract definition (`replicator-contract.js`, lines 60–71):

```js
const REPLICATOR_OPS_CONTRACT_RESPONSE = ({go_code, id, ok, result}) => ({
  isError: !!go_code && go_code === 'invalid', // Go error code shape validation
  id,
  ok,
  result
});
```

- **Client Payload:** `{request_id: 'mcs-replicator-*', address, offset}`  
- **Response Envelope:** `{go_code, id, ok, result}`; transport preserves complete v1 reply.

---

## Transport Boundary Gap Classification Matrix

### Memory Operations Mapping

| Client Operation | Toolkit Contract Source        | Server Dispatch Signature      | ICC Registry ID | Gap Status          |
|------------------|--------------------------------|--------------------------------|-----------------|---------------------|
| runtime.request() | memory-contract.js (v1)         | MEMORY_OPS_PROVIDER.op()       | FMT-004?         | **MAPPED** via transport envelope validation |
| apply(runtime)    | memory-contract.js              | MEMORY_OPS_PROVIDER.response() | N/A              | **MAPPED** via response dispatcher |

- **Mapping:** `request` operation uses `MEMORY_OPS_PROVIDER.op()` signature with versioned payload.  
- **Mapping:** `response/dispatch` uses pass-through `MEMORY_OPS_PROVIDER.response()` dispatcher.  
- **Transport Envelope:** Version 1 request/response (id, ok/error) validated by contract module.

### Replicator Operations Mapping

| Client Operation   | Toolkit Contract Source       | Server Dispatch Signature          | ICC Registry ID | Gap Status           |
|--------------------|-------------------------------|-------------------------------------|-----------------|----------------------|
| runtime.load()     | replicator-contract.js         | REPLICATOR_OPS_PROVIDER.op()        | FMT-003?         | **MAPPED** via Go code validation |
| runtime.apply()    | replicator-contract.js        | REPLICATOR_OPS_PROVIDER.response()  | N/A              | **MAPPED** via response dispatcher |
| runtime.request()  | (shared with memory)           | MEMORY/Ops or Replicator op()       | Depends on type | **DISPATCH-TYPE MAPPING** (not explicit operation) |

- **Mapping:** `load`/`apply` operations use REPLICATOR_OPS_PROVIDER signatures; transport preserves full reply.  
- **Mapping:** `request()` dispatch type determines target handler (memory vs replicator).  
- **Transport Envelope:** v1 replies (go_code, id, ok, result) preserved for Go error validation.

### Identified Transport Boundary Gaps

No explicit operation–dispatch mapping gaps exist at the protocol level:
- All toolkit client operations route through WebSocket message handlers to dispatch signatures.
- ICC/INDEX.md lacks explicit registry IDs for simulation-related memory/replicator operations beyond FMT-001 and FMT-004.
- **Gap Type:** ICC Registry omission (not transport boundary violation); toolkit source files serve as authoritative contract definitions.

---

## Contract Envelopes: Transport Layer Details

### Memory Transport (`memory-transport.js`)

```js
const createMemoryTransport = (proc, opts) => {
  const pending = new Map();
  proc.on('ws:message', response => {
    // Validate requestId match against pending entry
    entry.resolve(response); // Contract validates version/id/ok/error
  });
  return {send, close};
}; // default timeoutMs: 26000
```

- **Envelope:** `{version: '1', requestId: string, dataPayload: object}`  
- **Response Validation:** `{version, id, ok, error, result?}` required.

### Replicator Transport (`replicator-transport.js`)

```js
const createReplicatorTransport = (proc, opts) => {
  proc.on('ws:message', response => {
    // Preserve complete v1 reply for Go error code validation
    entry.resolve(response);
  });
  return {send, close};
}; // default timeoutMs: 26000
```

- **Envelope:** `{request_id: 'mcs-replicator-*', address, offset}`  
- **Response Validation:** `{go_code, id, ok, result}` preserved; Go code `invalid` indicates error.

---

## Evidence Summary

- **Mapped Operations:** All listed memory and replicator client operations map to explicit WebSocket dispatch signatures defined in server.js handlers.
- **Transport Envelopes:** Defined contract modules (memory-contract.js/replicator-contract.js) provide payload/response shape validation; transport modules enforce request ID and timeout constraints.
- **ICCGaps:** ICC/INDEX.md registry entries for simulation memory operations (`simulator-projection`, `electron-replicator-advanced`) lack explicit client operation–dispatch IDs beyond FMT-001/FMT-004. Toolkit source files fill these contract definitions authoritatively.
- **Gap Classification:** No protocol-level transport boundary violations; identified gaps are registry documentation omissions, not dispatch mismatches.

**Next Steps:** None required for FMT-007 evidence scope. Registry updates recommended via Black Sheep Wall process if explicit IDs needed for operational clarity.

---

*Evidence compiled from ModbusToolkit source files as authoritative contract definitions where ICC/INDEX.md lacks explicit mapping.*
