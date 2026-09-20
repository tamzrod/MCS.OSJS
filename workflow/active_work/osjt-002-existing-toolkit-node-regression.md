# OSJT-002 — Existing Toolkit Node regression

Status: QUEUED
Stage: TEST
Owner: OpenHands / independent JR under OPERATION CWAL
Previous: OSJT-001
Next: OSJT-003

## Primary outcome
Establish current basic Toolkit deterministic regression status on Linux.

## Bounded scope
OSJS/tests/*.test.js

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
From OSJS on the approved target: node --test tests/*.test.js. Capture complete output, Node version and exit code. Every test must pass; do not replace failed Unix-socket tests with Windows substitutes.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-001; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 0/1/0/1/0 = 2.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

