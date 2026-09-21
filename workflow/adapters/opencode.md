# OpenCode adapter — Ubuntu / Qwen

This adapter changes execution mechanics, NOT task authority. Identify OpenCode identity, read `workflow/IDENTITY_MAP.md`, `AGENTS.md` and `workflow/active_work/README.md`, then the ONE explicitly assigned ready task in `workflow/active_work/` and its execution packet. Admission to `active_work/` is approval; `ACTIVE`/`QUEUED` labels and `OSJT_QUEUE.md` are legacy scheduling hints, not another promotion gate. Consult the queue only if the assigned task needs its dependency or test reference. Do not use an old handoff or queue label to override a newer explicit human assignment; stop on substantive scope/stage/owner conflicts. Read the JR-only lifecycle in `operation cwal.md` only if explicitly assigned independent JR.

## Autonomous CODE lifecycle

1. Resolve exactly one execution-ready CODE task explicitly assigned to OpenCode. Its objective, allowed production/test files (including explicitly NEW paths), acceptance criteria, exact targeted test, environmental permissions and stop condition must be explicit. Other ready tasks do not block it. If assignment or required fields are absent or substantively contradictory, report the precise blocker once and STOP; never repair authority by guessing, rewriting ICC or deleting tasks.
2. Read only named files, directly relevant predecessor evidence and existing diagnostic findings. Do not repeat investigations without new evidence. If the cause is already established, implement directly.
3. Check `git status --short` and HEAD once before edits. Preserve pre-existing changes; if they overlap task scope or prevent attribution, STOP rather than reset/restore/clean/stash. Use a dedicated authorized worktree/branch; never silently switch branches. Do not turn unrelated Git cleanup into a task.
4. Edit only the task allowlist. Inspect the resulting diff. Make the smallest correction when a compiler or targeted test identifies a local defect; no unrelated refactor, dependency install, network service, sudo, operator data or production endpoint without explicit packet authorization.
5. Run the named targeted test first and capture command, exit code and actual output. Bounded corrections and retests are permitted for CODE only, with each attempt recorded; do not call preliminary self-tests independent JR PASS. Run broader tests only when explicitly required by the task.
6. Record task ID, source HEAD, changed paths, implementation summary, actual commands/exits, test evidence, limitations and next required action in the task-authorized report location or chat. Commit/push only if the assigned task explicitly grants branch/ref, file scope and transport permission. No force-push.
7. STOP after this task. Do not archive, assign or start the next task; do not edit ICC. A FAIL/BLOCKED result stops rather than spawning an automatic repair pre-task. No successor promotion or queue/ICC synchronization is required to begin a separately assigned ready task.

OpenCode's terminal, coding and Playwright capabilities are available only within the current task's explicit safety and file/target boundaries. Do not grant global auto-approval for shell actions, destructive commands, browser access or downloads. A locally installed Qwen model does not change permissions. The local Ubuntu workstation is not a disposable sandbox.

## Independent test handoff

For TEST/VERIFY, the coding agent may author an exact, source-pinned packet and request separately assigned JR. OpenCode must not both implement and independently certify the same TEST/VERIFY gate under one identity. JR executes the JR lifecycle in `operation cwal.md`, reports evidence and stops; the authorized workflow owner reviews evidence before closing that test. Closure does not automatically assign another ready task.
