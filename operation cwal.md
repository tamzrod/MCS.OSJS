# Operation CWAL

## Entry-point routing — mandatory

"Operation CWAL" invokes the workflow; it does not automatically assign JR identity. Execute only the single task named `ACTIVE` in root `handoff.md` when that task file also says `Status: ACTIVE`. `QUEUED`, `NOT EXECUTABLE`, parent, planning, archived, inferred-number, and mismatched tasks are not runnable. Never select a task by numerical order or treat mere presence in `workflow/active_work/` as execution authority. A missing/mismatched ACTIVE identity or packet is BLOCKED / STOP.

- OpenCode assigned a ready CODE/DISCOVERY/DESIGN task: follow `workflow/adapters/opencode.md` and the exact packet for bounded work, targeted checks, evidence, one allowlisted task commit, non-force push to `origin/opencode`, delivery verification, then STOP. Do not apply JR-only source-edit restrictions to authorized CODE work.
- OpenHands / independent JR assigned a ready TEST/VERIFY task: follow the JR lifecycle below, using its complete current test packet. JR does not implement or self-authorize a test.
- Missing task identity, genuinely conflicting assignments, or a missing essential execution/safety detail: report the precise blocker once and STOP. Do not select a different task, rewrite workflow files or perform speculative reconciliation. A stale queue label or redundant promotion wording alone is not a blocker when the approved task and assignment are unambiguous.

**After completing the one ACTIVE task:** write its required evidence, update only packet-authorized workflow state, commit only allowlisted paths, push non-force to the packet's branch, verify the remote SHA, report, and STOP. Do not begin the successor in the same invocation.

An executable packet must name exact read paths, exact write paths, forbidden paths/actions, ordered steps, acceptance, evidence, commit command/scope, push ref, and delivery check. Thin packets, invented/unverified paths, or broad directory writes/deletes are invalid. Completion transport is mandatory but grants no destructive cleanup, production access, unrelated edits, merge, force-push, or push to `main`. An OpenCode self-test is not an independent JR PASS. Archive operations—including manual file deletion, purging, repository rebasing/resetting, or any irreversible state change—are prohibited unless the exact ACTIVE packet and human authority name them. Neither agent may start a successor in the same invocation, and ICC changes require separate BLACK SHEEP WALL authority.

## Purpose and authority — JR only

For an explicitly assigned independent JR TEST/VERIFY task, Operation CWAL is the MCS.OSJS **JR test runner**. The coding agent prepares the complete, current, exact execution packet in `handoff.md`. JR executes only that packet, writes one report-only commit, pushes it non-force to the named branch, verifies delivery, returns genuine evidence, then STOPS. A TEST/VERIFY task is not done until the report commit is delivered. JR never fixes code, promotes work, updates ICC, selects follow-up tasks or infers product readiness from narrower tests.

When routed to JR, perform the whole authorized lifecycle, not just its tests:

```text
READ handoff.md + identify assigned ready test
→ CHECK CURRENT SOURCE AND TEST-ACTIVATION COMMIT (as packet requires)
→ VERIFY exact target / safety / permissions / preflight
→ EXECUTE exact tests ONCE (stop on first prohibited or failed gate)
→ CAPTURE raw stdout/stderr and exit status for EVERY executed command
→ EXECUTE packet's safe POST-CHECK even if a test failed
→ CLASSIFY PASS / FAIL / BLOCKED from required evidence
→ WRITE exact report to the mandatory packet-named report path
→ COMMIT/PUSH only that report after scope and freshness verification
→ VERIFY delivery or report the precise transport failure
→ RETURN verdict and evidence / report commit → STOP
```

**Completion gate:** Reaching the end of a command, a short summary, a partial PASS, or an unpushed local commit is not completion. Do not call the task complete while any required check, evidence write, scope audit, commit, non-force push, or remote delivery verification remains. If interrupted, state `INCOMPLETE`, preserve evidence, identify the next unexecuted action, and never rerun an already completed product test without a new packet. After verified delivery, STOP. Queued work never overrides this one-task boundary.

## 1. Packet validity and latest-commit gate

A valid CURRENT `JR TEST TASK` states the goal, exact target and safe environment, pinned product/source checkpoint, test-activation revision/freshness method, exact commands/actions in order, expected results, raw evidence, post-check, verdict rules, one exact report path/schema, exact stage/commit command, exact non-force push command/ref, delivery verification, and cleanup. Report commit/push authority is mandatory. An absent, incomplete, mismatched, thin, chat-only, no-push, or closed packet → BLOCKED / STOP.

