# FMT-003 — OS.js Toolkit UI inventory
Task ID: FMT-003
Task Name: Inventory existing OS.js Toolkit shell and launcher
Blocker Task: NONE
Status: PENDING
Assigned Agent: OpenCode — read-only discovery
Stage: DISCOVERY

## Objective
Identify implemented Toolkit entry, shell/navigation and launcher registration without claiming rendered behavior.
## Scope
Read `ICC/INDEX.md`, `OSJS/src/packages/MCSModbusToolkit/index.js`, `toolkit-renderer.js`, `metadata.json`, and direct launcher registration references discovered through OS.js client config. Read only relevant discovered paths; record them before access. Write only `planning/microtask/evidence/fmt-003-osjs-ui-inventory.md` after activation. No product/ICC/workflow edits or runtime actions.
## Execution
1. Record HEAD/branch/status; inspect entry and renderer to map existing panels.
2. Trace package metadata and launcher/icon registration from actual source.
3. Identify existing build/test script names from package manifests only; do not execute checks.
## Acceptance Criteria
1. Existing panels and entry/navigation are mapped to real source lines.
2. Launcher/icon configuration is classified implemented/unknown rather than inferred from intent.
3. Exact discovered build/test scripts are recorded without execution or PASS claims.
## Evidence
`planning/microtask/evidence/fmt-003-osjs-ui-inventory.md`: pinned SHA, source path/line map, known/unknown list, discovered commands.
## Completion and delivery
COMPLETE with all three outcomes evidenced; FAIL for incomplete outcomes; BLOCKED if source inaccessible. Only report and authorized task status may change. No commit/push authority until separately specified on promotion. STOP.
## Sizing
Implementation 0; environment 0; behavior 0; verification 1; decision/recovery 1 = 2/10. Shell/launcher only, not editor internals.