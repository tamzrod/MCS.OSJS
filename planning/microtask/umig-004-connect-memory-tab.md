# UMIG-004 — Connect Toolkit Memory Tab to Simulator Runtime

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-003
Next: UMIG-005

## Primary outcome
Make the copied Memory tab use the existing OS.js Simulator runtime contract.

## Scope
Implement the Toolkit-owned OS.js host adapter for Simulator load/apply/status, reusing or relocating the existing OS.js simulator message bridge without depending on a package slated for deletion. Map responses and errors to the copied renderer's expected shapes. Keep the approved Electron UI and final None/Random semantics; make status failures visible and preserve saved configuration.

## Non-scope
No new simulator protocol, MMA2 redesign, Windows named pipes, Replicator wiring, Diagnostics or UI re-layout.

## Acceptance
1. Memory loads existing definitions through OS.js runtime messaging.
2. Save & Apply and selected-device status use the existing backend and report real success/failure.
3. None/Random and saved settings are not silently changed by the adapter.

## Verification
Focused Simulator adapter fixtures followed by an OS.js Memory load/apply/status test against the existing runtime; record results without modifying unrelated backends.

## Dependencies
UMIG-003 and human promotion; existing Simulator runtime must be available for live verification.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 1 = 4 (one Simulator-only contract and verification branch).
