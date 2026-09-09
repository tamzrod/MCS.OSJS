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
→ RUN TASK-SPECIFIC TESTS
→ RUN REPOSITORY-NATIVE VERIFICATION FOR CHANGED EXECUTABLE SURFACES
→ VERIFY ACCEPTANCE CRITERIA
→ VERIFY INTENDED CHANGESET
→ ARCHIVE COMPLETED ACTIVE TASK
→ ADVANCE ONLY ITS EXPLICIT Next FROM QUEUED TO ACTIVE
→ UPDATE handoff.md
→ STAGE + REVIEW INTENDED CHANGESET
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

## 6. Pre-Completion Exit Gate

Implementation, focused tests, and completion verification are separate gates.

Before an ACTIVE task may become COMPLETED:

```text
ACTIVE IMPLEMENTATION
→ RUN TASK-SPECIFIC TESTS
→ RUN REPOSITORY-NATIVE VERIFICATION FOR EACH CHANGED EXECUTABLE SURFACE
→ VERIFY ACCEPTANCE CRITERIA
→ VERIFY INTENDED CHANGESET
→ ONLY THEN ARCHIVE / ADVANCE
```

Repository-native verification means the smallest existing repository command that actually parses, compiles, bundles, builds, or otherwise validates the changed production surface in the way that repository normally consumes it.

Examples of weaker checks that do not replace a native build gate:

- `node --check`;
- standalone parser invocation;
- isolated helper/unit tests;
- grep or raw-byte inspection;
- lint alone.

They may supplement the native gate but may not replace it.

If the required native gate cannot run:

```text
KEEP TASK ACTIVE
→ DO NOT ARCHIVE
→ DO NOT ADVANCE Next
→ DO NOT CLAIM COMPLETION
→ REPORT BLOCKED: REQUIRED VERIFICATION UNAVAILABLE
→ STOP
```

If the required native gate fails:

```text
KEEP TASK ACTIVE
→ FOLLOW AGENTS.md EXECUTION GUARDRAIL
→ DO NOT ARCHIVE
→ DO NOT ADVANCE Next
→ DO NOT PUSH A COMPLETION CLAIM
```

A higher-fidelity failed gate invalidates any earlier weaker verification claim for that affected surface.

## 7. Commit Integrity Gate

Before a completion commit:

```text
INSPECT WORKING TREE
→ STAGE THE INTENDED TASK CHANGESET EXPLICITLY
→ VERIFY STAGED PATHS INCLUDE NEW / RENAMED / DELETED FILES
→ REVIEW STAGED DIFF
→ CONFIRM NO REQUIRED TASK FILE IS UNTRACKED OR OMITTED
→ COMMIT
```

Do not rely on `git commit -am` as staging authority when a task may create files.

## 8. Deterministic Completion and Advancement

A microtask is complete only after implementation, task-specific tests, repository-native verification, acceptance verification, required evidence, workflow-state update, commit, push, and `origin/main` verification.

After the ACTIVE task passes its complete exit gate:

```text
READ ACTIVE TASK'S Next
→ ARCHIVE COMPLETED ACTIVE TASK
→ Next: none ? NO SUCCESSOR
→ Next: <ID> ? VERIFY THAT EXACT FILE IS QUEUED
→ VERIFY SUCCESSOR'S Previous MATCHES COMPLETED TASK
→ CHANGE ONLY THAT SUCCESSOR TO ACTIVE
→ UPDATE handoff.md TO MATCH
→ PASS COMMIT INTEGRITY GATE
→ COMMIT COMPLETION + ADVANCEMENT STATE
→ PUSH main
→ VERIFY origin/main
```

CWAL must not search for an eligible successor, sort QUEUED tasks, or solve a dependency graph.

If `Next` names a missing task, a non-QUEUED task, or a task whose `Previous` does not name the completed task, stop and report invalid Active Work state.

If `Next: none`, archive the completed task, update handoff to no current task, commit/push/verify, then stop.

The push to `main` is the checkpoint boundary between microtasks.

## 9. Continuous Active-Work Loop

After the completion/advancement push is verified:

```text
COUNT Status: ACTIVE
        │
        ├── 1 → execute that ACTIVE task
        ├── 0 → stop
        └── 2+ → stop: invalid Active Work state
```

A successor may be activated only after the predecessor's complete exit gate, workflow-state commit, push, and `origin/main` verification all pass.

If a higher-fidelity verification later proves that a pushed completion commit is invalid, CWAL must stop immediately. It must not execute the successor. Report the contradictory repository/workflow state and await recovery direction.

Do not request new human authorization for a task that was already `QUEUED` and became `ACTIVE` through the explicit predecessor/Next transition.

Do not inspect Planning when Active Work drains.

CWAL does not implicitly invoke THE GATHERING or any other StarCraft directive when its Active Work loop ends.

## 10. Never Guess

If required authority, task-state consistency, predecessor/successor linkage, selected-branch context, dependencies, repository-native verification, runtime evidence, push verification, or repository state is missing, stop and report it rather than inferring it.

If a push fails, `origin/main` cannot be verified, required verification is unavailable, or completion state is not safely persisted, do not advance execution.
