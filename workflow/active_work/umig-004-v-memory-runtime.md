# UMIG-004-V — VERIFY: Live Memory Behavior

Status: ACTIVE — corrected sandbox-local Docker TARGET PREFLIGHT READY reviewed at JR report `37dc61f533c8e7221610c8bf5858ddbe01790312`; executable live VERIFY packet published, NOT RUN / NOT PASS.
Stage / owner: VERIFY / OpenHands/JR via `operation cwal.md`; evidence review and advancement / ChatGPT
Previous: UMIG-004-T (COMPLETE/PASS, `workflow/archive/umig-004-t-memory-adapter.md`)
Next: UMIG-005 (QUEUED; promote only after real independent live VERIFY PASS reviewed by ChatGPT)

## Single outcome
Directly verify Toolkit Memory canonical load, explicit Save & Apply, persisted None/Random configuration, selected-device status and unavailable/recovery in a genuinely disposable Simulator/MMA2/OS.js environment with real browser evidence.

## Accepted target and packet
ChatGPT reviewed the corrected JR preflight at `37dc61f` and its commit-only handoff diff. An idle daemon was a preparation issue, not a missing engine; newly started daemon inside the existing disposable OpenHands sandbox had zero containers/volumes and no production resources, test-only port free, and resolved Compose config valid. A separate VM is not required. If sandbox is recycled, reconfirm its independently sandbox-local fresh daemon and zero pre-existing resources before test. No access to operator's Docker daemon. Test-only `deploy/verify/compose.yaml` now seeds explicit `devices: []` and provides MMA2 supervisor test-volume write permissions; these additions need fresh read-only Compose config inspection before `up` and have NOT been executed/verified yet.

Current full authority: ONLY `## JR TEST TASK — CURRENT: UMIG-004-V LIVE VERIFY` in `handoff.md`, with exact host/project checks, safe build/start, direct Chromium UI actions, synthetic port 15020, snapshot checks, stop/recover behavior, project-owned teardown, evidence and report-only permissions. On a product discrepancy return FAIL; if isolation or an indispensable observation cannot be established return BLOCKED, never substitute unit/build results. On any failure perform only verified owned-resource cleanup.

## Acceptance
1. Actual isolated OS.js Memory loads the canonical seeded empty document (never fixture-as-live); a synthetic `VERIFY-SIM-1` on private port 15020/unit 1 is created only by explicit Save & Apply, acknowledged by real Simulator/MMA2 and retained after reload.
2. Real GUI shows runtime-derived selected status and persists FC3 None=0 / Random positive / None=0 across explicit edits and reload; missing runtime yields UNAVAILABLE/errors and recovery, not falsely healthy or fixture config.
3. Record direct browser screenshots, actual response/config/status evidence, sandbox daemon/project/volume scope, startup/cleanup and tracked repository status. No operator data or production service touched.

## Non-scope
No production/customer config, host Modbus port, external IP/endpoint, changes to real operator Compose, Replicator/Diagnostics integration, Electron, ICC or legacy UI retirement. The preflight READY and former unit/build PASS are not this task's VERIFY PASS.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
