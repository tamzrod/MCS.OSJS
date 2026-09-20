# OSJT-009 — Hydration projection

Status: QUEUED
Stage: CODE
Owner: Coding agent
Previous: OSJT-008
Next: OSJT-010

## Primary outcome
Project omitted policy, sealing and RBE from the matching effective port/unit memory; preserve explicit values and unknown extensions; malformed effective configuration returns an error.

## Bounded scope
simulator/advanced_projection.go; simulator/advanced_projection_test.go; simulator/osjs_toolkit_settings_test.go

Read this task and direct predecessor evidence, not the entire backlog. No Electron changes, legacy removal, production/operator data, dependency upgrades or unrelated cleanup. Do not enable network outputs or reuse the cancelled EM-003 packet.

## Acceptance / exact check
Source-only gate: inspect the focused diff and added regression cases; record changed paths and source checkpoint. Do not claim test/build/live PASS. If scope needs a new architecture decision or more than three production files, stop and split before coding.

## Evidence and advancement
Record the source revision, changed paths (CODE only), exact checks and original outcomes. CODE source inspection is not TEST PASS; mocked tests are not live VERIFY. Stop on a failed check or missing prerequisite and report the smallest needed repair; do not expand the task.

Before TEST/VERIFY activation, the coding agent supplies one current handoff packet with predecessor source SHA, approved target, this check, post-check and chat-only report authority. JR does not invent a packet, modify source, or push reports without separate authorization. QUEUED is not runnable. Advance only after the task gate is satisfied and independently reviewed where required, through matching Previous/Next links.

## Dependencies and size
Depends on OSJT-008; target/permissions derive from OSJT-001.
Five dimensions (implementation/environment/behavior/verification/decision): 1/0/1/1/0 = 3.
At most three production files plus focused tests per coding task. Split before editing if the task requires a broader change or a new architectural decision.

