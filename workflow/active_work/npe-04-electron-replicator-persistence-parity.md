# NPE-04 — Electron Replicator Destination Persistence Parity

Status: PROMOTED / QUEUED
Stage: CODE
Owner: Codex JR DEV
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
