# Electron Replicator Leds Telemetry

## Boundary

Owns telemetry for replicator memory operations (seals, blocks, migrations). Parent: INDEX.md. No children.

## Source Dependencies

- electron/renderer/app.js
- electron/renderer/style.css
- electron/replicator-runtime.js
- electron/test/replicator-leds.test.js

## Facts

LED indicators show memory state changes (sealed/unsealed blocks, pull/write events) and migration progress. Sealed write confirmation now observed at Modbus transport layer. Telemetry contract expanded: pre-write byte capture, seal status propagation through RBE pipeline. Baseline refresh to `9775593` captures enforcement shift; legacy observation nodes archived but retain their commit refs.

## Verification

node --test electron/test/*.test.js passed 28 tests. These are unit tests only and do not prove deployment acceptance.
