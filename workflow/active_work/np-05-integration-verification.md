# NP-05 — Integration Verification

Status: PROMOTED / QUEUED (human authorized 2026-10-08)
Stage: TEST/VERIFY
Owner: Independent JR TEST/VERIFY (separate identity; not Codex JR DEV) (operator assignment 2026-10-08)
Previous: NP-04
Next: NP-06

## Scope
Using merged MMA2 native persistence, verify config save, MMA2 restart, Modbus write, graceful restart, restored value, disabled-memory isolation and invalid-config rejection. Do not change product implementation except a narrowly scoped test fix with explicit authorization.

## Acceptance
Record exact commands and observed results; do not claim PASS without execution.

## Shared contract
MMA2 PR #22, docs/PERSISTENCE.md: persistence only at listeners[].memory[].persistence; MMA2 owns snapshot, flush, backup, restore and recovery. MCS owns UI, validation and YAML only. No user snapshot deletion. Do not refactor unrelated features.

## Execution
Exactly one assigned packet per invocation; read task-relevant files only. Respect predecessor completion and OPERATION CWAL routing. Return changed paths, focused test results, blockers and checkpoint. STOP; do not execute next task.
