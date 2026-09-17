# RLED-008 — Compact Electron LED Layout

Status: QUEUED
Previous: RLED-007
Next: RLED-009

## Primary outcome
Replace the selected Replicator device's verbose runtime row with the approved compact LEDs.

## Scope
In the Electron renderer, show SOURCE labels Network/TCP/Modbus with three circular indicators and DESTINATION label MMA2 with one circle; place hover/focus/tap targets without permanent value text. Preserve device editor fields and actions.

## Non-scope
No OS.js redesign, backend state invention or permanent status text.

## Acceptance
1. Exactly four clearly ordered labeled circles appear.
2. Configuration inputs and buttons remain usable.
3. Unknown indicators initialize gray.

## Verification
Inspect renderer markup/CSS and keyboard/touch semantics; Windows rendered acceptance in RLED-011.

## Dependencies
RLED-007.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
