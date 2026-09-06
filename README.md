# MCS.OSJS

MCS.OSJS is the new Modbus Consolidation System rebuilt around an OS.js application shell and proven Modbus runtime components.

This repository starts with the project scaffolding first. Application code, donor-code harvesting, rewiring, and UI implementation follow only through the defined planning and execution workflow.

## Construction Model

```text
IDEA
  ↓
BRAINSTORM
  ↓
MICROTASK
  ↓
PROMOTION
  ↓
OPERATION CWAL
  ↓
VERIFIED SOFTWARE
```

Context maintenance is supported by BLACK SHEEP WALL and ICC.

## Initial Product Direction

The intended appliance is composed of four clearly separated responsibilities:

```text
OS.js         = presentation / desktop shell
Orchestrator  = lifecycle and control authority
Replicator    = Modbus acquisition and replication
MMA2          = deterministic Modbus memory appliance
```

These boundaries are starting hypotheses for planning and must remain explicit until proven or revised by repository authority.
