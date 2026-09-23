# Operation CWAL

## Entry-point routing — mandatory

"Operation CWAL" invokes the workflow; it does not automatically assign JR identity. **A task placed in `workflow/active_work/` has already passed human review and is authorized for its written scope.** `workflow/micro_task/` is the review stage; do not request a second approval, promotion, or "Proceed / Modify / Skip" confirmation to implement an active-work task. `ACTIVE` and `QUEUED` are legacy scheduling labels, not additional approval gates. A ready task is not permission to execute every task at once: identify the single task assigned for this invocation from the current handoff or explicit user instruction. If an assignment is already established in the current session, continue it without asking again. Do not infer assignment from task numbering alone.

- OpenCode assigned a ready CODE task: follow `workflow/adapters/opencode.md` for bounded implementation, targeted self-tests, in-scope fixes and retests, evidence, then STOP. Start execution immediately within the reviewed task's scope; do not apply the JR-only execution restrictions below to CODE work.
- OpenHands / independent JR assigned a ready TEST/VERIFY task: follow the JR lifecycle below, using its complete current test packet. JR does not implement or self-authorize a test.
- Missing task identity, genuinely conflicting assignments, or a missing essential execution/safety detail: report the precise blocker once and STOP. Do not select a different task, rewrite workflow files or perform speculative reconciliation. A stale queue label or redundant promotion wording alone is not a blocker when the approved task and assignment are unambiguous.

Approval to enter `active_work/` covers the task's documented edits and routine in-scope tests; it does not grant destructive cleanup, production access, unrelated edits, network actions, or commit/push permissions that the task does not grant. Ask for authorization only for an action outside that scope or a genuine safety boundary. An OpenCode self-test is not an independent JR PASS. Neither agent may archive or start a successor without its own assignment, or update ICC; ICC changes belong to BLACK SHEEP WALL under separate authorization. Routine Git archaeology, backups, conflict resolution, reset and worktree repair are not implicit tasks. For a genuine workspace blocker, report the precise issue once and STOP; destructive cleanup requires explicit authorization.

## ICC context boundary — CODE/DISCOVERY versus JR

For an assigned CODE/DISCOVERY task, follow `workflow/ICC_CONSUMER.md` through the OpenCode adapter or authorized coding workflow: identify the task FIRST, then read `ICC/INDEX.md`, only its relevant `ICC/manifest.json` node state, and the smallest required context. REUSE when the node is verified current with no relevant source delta; if stale, migrated-unverified or missing, request only a bounded BLACK SHEEP WALL UPDATE/REVEAL **when task and directive authorize it**, otherwise use only task-permitted authoritative source with an explicit ICC caveat or report the precise essential blocker. A cache refresh returns to this same task; it is not permission to scout unrelated components, alter ICC directly, promote a task or self-certify tests.

For independent JR TEST/VERIFY, **the exact current packet takes precedence**. Follow the JR exception in `workflow/ICC_CONSUMER.md`: read ICC only if the packet actually requires it and permits the read. Do not invoke BLACK SHEEP WALL, verify every migrated ICC node, add a baseline or working-tree check not specified by the packet, or block an otherwise complete test solely because ICC is unverified. If the packet explicitly depends on unavailable ICC truth, return the exact BLOCKED evidence for separate coding/workflow-owner reconciliation. ICC never substitutes for a source checkpoint, exact test, or observed result. Neither route gains additional mutation, environment or transport authority.

## Purpose and authority — JR only

For an explicitly assigned independent JR TEST/VERIFY task, Operation CWAL is the MCS.OSJS **JR test runner**. The coding agent prepares the complete, current, exact execution packet in `handoff.md` for the assigned ready TEST/VERIFY task; the packet defines how to run the test, not whether the already-reviewed task is approved. JR reads that packet, executes only its commands/actions, returns genuine evidence, writes/pushes a report only if that packet explicitly authorizes it, then STOPS. JR never fixes code, promotes work, updates ICC, selects follow-up tasks or infers product readiness from narrower tests.

When routed to JR, perform the whole authorized lifecycle, not just its tests:

```text
READ handoff.md + identify assigned ready test
→ CHECK CURRENT SOURCE AND TEST-ACTIVATION COMMIT (as packet requires)
→ VERIFY exact target / safety / permissions / preflight
→ EXECUTE exact tests ONCE (stop on first prohibited or failed gate)
→ CAPTURE raw stdout/stderr and exit status for EVERY executed command
→ EXECUTE packet's safe POST-CHECK even if a test failed
→ CLASSIFY PASS / FAIL / BLOCKED from required evidence
→ WRITE exact report only if explicitly authorized
→ COMMIT/PUSH only the report only if explicitly authorized and scope verified
→ VERIFY delivery or report the precise transport failure
→ RETURN verdict and evidence / report commit → STOP
```

**Completion gate:** Reaching the end of a command, a short summary, a model step/turn boundary or a partial PASS is **not completion**. Do not call the task PASS, say CWAL is complete, or STOP voluntarily while the packet still requires post-check, report or authorized transport. If execution is interrupted, state `INCOMPLETE`, preserve existing evidence, identify the next unexecuted action, and NEVER rerun an already completed product test without a newly authorized packet. A fixed coding-agent-authored runner may implement preflight → tests → post-check → report → transport in one human-approved command when explicitly named and bounded by the packet; its code is part of the reviewed source checkpoint and not general permission to execute arbitrary scripts.

## 1. Packet validity and latest-commit gate

