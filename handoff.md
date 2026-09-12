# Handoff

## Current

ACTIVE:
- ELECTRON-002 — Compact Industrial Desktop Layout

QUEUED:
- ELECTRON-001 — NSIS + NSSM Service Installer — implementation authored; Windows install/repair/uninstall acceptance still required.
- REP-BLOCK-002 — Independent Pull Block Pollers — implemented; JR retest pending.
- REP-BLOCK-003 — Device / Pull Blocks Folder Tabs — implemented; JR rendered retest pending.

## ELECTRON-002 intent

The MCS desktop UI should follow the mentality of classic Modbus engineering utilities rather than a modern dashboard: every pixel should carry information or control, with compact inputs/buttons, narrow sidebars, dense tables, thin status areas, and little decorative whitespace.

The user's tested Electron screenshot contains the real Simulator editor, but `electron/renderer/index.html` on `main` is still the older placeholder shell. The visible Simulator layout in the screenshot matches the canonical UI in `OSJS/src/packages/ModbusSimulator`, and the Replicator has the corresponding canonical UI under `OSJS/src/packages/ModbusReplicator`.

Therefore this task intentionally tightens both surfaces:
- Electron shell/window sizing and chrome in `electron/`;
- canonical Simulator/Replicator density in `OSJS/src/packages/`.

Do not overwrite or discard a newer local/uncommitted Electron UI integration when pulling these commits. Reconcile the compact styles into that integration if it has not yet been pushed to `main`.

## Implemented in ELECTRON-002

- Electron default window reduced from 1120 x 760 to 900 x 560.
- Electron minimum size reduced from 900 x 600 to 760 x 460.
- Shell title/status strip, runtime LEDs, top tabs, content padding, diagnostics controls, and log spacing are compacted.
- Simulator sidebar reduced to 190 px, with 160 px compact breakpoint.
- Simulator inputs/buttons/device rows/runtime row/identity grid/FC table/status line are compacted.
- Simulator editor actions are left-aligned like a classic desktop utility.
- Replicator receives the same compact sidebar/control/status density.
- Replicator Device/Pull Blocks folder tabs and spreadsheet rows are shortened and narrowed while preserving the existing table-first interaction model.
- No backend calls, persistence semantics, runtime protocols, or polling behavior were changed by this pass.

## Verification status

Static repository verification completed by re-reading the modified sources on `main` after the edits:
- `electron/main.js` contains the 900 x 560 default and 760 x 460 minimum window.
- `electron/renderer/style.css` contains the compact shell chrome/status/tab/control spacing.
- `OSJS/src/packages/ModbusSimulator/index.scss` contains the 190 px sidebar, 22 px controls, 27 px FC rows, and compact status/actions.
- `OSJS/src/packages/ModbusReplicator/index.scss` contains the corresponding compact sidebar, controls, folder tabs, and spreadsheet rows.

Final acceptance is rendered Windows behavior:
1. build/run the Electron package;
2. confirm the normal window is useful at roughly 900 x 560 without large blank regions;
3. confirm Simulator Name/Enabled/Port/Unit plus FC1-FC4 remain visible and usable;
4. confirm Replicator Device/Pull Blocks remains usable at the smaller size;
5. confirm dropdowns/inputs keep focus while runtime/activity status updates continue;
6. confirm enlarging the window still lays out correctly.

Do not mark ELECTRON-002 complete until that rendered Windows check passes.

## Queued verification notes

ELECTRON-001 still needs the Windows installer sequence: clean Install -> run the same Setup again -> Repair/reconfigure -> run Setup again -> Uninstall.

REP-BLOCK-002/003 remain queued for their previous JR/human retest. The last JR run stopped at the first gate because `replicator/reader_test.go` was not gofmt-clean; later backend/deploy/rendered/e2e checks were therefore not executed in that run.
