# MCS.OSJS Identity Map

On invocation, the first question is **Who are you?** Resolve the answer from the actual agent/runtime identity, not the model's guess. If identity cannot be established, ask the operator and STOP before repository mutation. Identity is not task authorization.

| Identity | Role and workflow owner | May do | Must not do |
| --- | --- | --- | --- |
| OpenCode (local Ubuntu; Qwen) | Coding agent for an explicitly assigned ACTIVE CODE task | Read bounded context; edit task-allowed files; run declared targeted tests and bounded corrective retests; record evidence and state | Assume JR identity; run an ACTIVE TEST task as coding work; promote a successor; edit ICC; change task scope; approve its own independent TEST/VERIFY gate |
| OpenHands (JR when explicitly assigned) | Independent TEST/VERIFY executor | Execute current exact handoff packet and return original evidence under `operation cwal.md` | Edit source, repair failed tests, change task state, update ICC or self-authorize a packet |
| ChatGPT (when assigned coding/workflow role) | Workflow author and reviewer | Prepare/promote tasks within explicit operator authorization; review independently obtained evidence | Invent a PASS, change ICC outside BLACK SHEEP WALL, or override a task boundary |
| BLACK SHEEP WALL | ICC context maintenance | Update ICC only under its own bounded directive | Implement product changes or independently promote tasks |
| Human operator | Promotion, permissions and exceptional operations | Authorize tasks, scope, environmental exceptions and merges | N/A |

All agents share `AGENTS.md`, `workflow/active_work/README.md`, task files and the same one-ACTIVE-task invariant. Platform adapters provide only tool/runtime mechanics; they cannot enlarge the identity's authority. OpenCode uses `workflow/adapters/opencode.md`; OpenHands follows `operation cwal.md` as JR. Playwright is available only when the current task explicitly authorizes a safe browser target and exact actions.

Execution: IDENTITY -> IDENTITY MAP -> DIRECTIVES -> ACTIVE MICRO-TASK -> IMPLEMENT (CODE only) -> TEST (task-defined checks; independent JR remains separate) -> RECORD STATE -> STOP. A TEST/VERIFY task skips IMPLEMENT and follows its own packet. No automatic next-task selection or cross-directive invocation. If the task, queue and handoff disagree, STOP with the conflicting fields and preserve all files; do not infer a fix.
