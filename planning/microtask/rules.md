# Microtask Rules

## ICC-First Context Rule

Before repository-dependent task planning, select the task's semantic boundary and follow `workflow/ICC_CONSUMER.md`: read `ICC/INDEX.md`, the relevant node in `ICC/manifest.json`, and the minimum context. Validate per-node baseline/overlay/validity, not an index-wide snapshot. REUSE current unchanged context; for stale, unverified or missing territory request only an authorized bounded BLACK SHEEP WALL UPDATE/REVEAL, or use task-permitted authoritative source with an explicit caveat/report the precise blocker. Only BLACK SHEEP WALL edits ICC; Git source remains authoritative. A migrated unverified node is not current, but does not justify a global audit or automatically block task sizing.

## Core Rule

One task = one primary outcome. One task = one detailed task file. Planning does not authorize execution; human promotion moves a task into `workflow/active_work/` and synchronizes `handoff.md`.

## Execution roles and separate gates (2026-09-18 human decision)

For MCS Modbus Toolkit UMIG work, ChatGPT is the coding agent; OpenHands is the independent JR test runner. Do not imply OpenHands was called or executed unless its actual report exists. Keep three distinct task files when a feature needs all three stages:

- **CODE** (`UMIG-NNN`, existing stable ID): only source implementation and a handoff-ready source checkpoint. Owner: ChatGPT. Record exact files and commit/diff; read back changes. No build, unit-test, runtime or acceptance PASS is claimed by source inspection.
- **TEST** (`UMIG-NNN-T`): deterministic build, package discovery, unit/fixture or static verification of the committed coding output. Owner: OpenHands acting as JR. No product source fixes, feature implementation or workflow advancement. Return raw commands, exit codes and results.
- **VERIFY** (`UMIG-NNN-V`): independent rendered UI, live runtime, installed behavior, data-safety or deployment acceptance, in a specified safe target. Owner: OpenHands acting as JR. No product fixes or speculative substitute evidence. Return actual observations and PASS/FAIL/BLOCKED.

If a task is inherently human approval (`UMIG-001`) or solely a verification gate (`UMIG-007`, `UMIG-007A`), keep that identity; do not manufacture coding work. Test and verify may be a single task only when there genuinely is one indivisible check, but never combine coding with either. Preserve explicit `Previous`/`Next` across stages; never auto-promote a Planning successor.

**Handoff:** Once CODE source is committed and inspected, the coding agent records its source-only outcome, archives/advances through authorized Active Work, and writes one explicit current `JR TEST TASK` in `handoff.md` for the active OpenHands test stage. The packet contains GOAL, EXACT COMMAND/ACTION, EXPECTED RESULT, EVIDENCE, safe setup and report-write authority. OpenHands follows `operation cwal.md`: runs only that packet, changes no product files, reports evidence, and stops. The coding agent reviews the report and alone decides verified completion and advancement. On FAIL/BLOCKED: do not mark passed or auto-advance; create/authorize a bounded coding repair when needed, then retest. Source review is never a replacement for a test, and automated tests never substitute for required actual UI/runtime verification.

## Preferred Size

Score five dimensions, each 0, 1 or 2: implementation surface, environment/dependency uncertainty, behavioral surface, verification surface and decision/recovery surface. Total 0-3 is preferred; 4-5 split unless tightly coupled with one deterministic workflow; 6-7 split; 8-10 must split. Never artificially lower a score because multiple phases belong to one feature.

## Mandatory Split Triggers

Split for more than 3 independent acceptance outcomes, more than 3 independent implementation verbs, multiple verification workflows, multiple architecture decisions, independently verifiable sequential subtasks, donor/import/toolchain plus live behavior, unknown environment discovery plus product implementation, or an early failure forcing later steps to be reinterpreted. Coding, TEST and VERIFY stages are distinct by default, even for a small feature.

For imported components: establish/code first; OpenHands build/static TEST next; OpenHands runtime/UI VERIFY last. If an integration needs separate unit and live probes, keep separate tests and verification tasks. A task must stay within one bounded semantic branch and working set; split when in doubt.

## Required Task Shape

Every task has one ID/title, stage and owner, one primary outcome, scope, non-scope, no more than three independently testable acceptance outcomes, a stage-appropriate evidence/handoff requirement, dependencies, Previous/Next and five-dimension sizing. TEST/VERIFY tasks must spell out exact commands or repeatable actions, expected observations and evidence. Do not allow a tester to choose its own scope from neighboring tasks. Record actual failures and environment blockers separately.

## Promotion Boundary

A human approves promotion into `workflow/active_work/`. An explicitly human-authorized ordered sequence may move between its already QUEUED tasks after evidence; a successor still in Planning needs separate promotion. Maintain exactly one ACTIVE. Only the coding agent manages implementation/task state; OpenHands/JR may update only the explicitly authorized handoff report. Only BLACK SHEEP WALL updates ICC.
