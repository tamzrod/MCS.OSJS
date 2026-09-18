# UMIG-002-V — VERIFY: One Toolkit Window

Status: COMPLETE — independent OpenHands/JR rendered one-window verification PASS.
Stage / owner: VERIFY / OpenHands (JR via Operation CWAL); completion reviewed by ChatGPT
Previous: UMIG-002-T
Next: none — terminal task; no successor is activated without human promotion.

## Primary outcome
Directly verify that the built and discovered MCSModbusToolkit placeholder opens as exactly one OS.js window while the existing desktop and legacy application availability remain intact.

## Executed verification and evidence
OpenHands/JR ran the exact current packet in `handoff.md` in a disposable checkout at `758a585817ce83bec40fc0fae764132dfece3efd`, using sandbox-local Node 16.20.2/npm 8.19.4 and real Chromium via Playwright. Full independent evidence is preserved in the `handoff.md` version at commit `7349caace4762cedcef8109371fb8b02f7a0e8c2`. ChatGPT reviewed the committed report and source task before closure.

1. `npm run build:local-packages`, `npm run package:discover`, and `npm run build` each exited 0; Toolkit built and appeared among seven discovered packages. The test server started on port 18209 and `/healthz` responded successfully.
2. A clean session began with zero windows. Exactly one real Start-menu click launched exactly one `MCSModbusToolkitWindow` showing the literal `NOT CONNECTED — PLACEHOLDER ONLY` text. Browser-console counts showed one Toolkit package launch and one corresponding window constructor; `page_errors` was empty.
3. Desktop, Start menu and taskbar remained functional: clock advanced, window minimize/restore worked using both window control and taskbar entry, and legacy Modbus Simulator and Modbus Replicator remained listed without being launched.

Pre/post tracked checkout status was clean. The test server was stopped, port 18209 was free again and disposable test data were removed. The prior ICC-prerequisite BLOCKED report at `3831e33338a77fcea2e70e96d3f5127193fc5c10` remains historical; the later rendered PASS does not reclassify it. The original build/discovery FAIL at `752a541` is also preserved.

## Non-blocking observations
JR observed four boot-time 404s for Toolkit/legacy app icon paths and a login sound; these were outside the specified acceptance gate and do not establish that the missing resources are fixed. OS.js session snapshots can restore windows across reloads, so the recorded one-window result came from a clean test session.

## Acceptance boundary and continuation
UMIG-002-T build/discovery and UMIG-002-V rendered placeholder launch are established PASS. This does not establish functional Memory, Simulator, Replicator, Diagnostics or Modpoll integration, backend connectivity, default-icon cutover, removal of legacy packages, portable packaging, Windows/Electron acceptance or final UI migration. This task has `Next: none`; stop and await separate human authorization/promotion for the next work item. No ICC write or product source change is part of this closure.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
