# UMIG-009 — CODE: Retire Legacy OS.js UI Packages

Status: PLANNED / BLOCKED — launcher cutover verified and human approval required.
Stage / owner: CODE / ChatGPT
Previous: UMIG-008-V
Next: UMIG-009-T

## Primary outcome
Author removal of superseded `ModbusSimulator` and `ModbusReplicator` OS.js UI packages from discovery/build without touching backend systems.

## Scope
Only after all Toolkit gates pass, remove obsolete UI package source/metadata/discovery references. Ensure every required Simulator/Replicator OS.js bridge has been moved into Toolkit. Record pre-removal known-good commit for rollback.

## Non-scope
No test execution, user config deletion, Go runtime/MMA2 removal, OS.js desktop/theme removal, Windows Electron changes or cross-deployment links.

## Coding acceptance / handoff
1. Legacy UI packages and project-owned references removed; Toolkit remains registered.
2. Existing backend bridges, data, services and desktop shell source preserved.
3. Record before/after paths, source commit and rollback SHA for OpenHands; do not claim post-removal behavior verified.

## Dependencies
UMIG-008-V PASS, focused tab tests and independent deployment verified, explicit human retirement approval/promotion.

## Sizing
Surface 1, environment 0, behavior 1, verification 0, recovery 1 = 3.
