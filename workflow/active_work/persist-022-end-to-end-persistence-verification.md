# PERSIST-022 — End-to-End Persistence Verification

Status: VERIFY PASS — independent JR on 2026-10-08
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

## Independent JR evidence

Verdict: **PASS**

Tested HEAD: `6175a77688a85e4aec4706c9e01079fa235f7963`  
Pinned product checkpoint: `c8f31b2443b69201f36464db72497ed8271aeb7d`

Executed once, in order:

1. `cd simulator && go test -mod=readonly -count=1 ./...` → PASS, exit 0.
2. `cd simulator && MCS_RUN_PERSIST_E2E=1 go test -mod=readonly -count=1 -run TestPersistenceDisposableEndToEnd -v .` → PASS, exit 0.

Observed lifecycle:
- first boot without snapshot stayed sealed and Modbus read was rejected with `0x06`;
- live value `0x1234` was written and persisted;
- restart began sealed and rejected Modbus while sealed;
- startup restore completed, verified all required areas, and final unseal committed last;
- runtime reported healthy/unsealed and Modbus returned restored `0x1234`;
- corrupted snapshot failed closed, remained sealed, reported a classified non-none failure, and Modbus remained rejected with `0x06`.

Post-check: working tree clean; changes after pinned checkpoint were workflow-only; no stray harness/MMA2 processes; disposable temp resources cleaned. Independent JR changed no source or workflow files; report was chat-only per pinned packet.

## Dependencies
Requires genuine completion of the already-promoted continuation chain:
`PERSIST-R01 → R02 → R03 → UI01 → R04 → R05`.

PERSIST-R05 must deliver a committed disposable real-runtime harness and record its exact command, target, expected observations and cleanup. Only then may the workflow owner/JR DEV pin this VERIFY packet exactly in `handoff.md` and switch execution identity to **OpenHands / independent JR**.

JR DEV must not execute or self-certify PERSIST-022.

## Sizing
0/2/0/2/2=6.

## CWAL
VERIFY complete — PASS. This packet is closed. No further persistence task is activated.
