# PERSIST-004 — Locked System RBE Behavior

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: PERSIST-003
Next: PERSIST-005

## Primary outcome
Represent persistence RBE rules as system-owned and non-editable/non-deletable through supported configuration/UI paths.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add new persistence runtime behavior.

## Acceptance
1. System persistence RBE is visibly identified as system-owned/locked.
2. User edit/delete attempts cannot mutate the system-derived range or ownership.
3. Ordinary user RBE behavior is unchanged.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-003. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/1=6; ownership/lock behavior only.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
