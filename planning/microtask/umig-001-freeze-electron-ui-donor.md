# UMIG-001 — Freeze the Final Electron UI Donor

Status: PLANNED / BLOCKED — wait for human confirmation that Electron UI is finished and accepted. Not ACTIVE.
Previous: none
Next: UMIG-002

## Primary outcome
Record one approved, immutable Electron renderer commit as the OS.js UI migration reference.

## Scope
After the human declares Electron UI finished, record the exact source commit and the final Memory, Replicator and Diagnostics tab surfaces, including renderer HTML/CSS/JS and any UI assets. Identify native-only calls that need an OS.js adapter. Record the reference in a small migration donor note; do not use an older in-progress renderer.

## Non-scope
No Electron modifications, OS.js port, workflow promotion, UI redesign or runtime changes.

## Acceptance
1. Human approval of the finished Electron UI and the exact commit SHA are recorded.
2. The donor note identifies the final renderer files/assets and native-only API boundary.

## Verification
Check the approved commit and inspect only the listed donor paths; confirm the donor note points to that SHA.

## Dependencies
Hard gate: Electron UI completion, its remaining acceptance/verification as determined by the human, and explicit human selection of this task for promotion. Do not infer completion from packaging or a prior partial test. Before promotion, validate the current planning ICC branch through BLACK SHEEP WALL if stale.

## Sizing
Surface 1, environment 0, behavior 0, verification 1, recovery 0 = 2.
