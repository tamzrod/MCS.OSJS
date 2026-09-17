# RREC-004 — Installed Package Proof

Status: QUEUED
Previous: RREC-003
Next: RLED-003

## Primary outcome
Prove the installer contains matching current Electron and Go runtime artifacts and preserves shared configuration.

## Scope
Build both Go backends and NSIS installer, compare hashes, inspect package contents, and perform safe repair verification with configuration backup.

## Non-scope
No LED acceptance; RLED-011 remains final communications UI acceptance.

## Acceptance
1. Bundled backend hashes match freshly built binaries.
2. Installed Electron and services use the same ProgramData root.
3. Repair preserves existing configuration.

## Verification
Artifact hashes, service configuration, named-pipe query and non-destructive installed repair evidence.

## Dependencies
RREC-003.

## Sizing
Surface 1, environment 2, behavior 0, verification 2, recovery 1 = 6; explicit environment acceptance task.