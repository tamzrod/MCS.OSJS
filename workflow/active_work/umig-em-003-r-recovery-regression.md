# UMIG-EM-003-R — CODE: Split Replicator restart-failure regression

Status: ACTIVE — human approved focused correction 2026-09-20; source-only authored checkpoint, awaiting coding-agent readback and archiving. Exactly one active task.
Stage / owner: CODE / ChatGPT. Previous: UMIG-EM-003-T (archived FAIL). Next: UMIG-EM-003-R-T (QUEUED independent OpenCode JR TEST).

Goal: Resolve the contradictory legacy test `TestApplyRestartFailureRestoresPreviousPollers`, preserving the approved contract: pre-commit composition failures restore previous pollers; once MMA2 configuration/ownership is committed, missing restart ACK requires recovery and old pollers MUST remain stopped. No automatic rollback, product behavior weakening or unapproved runtime operations.

Source-only change: `replicator/multiblock_regression_test.go` replaces the ambiguous test with `TestApplyPreCommitFailureRestoresPreviousPollers` (invalid advanced input survives advisory preflight, fails inside locked composer before Commit; original poller resumes; no config/owners/restart artifact) and `TestApplyPostCommitRestartFailureStopsPreviousPollers` (valid edited scan-rate, config/owners committed, ACK missing, no stale poller or falsely saved document, retained restart request). Reuse a local temp fixture/helper. Retain original unrelated validation tests. No production code change.

CODE gate: inspect exact diff/source, confirm tests reflect documented `RuntimeManager.Apply` commit boundary and no other product edits, checkpoint SOURCE ONLY (not Go TEST PASS). Never edit ICC. Once checked, archive this CODE and activate only the named QUEUED `UMIG-EM-003-R-T` with a fresh source-pinned exact OpenCode packet; no automatic VERIFY promotion.
