# SIM-017 — End-to-End Simulator/MMA2 Verification

Status: ACTIVE — promoted for sequential execution after SIM-016.

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

## Dependencies

SIM-016.

## Sizing

Implementation 0, environment 1, behavioral 1, verification 2, decision/recovery 0 = 4. Verification-only capstone with one architecture-level end-to-end proof.