# PERSIST-022 — End-to-End Persistence Verification

Status: QUEUED — HUMAN PROMOTED; dependency-gated until PERSIST-R05 completes
Stage: VERIFY
Owner: OpenHands / independent JR
Previous: PERSIST-R05
Next: none

## Primary outcome
Independently verify the complete safe persistence lifecycle on a disposable configuration: save state → restart sealed → restore → verify → unseal → Modbus sees restored state.

## Scope
The current handoff must pin the exact source checkpoint, disposable target, preflight, exact ordered actions/commands, expected observations, evidence destination and report permissions before this packet can be selected.

## Non-scope
No production/customer data, no source fixes, no invented commands, no global service actions.

## Acceptance
1. Execute only the exact safe disposable runtime/actions pinned in the current handoff and capture raw evidence.
2. Observe that Modbus access remains sealed until restoration completes and restored values are correct after unseal.
3. Exercise at least one authorized failure case showing failed restore remains sealed, if included in the promoted exact packet.

## Evidence / handoff
JR changes no product source, follows `operation cwal.md`, captures genuine evidence and STOPS after verdict.

## Dependencies
Requires genuine completion of the already-promoted continuation chain:
`PERSIST-R01 → R02 → R03 → UI01 → R04 → R05`.

PERSIST-R05 must deliver a committed disposable real-runtime harness and record its exact command, target, expected observations and cleanup. Only then may the workflow owner/JR DEV pin this VERIFY packet exactly in `handoff.md` and switch execution identity to **OpenHands / independent JR**.

JR DEV must not execute or self-certify PERSIST-022.

## Sizing
0/2/0/2/2=6.

## CWAL
Human promotion is already granted, but execution is dependency-gated. Do not run this packet until genuine PERSIST-R05 delivery has pinned the exact safe independent-JR execution packet. After R05, JR DEV may activate this packet for the next invocation but must STOP; the next invocation runs under independent JR identity.
