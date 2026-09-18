# UMIG-002 — Scaffold One OS.js Toolkit Package

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-001
Next: UMIG-003

## Primary outcome
Create a buildable OS.js package that opens one MCS Modbus Toolkit window.

## Scope
Add the minimal package metadata, build configuration and window bootstrap in `OSJS/src/packages/`, following existing OS.js package patterns. Keep the existing desktop shell, current application packages and shortcuts unchanged. A placeholder content root is enough; no renderer import or runtime calls yet. The OS.js Toolkit belongs to the OS.js build and deployment only: it must not require the Electron package, installation, renderer source directory, build output or native dependencies.

## Non-scope
No Electron IPC, copied UI, backend behavior, launcher cutover, shared build pipeline or deletion of legacy apps.

## Acceptance
1. OS.js package build/discovery recognizes the new Toolkit package without running an Electron build or installing Electron.
2. The package mounts one OS.js window with a placeholder content root, without removing existing apps.

## Verification
Run the existing OS.js local-package build/discovery gate independently of Electron; launch the placeholder in an isolated OS.js test environment. Record build and launch results separately; do not claim installed acceptance.

## Dependencies
UMIG-001 completed, then human promotion of UMIG-002. Preserve existing OS.js shell and active work.

## Sizing
Surface 1, environment 1, behavior 0, verification 1, recovery 0 = 3.
