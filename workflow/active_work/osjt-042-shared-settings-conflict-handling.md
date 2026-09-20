# OSJT-042 — Shared settings conflict handling

Status: QUEUED
Stage: VERIFY
Owner: OpenHands / independent JR under OPERATION CWAL
Previous: OSJT-041
Next: OSJT-043

## Primary outcome
Demonstrate revision-aware shared settings without output exposure.

## Bounded scope
Approved disposable OS.js Toolkit + isolated synthetic MMA2/runtime data only

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
Open two Toolkit windows, load the same shared revision, save an approved non-network setting in one, and verify the stale second save is rejected without mutation. Verify Close/Discard behavior and restore the original value. Keep RBE/access-event network outputs disabled unless separately approved. Use only the target and permissions recorded by OSJT-001. Record actual UI actions, replies, before/after values and cleanup evidence. Do not use production, public outputs, or the revoked runner.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-041; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 0/1/1/1/0 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

