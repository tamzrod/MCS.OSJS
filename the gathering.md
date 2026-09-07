# THE GATHERING

## Purpose

`THE GATHERING` is the read-only workflow status command for MCS.OSJS.

It gathers the project's current workflow state into one concise report so the human can see what is active, what is being planned, what is pending or blocked, and where execution currently stands.

> Gather status. Do not change status.

## Authority

`THE GATHERING` has explicit read authority across the workflow boundaries required to produce the report:

- `handoff.md`;
- `workflow/active_work/`;
- `planning/microtask/`;
- workflow items that are pending, blocked, or completed when those states are represented by the authoritative workflow sources.

This cross-boundary read authority exists only for status reporting.

`THE GATHERING` cannot:

- execute Active Work;
- promote a microtask;
- create or modify planning;
- choose the next task;
- reorder work;
- change task status;
- widen an existing task;
- modify `handoff.md`;
- modify repository files.

## Context Rule

Repository truth remains authoritative.

Use ICC first where a current workflow context exists. Validate the relevant ICC baseline against current HEAD and working-tree state. If the relevant workflow context is stale, refresh only the affected workflow context through BLACK SHEEP WALL before reporting it.

Do not perform a repository-wide audit for a status report.

## Report Flow

```text
THE GATHERING
→ READ BLACK_SHEEP_WALL.md
→ READ ICC/INDEX.md
→ SELECT WORKFLOW STATUS BOUNDARY
→ VALIDATE RELEVANT ICC CONTEXT
→ REFRESH ONLY AFFECTED WORKFLOW CONTEXT IF STALE
→ READ AUTHORITATIVE WORKFLOW SOURCES AS REQUIRED
→ REPORT STATUS
→ MAKE NO CHANGES
→ STOP
```

## Required Report

The report should be concise and organized around these views.

### Active Work

Show the current Active Work inventory and identify the task currently authorized for execution according to repository truth.

Include dependency/blocking state where it is explicitly represented.

### Microtask Planning

Show the current microtask planning inventory and its represented state.

Do not interpret a planned microtask as promoted or executable unless repository authority explicitly says so.

### Pending / Blocked

Show work that is waiting, pending, or blocked when that state can be established from the authoritative workflow sources.

Do not invent a queue from file ordering alone.

### Handoff / Current Position

Show the current handoff position, including the current task and explicitly established next/dependency information.

### Completed

Summarize completed workflow items when useful to understand current position. Keep this compact; THE GATHERING is primarily a current-status view, not a project history report.

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
= PLAN + PROMOTE + EXECUTE + MODIFY
```

If repository sources disagree, report the disagreement rather than resolving it by assumption.

Never guess.
