# NPE-05 — Electron Legacy Persistence Cleanup

Status: PROMOTED / QUEUED
Stage: CODE
Owner: Codex JR DEV
Previous: NPE-04
Next: NPE-06

## Purpose
Remove the obsolete external persistence implementation from the Electron product after native config wiring is proven.

## Scope
Remove only code proven obsolete because MMA2 now owns persistence runtime, including as applicable:

- Electron snapshot capture/restore helper;
- Electron Raw Ingest persistence restore;
- persistence unlock-coil writes;
- persistence-specific 0x06 watchdog logic;
- persistence-generated/system RBE behavior;
- persistence-specific State Sealing ownership/prerequisites;
- old filesystem snapshot orchestration outside MMA2;
- obsolete tests for those retired behaviors.

Keep:
- ordinary RBE;
- ordinary State Sealing;
- Raw Ingest as an independent feature;
- generic MMA2 restart/apply lifecycle;
- user data/snapshots;
- unrelated diagnostics and settings.

Never delete user snapshot files as part of cleanup.

## Acceptance
Electron focused/full Node tests and relevant Go config tests pass; source search shows no Electron-owned persistence runtime/restore engine remains.

No OS.js Toolkit migration in this packet. STOP after delivery.
