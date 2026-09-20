# Handoff

## Current authority — 2026-09-20

The human approved a focused CODE correction of the single Replicator regression after independent OpenCode JR `UMIG-EM-003-T` FAIL. Full original 421-line JR evidence and prior exact packet are retained byte-for-byte in `workflow/archive/umig-em-003-t-failed-report.md` and in report commit `2d61394adc671c7e6cfe728a012b047eec7a0c7d`. The verified failure was `TestApplyRestartFailureRestoresPreviousPollers` (NOT `TestRuntimeManagerApplyLifecycleAndStatus`, which passed). MMA2 composer and Simulator suite PASS applied only to the earlier source checkpoint. No live VERIFY.

SOLE ACTIVE: `workflow/active_work/umig-em-003-r-recovery-regression.md` (CODE / ChatGPT); `workflow/active_work/umig-em-003-r-t-recovery-test.md` is QUEUED TEST. Former TEST `UMIG-EM-003-T` archived FAIL. Source-only test change in `replicator/multiblock_regression_test.go` distinguishes safe pre-commit poller restoration versus committed-but-unacknowledged post-commit fail-closed recovery. Product runtime behavior is unchanged; no compiler/test execution or correctness PASS is claimed for the new test source. Only BLACK SHEEP WALL changes ICC.

## JR TEST TASK — NONE / NOT AUTHORIZED

OpenCode is not currently authorized to run another packet while CODE remains ACTIVE. ChatGPT must read back the source-only correction and checkpoint, archive CODE, activate ONLY the explicitly queued focused TEST, write its exact fresh source-pinned runner packet, then inform the human to update the clean detached JR worktree and approve the exact safe command. No automatic OpenHands or live VERIFY. Do not rerun the old `UMIG-EM-003-T` runner against a different HEAD.

## Next action

ChatGPT source review and coding checkpoint. User separately requested evaluating a balanced local JR-Qwen reasoning setting; do not modify the local Legion OpenCode config without user running a narrowly scoped command and do not claim a measured quality improvement without a measured comparison.
