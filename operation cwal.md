# Operation CWAL

## Purpose

Operation CWAL is the mandatory execution-discipline mode for MCS.OSJS. It prevents context drift and unauthorized work.

When the user says **Operation CWAL**, immediately follow:

```text
OPERATION CWAL
→ READ AGENTS.md
→ READ BLACK_SHEEP_WALL.md
→ READ ICC/INDEX.md FIRST
→ LOCATE AUTHORITY
→ SELECT ACTIVE TASK SEMANTIC BRANCH
→ SET THAT BRANCH AS THE NAVIGATION CEILING
→ VALIDATE ONLY RELEVANT ICC INSIDE THAT BRANCH
→ USE ICC IF CURRENT
→ REFRESH ONLY STALE AFFECTED CONTEXT INSIDE THAT BRANCH
→ FOLLOW WORKFLOW
→ ACT
→ VERIFY
→ NEVER GUESS
```

## 1. ICC-First Does Not Mean ICC-Wide

CWAL reads `ICC/INDEX.md` before reopening repository source files, but it must not validate or refresh the whole ICC registry.

The index is used to locate context after execution authority is known. Registry presence is not relevance.

After reading the index, CWAL immediately locates the current Active Work and selects the semantic branch required by that task. That branch becomes the navigation ceiling.

```text
READ ICC INDEX
→ LOCATE ACTIVE WORK
→ SELECT TASK BRANCH
→ VALIDATE TASK BRANCH ONLY
```

If current HEAD differs from the ICC baseline, compare the baseline to HEAD only far enough to determine whether declared dependencies of the selected branch changed. Do not inventory, inspect, or repair stale sibling contexts.

If HEAD is unchanged, inspect only uncommitted dependencies inside the selected branch whose state differs from the audited overlay.

A stale ICC context outside the Active Work branch is irrelevant to the current execution. Do not refresh it, open it, or repair it.

BLACK SHEEP WALL delegated from CWAL operates only inside this selected branch.

## 2. Locate Authority

Read `handoff.md` and locate the current Active Work task.

For execution state, the authoritative source is:

```text
workflow/active_work/
```

`handoff.md` is the current execution handoff and must agree with Active Work.

Do not read Planning or future tasks to decide what to execute.

If authority cannot be located or conflicts cannot be resolved from repository rules, stop and report the conflict.

## 3. Select Semantic Branch

The active task selects the semantic branch.

Allowed navigation:

```text
ACTIVE TASK
→ SELECT ITS ICC CONTEXT
→ ZOOM IN TO REQUIRED CHILD DETAIL
→ EXECUTE
```

Not allowed:

```text
ACTIVE TASK
→ ICC INDEX
→ CHECK EVERY STALE CONTEXT
→ REPAIR GOVERNANCE
→ REPAIR LICENSING
→ REPAIR NETWORKING
→ REPAIR PLANNING
→ RETURN TO TASK
```

CWAL may cross into another semantic branch only when the active task itself explicitly requires that boundary for its stated outcome or verification.

## 4. Follow Workflow

Execute only the current human-authorized Active Work task and only within its defined scope.

Do not promote work, invent work, expand scope, or silently resolve architectural questions.

When Active Work contains ordered microtasks, complete and verify the current microtask before reading implementation detail for later microtasks unless that later detail is an explicit dependency of the current task.

## 5. Act

Investigate, implement, test, and document only what the current active microtask authorizes.

Repository source files are opened when synchronized ICC context inside the selected branch is insufficient for the implementation detail or when changed source inside that branch requires direct inspection.

Do not perform unrelated ICC maintenance while executing product work.

## 6. Verify

Verify acceptance criteria against repository/runtime evidence. Do not mark work complete merely because code changed.

On completion, update the repository state required by the workflow, including Active Work and `handoff.md`.

## 7. Never Guess

If required authority, selected-branch context, dependencies, or evidence are missing: stop and report the missing authority or dependency rather than inferring it.
