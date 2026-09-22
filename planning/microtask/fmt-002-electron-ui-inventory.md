# FMT-002 — Electron UI inventory
Task ID: FMT-002
Task Name: Inventory Electron Toolkit shell and navigation entry points
Blocker Task: NONE
Status: PENDING
Assigned Agent: OpenCode — read-only discovery
Stage: DISCOVERY

## Objective
Identify actual Electron UI entry, shell/navigation definitions and links to screen modules without deeply analyzing each editor.
## Scope
Read `ICC/INDEX.md`, `electron/package.json`, `electron/main.js`, `electron/preload.js`, and only UI files directly referenced by those entries. Confirm actual paths from repository tree before opening. Write only `planning/microtask/evidence/fmt-002-electron-ui-inventory.md` after activation. No source, ICC, runtime, workflow or production changes.
## Execution
1. Record HEAD/branch/status and resolve UI entry paths from actual manifest and imports.
2. Record shell/navigation routes and directly linked screen modules and reusable asset entry paths.
3. Classify Electron-only IPC/runtime dependencies separately from portable UI; mark rendered appearance unknown without screenshots. Do not deep-audit individual views; report them for separate packets if needed.
## Acceptance Criteria
1. Exact entry, shell and navigation paths are source-linked.
2. Direct screen/asset entry references and Electron-only dependencies are classified.
3. Unknown or missing rendered evidence is explicitly recorded without fabricated parity.
## Evidence
`planning/microtask/evidence/fmt-002-electron-ui-inventory.md`: pinned SHA, path/line matrix, screen entry list, unknowns and screenshot availability.
## Completion and delivery
COMPLETE only with all acceptance evidence; FAIL for missing outcome; BLOCKED for inaccessible source. Only evidence report and authorized task status may change. No commit/push authority until promotion specifies exact ref and command. STOP after one task.
## Sizing
Implementation 0; environment 0; behavior 0; verification 1; decision/recovery 1 = 2/10. Entry-point inventory only; split editor deep dives.