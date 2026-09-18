# UMIG-004 — CODE: Connect Toolkit Memory Adapter

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: CODE / ChatGPT
Previous: UMIG-003-V
Next: UMIG-004-T

## Primary outcome
Author the Toolkit-owned OS.js Simulator messaging adapter for Memory load/apply/status.

## Scope
Reuse or relocate the existing Simulator OS.js server bridge into Toolkit-owned code so retiring old UI does not break it. Map backend responses/errors to copied Memory renderer, preserve saved definitions and None/Random behavior, make unavailable states explicit.

## Non-scope
No live testing, protocol/MMA2 redesign, Windows pipe use, other tab, old UI deletion or fabricated status.

## Coding acceptance / handoff
1. Toolkit has a self-owned Simulator bridge and Memory adapter for load/apply/status.
2. Existing semantics and error handling are retained without legacy-package import.
3. Record exact files/commit and message contract for OpenHands without claiming a test pass.

## Dependencies
UMIG-003-V PASS and human promotion.

## Sizing
Surface 2, environment 0, behavior 1, verification 0, recovery 0 = 3.
