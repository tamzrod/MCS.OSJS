# UMIG-CF-003 — CODE: Prepare Dormant Diagnostics Observation Model

Status: COMPLETE — SOURCE-ONLY CODE checkpoint, 2026-09-19; no build, unit test, rendered GUI, runtime or deployment test was executed or passed.
Stage / owner: CODE / ChatGPT
Previous: none — independently human-approved preparation, not UMIG-006 advancement.
Next: none — STOP; explicit next task selection required.

## Approval, provenance and checkpoint
Human approved the suggested isolated Diagnostics preparation after completed UMIG-CF-002. Baseline before task: `a0ae4fef94706cc81766d0aeabdfe153d069059e`; task activation `8eb25a3a7e47bfa304e0a39c24231abe2e4cb4a3`; source-only code commit `1123ad0d14551d6b628104158b1419363cca43b2`. Both code files were fetched and read back; bounded compare `8eb25a3..1123ad0` confirmed only these two additions:
- `OSJS/src/packages/MCSModbusToolkit/diagnostics-model.js`: pure read-only, unimported model. It expects future adapter-provided `{fresh:true,expectedName,result}` values, maps validated Simulator `{status:{name,mma2_status,device_status}}` and direct Replicator `{name,enabled,running,source_status,blocks}` device observations, and keeps absent/stale/mismatched data UNKNOWN. Explicit errors yield UNAVAILABLE and preserve their text. No device status is extrapolated into overall service health: global MMA2/Simulator/Replicator remain UNKNOWN. Windows runtime mode/paths stay UNAVAILABLE, Start/Stop disabled.
- `OSJS/tests/toolkit-diagnostics-model.test.js`: authored only for future OpenHands/JR. Exact proposed focused command: `cd OSJS && node tests/toolkit-diagnostics-model.test.js`; NOT RUN. Covers empty/valid/stale/mismatched/malformed/failure inputs and no fabricated global status.

## Limitations and deferred acceptance
The model is deliberately not imported into `OSJS/src/packages/MCSModbusToolkit/index.js` or `toolkit-renderer.js` and not linked to any backend, OS.js server provider, service, timer or native API. The existing fixture Diagnostics tab and disabled controls are unchanged. No Go, Docker, MMA2, Electron, legacy OS.js UI, persisted data or ICC edits. No actual observation freshness mechanism exists yet: a later verified adapter must supply and invalidate `fresh` responsibly; do not pass `fresh:true` for stale caches.

UMIG-003-T build TEST and UMIG-003-V rendered VERIFY remain QUEUED/NOT RUN, and UMIG-004/005/006 full integration follows original predecessors and independent gates. This source-only archive does not imply functional Diagnostics, verified statuses, UI acceptance or permission for service controls.
