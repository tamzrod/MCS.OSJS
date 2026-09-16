# MEM-003 — Simulator Status in Diagnostics

Status: QUEUED
Previous: MEM-002
Next: MEM-004

## Primary outcome
Expose Simulator runtime state in Diagnostics instead of the header.

## Scope
Add a compact Diagnostics field bound to the existing runtime status stream. Update its text in place and retain logs/controls.

## Non-scope
No new backend status API or polling.

## Acceptance
1. Simulator status remains visible in Diagnostics.
2. Existing status events update the field in place.
3. No Simulator indicator in header.

## Verification
Inspect selector/binding; Windows live check in MEM-008.

## Dependencies
MEM-002.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
