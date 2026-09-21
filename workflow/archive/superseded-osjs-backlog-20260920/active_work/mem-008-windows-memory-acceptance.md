> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# MEM-008 — Windows Memory Acceptance

Status: QUEUED
Previous: MEM-007
Next: none

## Primary outcome
Verify the approved Memory tab and None/Random on a real Windows packaged Electron application.

## Scope
Build using `cd electron && npm run dist:win` with required local prerequisites. Verify Memory tab, Devices, Area, tight layout; MMA2/Replicator steady header, Simulator in Diagnostics; mixed modes, Save & Apply, persistence/relaunch, allocated memory readability with None, focus and resizing.

## Non-scope
No inference of Windows success from static source or Linux checks.

## Acceptance
1. Windows build/package and launch pass.
2. Labels/header/Diagnostics/focus pass hands-on UI test.
3. None/Random behavior and persisted values pass real memory reads.

## Verification
Return concrete Windows commands, screenshots or observations, and Modbus-read evidence. Do not archive without all gates passing.

## Dependencies
MEM-007; actual Windows machine.

## Sizing
Surface 0, environment 1, behavior 0, verification 2, recovery 0 = 3.
