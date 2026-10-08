# PERSIST-021 — Persistence Runtime Status

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
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
PERSIST-020 is delivered on GitHub main at `1528d3e2451f3588de46f7539127164588878f1a`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task: expose observational persistence status to the OS.js operator surface from the authoritative runtime restore/sealing state, including enabled/disabled, snapshot health, last save, and last restore outcome when available; restore failures must remain visibly sealed; the UI must not gain any bypass of restore or sealing gates. Prepare PERSIST-022 for the next invocation only after genuine completion, then STOP.
