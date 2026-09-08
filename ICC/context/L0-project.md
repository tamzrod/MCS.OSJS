# L0 Project Context

Baseline commit: 57c8714
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

## Current Execution State (handoff)

Handoff status: COMPLETE. MMA2-001/002 and Simulator tasks SIM-001 through SIM-008 are completed+verified; SIM-009 was superseded; the corrected SIM-010 → SIM-017 sequence is completed and verified.

SIM-017's real-MMA2 capstone completed at `57c8714`, proving independent boot, restore, safe apply/restart, changing Raw Ingest schedules, FC1–FC4 reads, negative boundaries, foreign preservation, and reboot resume. No current Active Work microtask remains. Detail in Zoom In `active-work`.

## Component Boundaries (starting hypotheses until revised by repository authority

- OS.js:  presentation/desktop shell — runnable base shell exists at OSJS/ (neutral, no SCADA backend coupling(.
- Orchestrator:  lifecycle/control authority — not yet implemented.
- Modbus Replicator:  acquisition/replication — not yet implemented.
- MMA2: deterministic Modbus memory appliance — imported, built, runtime-smoke-tested, and activated for simulator-owned configuration and raw ingest.
- Simulator (new program): OS.js Modbus device simulator using MMA2 as memory/runtime engine — configuration, ownership, scheduling, raw ingest, OS.js editor, Save & Apply routing, and runtime status are complete through SIM-007. It is not part of the base four-component identity.



## Unresolved / Open

- Orchestrator/Replicator architecture, internal transport, and persistence model remain planning-stage questions.
- MMA2 Modbus TCP port numbers remain architecture-task decisions governed by the network directive.
