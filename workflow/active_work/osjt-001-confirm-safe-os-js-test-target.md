# OSJT-001 — Confirm safe OS.js test target

Status: ACTIVE
Stage: PREP
Owner: Coding agent
Previous: NONE
Next: OSJT-002

## Primary outcome
Record one explicitly approved disposable Linux target, available Node/Go/dependency versions, source revision and isolated data/network boundaries. Reuse is allowed only after checking prior resources; never reuse or rerun the cancelled EM-003 packet.

## Bounded scope
docs/osjs-toolkit/test-target.md; OSJS/package.json; deploy/verify/compose.yaml

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
Read source manifests, git status and existing environment evidence. Record target identity, permission boundary and exact commands for later tasks. No Docker startup, downloads, service changes or operator-data writes. If no approved target exists, record BLOCKED and request only that missing decision.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on human authorization for initial target discovery; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 0/1/0/1/1 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

