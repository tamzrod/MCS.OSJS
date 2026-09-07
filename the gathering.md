# THE GATHERING

## Purpose

`THE GATHERING` is the read-only committed-status command for MCS.OSJS.

It reports the current workflow position from the repository's committed `HEAD` only.

> Gather committed status. Ignore everything outside the commit. Make no changes.

## Authority Boundary

`THE GATHERING` may inspect only the committed state needed for these four views:

1. **Active Work**
2. **Handoff**
3. **Microtask Planning**
4. **Brainstorm**

ICC is the primary status context and should satisfy most reads.

Primary ICC sources are:

- `ICC/INDEX.md`
- relevant committed workflow context under `ICC/context/`, especially context for Active Work, planning/microtasks, and brainstorm topics

Authoritative committed fallback sources are limited to:

- `handoff.md`
- `workflow/active_work/`
- `planning/microtask/`
- `planning/Brainstorm/`

Do not expand beyond these workflow surfaces merely to produce a status report.

## Commit Rule

The committed repository state at current `HEAD` is the complete observation boundary.

`THE GATHERING` must ignore:

- dirty working-tree state;
- uncommitted files or overlays;
- staged-but-uncommitted changes;
- temporary files;
- local runtime state;
- implementation/code outside the four workflow status surfaces;
- repository changes that are not part of current committed `HEAD`.

Do not compare ICC against the working tree.

Do not reconcile committed status with local/uncommitted state.

If ICC records working-tree or overlay information for another workflow, that information is irrelevant to THE GATHERING and must not trigger further inspection.

## ICC Rule

Use committed ICC first.

If the relevant ICC context at `HEAD` already contains the needed status, report from it.

If ICC is missing or insufficient for one of the four allowed views, read only the matching committed authoritative source:

- Active Work → `workflow/active_work/`
- Handoff → `handoff.md`
- Microtask Planning → `planning/microtask/`
- Brainstorm → `planning/Brainstorm/`

Do not refresh ICC.

Do not invoke BLACK SHEEP WALL.

Do not treat an ICC baseline mismatch as authorization to inspect outside the committed status boundary. Report only what current committed `HEAD` establishes.

## Report Flow

```text
THE GATHERING
→ IDENTIFY CURRENT COMMITTED HEAD
→ READ COMMITTED ICC STATUS CONTEXT
→ ACTIVE WORK
→ HANDOFF
→ MICROTASK PLANNING
→ BRAINSTORM
→ READ MATCHING COMMITTED FALLBACK SOURCE ONLY IF ICC IS INSUFFICIENT
→ REPORT
→ MAKE NO CHANGES
→ STOP
```

## Required Report

### Active Work

Show the committed Active Work inventory and identify the task currently authorized for execution.

Dependency, waiting, blocked, next, or completed information may be shown when it is already represented inside Active Work or Handoff. Do not search additional workflow surfaces to construct those categories.

### Handoff

Show the committed handoff position:

- current authorized work;
- explicit next/dependency information when present.

### Microtask Planning

Show the committed microtask planning inventory and represented state.

Planning is not promoted or executable unless the committed workflow authority explicitly says so.

### Brainstorm

Show the committed brainstorm inventory/topics and their represented state.

Brainstorm material is exploratory and does not authorize implementation.

## Output Shape

Conceptually:

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

```text
THE GATHERING
= COMMITTED HEAD
+ ICC FIRST
+ ACTIVE WORK
+ HANDOFF
+ MICROTASK
+ BRAINSTORM
+ REPORT

THE GATHERING
!= WORKING TREE
!= UNCOMMITTED STATE
!= BLACK SHEEP WALL
!= REPOSITORY AUDIT
!= IMPLEMENTATION INSPECTION
!= MODIFY
```

If the four allowed committed sources disagree, report the disagreement rather than resolving it by assumption.

Never guess.
