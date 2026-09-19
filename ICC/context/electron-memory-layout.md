# Electron Memory Layout

## Boundary

Owns the standalone Electron Memory editor layout, not OS.js or workflow execution state.
Parent / Zoom Out: INDEX.md. No children or cross-tree connectors.

## Source Dependencies

- electron/renderer/app.js
- electron/renderer/style.css
- electron/renderer/memory-advanced.js
- electron/test/memory-settings.test.js

## Current Facts

The Memory editor has Device Definition and Advanced Settings folder strips.
Advanced Settings contains RBE Rules, State Sealing, and Access Policy.
Shared MMA configuration is opened through an MMA Settings popup inside Advanced Settings, not beside Devices.
Device and shared drafts remain separate. Closing the popup retains the shared draft; Discard restores its saved snapshot.
Shared saves use mma-apply IPC. A modal prevents editing the device behind it.
Renderer tests use a fake DOM and do not prove visual acceptance.

## Baseline / Overlay

Source baseline: dcc6f0d.
Audited overlay: the app.js modal and style.css folder-strip changes and popup regression test in memory-settings.test.js.
This node is scoped to that overlay; unrelated ICC branches retain their existing baselines.
Verification: node --test electron/test/*.test.js passed 31 tests; human visual acceptance remains pending.
