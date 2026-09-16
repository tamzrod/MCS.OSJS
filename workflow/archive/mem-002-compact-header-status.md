# MEM-002 — Simplify Header Status

Status: COMPLETED — static verification passed 2026-09-17
Previous: MEM-001
Next: MEM-003

## Primary outcome
Keep only MMA2 and Replicator service indicators in the Electron header, aligned and non-blinking.

## Scope
Remove header Simulator indicator and header activity pulse, retain MMA2/Replicator in-place state updates. Preserve Simulator runtime process, status stream and logging.

## Non-scope
No process lifecycle or service changes.

## Acceptance
1. Header contains only MMA2 and Replicator aligned on one line at normal width.
2. Both show running/stopped without blinking.
3. Simulator runs unaffected.

## Verification
Static source verification passed: `electron/renderer/index.html` contains only MMA2 and Replicator header status nodes; `electron/renderer/app.js` updates only those two header nodes and separately updates Diagnostics Simulator state; `electron/renderer/style.css` keeps the runtime strip on one line and contains no activity-pulse animation path. Windows visual/runtime confirmation remains MEM-008.

## Dependencies
MEM-001.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
