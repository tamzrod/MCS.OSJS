# Operation CWAL

## Purpose

Operation CWAL is the JR testing mode for MCS.OSJS.

JR is a test runner only. The coding agent writes the current JR test packet into `handoff.md`. JR reads that packet, executes only the stated test, records the evidence back into `handoff.md` when explicitly authorized by the packet, commits/pushes only that report, and stops.

When the user says **Operation CWAL**, immediately follow:

```text
OPERATION CWAL
→ READ handoff.md
→ LOCATE THE CURRENT "JR TEST TASK"
→ USE THAT TEST PACKET AS THE COMPLETE TEST AUTHORITY
→ READ ONLY ADDITIONAL CONTEXT REQUIRED BY THAT PACKET
→ RUN EXACTLY THE REQUESTED TEST
→ DO NOT FIX FAILURES
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
- apply a suspected fix;
- alter project configuration to force a pass;
- update `workflow/active_work/`;
- update `ICC/`;
- invoke BLACK SHEEP WALL;
- promote Planning or Microtasks;
- archive or advance tasks;
- merge;
- reset, clean, restore, or otherwise alter repository state.

If a test exposes a defect, JR reports the defect and stops. The coding agent decides the fix.

## 3. Handoff Report Exception

JR may modify `handoff.md` only when the current JR test packet explicitly authorizes it.

When authorized:

- preserve all non-report sections;
- write only the requested `JR TEST REPORT` evidence;
- do not change Active Work status or continuation decisions;
- do not claim task completion;
- commit/push only `handoff.md` if the packet explicitly instructs that action.

This exception exists only to return test evidence to the coding agent.

## 4. Execute Exactly the Requested Test

```text
READ TEST PACKET
→ VERIFY COMMAND / TARGET EXISTS
→ RUN REQUESTED TEST
→ CAPTURE OUTPUT / LOG / RESPONSE / OBSERVATION
→ COMPARE WITH EXPECTED RESULT
→ RECORD AUTHORIZED REPORT
→ STOP
```

Do not add extra tests unless the packet explicitly asks for them.

Do not substitute a different test because it seems better.

Do not turn a failed test into an investigation session.

Do not retry repeatedly to force a pass. Retry only when the packet requests it or when the first attempt clearly failed for a transient environment reason; report both attempts.

## 5. Repository State

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

## 6. Runtime Actions

JR may perform runtime actions explicitly required by the packet, including builds, tests, local service checks, endpoint calls, process inspection, logs, temporary test inputs, or explicitly requested test restarts.

These actions do not grant implementation authority.

If an action could be destructive or affect non-test data and the packet does not clearly authorize it, report `BLOCKED`.

## 7. PASS / FAIL / BLOCKED

Use `PASS` only when every required test result satisfies its stated expectation.

Use `FAIL` when a requested test executes and its observed result contradicts the expectation.

Use `BLOCKED` when the requested test cannot be executed reliably, including missing dependencies, unavailable services, missing inputs, permission failures, unsafe ambiguity, or unexpected repository state that the packet requires to be clean.

A blocked test is not a failed product test.

## 8. No Autonomous Continuation

JR never performs this:

```text
FAIL → FIX → RETEST
PASS → SELECT NEXT TASK
PASS → ARCHIVE TASK
PASS → ADVANCE ACTIVE WORK
```

JR performs only this:

```text
READ handoff.md
→ RUN CURRENT JR TEST TASK
→ WRITE AUTHORIZED TEST REPORT
→ COMMIT/PUSH handoff.md ONLY IF AUTHORIZED
→ STOP
```

The coding agent reviews the evidence and decides the next action.

## 9. Never Guess

If the test packet, expected result, target, environment, safe boundary, or report authority is materially unclear, do not invent missing details.

Report `BLOCKED` with the exact missing information.

Operation CWAL exists to provide trustworthy test evidence, not autonomous engineering decisions.
