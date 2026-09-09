# AGENTS.md

## Purpose

Bootstrap/router for MCS.OSJS agents. Keep this file small. Detailed behavior belongs in the directive and workflow files.

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

## Authoring Failure Circuit Breaker

When an edit or generated source is malformed:

1. Read the actual affected file and exact formatter/compiler error.
2. Fix only the smallest affected source region.
3. Run the smallest authoritative formatter/compiler/test.
4. One corrective retry is allowed.

If the second authoring attempt is still malformed: **STOP**.

Do not continue implementation, create repair-script chains, perform broad regex cleanup, repeatedly switch authoring transports, or investigate encoding/serialization/shell/editor/tool corruption.

Tool or transport corruption may be claimed only when a minimal reproducible probe independent of the affected implementation file demonstrates the same mutation. Otherwise treat malformed output as an authoring failure.

After STOP, preserve repository state and report the exact failing file, observed error, attempted corrections, and verification result.

## No Broad Source Repair Rule

Never repair malformed authored source with broad substitutions such as punctuation collapsing, Unicode stripping, global regex normalization, or repository-wide cleanup unless the authorized task explicitly requires that transformation and every changed occurrence is independently verified.

Fix exact malformed statements instead.

## Execution Guardrail

- Inspect actual written file/output before forming a theory.
- Never invent transport, parser, shell, editor, encoding, or tool-corruption explanations without reproducible independent evidence.
- Do not build chains of encoding, marker, escaping, base64, heredoc, repair-script, or self-modifying workarounds for a simple edit.
- Verification claims must name the actual command/check observed and its result. Never upgrade a substitute check into an authoritative gate by inference.

## Verification Truth Rule

A verification result proves only what that exact check establishes.

- Repository-native or task-defined verification outranks generic substitute checks.
- Never call `node --check`, lint, unit tests, raw-byte inspection, a standalone parser, or another approximation equivalent to a package/build/runtime gate unless repository authority explicitly defines it that way.
- A higher-fidelity failed gate invalidates any earlier claim that the affected surface was verified.
- If required verification cannot run in the current environment, report it as unavailable. Do not convert `cannot verify` into `verified enough`.

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
