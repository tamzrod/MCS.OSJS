# MEM-002 — Simplify Top Runtime Status

Status: QUEUED
Previous: MEM-001
Next: MEM-003

## Primary Outcome
Remove the Simulator runtime indicator from the top header so Memory is not visually treated as a top-level process status.

## Scope
- Remove the top `Simulator` runtime status item.
- Remove Simulator header activity pulse handling.
- Keep MMA2 and Replicator top status indicators aligned and functional.
- Do not stop, rename, or otherwise alter the Simulator process.

## Non-Scope
- No Diagnostics relocation yet.
- No MMA2 or Replicator runtime semantics changes.
- No backend/service changes.

## Acceptance Criteria
1. Header contains MMA2 and Replicator status only.
2. No Simulator LED/text remains in the header.
3. Simulator runtime continues to run and remain queryable internally.

## Verification
Static source re-read of `electron/renderer/index.html` and `electron/renderer/app.js`; final rendered verification is MEM-008.

## Dependencies
MEM-001.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 0. Total 3.