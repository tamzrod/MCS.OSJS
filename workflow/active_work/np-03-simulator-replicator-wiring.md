# NP-03 — Simulator & Replicator Wiring

Status: PROMOTED / QUEUED (human authorized 2026-10-08)
Stage: CODE
Previous: NP-02
Next: NP-04

## Scope
Wire native per-memory persistence through Simulator and Replicator config compose/update paths. Preserve foreign memories and ownership; no runtime persistence engine.

## Acceptance
Focused simulator and replicator Go tests; config round-trip.

## Shared contract
MMA2 PR #22, docs/PERSISTENCE.md: persistence only at listeners[].memory[].persistence; MMA2 owns snapshot, flush, backup, restore and recovery. MCS owns UI, validation and YAML only. No user snapshot deletion. Do not refactor unrelated features.

## Execution
Exactly one assigned packet per invocation; read task-relevant files only. Respect predecessor completion and OPERATION CWAL routing. Return changed paths, focused test results, blockers and checkpoint. STOP; do not execute next task.
