# OpenCode adapter — persistent Ubuntu workstation / Qwen

OpenCode works only on branch `opencode`. Record branch, HEAD and `git status --short` before execution. Preserve unrelated work. Never switch, reset, clean, restore, stash, rebase, merge, force-push, touch another worktree, or edit ICC. Fetch/pull only with explicit authorization. Never push main.

## OPERATION CWAL — continuous queue routing

Read `AGENTS.md`, `operation cwal.md`, repository-root `handoff.md` and the queue explicitly named by the root handoff. `workflow/handoff.md` is a non-authoritative pointer, not a status source. The approved queue invocation authorizes selecting and executing eligible packets assigned to OpenCode within their exact scope; it does not authorize new tasks or override packet gates. Historical OTR packets and retired OTR queues are non-executable.

If root handoff names an approved non-retired queue and says `QUEUE SELECTION`, or identifies a correctly checkpointed ACTIVE task, do not report NO ACTIVE TASK. Resume a valid ACTIVE task first; otherwise validate the queue and select an eligible task under CWAL dependency rules. If ACTIVE conflicts with packet or queue state, reconcile from actual repository evidence or report BLOCKED on an unsafe integrity conflict; never silently reset statuses.

## Execute without redundant confirmation

Once a packet is eligible, assigned to OpenCode and all gates are satisfied, checkpoint exactly one task ACTIVE in its packet, queue and root handoff, then execute immediately. Execute an already valid ACTIVE packet immediately. Do not ask whether to execute, activate, pause, select the next task or review changes when the approved queue already authorizes that decision. Do not finish with `Would you like me to...` while authorized OpenCode work remains. Status summaries, plans and successful commits or pushes are not terminal conditions.

Read the selected packet, verify branch/freshness and exact read/write scope, execute only authorized actions, preserve raw evidence and record an evidence-backed terminal result. COMPLETE requires actual packet acceptance criteria and required independent verification and delivery; file existence, chat summaries and checkpoint labels alone are insufficient. Never invent endpoints, commands, tests or PASS or impersonate JR. Product YAML, services, devices, sudo, installs, downloads and live side effects require exact packet authority. Commit/push only where specifically authorized; verify required delivery. Inspect changed and staged paths before each commit; never stage unrelated files.

## Task boundary versus queue boundary

A packet's STOP, RETURN or delivery instruction ends that task's execution, not the approved CWAL coordinator loop, unless it explicitly revokes the queue or requires a human decision. Finish evidence, terminal status and consistent packet/queue/root-handoff checkpoint first. Rescan and automatically select another eligible task assigned to OpenCode. Run one task at a time; never mark multiple tasks ACTIVE. Do not retry FAIL/BLOCKED in the same invocation; keep dependents ineligible and continue independent eligible work. For JR or another agent's task, prepare exact handoff without impersonation and continue other eligible OpenCode work.

Stop this OpenCode runner only when no eligible OpenCode task remains, all tasks are terminal, a mandatory role handoff prevents further work in this session, or a real safety, integrity or missing-authorization gate blocks all remaining eligible work. Record exact reason, task IDs, evidence and blockers. A blocked build task does not stop independent discovery/design tasks. Do not claim the whole queue COMPLETE while unresolved tasks remain.
