# MCS.OSJS External Network Exposure Directive

## Status

**Project directive.**

This document defines the intended externally exposed network surface of the MCS.OSJS appliance.

MCS.OSJS is planned as a single software appliance / single container containing separately bounded internal components.

## Network mode requirement

The MCS.OSJS appliance **must run with Docker host networking**:

```yaml
network_mode: "host"
```

The purpose is to keep external Modbus TCP traffic out of Docker bridge/NAT translation. MMA2 listeners must bind directly to the appliance host network stack rather than being exposed through Docker `ports:` mappings.

This follows the deployment method already proven by the `tamzrod/replicator-stack` donor, where the MMA and Replicator services use `network_mode: "host"`.

For the appliance container, do **not** use Docker port publishing such as:

```yaml
ports:
  - "502:502"
  - "503:503"
```

for MMA2 Modbus TCP endpoints.

Conceptually:

```text
Linux appliance host network stack
│
├── OS.js management listener
├── MMA2 Modbus TCP listener A
├── MMA2 Modbus TCP listener B
├── MMA2 Modbus TCP listener C
└── ... as configured
        ↑
        direct host-network listeners
        no Docker bridge/NAT translation
```

Because host networking removes Docker's bridge as an exposure boundary, internal-only components must not bind externally reachable addresses merely for internal communication.

## External exposure rule

The appliance must expose only the interfaces that are part of the external product contract:

1. **One OS.js management / operator endpoint.**
2. **One or more MMA2 Modbus TCP endpoints.**

The number of MMA2 TCP ports is configuration-dependent and may be greater than one.

Conceptually:

```text
MCS.OSJS container (network_mode: host)
│
├── OS.js
│   └── one externally exposed management/UI listener
│
├── Orchestrator
│   └── internal only
│
├── Modbus Replicator
│   └── internal only
│
└── MMA2
    ├── externally exposed Modbus TCP listener A
    ├── externally exposed Modbus TCP listener B
    ├── externally exposed Modbus TCP listener C
    └── ... as configured
```

Therefore the external port count is:

```text
1 OS.js port + N MMA2 Modbus TCP ports
```

where `N >= 1` when Modbus TCP service is enabled.

## Ownership of external ports

Only these components should own externally reachable listeners by default:

- **OS.js** — operator / management access;
- **MMA2** — Modbus TCP service for external Modbus clients.

The **Orchestrator** and **Modbus Replicator** are internal appliance components and must not create externally reachable listeners merely for internal communication.

## Internal communication

This directive does **not** require a specific internal transport.

Internal component communication may use whichever mechanism is selected by the architecture, including loopback TCP, Unix-domain socket, in-process API, or another local mechanism.

With `network_mode: "host"`, any TCP service intended to remain internal must be deliberately constrained, for example by binding to loopback rather than `0.0.0.0` when TCP is used.

The important distinction is:

> **Host networking is required for the appliance, but host-network access does not authorize every component to expose a network service.**

## MMA2 and client filtering

MMA2 may retain its network/host adapter for externally connected Modbus clients, including behavior that depends on client network identity such as client-IP filtering.

External MMA2 listeners remain true network boundaries and may therefore require separate TCP listeners and separate policy handling.

Host networking is required specifically so these external Modbus connections are not translated through Docker bridge/NAT before reaching MMA2.

Internal MCS.OSJS components must not be forced to masquerade as external Modbus clients solely to satisfy the external client-IP filtering model. The exact internal MMA2 access path remains an architectural decision until explicitly settled.

## Non-goals

This directive does not yet define:

- exact MMA2 TCP port numbers;
- how many MMA2 listeners a deployment must create;
- how MMA2 memory spaces map to listeners;
- whether internal components use loopback TCP, Unix sockets, or direct APIs;
- TLS termination or reverse-proxy behavior;
- authentication or authorization policy.

Those details must be defined in the appropriate architecture and implementation tasks.

## Enforcement rule

When implementing deployment or networking, JR must preserve the following boundary unless a later authorized architecture decision replaces it:

> **MCS.OSJS runs with Docker `network_mode: "host"` to avoid Docker bridge/NAT translation. OS.js and MMA2 are the only components that may own externally reachable listeners by default. Orchestrator and Modbus Replicator remain internal-only.**

MMA2 Modbus TCP endpoints must not be implemented using Docker `ports:` publishing as a substitute for the required host-network mode.

## OS.js management port

Assigned per OSJS-003: the shell listens on **TCP 18209** by default (a deterministic, unoccupied, non-colliding port for this appliance):

```text
OSJS/src/server/config.js    default port (PORT env var overrides)
OSJS/Dockerfile              EXPOSE 18209
OSJS/README.md               current standalone run instructions
```

The full OS.js desktop UI (HTTP + WebSocket) serves on this port. In the final appliance deployment, this listener shares the host network namespace under the required `network_mode: "host"`; Docker `ports:` publishing is not required for the appliance container.

This assignment resolves the exact TCP port for the OS.js endpoint only; MMA2 Modbus TCP port numbers remain architecture-task decisions per the directive above.
