# SIM-006 — Route Save & Apply to the Correct Parameter Consumer

Status: COMPLETED — verified 2026-09-08.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Primary Outcome

`Save & Apply` classifies changed simulator parameters and sends MMA2 structural changes and random-runtime timing changes to their correct execution paths.

## Scope

- Validate the edited simulator definition before activation.
- Compare edited values with the currently active valid definition.
- Classify Port, Unit ID, FC Start, and FC Count changes as MMA2 structural changes.
- Classify FC `Randomize Every (ms)` changes as random-runtime changes.
- Route MMA2 structural changes through SIM-002.
- Route timing-only changes through SIM-003.
- Preserve the previously active valid state when validation or downstream activation fails.
- Surface the resulting success/error state back to the UI.

## Non-Scope

- No new MMA2 configuration semantics.
- No new scheduler behavior.
- No raw-ingest implementation.
- No additional simulator UI features.

## Acceptance Criteria

1. One structural edit follows the MMA2 activation path and triggers restart/reload only when that path requires it.
2. One timing-only edit updates the random scheduler and does not restart MMA2.
3. One invalid/conflicting edit is rejected and the previous active valid runtime state remains unchanged.

## Verification

From the UI, perform exactly three saves: one structural change, one timing-only change, and one invalid/conflicting change. Record which path is invoked and prove final runtime state matches the expected result for each case.

**Evidence (2026-09-08):** Added `simulator/apply.go` with validation-first document classification and explicit `mma2-structural`, `random-runtime`, and `no-change` outcomes. Structural changes route through the SIM-002 ownership/composition consumer and rebuild the simulator schedulers; timing-only changes route through `Scheduler.UpdateTiming` and do not touch MMA2 configuration. The bridge now returns the selected path/message and the OS.js window surfaces it. Browser verification performed exactly three saves: port 15020→15021 reported the MMA2 structural path; FC1 timing 1000→1250 ms reported random-runtime update without restart; a second device conflicting on `(15021,1)` returned HTTP 422. The final persisted document retained one device at port 15021 with FC1 interval 1250, while effective MMA2 config and ownership retained the single simulator reservation. `apply_test.go` and the applying-bridge test prove route counts, validation-before-consumer behavior, and preservation of prior persisted state on downstream rejection. `gofmt -l .`, `go vet ./...`, `go test -count=1 ./...`, and `npm run build:local-packages` pass.

## Dependencies

- SIM-002.
- SIM-003.
- SIM-005.

## Sizing

**3 / 10 — Small.** One routing/transaction outcome with three tightly related acceptance cases.
