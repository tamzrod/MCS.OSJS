# THERE IS NO COW LEVEL

## Purpose

`THERE IS NO COW LEVEL` is the emergency sandbox-rescue command for MCS.OSJS.

Use it when the current JR/agent session or sandbox has become unreliable, confused, corrupted, stuck in a workaround loop, or otherwise unsafe to continue.

Its job is not to finish the current task. Its job is to preserve valid work, create a trustworthy continuation point on `main`, and make the current sandbox disposable.

> Rescue what is valid. Record exactly where work stopped. Push to main. Stop.

## Authority

`THERE IS NO COW LEVEL` is explicitly authorized to:

- inspect the current working tree and current `HEAD`;
- read the current Active Work authority and `handoff.md`;
- inspect changed files needed to classify current sandbox work;
- keep valid in-scope changes;
- repair only minimal mechanical damage required to make rescued work coherent and verifiable;
- discard clearly accidental, malformed, temporary, or failed-experiment artifacts created by the broken session;
- update `handoff.md` with the exact continuation state;
- commit the rescued repository state;
- push the rescue commit to `main`.

It does not gain authority to invent new product behavior, widen implementation scope, promote planning, or complete unrelated work.

## Rescue Principle

A rescue is a state-transfer operation, not a development sprint.

```text
BROKEN SANDBOX
→ IDENTIFY TRUSTWORTHY STATE
→ SALVAGE VALID WORK
→ VERIFY SALVAGED STATE
→ WRITE CONTINUATION STATE
→ COMMIT
→ PUSH MAIN
→ STOP
→ OLD SANDBOX MAY BE DESTROYED
```

When uncertain whether a change is intentional and valid, preserve evidence in `handoff.md` and leave the questionable implementation out of the rescue commit rather than guessing.

## Read Scope

Read only what is required to rescue the current work:

1. `AGENTS.md`
2. `ICC/INDEX.md` and the smallest relevant ICC context
3. `handoff.md`
4. current file under `workflow/active_work/`
5. current `git status`, staged/unstaged diff, and commits since the last known good state
6. changed source/test/config files directly involved in the current work

Do not perform a repository-wide audit.

If ICC itself needs maintenance, invoke `BLACK SHEEP WALL` only for the affected semantic branch. THERE IS NO COW LEVEL must not edit `ICC/` directly.

## Rescue Classification

Classify every current sandbox change into one of these buckets:

### KEEP

Change is clearly part of authorized Active Work, coherent, and supported by repository state.

### REPAIR

Change is clearly intended and in scope but has small mechanical damage such as formatting, syntax, incomplete rename, or an obviously interrupted edit. Repair only enough to restore coherent rescued state.

### DROP

Change is clearly accidental or disposable, including:

- temporary probes;
- generated scratch files;
- failed workaround artifacts;
- malformed duplicate files;
- debugging debris;
- changes outside authorized scope introduced by the broken session.

### UNCERTAIN

Intent or correctness cannot be established cheaply and safely.

Do not guess. Exclude it from the rescue commit when exclusion is safe, and record it in `handoff.md` as unresolved evidence for the next sandbox.

## Anti-Rabbit-Hole Rule

The global execution guardrail in `AGENTS.md` remains mandatory during rescue.

Additionally:

- do not investigate why JR, OpenHands, the shell, or an editing tool behaved badly unless that diagnosis is required to preserve repository state;
- do not continue failed workaround chains;
- do not redesign the current implementation during rescue;
- do not create speculative replacement code merely to make tests green;
- do not spend rescue time proving a tool bug.

The goal is a trustworthy checkpoint, not an explanation of the failed sandbox.

## Verification

Verify the smallest meaningful surface for the rescued changes.

Preferred order:

```text
FORMAT / STATIC CHECK IF APPLICABLE
→ TARGETED TESTS FOR CHANGED AREA
→ BROADER TEST ONLY WHEN CHEAP AND REQUIRED TO ESTABLISH SAFETY
```

