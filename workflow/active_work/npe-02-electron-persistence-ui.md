# NPE-02 ‚Äî Electron Persistence UI

Status: CODE COMPLETE ‚Äî delivered on main at `a90f52a1b7ce2830dbc2aa8b6a65041962f21355` (OpenHands JR DEV); awaiting independent verification (NPE-06)
Stage: CODE
Owner: OpenHands JR DEV
Previous: NPE-01
Next: NPE-03

## Purpose
Implement the native MMA2 persistence configuration experience in the Electron desktop first.

## Scope
In Electron Memory Advanced Settings, expose only native persistence configuration:

- Enabled
- Storage location:
  - default/native location when directory omitted
  - optional custom directory override
- Persisted memory:
  - default = All Allocated Areas
  - optional Selected Ranges
- selected ranges for coils, discrete inputs, holding registers and input registers, bounded by each allocated area.

Persistence UI must NOT expose:
- persistence RBE
- persistence lock coil
- persistence State Sealing dependency
- restore/unseal controls
- external snapshot manager controls

State Sealing remains a separate independent tab/feature.

## UX rule
The simple default is:

```
Persistence
[‚úì] Enabled
Persist: All Allocated Areas
Directory: Default
```

Custom ranges/directory are advanced overrides, not required fields.

## Acceptance
Focused Electron renderer tests cover default state, enable/disable, optional directory, whole-memory default, custom range editing, validation, save/reload and independence from State Sealing/RBE.

No OS.js Toolkit changes in this packet. STOP after delivery.

## Delivery evidence (OpenHands JR DEV)

Delivered SHA: `a90f52a1b7ce2830dbc2aa8b6a65041962f21355` on `main` (base `4335c80`).

Changed paths (all authorized; no OS.js Toolkit changes):
- `electron/renderer/memory-advanced.js` — Persistence tab now exposes only native configuration: Enabled (writes `persistence.enabled` only, no sealing/RBE mutation), optional `Directory` (empty = MMA2 native default), `Persisted memory` default "All Allocated Areas" (ranges omitted) with optional "Selected Ranges" bounded by each allocated area, per-area add/edit/delete of ranges, and in-UI range validation. Removed: persistence-owned derived RBE locked rows, lock-coil + sealed-response controls, "managed by Persistence" wording, and the sealing-tab redirect. State Sealing remains an independent tab. Exports changed to `persistenceEnabled`/`persistenceRangeError`/`allocatedAreas`.
- `electron/test/memory-advanced.test.js` — focused tests: enable/disable without touching sealing/RBE; default all-areas read-only view and absence of removed controls; optional directory round-trip; selected-range editing, in-area validation and all-areas reset; independence from State Sealing and RBE.

Commands and results:
- `cd electron && node --test test/memory-advanced.test.js` → 14 tests PASS, exit 0.
- `cd electron && node --test test/*.test.js` → 55 PASS, 1 FAIL. The single failure, `persistence-lifecycle.test.js` "Save & Apply snapshot capture recovers a previously sealed persistence runtime" (`Modbus coils read rejected with 0x06`), is PRE-EXISTING and unrelated: it is confirmed failing with this packet's changes stashed, and it exercises the legacy `electron/persistence.js` runtime (snapshot capture / Raw-Ingest restore) that NPE-05 removes, not the NPE-02 UI.

This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.
