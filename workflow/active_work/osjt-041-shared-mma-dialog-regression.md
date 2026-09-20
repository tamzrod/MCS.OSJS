# OSJT-041 — Shared MMA dialog regression

Status: QUEUED
Stage: TEST
Owner: OpenHands / independent JR under OPERATION CWAL
Previous: OSJT-040
Next: OSJT-042

## Primary outcome
Add a separate shared draft dialog reachable from either Advanced Settings editor; Close retains draft, Discard restores it, Save sends revision and conflict asks for reload; no automatic output enabling. (independent deterministic verification only).

## Bounded scope
OSJS/tests/toolkit-shared-dialog.test.js

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
From OSJS: node --test tests/toolkit-shared-dialog.test.js. The file must contain nonzero assertions/cases for the predecessor behavior. Then record full output and exit 0; no source edits.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-040; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 1/0/1/1/0 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

