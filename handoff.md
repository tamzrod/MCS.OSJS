# Handoff

## Human priority — 2026-09-19

Human asked ChatGPT to continue coding while OpenHands is busy and explicitly approved the proposed Replicator code-first preparation. Independent testing remains OpenHands/JR's later responsibility. Existing Docker Compose deployment is reported working by operator and remains untouched. Frozen one-time Electron UI donor is `1c971b9a6e00bafadf329df8821421a40cfc079c`. Portable Electron and Modpoll are separate.

## Task selection and authority

**Sole ACTIVE: UMIG-CF-002 — dormant Replicator contract CODE preparation only.** This human-promoted standalone source task is not an advancement of UMIG-003-T/V or UMIG-005. UMIG-CF-001 Memory source checkpoint is archived at `workflow/archive/umig-cf-001-memory-contract-prep.md` (code SHA `1159d690ddfe85d7ff17fe0e556eb94899cd84c6`). UMIG-003 fixture CODE at `fede1715fadd5900da12fd9630793e3514117caf` is archived. UMIG-003-T TEST, UMIG-003-V rendered VERIFY, UMIG-004 through UMIG-007A remain QUEUED/NOT PASSED with their original gates; DOCKER-001 QUEUED/paused; UMIG-008/009 cutover/retirement in Planning. Preserve Docker volumes, user data and legacy apps.

ChatGPT owns CODE, source-only checkpoints and later report review/workflow advancement. OpenHands/JR only TEST/VERIFY using `operation cwal.md` under an exact current packet, with no product/workflow/ICC writes beyond the authorized handoff report. Only BLACK SHEEP WALL edits ICC; stale cache does not prove execution status.

## Current authorized CODE boundary

Prepare only standalone `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js` and `OSJS/tests/toolkit-replicator-contract.test.js`, deriving v1 wire semantics from `replicator/runtime_api.go` and the existing Replicator OS.js relay. Inject transport; no runtime transport, UI import, Save & Apply enablement, live messages, Docker/Go/legacy/Electron changes or service controls. Read back changed source and inspect bounded compare; never claim source inspection as build/test PASS. After source-only closure, no ACTIVE task until explicit selection. The earlier dormant Memory contract is likewise unconnected and untested.

## Deferred JR verification

There is NO CURRENT `JR TEST TASK` during this CODE task. Original UMIG-003-T packet and pending report are preserved at historical `bc8fe7330969259a8e39a0fa4f533d88078b79cd`; it must be republished for latest HEAD when activated, then UMIG-003-V rendered VERIFY only after reviewed TEST PASS. Future focused tests `cd OSJS && node tests/toolkit-memory-contract.test.js` and `cd OSJS && node tests/toolkit-replicator-contract.test.js` are proposed only; neither is claimed executed. Historical UMIG-002 tests establish the earlier placeholder only.

## Next action and recommendation

Next action (ChatGPT): complete only UMIG-CF-002 source preparation, read back and archive; stop at the source-only gate. Recommendation: retain original test dependencies; do not wire live writes, remove legacy UIs or rebuild operator Docker deployment without subsequent authorization and independent acceptance.
