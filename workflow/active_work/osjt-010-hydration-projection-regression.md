# OSJT-010 — Hydration projection regression

Status: ACTIVE
Stage: TEST
Owner: OpenHands / independent JR under OPERATION CWAL
Previous: OSJT-009
Next: OSJT-011

## Primary outcome
Project omitted policy, sealing and RBE from the matching effective port/unit memory; preserve explicit values and unknown extensions; malformed effective configuration returns an error. (independent deterministic verification only).

## Bounded scope
simulator/advanced_projection_test.go

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
From simulator: go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionMerge' . Capture full output showing matching tests actually ran and passed; zero matching tests is not PASS. No source edits.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-009; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 1/0/1/1/0 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.
