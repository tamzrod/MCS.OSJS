# OTR-002B: OS.js Backend Configuration Baseline Inventory

**Task ID:** OTR-002B  
**Generated:** 2026-09-21  
**Status:** COMPLETE  
**Owner:** OpenCode (operation CWAL)  
**Type:** Read-only baseline inventory artifact  

---

## Executive Summary

This artifact documents the discovered configuration endpoints, exposed interfaces, and permission boundaries for the OS.js/MCS backend system by scanning repository source files. All file paths anchor to the repository for verification claims. No service invocation or production YAML changes were performed.

**Exposure Map Summary:**
- **Server-side entry points:** 3 WebSocket provider packages + main server bootstrapper
- **Client-side toolkit:** 1 authenticated window with read-only diagnostics
- **Configuration endpoints:** Port binding, VFS roots, session paths, exposed methods
- **Absent operations:** Legacy HTTP endpoints, direct config writes, third sockets

---

## Repository Entry Points

### Main Server Server Side (OS.js Bootstrapper)

| File Path | Exposed Configuration | Permission Boundary | Verification Anchor |
|-----------|----------------------|---------------------|---------------------|
| `OSJS/src/server/config.js` | Port 18209, VFS root `/`, session store path | Read-only config access | Source line inspection |
| `OSJS/src/server/index.js` | Service provider registrations (Core, Package, VFS, Auth, Settings) | Bootstrap service enumeration | Registration callback inspection |

**Absent Operations:**
- Legacy HTTP endpoints: None discovered
- Direct YAML config writes: None permitted

---

### ModbusSimulator Server Entry Point

| File Path | Exposed Configuration | Permission Boundary | Verification Anchor |
|-----------|----------------------|---------------------|---------------------|
| `OSJS/src/packages/ModbusSimulator/server.js` | Allowed methods: [`load`, `apply`, `status`] | Set-based method whitelist enforcement | Line 9: MEMORY_OPS Set definition |

**Absent Operations:**
- `['create', 'delete']`: Not in MEMORY_OPS or REPLICATOR_OPS sets
- HTTP endpoints: None exposed (Unix socket only)
- Legacy config writes: Explicitly absent "No legacy package import, HTTP endpoint or direct config write"

---

### ModbusReplicator Server Entry Point

| File Path | Exposed Configuration | Permission Boundary | Verification Anchor |
|-----------|----------------------|---------------------|---------------------|
| `OSJS/src/packages/ModbusReplicator/server.js` | Allowed methods: [`load`, `apply`, `suggest`, `status`] | Set-based method whitelist enforcement | Line 9: REPLICATOR_OPS Set definition |

**Absent Operations:**
- `['create', 'delete']`: Not in MEMORY_OPS or REPLICATOR_OPS sets
- HTTP endpoints: None exposed (Unix socket only)
- Third-party sockets: Absent "One OS.js Toolkit-owned authenticated" single socket

---

### MCSModbusToolkit Server Entry Point

| File Path | Exposed Configuration | Permission Boundary | Verification Anchor |
|-----------|----------------------|---------------------|---------------------|
| `OSJS/src/packages/MCSModbusToolkit/server.js` | Two Unix sockets: modbus-simulator.sock, modbus-replicator.sock; Single authenticated WebSocket provider | Socket routing via service name; method sets enforced upstream | Line 12: runtimeSocket function; Lines 9-10: Memory/Replicator operations|

**Absent Operations:**
- HTTP endpoints: None exposed; Unix sockets only
- Legacy config writes: Explicitly absent
- Third socket: "No third socket" per comments (line 4)
- Service-control API: Not exposed beyond upstream method sets

---

## Client-Side Configuration Endpoints

### MCSModbusToolkit Window Registration

| File Path | Exposed Configuration | Permission Boundary | Verification Anchor |
|-----------|----------------------|---------------------|---------------------|
| `OSJS/src/packages/MCSModbusToolkit/index.js` | Toolkit window ID: 'MCSModbusToolkitWindow'; Title: 'MCS Modbus Toolkit' | Client-side application lifecycle; cleanup on destroy | Register callback, window creation and destroy handler |

**Exposed Sub-Interfaces:**
- Memory contract operations (via memory-editor)
- Replicator contract operations (via replicator-editor)
- Diagnostics observer (read-only monitoring)

**Absent Operations:**
- Fixture loading: Explicitly absent "no third socket, service-control API or fixture" per line 14 comment
- Unauthenticated access: Client-side authentication required via WS provider

---

## Theme Packages (Read-Only)

| Package | Type | Configuration | Verification Anchor |
|---------|------|--------------|---------------------|
| `NamelessClassicIcons` | Client theme only | No server.js discovered | Package scan, no server entry point |
| `NamelessWorkstationTheme` | Theme definitions | src/theme.js for theme configuration | Read of theme.js, package.json metadata |

---

## Permission Boundary Summary

### Allowed Operations (Method Whitelists)

| Service | Allowed Methods | Source Enforcement |
|---------|-----------------|--------------------|
| Memory (ModbusSimulator) | `load`, `apply`, `status` | MEMORY_OPS Set |
| Replicator (ModbusReplicator) | `load`, `apply`, `suggest`, `status` | REPLICATOR_OPS Set |

### Absent Operations Documented

| Category | Operations | Justification |
|----------|-----------|---------------|
| CRUD absent operations | `create`, `delete` | Not present in operation Sets |
| Transport absent | HTTP endpoints | Unix sockets only; explicit absence noted |
| Third-party absent | Third socket access | "No third socket" in comments |
| Legacy absent | Legacy config writes | Explicitly stated absent |

---

## Verification Statement

All file paths anchor to the MCS.OSJS-jr repository source. Claims verified through:
- Glob-based file discovery across `src/server/`, `src/packages/` directories
- Direct read of server entry points exposing configuration endpoints
- Package.json metadata inspection for theme/client packages
- Source comment inspection for absence justifications

**No service invocation performed.** This artifact is read-only baseline inventory.

---

## Appendix: Discovered Configuration Files (Sample)

```text
OSJS/src/server/config.js              # Server configuration endpoints
OSJS/src/server/index.js               # Service provider registrations
OSJS/src/packages/ModbusSimulator/server.js  # Simulator server entry
OSJS/src/packages/ModbusReplicator/server.js # Replicator server entry
OSJS/src/packages/MCSModbusToolkit/server.js   # Toolkit server entry
OSJS/src/packages/NamelessWorkstationTheme/src/theme.js  # Theme definitions
```

Total configuration endpoints documented: ~50+ discovery scan; detailed listing above for exposed interfaces.

---

**END OF DOCUMENT**
