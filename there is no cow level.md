# THERE IS NO COW LEVEL

## Purpose

`THERE IS NO COW LEVEL` is the emergency sandbox-rescue command for MCS.OSJS.

Use it when the current JR/agent session or sandbox has become unreliable, confused, corrupted, stuck in a workaround loop, or otherwise unsafe to continue.

Its job is not to finish the current task. Its job is to preserve valid work, create a trustworthy continuation point on `main`, and make the current sandbox disposable.

> Preserve current truth. Record exactly where work stopped. Push to main. Stop.

## Rescue Override

Invocation immediately stops normal feature-development and debugging behavior.

An unfinished, failing, malformed, or partially verified implementation is valid handoff state. Rescue does not require implementation to be fixed or tests to pass.

Once invoked:

```text
STOP FEATURE DEVELOPMENT
→ STOP REPAIR ATTEMPTS
→ STOP TEST-FIX LOOPS
→ STOP TOOL DIAGNOSIS
→ CHECKPOINT CURRENT TRUTH
```

Do not delay rescue to make the task cleaner, passing, complete, or easier for the next sandbox. Do not rewrite broken tests, chase formatter/compiler/vet/lint/test failures, diagnose JR/OpenHands/editor/tool behavior, retry failed authoring approaches, invent workaround chains, or complete the current microtask first.

Instead, record the exact unfinished or broken state, relevant failure, and next action in `handoff.md`, then checkpoint and push it.

## Authority

`THERE IS NO COW LEVEL` may inspect current `HEAD`, working-tree state, Active Work, `handoff.md`, and changed files required to understand the checkpoint. It may update `handoff.md`, commit the intended checkpoint, and push it to `main`.

It does not gain authority to invent product behavior, widen implementation scope, promote planning, redesign implementation, or finish unrelated work.

It must not edit `ICC/` directly. If ICC maintenance is genuinely required for the rescue, request only the smallest affected branch through `BLACK SHEEP WALL`.

## Checkpoint Rule

Inspect only enough state to establish what must be preserved. Classify obvious temporary probes, scratch files, or known failed-workaround debris separately so they do not obscure the checkpoint. If intent is uncertain, do not investigate deeply; record the uncertainty in `handoff.md` and choose the option that loses the least recoverable work.

Verification is observational, not corrective. Use existing evidence first. If a minimum verification command is needed and fails:

```text
RECORD FAILURE
→ DO NOT FIX IT
→ CONTINUE HANDOFF
```

Do not rerun failing verification merely to obtain a passing rescue state. Never label failing or unverified work as complete.

## Handoff Update

Before committing, update `handoff.md` so a fresh sandbox can continue without reconstructing the broken session. Record:

- current Active Work item;
- exact implementation point reached;
- state preserved;
- known unfinished, malformed, failing, or unverified work;
- verification already performed and result;
- exact failing command/error when relevant;
- unresolved or excluded changes that matter;
- exact next action for the new sandbox;
- that the previous sandbox was rescued and is not authoritative beyond the rescue commit.

Keep `handoff.md` as continuation state, not a transcript or postmortem.

## Commit and Push Rule

```text
REVIEW CHECKPOINT DIFF
→ ENSURE handoff.md TRUTHFULLY DESCRIBES IT
→ COMMIT INTENDED CHECKPOINT STATE
→ PUSH TO main
→ CONFIRM REMOTE main CONTAINS THE RESCUE COMMIT
```

Do not require clean tests or completed implementation before committing the rescue checkpoint. Do not stop at a local commit. If push fails, do not claim rescue completion; report the exact failure and preserve the local commit SHA.

## Sandbox Destruction Boundary

After the rescue commit is confirmed on remote `main`, output:

```text
RESCUE COMPLETE
main: <rescue commit SHA>
old sandbox: SAFE TO DESTROY
new sandbox: CONTINUE FROM handoff.md
```

Then stop all implementation work in the old sandbox. `THERE IS NO COW LEVEL` does not itself destroy an external/cloud sandbox unless the environment explicitly provides that capability.

If rescue cannot safely reach `main`, stop mutating, preserve the best known local state, report status and exact blocker, and do not claim the sandbox is safe to destroy.

## Command

```text
THERE IS NO COW LEVEL
→ FREEZE FEATURE DEVELOPMENT + DEBUGGING
→ READ CURRENT ACTIVE WORK + HANDOFF + RELEVANT REPOSITORY STATE
→ INSPECT HEAD + STATUS + CURRENT DIFF
→ IDENTIFY WHAT MUST BE PRESERVED
→ DO NOT FIX FAILING IMPLEMENTATION OR TESTS
→ RECORD UNFINISHED/BROKEN/UNVERIFIED STATE IN handoff.md
→ RECORD EXACT FAILURE + NEXT ACTION WHEN RELEVANT
→ REVIEW CHECKPOINT DIFF
→ COMMIT CHECKPOINT STATE
→ PUSH TO main
→ CONFIRM REMOTE main HAS RESCUE COMMIT
→ REPORT SAFE TO DESTROY + NEW-SANDBOX CONTINUATION
→ STOP
```

## Core Invariant

```text
THERE IS NO COW LEVEL
= FREEZE + PRESERVE CURRENT TRUTH + HANDOFF + COMMIT + PUSH MAIN + STOP

THERE IS NO COW LEVEL
!= FINISH FEATURE
!= FIX TESTS FIRST
!= MAKE BUILD PASS FIRST
!= TOOL-FAILURE INVESTIGATION
!= SPECULATIVE REWRITE
!= NEW SCOPE
```

An unfinished state is a valid handoff state.

Never guess.
