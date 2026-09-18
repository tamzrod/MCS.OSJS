# UMIG-008-V — VERIFY: One-Icon Desktop Cutover

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-008-T
Next: UMIG-009

## Primary outcome
Observe exactly one Toolkit desktop icon with Start menu and taskbar preserved.

## Instruction / expected / evidence
Launch a clean/default isolated OS.js desktop; count icons, double-click or Enter Toolkit shortcut, exercise Start menu, taskbar window button, tray, clock and normal window management. Inspect a pre-existing saved-session case non-destructively. Expected: exactly one default Toolkit icon, a single launched Toolkit window, functional shell, user data intact and legacy Start-menu fallback still available. Return screenshots/action log, icon count, session observations and HEAD. Unavailable GUI/saved-session fixture = BLOCKED for any required part.

## Non-scope
No user profile wipe, kiosk conversion, product fix, old package removal or Windows changes.

## Dependencies
UMIG-008-T PASS, human cutover approval, current JR packet.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
