# PERSIST-010 — Persistence Configuration UI

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: PERSIST-009
Next: PERSIST-011

## Primary outcome
Expose persistence enablement and locked derived RBE state in the existing memory Advanced Settings UI.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement disk/runtime status UI yet.

## Acceptance
1. User can enable/disable persistence only through supported draft/apply flow.
2. UI rejects persistence ON when State Sealing is disabled and explains why.
3. Derived persistence RBE entries display as locked/system-owned while user RBE remains editable.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-009. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/1=6; UI only.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
