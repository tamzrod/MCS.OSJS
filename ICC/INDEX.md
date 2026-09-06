# ICC Index

## Purpose

This directory is the Incremental Context Compaction (ICC) cache for MCS.OSJS.

ICC provides compact, dependency-scoped repository understanding for agents.

## Lifecycle

```text
AUDIT
→ COMPACT
→ CACHE
→ repository changes
→ DEPENDENCY CHECK
→ INVALIDATE AFFECTED ONLY
→ REFRESH
```

## Rules

- One context file = one semantic boundary.
- Context files summarize current established state, contracts, dependencies, and unresolved questions.
- Context files do not store chat history.
- Each context declares the repository sources that materially define it.
- Known-stale context is never consumed.
- Missing context is created only when required by an operation.
- Unrelated valid context is left untouched.

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
