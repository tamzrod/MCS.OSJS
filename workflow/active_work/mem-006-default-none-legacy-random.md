# MEM-006 — Default None, Preserve Legacy Random

Status: QUEUED
Previous: MEM-005
Next: MEM-007

## Primary outcome
New Memory devices default None; existing saved positive intervals remain Random.

## Scope
Zero all four intervals on newly created Electron devices only. Derive selector on load from saved interval values. Keep Duplicate and existing persistence unchanged.

## Non-scope
No migration of documents, config paths or services.

## Acceptance
1. Newly added device defaults to None in every area.
2. Existing saved positive intervals remain Random through load/save.
3. Duplicating a device preserves its mode/interval values.

## Verification
Inspect defaults and load/save mapping; Windows reload in MEM-008.

## Dependencies
MEM-005.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
