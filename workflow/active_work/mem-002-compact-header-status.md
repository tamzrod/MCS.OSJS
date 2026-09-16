# MEM-002 — Simplify Header Status

Status: QUEUED
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
Static markup/JS/CSS inspection; actual Windows visual/runtime checks in MEM-008.

## Dependencies
MEM-001.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
