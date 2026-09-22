# OpenCode adapter — persistent Ubuntu workstation / Qwen

OpenCode works only on branch opencode. Record branch, HEAD and git status --short before execution. Preserve unrelated work. Never switch, reset, clean, restore, stash, rebase, merge, force-push, touch another worktree, or edit ICC. Fetch/pull only with explicit authorization. Never push main.

## Single-task routing
Read AGENTS.md, operation cwal.md, root handoff.md and the queue explicitly named by handoff. If Current task is NONE, report NO ACTIVE TASK and STOP. Old OTR packets and the retired OTR_PROMOTION_QUEUE.md have no execution authority.

If a task is explicitly named, require its packet to exist and say ACTIVE, verify branch/freshness and exact read/write scope, execute only that packet, preserve raw evidence and record COMPLETE, FAIL or BLOCKED. Perform only packet-authorized commit/push and verify delivery if required. Then STOP regardless of outcome.

Never automatically activate or execute a successor, scan for another eligible task, fall back to numerical order, or run a continuous queue. A successor requires separate human-authorized activation and a later invocation. Do not infer endpoints, tests or PASS; do not impersonate independent JR verification. Product YAML, services, devices, sudo, installs, downloads and live side effects require exact packet authority.
