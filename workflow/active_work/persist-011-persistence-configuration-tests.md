# PERSIST-011 — Persistence Configuration Tests

Status: ACTIVE — independent JR TEST/VERIFY; exact packet prepared in `handoff.md`
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
PERSIST-010 is delivered on GitHub main at `190e464f6106c3e640d21b1f9352328447464bef`. This packet is the sole current ACTIVE assignment in independent JR TEST/VERIFY mode. Execute only the exact read-only verification packet pinned in `handoff.md`, make no product fixes or new tests, report PASS/FAIL/BLOCKED/INCOMPLETE with raw evidence, and STOP.
