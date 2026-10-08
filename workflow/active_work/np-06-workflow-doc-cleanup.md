# NP-06 — Workflow & Documentation Cleanup

Status: PROMOTED / QUEUED (human authorized 2026-10-08)
Stage: CODE
Previous: NP-05
Next: none

## Scope
Retire old PERSIST workflow packets to archive without losing history, update handoff/queue and docs to state MMA2 is sole persistence owner; remove dead references and obsolete tests only after NP-05 passes.

## Acceptance
No active legacy persistence task routing; links and docs accurate.

## Shared contract
MMA2 PR #22, docs/PERSISTENCE.md: persistence only at listeners[].memory[].persistence; MMA2 owns snapshot, flush, backup, restore and recovery. MCS owns UI, validation and YAML only. No user snapshot deletion. Do not refactor unrelated features.

## Execution
Exactly one assigned packet per invocation; read task-relevant files only. Respect predecessor completion and OPERATION CWAL routing. Return changed paths, focused test results, blockers and checkpoint. STOP; do not execute next task.
