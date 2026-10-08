# PERSIST-R05 — Disposable Persistence E2E Harness

Status: QUEUED — HUMAN PROMOTED 2026-10-08
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
Already human-promoted. Execute when selected by handoff. After delivery, pin and activate PERSIST-022 exactly for independent JR, then STOP because VERIFY requires a new independent-JR identity; do not self-certify.
