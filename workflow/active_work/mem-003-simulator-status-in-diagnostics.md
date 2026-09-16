# MEM-003 — Move Simulator Status to Diagnostics

Status: QUEUED
Previous: MEM-002
Next: MEM-004

## Primary Outcome
Keep Simulator runtime health visible in Diagnostics after removing it from the top header.

## Scope
- Add a compact Simulator runtime status line to Diagnostics.
- Reuse the existing runtime status data already delivered to the renderer.
- Keep the top header free of Simulator status.

## Non-Scope
- No new runtime API.
- No backend/service changes.
- No simulation-mode behavior changes.

## Acceptance Criteria
1. Diagnostics displays current Simulator runtime status.
2. Existing runtime status updates drive the Diagnostics value.
3. No duplicate Simulator indicator appears in the header.

## Verification
Static source re-read; final rendered verification is MEM-008.

## Dependencies
MEM-002.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 0. Total 3.