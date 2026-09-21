# OTR numbered queue

Human-approved roadmap. One task per invocation. STOP after the ACTIVE packet.

## Numbered list

001A COMPLETE — Electron shell inventory. Archive: `workflow/archive/otr-001a-electron-shell-inventory.md`
001B COMPLETE — Electron editor inventory. Archive: `workflow/archive/otr-001b-electron-editor-inventory.md`
001C COMPLETE-UNTRUSTED — Electron IPC list exists in `workflow/archive/otr-001c-electron-backend-boundary-inventory.md`. Do not rerun. Do not treat as a full contract map.
002A ACTIVE — OS.js Toolkit UI baseline. Packet: `otr-002a-osjs-ui-baseline.md`
002B QUEUED — OS.js Toolkit backend baseline. Packet: `otr-002b-osjs-backend-baseline.md`. After 002A report exists.
003A QUEUED — UI source-to-target map. Packet: `otr-003a-ui-parity-map.md`. After 002A.
003B QUEUED — Backend contract map. Packet: `otr-003b-backend-contract-map.md`. After 002B and 003A.
004  NOT EXECUTABLE parent — `otr-004-toolkit-shell-replica.md`
005  NOT EXECUTABLE parent — `otr-005-simulator-editor-ui.md`
006  NOT EXECUTABLE parent — `otr-006-replicator-editor-ui.md`
007  NOT EXECUTABLE parent — `otr-007-mma-memory-ui.md`
008  NOT EXECUTABLE parent — `otr-008-config-read-adapter.md`
009  NOT EXECUTABLE parent — `otr-009-validation-parity.md`
010  NOT EXECUTABLE parent — `otr-010-safe-config-save.md`
011  NOT EXECUTABLE parent — `otr-011-scoped-restart-adapter.md`
012  NOT EXECUTABLE parent — `otr-012-applied-state-status.md`
013  NOT EXECUTABLE parent — `otr-013-ui-parity-verification.md`
014  NOT EXECUTABLE parent — `otr-014-end-to-end-acceptance.md`

## Rejected archives
Do not use:
- `workflow/archive/otr-002b-baseline-inventory.md`
- `workflow/archive/otr-003a-ui-parity-map-inventory.md`
Reason: invented target paths. Real Toolkit lives in `OSJS/src/packages/MCSModbusToolkit/`.

## Selection rule
If handoff Current task exists and its packet file exists, run that packet.
Else run the first ACTIVE or first QUEUED discovery packet that has predecessor evidence on disk under `workflow/active_work/evidence/`.
If neither exists, STOP and report the missing file name.
