# Handoff

## Human priority — 2026-09-19

The human instructed ChatGPT to do coding while OpenHands is busy and leave independent testing to OpenHands later. This reprioritizes work, not verification truth. The existing Docker Compose deployment is reported working by the operator and is untouched. The frozen one-time Electron renderer donor is `1c971b9a6e00bafadf329df8821421a40cfc079c`. Portable Electron and Modpoll are separate.

## Task selection and authority

**There is currently NO ACTIVE task; STOP until an explicit next task is selected.** The human-authorized independent code-first preparation UMIG-CF-001 reached its SOURCE-ONLY gate and is archived at `workflow/archive/umig-cf-001-memory-contract-prep.md`, code commit `1159d690ddfe85d7ff17fe0e556eb94899cd84c6`. UMIG-003 CODE's original three-tab fixture checkpoint remains archived at `fede1715fadd5900da12fd9630793e3514117caf`. UMIG-003-T TEST and UMIG-003-V rendered VERIFY are both QUEUED/NOT RUN; UMIG-004 through UMIG-007A remain QUEUED and their dependency/PASS gates are not waived. DOCKER-001 remains QUEUED/paused; UMIG-008/009 cutover/retirement stay in Planning pending human approval.

ChatGPT owns CODE, source checkpoints, review of real independent test evidence and workflow advancement. OpenHands/JR only TEST/VERIFY under `operation cwal.md`, without source, workflow or ICC edits; only authorized report section updates. Only BLACK SHEEP WALL edits ICC. The Active Work directory plus this handoff, not stale ICC, selects work. User data, Docker volumes and legacy UIs remain preserved.

## Code-first checkpoint and safe boundary

UMIG-CF-001 added `OSJS/src/packages/MCSModbusToolkit/memory-contract.js` (injected version-1 Simulator Memory request/response adapter, load/apply/status, correlation and fail-closed validation) and `OSJS/tests/toolkit-memory-contract.test.js` (future JR-only unit cases). Both were read back and the bounded `1dab09e..1159d69` diff contains just these two source paths. This code is deliberately dormant: it is NOT imported by the Toolkit entry point or renderer, does not connect to the backend, and cannot enable Save & Apply. No build, test command, GUI observation, live backend acceptance or Docker deploy was run for this code. Actual UMIG-004 integration requires UMIG-003-T and UMIG-003-V independent PASS followed by authorized promotion.

## Deferred JR test packets

There is NO CURRENT `JR TEST TASK` while no test is ACTIVE. Original UMIG-003-T full packet and PENDING report remain preserved at immutable commit `bc8fe7330969259a8e39a0fa4f533d88078b79cd` in `handoff.md`. When OpenHands is available, ChatGPT must explicitly reactivate UMIG-003-T as sole ACTIVE, republish exact updated packet for latest HEAD and then separately run the rendered UMIG-003-V after TEST PASS. The dormant Memory contract has a separate future focused test `cd OSJS && node tests/toolkit-memory-contract.test.js`; this is a proposed command, NOT an observed result or automatically authorized JR task. Historical UMIG-002-T and UMIG-002-V PASS concern only the earlier placeholder.

## Next action and recommendation

Next action (human): select another independently bounded code-only preparation task (for example Replicator contract) while OpenHands remains busy, or request resumption of the deferred JR queue later. Recommendation: continue code-first only behind disconnected interfaces; preserve original verification gates and do not enable real writes, cutover or retire legacy UI before independent acceptance.
