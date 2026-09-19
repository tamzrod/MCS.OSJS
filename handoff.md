# Handoff

## Current human direction — 2026-09-19

Continue approved staged Docker-hosted OS.js MCS Modbus Toolkit migration. ChatGPT codes and reviews independent evidence; OpenHands/JR TEST/VERIFY only via `operation cwal.md`. Preserve operator-reported working Docker Compose, volumes, persisted configurations, legacy OS.js apps and separate Windows Electron. Frozen Electron UI donor `1c971b9a6e00bafadf329df8821421a40cfc079c`; Modpoll and portable Electron remain out of scope. Only BLACK SHEEP WALL edits ICC; stale ICC is not task authority.

## Task selection

**SOLE ACTIVE: UMIG-004 — CODE: Connect Toolkit Memory Adapter.** `workflow/active_work/umig-004-connect-memory-tab.md` governs bounded implementation; `workflow/active_work/umig-004-t-memory-adapter.md` remains QUEUED. UMIG-004-V and subsequent Toolkit tasks remain QUEUED with their original predecessor/PASS gates. DOCKER-001 is QUEUED/paused, not independently certified by these UI tests. UMIG-008/009 launcher cutover/legacy removal remain Planning and require separate approval. No CURRENT JR TEST TASK while a CODE task is ACTIVE; OpenHands must not run an obsolete packet.

## Reviewed independent verification

UMIG-003-T BUILD/STATIC PASS is archived at `workflow/archive/umig-003-t-renderer-build.md`; full JR report at immutable `b002c9f852456b3447ae0bef41ae4f69eb3073e1` (`handoff.md`). UMIG-003-V RENDERED FIXTURE/SHELL VERIFY PASS is archived at `workflow/archive/umig-003-v-renderer-scope.md`; full JR report and test evidence at immutable `6f82196700b1c312652fdd7584a508e5452c4822` (`handoff.md`). ChatGPT reviewed the report and verified JR's commit modified only the authorized handoff report section. JR reports clean disposable Node 16/Chromium 152, successful build/discovery/serve, one Start-menu launch to one Toolkit window, direct Memory/Replicator/Diagnostics/Memory tab switching with one visible panel, correct fixture labels and UNKNOWN/UNAVAILABLE status, disabled unsafe controls, neutral COMMS, no CSS shell leakage, no backend calls, preserved legacy apps, and safe cleanup. Four pre-existing icon/sound 404s were documented as non-gating. These PASS results establish fixture-only display, not live backends, Docker health, configuration writes, dormant contract unit tests, Windows behavior or final pixel parity. Screenshots/logs reported under JR's disposable `/tmp/` locations are not attached to this repository.

## Current CODE boundary and handoff requirements

UMIG-004 implements Toolkit-owned Simulator/Memory load/apply/status transport and renderer mapping using existing OS.js bridge and Docker-shared Unix socket, without importing legacy UI or modifying live deployment. Review dormant `OSJS/src/packages/MCSModbusToolkit/memory-contract.js` and its UNRUN test, old Simulator bridge, runtime v1 schema, and existing fixture UI before coding. Preserve saved definitions, None/Random behavior, explicit errors and UNKNOWN/UNAVAILABLE for missing observations. The renderer's fixture values must not be mistaken for live configurations. Record a bounded source diff/commit and precise message contract. Only after CODE source review should ChatGPT archive UMIG-004, activate UMIG-004-T, publish one exact new JR TEST TASK and leave actual tests to OpenHands. Do not mark tests PASS from code review or run live configuration writes.

## Next action and recommendation

Next action (ChatGPT CODE): inspect existing Simulator bridge/runtime protocol and implement UMIG-004 within its active scope, then source-review and stage the independent UMIG-004-T packet. Recommendation: preserve existing Docker installation and legacy UIs until the separately gated adapter TEST and runtime VERIFY establish safe live behavior; do not advance Replicator/Diagnostics or launcher cutover based solely on rendered fixtures.
