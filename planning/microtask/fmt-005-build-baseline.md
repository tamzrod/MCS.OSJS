# FMT-005 — Build baseline
Task ID: FMT-005
Task Name: Run one authorized non-destructive Toolkit build baseline
Blocker Task: FMT-001, FMT-003
Status: PENDING
Assigned Agent: OpenHands/JR — independent TEST runner
Stage: TEST

## Objective
Execute exactly one source-verified Toolkit build command and report its actual result, without repairing code.
## Scope
Read `ICC/INDEX.md`, accepted FMT-001/FMT-003 evidence, and exact package/build script they identify. Write only `planning/microtask/evidence/fmt-005-build-baseline.md` when promoted. No source, ICC, production, service or workflow edits beyond authorized task status. No installation, cleanup or generated tracked files unless separately approved.
## Execution
1. On promotion, human/JMGR must replace this gate with ONE exact command, cwd, environment, expected result and allowed disposable build output path derived from FMT-003; absent these, report BLOCKED and STOP.
2. Record HEAD, branch, git status and environment; execute only that command once.
3. Capture full stdout/stderr, exit code, post-run git status and any generated output paths; do not repair or reinterpret failure.
## Acceptance Criteria
1. Exact approved command and expected result are recorded before execution.
2. Raw exit/output and pre/post repository state are captured.
3. PASS only if observed exit and expected artifacts match; otherwise FAIL/BLOCKED as appropriate.
## Evidence
`planning/microtask/evidence/fmt-005-build-baseline.md`: source SHA, exact command/cwd/env, raw transcript, exit, expected/observed, status diff.
## Completion and delivery
COMPLETE means independent test report delivered with truthful PASS/FAIL, not necessarily a passing build; FAIL if test instructions cannot be fulfilled, BLOCKED if command/permission absent. JR does not fix code, promote successors or push without separately specified exact authority. STOP.
## Sizing
Implementation 0; environment 1; behavior 0; verification 1; decision/recovery 0 = 2/10. One build command only.