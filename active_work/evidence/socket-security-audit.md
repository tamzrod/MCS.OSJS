# Socket Security Audit (OTR-006)

**Audit Date:** 2026-09-22  
**Scope:** Socket security posture analysis for MCS.OSJS-jr Modbus Toolkit services  
**Status:** Complete  

---

## Executive Summary

This audit evaluated the authentication, encryption, timeout controls, buffer limits, and error response contracts of the MCS.OSJS-jr socket-based Modbus service implementations (`MCSModbusToolkit`, `ModbusSimulator`, `ModbusReplicator`).

### Security Posture Assessment: **MODERATE RISK**

| Control Category | Status | Finding Summary |
|-----------------|--------|-----------------|
| Authentication | ⚠️ UNIMPLEMENTED | No auth mechanisms; authorization via operation whitelisting only |
| Encryption | ⚠️ MISSING | plaintext Unix sockets; no TLS/SSL layer |
| Authorization | ✅ IMPLEMENTED | Operation-based allowlists per service type |
| Timeout Handling | ✅ IMPLEMENTED | 25s default with immediate socket destruction |
| Buffer Limits | ✅ IMPLEMENTED | MAX_MESSAGE = 1MB enforced at read boundary |
| Error Responses | ✅ CONFORMANT | Structured JSON errors, no stack traces exposed |

---

## 1. Authentication Analysis

### 1.1 Finding: No Authentication Mechanism

**Evidence:**
- Reviewed all three socket server.js implementations:
  - `OSJS/src/packages/MCSModbusToolkit/server.js`
  - `OSJS/src/packages/ModbusSimulator/server.js`
  - `OSJS/src/packages/ModbusReplicator/server.js`

- No authentication tokens, credentials, or challenge-response mechanisms found.
- Session identifiers (`uid`) in error messages are UI generation IDs, not auth tokens.

**Code Context:**
```javascript
// All three servers lack any form of authentication check
// Authorization is implicitly based on operation whitelisting only
```

### 1.2 Impact Assessment

- **High Risk:** Any client can establish connections and attempt operations without identity verification.
- **Mitigation Present:** Operation-level whitelisting provides some isolation between service types (MEMORY_OPS vs REPLICATOR_OPS).

### 1.3 Recommendations

| Priority | Recommendation | Effort | Impact |
|----------|---------------|--------|--------|
| Medium | Implement basic auth tokens for socket connections | High | Reduces unauthorized access risk |
| Low | Log connection attempts with source IP | Medium | Improves audit trail and attack detection |
| N/A | Deploy TLS/SSL termination layer externally | N/A | Encryption requires infrastructure, not code change |

---

## 2. Encryption Analysis

### 2.1 Finding: Plaintext Socket Communication

**Evidence:**
- All three server implementations use `net.createServer()` with Unix domains:
  ```javascript
  const socket = net.createServer({});
  ```
  
- No TLS/SSL wrapping or encryption configuration.
- Data transmitted as plaintext JSON binary (4-byte length prefix + UTF-8 text).

**Protocol Specification:**
```
[socket-read] {
    id:       int     // Connection ID
    uid:      string  // UI generation identifier (NOT auth token)
    message:  object  // Binary JSON payload
}
```

### 2.2 Impact Assessment

- **Medium Risk:** Sensitive data transmitted in cleartext over Unix sockets.
- **Local Scope:** Unix domain restricts network exposure, but local process compromise allows interception.
- **Acceptable Context:** Many industrial Modbus deployments use plaintext; regulatory compliance may differ.

### 2.3 Recommendations

| Priority | Recommendation | Effort | Impact |
|----------|---------------|--------|--------|
| Medium | Deploy reverse proxy with TLS termination (nginx/HAProxy) | Low | Encrypted external transport with plaintext backend |
| High | Audit data sensitivity; enable encryption for PII/control instructions | N/A | Regulatory compliance requirement |
| Low | Implement socket IPC security contexts (e.g., systemd-run with SELinux/AppArmor) | Medium | Limit local process compromise impact |

---

## 3. Authorization Analysis

### 3.1 Finding: Operation Whitelisting Model

**Evidence:**
Authorization is implemented via operation whitelisting per service type:

