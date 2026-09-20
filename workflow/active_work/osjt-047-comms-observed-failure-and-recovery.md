# OSJT-047 — COMMS observed failure and recovery

Status: QUEUED
Stage: VERIFY
Owner: OpenHands / independent JR under OPERATION CWAL
Previous: OSJT-046
Next: OSJT-048

## Primary outcome
Demonstrate LEDs reflect actual synthetic communication evidence.

## Bounded scope
Approved disposable OS.js Toolkit + isolated synthetic MMA2/runtime data only

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
On the isolated synthetic source, observe successful reads and writes, stop only that test source, verify stale green clears and layer errors are truthful, restart only that source and observe recovery. Record timestamps and respect the approved target permissions. Use only the target and permissions recorded by OSJT-001. Record actual UI actions, replies, before/after values and cleanup evidence. Do not use production, public outputs, or the revoked runner.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-046; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 0/1/1/1/0 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

