# UMIG-EM-003-T — TEST: Shared writer lock (OpenCode JR)

Status: ARCHIVED — COMPLETE / FAIL, 2026-09-20. No TEST PASS, no promotion of UMIG-EM-003-V.
Source checkpoint tested: `949d7a8eb328dfa66f8359ffe81cfe359b8db2b7`; activated test HEAD: `1f5c30856d546274e9fea7e1c9a1ef43297fdbfc`.
Authoritative verbatim report: `handoff.md` at report commit `2d61394adc671c7e6cfe728a012b047eec7a0c7d`; original handoff snapshot additionally retained in `workflow/archive/umig-em-003-t-failed-report.md`.

Observed: `mma2composer` -race PASS; `simulator` -race PASS (opt-in live E2E intentionally SKIP); `replicator` -race FAIL (exit 1). Actual failing test `TestApplyRestartFailureRestoresPreviousPollers`, `multiblock_regression_test.go:73`; `TestRuntimeManagerApplyLifecycleAndStatus` PASS despite JR chat misidentification. Post-test status clean and HEAD unchanged; handoff-only report push confirmed.

Reason: old test expects previous pollers to resume after a committed MMA2 config fails restart ACK. Approved fail-closed contract prohibits stale pollers against potentially new config; preserve the underlying FAIL evidence, separate pre-commit restoration from post-commit recovery in coding task `UMIG-EM-003-R`. Only coding agent may advance. No live VERIFY or production readiness established.
