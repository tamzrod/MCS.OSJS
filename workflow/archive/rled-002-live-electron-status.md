# RLED-002 — Replace Fake Electron Replicator Status

Status: COMPLETED — focused IPC/status and packaging verification passed 2026-09-17
Previous: RLED-001
Next: RLED-003

## Primary outcome
Return the real Go runtime `status` result through the existing Electron `replicatorCall` IPC operation.

## Scope
Use the RLED-001 client and existing data root. Package its module. Remove the hardcoded RUNNING/CONFIGURED response, propagate runtime-unavailable errors honestly. Preserve load/apply and existing IPC names.

## Non-scope
No new polling loop, service manager or UI LED yet.

## Acceptance
1. Status calls return real per-device and block data.
2. Unavailable runtime rejects rather than reporting green.
3. Load/apply remain unchanged.

## Verification
Verification passed: Node syntax checks passed for main, renderer and both bridge modules; all ten named-pipe and IPC fixtures passed; npm run pack:win succeeded; app.asar contains replicator-runtime.js and replicator-ipc.js. Windows live acceptance remains RLED-011.

## Dependencies
RLED-001.

## Sizing
Surface 2, environment 0, behavior 1, verification 1, recovery 0 = 4; tightly coupled IPC wiring and packaging with one test workflow.
