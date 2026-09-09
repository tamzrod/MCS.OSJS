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
→ READ handoff.md
→ LOCATE CURRENT ACTIVE MICROTASK
→ SELECT ITS SEMANTIC BRANCH
→ USE ICC AS READ-ONLY CONTEXT
→ IF REQUIRED ICC CONTEXT IS STALE/MISSING: INVOKE BLACK SHEEP WALL FOR THAT BRANCH
→ FOLLOW WORKFLOW
→ IMPLEMENT
→ TEST
→ VERIFY ACCEPTANCE CRITERIA
→ UPDATE EVIDENCE / ACTIVE STATE
→ UPDATE handoff.md
→ COMMIT
→ PUSH TO main
→ VERIFY origin/main CONTAINS THE COMPLETION COMMIT
→ IF ANOTHER ALREADY-AUTHORIZED ACTIVE MICROTASK EXISTS: CONTINUE AUTOMATICALLY
→ OTHERWISE STOP
→ NEVER GUESS
```

## 1. Locate Authority First

Read `handoff.md` and locate the current Active Work microtask.

For execution state, the authoritative source is:

```text
workflow/active_work/
```

`handoff.md` is the persisted execution handoff and must agree with Active Work.

Do not read Planning or future work to choose execution targets.

If authority cannot be located or conflicts cannot be resolved from repository rules, stop and report the conflict.

Human authorization happens when work is promoted into Active Work. Once multiple microtasks are already authorized there, CWAL does not require another human approval between them.

## 2. ICC Is Read-Only to CWAL

CWAL may read ICC as compressed repository context, but it must never modify ICC itself.

After locating the active microtask, select only the semantic branch required by that task. That branch becomes the navigation ceiling.

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

Not allowed:

```text
ACTIVE MICROTASK
→ ICC INDEX
→ CHECK EVERY STALE CONTEXT
→ REPAIR GOVERNANCE
→ REPAIR LICENSING
→ REPAIR NETWORKING
→ REPAIR PLANNING
→ RETURN TO TASK
```

CWAL may cross into another semantic branch only when the active microtask itself explicitly requires that boundary for its stated outcome or verification.

## 5. Execute Only Current Authorized Work

Investigate, implement, test, and document only what the current active microtask authorizes.

Do not promote work, invent work, expand scope, or silently resolve architectural questions.

When Active Work contains ordered microtasks, do not preload implementation detail for later microtasks unless that detail is an explicit dependency of the current one.

Repository source files are opened only when synchronized ICC context inside the selected branch is insufficient for implementation or when changed source inside that branch requires direct inspection.

## 6. Microtask Completion Boundary

A microtask is not complete merely because the code works locally.

A microtask is complete only when all of the following have happened:

```text
IMPLEMENTATION COMPLETE
→ TEST COMPLETE
→ ACCEPTANCE CRITERIA VERIFIED
→ REQUIRED EVIDENCE RECORDED
→ ACTIVE WORK STATE UPDATED AS REQUIRED
→ handoff.md UPDATED
→ COMPLETION STATE COMMITTED
→ COMMIT PUSHED TO main
→ origin/main VERIFIED TO CONTAIN THAT COMMIT
```

The push to `main` is the checkpoint boundary between microtasks.

Never begin the next microtask before the completed microtask and its handoff state are successfully pushed to `main` and verified there.

## 7. Continuous Active-Work Loop

After the push is verified, inspect only the already-authorized Active Work state to determine whether another microtask is next.

```text
MICROTASK COMPLETE
        ↓
UPDATE handoff.md
        ↓
COMMIT
        ↓
PUSH main
        ↓
VERIFY origin/main
        ↓
NEXT ALREADY-AUTHORIZED ACTIVE MICROTASK?
        │
        ├── YES
        │    ↓
        │  SELECT ITS SEMANTIC BRANCH
        │    ↓
        │  READ ICC AS READ-ONLY CONTEXT
        │    ↓
        │  IF STALE/MISSING: DELEGATE THAT BRANCH TO BLACK SHEEP WALL
        │    ↓
        │  IMPLEMENT → TEST → VERIFY
        │    ↓
        │  UPDATE handoff → COMMIT → PUSH → VERIFY
        │    ↓
        │  LOOP
        │
        └── NO
             ↓
            STOP
```

Do not stop between already-authorized Active Work microtasks merely to request human authorization again.

Do not inspect Planning to find additional work when Active Work becomes empty.

## 8. Never Guess

If required authority, selected-branch context, dependencies, runtime evidence, push verification, or repository state is missing, stop and report the missing authority or dependency rather than inferring it.

If a push fails, `origin/main` cannot be verified, or the completion state is not safely persisted, do not advance to the next microtask.
