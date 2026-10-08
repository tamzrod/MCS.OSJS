# NPO-01 — OS.js Toolkit Persistence UI

Status: PROMOTED / QUEUED
Stage: CODE
Owner: Codex JR DEV
Previous: NPE-06 PASS
Next: NPO-02

## Purpose
After Electron native persistence is independently verified, implement the same native MMA2 persistence configuration UX in the OS.js Toolkit.

## Scope
Mirror the verified Electron contract, not the retired external persistence design:

- Enable/Disable
- Default storage location with optional directory override
- All Allocated Areas default
- optional Selected Ranges
- no persistence RBE controls
- no persistence lock coil
- no State Sealing dependency
- no restore/unseal controls

Preserve Toolkit authorization and existing configuration safety boundaries.

## Acceptance
Focused OS.js Toolkit UI/model tests cover defaults, custom ranges, optional directory, validation and reload parity with Electron/native MMA2 contract.

STOP after delivery.
