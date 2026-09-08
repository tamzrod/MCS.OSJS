# AGENTS.md

## Purpose

This file is the bootstrap/router for agents working in MCS.OSJS.

It defines **how to locate authority**, not the detailed behavior of each operation. Detailed rules belong in their directive files. Do not turn this file into project history, troubleshooting notes, implementation recipes, or an append-only diary.

## 1. Resolve the Command First

Before applying any generic repository workflow, identify whether the user invoked a StarCraft directive.

Exact directive phrases are command selectors:

| Command | Authority | Mode |
| --- | --- | --- |
| `BLACK SHEEP WALL` | `BLACK_SHEEP_WALL.md` | Context audit / ICC synchronization |
| `OPERATION CWAL` | `operation cwal.md` | Continuous execution of already-authorized Active Work |
| `THE GATHERING` | `the gathering.md` | Read-only committed workflow status |

When a StarCraft directive is invoked:

```text
IDENTIFY DIRECTIVE
→ READ AGENTS.md
→ READ THAT DIRECTIVE FILE
→ FOLLOW ITS OWN READ / WRITE / STOP BOUNDARIES
→ DO NOT IMPORT BEHAVIOR FROM ANOTHER DIRECTIVE UNLESS IT EXPLICITLY CALLS IT
→ NEVER GUESS
```

A directive file overrides the generic bootstrap where its rules are more specific.

## 2. Generic Repository Operation

For repository-dependent work that is **not** a StarCraft directive:

```text
READ ICC/INDEX.md FIRST
→ CHECK WHETHER RELEVANT ICC CONTEXT MATCHES CURRENT REPOSITORY STATE
→ USE ICC IF CURRENT
→ IF STALE OR MISSING, USE BLACK_SHEEP_WALL.md ONLY FOR THE AFFECTED CONTEXT
→ LOCATE AUTHORITATIVE REPOSITORY SOURCE
→ ACT WITHIN USER-AUTHORIZED SCOPE
→ VERIFY AGAINST REPOSITORY SOURCE
→ NEVER GUESS
```

Repository files are authoritative. ICC is a compact context cache and is trusted only within the baseline/overlay rules defined by BLACK SHEEP WALL.

## 3. StarCraft Boundaries

### BLACK SHEEP WALL

Owns repository-context auditing and ICC synchronization.

It does **not** grant implementation authority, select future work, or widen scope.

### OPERATION CWAL

Owns execution discipline.

CWAL executes only work already authorized in `workflow/active_work/`. It must not inspect Planning to choose future work.

Completion of one microtask is **not** a stop condition. If another already-authorized Active Work microtask exists, CWAL continues automatically according to `operation cwal.md`.

### THE GATHERING

Owns committed workflow-status reporting.

THE GATHERING is read-only, observes committed `HEAD`, ignores uncommitted state, makes no changes, and must **not** invoke BLACK SHEEP WALL unless its own directive is explicitly changed to allow it.

## 4. Workflow Authority

```text
Planning / Brainstorm
    = human-owned future or exploratory work

planning/microtask/
    = defined work awaiting promotion

workflow/active_work/
    = implementation authority

handoff.md
    = persisted execution position / continuation state

ICC/
    = compressed context, never execution authority
```

Promotion into Active Work is the authorization boundary. Planning alone never authorizes implementation.

## 5. Permanent Invariants

1. Resolve the invoked operation before choosing tools or context paths.
2. Read the smallest authoritative context required for the operation.
3. Do not widen scope because related work is visible.
4. Do not treat ICC, Planning, Brainstorm, or historical notes as implementation authority.
5. Do not duplicate detailed directive logic in `AGENTS.md`; route to the directive file.
6. Do not store task-specific discoveries, runbooks, debugging history, or completed-work summaries here.
7. Verify claims against the repository state permitted by the active operation.
8. If authoritative sources conflict and the active directive does not define resolution, report the conflict instead of guessing.
