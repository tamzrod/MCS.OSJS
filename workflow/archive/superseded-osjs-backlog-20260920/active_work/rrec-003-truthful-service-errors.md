> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# RREC-003 — Truthful Service and Runtime Errors

Status: COMPLETE
Previous: RREC-002
Next: RREC-004

## Primary outcome
Stop displaying false service health and expose actionable Replicator runtime errors.

## Scope
Query all actual Windows service states; keep runtime status fail-closed; show selected device and underlying pipe/runtime error in diagnostics/logs.

## Non-scope
No communications LEDs or service restart control.

## Acceptance
1. Missing/stopped services are never green.
2. Runtime errors remain visible and distinguishable.
3. Successful status clears stale error diagnostics.

## Verification
Focused Node/status tests and packaged UI inspection.

## Dependencies
RREC-002.

## Sizing
Surface 2, environment 0, behavior 2, verification 1, recovery 0 = 5; tightly coupled truth surfaces.
## Completion evidence
- All three backend labels now come from actual Windows service queries.
- Replicator pipe/runtime failures show the selected device and underlying error in both the status area and diagnostics log.
- Repeated identical poll errors are deduplicated; a successful poll clears the stale runtime error.
- Focused Node tests passed: 11/11.
- npm run pack:win passed, and packaged app.asar contains main.js, renderer/app.js, and runtime-status.js.
