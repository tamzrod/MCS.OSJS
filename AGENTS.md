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

## Controlled Editing Path

For substantive source/config/document edits, especially multiline replacements or generated blocks, use `tools/experimental-editor/editor.py` as the preferred repository write path. Read and follow `tools/experimental-editor/README.md` before first use in a run.

Normal shell/git operations, file reads, directory creation, deletion, and trivial exact edits do not need to be routed through it. Do not switch among heredocs, `str_replace`, generator scripts, ad-hoc Python rewrites, or alternate editors as a recovery strategy after malformed authoring appears.

A successful `SAFEEDIT_OK` proves only that the editor wrote exactly the bytes it received. It does not prove those authored bytes were semantically correct.

## Authoring Failure Circuit Breaker

A normal localized syntax or compile error is not authoring corruption. Use the authoritative formatter/compiler location, make one exact local correction, and rerun the smallest authoritative gate. Do not manually recount large delimiter nests when the parser/compiler can identify the failure.

Treat repeated unusual malformed authoring as patterned corruption when a corrective or fresh write produces new unexplained mutations of the same family, such as duplicated punctuation, joined/split words, delimiter substitution, mangled operators, or unrelated syntax mutation.

**Pre-execution malformed authoring counts.** If such mutations are already present in text supplied to an editor, shell, API, or other tool before execution, they qualify as authored corruption. A byte-identical transport/editor probe can rule out mutation in that tested transport/write path, but it cannot downgrade repeated malformed authored input into an ordinary typo.

Single syntax mistakes, one brace mismatch, the same still-unfixed error, and cascading parser errors do not qualify.

### Fresh Authoring Recovery

After patterned corruption appears following one exact corrective attempt:

1. Stop authoring that approach immediately.
2. Preserve/restore the last known-good repository state and abandon malformed temporary artifacts.
3. Stop using the malformed block as implementation input.
4. Reload authoritative known-good source.
5. Redo only the smallest coherent affected transformation through the controlled editing path.
6. Run the smallest authoritative formatter/compiler/test immediately.

If the fresh attempt is normal, continue. If the same unusual corruption appears again in the fresh attempt, **STOP immediately; do not repair it again or switch editing mechanisms.**

Tool or transport corruption may be claimed only when an independent minimal reproducible probe demonstrates mutation in that layer. Otherwise report observed authoring corruption without inventing a lower-level cause.

## No Broad Source Repair Rule

Never repair malformed authored source with punctuation collapsing, Unicode stripping, global regex normalization, repository-wide cleanup, or similar broad substitutions unless the authorized task explicitly requires that exact transformation and every changed occurrence is independently verified.

## Execution Guardrail

- Inspect actual written output before forming a theory.
- Verification claims must name the actual command/check observed and its result.
- Do not build encoding, escaping, base64, heredoc, transport-switch, repair-script, or self-modifying workaround chains for a simple edit.
- Do not infer transport, parser, shell, editor, encoding, or tool corruption without reproducible evidence specific to that layer.

## Verification Truth Rule

A verification result proves only what that exact check establishes.

- Repository-native or task-defined verification outranks generic substitutes.
- Never call `node --check`, lint, unit tests, raw-byte inspection, a standalone parser, or another approximation equivalent to a package/build/runtime gate unless repository authority explicitly defines it that way.
- A higher-fidelity failed gate invalidates any earlier claim that the affected surface was verified.
- If required verification cannot run, report it as unavailable rather than verified.

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
