# NP-02 — Persistence UI

Status: SUPERSEDED — 2026-10-08
Stage: HISTORICAL / DO NOT EXECUTE
Owner: Codex JR DEV (operator assignment 2026-10-08)
Previous: NP-01
Next: NP-03

## Scope
Replace Memory Advanced Settings persistence controls with Enable, Directory, and All allocated areas / Selected ranges. Reuse existing area/range editor; no new snapshot or RBE controls. Preserve unrelated settings.

## Acceptance
Focused OSJS UI tests: defaults, validation, edit/save/reload.

## Shared contract
MMA2 PR #22, docs/PERSISTENCE.md: persistence only at listeners[].memory[].persistence; MMA2 owns snapshot, flush, backup, restore and recovery. MCS owns UI, validation and YAML only. No user snapshot deletion. Do not refactor unrelated features.

## Execution
Exactly one assigned packet per invocation; read task-relevant files only. Respect predecessor completion and OPERATION CWAL routing. Return changed paths, focused test results, blockers and checkpoint. STOP; do not execute next task.


## Superseded routing
This packet is superseded by the human-approved phased native-persistence chain:
NPE-01 → NPE-02 → NPE-03 → NPE-04 → NPE-05 → NPE-06 (Electron complete + verify) → NPO-01 → NPO-02 → NPO-03 → NPO-04 (OS.js Toolkit complete + verify) → NPF-01.

Do not execute this packet.
