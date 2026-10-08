# NP-04 — Legacy Engine Removal

Status: PROMOTED / QUEUED (human authorized 2026-10-08)
Stage: CODE
Previous: NP-03
Next: NP-05

## Scope
Remove old MCS snapshot writer, loader, restore, watchdog, persistence-generated RBE and persistence-specific sealing prerequisites, including obsolete Electron adapters only where proven unused. Keep ordinary RBE, State Sealing, FC43, memory ownership, safe YAML writer and user data. Delete obsolete tests only if replaced.

## Acceptance
Go and Node focused suites build/pass; no MCS snapshot runtime remains.

## Shared contract
MMA2 PR #22, docs/PERSISTENCE.md: persistence only at listeners[].memory[].persistence; MMA2 owns snapshot, flush, backup, restore and recovery. MCS owns UI, validation and YAML only. No user snapshot deletion. Do not refactor unrelated features.

## Execution
Exactly one assigned packet per invocation; read task-relevant files only. Respect predecessor completion and OPERATION CWAL routing. Return changed paths, focused test results, blockers and checkpoint. STOP; do not execute next task.
