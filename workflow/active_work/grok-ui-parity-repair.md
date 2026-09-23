# GROK-UI-001 - Advanced editor and observed status parity

Status: COMPLETE
Stage: CODE
Owner: Codex, explicitly assigned by human

## Scope

Human requests fixing the missing Electron-like advanced settings and nonworking status LEDs on grok, then pushing back to grok. Baseline: e939a98.
Scope is the OSJS MCSModbusToolkit package, its focused tests and supporting scoped documentation. Do not modify Electron, operator data, runtime services or unrelated OTR/OSJT queues.

## Implementation

- Port the browser-only Electron advanced editor and COMMS renderer into the self-contained OS.js build context.
- Connect per-device Memory and Replicator folder tabs to their canonical drafts and existing apply/discard paths.
- Map actual COMMS observations, preserve UNKNOWN for absent/stale telemetry and clear status on identity changes, failure or destruction.
- Update header indicators from selected-device observations without claiming host-service health.
- Shared root MMA settings are not exposed by the current OS.js relay; clearly disclose unavailable output configuration, never invent a port or enable network outputs.

## Verification

- From OSJS: node tests/toolkit-ui-parity.test.js (deterministic DOM with mocked transports; no live runtime).
- Run existing non-socket Toolkit Node regression scripts individually and report each exit code.
- From OSJS: NODE_OPTIONS=--openssl-legacy-provider node node_modules/webpack/bin/webpack.js --config src/packages/MCSModbusToolkit/webpack.config.js (Windows: set NODE_OPTIONS using PowerShell).
- git diff --check.
- No claim of Ubuntu live acceptance from Windows DOM/build checks. Record unavailable gates accurately. Commit and push only scoped changes to grok, as requested.

## Observed results (2026-09-23)

- `node tests/toolkit-ui-parity.test.js`: exit 0; Memory and Replicator DOM/persistence/status scenarios passed with mocked transports.
- Individual `node tests/<name>.test.js` checks: all exit 0 for toolkit-diagnostics-editor, toolkit-diagnostics-model, toolkit-diagnostics-observer, toolkit-fixtures, toolkit-memory-adapter, toolkit-memory-contract, toolkit-replicator-adapter, toolkit-replicator-contract, toolkit-replicator-errors and toolkit-replicator-transport.
- Initial Toolkit webpack build: FAIL on Electron logical assignment syntax. Package-local Babel target corrected; production webpack command above then exited 0 (webpack 4.47.0). Sass emitted a legacy API deprecation warning.
- Windows Node v24.15.0 used; npm reported the existing OS.js Node <17 engine mismatch. No dependency manifests changed. Socket relay tests and live Ubuntu validation were not run.
- `git diff --check`: exit 0. No independent JR/live gate claimed and no OTR/OSJT successor advanced.
- Shared root MMA settings remain outside the current relay contract, explicitly shown unavailable rather than simulated. See package ADVANCED_SETTINGS.md.
