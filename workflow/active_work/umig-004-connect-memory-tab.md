# UMIG-004 — CODE: Connect Toolkit Memory Adapter

Status: ACTIVE — promoted by ChatGPT after reviewed UMIG-003-V independent real-browser PASS at `6f82196700b1c312652fdd7584a508e5452c4822`, 2026-09-19. CODE not yet implemented; TEST/VERIFY pending.
Stage / owner: CODE / ChatGPT
Previous: UMIG-003-V (COMPLETE; `workflow/archive/umig-003-v-renderer-scope.md`)
Next: UMIG-004-T (QUEUED; only after source checkpoint and handoff)

## Primary outcome
Author the Toolkit-owned OS.js Simulator messaging adapter for Memory load/apply/status.

## Scope
Reuse or relocate the existing Simulator OS.js server bridge into Toolkit-owned code so retiring old UI does not break it. Map backend responses/errors to copied Memory renderer, preserve saved definitions and None/Random behavior, make unavailable states explicit. Use existing Docker-shared Unix-socket transport and services rather than Windows pipe calls. Reuse dormant `OSJS/src/packages/MCSModbusToolkit/memory-contract.js` only after reviewing it against real Simulator v1 contracts; its authored unit test remains unrun.

## Non-scope
No live testing, protocol/MMA2 redesign, Windows pipe use, other tab, old UI deletion or fabricated status. Do not modify operator's working Docker deployment or persisted config. A CODE checkpoint is not permission to apply real user configuration or claim backend behavior verified.

## Coding acceptance / handoff
1. Toolkit has a self-owned Simulator bridge and Memory adapter for load/apply/status.
2. Existing semantics and error handling are retained without legacy-package import.
3. Record exact files/commit and message contract for OpenHands without claiming a test pass.

## Dependencies
Satisfied: UMIG-003-V independent rendered PASS reviewed and archived, sole ACTIVE promotion. UMIG-004-T remains QUEUED until this CODE stage meets its source-specific gate.

## Sizing
Surface 2, environment 0, behavior 1, verification 0, recovery 0 = 3.
