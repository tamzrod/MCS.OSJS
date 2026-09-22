# OS.js advanced settings and status

Memory and Replicator expose Device Definition / Advanced Settings folder tabs.
The advanced editor contains RBE Rules, State Sealing and Access Policy, including
comma-separated source IP/CIDR values, resettable presets and custom FC controls.
Memory edits remain in `mma2`; Replicator destination edits remain in
`mma2_advanced`. Existing Save & Apply and Discard own persistence. Replicator
requires the backend's `mma2_advanced` capability before showing editable fields.

`memory-advanced.js` and `comms-status.js` are browser-only adaptations of the
Electron modules at e939a98. They are local so the OSJS-only Docker build context
does not depend on Electron files. Keep future fixes synchronized deliberately.
The package Babel target also transpiles their modern syntax for webpack 4.

COMMS indicators consume runtime `comms` and per-layer observations. A successful
aggregate source poll alone never makes individual LEDs green. Missing telemetry,
errors, unmatched names, unsaved identities and observations older than six seconds
clear the LEDs. Detail tooltips expose actual backend observations. Header lights
describe the selected devices' runtime observations, not global OS service probes.

## Remaining boundary

The current OS.js relay does not expose shared root MMA configuration. The RBE
port is explicitly displayed as Unavailable; the editor does not fabricate a port,
modify shared settings or start outputs. Device RBE rules still require an existing
valid root output; backend validation remains authoritative. Shared MMA settings
popup/API parity is separate work, not implemented by this UI repair.

## Checks

From OSJS, run `node tests/toolkit-ui-parity.test.js` for deterministic DOM tests
with mocked transports. Run the existing non-socket `tests/toolkit-*.test.js`
scripts individually. Build with `NODE_OPTIONS=--openssl-legacy-provider
NODE_ENV=production node node_modules/webpack/bin/webpack.js --config
src/packages/MCSModbusToolkit/webpack.config.js` (PowerShell: set these environment
variables separately). These checks do not establish live Ubuntu runtime behavior.
