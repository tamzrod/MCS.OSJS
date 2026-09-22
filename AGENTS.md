# AGENTS.md

## Identity bootstrap

First question: **Who are you?** Establish actual agent/runtime identity, then read `workflow/IDENTITY_MAP.md`. For OpenCode work, load `workflow/adapters/opencode.md`. Unknown identity must ask operator and STOP before mutation. Identity does not authorize a task. Platform adapters never override the shared directive or per-task permissions.

## Purpose

Bootstrap/router for MCS.OSJS agents. Keep this file small. Detailed behavior belongs in directive, workflow and tool files.

## StarCraft Directives

Resolve these commands before generic repository workflow:

- `BLACK SHEEP WALL` → `BLACK_SHEEP_WALL.md`
- `OPERATION CWAL` → `operation cwal.md`
- `THE GATHERING` → `the gathering.md`
- `THERE IS NO COW LEVEL` → `there is no cow level.md`

When invoked:

```text
IDENTITY → IDENTITY MAP → IDENTIFY DIRECTIVE
→ READ ITS FILE
→ FOLLOW ITS OWN SCOPE / READ / WRITE / STOP RULES
→ NEVER GUESS
```

Do not import behavior from another directive unless the active directive explicitly calls it. `OPERATION CWAL` processes a human-approved queue continuously, one eligible task at a time, respecting dependencies and per-packet roles. OpenCode CODE/DISCOVERY/DESIGN and independent JR TEST/VERIFY are distinct identities; invocation never grants JR identity or permission to claim independent PASS.

## Generic Repository Work

For repository-dependent work outside a StarCraft directive:

```text
READ ICC/INDEX.md FIRST
→ USE CURRENT RELEVANT ICC CONTEXT
→ IF STALE OR MISSING, REFRESH ONLY AFFECTED CONTEXT THROUGH BLACK_SHEEP_WALL.md
→ LOCATE AUTHORITATIVE REPOSITORY SOURCE
→ ACT WITHIN USER-AUTHORIZED SCOPE
→ VERIFY
```

## Editing

Use normal repository editing tools appropriate to the task. Keep edits small and direct. If an authored edit is malformed: reload/restore the last known-good affected source; retry the smallest edit once; run repository-native verification; if malformed again STOP and report. Do not diagnose transport, encoding, shell or editor during the task or build repair scripts. A normal formatter/compiler error is not automatically malformed authoring; follow authoritative error location, make one local correction and rerun the smallest required gate.

## Execution guardrail and verification truth

Inspect actual edit result before continuing. Verification claims name the exact observed check and result. Repository-native or task-defined verification outranks generic substitutes. Never equate a generic syntax check, lint, unit test or standalone parser with a package/build/runtime gate unless repository authority explicitly defines it. A higher-fidelity failed gate invalidates earlier verification claims for that surface.

## Completion integrity

No directive may mark implementation COMPLETE or verified, archive it or release a dependent task without its required completion evidence and delivery. A checkpoint may preserve incomplete work only where explicitly allowed. CWAL may continue to an independent eligible task after a task FAIL/BLOCKED, but must not bypass a failed prerequisite.

## Directive composition

Do not invoke another StarCraft directive merely because the first finished. Cross-directive invocation requires explicit user invocation or active directive delegation. Current delegation: OPERATION CWAL → BLACK SHEEP WALL only for bounded stale/missing ICC context; THERE IS NO COW LEVEL → BLACK SHEEP WALL only for genuinely needed ICC maintenance; THE GATHERING → none; BLACK SHEEP WALL → return to caller. A directive ends at its own STOP boundary; CWAL's queue loop and its termination condition are defined in `operation cwal.md`.

## Authority

Repository files are authoritative. `workflow/active_work/` holds promoted work; `handoff.md` and approved queue hold continuation state. Planning and Brainstorm do not authorize execution; ICC is context, not execution authority. Do not store project history, troubleshooting notes, runbooks or duplicated detailed directive rules here.
