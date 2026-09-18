# UMIG-002-V — VERIFY: One Toolkit Window

Status: ACTIVE — human-authorized independent rendered verification; not run.
Stage / owner: VERIFY / OpenHands (JR via Operation CWAL)
Previous: UMIG-002-T
Next: none

## Primary outcome
Observe the built/discovered placeholder opening as exactly one OS.js Toolkit window without disturbing the existing desktop.

## Verification action
Follow only the current `JR TEST TASK — CURRENT: UMIG-002-V` packet in `handoff.md`. In a disposable OS.js test session, use the repository-native build/discovery/serve commands and the existing Start menu to launch `MCS Modbus Toolkit` exactly once. If a different launch method is necessary, STOP and report BLOCKED/request an updated packet rather than inventing one.

Expected: exactly one Toolkit window with the clearly marked `NOT CONNECTED — PLACEHOLDER ONLY` placeholder; desktop, Start menu and taskbar remain functional; legacy Simulator and Replicator applications remain available. No backend health or final cutover is claimed.

Evidence: HEAD SHA and pre/post git status, exact build/serve and browser/launch steps, URL, direct screenshot or rendered observation, Toolkit window count, visible placeholder text, desktop/menu/taskbar and legacy-package observations, startup/browser errors, cleanup. If a required GUI observation cannot be made, BLOCKED rather than PASS. An executable test contradicting a required result is FAIL. No source-only or HTTP-only substitute for rendered evidence.

## Non-scope
No functional Memory/Replicator/Diagnostics acceptance, icon cutover, deletion, product fixes, backend startup, Electron/Windows work, ICC writes or workflow advancement by JR.

## Dependencies
UMIG-002-T reported independent PASS at `e9d25e3d4c1b832126a161a6071fd4abf8b5546b`, reviewed and archived by the coding agent. Previous task names this successor; human authorization already exists. The current `handoff.md` packet must be present, and affected ICC context must be refreshed by BLACK SHEEP WALL before JR execution.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
