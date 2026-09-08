# SIM-001 — Establish Simulator Device Configuration Model

Status: DONE — completed+verified 2026-09-06 (CWAL).

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Completion evidence (recorded by CWAL 2026-09-06)

- Component boundary: repository-root `simulator/` (not `MMA2/`, not `OSJS/` packaged source).
- Host-mounted configuration root used: `OSJS_DATA_DIR` (verified by `deploy/docker-compose.yml` volume `osjs-data` -> `/data` and `OSJS/src/server/config.js`). Persist path: `$OSJS_DATA_DIR/config/simulator/devices.yaml`. `ConfigRootFromEnv` errors when `OSJS_DATA_DIR` is unset — no invented host path.
- Validation bounds taken from `MMA2/internal/config/validate.go` (port > 0, unit_id <= 255, start+count within 16-bit space, count 0 = unused area) plus brainstorm rule that a configured FC requires `Randomize Every (ms)` > 0.
- Verification command: `cd simulator && go test -count=1 -v ./...` (Go 1.22.2).
  - `TestRoundTripExactValues` PASS — Sim-PLC-1 name/enabled + MMA2 port 5020 unit 1 FC1-4 start/count + random-runtime 1000/5000/60000/10000 ms reloaded with exact struct equality.
  - `TestInvalidInputsDoNotReplaceLastValid` PASS — rejected port 0, unit_id 256, FC3 start+count overflow, FC1 interval 0, empty name; previous valid `devices.yaml` bytes unchanged.
  - `TestConfigRootFromEnvRefusesInventedPath` PASS.

## Primary Outcome

A simulator-owned persistent device definition can be created, validated, saved, and loaded while keeping MMA2 structural parameters separate from random-runtime parameters.

## Scope

- Define one simulator device definition containing `name` and `enabled`.
- Define the MMA2 parameter domain containing listener port, Unit ID, and FC1-FC4 Start/Count.
- Define the random-runtime parameter domain containing FC1-FC4 `Randomize Every (ms)`.
- Persist simulator-owned configuration under the repository-established host-mounted configuration root.
- Load the same persisted definition back into memory.
- Validate required field types and ranges before replacing the last valid persisted definition.
- Keep simulator-owned configuration separate from effective MMA2 runtime configuration.

## Non-Scope

- No MMA2 config activation or restart/reload.
- No shared MMA2 collision resolution.
- No random value generation.
- No raw-ingest writes.
- No OS.js simulator UI.
- No Replicator implementation.

## Acceptance Criteria

1. One valid simulator definition round-trips through save/load with exact values preserved in both parameter domains.
2. Persistent simulator configuration is written under the verified host-mounted configuration location, not an application or MMA2 source directory.
3. Representative invalid Port, Unit ID, FC Start/Count, and randomization interval inputs are rejected without replacing the previous valid definition.

## Verification

Create one valid definition, persist it, reload it, and compare all values exactly. Then attempt representative invalid updates and prove the previous valid persisted definition remains unchanged.

## Dependencies

- Approved Modbus Simulator brainstorm.
- Repository/runtime truth for the host-mounted configuration root. Do not invent a host path.

## Sizing

**3 / 10 — Small.** One configuration-model outcome, one persistence boundary, one validation workflow.
