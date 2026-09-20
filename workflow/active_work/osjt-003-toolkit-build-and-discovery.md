# OSJT-003 — Toolkit build and discovery

Status: QUEUED
Stage: TEST
Owner: OpenHands / independent JR under OPERATION CWAL
Previous: OSJT-002
Next: OSJT-004

## Primary outcome
Establish that the unified Toolkit is built and discoverable.

## Bounded scope
OSJS/package.json; OSJS/scripts/build-local-packages.js; OSJS/dist/

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
From OSJS with the repository-compatible provisioned Node environment: npm run build:local-packages && npm run package:discover && npm run build. Confirm MCSModbusToolkit discovery and its emitted bundle; preserve full output and artifact paths. Missing dependencies are BLOCKED, not permission to install.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-002; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 0/1/0/1/1 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

