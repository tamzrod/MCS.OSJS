# THE GATHERING

## Purpose

`THE GATHERING` is the read-only workflow status command for MCS.OSJS.

It gathers the project's current workflow state into one concise report so the human can see what is active, what is being planned, what is pending or blocked, and where execution currently stands.

> Gather status. Do not change status.

## Authority

`THE GATHERING` has explicit read authority across the workflow boundaries required to produce the report:

- `ICC/INDEX.md`;
- relevant workflow context under `ICC/context/`;
- `handoff.md`;
- `workflow/active_work/`;
- `planning/microtask/`;
- workflow items that are pending, blocked, or completed when those states are represented by authoritative workflow sources.

This cross-boundary authority is read-only and exists only for status reporting.

`THE GATHERING` cannot:

- execute Active Work;
- promote a microtask;
- create or modify planning;
- choose the next task;
- reorder work;
- change task status;
- widen an existing task;
- refresh ICC context;
- invoke BLACK SHEEP WALL as a maintenance action;
- modify `handoff.md`;
- modify repository files.

## Context Rule

Repository truth remains authoritative, but ICC is the primary status source for THE GATHERING.

Start with `ICC/INDEX.md` and the relevant workflow context under `ICC/context/`. Most status reporting should be satisfied from this context without re-reading the full repository workflow surface.

Use current HEAD and working-tree state only to determine whether the relevant ICC context still matches the repository state.

If the relevant ICC context is current, report from ICC.

If the relevant ICC context is stale, missing, incomplete, or conflicts with repository state:

1. do not refresh it;
2. do not invoke BLACK SHEEP WALL to repair it;
3. read only the minimum authoritative workflow sources needed to establish current status;
4. report the stale or conflicting context explicitly;
5. continue the status report using the best available repository truth.

A stale ICC context is a reportable condition, not authorization to modify anything.

Do not perform a repository-wide audit for a status report.

## Report Flow

```text
THE GATHERING
→ READ ICC/INDEX.md
→ SELECT WORKFLOW STATUS BOUNDARY
→ READ RELEVANT ICC CONTEXT
→ CHECK ICC BASELINE AGAINST CURRENT REPOSITORY STATE
→ IF CURRENT: REPORT FROM ICC
→ IF STALE/INCOMPLETE: READ MINIMUM AUTHORITATIVE SOURCES REQUIRED
→ REPORT ANY STALE OR CONFLICTING CONTEXT
→ REPORT STATUS
→ MAKE NO CHANGES
→ STOP
```

BLACK SHEEP WALL may be mentioned as the mechanism that can later refresh stale ICC context, but THE GATHERING must never run that refresh itself.

## Required Report

The report should be concise and organized around these views.

### Active Work

Show the current Active Work inventory and identify the task currently authorized for execution according to repository truth.

Prefer the relevant ICC workflow context when current. Include dependency/blocking state where it is explicitly represented.

### Microtask Planning

Show the current microtask planning inventory and its represented state.

Prefer ICC context when it already contains the current planning state. Read `planning/microtask/` only when needed to fill a missing or stale status boundary.

Do not interpret a planned microtask as promoted or executable unless repository authority explicitly says so.

### Pending / Blocked

Show work that is waiting, pending, or blocked when that state can be established from current ICC context or the authoritative workflow sources.

Do not invent a queue from file ordering alone.

### Handoff / Current Position

Show the current handoff position, including the current task and explicitly established next/dependency information.

Prefer current ICC context. Read `handoff.md` directly when validation or stale context requires it.

### Completed

Summarize completed workflow items when useful to understand current position. Keep this compact; THE GATHERING is primarily a current-status view, not a project history report.

## Stale Context Reporting

When ICC is stale or conflicts with repository truth, include a compact warning such as:

```text
CONTEXT WARNING
ICC workflow context is stale against current repository state.
Status below was read from the minimum authoritative workflow sources required.
BLACK SHEEP WALL refresh is required separately.
```

Do not repair the condition during THE GATHERING.

## Output Shape

Conceptually:

```text
MCS.OSJS — THE GATHERING

ACTIVE WORK
<current active-work status>

MICROTASK PLANNING
<current microtask planning status>

PENDING / BLOCKED
<waiting or blocked work established by repository truth>

HANDOFF
Current: <current authorized work>
Next:    <explicit next/dependency state, if established>

COMPLETED
<compact recent/relevant completion state>
```

Exact formatting may vary. Repository state must not.

## Core Invariant

```text
THE GATHERING
= OBSERVE + GATHER + REPORT

NOT
= PLAN + PROMOTE + EXECUTE + REFRESH + MODIFY
```

If repository sources disagree, report the disagreement rather than resolving it by assumption.

Never guess.
