# Governance and Operation Discipline

Baseline commit: ee19b8a
Working tree: clean
Source dependencies: AGENTS.md, BLACK_SHEEP_WALL.md, ICC/INDEX.md, operation cwal.md, handoff.md, workflow/active_work/README.md, "the gathering.md", "there is no cow level.md"
Parent: L0-project
Zoom In: (none; leaf node); indexing/active-work program via L0 half; implementation details via active-work node
Zoom Out: L0-project

## Directive Router (AGENTS.md)

`AGENTS.md` is now a small bootstrap/router, not the long operation chain it used to state. It maps
four StarCraft commands before any generic repository workflow:

- `BLACK SHEEP WALL` → `BLACK_SHEEP_WALL.md`
- `OPERATION CWAL` → `operation cwal.md`
- `THE GATHERING` → `the gathering.md`
- `THERE IS NO COW LEVEL` → `there is no cow level.md`

Rule: identify the directive, read its file, follow its own scope/read/write/stop rules, never guess,
and do not import behavior from another directive unless the active directive explicitly calls it.
Delegation recorded in `AGENTS.md`: `OPERATION CWAL` → `BLACK SHEEP WALL` only for bounded stale or
missing ICC context; `THERE IS NO COW LEVEL` → `BLACK SHEEP WALL` only when rescue genuinely requires
ICC maintenance; `THE GATHERING` → none; `BLACK SHEEP WALL` → return to caller.

## Operation Chain (as rooted in AGENTS.md)

For repository-dependent work that is not a StarCraft directive:
```text
READ ICC/INDEX.md FIRST
→ USE CURRENT RELEVANT ICC CONTEXT
→ IF STALE OR MISSING, REFRESH ONLY THE AFFECTED CONTEXT THROUGH BLACK_SHEEP_WALL.md
→ LOCATE AUTHORITATIVE REPOSITORY SOURCE
→ ACT WITHIN USER-AUTHORIZED SCOPE
→ VERIFY
```

## ICC State Model

Repository truth at BASELINE COMMIT + audited uncommitted working-tree overlay. Baseline commit is
the HEAD used at audit time (need not equal the commit that stores ICC files). For uncommitted files,
record deterministic fingerprints per audited path. A context is synchronized when: committed deps are
represented by the current baseline; no changed source dep is unincorporated; and uncommitted deps
match the audited overlay fingerprints. Known-stale context is never consumed as authoritative.
Unaffected synchronized context is not recomputed.

## Editing and Execution Guardrails (AGENTS.md)

- Use normal repository editing tools; keep edits small and direct.
- If an authored edit is malformed: reload the last known-good source, retry the smallest edit once,
  run the repository-native verification for that edit, and if the retry is also malformed, STOP and
  report. Do not diagnose the cause, test transports/encodings/shells/editors/byte paths, build
  repair scripts, switch editing mechanisms, or repeatedly repair malformed output.
- A normal formatter/compiler error is not malformed authoring: follow the error location, make one
  local correction, rerun the smallest required gate.
- Inspect the actual result of an edit before continuing. Verification claims must name the actual
  command/check observed and its result. If required verification cannot run, report it as
  unavailable rather than verified.

## Verification Truth Rule (AGENTS.md)

A verification result proves only what that exact check establishes. Repository-native or task-defined
verification outranks generic substitutes. Never call a generic syntax check, lint, unit test,
raw-byte inspection, or standalone parser equivalent to a package/build/runtime gate unless repository
authority explicitly defines it that way. A higher-fidelity failed gate invalidates any earlier claim
that the affected surface was verified.

## Completion Integrity Rule (AGENTS.md)

No directive may mark implementation complete, archive it, advance a successor, or describe it as
verified unless that directive's required completion gate actually passed. A checkpoint may preserve
incomplete or failing work only through a directive that explicitly permits incomplete checkpoints.

## Authority (AGENTS.md)

