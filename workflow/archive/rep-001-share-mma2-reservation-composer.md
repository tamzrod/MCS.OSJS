# REP-001 — Share MMA2 Reservation Composer

Status: COMPLETED 2026-09-10 — shared composer and Simulator migration verified.
Previous: none
Next: REP-002

## Primary Outcome

Move/generalize the already-proven Simulator MMA2 effective-config and `(port, unit_id)` ownership composition logic into a reusable internal package so both Simulator and Replicator can use the same implementation without duplicating ownership rules.

## Scope

- Extract only the producer-neutral MMA2 config/ownership mechanics currently proven in `simulator/mma2_config.go`.
- Preserve the existing effective config and ownership paths under `$OSJS_DATA_DIR/config/mma2/`.
- Parameterize producer identity so Simulator continues to use `simulator` and Replicator can later use `replicator`.
- Update Simulator to call the shared implementation with no intended behavior change.
- Preserve atomic candidate validation/write and foreign-owner collision rejection.

## Non-Scope

- No Replicator source Modbus connection.
- No Replicator polling/runtime.
- No raw-ingest extraction.
- No OS.js UI changes.
- No new MMA2 protocol or lifecycle behavior.

## Acceptance Criteria

1. Simulator uses the shared composer rather than a private duplicate of the ownership/composition implementation.
2. Existing `(port, unit_id)` collision semantics and effective MMA2 YAML behavior remain unchanged for Simulator.
3. The shared API accepts an explicit producer identity suitable for a later `replicator` caller.

## Verification

Run the existing Simulator MMA2 compose/ownership tests plus focused tests for producer identity and foreign-owner rejection in the shared package. All must pass without changing expected Simulator behavior.

## Dependencies

- Existing completed Simulator MMA2 ownership/composition implementation.
- `MMA2/pkg/configvalidate` validation behavior.

## Sizing

Implementation surface 1; environment uncertainty 0; behavioral surface 1; verification surface 1; decision/recovery surface 0. Total: 3 — good JR task.

## Completion Evidence

- `mma2composer` owns the producer-neutral effective-config, ownership, collision, atomic write, validation, and rollback implementation with explicit producer identity.
- Simulator now delegates composition to `mma2composer.New(s.Root, "simulator")`; its prior private implementation was removed while compatibility aliases and Store path helpers preserve existing callers and tests.
- Focused shared-package tests cover producer identity, foreign-owner rejection, first-come-first-save persistence, requested-port deletion, and byte-for-byte config restoration when the owners replace fails.
- `GOCACHE=/tmp/mcs-osjs-rep001-go-cache go test -count=1 ./...` and `go vet ./...` passed in `mma2composer/`.
- The same full test and vet gates passed in `simulator/`; the test suite reported `ok github.com/tamzrod/MCS.OSJS/simulator 1.400s`.
