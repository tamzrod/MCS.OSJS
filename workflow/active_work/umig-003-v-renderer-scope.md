# UMIG-003-V — VERIFY: Fixture Tabs and Shell Isolation

Status: ACTIVE — promoted after ChatGPT reviewed JR's independent UMIG-003-T PASS at `b002c9f852456b3447ae0bef41ae4f69eb3073e1`; current rendered VERIFY PENDING, NOT RUN, NOT PASS.
Stage / owner: VERIFY / OpenHands (JR via `operation cwal.md`)
Previous: UMIG-003-T (COMPLETE, `workflow/archive/umig-003-t-renderer-build.md`)
Next: UMIG-004 (QUEUED; only if this VERIFY genuinely PASSes and ChatGPT reviews evidence)

## Primary outcome
Directly observe three fixture-only Toolkit tabs in one OS.js window and verify that ShadowRoot styling does not damage the legacy OS.js desktop, Start menu, taskbar or window chrome. This is actual rendered browser verification, NOT static source review, backend integration or final visual parity.

## Current exact test authority
Execute ONLY `## JR TEST TASK — CURRENT: UMIG-003-V` in `handoff.md`: disposable checkout and OSJS_DATA_DIR, compatible Node 16, isolated port 18209, OS.js build/discover/serve, real browser Start-menu click, single window and three real tab clicks, visible fixtures/states, disabled actions, desktop/Start/taskbar/chrome/clock observation, screenshot evidence and safe cleanup. If GUI or required observation is unavailable report BLOCKED. JR may update only authorized `## JR TEST REPORT — UMIG-003-V`, commit/push only `handoff.md`, and stop.

## Required acceptance
1. A clean OS.js session starts with zero windows; one actual Start-menu Toolkit launch yields exactly one `MCSModbusToolkitWindow`. Memory, Replicator and Diagnostics are present as tab buttons; exactly one panel is visibly active at a time after actual clicks, including return to Memory. Capture screenshots or equivalent direct rendered evidence for all tabs, with labels/fixture values intact.
2. Memory shows fixture FC1–FC4 table, example device and explicit fixture notice; Replicator shows example source/destination, pull-block row, UNKNOWN ownership and COMMS (no false green); Diagnostics displays UNKNOWN Simulator service, UNAVAILABLE native runtime mode/paths, a fixture-only log and disabled Start/Stop. Header MMA2/Replicator stays UNKNOWN. All Save & Apply, CRUD, service and COMMS actions remain disabled; fields read-only. No unexplained green/healthy status or real backend interaction.
3. ShadowRoot CSS stays window-local; OS.js desktop, window chrome, clock, Start menu, taskbar and legacy Simulator/Replicator entries survive. Validate actual computed panel visibility, not merely DOM presence. Observe browser page/console errors and unexpected requests. Record target, real browser method, HEAD, screenshots/visible observations, exact actions, commands/exits, pre/post tracked state, and safe cleanup.

## Non-scope / gates
Do not run or manipulate user Docker stack, perform real Modbus requests, install Windows service components, run dormant contracts' unit tests, edit product to fix test results, claim final donor pixel parity or remove legacy apps. UMIG-004 and later integration gates remain QUEUED until independent VERIFY PASS is reviewed and advanced by ChatGPT. Missing safe GUI or indispensable evidence = BLOCKED, not PASS.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
