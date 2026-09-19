# UMIG-004-T — TEST: Toolkit Memory Adapter and Relay Contract

Status: COMPLETE — independent OpenHands/JR UNIT/BUILD TEST PASS reviewed by ChatGPT on 2026-09-19.
Stage / owner: TEST / OpenHands/JR via `operation cwal.md`; evidence review and closure / ChatGPT
Previous: UMIG-004 (archived CODE source checkpoint `508c6b031e675c696de3ff7bdb4291dd005d54ae`)
Next: UMIG-004-V (promoted sole ACTIVE; safe isolated runtime target NOT yet verified)

## Direct evidence and review
JR ran the exact UMIG-004-T packet in a tracked-clean disposable checkout at `54896ff1fbdbfa4c052c993d5b9b3d6e83759ec8`, with Node 16.20.2/npm 8.19.4 and a fake Unix socket owned by a temporary directory. Full report is preserved in immutable `handoff.md` at `0eba36e61ec30254ddccb19bb8ff88ff403131cf`. ChatGPT reviewed the report, governing task and commit diff: `54896ff..0eba36e` changed only `handoff.md`, with only the authorized UMIG-004-T report section replaced. This was an evidence review, not another test run.

JR reported that `node tests/toolkit-memory-contract.test.js`, `node tests/toolkit-memory-adapter.test.js`, `node tests/toolkit-memory-relay.test.js`, `npm run build:local-packages`, `npm run package:discover`, and `npm run build` all exited 0. Cases covered correlated v1 load/apply/status, errors and teardown, None/Random, empty document, auth/allowlist, Unix framing and unavailable handling. Five local packages built; seven discovered including MCSModbusToolkit; `main.js` (30,255 bytes) and `main.css` (121 bytes) emitted. Source and generated metadata register Toolkit `server.js`. No prohibited Electron/legacy/dormant imports. Scoped diff contains seven Toolkit source/test paths plus workflow/handoff; tracked checkout clean. No real backend, Docker, MMA2 or production data touched.

## Boundary
PASS establishes only independent UNIT/BUILD behavior at tested HEAD, not rendered Memory UI, real backend status/apply, Docker health, production write safety or legacy cutover. UMIG-004-V is next but MUST NOT execute live test until an isolated disposable target and exact instructions are verified and documented. No safe target is presently established.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
