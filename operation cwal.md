# Operation CWAL

## Purpose

Operation CWAL is the mandatory execution-discipline mode for MCS.OSJS. It prevents context drift and unauthorized work.

When the user says **Operation CWAL**, immediately follow:

```text
OPERATION CWAL
→ READ AGENTS.md
→ READ BLACK_SHEEP_WALL.md
→ RUN BLACK SHEEP WALL FOR ACTIVE-WORK SCOPE
→ LOCATE AUTHORITY
→ READ VALID EXECUTION CONTEXT
→ FOLLOW WORKFLOW
→ ACT
→ VERIFY
→ NEVER GUESS
```

## 1. Run BLACK SHEEP WALL

Before consuming ICC context or executing repository work, read `BLACK_SHEEP_WALL.md` and run its context-maintenance prelude for the current Active Work scope.

BLACK SHEEP WALL validates or refreshes only the context required by CWAL. It does not grant execution authority, choose a task, inspect Planning to select future work, or widen Active Work.

After the required context is valid, return to Operation CWAL.

## 2. Locate Authority

Read `handoff.md` and locate the current Active Work task.

For execution state, the authoritative source is:

```text
workflow/active_work/
```

`handoff.md` is the current execution handoff and must agree with Active Work.

Do not read Planning or future tasks to decide what to execute.

If authority cannot be located or conflicts cannot be resolved from repository rules, stop and report the conflict.

## 3. Read Valid Execution Context

Use `ICC/INDEX.md` and only the minimum relevant ICC context validated by BLACK SHEEP WALL for the authorized Active Work.

Known-stale context must not be consumed.

## 4. Follow Workflow

Execute only the current human-authorized Active Work task and only within its defined scope.

Do not promote work, invent work, expand scope, or silently resolve architectural questions.

## 5. Act

Investigate, implement, test, and document only what the active task authorizes.

## 6. Verify

Verify acceptance criteria against repository/runtime evidence. Do not mark work complete merely because code changed.

On completion, update the repository state required by the workflow, including Active Work and `handoff.md`.

## 7. Never Guess

If required authority, context, dependencies, or evidence are missing: stop and report the missing authority rather than inferring it.
