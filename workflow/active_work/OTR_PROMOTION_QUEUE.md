# OTR numbered queue

Human-approved roadmap. One task per invocation. STOP after the ACTIVE packet.

## Numbered list

001A COMPLETE — Electron shell inventory. Archive: `workflow/archive/otr-001a-electron-shell-inventory.md`
001B COMPLETE — Electron editor inventory. Archive: `workflow/archive/otr-001b-electron-editor-inventory.md`
001C COMPLETE-UNTRUSTED — Electron IPC list exists in `workflow/archive/otr-001c-electron-backend-boundary-inventory.md`. Do not rerun. Do not treat as a full contract map.
002A COMPLETE — OS.js Toolkit UI baseline discovered and written to evidence/otr-002a-report.md ✅
002B COMPLETE — OS.js Toolkit backend inventory complete, written to evidence/otr-002b-report.md for next invocation ✅
003A COMPLETE — UI source-to-target map. Packet: `otr-003a-ui-parity-map.md`. After 002A.
003B ACTIVE — Pending 003A activation. Packet: `otr-003b-backend-contract-map.md`. After 002B and 003A.
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
Run only the task named Current task in root `handoff.md`, and only when its packet exists and says
`Status: ACTIVE`. Never fall back to the first QUEUED or numerically next task. A missing packet,
status mismatch, or `NONE` means STOP and report the exact mismatch.

On successful completion, the current packet may atomically mark itself COMPLETE and its named
successor ACTIVE, commit/push that workflow state with its evidence, verify delivery, and STOP. The
successor runs only in a later invocation.
