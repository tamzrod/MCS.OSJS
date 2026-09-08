# SIM-009 — Serve Simulator Values Through MMA2

Status: ACTIVE — current authorized CWAL task.

## Primary Outcome

After managed MMA2 activation succeeds, the Modbus Simulator runs its configured schedules, sends generated values through MMA2 raw ingest, and serves those values to real Modbus TCP clients from its ownership-reserved `(port, unit_id)` endpoint.

## Scope

- Connect successful managed MMA2 activation to the existing simulator scheduler and raw-ingest client.
- Start or replace the selected device's scheduler only after its MMA2 endpoint and raw-ingest path are ready.
- Restore enabled persisted simulator devices on appliance startup through the same ownership validation, MMA2 activation, and scheduler path used by Save & Apply.
- Stop or retain scheduler state consistently when MMA2 activation or raw ingest fails; do not continue claiming healthy service while writes cannot reach MMA2.
- Surface truthful Device, MMA2, Raw Ingest, and FC timing status through the existing runtime-status UI.
- Prove generated FC1-FC4 values are observable through real Modbus TCP reads at the configured addresses.

## Non-Scope

- No changes to ownership-key semantics or effective-config ownership policy.
- No direct writes to MMA2 memory; simulator values continue to use raw ingest exclusively.
- No manual register/coil editor, charts, scripting, or additional signal modes.
- No Replicator or Memory Appliance implementation.
- No unrelated OS.js redesign.

## Acceptance Criteria

1. Saving one valid enabled simulator device results in RUNNING MMA2, healthy raw ingest, advancing configured FC schedules, and a reachable Modbus TCP `(port, unit_id)` endpoint.
2. Real Modbus clients read values produced by the simulator for every configured FC area, with address ranges and data types matching the persisted definition.
3. Restarting the appliance runtime restores enabled persisted devices through the same ownership-safe path, while activation or ingest failure produces truthful ERROR/STOPPED status and does not falsely advance a healthy-serving state.

## Verification Method

Use one end-to-end serving workflow:

```text
SAVE ENABLED SIMULATOR DEVICE
-> OWNERSHIP VALIDATION + MANAGED MMA2 ACTIVATION
-> START FC SCHEDULES
-> RAW INGEST GENERATED VALUES
-> READ FC1-FC4 THROUGH REAL MODBUS TCP CLIENT
-> VERIFY VALUES + RANGES + STATUS
-> RESTART APPLIANCE RUNTIME
-> VERIFY AUTOMATIC RESTORE + MODBUS SERVING
-> BREAK ACTIVATION OR INGEST
-> VERIFY TRUTHFUL FAILURE STATE
```

Record the saved device definition, ownership entry, MMA2 listener, raw-ingest success, scheduler timestamps, exact Modbus client reads for each configured FC, restart/restore evidence, and failure-state evidence.

## Dependencies

- SIM-008 completed and verified.
- Completed SIM-003 scheduler, SIM-004 raw-ingest client, SIM-006 Save & Apply routing, and SIM-007 runtime status.
- Human promotion completed 2026-09-08.

## Sizing Assessment

- Implementation surface: 1 — bounded orchestration across existing simulator runtime components.
- Environment/dependency uncertainty: 0 — MMA2, scheduler, raw ingest, and Modbus verification paths are established.
- Behavioral surface: 1 — one serving behavior from enabled definition to externally readable values.
- Verification surface: 1 — one continuous end-to-end serving and restart workflow.
- Decision/recovery surface: 1 — startup restoration and failure-state coordination are bounded.
- Total: **4 / 10**. Acceptable as one task because all actions form one end-to-end serving path and use one verification workflow.
