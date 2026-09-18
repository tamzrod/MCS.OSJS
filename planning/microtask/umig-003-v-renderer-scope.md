# UMIG-003-V — VERIFY: Fixture Tabs and Shell Isolation

Status: PLANNED / BLOCKED — human promotion required.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-003-T
Next: UMIG-004

## Primary outcome
Observe three fixture-rendered Toolkit tabs in one OS.js window without shell CSS leakage.

## Instruction / expected / evidence
Launch Toolkit in an isolated OS.js GUI as specified in the JR packet. Visit Memory, Replicator and Diagnostics, inspect visible fixture controls, unknown/unavailable COMMS state, taskbar and window chrome, and absence of Electron runtime references. Expected: all three tabs render, shell stays intact and no unexplained green when evidence is missing. Return screenshots/direct observations, console errors, target, exact actions and HEAD. Unavailable GUI = BLOCKED, not PASS.

## Non-scope
No backend load/apply, final visual parity, product fixes or legacy removal.

## Dependencies
UMIG-003-T PASS, approved donor SHA and current JR packet.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
