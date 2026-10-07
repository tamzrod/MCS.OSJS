# PERSIST-021 — Persistence Runtime Status

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: PERSIST-020
Next: PERSIST-022

## Primary outcome
Expose persistence health/status to the OS.js operator surface without changing persistence authority.

## Scope
Add observational status only.

## Non-scope
Do not add historian/telemetry persistence, alternate control authority, or any UI bypass of restore/sealing gates.

## Acceptance
1. Status includes enabled/disabled, snapshot health, last save and last restore outcome when available.
2. Restore failure is clearly visible while sealed state remains authoritative.
3. Status is observational; UI cannot bypass restore/sealing gates.

## Evidence / handoff
Record exact changed paths, source diff, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-020.

## Sizing
2/0/2/1/1=6.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`.
