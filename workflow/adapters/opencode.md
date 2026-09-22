# OpenCode adapter — persistent Ubuntu workstation / Qwen

## Environment

OpenCode works only on branch `opencode` in this worktree. At startup record branch, HEAD and
`git status --short`. Preserve unrelated work. Never switch, reset, clean, restore, stash, rebase,
merge, force-push, touch another worktree, or edit ICC. Fetch/pull only when explicitly requested.

This is a persistent workstation, not a disposable sandbox. Commit and push are not task-completion
requirements unless the operator or exact packet explicitly requests them. Never push to `main`.

## Continuous CWAL scheduler

1. Read `AGENTS.md`, root `operation cwal.md`, root `handoff.md`, and
   `workflow/active_work/OTR_PROMOTION_QUEUE.md`.
2. Run the handoff-named ACTIVE packet when its packet also says ACTIVE.
3. Close that task with durable evidence and exactly one terminal status:
   - COMPLETE: acceptance and required checks passed.
   - FAIL: an executed check contradicted expectations.
   - BLOCKED: required authority, input, environment or observable evidence was unavailable.
4. On COMPLETE, activate its dependency-satisfied successor and execute it immediately.
5. On FAIL/BLOCKED, preserve evidence, mark direct dependents SKIPPED-BLOCKED, then select the next
   independent QUEUED task whose prerequisites are COMPLETE and execute it.
6. Never retry a failed command merely to obtain PASS. Corrective CODE requires its own eligible
   packet; unsafe or destructive recovery still requires explicit authority.
7. Continue until no eligible tasks remain. Then set handoff to NONE and return one exhaustion
   summary listing COMPLETE, FAIL, BLOCKED and SKIPPED-BLOCKED tasks plus evidence paths.

Task boundaries remain strict even though the invocation continues: finish evidence/state for one
task before opening another, use only its allowlisted files/actions, and never infer a missing API,
test or PASS. Product YAML, services, devices, sudo, installs, downloads and live side effects need
exact packet authority. OpenCode does not impersonate independent JR verification.