Repository files are authoritative. `workflow/active_work/` is implementation authority. Planning and
Brainstorm do not authorize implementation. `handoff.md` is execution continuation state. `ICC/` is
context, not authority. Do not store project history, troubleshooting notes, runbooks, task-specific
discoveries, or duplicated directive rules in `AGENTS.md`.

## BLACK SHEEP WALL (modes

- Direct invocation = repository-wide context compaction:  current HEAD -> audit all relevant context -> discover semantic boundaries -> compact into ICC/context/ -> register in index -> stamp baseline -> record audited uncommitted overlays. Subsequent runs:  do NOT rescan unchanged files; inspect only committed files changed since baseline + uncommitted files differing from audited overlay; refresh only affected context; preserve unaffected. If HEAD unchanged, only new/modified/renamed/deleted uncommitted files differing from audited overlay are inspected.
- Delegated branch refresh = invoked by CWAL/planning/etc with a selected semantic node as navigation ceiling; refresh only that node/required children; never repair stale siblings or widen scope.

`BLACK_SHEEP_WALL.md` was rewritten by this delta and is now explicit that BLACK SHEEP WALL keeps ICC
as a compact semantic cache, owns every ICC write, patches the deepest affected node first, propagates
only demonstrated semantic impact, keeps `ICC/INDEX.md` stable unless routing/structure/state changes,
never advances baseline/overlay ahead of verified re-read node writes, and returns to its caller at
its own STOP boundary.

## Operation CWAL

`operation cwal.md` was rewritten by this delta. JR is a test runner only. The current `JR TEST TASK`
in `handoff.md` is the complete execution authority; if it is absent or materially incomplete, JR
reports BLOCKED and stops. JR must not edit source, create/delete implementation files, refactor,
apply suspected fixes, alter configuration to force a pass, update `workflow/active_work/` or `ICC/`,
invoke BLACK SHEEP WALL, promote planning, archive/advance tasks, merge, or reset/clean/restore
repository state. Sandbox test-environment preparation is allowed and is not coding; use BLOCKED for a
missing dependency only when safe sandbox-local preparation is unavailable or would require
prohibited repository/product modification or unavailable privilege. JR may modify `handoff.md` only
when the packet explicitly authorizes it, then preserve every other section, report evidence only, and
commit/push only `handoff.md` if instructed.

## Workflow Boundaries (per AGENTS

- Brainstorm/Planning:  ICC first, refresh affected only through BSW; human-owned, no implementation authority
- Microtask:  ICC first, refresh stale affected; follow planning/microtask/rules.md
- Promotion:  ICC first, refresh stale; promote only human-selected microtask; sync handoff.md
- CWAL:  ICC first for Active Work context; executes only authorized Active Work.

`workflow/active_work/README.md` now states the ordered-sequence protocol: a human may promote an
ordered sequence whose first task is ACTIVE and later tasks QUEUED; after the ACTIVE task meets its own
stage gate the coding agent archives it, follows its explicit `Next`, checks the successor exists as
QUEUED with a matching `Previous`, activates only that successor, synchronizes `handoff.md`, and
commits/pushes that state together. If `Next` is absent, the successor is missing/in Planning/not
QUEUED/mismatched, or `handoff.md` conflicts, stop for human repair; never infer. OpenHands/JR never
advances.

## Key Operational Rules

- The ICC index is a registry, not a navigation permission menu. Registry presence is not relevance; crossing semantic branches requires the current authorized task to explicitly require that boundary for implementation/understanding/testing/verification.
- Zoom Out must not climb above the task-selected navigation ceiling and re-enter an unrelated sibling branch.
- An operation must not reopen unchanged repository source files when synchronized ICC contains the needed context.
- BLACK SHEEP WALL cannot grant execution authority, promote work, choose future work, let CWAL inspect Planning to select tasks, let brainstorming implement code, widen a bounded operation, or convert unrelated ICC staleness into work for the invoking operation.
