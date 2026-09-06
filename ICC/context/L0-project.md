# L0 Project Context

Baseline commit: 500376cfb5c222298aadfcf035aad0af0a635773
Working tree: clean
Source dependencies: README.md, PROJECT_IDENTITY.md, handoff.md
Parent: none
Zoom In: governance, donor-licensing, network-exposure, planning-workflow, osjs-shell, active-work
Zoom Out: none

## Identity

MCS.OSJS = Modbus Consolidation System; standalone repo/lineage rebuilt around an OS.js application shell
and proven Modbus runtime components. Planned as single-container appliance: OS.js presentation, Orchestrator
lifecycle/control authority, Modbus Replicator acquisition, MMA2 deterministic Modbus memory appliance. It is NOT part of Nameless SCADA, but reuses proven generic parts from it( and may later be migrated into it after proving standalone-first.

## Construction Strategy

Staged: scaffold -> donor inventory(Nameless SCADA, then Modbus Replicator stack( -> program architecture+rewiring -> UI plan -> microtasks -> promotion -> Operation CWAL execution -> verify appliance. Completed so far:  scaffolding, Nameless SCADA donor inventory + OS.js shell harvest (executed,head of sim/MMA2 work(; Modbus Replicator stack inventory and remaining architecture/UI stages proceed through planning+workflow.

## Development Model

Brainstorm -> microtask -> promotion -> CWAL -> verified change. Context maintenance via BLACK SHEEP WALL and ICC. One task = one primary outcome. Reuse proven machinery, but preserve explicit boundaries; no component gains authority merely for convenience. Standalone-first:  stabilize the standalone appliance before any optional Nameless SCADA migration.



## Authority Boundaries

Project identity document defines identity/direction only:  no implementation authority, no settled architecture. Detailed architecture, donor selection, rewiring, UI behavior, and execution scope established via planning + workflow docs. Repository files are authoritative; conversation memory is not a substitute.

## Current Execution State (handoff

Handoff status: ACTIVE. Active Work execution order:

1. `workflow/active_work/mma2-basic-install-test.md` — MMA2-001 (import+build( then MMA2-002 (runtime write/read/restart smoke test(, preserving internal order as repository truth.
2-8. `workflow/active_work/sim-001…sim-007.md` — simulator program in dependency order SIM-001 → SIM-007. A later task does not authorize skipping an incomplete dependency. JR executes only work present in workflow/active_work/ + reflected in handoff; Planning is not used to choose or widen work.

## Component Boundaries (starting hypotheses until revised by repository authority

- OS.js:  presentation/desktop shell — runnable base shell exists at OSJS/ (neutral, no SCADA backend coupling(.
- Orchestrator:  lifecycle/control authority — not yet implemented.
- Modbus Replicator:  acquisition/replication — not yet implemented.
- MMA2:  deterministic Modbus memory appliance — import/activation planned in active work (MMA2-001/2, SIM-002/SIM-004(.
- Simulator (new program(: OS.js Modbus device simulator using MMA2 as memory/runtime engine — active work SIM-001…SIM-007. Not part of the base four-component identity;and an OS.js application planned from the approved simulator brainstorm.



## Unresolved / Open

- Shared MMA2 configuration-authority implementation (who composes/validates effective MMA2 config from multiple producers( is not yet settled (only a boundary contract in the simulator brainstorm..
- Orchestrator/Replicator architecture, internal transport, persistence model:  remain planning-stage questions..
- MMA2 Modbus TCP port numbers:  architecture-task decisions per network directive.

