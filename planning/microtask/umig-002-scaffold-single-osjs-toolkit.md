# UMIG-002 — Scaffold One OS.js Toolkit Package

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-001
Next: UMIG-003

## Primary outcome
Create a buildable OS.js package that opens one MCS Modbus Toolkit window.

## Scope
Add the minimal package metadata, build configuration and window bootstrap in `OSJS/src/packages/`, following existing OS.js package patterns. Keep the existing desktop shell, current application packages and shortcuts unchanged. A placeholder content root is enough; no renderer import or runtime calls yet.

## Non-scope
No Electron IPC, copied UI, backend behavior, launcher cutover or deletion of legacy apps.

## Acceptance
1. Package build/discovery recognizes the new Toolkit package.
2. The package mounts one OS.js window with a placeholder content root, without removing existing apps.

## Verification
Run the existing local-package build/discovery gate for the new package; launch the placeholder in an isolated OS.js test environment. Record build and launch results separately; do not claim installed acceptance.

## Dependencies
UMIG-001 completed, then human promotion of UMIG-002. Preserve existing OS.js shell and active work.

## Sizing
Surface 1, environment 1, behavior 0, verification 1, recovery 0 = 3.
