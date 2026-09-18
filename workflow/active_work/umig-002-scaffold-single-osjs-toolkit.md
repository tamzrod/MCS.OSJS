# UMIG-002 — Scaffold One OS.js Toolkit Package

Status: ACTIVE — human redirected work to OS.js Toolkit migration on 2026-09-18. Scaffold source staged; build/discovery and launch verification remain pending.
Previous: none (human-approved early staging before donor freeze)
Next: UMIG-001 (requires separate human promotion when a UI snapshot is selected)

## Primary outcome
Create a buildable OS.js package that opens one MCS Modbus Toolkit window.

## Scope
Add minimal metadata, OS.js-owned icon, build configuration, and one-window placeholder bootstrap in `OSJS/src/packages/`, following existing package patterns. Keep the current desktop shell, legacy application packages, desktop shortcuts, backend services and user configurations unchanged. Do not import Electron code or expose unfinished Memory/Replicator/Diagnostics as working.

## Non-scope
No Electron renderer copy or final donor approval, IPC, backend behavior, desktop launcher cutover, shared build pipeline or deletion of legacy apps.

## Acceptance
1. OS.js local-package build and discovery recognize `MCSModbusToolkit` without Electron installation or build.
2. The package opens one placeholder OS.js window without removing or changing existing apps.

## Verification
From the OS.js project, run `npm run build:local-packages` and `npm run package:discover`, then launch `MCSModbusToolkit` in an isolated OS.js test session. Record build/discovery and rendered launch results separately. Source presence or syntax alone does not close this task. If runtime access is unavailable, report BLOCKED and leave ACTIVE.

## Dependencies
Human direction on 2026-09-18 authorizes starting the independent placeholder before choosing the Electron donor snapshot. UMIG-001 remains the separate gate for UI copying. Do not advance to it automatically until human promotion.

## Sizing
Surface 1, environment 1, behavior 0, verification 1, recovery 0 = 3.
