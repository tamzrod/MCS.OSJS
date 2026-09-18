# Handoff

## Human priority — 2026-09-19

Human requested ChatGPT continue bounded coding while OpenHands/JR is busy, with independent TEST/VERIFY deferred. They approved isolated Memory, Replicator and Diagnostics preparations. Existing Docker Compose is operator-reported working and has not been touched. Frozen Electron renderer donor: `1c971b9a6e00bafadf329df8821421a40cfc079c`. Portable Electron and Modpoll are separate.

## Task selection and authority

**NO ACTIVE task — STOP pending explicit next selection.** Human-promoted UMIG-CF-001 Memory source-only contract is archived at `1159d690ddfe85d7ff17fe0e556eb94899cd84c6`, UMIG-CF-002 Replicator contract at `b00bb29dfae4597016820483d00209383a91cf8b`, and UMIG-CF-003 Diagnostics model at `1123ad0d14551d6b628104158b1419363cca43b2` (`workflow/archive/umig-cf-003-diagnostics-model-prep.md`). UMIG-003 three-tab fixture CODE checkpoint `fede1715fadd5900da12fd9630793e3514117caf` is archived. UMIG-003-T build TEST and UMIG-003-V rendered VERIFY are QUEUED/NOT RUN, as are UMIG-004 through UMIG-007A with original dependency/PASS gates intact. DOCKER-001 stays QUEUED/paused; UMIG-008/009 cutover and legacy removal remain Planning. Preserve volumes, persisted configuration, Electron and legacy OS.js apps.

ChatGPT owns CODE, source checkpoints, review of independent evidence and workflow advancement. OpenHands/JR owns only independently instructed TEST/VERIFY under `operation cwal.md`, editing solely the authorized report section. Only BLACK SHEEP WALL edits ICC; authoritative selection is Active Work plus handoff, not stale ICC.

## Code-first checkpoints and deferred behavior

Toolkit source now contains three **dormant, unimported** code-only components: `memory-contract.js` (Simulator v1 load/apply/status), `replicator-contract.js` (Replicator v1 load/apply/status/suggest), and `diagnostics-model.js` (pure per-device observation mapping). Each has an authored, unrun focused test in `OSJS/tests/`. UMIG-CF-003 model consumes future adapter-provided fresh, identity-matched Simulator `result.status` and direct Replicator status results; absent/stale/malformed data shows UNKNOWN, explicit errors UNAVAILABLE. Global MMA2/Simulator/Replicator service status stays UNKNOWN regardless of per-device observations; native mode/paths stay UNAVAILABLE; service Start/Stop disabled. No freshness mechanism, live bridge or Diagnostics UI binding is implemented yet. The two Diagnostics files were read back and source compare `8eb25a3..1123ad0` included exactly those two paths; no tests were executed or claimed PASS. No UI, network, backend, Go, Docker, Windows service or ICC changes from these prep tasks.

## Deferred JR gates

There is NO CURRENT `JR TEST TASK` with no active test. Original UMIG-003-T exact packet and PENDING report are preserved at immutable `bc8fe7330969259a8e39a0fa4f533d88078b79cd`. When OpenHands is available, ChatGPT must activate UMIG-003-T as sole ACTIVE and republish its packet for then-current HEAD; advance to UMIG-003-V only after genuine reviewed TEST PASS. Proposed focused JR tests (NOT RUN): `cd OSJS && node tests/toolkit-memory-contract.test.js`, `node tests/toolkit-replicator-contract.test.js`, and `node tests/toolkit-diagnostics-model.test.js`; each requires an independently authorized current packet. UMIG-004/005/006 real wiring requires original gates and explicit promotion. Historical UMIG-002-T/V PASS establishes only the earlier placeholder.

## Next action and recommendation

Next action (human): select another narrowly bounded code-only preparation task or resume the deferred test queue when OpenHands becomes available. Recommendation: prioritize independent fixture/build and rendered gates before wiring real backend writes or diagnosing UI, and do not cut over launcher, remove legacy apps or modify the working Docker deployment without separate approval.
