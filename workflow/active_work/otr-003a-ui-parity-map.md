# OTR-003A — UI source-to-target parity map

Status: QUEUED
Stage: DESIGN
Owner: OpenCode
Previous: OTR-002A
Next: OTR-003B
Size: I0 E0 B1 V1 D1 = 3

## Gate
Need `workflow/active_work/evidence/otr-002a-report.md`. If missing, STOP.

## One outcome
One table mapping Electron renderer UI files to existing OS.js Toolkit files. No code changes.

## Allowed writes
- `workflow/active_work/evidence/otr-003a-map.md`
- `handoff.md` Current task after COMPLETE
- `workflow/active_work/OTR_PROMOTION_QUEUE.md` 003A line only

## Exact files to read
Electron (do not edit):
- `electron/renderer/index.html`
- `electron/renderer/style.css`
- `electron/renderer/app.js`
- `electron/renderer/comms-status.js`
- `electron/renderer/diagnostics.js`
- `electron/renderer/memory-advanced.js`

OS.js (do not edit):
- `OSJS/src/packages/MCSModbusToolkit/toolkit-renderer.js`
- `OSJS/src/packages/MCSModbusToolkit/renderer.css`
- `OSJS/src/packages/MCSModbusToolkit/index.js`
- `OSJS/src/packages/MCSModbusToolkit/memory-editor.js`
- `OSJS/src/packages/MCSModbusToolkit/replicator-editor.js`
- `OSJS/src/packages/MCSModbusToolkit/diagnostics-editor.js`

Ignore `*.bak` files.

## Numbered steps
1. Confirm branch `opencode` and record HEAD.
2. Read listed files only.
3. Write a markdown table with columns: Electron path | OS.js path | copy/adapt/reimplement/gap | notes with line anchors.
4. List the three tabs on each side. Do not invent a fourth tab.
5. Propose the next CODE slice as exactly one file: `OSJS/src/packages/MCSModbusToolkit/toolkit-renderer.js` for shell/nav only. Do not implement it.
6. Verdict MAP COMPLETE or BLOCKED.
7. STOP.

## Forbidden targets
Never write `OSJS/packages/toolkit/app/views/`. That path does not exist.
