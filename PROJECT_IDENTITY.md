# Project Identity

## Project

**MCS.OSJS**

**MCS** means **Modbus Consolidation System**.

MCS.OSJS is a new implementation lineage of the original MCS concept. It is a standalone project and repository. It is **not** part of Nameless SCADA, although it is expected to reuse proven generic parts from Nameless SCADA and may later be migrated into Nameless SCADA as an application.

## Repository

**GitHub:** `tamzrod/MCS.OSJS`  
**Default branch:** `main`

## Purpose

MCS.OSJS is intended to become a self-contained Modbus consolidation appliance with an OS.js-based operator environment.

The project will combine proven components from existing work, rewire them into a new application, and place them behind a coherent operator UI rather than rebuilding every subsystem from zero.

The immediate objective is not to reproduce Nameless SCADA. The objective is to build a focused Modbus Consolidation System that can stand alone first.

## Product Direction

The intended deployment direction is a **single software appliance / single container** containing separately bounded internal components.

Current planned component set:

```text
MCS.OSJS appliance
│
├── OS.js
├── Orchestrator
├── Modbus Replicator
└── MMA2
```

A single container does not imply a single monolithic runtime. Internal responsibilities and failure boundaries should remain explicit.

The exact contracts, process ownership, configuration authority, lifecycle behavior, persistence model, and internal communication paths are architectural decisions that must be planned before implementation. They must not be silently inferred from this identity document.

## Source Lineage

MCS.OSJS is expected to harvest reusable parts from two existing systems.

### Nameless SCADA

Primary expected donor area:

- OS.js shell and generic desktop/application infrastructure.

Only reusable, application-neutral parts should be harvested by default. Nameless SCADA-specific architecture must not be imported merely because the code exists there.

### Modbus Replicator Stack

Primary expected donor areas:

- Modbus Replicator;
- MMA2;
- existing deterministic Modbus behavior;
- proven configuration, status, polling, writing, and memory-contract behavior where appropriate.

Existing proven semantics should be preserved unless a deliberate architectural decision changes them.

## Construction Strategy

The project is intentionally built in stages.

```text
1. ERECT SCAFFOLDING
        ↓
2. INVENTORY NAMELESS SCADA DONOR PARTS
        ↓
3. INVENTORY MODBUS REPLICATOR DONOR PARTS
        ↓
4. PLAN PROGRAM ARCHITECTURE + REWIRING
        ↓
5. PLAN UI
        ↓
6. DECOMPOSE INTO MICROTASKS
        ↓
7. PROMOTE AUTHORIZED WORK
        ↓
8. EXECUTE THROUGH OPERATION CWAL
        ↓
9. VERIFY THE APPLIANCE
```

The scaffolding exists to support construction. It is not the application itself.

## Development Model

MCS.OSJS uses the project discipline established by its repository scaffolding:

```text
BRAINSTORM
    ↓
MICROTASK
    ↓
PROMOTION
    ↓
OPERATION CWAL
    ↓
VERIFIED CHANGE
```

Repository truth is supported by BLACK SHEEP WALL and ICC context maintenance.

The development model exists so that JR can execute bounded, independently verifiable work without needing the entire application in working context at once.

## Task Principle

> **One task = one primary outcome.**

Work should be decomposed until each task can be independently understood, implemented, verified, and completed.

Complexity should be reduced by decomposition rather than by making execution instructions increasingly complicated.

## Architectural Principle

> **Reuse proven machinery, but preserve explicit boundaries.**

The project should distinguish clearly between:

- reusable donor code;
- newly designed application glue;
- runtime authority;
- presentation/UI;
- Modbus acquisition and replication;
- Modbus memory behavior.

No component should gain authority merely because it is convenient to connect it there.

## Standalone-First Principle

MCS.OSJS should become a complete standalone application before any later migration into Nameless SCADA is treated as a requirement.

The desired long-term path is:

```text
MCS.OSJS standalone appliance
        ↓
stabilize behavior + contracts + UI
        ↓
prove the application independently
        ↓
optionally migrate the mature MCS capability
into Nameless SCADA as an application
```

Future Nameless SCADA integration should therefore be treated as a later transplantation of a mature application capability, not as an excuse to couple MCS.OSJS to Nameless SCADA during its initial construction.

## Authority

This document defines project identity and high-level direction only.

It does **not** authorize implementation and does not settle unresolved architecture.

Detailed architecture, donor selection, rewiring, UI behavior, and execution scope must be established through the repository's planning and workflow documents.

Where a design question remains unresolved, JR must not silently choose a design and present it as established project truth.

## Project Operating Principle

> **Build the scaffolding first, harvest proven parts second, design the rewiring explicitly, then build the UI and application through small verified tasks.**
