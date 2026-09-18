# UMIG-008 — Switch OS.js Launchers to One Toolkit

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-007A
Next: UMIG-009

## Primary outcome
Make MCS Modbus Toolkit the only default desktop icon and normal MCS application launch target while retaining the OS.js desktop, Start menu and taskbar.

## Scope
For a fresh/default OS.js workstation desktop, replace project-owned legacy Modbus desktop shortcuts with exactly one visible icon labeled `MCS Modbus Toolkit`, which launches the single Toolkit window. Preserve the existing desktop background, window management, working Start menu and bottom taskbar (including window buttons, tray and clock); do not switch to kiosk mode or remove OS.js desktop infrastructure. Update project-owned configured default launch/auto-start references to Toolkit. Handle existing saved-session references to retired names deliberately without silently deleting user settings or user-managed desktop files/icons. Legacy packages may remain discoverable through the Start menu as fallback until UMIG-009 retires their UI packages. Only the OS.js deployment/launchers change; Windows Electron remains a separate standalone installation.

## Non-scope
Do not delete legacy UI packages in this task, erase user profiles or user-managed desktop content, alter backend services, link deployments, change Electron launch behavior, remove the Start menu/taskbar, or impose kiosk/fullscreen mode.

## Acceptance
1. On a clean/default desktop, exactly one desktop icon is visible: `MCS Modbus Toolkit`; double-click/Enter launches one Toolkit window, without duplicate or legacy desktop shortcuts.
2. OS.js desktop, Start menu and taskbar remain visible and functional, and new project-owned launch/auto-start references target Toolkit; no unexpected loss of saved user/session content.
3. Legacy package source remains available as fallback until UMIG-009, and Windows Electron deployment is unchanged.

## Verification
Inspect the project-owned shortcut provider and launcher/auto-start mappings, then launch from a clean OS.js session. Visually check icon count/label and single-window launch; exercise Start menu and taskbar window buttons, tray and clock. Inspect existing saved-session behavior without wiping user data, and confirm the diff touches no Electron packaging/launcher files. Record actual results; no runtime acceptance is implied by shortcut inspection alone.

## Dependencies
UMIG-007 visual gate, UMIG-007A independent-deployment gate and successful focused runtime gates UMIG-004/005/006; human approval of cutover/promotion. Existing saved-session cleanup or user-managed icon removal, if needed, requires an explicitly approved separate scope.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 1 = 4 (one bounded OS.js launcher cutover; leave legacy package retirement to UMIG-009).
