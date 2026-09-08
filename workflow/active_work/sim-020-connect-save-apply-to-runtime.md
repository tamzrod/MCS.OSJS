# SIM-020 — Connect OS.js Save & Apply to the Runtime

Status: ACTIVE — promoted by user 2026-09-08; execute after SIM-019.

## Primary outcome

Make the OS.js **Save & Apply** action submit the complete edited document to the real local `ApplyRouter` and update the UI's persisted snapshot only after the runtime accepts it.

## Scope

- Load the canonical Simulator document through the SIM-018 local boundary.
- Submit the normalized complete document through the SIM-019 runtime owner.
- Display the returned classification and outcome: structural apply/restart, timing-only update, or no change.
- Keep the prior applied document and Discard snapshot when validation, ownership, restart/readiness, or persistence fails.
- Remove the misleading success message that currently means only OS.js settings were saved.

## Non-scope

- No runtime-status panel or polling; that belongs to SIM-021.
- No new MMA2 control operation.
- No second canonical document in per-user OS.js settings.

## Acceptance criteria

1. “Applied” is displayed only after `ApplyRouter.Apply` succeeds.
2. Structural edits report restart/readiness outcome; timing-only edits report that MMA2 was not restarted.
3. A rejected edit leaves the last applied document available to Discard and shows the actual failure.
4. Reloading the window shows the canonical Simulator document used by boot restore, not a divergent settings copy.

## Verification

In the rebuilt OS.js application, perform one structural save, one timing-only save, and one rejected duplicate reservation; compare UI outcome, canonical document, MMA2 restart artifact behavior, and scheduler state.

## Dependencies

- SIM-019.

## Sizing

Implementation 1, environment 0, behavioral 2, verification 1, decision/recovery 0 = 4. Tightly coupled UI apply transaction with one browser verification workflow.
