# Active Work

This directory contains only human-promoted work authorized for JR execution.

## Execution States

Every executable task file in this directory must have exactly one of these states:

- `ACTIVE` — the one task JR executes now.
- `QUEUED` — already human-authorized, waiting its turn. No additional approval is required when it becomes ACTIVE.

Completed or retired tasks belong in `workflow/archive/`, not here.

## Current-Task Invariant

```text
EXACTLY ONE ACTIVE = execute it
ZERO ACTIVE        = stop; no executable current task
TWO OR MORE ACTIVE = stop; invalid Active Work state
```

JR must never infer the current task from filename order, task number, dependency sorting, handoff prose, or Planning.

`workflow/active_work/` is the authority for current-task selection. `handoff.md` is only a continuation summary and must agree with it. A disagreement is a conflict: stop rather than guess.

## Ordered Authorized Work

When a human promotes an ordered sequence, promotion authorizes the entire sequence but marks only its first executable task `ACTIVE`. Remaining authorized tasks are `QUEUED`.

Each queued task must explicitly name its predecessor. Each non-final task should explicitly name its successor using `Next:`. This makes advancement deterministic without dependency scheduling.

When the ACTIVE task completes successfully:

```text
ACTIVE task verified
→ archive completed task
→ read its explicit Next
→ verify that exact task is QUEUED and its Previous names the completed task
→ change that task to ACTIVE
→ update handoff summary
→ commit and push the completion/advancement state together
```

If `Next` is absent, the authorized sequence is complete and JR stops after persisting completion state.

If `Next` is named but missing, not QUEUED, or its `Previous` does not match, stop and report invalid Active Work state.

## ICC Context Rule

Active Work is consumed through Operation CWAL. ICC is read-only to CWAL. If required ICC context is stale or missing, CWAL delegates that semantic branch to BLACK SHEEP WALL. Only BLACK SHEEP WALL may modify ICC.

BLACK SHEEP WALL cannot select Active Work, inspect Planning to choose work, grant authority, or widen the active task.

Do not place future backlog, brainstorm material, or speculative work here.
