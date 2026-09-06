# SIM-006 — Route Save & Apply to the Correct Parameter Consumer

Status: ACTIVE — human-promoted for JR execution.

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

## Dependencies

- SIM-002.
- SIM-003.
- SIM-005.

## Sizing

**3 / 10 — Small.** One routing/transaction outcome with three tightly related acceptance cases.