# THE GATHERING

## Purpose

`THE GATHERING` is the read-only committed-status command for MCS.OSJS. It reports the current workflow position from the repository's committed `HEAD` only.

> Gather committed status. Ignore everything outside the commit. Make no changes.

## Invocation Boundary

THE GATHERING runs only when the user explicitly invokes it or another directive explicitly delegates to it. Completion of OPERATION CWAL, BLACK SHEEP WALL, generic repository work, or any other operation does not implicitly invoke THE GATHERING.

## Authority Boundary

Inspect only the committed state needed for four views: **Active Work, Handoff, Microtask Planning, Brainstorm**. `ICC/INDEX.md`, the relevant entry in `ICC/manifest.json` and committed workflow context under `ICC/context/` may route or summarize those views. Authoritative committed fallback sources are restricted to `handoff.md`, `workflow/active_work/`, `planning/microtask/`, and `planning/Brainstorm/`. Do not expand beyond these workflow surfaces to make a report.

## Commit Rule

Current committed `HEAD` is the complete observation boundary. Ignore dirty working-tree state, uncommitted overlays, staged changes, temporary files, local runtime state, product source and changes outside committed HEAD. Do not compare ICC against the working tree or reconcile committed status with local/uncommitted state. ICC overlay metadata used by other operations is irrelevant here and must not initiate inspection.

## Current-State Only Rule

This is not a historical completion audit. Do not inspect `workflow/archive/`, implementation source, old task histories, commit history beyond identifying current HEAD or prior completion commits. Only follow an explicit reference from one of the four permitted status surfaces when strictly necessary to interpret that current status, then read the minimum. ZERO ACTIVE means report ZERO ACTIVE; do not search history.

## ICC read-only validity rule — mandatory

Follow the THE GATHERING exception in `workflow/ICC_CONSUMER.md`. Read the committed index/manifest for navigation. Use cached workflow status as a current finding **only when the selected node's committed source dependencies have demonstrably been reconciled with the observed HEAD**. A `null` baseline, `unverified`, `partial`, `stale`, a remote-only historical reviewed HEAD, or unavailable committed comparison cannot certify current status. Do not infer freshness from a file's existence, an index-level snapshot, or the absence of an explicit stale flag.

For each view whose cached status is not demonstrably current, bypass the cached conclusion and read **only** its matching authoritative committed source: Active Work → `workflow/active_work/`; Handoff → `handoff.md`; Microtask Planning → `planning/microtask/`; Brainstorm → `planning/Brainstorm/`. For actual execution identity, report what the committed task and handoff say, not what stale ICC says. If they disagree, report the conflict rather than resolve it. Reading a permitted fallback is normal successful Gathering, not BLOCKED solely because ICC was migrated as unverified.

Never refresh ICC, invoke BLACK SHEEP WALL, run the ICC checker, certify the working tree, or write any file. Do not read product source to establish workflow status.

## Report Flow

```text
THE GATHERING
→ IDENTIFY CURRENT COMMITTED HEAD
→ READ COMMITTED ICC/INDEX.md + RELEVANT MANIFEST NODE STATE
→ FOR EACH VIEW: USE ICC ONLY IF COMMITTED DEPENDENCIES ARE PROVEN CURRENT
→ OTHERWISE READ MATCHING AUTHORITATIVE COMMITTED FALLBACK
→ ACTIVE WORK → HANDOFF → MICROTASK PLANNING → BRAINSTORM
→ REPORT ACTUAL STATUS / ANY CONFLICT
→ MAKE NO CHANGES → STOP
```

## Required Report

### Active Work

Show the committed Active Work inventory and identify the task currently authorized for execution. Dependency, waiting, blocked, next, or completed information may be shown only when already represented inside Active Work or Handoff; do not build these categories from other surfaces.

### Handoff

Show current authorized work and explicit next/dependency information when present.

### Microtask Planning

Show committed planning inventory/state. Planning is not promoted or executable unless committed workflow authority explicitly says so.

### Brainstorm

Show committed brainstorm topics and represented state. Brainstorm is exploratory and does not authorize implementation.

## Output Shape

```text
MCS.OSJS — THE GATHERING
HEAD: <commit>

ACTIVE WORK
<status>

HANDOFF
Current: <current authorized work>
Next:    <explicit next/dependency state if present>

MICROTASK PLANNING
<planning inventory/status>

BRAINSTORM
<brainstorm inventory/status>
```

Keep the report concise. This is a current committed-status view, not a repository audit or project history report.

## Core Invariant

THE GATHERING = explicit invocation + committed HEAD + ICC-first navigation with proved committed validity or authoritative committed fallback + four workflow views + report + STOP. It is NOT working-tree inspection, uncommitted reconciliation, BLACK SHEEP WALL, a repository/historical audit, product inspection, mutation or an implicit post-CWAL action. If permitted committed sources disagree, report the disagreement. Never guess.
