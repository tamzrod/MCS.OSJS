# External Network Exposure Directive

Baseline commit: 500376cfb5c222298aadfcf035aad0af0a635773
Working tree: clean
Source dependencies: docs/NETWORK_EXPOSURE.md, deploy/docker-compose.yml, OSJS/Dockerfile, OSJS/src/server/config.js
Parent: L0-project
Zoom In:(none; leaf node(
Zoom Out: L0-project

## Rule (project directive

Appliance must run with Docker host networking:

```yaml
network_mode:  "host"
```

Purpose:  keep external Modbus TCP traffic out of Docker bridge/NAT translation; MMA2 listeners bind directly to host network stack, NOT exposed via Docker `ports:` mappings. This follows the deployment method proven by `tamzrod/replicator-stack` donor (MMA/Replicator services use host networking(. Do NOT use Docker port publishing (e.g. `ports: "502:502"`( for MMA2 Modbus TCP endpoints..

## External Exposure Surface

Appliance exposes only interfaces that are part of external product contract:

1. One OS.js management/operator endpoint.
2. One or more MMA2 Modbus TCP endpoints (configuration-dependent, N >= 1 when Modbus TCP enabled..

External port count = 1 OS.js port + N MMA2 Modbus TCP ports. Only OS.js and MMA2 may own externally reachable listeners by default. Orchestrator and Modbus Replicator are internal-only and must not create externally reachable listeners merely for internal communication. Internal port != externally published appliance port; do not expose internal service at container/host boundary solely because donor code communicates over TCP. With host networking, any TCP service meant to remain internal must be deliberately constrained (e.g. bind loopback rather than 0.0.0.0(..

## MMA2 and Client Filtering

MMA2 may retain host network/adapter for externally connected Modbus clients, incl. client-IP filtering behavior. External MMA2 listeners remain true network boundaries requiring separate listeners/policy. Host networking is required specifically so external Modbus connections are not translated through bridge/NAT. Internal MCS.OSJS components must not be forced to masquerade as external Modbus clients to satisfy the external client-IP filtering model. Exact internal MMA2 access path remains architectural decision until explicitly settled..

## Non-Goals

Not defined by directive:  exact MMA2 TCP port numbers; how many MMA2 listeners a deployment must create; how MMA2 memory spaces map to listeners; internal transport choice (loopback TCP, Unix sockets, direct APIs(; TLS/reverse-proxy; auth/z policy; container runtime publish syntax. All TBD in architecture/implementation tasks. Enforcement: JR must preserve boundary unless a later authorized architecture decision replaces it..

## OS.js Management Port (assigned OSJS-003

The shell listens on TCP 18209 by default:

- `OSJS/src/server/config.js` — default port (PORT env var overrides(;
- `OSJS/Dockerfile` — EXPOSE 18209;
- `OSJS/README.md` — current standalone run instructions (port recorded in docs/NETWORK_EXPOSURE.md(;
- `deploy/docker-compose.yml` — standalone shell deployment publishes `18209` via `ports:` (OSJS_PORT env override(;, OSJS_DATA_DIR=/data volume-mounted persistence (named volume osjs-data(, /healthz healthcheck..

The full OS.js desktop UI (HTTP + WebSocket( serves on this port. In the final appliance deployment, this listener shares host network namespace under required host networking; Docker `ports:` publishing not required for the appliance container. This resolves the OS.js TCP port only; MMA2 Modbus TCP port numbers remain architecture-task decisions per the directive..