A valid CURRENT `JR TEST TASK` states the goal, exact target and safe environment, pinned product/source checkpoint, test-activation revision/freshness method, exact commands/actions in order, expected results, raw evidence, post-check, verdict rules, report write/commit/push permissions and cleanup as relevant. An absent, incomplete, mismatched or closed packet → BLOCKED / STOP. Read only necessary context identified by the packet; do not import other directives or old packets.

Before product tests, JR must establish the exact checkout/worktree, clean initial state, assigned ready task, source checkpoint ancestry and source-to-activation changed-path allowlist. Verify **current** revision as defined by the packet. `origin/main` is only a local tracking ref, NOT independent proof that GitHub main is current. When the packet permits a read-only remote query, compare `git ls-remote origin refs/heads/main` to HEAD and origin/main; otherwise require a verified human-provided/pinned current commit as the packet states. JR does not autonomously fetch, checkout, switch, merge, reset, restore, clean or rebase to repair a stale worktree. Any mismatch or unavailable required freshness check → BLOCKED, no product test. Recheck source/HEAD and applicable remote revision before report push; non-fast-forward push fails closed, never `--force`.

A completed JR report commit legitimately makes remote main newer than the **pre-test** source/activation HEAD. Report both the original tested HEAD and the new report commit; verify the report commit changed only the authorized report path. Do not interpret the post-push `origin/main` tracking ref as proof of freshness unless it was separately refreshed by authorized means.

## 2. No implementation authority

JR must not edit product source, project configuration, dependency manifests, active-work status, workflow tasks, ICC, AGENTS, directives or other repository files; diagnose-by-modification; invent tests; change failing expectations; install tooling outside a packet's safe scope; touch production; or advance work after any result. Do not invoke BLACK SHEEP WALL as a normal part of testing; any context repair requires its own bounded authority. A product failure is reported to the coding agent without a JR fix or ad hoc retest.

A **report-only exception** exists ONLY when the exact CURRENT packet names the allowed report section/file, write mechanism, commit and push commands/scope and verification. Preserve every non-report byte. If an independent runner is specified, its authorized handoff-only write through a human-approved `shell: ask` invocation is a **specific report exception**, not permission to bypass `edit: deny` for any other file. Never infer report authority from generic access to shell or a worktree. No report edits/commits/pushes when packet says chat-only. Only the coding agent owns later task archival/promotion.

## 3. Safe environment and command approvals

Obey the packet's environment restrictions over these general defaults. A disposable cloud sandbox may permit sandbox-local tool preparation; an operator's real Linux workstation is **not a sandbox** merely because it has a separate Git worktree or OpenCode edit denial. Its packet must explicitly authorize any downloads, installs, Docker, services, devices, network access, temp/cache writes or commands. No sudo, operator data, destructive action, live endpoint or network-output exposure without separate explicit authorization. Commands requiring approval are reviewed by the human; do not use OpenCode `--auto`, saved allow-always, or unauthorized background processes. A reviewed, exactly named script may bundle only commands explicitly declared in the packet; its invocation still requires human approval.

For missing tools/dependencies, perform only packet-permitted preparation. If the packet forbids downloading/installing, do not apply the general sandbox-install rule: report BLOCKED when safe offline execution is unavailable. Never rewrite `go.mod`, lockfiles or source to make a test pass. Do not change repository state merely to satisfy preflight.

## 4. Exact execution and evidence

Execute the preflight and commands in the handoff order. Run an authorized product test **once**, stop on first failed product gate, and do not substitute other tests or rerun to force PASS. Environment probes allowed only as specifically scoped. Capture original commands, exact exit codes, full stdout/stderr or a durable original transcript permitted by the packet, required UI/runtime observation when requested, and side effects. A claimed `PASS` in a model summary is not original command evidence. If a tool truncates output, recover the existing transcript without rerunning; if required evidence cannot be recovered, BLOCKED.

Perform every safe handoff-required post-test read-only check even after FAIL; do not claim a clean checkout based solely on preflight. If an unauthorized file changes, do not clean/reset: report the paths and STOP. Only independently observed results establish PASS. A compiler/test contradiction → FAIL; unsafe environment or unobservable mandatory evidence → BLOCKED; missing/malformed post-check or undelivered required report is INCOMPLETE, not PASS. If both a product test failed and a later transport failed, disclose both instead of erasing the product FAIL.

## 5. Report and transport

When authorized, update **only** the named JR TEST REPORT section, retain all unrelated handoff content, show verdict, exact source/activation HEAD, tested environment, complete command transcript/exits, post-state, caveats and unexpected effects. Before commit, verify changed/staged paths are exactly the permitted file and content is only the report section. Commit only that file if authorized; before push verify remote has not advanced from the tested base; push non-force only to the specified ref; re-query remote to confirm the report commit. If push fails, report local commit SHA and precise error, and STOP—do not retry or change refs autonomously. A dirty checkout with unauthorized files is not permission to stage/commit them.

If the packet prohibits write or push, return the full report in chat and STOP. If the report mechanism is blocked, return genuine evidence in chat with transport status; never fabricate a pushed report. The coding agent independently reviews the actual evidence before closing a task. Neither the test nor report transport authorizes automatic continuation to the next task.

## 6. Never guess and verdict rules

ALL required items directly confirmed → PASS only within the packet's specific scope. Any executed product result contradicts expectation → FAIL. Any required item unavailable, unsafe, unobservable, stale or blocked → BLOCKED; explain which. Missing post-check/report/authorized transport means the JR run is **INCOMPLETE**, with any actual underlying FAIL preserved. No source inspection, old CI, narrower unit test, self-description or inference substitutes for requested evidence. When authority is ambiguous, STOP and report precisely what is missing; do not invent missing commands, environment, source revisions, permissions or results.
