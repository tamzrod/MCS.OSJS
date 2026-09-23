# Shared MMA settings through OS.js

## Boundary

Bounded BLACK SHEEP WALL refresh for GROK-UI-002. Parent: INDEX.md.
Connector: osjs-toolkit-advanced-status.md for device/editor status wiring.
No Electron, operator-data or unrelated queue context is refreshed.

## Facts

- Simulator handles mma-load/mma-apply under its mutation mutex and the existing shared writer lock.
- Shared patches permit only rbe, access_events and debug. Null deletes; omitted keys persist.
- The disk-byte SHA-256 revision is compared before mutation; complete candidate validation precedes an atomic config-file replacement.
- Device and ownership files are not changed. Other YAML fields remain in the candidate.
- A matching restart acknowledgment is awaited; committed-but-unacknowledged errors remain explicit and no rollback is attempted.
- Concurrent retry IDs are rechecked/cached inside the mutation boundary.
- The authenticated Toolkit relay allowlists shared operations only on the Simulator route.
- Memory contract preserves committed-error details; shared dialog retains separate drafts and requires reload after conflicts/uncertain results.
- Both advanced editors open the shared dialog and show its last saved RBE port; enabling output prefills editable :9001.
- Updated Simulator and Toolkit must be deployed together. Older backend failures are displayed, not bypassed with direct writes.
- Supervisor acknowledgment is process-start evidence, not a listener-health or free-port guarantee.
- Full Simulator tests, 12 non-socket Toolkit scripts and production package build pass locally on Windows.
- No live Ubuntu runtime, real supervisor restart or operator acceptance has been verified.

## Baseline and overlay

Baseline commit: 0b7ee17. Working tree: dirty.
Audited source dependencies with git blob fingerprints:
- simulator/shared_settings.go: 83c76a95bc143542d8621c607826b8eed2652f58
- simulator/shared_settings_test.go: fa0a362f52eb746efb6777a85764f6a8c6af8b42
- simulator/runtime_server.go: e9eaa3a9197c4a7fd8e7e0e8f74b6122a1609f3f
- OSJS/src/packages/MCSModbusToolkit/shared-settings.js: 5ac07b25e42bb807df43f7b5a572b6fd7734974b
- OSJS/src/packages/MCSModbusToolkit/server.js: 450e88ff2dbb20051332dffd691d8c0b38b7d1b5
- OSJS/src/packages/MCSModbusToolkit/memory-contract.js: 1153f29795dee7844ef6216693bbb65f992201c6
- OSJS/src/packages/MCSModbusToolkit/memory-transport.js: a3e997e627e71cab661de1baa11b6c8c911fec46
- OSJS/tests/toolkit-shared-settings.test.js: 12c9031378d54a358cb5ad73cb40e3b92caac002
