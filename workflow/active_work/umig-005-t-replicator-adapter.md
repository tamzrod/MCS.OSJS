# UMIG-005-T — TEST: Replicator Adapter Contract

Status: ACTIVE — promoted 2026-09-19 after UMIG-005 CODE source-only checkpoint; TEST NOT RUN / NO PASS.
Stage / owner: TEST / OpenHands (JR via `operation cwal.md`)
Previous: UMIG-005 (COMPLETE, `workflow/archive/umig-005-connect-replicator-tab.md`)
Next: UMIG-005-V (QUEUED, prerequisite independent TEST PASS and later safe target)

## Primary outcome
Independently prove Toolkit-owned Replicator contract, adapter validation/status, WebSocket transport, dual-socket relay routing and OS.js build/discovery without touching an actual device or deployed config.

## Exact instruction
The single current `## JR TEST TASK — CURRENT: UMIG-005-T` packet in `handoff.md` is the only execution authority. Run its focused Node tests against mock and mkdtemp Unix server; ensure one Go-shaped apply, typed backend ownership/missing-device errors, inspect/suggest mapping, FC1–FC4 Pull Block validation/gaps, per-block truthful status, unknown COMMS, correlated timeout/unavailable, Memory regression and build/discovery. Capture full command/exit, source HEAD, source/build artifacts, forbidden-import review, tracked clean state and unexpected behavior. Tests and build have NOT been run by CODE; do not infer PASS from authored assertions.

## Non-scope
No real device/MMA2 writes, live Replicator runtime, disposable Compose modification, Docker, browser/live acceptance, product fixes, UI redesign, production config, ICC or workflow edits by JR. Diagnostics stays fixture-only. UMIG-005-V remains QUEUED until ChatGPT accepts actual TEST evidence.

## Dependencies
UMIG-005 source checkpoint `11bea98391b4953ca2356385bf2ffd7164bb657c` plus source-only typed-error test in the task handoff commit. Use current HEAD and packet, do not run older JR instructions.

## Sizing
Surface 0, environment 0, behavior 0, verification 1, recovery 0 = 1.
