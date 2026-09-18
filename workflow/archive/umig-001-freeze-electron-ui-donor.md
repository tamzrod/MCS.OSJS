# UMIG-001 — Human Approval: Freeze Electron UI Donor

Status: COMPLETE — human explicitly approved the proposed `1c971b9` donor on 2026-09-19; source provenance inspected and recorded. This is a DONOR DECISION gate, not a CODE/TEST/VERIFY result.
Stage / owner: DONOR DECISION / Human approval; ChatGPT provenance record.
Previous: UMIG-002-V (archived rendered PASS; early scaffold sequence ended there; separate Toolkit sequence explicitly promoted)
Next: UMIG-003 (promoted to ACTIVE only after this record)

## Primary outcome and approval

**Frozen one-time Electron UI donor:** `1c971b9a6e00bafadf329df8821421a40cfc079c` in `tamzrod/MCS.OSJS`.

Approval evidence: ChatGPT proposed this exact short SHA in the previous handoff conversation; the human replied `approved` on 2026-09-19. The full SHA resolves in Git to a commit that repairs Linux Docker IPC; its snapshot also contains the Electron renderer. The approval pins the *renderer snapshot at this commit*, not the Docker repair as a UI change, not a newer tip, and not a claim that Windows behavior has passed testing.

## Source-only donor inventory at pinned SHA

- `electron/renderer/index.html` — three tabs named Memory (`panel-simulator`), Replicator (`panel-replicator`) and Diagnostics (`panel-diagnostics`); header/runtime strip and diagnostics placeholders. Includes `comms-status.js`, then `app.js`, plus `style.css`.
- `electron/renderer/style.css` — gray desktop-tool visual styling, tab layout, device/sidebar/editor forms, FC and pull-block tables, COMMS LEDs/tooltips, diagnostics and responsive layouts. Its `html`, `body`, `*`, bare `button`, `input`, `select`, `textarea`, `.tabs` and other global selectors MUST be scoped to the Toolkit content root before copying so OS.js desktop/window chrome is untouched.
- `electron/renderer/app.js` — DOM composition and form/validation logic for Memory devices, None/Random intervals and FC1–FC4; Replicator devices, destination, pull blocks, UI validation, polling, status and diagnostics. Copy/adapt visual and fixture-only behavior, not live calls or host/process control. Its top-level DOM hooks, timers and document-wide click/keydown handlers require Toolkit-window-scoped lifecycle and cleanup rather than injecting another desktop page.
- `electron/renderer/comms-status.js` — four indicators (Network, TCP, Modbus, MMA2), diagnostic tooltip model, freshness/activity logic and default UNKNOWN. Adapt the browser-global `window.mcsComms` hookup into Toolkit-owned code; fixture/absent observations MUST stay UNKNOWN, not fabricated OK/green.
- Asset reference: `electron/renderer/index.html` points favicon at `../build/replicator-logo.svg`; the actual source is `electron/build/replicator-logo.svg`. `electron/build/icon.png` is a separate Electron window/package icon, and `electron/build/icon.ico`/`installer.nsh` are Windows packaging assets, NOT required renderer imports. If the SVG or PNG is actually reused, copy it into Toolkit-owned OS.js assets and update its path; never reference the Electron tree or installer from Docker. Keep OS.js package metadata/assets independent.
- Historical `.bak` renderer files are NOT the approved donor files; do not copy them.

## Native/host boundary and observed limitations

`electron/preload.js` exposes `window.mcsDesktop` through `contextBridge`/`ipcRenderer`: `getRuntimeStatus`, `getRuntimePaths`, `startAll`, `stopAll`, `simulatorCall`, `replicatorCall`, `onRuntimeStatus`, `onRuntimeActivity`, `onRuntimeLog`. The renderer directly calls these for load/apply/status, diagnostics, subscriptions and Windows/Electron path/service controls. `electron/main.js` supplies BrowserWindow, IPC handlers, `app.getPath`, `process.resourcesPath`/ProgramData paths, filesystem YAML composition/restart requests, Windows services, child processes, local sockets and runtime status. These host files and `electron/replicator-runtime.js`, `replicator-ipc.js`, `runtime-status.js` are boundary references only: never copy/import/execute them inside OS.js or make them a build/runtime dependency.

Docker-hosted OS.js already has its own Simulator and Replicator server relays backed by Linux Unix sockets under `OSJS_DATA_DIR`; their integration belongs to later UMIG-004/005 tasks. UMIG-003 is visual/fixture-only and MUST NOT call the Electron bridge or issue real writes. The donor's Diagnostics Start/Stop buttons invoke `window.mcsDesktop.startAll/stopAll`, and its Windows-service/NSSM, binary and data path displays are not portable. Disable unsupported controls and render UNKNOWN/UNAVAILABLE fixture states until UMIG-006 maps real OS.js diagnostics. The donor `setStatuses` defaults missing runtime reports to STOPPED and can mark RUNNING green; replace absence-based defaults with UNKNOWN in the OS.js copy. Do not infer COMMS health from a missing or stale status. The Electron main-process Simulator status and Replicator bridge are not proof of corresponding Docker status parity; any COMMS mismatch requires independent later tests.

## Copy and independence contract

ONE-TIME copy from the pinned SHA into `OSJS/src/packages/MCSModbusToolkit` only. No symlink, dynamic import, shared source directory, silent synchronization, Electron package/renderer/preload/main dependencies, installer service calls, or additional Toolkit container. Windows Electron and Docker OS.js continue as independent products. Preserve the existing disconnected placeholder and legacy Simulator/Replicator UI packages until separate verified replacement/cutover gates. The proposed Modpoll tab, portable Electron distribution, backend/MMA2 redesign and launcher retirement are outside this decision.

## Acceptance and handoff result

1. Exact full donor SHA and human approval recorded above: decision accepted.
2. Donor HTML/CSS/JS and asset sources plus native boundary documented above: source inventory inspected.
3. Explicit one-time copy and independence established above. No file copy, product edit, build, GUI run, runtime proof or acceptance PASS is claimed by this decision record. UMIG-003 is the sole authorized successor; its separate CODE gate and JR TEST/VERIFY gates remain pending.