Record failures honestly. A rescue may still be committed when unfinished Active Work is intentionally preserved, provided `handoff.md` clearly states what is incomplete or failing and the committed state is not misleading.

Do not label failing or unverified work as complete.

## Handoff Update

Before committing, rewrite `handoff.md` so a fresh sandbox can continue without reconstructing the broken session.

It must establish, at minimum:

- current Active Work item;
- what was successfully rescued;
- exact implementation point reached;
- verification performed and result;
- unresolved or excluded changes that matter;
- exact next action for the new sandbox;
- any command/test needed to reproduce the current state;
- explicit statement that the previous sandbox was rescued and must not be treated as authoritative beyond the rescue commit.

Keep `handoff.md` as continuation state, not a transcript or postmortem.

## Commit and Push Rule

After classification, minimal repair, verification, and handoff update:

```text
REVIEW FINAL DIFF
→ ENSURE ONLY RESCUED STATE IS INCLUDED
→ COMMIT ALL INTENDED RESCUE CHANGES
→ PUSH TO main
→ CONFIRM REMOTE main CONTAINS THE RESCUE COMMIT
```

Do not stop at a local commit.

Do not leave intended rescue changes uncommitted after reporting success.

If push to `main` fails, do not claim rescue completion. Report the exact failure and preserve the local commit SHA so the user can recover it.

## Sandbox Destruction Boundary

`THERE IS NO COW LEVEL` does not itself destroy an OpenHands/cloud sandbox unless the execution environment explicitly provides a safe sandbox-destruction capability.

After the rescue commit is confirmed on remote `main`, output:

```text
RESCUE COMPLETE
main: <rescue commit SHA>
old sandbox: SAFE TO DESTROY
new sandbox: CONTINUE FROM handoff.md
```

Then stop all implementation work in the old sandbox.

The user may destroy the old sandbox. A fresh sandbox should pull `main`, read `AGENTS.md`, then follow `handoff.md` and the current Active Work authority.

## Failure Handling

If rescue cannot safely reach `main`:

```text
STOP MUTATING
→ PRESERVE BEST KNOWN LOCAL STATE
→ REPORT git status
→ REPORT LOCAL COMMIT SHA IF ONE EXISTS
→ REPORT EXACT BLOCKER
→ DO NOT CLAIM SAFE TO DESTROY
```

Never tell the user to destroy the sandbox until the rescue state is confirmed on remote `main`.

## THERE IS NO COW LEVEL Command

```text
THERE IS NO COW LEVEL
→ FREEZE FEATURE DEVELOPMENT
→ READ AGENTS + CURRENT ICC ROUTE + ACTIVE WORK + HANDOFF
→ INSPECT HEAD + STATUS + CURRENT DIFF
→ CLASSIFY CHANGES: KEEP / REPAIR / DROP / UNCERTAIN
→ SALVAGE ONLY AUTHORIZED TRUSTWORTHY WORK
→ RUN MINIMUM MEANINGFUL VERIFICATION
→ UPDATE handoff.md WITH EXACT CONTINUATION STATE
→ REVIEW FINAL DIFF
→ COMMIT RESCUE STATE
→ PUSH TO main
→ CONFIRM REMOTE main HAS RESCUE COMMIT
→ REPORT SAFE TO DESTROY + NEW-SANDBOX CONTINUATION
→ STOP
```

## Core Invariant

```text
THERE IS NO COW LEVEL
= SALVAGE
+ HANDOFF
+ COMMIT
+ PUSH MAIN
+ STOP

THERE IS NO COW LEVEL
!= FINISH FEATURE
!= REPOSITORY AUDIT
!= TOOL-FAILURE INVESTIGATION
!= SPECULATIVE REWRITE
!= NEW SCOPE
!= CLAIM SUCCESS BEFORE REMOTE CONFIRMATION
```

Never guess.
