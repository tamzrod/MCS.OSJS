# Microtask Rules

## Context and authority
Before repository-dependent planning, read `ICC/INDEX.md` and relevant valid context; verify against actual Git source. Stale ICC requires a separately authorized bounded BLACK SHEEP WALL refresh; only BLACK SHEEP WALL edits ICC. Planning is not execution authority. Human approval is required to promote a task into a fresh executable queue. Historical OTR packets and retired queues remain non-executable.

## Core rule
One microtask = one primary outcome = one detailed task file. Keep tasks small, independently verifiable, and reuse source-verified existing work rather than recreating it. Never infer completion from an old report or status alone.

## Required task format
Every new task packet must contain these exact, prominent fields:

Task ID: <unique stable ID>
Task Name: <one specific outcome>
Blocker Task: NONE | <comma-separated exact Task IDs>
Status: PENDING | READY | ACTIVE | COMPLETE | FAIL | BLOCKED
Assigned Agent: <explicit agent and role>
Stage: DISCOVERY | DESIGN | CODE | TEST | VERIFY

Then include:
- Objective: one measurable result.
- Scope: exact read paths, exact permitted write paths, and forbidden/out-of-scope actions.
- Execution: bounded steps; exact commands/actions for TEST/VERIFY.
- Acceptance Criteria: at most three independently testable outcomes and explicit expected results.
- Evidence: exact report path, required raw observations, source/commit reference and handoff requirements.
- Completion and delivery: verdict rules, permitted status edits, exact commit/push/ref and remote verification when required; otherwise explicitly state no commit/push authority.
- Sizing: five-dimension score and rationale.

Do not use Previous/Next as routing or dependency fields. Task ID is an identifier, not an execution order. Every Blocker Task ID must resolve to a real task; no circular dependencies, self-dependencies or references to historical/retired tasks as executable blockers. When previously completed work is reused, cite source and verified evidence as an input, not an invented COMPLETE task.

## Dependency-based readiness and selection
PENDING means not executed; waiting for blockers is PENDING, not BLOCKED. READY means all listed blocker tasks are COMPLETE with their required evidence independently verified, and the packet and other safety/authority gates are satisfied. `Blocker Task: NONE` removes task dependencies only; it does not waive approval, packet validity, environment or safety gates.

When a human authorizes selection from a newly approved queue, inspect every PENDING task, resolve its blocker IDs and evidence, and identify all eligible READY tasks. Select exactly ONE by the lowest stable Task ID as a deterministic tie-breaker, not as a sequential dependency. Explicitly activate only that task in root `handoff.md`, its packet and the queue before execution; do not execute a merely READY task without activation. If no task is eligible, report why and STOP. If the handoff is NONE or the queue is retired, report NO ACTIVE TASK and STOP; selection does not override the reset.

Operation CWAL executes only the explicitly handoff-named ACTIVE task and then STOPS after COMPLETE, FAIL or BLOCKED. It never selects, activates or executes a successor in the same invocation, and never falls back to another independent task on failure. Next selection/activation is a separate authorized action/invocation. A FAIL/BLOCKED prerequisite does not make dependent tasks runnable; preserve their PENDING status and report the dependency obstacle. Never fabricate PASS, alter expectations to pass or silently retry.

## Execution roles and independent gates
Keep CODE implementation, independent TEST and live/UI VERIFY in separate tasks by default, with explicit blocker IDs between them when genuinely dependent. A coding agent may record source-only completion but cannot claim independent TEST/VERIFY PASS. Independent JR executes only its exact test packet, makes no product fixes or workflow promotion, captures raw commands/exits/observations, writes only its authorized report, verifies required delivery, and STOPS. Coding-agent review of the report is required before a dependent task becomes READY. A human-approval-only task or standalone verification task need not manufacture CODE work. Assign each task's actual agent explicitly; do not infer agent identity from the directive invocation.

## Preferred size
Score implementation surface, environment/dependency uncertainty, behavioral surface, verification surface and decision/recovery surface, each 0–2. Total 0–3 preferred; 4–5 split unless tightly coupled in one deterministic workflow; 6–7 split; 8–10 must split. Never artificially lower a score to combine stages.

## Mandatory split triggers
Split for more than three independent acceptance outcomes or implementation verbs, multiple verification workflows or architecture decisions, independently verifiable sequential subtasks, donor/import/toolchain plus live behavior, unknown environment discovery plus implementation, or an early failure that changes later steps. Keep one bounded semantic branch and working set. CODE, TEST and VERIFY are distinct by default.

## Promotion boundary
Only human-approved fresh tasks may be promoted. At most one task is ACTIVE across the queue; multiple independent tasks may be READY or PENDING. Promotion and selection do not authorize execution of more than one task per CWAL invocation. Product tasks cannot edit ICC. No destructive cleanup, live actions, merges, force pushes or main pushes without exact separate authority.
