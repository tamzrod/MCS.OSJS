# AGENTS.md

## Purpose

Bootstrap/router for MCS.OSJS agents. Keep this file small. Detailed behavior belongs in the directive and workflow files.

## StarCraft Directives

Resolve these commands before any generic repository workflow:

- `BLACK SHEEP WALL` → `BLACK_SHEEP_WALL.md`
- `OPERATION CWAL` → `operation cwal.md`
- `THE GATHERING` → `the gathering.md`
- `STAYING ALIVE` → `staying alive.md`

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

## Execution Guardrail

When an edit, generated file, or command fails because of malformed output or syntax:

- Inspect the actual written file/output before forming a theory.
- Retry the same authoring approach at most once.
- After two failures, stop workaround experimentation and switch to a deterministic edit/patch method.
- Do not invent transport, parser, shell, or tool-corruption explanations without reproducible evidence.
- Do not build chains of encoding, marker, escaping, or self-modifying workarounds for a simple edit.
- After the deterministic edit, run the smallest relevant formatter/compiler/test.
- If it still fails, stop and report the exact observed failure and repository state; do not continue an open-ended debugging loop.

## Authority

- Repository files are authoritative.
- `workflow/active_work/` is implementation authority.
- Planning and Brainstorm do not authorize implementation.
- `handoff.md` is execution continuation state.
- `ICC/` is context, not authority.

Do not store project history, troubleshooting notes, runbooks, task-specific discoveries, or duplicated directive rules in this file.
