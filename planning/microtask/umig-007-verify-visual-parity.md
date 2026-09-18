# UMIG-007 — Verify Single-Window Electron UI Parity

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-006
Next: UMIG-008

## Primary outcome
Verify that the new OS.js Toolkit visually matches the approved Electron donor in one window.

## Scope
Compare rendered Memory, Replicator and Diagnostics tabs against screenshots of the pinned Electron build using the same fixture devices and window dimensions. Check tabs, forms, spacing, active/selected states, scrolling and that OS.js desktop/taskbar/window chrome remains intact. Record only concrete visual differences; correct visual defects through separately scoped work if found.

## Non-scope
No backend protocol changes, acceptance by screenshots alone for runtime functions, or old app removal.

## Acceptance
1. All three tab surfaces have recorded side-by-side parity evidence.
2. No Toolkit CSS leaks into OS.js shell surfaces.
3. Visual discrepancies are resolved or explicitly recorded as blockers before cutover.

## Verification
One rendered visual inspection workflow with matched Electron/OS.js dimensions and screenshots. Report precisely what was inspected; do not call it a runtime test.

## Dependencies
UMIG-006, approved donor snapshot and human promotion. UMIG-004/005/006 retain their own focused runtime gates.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
