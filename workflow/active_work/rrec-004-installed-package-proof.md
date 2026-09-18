# RREC-004 — Installed Package Proof

Status: QUEUED — PAUSED by human switch to OS.js Toolkit migration on 2026-09-18. Not completed or retired; Windows installer evidence remains unverified.
Previous: RREC-003
Next: RLED-003 (legacy Windows chain only; do not automatically advance while paused)

## Primary outcome
Prove the installer contains matching current Electron and Go runtime artifacts and preserves shared configuration.

## Scope
Build both Go backends and NSIS installer, compare hashes, inspect package contents, and perform safe repair verification with configuration backup.

## Non-scope
No LED acceptance; RLED-011 remains separate Windows communications acceptance. No OS.js migration work under this task.

## Acceptance
1. Bundled backend hashes match freshly built binaries.
2. Installed Electron and services use the same ProgramData root.
3. Repair preserves existing configuration.

## Verification
Artifact hashes, service configuration, named-pipe query and non-destructive installed repair evidence. Prior package builds or screenshots alone do not satisfy these gates.

## Dependencies
RREC-003; explicit human decision to resume Windows installer verification. Do not infer completion from the user's acceptance of the LED appearance.

## Sizing
Surface 1, environment 2, behavior 0, verification 2, recovery 1 = 6; must be split under the current microtask rules before resuming if the prescribed proof cannot be performed as one bounded workflow.
