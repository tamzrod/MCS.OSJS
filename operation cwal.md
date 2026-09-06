# Operation CWAL

## Purpose

Operation CWAL is the mandatory execution-discipline mode for MCS.OSJS. It prevents context drift and unauthorized work.

When the user says **Operation CWAL**, immediately follow:

```text
OPERATION CWAL
→ READ CONTEXT
→ LOCATE AUTHORITY
→ FOLLOW WORKFLOW
→ ACT
→ VERIFY
→ NEVER GUESS
```

## 1. Read Context

Read `AGENTS.md`, `ICC/INDEX.md`, the minimum relevant valid ICC context, `handoff.md`, and the current Active Work task.

Do not read Planning or future tasks to decide what to execute.

## 2. Locate Authority

For execution state, the authoritative source is:

```text
workflow/active_work/
```

`handoff.md` is the current execution handoff and must agree with Active Work.

If authority cannot be located or conflicts cannot be resolved from repository rules, stop and report the conflict.

## 3. Follow Workflow

Execute only the current human-authorized Active Work task and only within its defined scope.

Do not promote work, invent work, expand scope, or silently resolve architectural questions.

## 4. Act

Investigate, implement, test, and document only what the active task authorizes.

## 5. Verify

Verify acceptance criteria against repository/runtime evidence. Do not mark work complete merely because code changed.

On completion, update the repository state required by the workflow, including Active Work and `handoff.md`.

## 6. Never Guess

If required authority, context, dependencies, or evidence are missing: stop and report the missing authority rather than inferring it.
