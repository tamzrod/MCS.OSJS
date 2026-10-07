# PERSIST-022 — End-to-End Persistence Verification

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: VERIFY
Owner: OpenHands / independent JR
Previous: PERSIST-021
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
Requires genuine completion evidence for PERSIST-021 and a complete current VERIFY packet in `handoff.md`.

## Sizing
0/2/0/2/2=6.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`.
