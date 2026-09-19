# UMIG-007 — VERIFY: Three-Tab Visual Parity

Status: QUEUED / DEFERRED — predecessor UMIG-006-V COMPLETE/PASS; 2026-09-19 human direction now follows the CURRENT Electron MCS Toolkit model. Do not activate visual-only parity before functional/model alignment and an approved current-model baseline.
Stage / owner: VERIFY / OpenHands JR; packet, review and advancement / ChatGPT.
Previous: UMIG-006-V (archived PASS; historical chronological link only). New model-alignment dependencies are described below and must be formally linked through approved microtasks before activation.
Next: UMIG-007A (QUEUED; no cutover approval implied).

## Primary outcome
Independently compare the functional OS.js Toolkit's three-tab appearance with the human-approved CURRENT Electron model, after the model has been mapped and implemented through separate CODE/TEST/VERIFY gates. Do not treat the old frozen one-time renderer screenshot or source as the current target.

## Changed dependency / source truth
Original frozen one-time visual donor `1c971b9a6e00bafadf329df8821421a40cfc079c` remains immutable history and explains current Toolkit provenance. Human's newer direction is Electron MCS Toolkit at `ff6846f973ad973bd3270b0ff2d707ce5c2616c9` for the current review; a further Electron push requires a new assessed diff/re-pin. That Electron model introduces Device Definition/Advanced folder tabs, per-device RBE/State Sealing/Access Policy, a separate shared MMA Settings modal and new Go/MMA2 RBE behavior. OS.js has not yet implemented those features. The earlier UMIG-006-V PASS is valid at its tested SHA, not a regression pass for upgraded MMA2 or feature parity.

First finalize the PLANNED `planning/microtask/umig-em-001-electron-model-contract.md` design/compatibility gate and separately authorize its resulting small CODE/TEST/VERIFY tasks. Current-model Memory, Replicator and Diagnostics donor screenshots or an explicit human-approved reproducible capture procedure must then be fixed at comparable window size/state, with durable accessible artifacts. The image previously shared in chat is an isolated Electron Memory Advanced/RBE view with design annotations; it is NOT a three-tab approved baseline or a runtime acceptance test. Original `/tmp/jr-shots/` artifacts are not durable donor evidence.

## VERIFY action AFTER all dependencies
Under a separate exact task-specific `handoff.md` packet, compare actual OS.js Memory (definition, advanced, shared modal), Replicator and Diagnostics against current-model Electron at matched dimensions and safe comparable synthetic states. Check navigation, forms, control availability, spacing, selection, scrolling, taskbar/window chrome and ShadowRoot CSS isolation. Separate platform-specific controls (Windows service/bin paths) and unsupported RBE network output from visual discrepancies; never fabricate health, COMMS, logs or backend results. Require paired screenshots/observations and discrepancies, not source-only inference. Missing baseline/target => BLOCKED; proven discrepancy => FAIL for independently authorized repair.

## Boundaries
No code fixes/debugging by JR, production user data, enabled RBE network listener without separate safety review, protocol changes, legacy retirement, cutover, ICC or general CWAL changes. QUEUED is not executable; no JR packet exists yet.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
