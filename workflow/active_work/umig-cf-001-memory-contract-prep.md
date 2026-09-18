# UMIG-CF-001 — CODE: Prepare Dormant Memory Runtime Contract

Status: ACTIVE — human explicitly requested coding continue on 2026-09-19 while OpenHands is occupied; independent TEST/VERIFY postponed, never pre-passed.
Stage / owner: CODE / ChatGPT
Previous: none — independent human-promoted preparation, not an advancement from UMIG-003-T or UMIG-003-V.
Next: none — STOP after source checkpoint; separately choose next preparation or resume the test queue.

## Goal
Prepare a small Toolkit-owned, transport-injected Memory adapter for the existing Simulator version-1 request/response contract. Later UMIG-004 may wire it to a Toolkit-owned OS.js server bridge, but only after the original UMIG-003 TEST and VERIFY gates actually PASS and UMIG-004 is activated.

## Authorized change
- Add only a standalone, *unimported* Toolkit module `OSJS/src/packages/MCSModbusToolkit/memory-contract.js` and focused future JR unit tests `OSJS/tests/toolkit-memory-contract.test.js`.
- Mirror Go Simulator RuntimeRequest/RuntimeResponse semantics (`version`, `request_id`, `operation`, `payload`, `ok`, `result`, `error`). Provide load/apply/status entry points, validate payload before dispatch, correlate responses, propagate actual failure code/message, and reject malformed/missing responses as UNKNOWN rather than fabricate success.
- Transport is an injected function; no direct WebSocket, Unix socket, filesystem, Electron, network, timer, import from legacy UI, or production write. Do not import this module into `index.js` or the fixture renderer, register a server provider, alter Compose, enable Save & Apply, or modify current running behavior.
- Record source checkpoint and deferred tests explicitly. Only ChatGPT authors code; JR later tests.

## Completion gate
Read back the exact two source files and compare the bounded diff. Record a source-only checkpoint without claiming build, unit test, GUI or live behavior. Do not archive as independently verified. Restore the queued JR gate before claiming completion of the actual migration.

## Safety / constraints
UMIG-003-T is QUEUED/NOT RUN; UMIG-003-V remains QUEUED; UMIG-004 remains QUEUED and blocked. This is independent prep only, not a shortcut through their Previous/Next links. Preserve working deployment, persistent volume, legacy apps and ICC. No automatic advancement.
