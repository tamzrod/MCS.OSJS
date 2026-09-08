# SIM-017 — End-to-End Simulator/MMA2 Verification

Status: COMPLETED 2026-09-08.

## Primary outcome

Prove the corrected Simulator/MMA2 architecture end to end with real Modbus reads and explicit negative-boundary checks.

## Scope

Verify the completed sequence:

`BOOT -> MMA2 auto-start -> Simulator restore -> edit simulation -> Save & Apply -> Simulator-owned shared MMA2 entry update -> MMA2 RESTART only -> MMA2 ready -> schedules running -> Raw Ingest -> real Modbus FC1/FC2/FC3/FC4 reads -> reboot -> automatic resume`.

Also verify architectural negatives: no simbridge, no Simulator config API, no Simulator-started MMA2, no Simulator STOP operation, no foreign shared-config damage, and no restart after rejected config.

## Non-scope

- No new architecture or feature behavior.
- No protocol expansion.
- No new MMA2 lifecycle commands.

## Acceptance criteria

1. MMA2 auto-starts independently on boot and serves the persisted shared config.
2. Save & Apply safely updates only Simulator-owned MMA2 configuration and requests RESTART.
3. Scheduler/Raw Ingest produces changing configured values.
4. A real Modbus TCP client reads expected values from FC1, FC2, FC3, and FC4.
5. Full reboot restores MMA2 and enabled simulation without manual apply.
6. All negative architecture boundaries listed in Scope are proven.

## Verification

Use a real enabled Simulator definition with FC1-FC4 ranges and a real Modbus TCP client. Record listener/read results, scheduler/raw-ingest status, restart behavior, shared-config preservation, and reboot restoration.

## Completion evidence

- Added verification-only `simulator/e2e_test.go`; the harness builds and independently starts MMA2, while Simulator code remains limited to composing config and emitting RESTART requests.
- Boot restore armed an enabled four-area device after independently started MMA2 became ready. Scheduler Raw Ingest changed FC3 register values, and a real Modbus TCP client successfully read FC1, FC2, FC3, and FC4 at their configured ranges.
- A structural port edit emitted exactly one restart request. The independent harness restarted MMA2, the new listener became ready, and reads succeeded through the reloaded config.
- A duplicate-reservation edit was rejected with shared config byte-unchanged and no restart request. A foreign-owned listener/reservation remained available before and after apply and reboot.
- A full stop/start of MMA2 plus reconstruction of the Simulator router restored the enabled schedule without manual apply and without creating a restart request.
- `MCS_RUN_E2E=1 go test -race -count=1 -run TestEndToEndSimulatorMMA2Architecture -v` passes. Simulator `gofmt -l .`, `go vet ./...`, and `go test -count=1 ./...` pass. Static searches find no simbridge/config API or forbidden MMA2 process-control path.

## Dependencies

SIM-016.

## Sizing

Implementation 0, environment 1, behavioral 1, verification 2, decision/recovery 0 = 4. Verification-only capstone with one architecture-level end-to-end proof.
