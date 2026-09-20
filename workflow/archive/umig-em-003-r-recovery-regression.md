# UMIG-EM-003-R — CODE: Split Replicator restart-failure regression

Status: ARCHIVED — COMPLETE / SOURCE ONLY on 2026-09-20. PRODUCT/TEST SOURCE checkpoint: `538324a472eb15ff8ef97ee66f826fa02ba9462f`.
Previous: `UMIG-EM-003-T` archived FAIL with original independent report `2d61394adc671c7e6cfe728a012b047eec7a0c7d`. Next: `UMIG-EM-003-R-T` now ACTIVE / independent TEST.

Human approved correction of the obsolete expectation in `replicator/multiblock_regression_test.go`. Source readback verified precisely one product test file changed from failing JR report HEAD; no runtime/product implementation changes. Old ambiguous test replaced by `TestApplyPreCommitFailureRestoresPreviousPollers` (unsupported advanced input fails locked composition before Commit; restore previous runtime, unchanged persisted document, no config/owners/restart request) and `TestApplyPostCommitRestartFailureStopsPreviousPollers` (config/owner commit succeeds, ACK fails; pollers stopped, no false document save, request retained). Shared temp fixture is isolated, unrelated validation regression tests preserved. Reviewed against `RuntimeManager.Apply`'s `committed` branch and approved fail-closed architecture.

SOURCE ONLY: Go compiler, tests and runtime for this newly authored correction are NOT verified by the coding agent. Sole successor TEST gets exact pinned one-command packet, not permission to edit/fix. No ICC edits or live VERIFY.
