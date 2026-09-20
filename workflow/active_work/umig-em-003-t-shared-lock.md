# UMIG-EM-003-T — TEST: Shared writer lock (one OpenCode JR trial)

Status: ACTIVE — human-approved one-task OpenCode V2 JR exception; coding agent archived UMIG-EM-003 SOURCE ONLY at `949d7a8eb328dfa66f8359ffe81cfe359b8db2b7` on 2026-09-20. No TEST result yet.
Stage / owner: TEST / independent OpenCode V2 on separate Legion JR worktree, experimental one-task exception to OpenHands; OpenHands remains fallback. ChatGPT reviews report and owns state changes.
Previous: UMIG-EM-003 (archived COMPLETE / SOURCE ONLY).
Next: UMIG-EM-003-V (PLANNED; separate human promotion and independently safe runtime target required).

## Primary outcome
Independently compile/test the committed shared Linux flock implementation and Simulator/Replicator writer integrations using ONLY existing Go unit, synthetic fixture, concurrency and race tests on a clean Legion JR worktree. No Docker, sudo, operator Modbus endpoints, installed config or persistent services; existing Go fixture tests may use ephemeral loopback listeners in temp test environments ONLY.

## Acceptance
1. Exact handoff preflight pins CODE source `949d7a8e`, verifies current HEAD matches activated test handoff, clean worktree, Go >=1.25, dependencies already available offline and approved single-test permissions. No checkout mutation, installs, downloads or unauthorized shell commands by JR; ambiguity => BLOCKED.
2. Run ONLY the exact module suite and named focused race/concurrency commands in handoff, with a bounded `timeout`, Go module read-only/offline, no source changes, no Docker or host service actions. Check lock timeout/process crash release, foreign producer preservation under race, direct writer no-mutation and committed-ACK failure truth. Report verbatim stdout/stderr and exits; stop on first failure/blocker.
3. OpenCode has edit permission DENIED: return evidence and PASS/FAIL/BLOCKED IN CHAT ONLY (explicitly NO handoff edit/commit/push), then STOP. ChatGPT will adjudicate and update handoff later. The absence of a written report is not a test outcome; cannot claim PASS until evidence reviewed.

`handoff.md` is sole exact JR test authority; this task is not an independent command packet. `~/opencode-smoke/opencode.jsonc` still includes an old no-CWAL sentence and user must remove ONLY that sentence before starting JR while keeping shell ask, edit deny, subagent deny and no `--auto`. Worktree is not a security sandbox. Sizing 0/1/0/1/1=3. TEST PASS does NOT imply live VERIFY or production readiness. ICC only BLACK SHEEP WALL.
