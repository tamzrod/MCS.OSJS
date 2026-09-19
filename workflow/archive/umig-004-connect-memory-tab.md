# UMIG-004 — CODE: Connect Toolkit Memory Adapter

Status: COMPLETE — source-only CODE checkpoint, 2026-09-19. No independent build, unit, browser or runtime test has been executed or passed for this stage.
Stage / owner: CODE / ChatGPT
Previous: UMIG-003-V (reviewed and archived independent rendered fixture VERIFY PASS at `6f82196700b1c312652fdd7584a508e5452c4822`)
Next: UMIG-004-T (promoted sole ACTIVE with exact independent JR test packet)

## Scoped source checkpoint
Baseline `f596670527956302031e1ba4aa804386b0cfc133`; final CODE source `508c6b031e675c696de3ff7bdb4291dd005d54ae`. All seven changed paths were inspected/read back; bounded GitHub compare baseline..source confirmed only Toolkit source/tests:
- `OSJS/src/packages/MCSModbusToolkit/server.js`: Toolkit-owned, authenticated allowlisted `load/apply/status` provider, existing shared `$OSJS_DATA_DIR/run/modbus-simulator.sock` with max-1-MiB, 4-byte big-endian JSON frames; forwards complete v1 envelope, maps socket errors to request-correlated RUNTIME_UNAVAILABLE. No new HTTP endpoint, service lifecycle or legacy UI import.
- `OSJS/src/packages/MCSModbusToolkit/metadata.json`: registers its own `server.js` provider; old Simulator metadata and application unchanged.
- `OSJS/src/packages/MCSModbusToolkit/memory-transport.js`: per-process pending request IDs, 26-second timeout and teardown; forwards full envelope for prepared `memory-contract.js` to validate.
- `OSJS/src/packages/MCSModbusToolkit/index.js`: creates transport and contract, replaces ONLY Memory fixture pane with Toolkit-owned editor before attaching window; retains one OS.js window and fixture-only Replicator/Diagnostics, disconnects and disposes on window close.
- `OSJS/src/packages/MCSModbusToolkit/memory-editor.js`: exclusively canonical load/apply document (never fixture-as-live), explicit editable Memory FC1–FC4, Add/Duplicate/Delete/Discard and manual Save & Apply, JSON snapshot and backend-provided apply result, None=0 and positive Random with session-only interval restoration, range/uint32/duplicate validation, save possible after deleting final device, failure keeps local draft and error, status polling selected persisted name only with stale generation protection; absent/wrong/error observations UNKNOWN/UNAVAILABLE, not false healthy. No automatic config writes on mount.
- `OSJS/tests/toolkit-memory-adapter.test.js`: authored-only deterministic contract/transport/status/None-Random/validation checks.
- `OSJS/tests/toolkit-memory-relay.test.js`: authored-only local mock Unix socket/auth/allowlist/framing/unavailable checks, creates and cleans only its own disposable temporary directory.

Existing `memory-contract.js` read and reused unchanged. Reviewed Go `simulator/runtime_server.go` v1 load/apply/status shapes, `simulator/apply.go` IDLE/None behavior, `simulator/validate.go` zero-interval validity, and legacy server's OS.js WebSocket/Unix transport. No code changes to Go, Docker, MMA2, Electron, legacy UI, Replicator/Diagnostics, ICC or user data; no deployment was run.

## Verification boundary
CODE source readback and bounded diff are not TEST PASS. Focused legacy-independent, deterministic tests and local-package build/discovery have NOT RUN. No actual backend connection/configuration write or rendered live editor accepted. UI now contains a user-triggered apply path in source, but operator deployment remains untouched. Independent UMIG-004-T must PASS and be reviewed before separately activating UMIG-004-V safe runtime VERIFY. Preserve operator deployment and volumes, legacy apps and original successor gates.

## Sizing
Surface 2, environment 0, behavior 1, verification 0, recovery 0 = 3.
