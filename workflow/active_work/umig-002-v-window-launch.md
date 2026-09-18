# UMIG-002-V — VERIFY: One Toolkit Window

Status: QUEUED — human authorized independent OpenHands verification; not run.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-002-T
Next: none

## Primary outcome
Observe the installed/discovered placeholder opening as one OS.js window without disturbing the existing desktop.

## Verification action
In a disposable OS.js test session built by UMIG-002-T, start OS.js with its repository-native serve action, open its existing Start menu, launch `MCS Modbus Toolkit` once, and observe the resulting window. If an alternative launch method is needed, stop and request an exact packet update; do not invent one.
Expected: exactly one Toolkit window with the clearly marked disconnected placeholder; OS.js desktop, Start menu and taskbar remain functional, and legacy Simulator/Replicator applications remain present. No backend health is claimed.
Evidence: exact start/launch steps, running URL/target, screenshot or direct rendered observation, window count, legacy package availability, any error logs and HEAD. An unavailable GUI is BLOCKED, not PASS.

## Non-scope
No functional Memory/Replicator/Diagnostics acceptance, icon cutover, deletion, product fixes or workflow advancement.

## Dependencies
UMIG-002-T reported PASS, coding agent reviewed evidence and authored a current `JR TEST TASK`.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
