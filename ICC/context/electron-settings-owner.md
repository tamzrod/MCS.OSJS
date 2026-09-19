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

Source baseline: 1ca1740.
Scoped overlay: account-selection page, settings-access helper, distribution documentation, and tests.
Other ICC nodes and workflow records retain their existing state.

## Facts

Setup explicitly requests a Windows user account, resolves it to a user SID, and rejects groups.
It grants Modify only on ProgramData/MCS Modbus Toolkit/runtime/config and descendants.
The helper checks the exact settings path and rejects reparse points before calling icacls without a shell.
SYSTEM, administrators, binaries, service permissions and unrelated pre-existing ACL entries are not reset.
The account/SID are retained in HKLM/Software/MCS Modbus Toolkit for repairs and upgrades.
Changing the selected account removes the prior installer-managed grant after granting the new account.
Silent setup without a saved owner fails before installation; silent upgrades reuse the saved SID.
The distribution build compiles the Windows Go helper before packaging NSIS.
This is not a complete per-user security boundary for runtime APIs.

## Verification Boundary

Helper tests exercise user/group lookup and settings-only ACL operations in temporary directories.
Node tests check installer source wiring, not interactive Setup behavior.
NSIS packaging is the required installer compilation gate; human distribution acceptance remains separate.
