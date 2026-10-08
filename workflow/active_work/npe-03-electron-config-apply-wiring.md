# NPE-03 — Electron Native Persistence Config / Apply Wiring

Status: PROMOTED / QUEUED
Stage: CODE
Owner: Codex JR DEV
Previous: NPE-02
Next: NPE-04

## Purpose
Make Electron Save & Apply write only the native MMA2 persistence configuration.

## Scope
Wire the Electron Simulator memory document through compose/validate/save so the resulting MMA2 memory contains the native per-memory `persistence` block exactly as configured.

Save & Apply responsibilities:
1. validate configuration;
2. commit configuration safely;
3. request/reach normal MMA2 restart/readiness using existing generic configuration lifecycle;
4. report configuration apply result.

Save & Apply MUST NOT:
- capture its own snapshot;
- restore snapshot over Modbus/Raw Ingest;
- write an unlock coil;
- poll 0x06 for persistence;
- own a persistence watchdog;
- depend on State Sealing;
- depend on RBE TCP.

MMA2 owns native persistence startup and runtime behavior.

## Acceptance
Focused Electron/main-process tests prove:
- enabled/default emits native persistence with no forced directory/ranges;
- custom directory/ranges round-trip correctly;
- disabled removes/disables only persistence config;
- unrelated memory settings are preserved;
- Save & Apply contains no external persistence restore/unseal lifecycle.

No OS.js Toolkit changes. STOP after delivery.
