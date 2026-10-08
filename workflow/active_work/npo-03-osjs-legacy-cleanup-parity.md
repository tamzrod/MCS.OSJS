# NPO-03 — OS.js Toolkit Legacy Persistence Cleanup / Parity

Status: PROMOTED / QUEUED
Stage: CODE
Owner: Codex JR DEV
Previous: NPO-02
Next: NPO-04

## Purpose
Remove stale Toolkit assumptions from the previous external persistence architecture and establish parity with the verified Electron experience.

## Scope
Remove obsolete persistence-specific State Sealing/RBE/restore/status assumptions only where they are no longer valid. Preserve independent State Sealing, RBE and generic diagnostics.

Update persistence diagnostics wording to describe native MMA2 ownership rather than external restore/unseal lifecycle.

## Acceptance
Toolkit focused/full Node tests pass; persistence UI/config semantics match the Electron/native contract.

STOP after delivery.
