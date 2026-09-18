# UMIG-003 — CODE: Copy Approved Renderer into OS.js

Status: ACTIVE — UMIG-001 donor decision complete and archived 2026-09-19; source implementation NOT STARTED.
Stage / owner: CODE / ChatGPT
Previous: UMIG-001 (COMPLETE, `workflow/archive/umig-001-freeze-electron-ui-donor.md`)
Next: UMIG-003-T (QUEUED; no auto-activation before CODE source checkpoint)

## Primary outcome
Author a self-contained, OS.js-owned copy of the approved three-tab Toolkit renderer.

## Frozen donor and source provenance

Human-approved Electron renderer commit: `1c971b9a6e00bafadf329df8821421a40cfc079c`. Read `workflow/archive/umig-001-freeze-electron-ui-donor.md` for the exact `electron/renderer/index.html`, `style.css`, `app.js`, `comms-status.js` and optional local SVG/PNG asset inventory, host/API boundary and scoped-CSS requirements. Never drift to latest Electron tip or silently sync with it.

## Scope
Copy/adapt the approved Memory, Replicator and Diagnostics tab DOM/layout into `OSJS/src/packages/MCSModbusToolkit`, replace standalone HTML/BrowserWindow bootstrap with existing OS.js window lifecycle, scope all CSS to Toolkit content (including global selectors), supply fixture-only tab data and explicit UNKNOWN/UNAVAILABLE status/LEDs. Disable native Windows service Start/Stop and unsupported path/actions. Keep the app inside the existing `osjs-shell` Docker image; do not add another container.

## Non-scope
No build/tests or verification claims; no real backend calls/Save & Apply, Electron preload/main/runtime imports, legacy package deletion, Modpoll, portable app, new service controls or visual-parity PASS.

## Coding acceptance / handoff
1. Three-tab Toolkit renderer and required local assets authored and committed from pinned SHA with fixture-only rendering and no Electron imports.
2. OS.js desktop/window chrome unaffected by Toolkit-scoped styling; no missing/stale observation displayed as healthy, and unsupported actions cannot execute.
3. Read back changed source, record exact commit, donor provenance, fixture setup and host gaps in `handoff.md`, then archive CODE and activate UMIG-003-T only on a valid source checkpoint. No build, GUI or backend test PASS may be inferred from source review.

## Dependencies
UMIG-001 accepted and archived; sole ACTIVE confirmed. Later JR TEST and VERIFY have distinct acceptance stages. Only BLACK SHEEP WALL may update ICC; stale cache is not an automatic blocker for this source task.

## Sizing
Surface 2, environment 0, behavior 0, verification 0, recovery 1 = 3.
