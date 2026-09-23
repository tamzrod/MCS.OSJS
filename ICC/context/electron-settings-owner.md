# Electron Installer Settings Owner

## Boundary

Owns Windows installer settings-access setup, not runtime API authorization or service-control security.
Parent / Zoom Out: [L0-project](L0-project.md). No children or required cross-tree connectors.

## Source Dependencies

- electron/build/installer.nsh
- electron/build/settings-access/go.mod
- electron/build/settings-access/main_windows.go
- electron/build/settings-access/main_windows_test.go
- electron/test/installer-permissions.test.js
- electron/package.json
- electron/SETTINGS_PERMISSIONS.md

## Baseline / Overlay

Source baseline: 449cfda.
Scoped historical overlay: Current user / All users radio selection, automatic launch-account capture, updated helper scope checks and tests. This routing-only update does not re-audit current source or local overlay.

## Facts

Setup offers Current user / All users without account entry. Current user is captured before NSIS UAC elevation and retrieved from the outer instance. Both choices control settings access; app and backend services stay system-wide. Only explicit All users scope accepts the built-in Users group, while Current user requires an individual user SID.
The helper grants Modify only on ProgramData/MCS Modbus Toolkit/runtime/config and descendants, checks the exact path, and rejects reparse points before invoking icacls without a shell. SYSTEM, administrators, binaries, service permissions and unrelated pre-existing ACL entries are not reset. Scope/SID persist in HKLM/Software/MCS Modbus Toolkit for repairs and upgrades. Changing scope removes the former installer-managed grant after granting the new principal. Setup defaults to Current user; silent upgrades reuse the saved scope/SID. Distribution packaging builds the Windows helper before NSIS. This is not a complete per-user runtime API security boundary.

## Historical Verification Boundary

Helper tests exercised temporary settings-only ACL operations and Node tests checked installer wiring, not interactive Setup. NSIS compilation and human distribution acceptance remain separate; no test was executed by this ICC routing edit.
