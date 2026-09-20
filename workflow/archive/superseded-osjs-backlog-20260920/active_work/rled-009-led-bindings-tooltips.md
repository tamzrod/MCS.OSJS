> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# RLED-009 — Bind LED States and Tooltip Details

Status: QUEUED
Previous: RLED-008
Next: RLED-010

## Primary outcome
Render real runtime status as LED colors and accessible details without interrupting editing.

## Scope
Map Network/TCP/Modbus/MMA2 states to green/yellow/red/gray, and show observed endpoint, FC/block, exception code, errors and last-result time in circle tooltip on hover/focus/tap. Patch existing DOM nodes in place on polling; handle status request failures as unknown.

## Non-scope
No status simulation or backend protocol changes.

## Acceptance
1. All colors reflect real observed status, not CONFIGURED or process existence.
2. Hover/focus/tap reveals the correct per-indicator evidence.
3. Status refresh preserves focus and unsaved editor values.

## Verification
Focused renderer/state mapping checks, static source inspection and Windows interactive acceptance in RLED-011.

## Dependencies
RLED-008.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
