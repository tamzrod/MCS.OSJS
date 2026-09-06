# MCS.OSJS External Network Exposure Directive

## Status

**Project directive.**

This document defines the intended externally exposed network surface of the MCS.OSJS appliance.

MCS.OSJS is planned as a single software appliance / single container containing separately bounded internal components. The container boundary does not require every internal component to publish a host-accessible TCP port.

## External exposure rule

The appliance must expose only the interfaces that are part of the external product contract:

1. **One OS.js management / operator endpoint.**
2. **One or more MMA2 Modbus TCP endpoints.**

The number of MMA2 TCP ports is configuration-dependent and may be greater than one.

Conceptually:

```text
MCS.OSJS container
│
├── OS.js
│   └── one externally exposed management/UI port
│
├── Orchestrator
│   └── internal only
│
├── Modbus Replicator
│   └── internal only
│
└── MMA2
    ├── externally exposed Modbus TCP port A
    ├── externally exposed Modbus TCP port B
    ├── externally exposed Modbus TCP port C
    └── ... as configured
```

Therefore the external port count is:

```text
1 OS.js port + N MMA2 Modbus TCP ports
```

where `N >= 1` when Modbus TCP service is enabled.

## Ownership of external ports

Only these components should own host-published ports by default:

- **OS.js** — operator / management access;
- **MMA2** — Modbus TCP service for external Modbus clients.

The **Orchestrator** and **Modbus Replicator** are internal appliance components and must not receive externally published ports merely for internal communication.

## Internal communication

This directive does **not** require a specific internal transport.

Internal component communication may use whichever mechanism is selected by the architecture, including an internal TCP listener, loopback connection, Unix-domain socket, in-process API, or another local mechanism.

The important distinction is:

> **An internal port is not automatically an externally published appliance port.**

Do not expose an internal service at the container/host boundary solely because donor code currently communicates over TCP.

## MMA2 and client filtering

MMA2 may retain its network/host adapter for externally connected Modbus clients, including behavior that depends on client network identity such as client-IP filtering.

External MMA2 listeners remain true network boundaries and may therefore require separate TCP listeners and separate policy handling.

Internal MCS.OSJS components must not be forced to masquerade as external Modbus clients solely to satisfy the external client-IP filtering model. The exact internal MMA2 access path remains an architectural decision until explicitly settled.

## Non-goals

This directive does not yet define:

- exact TCP port numbers (except the OS.js management port, fixed below in "OS.js management port"; MMA2 port numbers remain undecided);
- how many MMA2 listeners a deployment must create;
- how MMA2 memory spaces map to listeners;
- whether internal components use TCP, Unix sockets, or direct APIs;
- TLS termination or reverse-proxy behavior;
- authentication or authorization policy;
- container runtime syntax for publishing ports.

Those details must be defined in the appropriate architecture and implementation tasks.

## Enforcement rule

When implementing deployment or networking, JR must preserve the following boundary unless a later authorized architecture decision replaces it:

> **MCS.OSJS exposes one OS.js management endpoint and one or more MMA2 Modbus TCP endpoints. Orchestrator and Modbus Replicator remain internal-only and do not receive host-published ports by default.**

## OS.js management port

Assigned per OSJS-003:the shell listens on **TCP 18209** by default
(a deterministic, unoccupied, non-colliding port for this appliance:

```text
OSJS/src/server/config.js    default port (PORT env var overrides)
OSJS/Dockerfile                EXPOSE 18209
OSJS/README.md                run instructions use -p 18209:18209
```

The full OS.js desktop UI (HTTP+S WebSocket) serves on this port. This
assignment resolves the "exact TCP port numbers" non-goal for the OS.js
endpoint only;the MMA2 Modbus TCP port numbers remain architecture-task
decisions per the directive above.
