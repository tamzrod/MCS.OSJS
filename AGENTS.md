# AGENTS.md

## Purpose

This file is the bootstrap/router for agents working in MCS.OSJS. It is not a project diary and must not become an append-only history file.

## First Rule

For every repository-dependent operation:

```text
LOCATE COMMAND
→ READ BLACK_SHEEP_WALL.md
→ RUN BLACK SHEEP WALL WITHIN INVOKING SCOPE
→ LOCATE AUTHORITY
→ USE VALID ICC CONTEXT
→ ACT WITHIN OPERATION SCOPE
→ VERIFY AGAINST REPOSITORY SOURCE
→ NEVER GUESS
```

BLACK SHEEP WALL is mandatory context infrastructure for repository-dependent work. It is not optional merely because the invoking operation already names files to read.

## Commands

- `Operation CWAL` → read `operation cwal.md`; that operation must invoke BLACK SHEEP WALL before consuming execution context.
- `BLACK SHEEP WALL` → read `BLACK_SHEEP_WALL.md` and perform the context-maintenance operation directly.

## Mandatory Context Prelude

Brainstorm, Planning, Microtask, Promotion, Operation CWAL, verification, and other repository-dependent operations must begin by reading `BLACK_SHEEP_WALL.md` and running its context prelude within the scope of the invoking operation.

BLACK SHEEP WALL:

- validates the ICC context needed by the operation;
- refreshes stale or missing context only when required;
- returns control to the invoking operation after context is valid;
- does not grant authority, select work, or widen scope.

An operation must not bypass BLACK SHEEP WALL by reading ICC files directly and assuming they are current.

## Workflow Boundaries

### Brainstorm / Planning

Run BLACK SHEEP WALL for the planning scope first. Planning is human-owned. Brainstorming may define problems, alternatives, contracts, and proposed tasks, but it does not authorize implementation.

### Microtask

Run BLACK SHEEP WALL for the relevant planning/task scope first. Microtasks translate approved intent into small, independently verifiable work items. Follow `planning/microtask/rules.md`.

### Promotion

Run BLACK SHEEP WALL for the selected task and workflow state first. Promotion moves a human-selected microtask into `workflow/active_work/` and updates `handoff.md`.

### Operation CWAL

Run BLACK SHEEP WALL for the Active Work scope first. CWAL executes only authorized Active Work. It must not inspect Planning to select future work.

## Repository Authority

Current repository files are authoritative. Conversation memory and historical summaries are not substitutes for repository state.
