# UMIG-008 — CODE: Switch OS.js Toolkit Launchers

Status: PLANNED / BLOCKED — independent deployment and cutover approval required.
Stage / owner: CODE / ChatGPT
Previous: UMIG-007A
Next: UMIG-008-T

## Primary outcome
Author project-owned launcher cutover to exactly one default `MCS Modbus Toolkit` desktop icon.

## Scope
Replace project-owned old desktop shortcuts/default auto-start mapping; keep OS.js desktop, Start menu, bottom taskbar, window buttons, tray, clock and user-managed icons/files/settings intact. Legacy packages remain as Start-menu fallback until UMIG-009. Handle saved-session references non-destructively; Windows Electron remains separate.

## Non-scope
No test/visual PASS, kiosk, backend changes, user-profile purge, deleting legacy packages or Electron changes.

## Coding acceptance / handoff
1. Project-owned default launcher/auto-start code targets Toolkit and one project icon.
2. Legacy packages, shell chrome and user-managed data are not deleted.
3. Record exact diff/commit and saved-session handling for OpenHands; do not claim runtime success.

## Dependencies
UMIG-007 and UMIG-007A PASS, focused tab verification PASS, explicit human cutover approval/promotion.

## Sizing
Surface 1, environment 0, behavior 1, verification 0, recovery 1 = 3.
