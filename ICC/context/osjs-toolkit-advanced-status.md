# OS.js Toolkit — advanced editor and communications status

Parent: [osjs-shell](osjs-shell.md)
Zoom Out: [osjs-shell](osjs-shell.md)
Zoom In: none
Connector: [osjs-shared-mma-settings](osjs-shared-mma-settings.md) for shared root configuration; follow only if required.

## Semantic boundary

Toolkit advanced-settings editor and rendered communications status only. This node does not own runtime process lifecycle, a third diagnostics transport, operator acceptance, or task state.

## Established facts from scoped grok UI repair

- Both canonical editors render Device Definition / Advanced Settings tabs; empty warning containers are hidden and actual warnings remain visible.
- Browser-only advanced editor handles RBE, state sealing and policy; Memory drafts use `mma2`, Replicator advanced drafts use `mma2_advanced` and require backend capability.
- Existing apply/discard paths preserve drafts on failure and restore persisted snapshots on discard.
- Replicator LEDs consume actual communications observations, not aggregate success. Missing or stale evidence remains UNKNOWN; poll generations reject stale device identity and timers are disposed on destroy.
- Header indicators describe selected-device observations, not host service health. Shared root editor and RBE settings live in the linked shared-settings node.
- Historical local production package build and mocked DOM checks passed; live Ubuntu/operator acceptance was not verified. These are not a fresh PASS for another checkout.

## Source dependencies and state

`OSJS/src/packages/MCSModbusToolkit/index.js`, `memory-editor.js`, `replicator-editor.js`, `advanced-editor.js`, `memory-advanced.js`, `comms-status.js`, `renderer.css`, `webpack.config.js`, `OSJS/tests/toolkit-ui-parity.test.js`.

Historical scoped source baseline: `0b7ee17` with audited uncommitted fingerprints in the prior version of this document (recoverable in Git history). The local overlay is not visible to this remote review; no assumption of clean state. On current checkout, compare the above dependencies with the historical audited source/fingerprints and refresh only changed files before treating these details as current. The entry-point wiring was separately observed at remote `main` `45cd3d2d81831306c6943f844e49b7deccbfc5ad`.
