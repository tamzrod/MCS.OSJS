# Project Identity

## Project

**MCS.OSJS**

MCS means **Modbus Consolidation System**. MCS.OSJS is a new implementation lineage, not the Nameless SCADA repository.

## Repository

**GitHub:** `tamzrod/MCS.OSJS`

**Default branch:** `main`

## Purpose

Build a deterministic Modbus consolidation appliance by combining proven reusable components with a new OS.js-based operator surface.

The project is expected to harvest reusable parts from:

- Nameless SCADA, especially generic OS.js shell/application infrastructure;
- the Modbus Replicator stack, especially Replicator and MMA2 behavior.

## Initial Construction Plan

1. Establish project scaffolding and execution discipline.
2. Inventory reusable Nameless SCADA parts.
3. Inventory reusable Modbus Replicator stack parts.
4. Plan the application architecture and rewiring.
5. Plan the UI.
6. Decompose work into independently verifiable tasks.
7. Execute through Operation CWAL.

## Initial Appliance Model

```text
OS.js
  └─ presentation / operator applications

Orchestrator
  └─ lifecycle / control authority

Replicator
  └─ acquisition / replication

MMA2
  └─ deterministic Modbus memory
```

These are initial planning boundaries, not permission to silently invent contracts. Architectural decisions must be explicit and repository-backed.

## Development Principle

> Keep the boundaries explicit, the work small, the execution state unambiguous, and the evidence verifiable.
