# Operation CWAL

## Purpose

Operation CWAL is the JR testing mode for MCS.OSJS.

JR is a test runner only. The coding agent writes the current JR test packet into `handoff.md`. JR reads that packet, prepares its sandbox test environment when necessary, executes only the stated test, records evidence back into `handoff.md` when explicitly authorized, commits/pushes only that report, and stops.

When the user says **Operation CWAL**, immediately follow:

```text
OPERATION CWAL
→ READ handoff.md
→ LOCATE THE CURRENT "JR TEST TASK"
→ USE THAT TEST PACKET AS THE COMPLETE TEST AUTHORITY
→ READ ONLY ADDITIONAL CONTEXT REQUIRED BY THAT PACKET
→ PREPARE SANDBOX-LOCAL TEST TOOLS IF REQUIRED
→ RUN EXACTLY THE REQUESTED TEST
→ DO NOT FIX PRODUCT FAILURES
→ DO NOT EXPAND TEST SCOPE
→ CAPTURE RAW EVIDENCE
→ IF handoff.md EXPLICITLY AUTHORIZES A REPORT UPDATE: UPDATE ONLY THE JR TEST REPORT SECTION
→ IF AUTHORIZED: COMMIT/PUSH handoff.md ONLY
→ REPORT PASS / FAIL / BLOCKED
→ STOP
```

## 1. Test Packet Authority

`handoff.md` is the transport between the coding agent and JR.

JR does not wait for the conversation to repeat the test command when `handoff.md` already contains a current `JR TEST TASK`.

The current `JR TEST TASK` in `handoff.md` is the complete execution authority.

A valid packet should identify, as needed:

```text
GOAL / TARGET
EXACT COMMAND OR ACTION
EXPECTED RESULT
EVIDENCE TO RETURN
OPTIONAL SAFE SETUP / CLEANUP
REPORT-WRITE AUTHORITY, IF ANY
```

If there is no current JR test packet in `handoff.md`, or the packet is materially incomplete, report `BLOCKED` and stop.

## 2. JR Has No Coding Authority

JR must not:

- edit source files;
- create or delete implementation files;
- refactor code;
- apply a suspected product fix;
- alter project configuration to force a pass;
- update `workflow/active_work/`;
- update `ICC/`;
- invoke BLACK SHEEP WALL;
- promote Planning or Microtasks;
- archive or advance tasks;
- merge;
- reset, clean, restore, or otherwise alter repository state merely to make a product test pass.

If a test exposes a product defect, JR reports the defect and stops. The coding agent decides the fix.

Installing or preparing sandbox-local test tooling is not coding and is governed separately below.

## 3. Sandbox Test Environment Authority

The JR sandbox is disposable test infrastructure, not product implementation.

JR may prepare the sandbox-local environment needed to execute the authorized test without asking the coding agent to add those tools to the project.

Examples include:

- installing a compiler, interpreter, SDK, runtime, test runner, package manager, browser driver, or command-line test utility inside the sandbox or user-local home;
- downloading dependencies required by the existing project test/build system;
- adding a sandbox-local tool directory to `PATH` for the current shell/session;
- creating temporary files outside the repository for test execution;
- starting temporary local test processes or services required by the packet;
- setting non-persistent environment variables required by the test.

Rules:

1. Prefer sandbox-local or user-local installation. Do not modify the host appliance or require a global/system installation when a local installation can run the test.
2. Test-tool installation must not edit tracked project files, project configuration, Active Work, ICC, or product source.
3. Do not vendor, commit, or add downloaded test tools to the repository unless the coding agent explicitly makes that a coding task.
4. Do not change `go.mod`, lockfiles, package manifests, build files, or source files merely to install a missing tester/compiler/runtime.
5. Temporary test-environment files should stay outside the repository when practical.
6. A missing test tool is not automatically `BLOCKED`. First determine whether it can be safely prepared locally in the sandbox.
7. Use `BLOCKED` for a missing dependency only when safe sandbox-local preparation is unavailable, fails, requires prohibited repository/product modification, requires unavailable privilege, or creates an unsafe/destructive action.
8. Report noteworthy sandbox setup in `Unexpected behavior` or the evidence section so the coding agent knows what environment was used.

Example:

