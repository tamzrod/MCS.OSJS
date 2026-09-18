# UMIG-007 — VERIFY: Three-Tab Visual Parity

Status: QUEUED — promoted 2026-09-18; wait for UMIG-006-V PASS and donor snapshot.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-006-V
Next: UMIG-007A

## Primary outcome
Independently verify single-window OS.js Toolkit appearance against approved Electron donor.

## Verification action
In an isolated OS.js test desktop, open all three Toolkit tabs at matched dimensions with the same fixture devices as the approved donor screenshots; compare tabs, forms, spacing, selection, scrolling and OS.js desktop/taskbar/chrome. Do not equate visual inspection with backend success.
Expected: pinned donor appearance matches or concrete deviations are logged; CSS never affects OS.js shell. Evidence: pinned SHA, side-by-side screenshots or direct rendered observations per tab, dimensions, discrepancy list and current HEAD. Required unavailable visuals = BLOCKED; actual discrepancies = FAIL pending separately authorized coding fix.

## Non-scope
No code fixes, backend protocol changes, old UI removal or acceptance based on screenshots for runtime behavior.

## Dependencies
UMIG-006-V PASS, donor snapshot and exact JR TEST TASK.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
