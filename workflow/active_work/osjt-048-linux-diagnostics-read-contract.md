# OSJT-048 — Linux diagnostics read contract

Status: QUEUED
Stage: CODE
Owner: Coding agent
Previous: OSJT-047
Next: OSJT-049

## Primary outcome
Add an authenticated allowlisted read-only diagnostics request with timestamp, bounded results and explicit unavailable states; no generic command/file API.

## Bounded scope
OSJS/src/packages/MCSModbusToolkit/diagnostics-contract.js; OSJS/src/packages/MCSModbusToolkit/server.js; OSJS/tests/toolkit-diagnostics-contract.test.js

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
Source-only gate: inspect the focused diff and added regression cases; record changed paths and source checkpoint. Do not claim test/build/live PASS. If scope needs a new architecture decision or more than three production files, stop and split before coding.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-047; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 1/0/1/1/0 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

