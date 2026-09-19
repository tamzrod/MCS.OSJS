# Windows settings owner

Interactive Setup asks for the individual Windows account that will use the app, such as COMPUTER\user or DOMAIN\user. No account is inferred from the elevated installer identity. Windows resolves the selection to a SID; nonexistent accounts, groups, and service identities are rejected.

Setup grants that SID Modify access only to %ProgramData%\MCS Modbus Toolkit\runtime\config, including existing settings and inherited access for future service-created files. Program Files and service-control permissions are unchanged. SYSTEM and administrators retain their existing access.

The account and SID are stored under HKLM\Software\MCS Modbus Toolkit for upgrades and repairs. Selecting another owner removes the previous installer-managed account grant after adding the new one. Other pre-existing ACL entries are not reset. This is settings-write authorization, not a complete per-user access-control boundary for all runtime APIs.

Fresh installations and older installations without a saved owner require interactive Setup. Silent upgrades reuse the saved SID; no blanket Users or Everyone grant is used. A deleted owner account must be replaced through interactive repair.

Permission or account-validation failures stop Setup. Redirected settings directories or entries are rejected. After installation, reopen the desktop normally and verify Save & Apply without Run as administrator.

Build: npm run dist:win in electron (requires Go and Windows; builds the account helper before NSIS packaging).
Helper tests: go test ./... in electron/build/settings-access.
Renderer/installer contract tests: node --test electron/test/*.test.js from the repository root.
