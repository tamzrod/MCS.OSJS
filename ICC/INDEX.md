# ICC Index

## Purpose

This directory is the Incremental Context Compaction (ICC) cache for MCS.OSJS.

ICC provides compact, dependency-scoped repository understanding for agents.

## Access Rule

Repository-dependent operations consume ICC through BLACK SHEEP WALL.

Before an operation relies on ICC context, it must read `BLACK_SHEEP_WALL.md` and run that prelude within the invoking operation's scope. BLACK SHEEP WALL determines whether the required context is valid, stale, or missing and performs only the necessary refresh.

Directly reading an ICC context file does not by itself establish that the context is current.

## Lifecycle

```text
INVOKING OPERATION
→ BLACK SHEEP WALL
→ AUDIT REQUIRED CONTEXT
→ COMPACT
→ CACHE
→ repository changes
→ DEPENDENCY CHECK
→ INVALIDATE AFFECTED ONLY
→ REFRESH WHEN REQUIRED
→ RETURN TO INVOKING OPERATION
```

## Rules

- One context file = one semantic boundary.
- Context files summarize current established state, contracts, dependencies, and unresolved questions.
- Context files do not store chat history.
- Each context declares the repository sources that materially define it.
- Known-stale context is never consumed.
- Missing context is created only when required by an operation.
- Unrelated valid context is left untouched.
- BLACK SHEEP WALL is the mandatory validation/maintenance gateway for repository-dependent ICC consumption.
- BLACK SHEEP WALL inherits the invoking operation's scope and never grants authority.

## Zoom Model

```text
L0 — broadest project/system view
 ↓
L1
 ↓
...
 ↓
LX — as deep as required
```

Parent files may use `## Zoom In`; children may use `## Zoom Out`.

## Registry

No semantic context has been audited yet. The repository is at scaffolding stage.

Create context files under `ICC/context/` only after their source boundaries exist and can be audited.
