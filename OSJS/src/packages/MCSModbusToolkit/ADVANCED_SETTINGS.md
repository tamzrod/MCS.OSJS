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

## Shared MMA settings

Open MMA Settings or RBE TCP Settings inside either Advanced Settings editor.
The shared dialog edits RBE TCP output, debug logging and access-event settings.
Enabling RBE prefills editable `:9001`; existing listener addresses are preserved.
Saving shared settings does not save device drafts. Closing retains the shared
draft, Discard restores the saved snapshot, and Reload explicitly replaces it.
Both device editors display the last saved RBE port, not an unsaved draft value.

The authenticated Toolkit relay routes `mma-load` and `mma-apply` only to the
Simulator runtime. Apply accepts only `rbe`, `debug` and `access_events` patches;
null removes a key and omission preserves it. SHA-256 revision comparison,
full MMA2 candidate validation and the shared writer lock protect the transaction.
Listeners, ownership and device files are not rebuilt by a shared-settings save.
Validation rejects disabling RBE output while existing memory rules require it.

After atomic persistence, the runtime requests restart and waits for the matching
configuration acknowledgment. A failure after commit reports that the configuration
was saved and recovery is required; it never silently rolls back. The UI retains
the draft and requires Reload after conflict, timeout or ambiguous/committed failure.
Supervisor acknowledgment only proves restart initiation, not healthy listeners.
An occupied port can still cause runtime startup failure: check LEDs/diagnostics
after save. Save success is not a bind-availability guarantee.

Deploy the updated Simulator runtime together with the updated Toolkit package.
An older Simulator reports shared settings unavailable; there is no filesystem
fallback from the browser or OS.js server and no Windows service-control API.

## Checks

From OSJS, run `node tests/toolkit-ui-parity.test.js` for deterministic DOM tests
with mocked transports and `node tests/toolkit-shared-settings.test.js` for shared
dialog, contract and mocked framed-relay integration. From simulator, run
`go test -count=1 -timeout=120s ./...`. Run the existing non-socket `tests/toolkit-*.test.js`
scripts individually. Build with `NODE_OPTIONS=--openssl-legacy-provider
NODE_ENV=production node node_modules/webpack/bin/webpack.js --config
src/packages/MCSModbusToolkit/webpack.config.js` (PowerShell: set these environment
variables separately). These checks do not establish live Ubuntu runtime behavior.
