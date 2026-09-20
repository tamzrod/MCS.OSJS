# UMIG-EM-002-R — TEST: current Replicator Go advanced-settings regression

Status: ACTIVE — human authorized iterative migration/test work 2026-09-20.
Stage / owner: TEST / OpenHands JR; evidence adjudication and advancement / ChatGPT.
Previous: UMIG-EM-002-B (archived COMPLETE/PASS for previous OS.js source).
Next: UMIG-EM-002-V (QUEUED; disposable live VERIFY, separate packet after this report).

## Primary outcome
Prove that the *current* Go Replicator at source checkpoint `c98eacef8a6af8b0786a09766a4a034e25e82ce6` passes its full repository-native package test suite, including newly committed advanced-settings, legacy-document persistence, capability and composition tests. This source was modified after the prior Go and Toolkit JR reports; those historical passes do not apply to it.

## Scope and acceptance

1. In JR's own disposable clean checkout, guarded-fast-forward to the precisely pinned source (only this task, predecessor archival, queued successor, and handoff documentation may differ from the pinned checkpoint); refuse unrelated changes or unsafe Git state.
2. With sandbox-local Go >= 1.25, independently run `(cd replicator && go test -race -count=1 -timeout=90s -v ./...)` once, report actual per-test results, evidence of advanced-settings tests executed, any warnings and exit code. No source/test/config edits or retries to force PASS.
3. Post-test tracked tree stays clean; return truthful PASS/FAIL/BLOCKED. A PASS certifies only Go unit/regression results, NOT live runtime, Electron UI, OS.js advanced UI, or deployment.

The exact preflight, invocation, expected evidence, report-write permission and failure stop gate are in the sole current `JR TEST TASK` in `handoff.md`. No other command or product test is authorized. No operator data, Docker, service startup, listener enablement, production, network exposure, installation, ICC or workflow edits. OpenHands updates only its authorized handoff report, commits/pushes only handoff if origin is unchanged and stops.

Sizing: implementation 0 / environment 1 / behavior 0 / verification 1 / recovery 1 = 3. No coding in this task.
