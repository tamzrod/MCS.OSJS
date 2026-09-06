# Brainstorm: Port Configuration and Socket Deployment

Status: brainstorm material. No architectural decision is made here. This topic frames problems, alternatives, and open questions for later planning.

## Problem Framing

MCS.OSJS is a single-container Modbus consolidation appliance with separately bounded internal components: OS.js, Orchestrator, Modbus Replicator, and MMA2.



Every candidate connection between these components, and from the appliance to the outside world, passes through a socket bound to a port.in

- Which components expose sockets, and which only connect outward?
- How are port numbers assigned, validated,and persisted?
- Who owns port configuration authority: one central config, or per-component config?
- When are sockets bound, and what happens when the bound fails or the port vanishes?
- How do internal sockets(map to container ports, and which ports are exposed externally?
- tcp vs udp; unix sockets vs network sockets - which is appropriate where?

These are prerequisites for the future program architecture and rewiring plan (construction stage 4,and must be planned before contracts are written.

## Constraint Input: MMA2 Host Network Operation

Status: input from the human for brainstorming. Not a decision; not a contract yet.

MMA2 must operate on the host network adapter, not merely inside the container's internal network, because it needs to filter IP traffic.in

This changes the socket deployment picture, because MMA2's sockets cannot be treated like ordinary container-internal sockets..



Planning questions this raises (left open here:

- Does "host network adapter" mean all host interfaces, or a specific adapter set that MMA2 filters?
- Is MMA2 passive (observing/mirroring IP) or active(intercepting/forwarding IP)? Raw sockets, TUN/TAP, BPF/eBPF,and iptables hooks have very different socket and privilege requirements. The filter model must be known before socket contracts can be planned.

- If MMA2 runs in the host network namespace, how does it communicate with container-resident components (Orchestrator, Replicator, OS.js)?
- Under which OS privileges does MMA2 run,and what does that imply for the single-container appliance model?
- Which MMA2 sockets bind on host interfaces,and which (if any) remain container-internal?(and how is that distinction configured)
- Does the host-adapter requirement force a different deployment shape than "one container"?(e.g., host network mode, a privileged sidecar, or a host-level process) or can it be satisfied inside the container model?

## Initial Scope Candidates

Brainstorm only. No selection is implied by listing:

- Modbus acquisition sockets (Replicator side)
- Internal inter-process sockets between Orchestrator, Replicator,and MMA2
- OS.js presentation socket(s)and the web/desktop entry point
- Socket address/port configuration schema and validation rules
- Port conflict and reuse semantics (bind failure, restart, rebind,
- Container port mapping and external exposure policy
- Interface binding policy (loopback vs all interfaces; later TLS/auth if out-of-scope initially)

## Known Open Questions

- Is the Modbus acquisition path TCP, serial, or both - and what does that imply for socket types?
- Is internal IPC network sockets, Unix domain sockets, or a mix? (single-container constraint influences this.
- Should there be a single configuration authority for all ports (e.g., Orchestrator-owned), or per-component config?
- What is the default port namespace,and how are collisions detected?
- Who owns socket lifecycle (bind, accept, shutdown, restart), and how does it interact with component lifecycle?
- Which ports (if any) are exposed outside the container, and what is the mapping policy?
- Does initial planning exclude auth/TLS on sockets,orinlude it as a contract from the start?

## Candidate Directions (not decisions)

- A single configuration source per appliance, validated centrally, consumed by each component at its own startup.

- Fixed well-known defaults with overrides in an appliance-level configuration, kept explicit in the eventual program architecture plan.

## Out of Scope Here

- Choosing the design (left for the architecture planning stage)
- Microtask decomposition (comes after planning assumes shape)
- Any implementation authorization (brainstorm grants none)
