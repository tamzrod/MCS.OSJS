# OpenCode adapter — persistent Ubuntu workstation / Qwen

## Permanent branch and environment — mandatory

OpenCode operates on the user's persistent workstation, NOT a disposable sandbox. Its sole authorized development branch is `opencode`. Before ANY task or edit, verify `git branch --show-current`, `git rev-parse HEAD`, `git status --short` and `git worktree list`. If current branch is not exactly `opencode`, STOP with the mismatch; never silently switch, checkout, reset, rebase, merge, clean, restore or stash to fix it. Do not implement on `main`, `osjt-*`, detached HEAD, or another agent's branch. Do not touch other worktrees. Preserve all pre-existing work; overlapping edits or an unsafe workspace => BLOCKED/STOP. Fetch/pull only when safe and explicitly requested; a pull must not overwrite local commits or dirty work. No automatic merging or conflict resolution.

## Transport boundary

The human has permanently authorized one non-force task-completion commit and push to the `opencode` branch for every executable task run through OPERATION CWAL. A task is not COMPLETE until its exact allowlisted changes and evidence are committed, `git push origin HEAD:opencode` succeeds, and `git ls-remote origin refs/heads/opencode` confirms the report commit. Never push to `main`, force-push, merge, create a PR, or include unrelated paths. A packet that says no commit, no push, local-only, or chat-only is invalid and must be repaired before execution. A BLOCKED/FAIL run may publish only its allowlisted evidence/state when the packet explicitly defines that terminal report path; it must not advance the queue. No global git reset, even for synchronization, without separate explicit approval and verified backups.

## CWAL routing — autonomous, one task per invocation

The operator's complete prompt is `pull latest and do operation cwal`. Treat OPERATION CWAL as workflow routing, NOT automatic JR identity. Read `workflow/IDENTITY_MAP.md`, `AGENTS.md`, `operation cwal.md`, root `handoff.md` and `workflow/active_work/OTR_PROMOTION_QUEUE.md` as needed. Execute only the single task named ACTIVE by root handoff whose packet also says `Status: ACTIVE`; never fall back to a queued or numerically next task. If a packet is incomplete or evidence unavailable, report the precise blocker and STOP, never ask 'what would you like me to do?'. Parent OTR-004..014 templates are not executable children.

Execute exactly ONE task per invocation: read the packet, inspect permitted sources, perform only its authorized edits/actions and targeted self-tests, capture actual evidence, write only packet-allowlisted workflow state, commit only allowlisted paths, push non-force to `origin/opencode`, verify delivery, report, and STOP. STOP is a hard per-task boundary. Do not self-trigger a second cycle or pretend chat output or an unpushed commit is persistent evidence. On FAIL/BLOCKED do not advance. No ICC edits: BLACK SHEEP WALL alone owns ICC.

CODE and independent TEST/VERIFY remain separate. OpenCode must not act as OpenHands/JR or certify its own changes independently. No production YAML, services, devices, sudo, installs, downloads or other workstation side effects without exact packet authority. Preserve unrelated changes and never invent tests or completion evidence.
