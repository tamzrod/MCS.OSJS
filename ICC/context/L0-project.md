# L0 Project Context

Baseline commit: cb869da3caeed040478e67ecb0a01c5f93b3a66b
Working tree: clean
Source dependencies: README.md, PROJECT_IDENTITY.md, handoff.md

## Identity

MCS.OSJS = Modbus Consolidation System; standalone repo/lineage rebuilt around an OS.js application shell
and proven Modbus runtime components. Planned as single-container appliance: OS.js presentation, Orchestrator
lifecycle/control authority, Modbus Replicator acquisition, MMA2 deterministic Modbus memory appliance.

Boundaries are starting hypotheses for planning and must remain explicit until revised by repository authority.

It is NOT part of Nameless SCADA, but expected to reuse proven generic parts from it( and may later be migrated into it as an application after proving standalone-first).

## Construction Strategy

Staged: scaffold -> donor inventory (Nameless SCADA, then Modbus Replicator stack) -> program architecture+rewiring -> UI plan -> microtasks -> promotion -> Operation CWAL execution -> verify appliance.



## Development Model

Brainsstorm -> microtask -> promotion -> CWAL -> verified change. Context maintenance via BLACK SHEEP WALL and ICC. One task = one primary outcome

## Authority Boundaries

Project identity document defines identity/direction only:. no implementation authority, no settled architecture. Detailed architecture, donor selection, rewiring, UI behavior, and execution scope established via planning + workflow docs. Repository files are authoritative; conversation memory is not a substitute. Reuse proven machinery, but preserve explicit boundaries; no component gains authority merely for convenience.

Handoff status: IDLE; no active work present in workflow/active_work/.