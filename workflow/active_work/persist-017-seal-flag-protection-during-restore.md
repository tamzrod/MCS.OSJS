# PERSIST-017 — Seal-Flag Protection During Restore

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-016
Next: PERSIST-018

## Primary outcome
Prevent persisted sealing-flag value from causing premature unseal during snapshot restoration.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add a second sealing source of truth or runtime reseal semantics.

## Acceptance
1. The configured State Sealing address remains sealed/zero throughout area restoration.
2. A snapshot containing prior flag=1 cannot unseal during restore.
3. Behavior follows the authoritative State Sealing area/address configuration; no duplicate flag setting is added.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-016. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/2=6.

## CWAL
PERSIST-016 is delivered on GitHub main at `2fbabc3481b6435ca38596ea7dd060e36ed0655d`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, preserve the authoritative State Sealing area/address as the only sealing source of truth, deliver under standing JR DEV authority, prepare PERSIST-018 for the next invocation only after genuine completion, and STOP.
