# OSJT-033 — Shared MMA revision guard regression

Status: QUEUED
Stage: TEST
Owner: OpenHands / independent JR under OPERATION CWAL
Previous: OSJT-032
Next: OSJT-034

## Primary outcome
Require expected_revision and compare it under the shared lock; stale/missing revision returns a typed error with zero writes; prepare the validated candidate without releasing the transaction lock. (independent deterministic verification only).

## Bounded scope
simulator/osjs_toolkit_settings_test.go

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
From simulator: go test -count=1 -timeout=90s -v -run '^TestSharedSettingsRevision' . Capture full output showing matching tests actually ran and passed; zero matching tests is not PASS. No source edits.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-032; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 1/0/1/1/0 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

