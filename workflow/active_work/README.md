# Active Work

This directory contains only human-promoted tasks. Planning and Brainstorm do not authorize execution.

## Execution states and authority

Executable task files have `ACTIVE` or `QUEUED`; completed/retired tasks belong in `workflow/archive/`. Exactly one ACTIVE means execute only that task; zero or multiple ACTIVE means stop. Never infer a task from numbering, handoff, dependencies or Planning. This directory selects the task, and `handoff.md` must agree. A conflict means stop. Do not delete, archive or rewrite a conflicting task merely to make the state look consistent; report the exact conflict for authorized reconciliation.

## Stage ownership

UMIG CODE tasks: ChatGPT coding agent owns product changes, source readback, source checkpoint, test packets and workflow advancement unless the ACTIVE task explicitly assigns bounded CODE execution to OpenCode. OpenCode CODE execution follows `workflow/IDENTITY_MAP.md` and `workflow/adapters/opencode.md`: it may implement, run declared targeted tests, make bounded fixes, record results and STOP within one assigned task. It cannot independently certify a TEST/VERIFY stage or advance a successor. UMIG TEST/VERIFY tasks: OpenHands normally acts as JR and follows `operation cwal.md`; JR runs only the current test packet and returns raw evidence, never edits product files, modifies task state or fixes bugs. **One-task human exception (2026-09-20): `UMIG-EM-003-T` may use experimental OpenCode V2 JR in a separate Legion worktree ONLY after that task is ACTIVE with an exact local-only Go test packet. This is not a general JR replacement and does not permit Docker, operator data/services, sudo or automatic shell approval.** The coding agent reads actual independent JR evidence before deciding PASS/completion. A source-only checkpoint is not proof of a test or runtime acceptance. FAIL/BLOCKED stops automatic advancement; fix only within an explicitly authorized coding task.

For every ACTIVE task, specify objective, agent owner, stage, allowed files and explicitly NEW paths, acceptance criteria, exact tests, safe environment, evidence/report location, commit/push permission and STOP condition. Missing fields are a blocker, not an invitation to infer. Established diagnostic findings should be reused; do not repeat investigation without new evidence. CODE may correct a focused failing check within its original scope and record retests. TEST/VERIFY never repairs. No global shell auto-approval or unbounded workstation access.

## Ordered authorized work

A human may promote an ordered sequence; its first task is ACTIVE and later tasks QUEUED. Each queued task names `Previous`; each non-final task names `Next`.

After the ACTIVE task meets its *own stage-specific* gate: coding agent archives it, follows its explicit Next, checks the named successor exists as QUEUED and links back via Previous, changes only that successor to ACTIVE, synchronizes handoff, and commits/pushes this state together. JR never performs advancement. OpenCode's autonomous single-task adapter does not itself grant advancement: the separately authorized workflow owner must review the gate and perform the transition. If Next is absent, stop. If the successor is missing, in Planning, not QUEUED, or has mismatched Previous, stop for human promotion/repair; never infer it.

When a TEST/VERIFY task becomes ACTIVE, coding agent must put one unambiguous `JR TEST TASK` packet in `handoff.md` with the exact commands/actions, expected result, evidence and safe boundary before invoking the explicitly authorized JR. Tester reports PASS/FAIL/BLOCKED but coding agent controls closure. No destructive tests on user data without an explicit safe setup/backup and permission.

## ICC context rule

Read relevant current ICC via the prescribed workflow. If stale/missing, request a bounded BLACK SHEEP WALL refresh. Only BLACK SHEEP WALL modifies ICC. Staleness does not authorize tester scope expansion, cross-branch browsing or unverified completion.

Do not store future backlog, brainstorm or speculative work here.

## OSJT execution map

For the active OS.js Toolkit sequence, read `OSJT_QUEUE.md` after this file and before opening an individual task. The queue map is a routing aid: the individual ACTIVE task and matching `handoff.md` remain execution authority. If the map, task, or handoff disagree, stop and reconcile them before editing or testing. In particular, an OSJT TEST task cannot become an OpenCode CODE task merely because OpenCode is the current runtime.
