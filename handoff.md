# Handoff

## Human priority — 2026-09-19

Human asked ChatGPT to code while OpenHands is busy and subsequently approved isolated Replicator code preparation. OpenHands retains independent TEST/VERIFY ownership later. Existing Docker Compose is reported working by the operator and was not touched. Frozen UI donor `1c971b9a6e00bafadf329df8821421a40cfc079c`; portable Electron and Modpoll are separate.

## Task selection and authority

**NO ACTIVE task — STOP pending explicit next selection.** Independently human-promoted UMIG-CF-002 Replicator contract preparation reached its source-only gate and is archived at `workflow/archive/umig-cf-002-replicator-contract-prep.md`, code SHA `b00bb29dfae4597016820483d00209383a91cf8b`. Earlier UMIG-CF-001 Memory preparation is archived at `workflow/archive/umig-cf-001-memory-contract-prep.md`, code SHA `1159d690ddfe85d7ff17fe0e556eb94899cd84c6`. The three-tab UMIG-003 fixture CODE source checkpoint `fede1715fadd5900da12fd9630793e3514117caf` is archived; UMIG-003-T build TEST and UMIG-003-V rendered VERIFY are both QUEUED/NOT RUN. UMIG-004 through UMIG-007A remain QUEUED, with original predecessors and PASS gates intact. DOCKER-001 stays QUEUED/paused; UMIG-008/009 cutover/retirement remain in Planning. User Docker data, persisted configurations and legacy apps remain intact.

ChatGPT owns CODE, source checkpoints, review of real JR evidence and workflow advancement. OpenHands/JR only TEST/VERIFY under `operation cwal.md` and an exact current test packet, editing only the permitted report section. Only BLACK SHEEP WALL edits ICC. Handoff and Active Work, not stale ICC, govern task selection.

## Code-first checkpoints and deferred functionality

UMIG-CF-001 introduced dormant `OSJS/src/packages/MCSModbusToolkit/memory-contract.js`, transport-injected Simulator v1 load/apply/status, and an unrun unit test. UMIG-CF-002 introduced dormant `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js`, transport-injected Replicator v1 load/apply/status/suggest with request IDs, error propagation and fail-closed envelope/result checks, plus an unrun unit test. Only the two Replicator files changed in the `2ebe63c..b00bb29` source diff; both were read back. Neither contract is imported into Toolkit UI or wired to OS.js server relay, and all fixture Save & Apply and Start/Stop controls remain disabled. No build, test, rendered GUI, live status, configuration write, or Docker deployment verification was executed for these code-first checkpoints.

## Deferred JR tests

There is NO CURRENT `JR TEST TASK` while no TEST is ACTIVE. Original UMIG-003-T packet and PENDING report remain preserved at immutable commit `bc8fe7330969259a8e39a0fa4f533d88078b79cd`. When OpenHands becomes available, ChatGPT must activate UMIG-003-T as sole ACTIVE, reissue its precise packet for then-current HEAD and review evidence before advancing to rendered UMIG-003-V. Separate focused test commands `cd OSJS && node tests/toolkit-memory-contract.test.js` and `cd OSJS && node tests/toolkit-replicator-contract.test.js` are PROPOSED and have NOT BEEN RUN or authorized as an active JR task. UMIG-004/005 live integration may begin only after original gates and explicit promotion. Historical UMIG-002-T/V passes establish only the earlier placeholder.

## Next action and recommendation

Next action (human): select another independent code-only preparation task, or resume the deferred JR test queue when OpenHands is available. Recommendation: maintain disconnected contracts until independent verification; do not enable writes, cut over launcher, retire legacy apps or alter the working Docker deployment prematurely.
