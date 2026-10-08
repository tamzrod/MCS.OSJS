# NP-01 — Native Schema & Validation

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); delivered on main at `757fe00b6e8b057b92cb693781fe1885ce362dd6`
Stage: CODE
Previous: none
Next: NP-02

## Scope
Implement MMA2 native listeners[].memory[].persistence schema and candidate validation in mma2composer. enabled false/omitted disables; enabled true requires directory; optional ranges must be inside allocated areas, nonoverlapping and nonempty if specified. Reject root persistence. Preserve existing unrelated configuration.

## Acceptance
Focused composer Go tests for valid/invalid YAML and round-trip.

## Shared contract
MMA2 PR #22, docs/PERSISTENCE.md: persistence only at listeners[].memory[].persistence; MMA2 owns snapshot, flush, backup, restore and recovery. MCS owns UI, validation and YAML only. No user snapshot deletion. Do not refactor unrelated features.

## Execution
Exactly one assigned packet per invocation; read task-relevant files only. Respect predecessor completion and OPERATION CWAL routing. Return changed paths, focused test results, blockers and checkpoint. STOP; do not execute next task.
