# UMIG-005 — CODE: Connect Toolkit Replicator Adapter

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: CODE / ChatGPT
Previous: UMIG-004-V
Next: UMIG-005-T

## Primary outcome
Author the Toolkit-owned OS.js Replicator adapter for load/apply/status/suggest.

## Scope
Move or reuse the existing OS.js Replicator bridge inside Toolkit ownership, retaining one Go apply transaction, destination ownership checks, Pull Block validation, actual errors and available per-block status. Do not fabricate green COMMS observations unsupported by the backend.

## Non-scope
No live tests, new protocol, MMA2 composer changes, Windows service calls, Memory changes or UI redesign.

## Coding acceptance / handoff
1. Toolkit owns load/apply/status/suggest bridge and adapter without legacy-package dependency.
2. Ownership/validation and genuine unknown/error behavior remain explicit.
3. Record exact changed files/commit and expected response mapping for OpenHands; no test claims.

## Dependencies
UMIG-004-V PASS, approved available backend and human promotion.

## Sizing
Surface 2, environment 0, behavior 1, verification 0, recovery 0 = 3.