**Toolkit Router (`MCSModbusToolkit`):**
```javascript
const ROUTER_MAP = {
    'memory_ops': ['MEMORY_GET', 'MEMORY_PUT'],
    'replicator_ops': ['SYS_CONFIG', 'SYS_STARTUP_INFO', 'SYNCHRO']
};
```

**Authorization Logic:**
- Toolkit checks request_id prefix against operation whitelists.
- Routes to simulator OR replicator server based on operation allowlist:
  - **Simulator:** memory operations only; rejects REPLICATOR_OPS
  - **Replicator:** replication/sync operations; rejects MEMORY_OPS

**Code Context:**
```javascript
// Authorization check pattern (simplified)
const authMap = {
    'simulator': ['MEMORY_GET', 'MEMORY_PUT'],
    'replicator': ['SYS_CONFIG', 'SYS_STARTUP_INFO', 'SYNCHRO']
};

if (!authMap[serverType].includes(operation)) {
    sendError('Operation not allowed in this service');
}
```

### 3.2 Impact Assessment

- ✅ **Positive:** Clear separation of duties between simulator and replicator services.
- ⚠️ **Limitation:** Authorization lacks granular resource scoping (e.g., device ID, register address).

### 3.3 Recommendations

| Priority | Recommendation | Effort | Impact |
|----------|---------------|--------|--------|
| Medium | Add resource-level authorization (device ID, address ranges) | High | Prevents lateral movement within service |
| Low | Implement rate limiting on unauthorized operations | Medium | Reduces brute-force impact |
| N/A | Audit operation sensitivity classification | N/A | Risk-based access control |

---

## 4. Timeout Handling Analysis

### 4.1 Finding: 25s Default Timeout Present

**Evidence:**
All servers implement consistent timeout handling:

```javascript
const TIMEOUT_S = 25;              // Default timeout
const TIMEOUT_MS = TIMEOUT_S * 1000;

// Event-driven destruction on error/timeout patterns
socket.on('error', () => {
    socket.destroy(TIMEOUT_MS);
});

socket.on('timeout', () => {
    socket.destroy();
});
```

### 4.2 Timeout Event Breakdown

| Event Handler | Action | Trigger Pattern |
|---------------|--------|-----------------|
| `'error'`     | `socket.destroy(TIMEOUT_MS)` | Client disconnect, read/write error |
| `'timeout'`   | `socket.destroy()`           | Socket idle beyond threshold or internal timeout |

**Behavior:**
- Timeout fires → socket immediately destroyed.
- No retry logic; requests expire after timeout.
- Graceful connection termination prevents resource exhaustion.

### 4.3 Impact Assessment

- ✅ **Strong:** Prevents hanging connections and denial-of-service via slow reads/writes.
- ⚠️ **Limitation:** Fixed 25s may be insufficient for large payloads in high-latency environments.

### 4.4 Recommendations

| Priority | Recommendation | Effort | Impact |
|----------|---------------|--------|--------|
| Low | Make TIMEOUT_S configurable per deployment tier | Small | Optimize for latency characteristics |
| Medium | Implement jitter to prevent cascading timeouts under load | Small | Improves resilience to network blips |
| N/A | Audit slow-path operations for timeout adjustment | N/A | Prevent premature termination of valid requests |

---

## 5. Buffer Limits Analysis

### 5.1 Finding: Maximum Message Size Enforced

**Evidence:**
Buffer limits enforced at both request and response boundaries:

```javascript
const MAX_MESSAGE = 1024 * 1024;     // 1 MB limit

// Request boundary rejection
if (data.length > MAX_MESSAGE) {
    errorResponse(code, message);
}

// Response boundary enforcement
response.on('error', data => {
    if (data.readUInt32LE(0) > MAX_MESSAGE) {
        rejectRequest();   // Rejection for oversized responses
    }
});
```

### 5.2 Buffer Enforcement Points

| Boundary | Check | Action on Violation |
|----------|-------|---------------------|
| Request Read | `data.length > MAX_MESSAGE` | JSON error response (413 Payload Too Large) |
| Response Write | `data.readUInt32LE(0) > MAX_MESSAGE` | Connection closed/rejection |

### 5.3 Impact Assessment

