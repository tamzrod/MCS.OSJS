# External Network Exposure Directive

Baseline commit: cb869da3caeed040478e67ecb0a01c5f93b3a66b
Working tree: clean
Source dependencies: docs/NETWORK_EXPOSURE.md

## Rule (project directive

Appliance exposes only the interfaces that are part of the external product contract:

1. One OS.js management/operator endpoint
2. One or more MMA2 Modbus TCP endpoints (configuration-dependent, N >= 1 when Modbus TCP enabled



External port count = 1 OS.js port + N MMA2 Modbus TCP ports. Orchestrator and Modbus Replicator are internal-only and must not receive host-published ports merely for internal communication. Internal port != externally published appliance port; do not expose internal service at container/host boundary solely because donor code communicates over TCP.



## Internal Communication

No specific internal transport required: internal TCP listener, loopback, Unix-domain socket, in-process API, or other local mechanism are candidate choices; architecture decides.

 MMA2 may retain host network/adapter for externally connected Modbus clients, incl. client-IP filtering behavior; external listeners remain true network boundaries requiring separate listeners/policy. Internal components must not masquerade as external Modbus clients to satisfy client-IP filtering model. Exact internal MMA2 access path = open architectural question, not settled by this directive.

 Non-goals of directive: exact port numbers, MMA2 listener count, memory-space-to-listener mapping, internal transport choice, TLS/reverse-proxy, auth/z policy, container runtime publish syntax;all TBD in architecture/implementation tasks.

 Enforcement: JR must preserve boundary unless later authorized architecture decision replaces it.