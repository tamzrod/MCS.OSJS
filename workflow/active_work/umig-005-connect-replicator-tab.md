# UMIG-005 — CODE: Connect Toolkit Replicator Adapter

Status: ACTIVE — promoted 2026-09-19 after ChatGPT reviewed UMIG-004-V direct live behavior and archived it with a documented test-command deviation. CODE pending; no Replicator integration or tests for this task performed.
Stage / owner: CODE / ChatGPT
Previous: UMIG-004-V (COMPLETE — live behavioral VERIFY accepted with explicit packet restart defect, `workflow/archive/umig-004-v-memory-runtime.md`)
Next: UMIG-005-T (QUEUED)

## Primary outcome
Author the Toolkit-owned OS.js Replicator adapter for load/apply/status/suggest.

## Scope
Move or reuse the existing OS.js Replicator bridge inside Toolkit ownership, retaining one Go apply transaction, destination ownership checks, Pull Block validation, actual errors and available per-block status. Reuse the Docker-hosted Go service and its Unix-socket endpoint. Do not fabricate green COMMS observations unsupported by the backend. The existing dormant `replicator-contract.js` is a source-only preparation, not verified or wired.

## Non-scope
No live tests, new protocol, MMA2 composer changes, Windows service calls, Memory changes, UI redesign, production deployment, ICC editing or legacy UI retirement. Do not mistake the Memory VERIFY for Replicator acceptance.

## Coding acceptance / handoff
1. Toolkit owns load/apply/status/suggest bridge and adapter without legacy-package dependency.
2. Ownership/validation and genuine unknown/error behavior remain explicit.
3. Record exact changed files/commit and expected response mapping for OpenHands; source-only CODE, no build, unit, GUI or live PASS claims. After source checkpoint, archive CODE and promote only the explicit UMIG-005-T successor with a current JR packet.

## Dependencies
UMIG-004-V reviewer-accepted live behavioral gate documented in archive; approved available backend must be checked before writes; execute only while sole ACTIVE.

## Sizing
Surface 2, environment 0, behavior 1, verification 0, recovery 0 = 3.
