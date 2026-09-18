# UMIG-008-T — TEST: Toolkit Launcher Configuration

Status: PLANNED / BLOCKED — cutover promotion required.
Stage / owner: TEST / OpenHands (JR)
Previous: UMIG-008
Next: UMIG-008-V

## Primary outcome
Validate project-owned shortcut/auto-start configuration and package discovery after coding cutover.

## Instruction / expected / evidence
On a disposable checkout, run the exact launcher/config checks and OS.js package discovery command from current JR packet. Expected: one project-owned desktop Toolkit shortcut, default startup references point to Toolkit, legacy packages remain discoverable as fallback, no Windows Electron modification. Record commands/exit codes, matching paths, discovery entries, diff, HEAD. Missing target = BLOCKED; incorrect mappings = FAIL.

## Non-scope
No GUI interaction, user-setting deletion, legacy package removal or code edits.

## Dependencies
UMIG-008 code checkpoint, cutover approval/promotion and current JR packet.

## Sizing
Surface 0, environment 0, behavior 0, verification 1, recovery 0 = 1.
