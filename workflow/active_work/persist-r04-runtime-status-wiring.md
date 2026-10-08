# PERSIST-R04 — Runtime Status Wiring

Status: QUEUED — HUMAN PROMOTED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-UI01
Next: PERSIST-R05

## Primary outcome
Feed real runtime save/restore observations into the existing persistence status projection and operator diagnostics.

## Scope
- Add the real non-test runtime call site for `PersistenceRuntimeStatusFromPlan`.
- Carry actual last-save and last-restore observations from PERSIST-R02/R03 into `DeviceRuntimeStatus.Persistence`.
- Keep status observational: restore failure remains visibly sealed and the UI gains no control authority.

## Non-scope
No persistence writes/restores in the status layer, no new retry/control path, no historian metrics, no UI bypass, no ICC edits.

## Acceptance
1. Runtime status reflects actual configured/snapshot/restore/sealed/healthy state after real save/restore activity rather than configured-only placeholders.
2. Last save/restore observations appear when available and remain absent/unknown when genuinely unobserved.
3. Diagnostics remains read-only and fail-closed for malformed/stale observations.

## Evidence / handoff
Record real status call sites, changed paths, focused Go/JS tests, bounded regression and delivered source SHA. After genuine delivery, automatically arm PERSIST-R05 for the next invocation and STOP.

## Dependencies
Requires genuine PERSIST-UI01 and PERSIST-R03 delivery.

## Sizing
1/0/1/1/0=3.

## CWAL
Already human-promoted. Execute only when selected by handoff. After delivery, automatically activate PERSIST-R05 and STOP.
