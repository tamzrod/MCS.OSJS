# SIM-020 — Connect OS.js Save & Apply to the Runtime

Status: COMPLETED 2026-09-09 — canonical load and real Save & Apply verified.

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

## Completion evidence

- The window no longer reads or writes `mcs/modbus-simulator.document`; it loads the canonical Go `Store` document through runtime `load`.
- Save & Apply sends the normalized complete document through runtime `apply`, stays pending for the real response, and advances the document/Discard snapshot only on success.
- Successful responses display the actual apply classification message and completion time. Failure retains the edited form and prior applied Discard snapshot with the backend error.
- Runtime connection framing now remains open beyond the legitimate 20-second MMA2 readiness timeout, so unavailable-MMA2 errors reach the UI instead of becoming a misleading client timeout.
- Rebuilt-browser verification proved structural success, timing-only success with the explicit “without restarting MMA2” result, duplicate-reservation rejection, Discard restoration to one prior device, and canonical one-device reload.
- JS syntax, OS.js local-package/full builds, full Go tests, and Go vet pass on 2026-09-09.

## Dependencies

- SIM-019.

## Sizing

Implementation 1, environment 0, behavioral 2, verification 1, decision/recovery 0 = 4. Tightly coupled UI apply transaction with one browser verification workflow.
