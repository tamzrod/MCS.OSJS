# Electron Installer Settings Owner

## Boundary

Owns the Windows installer settings-access setup, not runtime API authorization or service-control security.
Parent / Zoom Out: INDEX.md. No children or required cross-tree connectors.

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
Scoped overlay: Current user / All users radio selection, automatic launch-account capture, updated helper scope checks and tests.
Other ICC nodes and workflow records retain their existing state.

## Facts

Setup offers Current user / All users without account entry. Current user is captured before NSIS UAC elevation and retrieved from the outer instance.
The two choices control settings access; the application and backend services remain system-wide.
Only explicit All users scope accepts the built-in Users group. Current user still requires an individual user SID.
It grants Modify only on ProgramData/MCS Modbus Toolkit/runtime/config and descendants.
The helper checks the exact settings path and rejects reparse points before calling icacls without a shell.
SYSTEM, administrators, binaries, service permissions and unrelated pre-existing ACL entries are not reset.
The scope/SID are retained in HKLM/Software/MCS Modbus Toolkit for repairs and upgrades.
Changing scope removes the prior installer-managed grant after granting the new principal.
New setup defaults to Current user; silent upgrades reuse the saved scope and SID.
The distribution build compiles the Windows Go helper before packaging NSIS.
This is not a complete per-user security boundary for runtime APIs.

## Verification Boundary

Helper tests exercise user/group lookup and settings-only ACL operations in temporary directories.
Node tests check installer source wiring, not interactive Setup behavior.
NSIS packaging is the required installer compilation gate; human distribution acceptance remains separate.
