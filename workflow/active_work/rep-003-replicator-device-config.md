# REP-003 — Replicator Device Configuration

Status: QUEUED
Previous: REP-002
Next: REP-004

## Primary Outcome

Add a minimal persisted Replicator configuration model describing one external Modbus source and one MMA2 destination reservation, using the same storage conventions and validation style already proven by the Simulator where applicable.

## Scope

- Add Replicator-owned persisted configuration under `$OSJS_DATA_DIR/config/replicator/`.
- Define the minimum source fields needed for the first test: source host, source port, source unit ID, function/area, start address, count, and polling interval.
- Define destination MMA2 fields needed for the first test: listener port, unit ID, destination area/start/count.
- Validate obvious invalid/zero/out-of-range values before persistence.
- Keep the first configuration shape limited to one simple 1:1 register range.

## Non-Scope

- No source Modbus connection yet.
- No polling loop.
- No data transformation/scaling.
- No multiple-source or multiple-range support.
- No OS.js UI.

## Acceptance Criteria

1. A valid minimal Replicator definition can be saved and loaded from the Replicator config path.
2. Invalid required source/destination parameters are rejected before persistence.
3. The persisted model contains enough information for later tasks to read a known Simulator register range and target a distinct Replicator-owned MMA2 reservation.

## Verification

Focused config validation and save/load round-trip tests using a temporary `OSJS_DATA_DIR`.

## Dependencies

- Repository storage conventions proven by the Simulator configuration work.
- REP-001 only if destination reservation types are intentionally shared at this layer; otherwise no runtime dependency.

## Sizing

Implementation surface 1; environment uncertainty 0; behavioral surface 1; verification surface 1; decision/recovery surface 0. Total: 3 — good JR task.
