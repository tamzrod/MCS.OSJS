# AGENTS.md

## Purpose

This file is the bootstrap/router for agents working in MCS.OSJS. It is not a project diary and must not become an append-only history file.

## First Rule

For every repository-dependent operation:

```text
LOCATE COMMAND
→ READ BLACK_SHEEP_WALL.md
→ READ ICC/INDEX.md FIRST
→ VALIDATE ICC BASELINE AGAINST CURRENT HEAD + WORKING TREE
→ USE ICC IF CURRENT
→ IF STALE, REFRESH ONLY AFFECTED CONTEXT THROUGH BLACK SHEEP WALL
→ LOCATE AUTHORITY
→ ACT WITHIN OPERATION SCOPE
→ VERIFY AGAINST REPOSITORY SOURCE
→ NEVER GUESS
```

BLACK SHEEP WALL is mandatory context infrastructure. ICC is the first-read context cache.

## BLACK SHEEP WALL Modes

### Direct Invocation

When the user says `BLACK SHEEP WALL`, perform repository-wide context compaction:

```text
current HEAD
→ audit repository context
→ compact established repository truth into ICC
→ stamp ICC with the HEAD commit used as baseline
→ record audited uncommitted file overlays
```

If an ICC baseline already exists, do not rescan unchanged repository files. Refresh only files changed since the recorded baseline and uncommitted files whose current content differs from the last audited overlay.

### Invoked by Another Operation

The operation reads ICC first. If the relevant context is synchronized with current HEAD and any audited working-tree overlay, use ICC without reopening source files. If it is stale or missing, BLACK SHEEP WALL refreshes only the affected context, then returns control to the invoking operation.

BLACK SHEEP WALL never grants authority, selects work, or widens scope.

## Workflow Boundaries

### Brainstorm / Planning

Use ICC first. Refresh affected context through BLACK SHEEP WALL only when the relevant ICC state is stale or missing. Planning is human-owned and does not authorize implementation.

### Microtask

Use ICC first, refresh only stale affected context, then follow `planning/microtask/rules.md`.

### Promotion

Use ICC first, refresh only stale affected context, then promote only the human-selected microtask and synchronize `handoff.md`.

### Operation CWAL

Use ICC first for Active Work context. Refresh only stale affected context. CWAL executes only authorized Active Work and must not inspect Planning to select future work.

## Repository Authority

Repository files are authoritative. ICC is a compact cache of repository truth and may be trusted only when its recorded baseline and working-tree overlay match the repository state being used.
