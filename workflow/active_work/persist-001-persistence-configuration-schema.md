# PERSIST-001 — Persistence Configuration Schema

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: none
Next: PERSIST-002

## Primary outcome
Add `persistence.enabled` per MMA2 memory instance without adding persistence-owned port, unit, area, start or count fields.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement RBE derivation, filesystem persistence, restore, or runtime behavior.

## Acceptance
1. Persistence is represented per existing Port → Unit ID → Memory identity.
2. `persistence.enabled` round-trips without changing existing configurations.
3. No duplicate memory-range source of truth is introduced.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for no predecessor. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/1/1/1=4; tightly coupled schema/round-trip change.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
