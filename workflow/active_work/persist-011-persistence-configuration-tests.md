# PERSIST-011 — Persistence Configuration Tests

Status: TEST PASS — independent JR 2026-10-08
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

## Independent JR result

Verdict: **PASS**.

Tested HEAD: `313e3be60a29fff6c967d78c0cf677b8b888ed56`
Pinned product checkpoint: `190e464f6106c3e640d21b1f9352328447464bef`

Exact ordered commands all exited 0:
1. `cd mma2composer && go test -mod=readonly ./...`
2. `cd simulator && go test -mod=readonly ./...`
3. `cd replicator && go test -mod=readonly ./...`
4. `cd OSJS && node tests/toolkit-persistence-ui.test.js`
5. `cd OSJS && node tests/toolkit-ui-parity.test.js`
6. `cd OSJS && node tests/toolkit-fc43.test.js`

Freshness gate passed; commits after the pinned product checkpoint were workflow-only. Post-check clean; no product/test file modified by JR. No reruns or substitutions.

## CWAL

Independent TEST/VERIFY complete with PASS. This packet is no longer ACTIVE. PERSIST-012 is the next eligible CODE packet.
