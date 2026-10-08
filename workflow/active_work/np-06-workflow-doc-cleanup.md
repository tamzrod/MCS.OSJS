# NP-06 — Workflow & Documentation Cleanup

Status: SUPERSEDED — 2026-10-08
Stage: HISTORICAL / DO NOT EXECUTE
Owner: Codex JR DEV (operator assignment 2026-10-08)
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


## Superseded routing
This packet is superseded by the human-approved phased native-persistence chain:
NPE-01 → NPE-02 → NPE-03 → NPE-04 → NPE-05 → NPE-06 (Electron complete + verify) → NPO-01 → NPO-02 → NPO-03 → NPO-04 (OS.js Toolkit complete + verify) → NPF-01.

Do not execute this packet.
