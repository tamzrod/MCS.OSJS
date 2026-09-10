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

Treat unusual malformed authoring as patterned corruption when the same unusual family appears again after one exact local correction or in a fresh authored attempt. Examples include duplicated punctuation, joined/split words, delimiter substitution, mangled operators, foreign punctuation, or unrelated syntax mutation.

**Pre-execution malformed authoring counts.** If such mutations are already visible in text supplied to an editor, shell, API, or other tool before execution, they count as authored corruption.

Single syntax mistakes, one brace mismatch, the same still-unfixed error, and cascading parser errors do not qualify.

### Mandatory state machine

Use this sequence exactly:

```text
FIRST unusual malformed authored payload
→ ONE exact local correction only

SAME unusual family appears again
→ ENTER FRESH AUTHORING RECOVERY IMMEDIATELY

FRESH AUTHORING RECOVERY
→ preserve/restore last known-good affected file
→ abandon malformed temporary/generated artifacts
→ stop referencing malformed authored text
→ reload authoritative known-good source
→ perform ONE smallest coherent fresh edit through the controlled editing path
→ run the smallest authoritative formatter/compiler/test

IF fresh attempt is normal
→ continue task

IF same unusual family appears during fresh attempt
→ HARD STOP
```

Once Fresh Authoring Recovery is triggered, **do not run authoring sanity probes, transport probes, byte probes, encoding investigations, shell/editor comparisons, or alternate-tool experiments before recovery.** Recovery is the next action.

A clean transport, shell, or editor probe **never authorizes continued authoring after patterned authoring corruption has already been established**. Such a probe can only say that the tested layer preserved the bytes it received; it cannot reclassify malformed authored input as an ordinary typo.

During or after Fresh Authoring Recovery, do not use `sed`, `perl`, ad-hoc Python rewrites, generator scripts, heredoc reconstruction, `str_replace`, token-by-token punctuation repair, broad substitutions, or editor switching to rescue the malformed artifact.

If the fresh attempt shows the same unusual corruption family, **STOP immediately. Do not correct it. Do not probe it. Do not try another transport. Do not switch editors. Preserve repository state and report the exact failing file, observed mutation pattern, first corrective attempt, fresh recovery attempt, and verification result.**

Tool or transport corruption may be claimed only when an independent minimal reproducible probe demonstrates mutation in that layer, and only when such diagnosis is actually needed. Do not delay or bypass the authoring circuit breaker to perform that diagnosis.

## No Broad Source Repair Rule

Never repair malformed authored source with punctuation collapsing, Unicode stripping, global regex normalization, repository-wide cleanup, or similar broad substitutions unless the authorized task explicitly requires that exact transformation and every changed occurrence is independently verified.

## Execution Guardrail

- Inspect actual written output before forming a theory.
- Verification claims must name the actual command/check observed and its result.
- Do not build encoding, escaping, base64, heredoc, transport-switch, repair-script, or self-modifying workaround chains for a simple edit.
- Do not infer transport, parser, shell, editor, encoding, or tool corruption without reproducible evidence specific to that layer.
- Once the authoring circuit breaker reaches Fresh Authoring Recovery, diagnostic curiosity does not override the state machine.

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
