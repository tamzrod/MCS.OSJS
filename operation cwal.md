# Operation CWAL

## Purpose

Operation CWAL is the mandatory execution-discipline mode for MCS.OSJS. It prevents context drift and unauthorized work while allowing already-authorized Active Work to proceed continuously.

**ICC ownership rule:** Operation CWAL never creates, edits, refreshes, patches, regenerates, or commits ICC state. `BLACK_SHEEP_WALL.md` is the only operation authorized to update `ICC/`.

When the user says **Operation CWAL**, immediately follow:

```text
OPERATION CWAL
→ READ AGENTS.md
→ READ BLACK_SHEEP_WALL.md
→ READ ICC/INDEX.md FIRST
→ READ workflow/active_work/README.md
→ LOCATE THE ONE TASK WITH Status: ACTIVE
→ VERIFY handoff.md AGREES
→ SELECT ITS SEMANTIC BRANCH
→ USE ICC AS READ-ONLY CONTEXT
→ IF REQUIRED ICC CONTEXT IS STALE/MISSING: INVOKE BLACK SHEEP WALL FOR THAT BRANCH
→ IMPLEMENT ONLY ACTIVE
→ TEST
→ VERIFY ACCEPTANCE CRITERIA
→ ARCHIVE COMPLETED ACTIVE TASK
→ ADVANCE ONLY ITS EXPLICIT Next FROM QUEUED TO ACTIVE
→ UPDATE handoff.md
→ COMMIT
→ PUSH TO main
→ VERIFY origin/main CONTAINS THE COMPLETION COMMIT
→ CONTINUE IF EXACTLY ONE ACTIVE TASK EXISTS
→ OTHERWISE STOP
→ NEVER GUESS
```

## 1. Select Current Work Deterministically

`workflow/active_work/` is the authority for current-task selection.

Apply exactly this rule:

```text
EXACTLY ONE Status: ACTIVE = execute it
ZERO Status: ACTIVE        = stop; no executable current task
TWO OR MORE ACTIVE         = stop; invalid Active Work state
```

Never infer current work from filename order, task number, dependency sorting, `handoff.md` prose, Planning, or repository history.

`Status: QUEUED` means already human-authorized and waiting its explicit turn. It does not require another human approval when its predecessor completes.

After locating the one ACTIVE task, read `handoff.md` only to verify continuation state. If handoff disagrees with Active Work, stop and report the conflict. Active Work remains the execution authority; CWAL must not silently repair or reinterpret the disagreement.

Do not inspect Planning or future work to choose execution targets.

## 2. ICC Is Read-Only to CWAL

CWAL may read ICC as compressed repository context, but it must never modify ICC itself.

After locating the ACTIVE microtask, select only the semantic branch required by that task. That branch becomes the navigation ceiling.

```text
ACTIVE MICROTASK
→ SELECT ITS ICC BRANCH
→ READ CURRENT CONTEXT
→ NEED MORE DETAIL? ZOOM IN INSIDE THAT BRANCH
→ STALE/MISSING? INVOKE BLACK SHEEP WALL FOR THAT BRANCH
→ RESUME CWAL USING THE RESULT
```

CWAL must not implement BLACK SHEEP WALL's delta, refresh, patch, baseline, overlay, or regeneration algorithm itself.

A stale ICC context outside the selected branch is irrelevant to the current microtask.

## 3. BLACK SHEEP WALL Delegation

BLACK SHEEP WALL is the sole writer and maintainer of ICC.

CWAL may invoke BLACK SHEEP WALL only when the selected task requires ICC context that is stale or missing. The selected semantic branch is the delegation boundary.

```text
CWAL NEEDS CONTEXT
→ ICC CURRENT? USE READ-ONLY
→ ICC STALE/MISSING? DELEGATE SELECTED BRANCH TO BLACK SHEEP WALL
→ BLACK SHEEP WALL OWNS ANY ICC CHANGE
→ RETURN TO CWAL
```

CWAL does not perform a post-task or post-push ICC update. A successful implementation push is a complete CWAL checkpoint without any ICC mutation afterward.

## 4. Semantic Navigation Boundary

Allowed navigation:

```text
ACTIVE MICROTASK
→ SELECT ITS ICC CONTEXT
→ ZOOM IN TO REQUIRED CHILD DETAIL
→ EXECUTE
```

CWAL may cross into another semantic branch only when the ACTIVE microtask explicitly requires that boundary for its stated outcome or verification.

Do not scan unrelated ICC branches or repair unrelated repository state.

## 5. Execute Only ACTIVE

Investigate, implement, test, and document only what the one ACTIVE microtask authorizes.

Do not execute QUEUED tasks early. Do not promote work from Planning, invent work, expand scope, or silently resolve architectural questions.

Repository source files are opened only when synchronized ICC context inside the selected branch is insufficient for implementation or when changed source inside that branch requires direct inspection.

## 6. Deterministic Completion and Advancement

A microtask is complete only after implementation, tests, acceptance verification, required evidence, workflow-state update, commit, push, and `origin/main` verification.

After the ACTIVE task passes its acceptance criteria:

```text
READ ACTIVE TASK'S Next
→ ARCHIVE COMPLETED ACTIVE TASK
→ Next: none ? NO SUCCESSOR
→ Next: <ID> ? VERIFY THAT EXACT FILE IS QUEUED
→ VERIFY SUCCESSOR'S Previous MATCHES COMPLETED TASK
→ CHANGE ONLY THAT SUCCESSOR TO ACTIVE
→ UPDATE handoff.md TO MATCH
→ COMMIT COMPLETION + ADVANCEMENT STATE
→ PUSH main
→ VERIFY origin/main
```

CWAL must not search for an eligible successor, sort QUEUED tasks, or solve a dependency graph.

If `Next` names a missing task, a non-QUEUED task, or a task whose `Previous` does not name the completed task, stop and report invalid Active Work state.

If `Next: none`, archive the completed task, update handoff to no current task, commit/push/verify, then stop.

The push to `main` is the checkpoint boundary between microtasks.

## 7. Continuous Active-Work Loop

After the completion/advancement push is verified:

```text
COUNT Status: ACTIVE
        │
        ├── 1 → execute that ACTIVE task
        ├── 0 → stop
        └── 2+ → stop: invalid Active Work state
```

Do not request new human authorization for a task that was already `QUEUED` and became `ACTIVE` through the explicit predecessor/Next transition.

Do not inspect Planning when Active Work drains.

## 8. Never Guess

If required authority, task-state consistency, predecessor/successor linkage, selected-branch context, dependencies, runtime evidence, push verification, or repository state is missing, stop and report it rather than inferring it.

If a push fails, `origin/main` cannot be verified, or completion state is not safely persisted, do not advance execution.
