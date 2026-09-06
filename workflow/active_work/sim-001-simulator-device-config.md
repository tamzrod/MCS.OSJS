# SIM-001 — Establish Simulator Device Configuration Model

Status: ACTIVE — human-promoted for JR execution.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

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