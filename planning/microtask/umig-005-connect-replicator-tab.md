# UMIG-005 — Connect Toolkit Replicator Tab to Runtime

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-004
Next: UMIG-006

## Primary outcome
Make the copied Replicator tab use the existing OS.js Replicator runtime contract.

## Scope
Implement the Toolkit-owned OS.js adapter for Replicator load/apply/status/suggest using the existing OS.js message bridge, without retaining a dependency on a package planned for retirement. Preserve the final approved Electron editor, but do not discard backend ownership checks, required-block constraints, errors or any verified per-block status behavior. Surface actual runtime errors rather than reporting fabricated healthy state.

## Non-scope
No new Replicator protocol, MMA2 composer changes, Windows service controls, Memory changes or UI redesign.

## Acceptance
1. Existing Replicator definitions load in Toolkit.
2. Save & Apply goes through one authoritative backend transaction with real errors exposed.
3. Selected-device status and available suggestion/ownership responses are truthful.

## Verification
Focused adapter fixture tests and a Replicator-only OS.js load/apply/status test with a real runtime. Preserve current configs and record any existing backend test blockers separately.

## Dependencies
UMIG-004, human promotion, and available verified Replicator backend. Electron UI completion does not imply backend readiness.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 1 = 4 (Replicator-only runtime surface).
