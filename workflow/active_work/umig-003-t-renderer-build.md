# UMIG-003-T — TEST: Copied Renderer Build and Fixtures

Status: QUEUED — promoted 2026-09-18; wait for UMIG-003 source checkpoint.
Stage / owner: TEST / OpenHands (JR)
Previous: UMIG-003
Next: UMIG-003-V

## Primary outcome
Prove the copied Toolkit renderer builds with fixture data and no Electron dependency.

## Instruction / expected / evidence
In an isolated checkout of the coded SHA, run `cd OSJS && npm run build:local-packages && npm run package:discover` and the focused Toolkit fixture test command explicitly recorded by the coding agent in the current JR packet. Expected: commands pass, Toolkit bundle/assets present, fixture tabs do not require `window.mcsDesktop`, Electron directories or old OS.js UI imports. Return exact commands/exit codes, artifact/import inspection, HEAD and git-status before/after. Missing fixture command or environment = BLOCKED; actual failure = FAIL.

## Non-scope
No live backend, rendered parity claim, source changes or workflow advancement.

## Dependencies
UMIG-003 source checkpoint, approved donor SHA and an exact current JR TEST TASK packet. JR runs only this task when sole ACTIVE.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
