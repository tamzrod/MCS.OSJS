# PERSIST-011 — Persistence Configuration Tests

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: TEST
Owner: OpenHands / independent JR
Previous: PERSIST-010
Next: PERSIST-012

## Primary outcome
Independently verify persistence configuration, sealing prerequisite, derived locked RBE synchronization, user-RBE compatibility and backward compatibility on the committed source checkpoint.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
No product fixes, no new tests, no runtime persistence verification.

## Acceptance
1. Run the exact repository-native targeted tests named in the current handoff packet and capture raw exits/output.
2. Required valid/invalid configuration cases and synchronization cases all pass.
3. Post-check confirms no unauthorized source/workflow changes by JR.

## Evidence / handoff
The current handoff must pin the exact source checkpoint, safe target, exact ordered commands/actions, expected observations, evidence and report permissions before execution. JR changes no product source and STOPS after its verdict.

## Dependencies
Requires genuine completion evidence for PERSIST-010. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
0/1/0/2/1=4; exact commands must be pinned in handoff at activation.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`; do not self-select this task while another ACTIVE assignment exists.
