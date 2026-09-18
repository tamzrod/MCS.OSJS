# UMIG-009 — Retire Legacy OS.js Modbus UI Packages

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-008
Next: none

## Primary outcome
Remove superseded separate Modbus UI packages from OS.js application discovery after verified Toolkit cutover.

## Scope
Only after one-window launcher, Memory, Replicator and Diagnostics checks pass, remove obsolete OS.js package discovery/build references and legacy UI package files that Toolkit no longer uses. Ensure any still-required OS.js server bridge code has been migrated into Toolkit before deletion. Keep a known-good commit for rollback. Electron's standalone Windows source, packaging and deployment must remain untouched; the OS.js Toolkit must retain its own copied UI files and build pipeline.

## Non-scope
Do not remove OS.js desktop/theme/icons, Go services, MMA2 configuration, user data or Electron app. No runtime protocol or storage migration, shared renderer, or cross-deployment coupling.

## Acceptance
1. Toolkit is the only MCS Modbus application shown by fresh OS.js package discovery.
2. Simulator/Replicator communication still works after old UI removal, without any Electron installation.
3. User configuration and backend services are unchanged, with a documented rollback commit.

## Verification
Run OS.js package build/discovery and a post-removal Toolkit smoke test against its own deployed services; inspect the changed file list to confirm no data, runtime or Electron files were touched. Preserve the successful UMIG-007A independent-deployment evidence.

## Dependencies
UMIG-008 completed; recorded UI parity, independent-deployment gate and each tab's focused runtime verification; explicit human promotion. Any failed acceptance blocks deletion.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 1 = 4 (one legacy UI retirement boundary).
