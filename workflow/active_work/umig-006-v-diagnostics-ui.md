# UMIG-006-V — VERIFY: Rendered Diagnostics Safety and Toolkit Reopen

Status: ACTIVE — promoted 2026-09-19 after independent UMIG-006-T UNIT/BUILD PASS review. LIVE VERIFY NOT RUN / NO PASS.
Stage / owner: VERIFY / OpenHands JR; review / ChatGPT
Previous: UMIG-006-T (COMPLETE/PASS, `workflow/archive/umig-006-t-diagnostics-fixtures.md`)
Next: UMIG-007 (QUEUED; no advancement before reviewed live evidence)

## Primary outcome
In a NEW disposable OpenHands Docker project, observe real Toolkit Diagnostics with canonical Memory and Replicator devices: read-only per-device statuses and errors; global service health UNKNOWN; native Windows controls disabled and paths UNAVAILABLE; no hidden writes on refresh. In the same browser session, close the actual Toolkit OS.js window via window chrome, then relaunch from Start and verify the persisted canonical Replicator device, owner, and configuration hash remain unchanged with no automatic apply. This last observation was missing from UMIG-005-V and is NOT yet PASS.

## Exact authority and acceptance
Only `## JR TEST TASK — CURRENT: UMIG-006-V LIVE VERIFY` in `handoff.md` authorizes commands, browser actions, stop conditions, evidence and cleanup. Reuse existing test-only `deploy/verify/compose.yaml` on a freshly verified sandbox-local daemon with a NEW project/volume; never reuse or delete earlier retained verification volumes. Seed through the Toolkit UI ONLY: one synthetic Memory source at private 15020/1 and one Replicator at private 15021/1, FC3 from 127.0.0.1:15020. Confirm healthy Diagnostics with first canonical device explicitly identified. Real window close/relaunch must be directly seen in the browser (tab switch or file read is insufficient); confirm Go-backed reload and no config/owner/hash/apply change. Make source 127.0.0.1:15999 unavailable through a single test-only GUI apply, confirm Replicator ERROR detail while Memory remains independently observed; restore through GUI, delete only both test-owned devices and verify original three config hashes. Scoped project-label-verified `down --remove-orphans` without `-v`; retain volume.

## Boundaries
No production Docker, customer/LAN endpoint, host Modbus port, legacy app or Windows IPC; no service-control invocation, Docker stop/start after initial test startup, backend/Compose/source fix, ICC edit, unscripted exploration or volume removal by JR. Missing safe infrastructure after permitted sandbox-local prep = BLOCKED; contradictory product behavior = FAIL. Do not infer production service health, Windows support, live log retrieval, four-layer COMMS probes, visual parity or migration cutover from this stage.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
