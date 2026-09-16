# MEM-006 — Default New Memory Areas to None

Status: QUEUED
Previous: MEM-005
Next: MEM-007

## Primary Outcome
Make newly created Memory devices default to no simulation while preserving Random behavior for existing saved devices with positive intervals.

## Scope
- Change new-device defaults so FC1-FC4 random intervals start at 0.
- Interpret existing positive persisted intervals as Random in the UI.
- Ensure Duplicate preserves the source device's current simulation modes/intervals.
- Do not rewrite existing persisted documents merely because they are loaded.

## Non-Scope
- No on-disk schema migration.
- No automatic conversion of existing positive intervals to None.
- No new generator types.

## Acceptance Criteria
1. A newly added Memory device shows None for all four areas.
2. Existing saved positive intervals show Random after load.
3. Duplicate preserves source simulation settings.
4. Load alone does not modify persisted configuration.

## Verification
Renderer source re-read and focused compatibility reasoning; final rendered/persistence check is MEM-008.

## Dependencies
MEM-005.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 0. Total 3.