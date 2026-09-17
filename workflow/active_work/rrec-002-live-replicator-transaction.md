# RREC-002 — Single Live Replicator Transaction

Status: COMPLETE
Previous: RREC-001
Next: RREC-003

## Primary outcome
Route Replicator load and apply through the existing live Go runtime transaction.

## Scope
Use the named-pipe client for load/apply/status. Remove Electron-local Replicator composition from the UI path while preserving response shapes and foreign reservations.

## Non-scope
No LED state model or Modbus protocol change.

## Acceptance
1. Load returns the runtime canonical document.
2. Apply updates persisted config and running pollers atomically.
3. Status immediately recognizes new/renamed devices.

## Verification
Focused Node dispatch tests, Go runtime apply tests, package build.

## Dependencies
RREC-001.

## Sizing
Surface 2, environment 0, behavior 2, verification 1, recovery 0 = 5; one coupled authority correction.
## Completion evidence
- Electron load, apply, and status now call the live Replicator named-pipe runtime.
- The Electron-local Replicator apply implementation was removed.
- Node focused tests passed: 10/10.
- Go apply, rollback, and foreign-reservation tests passed.
- npm run pack:win passed.
- The full Go suite remains limited by Windows E2E executable naming and the ambient OSJS_DATA_DIR test precondition.