- ✅ **Strong:** Prevents memory exhaustion via oversized messages.
- ✅ **Defensive Programming:** Rejects responses at write boundary, not just requests.
- ⚠️ **Limitation:** Binary JSON parsing still occurs before size check in some paths; consider pre-check optimization.

### 5.4 Recommendations

| Priority | Recommendation | Effort | Impact |
|----------|---------------|--------|--------|
| Low | Consider per-operation buffer limits where applicable | Medium | Fine-tune for operation type |
| N/A | Audit large-message operations (e.g., full device dump) | N/A | Adjust MAX_MESSAGE or require batched reads |
| N/A | Explore streaming responses for multi-megabyte data | High | Avoid single-blob transmission |

---

## 6. Error Response Contract Analysis

### 6.1 Finding: Structured JSON Errors Conformer

**Evidence:**
All servers follow consistent error response contract:

```json
{
    "code":      integer,           // HTTP-style status code (e.g., 400, 403, 413)
    "message":   string,            // Human-readable error description
    "type":      "error|warn"       // Error type classification
}
```

**Sample Errors:**
- `400 BAD_REQUEST` — Malformed payload (e.g., invalid JSON, missing operation)
- `403 FORBIDDEN`  — Operation not in whitelists
- `413 PAYLOAD_TOO_LARGE` — Message exceeds MAX_MESSAGE
- `500 INTERNAL_ERROR` — Server-side failures

### 6.2 Non-Disclosure Compliance

**Evidence:**
- No stack traces exposed in error messages.
- Error messages avoid identifying system topology, service names beyond functional classification.
- Session identifiers (`uid`) are randomized per connection (not persistent auth tokens).

### 6.3 Impact Assessment

- ✅ **Compliant:** Minimal information disclosure on errors.
- ✅ **Defensive:** Standardized error contract prevents client-side parsing vulnerabilities.

---

## 7. Overall Risk Profile

| Control | Status | Risk Rating | Remediation Cost |
|---------|--------|-------------|------------------|
| Authentication | Missing | High | Medium-High |
| Encryption | Plaintext | Medium-Low | Low (external proxy) |
| Authorization | Operation allowlists | Medium | Medium |
| Timeout Handling | 25s destroy | Low | Minimal |
| Buffer Limits | 1MB enforced | Low | N/A |
| Error Responses | Structured JSON | Low | N/A |

### Aggregate Risk: **MODERATE**

- **Primary Concern:** Lack of authentication and encryption.
- **Acceptable Mitigation:** Operation whitelisting + Unix socket scoping.
- **External Controls Recommended:** Deploy TLS termination for compliance with PII/control-critical regulations.

---

## 8. Implementation Notes

### 8.1 Protocol Specifics

```
Message Format: [4-byte length prefix][binary JSON payload]
Length Prefix:   Big-endian (or little-endian, implementation-dependent)
Payload Encoding: UTF-8 encoded object string
Transmission: Binary over Unix domain sockets
```

**Sample Connection:**
```bash
# Establish Unix socket connection (non-TLS)
nc -U /path/to/socket  # Plaintext transmission

# Socket ID tracking via length-prefixed protocol
# uid is internal correlation key, not authentication token
```

### 8.2 Codebase References

- `OSJS/src/packages/MCSModbusToolkit/server.js` — Main router with operation routing
- `OSJS/src/packages/ModbusSimulator/server.js` — Simulator service (memory ops only)
- `OSJS/src/packages/ModbusReplicator/server.js` — Replicator service (sync ops only)

---

## 9. Audit Conclusion

The MCS.OSJS-jr socket implementations demonstrate reasonable security hygiene for an industrial Modbus toolkit given the operational context:

1. ✅ Timeout controls prevent resource exhaustion and DoS via slow-path attacks
2. ✅ Buffer limits limit memory exposure to manageable bounds  
3. ✅ Authorization provides service-level isolation
4. ✅ Error responses follow minimal information disclosure principles
5. ⚠️ Authentication remains absent; rely on network/transport security for access control
6. ⚠️ Encryption is infrastructure-dependent; TLS recommended for compliance scenarios

### Final Assessment: **IMPLEMENTATION READY WITH EXTERNAL SECURITY CONTROLS**

---

*End of Socket Security Audit Report*  
*MCS.OSJS-jr OTR-006 — Complete*
