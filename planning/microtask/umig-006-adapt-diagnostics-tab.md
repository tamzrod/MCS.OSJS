# UMIG-006 — CODE: Adapt OS.js Diagnostics

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: CODE / ChatGPT
Previous: UMIG-005-V
Next: UMIG-006-T

## Primary outcome
Author Diagnostics mapping that shows only truthful OS.js-available observations.

## Scope
Preserve donor layout; map supported Simulator/Replicator/MMA2 observations/errors; disable Windows-only native service controls and filesystem paths when unsupported; show unknown/unavailable rather than green by default.

## Non-scope
No live testing, privileged new endpoint, backend service controls, Windows IPC or desktop redesign.

## Coding acceptance / handoff
1. Supported diagnostics map actual status/error contracts.
2. Unsupported native controls cannot call Windows APIs, and absent data is unknown/unavailable.
3. Record implementation files/commit and unsupported capabilities for OpenHands, without claiming tests.

## Dependencies
UMIG-005-V PASS and human promotion; any new service-control capability requires independent approval.

## Sizing
Surface 1, environment 0, behavior 1, verification 0, recovery 0 = 2.
