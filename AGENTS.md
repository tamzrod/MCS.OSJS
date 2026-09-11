# AGENTS.md

## Purpose

Bootstrap/router for MCS.OSJS agents. Keep this file small. Detailed behavior belongs in directive, workflow, and tool files.

## StarCraft Directives

Resolve these commands before any generic repository workflow:

- `BLACK SHEEP WALL` → `BLACK_SHEEP_WALL.md`
- `OPERATION CWAL` → `operation cwal.md`
- `THE GATHERING` → `the gathering.md`
- `THERE IS NO COW LEVEL` → `there is no cow level.md`

When invoked:

```text
IDENTIFY DIRECTIVE
→ READ ITS FILE
→ FOLLOW ITS OWN SCOPE / READ / WRITE / STOP RULES
→ NEVER GUESS
```

Do not import behavior from another directive unless the active directive explicitly calls it.

## Generic Repository Work

For repository-dependent work that is not a StarCraft directive:

```text
READ ICC/INDEX.md FIRST
→ USE CURRENT RELEVANT ICC CONTEXT
→ IF STALE OR MISSING, REFRESH ONLY THE AFFECTED CONTEXT THROUGH BLACK_SHEEP_WALL.md
→ LOCATE AUTHORITATIVE REPOSITORY SOURCE
→ ACT WITHIN USER-AUTHORIZED SCOPE
→ VERIFY
```

## Editing

Use the normal repository editing tools appropriate to the task. Keep edits as small and direct as practical.

If an authored edit is malformed:

1. Reload or restore the last known-good affected source.
2. Retry the smallest affected edit once.
3. Run the repository-native verification for that edit.
4. If the retry is also malformed, **STOP and report the failure.**

Do not diagnose the cause during the task. Do not test transports, encodings, shells, editors, or byte paths. Do not build repair scripts or switch editing mechanisms to rescue malformed generated text. Do not repeatedly repair malformed output.

A normal formatter/compiler error is not automatically malformed authoring. Follow the authoritative error location, make one local correction, and rerun the smallest required gate.

## Execution Guardrail

- Inspect the actual result of an edit before continuing.
- Verification claims must name the actual command/check observed and its result.
- Do not build workaround chains for a simple edit.
- If required verification cannot run, report it as unavailable rather than verified.

## Verification Truth Rule

A verification result proves only what that exact check establishes.

- Repository-native or task-defined verification outranks generic substitutes.
- Never call a generic syntax check, lint, unit test, raw-byte inspection, or standalone parser equivalent to a package/build/runtime gate unless repository authority explicitly defines it that way.
- A higher-fidelity failed gate invalidates any earlier claim that the affected surface was verified.

## Completion Integrity Rule

No directive may mark implementation complete, archive it, advance a successor, or describe it as verified unless that directive's required completion gate has actually passed.

A checkpoint may preserve incomplete or failing work only through a directive that explicitly permits incomplete checkpoints.

## Directive Composition Rule

A StarCraft directive ends at its own RETURN / STOP boundary.

Do not invoke another StarCraft directive merely because the first directive finished. Cross-directive invocation is allowed only when:

- the user explicitly invokes the other directive; or
- the active directive explicitly delegates to it.

Current delegation:

- `OPERATION CWAL` → `BLACK SHEEP WALL` only for bounded stale/missing ICC context.
- `THERE IS NO COW LEVEL` → `BLACK SHEEP WALL` only when rescue genuinely requires ICC maintenance.
- `THE GATHERING` → none.
- `BLACK SHEEP WALL` → return to caller.

## Authority

- Repository files are authoritative.
- `workflow/active_work/` is implementation authority.
- Planning and Brainstorm do not authorize implementation.
- `handoff.md` is execution continuation state.
- `ICC/` is context, not authority.

Do not store project history, troubleshooting notes, runbooks, task-specific discoveries, or duplicated directive rules in this file.