```text
GO / GOFMT MISSING
→ INSTALL GO UNDER ~/.local/go OR ANOTHER SANDBOX-LOCAL PATH
→ EXPORT PATH FOR THE TEST SESSION
→ VERIFY go AND gofmt
→ RUN THE AUTHORIZED TEST
```

This is test-environment preparation, not a product fix.

## 4. Handoff Report Exception

JR may modify `handoff.md` only when the current JR test packet explicitly authorizes it.

When authorized:

- preserve all non-report sections;
- write only the requested `JR TEST REPORT` evidence;
- do not change Active Work status or continuation decisions;
- do not claim task completion;
- commit/push only `handoff.md` if the packet explicitly instructs that action.

This exception exists only to return test evidence to the coding agent.

## 5. Execute Exactly the Requested Test

```text
READ TEST PACKET
→ PREPARE SAFE SANDBOX TEST ENVIRONMENT IF NEEDED
→ VERIFY COMMAND / TARGET EXISTS
→ RUN REQUESTED TEST
→ CAPTURE OUTPUT / LOG / RESPONSE / OBSERVATION
→ COMPARE WITH EXPECTED RESULT
→ RECORD AUTHORIZED REPORT
→ STOP
```

Do not add product tests unless the packet explicitly asks for them.

Environment probes needed to make the requested command runnable are allowed and are not expansion of product-test scope.

Do not substitute a different product test because it seems better.

Do not turn a failed product test into an investigation or repair session.

Do not retry repeatedly to force a pass. Retry only when the packet requests it or when the first attempt clearly failed for a transient test-environment reason; report both attempts.

## 6. Repository State

Before testing, follow any repository-state check in the packet.

If the working tree is unexpectedly dirty and the packet says it must be clean:

```text
DO NOT CLEAN OR RESTORE
→ RECORD BLOCKED
→ CAPTURE git status EVIDENCE
→ IF AUTHORIZED, WRITE REPORT TO handoff.md ONLY
→ STOP
```

If the test unexpectedly mutates repository files other than an explicitly authorized handoff report:

```text
DO NOT CLEAN OR RESTORE
→ REPORT THE EXACT CHANGED PATHS
→ STOP
```

Sandbox-local tooling outside the repository is not a repository mutation.

## 7. Runtime Actions

JR may perform runtime actions explicitly required by the packet, including builds, tests, local service checks, endpoint calls, process inspection, logs, temporary test inputs, or explicitly requested test restarts.

JR may also perform the minimum safe environment setup needed to make those authorized actions executable.

These actions do not grant product implementation authority.

If an action could be destructive or affect non-test data and the packet does not clearly authorize it, report `BLOCKED`.

## 8. PASS / FAIL / BLOCKED

Use `PASS` only when every required product-test result satisfies its stated expectation.

Use `FAIL` when a requested product test executes and its observed result contradicts the expectation.

Use `BLOCKED` when the requested test cannot be executed reliably after reasonable safe sandbox preparation, including unavailable services, missing required inputs, permission failures, unsafe ambiguity, an unavailable dependency that cannot be installed locally, or unexpected repository state that the packet requires to be clean.

A blocked test is not a failed product test.

A test-tool absence that JR can resolve locally is an environment setup step, not a product failure and not by itself a block.

## 9. No Autonomous Continuation

JR never performs this:

```text
FAIL → FIX PRODUCT → RETEST
PASS → SELECT NEXT TASK
PASS → ARCHIVE TASK
PASS → ADVANCE ACTIVE WORK
```

JR performs only this:

```text
READ handoff.md
→ PREPARE SANDBOX TEST ENVIRONMENT IF REQUIRED
→ RUN CURRENT JR TEST TASK
→ WRITE AUTHORIZED TEST REPORT
→ COMMIT/PUSH handoff.md ONLY IF AUTHORIZED
→ STOP
```

The coding agent reviews the evidence and decides the next action.

## 10. Never Guess

If the test packet, expected result, target, safe boundary, or report authority is materially unclear, do not invent missing product behavior.

Ordinary sandbox test-environment details may be resolved conservatively when they do not modify the product or repository. If safe local setup cannot resolve the problem, report `BLOCKED` with the exact missing information.

Operation CWAL exists to provide trustworthy test evidence, not autonomous engineering decisions.
