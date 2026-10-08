# PERSIST-R05 — Disposable Persistence E2E Harness

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); pinning PERSIST-022 independent-JR VERIFY packet
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-R04
Next: PERSIST-022

## Primary outcome
Add one committed repository-native disposable integration harness that PERSIST-022 independent JR can execute without modifying product source.

## Scope
Create a deterministic disposable integration test/runner using loopback/temp data only that proves the real runtime sequence:
`known value → RBE save → restart sealed → Modbus rejected while sealed → restore → verify → final unseal → Modbus reads restored value`.
Also include one deterministic failed-restore case that remains sealed.

## Non-scope
No production/customer data, no global services, no product behavior changes disguised as a test, no source repair by independent JR, no ICC edits.

## Acceptance
1. One committed exact command exercises the successful real-runtime lifecycle using disposable ports/data and asserts restored value through Modbus after final unseal.
2. The harness observes sealed Modbus rejection before commit and a failure case that remains sealed.
3. The harness cleans only its own disposable resources and provides stable evidence suitable for a one-run independent PERSIST-022 packet.

## Evidence / handoff
Record harness path, exact command, expected observations, cleanup behavior, bounded regression and delivered source SHA. After genuine delivery, prepare and activate the exact PERSIST-022 independent-JR VERIFY packet. This is the recursion boundary: JR DEV must STOP before executing or certifying PERSIST-022.

## Dependencies
Requires genuine PERSIST-R04 delivery and the full runtime path from R01-R03.

## Sizing
1/1/1/2/0=5.

## CWAL
PERSIST-R05 CODE is complete and no longer ACTIVE. This is the recursion boundary: PERSIST-022 is pinned below for independent JR and JR DEV STOPS here.

## Coding evidence (PERSIST-R05, OpenHands JR DEV)

- Source checkpoint base: `43c266229e4fd4089e43a89dd8cc078f0bd5ab98` (`main`, clean, = `origin/main` at run).
- Changed path (product):
  - A `simulator/persistence_e2e_test.go` (blob `7c833aa20fe3f6ee6005b4f6633b5d4cecc3b8e9`) — `TestPersistenceDisposableEndToEnd`, the committed disposable harness. It builds a real MMA2 from `../MMA2/cmd/mma2` into a `t.TempDir()` root and drives the real lifecycle over disposable loopback ports only. Gated behind `MCS_RUN_PERSIST_E2E` so ordinary `go test ./...` does not build MMA2.
- Exact command (from the repository root):
  `cd simulator && MCS_RUN_PERSIST_E2E=1 go test -mod=readonly -count=1 -run TestPersistenceDisposableEndToEnd -v .`
- Expected observations (all asserted by the harness):
  1. Phase A (live): first boot with no snapshot stays sealed; Modbus read is rejected `0x06`; writing the known value `0x1234` and unsealing (Raw Ingest) makes a live Modbus read return `0x1234`; a real RBE save event persists a changed snapshot and is observed as `LastSave`.
  2. Phase B (restart): fresh MMA2 is sealed before restore; building the runtime runs the real startup restore → verify → final unseal; `Healthy` is true, `Sealed` false; Modbus reads the restored `0x1234`.
  3. Phase C (fail-closed): a corrupted snapshot makes restore fail closed; `Healthy` false, `Sealed` true, a non-`none` classified reason, and Modbus stays rejected `0x06`.
- Cleanup: the harness uses only `t.TempDir()` roots and disposable processes; `t.Cleanup`/deferred stops and MMA2 interrupts clean up only its own resources. No production/customer data, no global services, no product behavior change.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0; the harness is a gated `SKIP` under ordinary `go test ./...`.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is PERSIST-022.
- Delivered source commit: `c8f31b2443b69201f36464db72497ed8271aeb7d` on GitHub main.
