# Microtask Rules

## Context and authority
Before repository-dependent planning, read `ICC/INDEX.md` and relevant valid context; verify against actual Git source. Stale ICC requires separately authorized bounded BLACK SHEEP WALL refresh; only BLACK SHEEP WALL edits ICC. Planning is not execution authority. Human approval is required to promote tasks into a fresh executable queue. Historical OTR packets and retired queues remain non-executable.

## Core rule
One microtask = one primary outcome = one detailed task file. Keep tasks small, independently verifiable, and reuse source-verified existing work rather than recreating it. Never infer completion from an old report or status alone.

## Required task format
Every new task packet must contain these exact prominent fields:

Task ID: <unique stable ID>
Task Name: <one specific outcome>
Blocker Task: NONE | <comma-separated exact Task IDs>
Status: PENDING | READY | ACTIVE | COMPLETE | FAIL | BLOCKED
Assigned Agent: <explicit agent and role>
Stage: DISCOVERY | DESIGN | CODE | TEST | VERIFY

Then include objective (one measurable result); scope (exact read/write paths and forbidden actions); bounded execution (exact commands/actions for TEST/VERIFY); at most three independently testable acceptance outcomes; evidence (exact report path, raw observations, source/commit reference and handoff); completion/delivery (verdict rules, permitted status edits, exact commit/push/ref and remote verification when required, or explicitly no commit/push authority); five-dimension sizing and rationale.

Do not use Previous/Next for routing or dependencies. Task ID is not an execution order. Every blocker must resolve to a real approved task in the same queue; no cycles, self-dependencies or historical/retired executable blockers. Reused completed work is source and verified evidence input, not invented COMPLETE.

## Continuous dependency-aware readiness and selection
PENDING means not executed; waiting for blockers remains PENDING, not BLOCKED. READY means every blocker is COMPLETE with independently verified required evidence and all packet, environment, role and safety gates are satisfied. `Blocker Task: NONE` waives dependencies only. A human-approved promoted queue authorizes CWAL to select and process eligible tasks continuously in one invocation. `handoff.md` must name the approved, non-retired queue; NONE or retired means NO ACTIVE QUEUE / STOP. Presence of a file alone never grants approval.

At each iteration inspect all nonterminal task headers and resolve blocker IDs and evidence, without loading every packet body. Select any eligible task, lowest stable Task ID as default tie-breaker, activate exactly one in handoff/packet/queue, execute it, record its actual terminal result and required delivery, clear ACTIVE, and rescan. If a desired task is blocked by prerequisites, first run an eligible prerequisite or choose an independent eligible task. Continue after FAIL/BLOCKED on other eligible tasks; dependents of failed/blocked prerequisites stay PENDING. Do not retry a failed/blocked task within the same invocation without explicit new authorization. Never run tasks concurrently, bypass a blocker, infer PASS, or alter expectations to pass.

Stop when all tasks are terminal or no eligible tasks remain. Report COMPLETE, FAIL, BLOCKED and unresolved PENDING with reasons and preserve a truthful handoff for later authorized continuation. A task unable to run because of role or safety authority remains PENDING with reason while other eligible tasks may proceed. Queue-integrity failures or unsafe global state require STOP. Do not automatically create/promote new tasks or execute another queue.

## Execution roles and independent gates
Keep CODE implementation, independent TEST and live/UI VERIFY in separate tasks by default with genuine blockers. A coding agent may record source-only completion but cannot claim independent TEST/VERIFY PASS. JR executes only its exact activated test packet, makes no product fixes or workflow promotion, captures raw commands/exits/observations, writes only its authorized report and stops its task; the queue coordinator may rescan only within its own permissions. Coding-agent review of JR evidence is required before dependent work becomes READY. A role handoff requiring another agent must checkpoint; do not impersonate that agent. Human-approval-only or standalone verification tasks need not manufacture CODE work.

## Preferred size
Score implementation surface, environment/dependency uncertainty, behavioral surface, verification surface and decision/recovery surface, each 0–2. Total 0–3 preferred; 4–5 split unless tightly coupled in one deterministic workflow; 6–7 split; 8–10 must split. Never artificially lower a score to combine stages.

## Mandatory split triggers
Split for more than three independent acceptance outcomes or implementation verbs, multiple verification workflows or architecture decisions, independently verifiable sequential subtasks, donor/import/toolchain plus live behavior, unknown environment discovery plus implementation, or an early failure that changes later steps. Keep one bounded semantic branch and working set. CODE, TEST and VERIFY are distinct by default.

## Promotion boundary
Only human-approved fresh tasks may be promoted. At most one task ACTIVE at once; multiple tasks may be READY or PENDING. Promotion of a queue permits continuous dependency-aware processing only under CWAL and each packet's explicit permissions; it does not authorize destructive cleanup, live actions, merges, force pushes, main pushes or ICC edits. No task or queue is COMPLETE without required evidence and delivery.
