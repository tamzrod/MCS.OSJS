# Handoff

## Human direction and authority — 2026-09-19

Continue approved staged OS.js MCS Modbus Toolkit migration. ChatGPT owns CODE, JR evidence review and workflow advancement; OpenHands/JR owns independent TEST/VERIFY via `operation cwal.md` and may update ONLY its authorized report, push ONLY `handoff.md` and STOP. Only BLACK SHEEP WALL edits ICC; stale ICC is not task authority. Preserve operator-reported working Docker Compose, production/user configuration and volumes, legacy apps and separate Windows Electron. Never run `docker compose down -v`.

## Current task and accepted evidence

**SOLE ACTIVE: UMIG-004-V — Memory live VERIFY, safe target NOT YET VERIFIED, NOT RUN / NOT PASS.** UMIG-004-T is COMPLETE after independent UNIT/BUILD PASS, archived in `workflow/archive/umig-004-t-memory-adapter.md`; full JR evidence at immutable `0eba36e61ec30254ddccb19bb8ff88ff403131cf` (`handoff.md`). ChatGPT reviewed its evidence, and verified JR commit `54896ff..0eba36e` changed only the authorized handoff JR report section. Three focused tests and three build/discovery commands each exited 0; Toolkit provider metadata/artifacts and import checks matched; fake Unix socket isolated and cleaned; tracked checkout clean. This is unit/build evidence ONLY, not actual GUI, working live backend or safe writes.

Earlier UMIG-003-T build/static and UMIG-003-V fixture/browser passes remain archived (JR evidence `b002c9f852456b3447ae0bef41ae4f69eb3073e1`, `6f82196700b1c312652fdd7584a508e5452c4822`). UMIG-004 CODE checkpoint `508c6b031e675c696de3ff7bdb4291dd005d54ae` archived SOURCE-ONLY. UMIG-005 and later Toolkit tasks QUEUED; DOCKER-001 QUEUED/paused and not independently certified; UMIG-008/009 cutover/removal still Planning and separately approval-gated. No product code or deployment changed in this advancement.

## UMIG-004-V safe-target gate

**Safe target: NOT PROVIDED / NOT VERIFIED.** The operator's Docker deployment is not disposable. `deploy/docker-compose.yml` defines fixed container names, a persistent `osjs-data` named volume and host networking for runtime/MMA2; a second unmodified Compose project on that host does not establish isolation. Require a documented and approved independent test host/sandbox or explicitly isolated test stack with its own data root/volume, runtime/MMA2, collision-free ports, test config/device, ownership and unambiguous cleanup/restore boundaries. No production volumes/sockets or speculative backup/restore. ChatGPT must verify this record and replace the preflight-only packet with exact setup, live browser/test actions, expected results and teardown BEFORE JR performs real testing.

## JR TEST TASK — CURRENT: UMIG-004-V TARGET PREFLIGHT ONLY

GOAL: Non-mutating verification that the safe-test-target prerequisite exists, NOT a live runtime test and NOT a path to PASS.

EXACT ACTIONS (in order): In a fresh, tracked-clean disposable checkout of `tamzrod/MCS.OSJS` at `origin/main`, record `git rev-parse HEAD` and `git status --porcelain`; run `git merge-base --is-ancestor 0eba36e61ec30254ddccb19bb8ff88ff403131cf HEAD` and record exit code. Read `workflow/active_work/umig-004-v-memory-runtime.md`, archived `workflow/archive/umig-004-t-memory-adapter.md`, and this handoff's safe-target section. Confirm UMIG-004-V sole ACTIVE and predecessor COMPLETE/PASS. Check whether a verified disposable target record and executable runtime packet exist here: they are explicitly absent at publication. Therefore if invoked now report **BLOCKED: no verified isolated target**, and STOP before dependency installation, browser launch, Docker, backend startup, any socket access, configuration write or cleanup. Do not guess a target, reuse production or use an obsolete test packet. If checkout or state checks conflict, report BLOCKED with evidence and stop. Do not invoke this deliberately just to reproduce the known blocker; obtain target evidence first.

EXPECTED / EVIDENCE: No live VERIFY result is possible until safe target and precise packet exist. If prematurely invoked return full HEAD, tracked status, ancestor exit, task-state check, missing target proof and BLOCKED. No PASS or failure of product inferred. After target approval, ChatGPT must replace this packet, not let JR improvise.

REPORT-WRITE AUTHORITY: JR may replace ONLY `## JR TEST REPORT — UMIG-004-V` with factual evidence if invoked, commit/push ONLY `handoff.md` then STOP. No code, workflow, ICC or other handoff edits; no successor selection.

## JR TEST REPORT — UMIG-004-V

PENDING — safe isolated target not supplied, live VERIFY not run. No outcome is predetermined from UNIT/BUILD PASS.

## Next action and recommendation

Next action: identify a separate disposable Simulator/MMA2/OS.js test environment or authorize planning one; ChatGPT verifies isolation and publishes exact JR live VERIFY instructions. Recommendation: preserve working Docker/user data and do not invoke live `OPERATION CWAL` until target is documented.
