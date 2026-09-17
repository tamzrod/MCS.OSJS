# RREC-001 — Canonical Windows Data Root

Status: COMPLETE
Previous: RLED-002
Next: RREC-002

## Primary outcome
Make ProgramData the explicit shared runtime root used by installer initialization and all Windows services.

## Scope
Resolve ProgramData explicitly for runtime directories, initial MMA2 config and NSSM AppDirectory. Preserve existing files and idempotent repair behavior.

## Non-scope
No config deletion, migration from unrelated roots, protocol or UI change.

## Acceptance
1. Installer and services explicitly target ProgramData.
2. Existing configuration is never overwritten.
3. Fresh and repair installer scripts remain idempotent.

## Verification
Static installer inspection, installer build, and installed-path check in RREC-004.

## Dependencies
RLED-002.

## Sizing
Surface 1, environment 1, behavior 1, verification 1, recovery 0 = 4; tightly coupled installer-path repair.
## Completion evidence
- The installer resolves C:\ProgramData\MCS Modbus Toolkit\runtime once and reuses it for configuration and NSSM AppDirectory.
- Existing MMA2 configuration remains protected by IfFileExists.
- npm run dist:win passed on 2026-09-17.