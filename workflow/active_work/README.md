# Active Work

This directory contains human-approved, execution-ready tasks. Moving a task here is its approval/promotion gate; no second promotion is required to start it. Planning and Brainstorm do not authorize execution. Multiple ready tasks may coexist here, but an agent executes only one explicitly assigned task per invocation.

## Readiness, assignment and authority

A task in `workflow/active_work/` is READY, not automatically assigned to every agent. The current assignment identifies one task ID, owner and stage in `handoff.md` (or an explicit current human instruction identifying that ready task). The assigned task file defines objective, scope and tests. If no task is assigned, report that once and STOP; do not pick a successor by numbering or infer an assignment from the queue. A stale handoff does not override a newer explicit human assignment, but a substantive conflict about scope, stage or ownership requires clarification rather than speculative edits. Other ready tasks do not block execution merely because they coexist in this directory.

`ACTIVE` and `QUEUED` in legacy task files or queue maps are historical scheduling labels, not an additional authorization or promotion gate. Do not rewrite every ready task, queue map, handoff or ICC just to synchronize those labels. Never delete or archive another task to manufacture a sole ACTIVE state. The queue is an ordering/reference aid, not execution authority. The current assigned task and its explicit execution packet govern what runs.

## Stage ownership

UMIG CODE tasks: ChatGPT coding agent owns product changes, source readback, source checkpoint, test packets and workflow advancement unless the assigned task explicitly gives bounded CODE execution to OpenCode. OpenCode CODE execution follows `workflow/IDENTITY_MAP.md` and `workflow/adapters/opencode.md`: implement, run declared targeted tests, make bounded fixes, record results and STOP within one assigned task. It cannot independently certify a TEST/VERIFY stage or start a successor. UMIG TEST/VERIFY tasks: OpenHands normally acts as JR and follows `operation cwal.md`; JR runs only the current test packet and returns raw evidence, never edits product files, modifies task state or fixes bugs. **One-task human exception (2026-09-20): `UMIG-EM-003-T` may use experimental OpenCode V2 JR in a separate Legion worktree ONLY with an exact local-only Go test packet. This is not a general JR replacement and does not permit Docker, operator data/services, sudo or automatic shell approval.** The coding agent reads actual independent JR evidence before deciding PASS/completion. A source-only checkpoint is not proof of a test or runtime acceptance. FAIL/BLOCKED stops automatic continuation; fixes require an authorized CODE task.

Each ready task must specify objective, agent owner, stage, allowed files and explicitly NEW paths, acceptance criteria, exact tests, safe environment, evidence/report location, commit/push permission and STOP condition. Missing execution-critical fields block that task, not unrelated ready work. Reuse established diagnostics; do not repeat investigation without new evidence. CODE may correct a focused failing check within its original scope and record retests. TEST/VERIFY never repairs. No global shell auto-approval or unbounded workstation access.

## Single-task lifecycle

Human approves a task by placing it in `active_work/` and assigns one ready task to an agent. The agent executes that task's authorized scope, captures evidence, reports and STOPS. Completion does not assign the next task. The workflow owner reviews stage-specific evidence and may archive the completed task separately; the next ready task needs only an explicit assignment, not another file-status promotion. `Previous`/`Next` fields describe dependencies and order, not permission to execute. Do not automatically archive, delete, rewrite statuses, update handoff, commit or push a transition as a prerequisite to running another ready task.

For an assigned TEST/VERIFY task, the coding agent/workflow owner provides one complete source-pinned `JR TEST TASK` packet in `handoff.md` before invoking independent JR. JR runs only that packet and reports PASS/FAIL/BLOCKED without changing task state. No destructive tests on user data without an explicit safe setup/backup and permission.

## ICC context rule

Read relevant current ICC via the prescribed workflow. If stale/missing, request a bounded BLACK SHEEP WALL refresh only when needed for the assigned task. Only BLACK SHEEP WALL modifies ICC. Stale context alone does not authorize tester scope expansion, cross-branch browsing or unverified completion; do not block unrelated execution to cosmetically synchronize ICC.

Do not store brainstorm or unapproved speculative work here.

## OSJT execution map

`OSJT_QUEUE.md` is an optional compact ordering and artifact reference. It is not a second approval ledger and its ACTIVE/QUEUED labels cannot override an explicit assignment of a ready task. Read the assigned task and its required packet first; consult the queue only for relevant dependencies or named test details. A queue-label mismatch alone is not a reason to edit ICC, delete task files, repair Git or block an otherwise valid assignment. A substantive task/packet mismatch is BLOCKED for that task only. An OSJT TEST task cannot become a CODE task merely because OpenCode is the current runtime.