Before product tests, JR must establish the exact checkout/worktree, clean initial state, assigned ACTIVE task, source checkpoint ancestry and source-to-activation changed-path allowlist. Verify **current** revision as defined by the packet. A local tracking ref is not independent proof that its remote branch is current. Compare the packet-named local tracking ref and `git ls-remote` result for the packet-named remote ref to HEAD. JR does not autonomously fetch, checkout, switch, merge, reset, restore, clean or rebase to repair a stale worktree. Any mismatch or unavailable required freshness check → BLOCKED, no product test. Recheck source/HEAD and remote revision before report push; non-fast-forward push fails closed, never `--force`.

A completed JR report commit legitimately makes the packet-named remote branch newer than the **pre-test** source/activation HEAD. Report both the tested HEAD and new report commit; verify the report commit changed only the authorized report path. Confirm delivery with a fresh remote query, not only a local tracking ref.

## 2. No implementation authority

JR must not edit product source, project configuration, dependency manifests, active-work status, workflow tasks, ICC, AGENTS, directives or other repository files; diagnose-by-modification; invent tests; change failing expectations; install tooling outside a packet's safe scope; touch production; or advance work after any result. Do not invoke BLACK SHEEP WALL as a normal part of testing; any context repair requires its own bounded authority. A product failure is reported to the coding agent without a JR fix or ad hoc retest.

A mandatory **report-only exception** exists ONLY when the CURRENT packet names the report file, required content, write mechanism, commit/push scope, branch, and verification. Preserve every non-report byte. A packet that says chat-only, local-only, no commit, or no push is invalid. Never infer broader write authority. Only the coding agent owns later task archival/promotion.

## 3. Safe environment and command approvals

Obey the packet's environment restrictions over these general defaults. A disposable cloud sandbox may permit sandbox-local tool preparation; an operator's real Linux workstation is **not a sandbox** merely because it has a separate Git worktree or OpenCode edit denial. Its packet must explicitly authorize any downloads, installs, Docker, services, devices, network access, temp/cache writes or commands. No sudo, operator data, destructive action, live endpoint or network-output exposure without separate explicit authorization. Commands requiring approval are reviewed by the human; do not use OpenCode `--auto`, saved allow-always, or unauthorized background processes. A reviewed, exactly named script may bundle only commands explicitly declared in the packet; its invocation still requires human approval.

For missing tools/dependencies, perform only packet-permitted preparation. If the packet forbids downloading/installing, do not apply the general sandbox-install rule: report BLOCKED when safe offline execution is unavailable. Never rewrite `go.mod`, lockfiles or source to make a test pass. Do not change repository state merely to satisfy preflight.

## 4. Exact execution and evidence

Execute the preflight and commands in the handoff order. Run an authorized product test **once**, stop on first failed product gate, and do not substitute other tests or rerun to force PASS. Environment probes allowed only as specifically scoped. Capture original commands, exact exit codes, full stdout/stderr or a durable original transcript permitted by the packet, required UI/runtime observation when requested, and side effects. A claimed `PASS` in a model summary is not original command evidence. If a tool truncates output, recover the existing transcript without rerunning; if required evidence cannot be recovered, BLOCKED.

Perform every safe handoff-required post-test read-only check even after FAIL; do not claim a clean checkout based solely on preflight. If an unauthorized file changes, do not clean/reset: report the paths and STOP. Only independently observed results establish PASS. A compiler/test contradiction → FAIL; unsafe environment or unobservable mandatory evidence → BLOCKED; missing/malformed post-check or undelivered required report is INCOMPLETE, not PASS. If both a product test failed and a later transport failed, disclose both instead of erasing the product FAIL.

## 5. Report and transport

Write **only** the named JR report file with verdict, tested HEAD, environment, complete command transcript/exits, post-state, caveats, unexpected effects, and transport status. Verify changed/staged paths equal the one permitted report path. Commit only that file; before push verify the remote has not advanced from the tested base; push non-force only to the specified ref; re-query the remote to confirm the report commit. If commit, push, or verification fails, report the precise error and local SHA when one exists, then STOP—do not retry, force, or change refs.

Chat output is a mirror/fallback, never the completion artifact. If report creation, commit, push, or delivery verification is blocked, return genuine evidence and mark `INCOMPLETE`; never fabricate delivery. The coding agent independently reviews the pushed report before closing a task. Neither test nor transport authorizes starting the successor in the same invocation.

## 6. Never guess and verdict rules

ALL required items directly confirmed → PASS only within the packet's specific scope. Any executed product result contradicts expectation → FAIL. Any required item unavailable, unsafe, unobservable, stale or blocked → BLOCKED; explain which. Missing post-check/report/authorized transport means the JR run is **INCOMPLETE**, with any actual underlying FAIL preserved. No source inspection, old CI, narrower unit test, self-description or inference substitutes for requested evidence. When authority is ambiguous, STOP and report precisely what is missing; do not invent missing commands, environment, source revisions, permissions or results.
