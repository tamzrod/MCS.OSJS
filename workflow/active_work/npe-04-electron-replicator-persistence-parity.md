# NPE-04 — Electron Replicator Destination Persistence Parity

Status: CODE COMPLETE — delivered on main at `b4580d6fa3bbee547caea4680aacd1a8c812b48b` (OpenHands JR DEV); awaiting independent verification (NPE-06)
Stage: CODE
Owner: OpenHands JR DEV
Previous: NPE-03
Next: NPE-05

## Purpose
Apply the same native per-memory persistence contract to Electron-managed Replicator destination memory where that destination owns MMA2 memory.

## Scope
Inspect the current Replicator destination configuration path and implement only the persistence configuration parity that is valid for its destination memory.

- same native `persistence` block;
- directory optional;
- ranges omitted = all allocated destination areas;
- custom ranges bounded to destination allocation;
- no persistence runtime engine in Replicator;
- no State Sealing or RBE dependency.

If the current Electron Replicator UI intentionally has no advanced memory configuration surface, keep the scope to model/config round-trip and document that UI parity belongs to a later explicit UI task rather than inventing controls.

## Acceptance
Focused Replicator/Electron tests prove native persistence config is preserved and composed without introducing a second runtime owner.

No OS.js Toolkit changes. STOP after delivery.

## Delivery evidence (OpenHands JR DEV)

Delivered SHA: `b4580d6fa3bbee547caea4680aacd1a8c812b48b` on `main` (base `3d20c32`).

Changed paths (all authorized; no OS.js Toolkit changes):
- `electron/renderer/app.js` — the Replicator advanced memory editor now passes `persistenceSupported: true`, so a Replicator destination exposes the same native Persistence tab (driven by the shared model; ranges bounded by the pull-block destination allocation via `replicatorParams`).
- `electron/test/memory-advanced.test.js` — `Replicator destination exposes native persistence parity bounded to pull-block areas` (enable, selected ranges, in-area validation, round-trip into `mma2_advanced.persistence`).
- `electron/test/memory-persistence.test.js` — `Replicator destination composes the same native persistence block without a second runtime` and `Replicator destination defaults persistence to all allocated areas` (native block preserved through `composeAll`; no capture/restore/unseal/sealing/RBE).

Command and results:
- `cd electron && node --test test/*.test.js` → 63 PASS, 1 FAIL (64 total, up from 61). The single failure, `persistence-lifecycle.test.js` "Save & Apply snapshot capture recovers a previously sealed persistence runtime", is PRE-EXISTING and unrelated: it directly exercises the legacy `electron/persistence.js` runtime that is no longer invoked and that NPE-05 removes.

This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.
