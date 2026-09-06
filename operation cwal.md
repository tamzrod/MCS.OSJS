# Operation CWAL

## Purpose

Operation CWAL is the mandatory execution-discipline mode for MCS.OSJS. It prevents context drift and unauthorized work.

When the user says **Operation CWAL**, immediately follow:

```text
OPERATION CWAL
→ READ AGENTS.md
→ READ BLACK_SHEEP_WALL.md
→ READ ICC/INDEX.md FIRST
→ VALIDATE RELEVANT ICC AGAINST CURRENT HEAD + WORKING TREE
→ USE ICC IF CURRENT
→ REFRESH ONLY STALE AFFECTED CONTEXT
→ LOCATE AUTHORITY
→ FOLLOW WORKFLOW
→ ACT
→ VERIFY
→ NEVER GUESS
```

## 1. ICC-First Context

CWAL reads ICC before reopening repository source files.

If the relevant ICC context is synchronized with the recorded baseline commit and audited working-tree overlay, use it directly.

If current HEAD differs from the ICC baseline, compare the baseline to HEAD and refresh only context affected by changed committed files. If HEAD is unchanged, inspect only uncommitted files whose current state differs from the recorded audited overlay.

Do not perform a full repository rescan unless the ICC baseline is missing, unusable, or cannot be incrementally reconciled.

BLACK SHEEP WALL does not grant execution authority, choose a task, inspect Planning to select future work, or widen Active Work.

## 2. Locate Authority

Read `handoff.md` and locate the current Active Work task.

For execution state, the authoritative source is:

```text
workflow/active_work/
```

`handoff.md` is the current execution handoff and must agree with Active Work.

Do not read Planning or future tasks to decide what to execute.

If authority cannot be located or conflicts cannot be resolved from repository rules, stop and report the conflict.

## 3. Follow Workflow

Execute only the current human-authorized Active Work task and only within its defined scope.

Do not promote work, invent work, expand scope, or silently resolve architectural questions.

## 4. Act

Investigate, implement, test, and document only what the active task authorizes.

Repository source files are opened when the synchronized ICC context is insufficient for the implementation detail or when changed source requires direct inspection.

## 5. Verify

Verify acceptance criteria against repository/runtime evidence. Do not mark work complete merely because code changed.

On completion, update the repository state required by the workflow, including Active Work and `handoff.md`.

## 6. Never Guess

If required authority, context, dependencies, or evidence are missing: stop and report the missing authority rather than inferring it.
