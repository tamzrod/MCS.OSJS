> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# UMIG-009-V — VERIFY: Toolkit After Legacy UI Retirement

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-009-T
Next: none

## Primary outcome
Prove the replacement Toolkit works and preserved services/data after old OS.js UI retirement.

## Instruction / expected / evidence
Use isolated OS.js deployment and non-destructive test devices/config. Launch Toolkit from the sole desktop icon; inspect Memory, Replicator and Diagnostics; exercise safe read/status paths to the existing simulator and replicator services. Confirm unchanged config/backend service state and known-good rollback commit. Expected: live bridges continue without legacy UI or Electron, user data untouched, Start menu/taskbar intact. Return actual UI/backend observations, before/after state, changed paths, rollback SHA and HEAD. Missing live services or safety proof = BLOCKED.

## Non-scope
No production resets, backend removal, silent bug fixes or automatic further work.

## Dependencies
UMIG-009-T PASS, earlier runtime and independent-deployment gates PASS, current JR packet.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
