# UMIG-004-T — TEST: Toolkit Memory Adapter and Relay Contract

Status: ACTIVE — promoted after source-only UMIG-004 CODE checkpoint `508c6b031e675c696de3ff7bdb4291dd005d54ae`; independent TEST PENDING, NOT RUN or PASS.
Stage / owner: TEST / OpenHands/JR (`operation cwal.md`)
Previous: UMIG-004 (COMPLETE, `workflow/archive/umig-004-connect-memory-tab.md`)
Next: UMIG-004-V (QUEUED; only after real independent TEST PASS reviewed by ChatGPT)

## Outcome and authority
On a disposable Linux/Node 16 checkout, independently test the prepared v1 Memory contract, Toolkit per-process transport and Toolkit-owned WebSocket-to-Unix-socket relay against deterministic fake runtime; verify local package builds/discovers and has no Electron/legacy UI import. The sole authoritative ordered packet, expected results, safe setup, evidence and report-write boundary are `## JR TEST TASK — CURRENT: UMIG-004-T` in `handoff.md`.

## Required acceptance
1. All three focused Node tests (Memory contract, Memory adapter and Toolkit relay) exit 0, including correlated load/apply/status, explicit failure and timeout/close handling where authored, None/Random, valid empty document, error/status fail-closed, provider authentication, allowed operations, request framing and isolated Unix-socket unavailable response.
2. OS.js local-package build, discover and main build exit 0; Toolkit metadata registers its own `server.js`; emitted runtime bundle does not import Electron/legacy Simulator/Replicator UIs; existing legacy source and Docker files unmodified.
3. Record exact commands/outputs/exits, HEAD and predecessor state, Node/npm versions, temporary socket/data paths and cleanup, source/metadata/bundle inspection and pre/post tracked status. No actual Simulator/MMA2 operations or production configuration writes.

## Non-scope
No live UI/browser (UMIG-004-V), Docker deploy/restart, real backend apply, Modbus network calls, product fixes, unrelated tests, ICC edits or workflow advancement. A failing test is FAIL; unavailable safe environment is BLOCKED. JR edits only the authorized handoff report, pushes only `handoff.md`, and stops.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
