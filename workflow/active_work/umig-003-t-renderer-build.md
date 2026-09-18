# UMIG-003-T — TEST: Toolkit Renderer Build and Fixtures

Status: ACTIVE — explicitly resumed by human on 2026-09-19; PENDING, NOT RUN, NOT PASS.
Stage / owner: TEST / OpenHands (JR via `operation cwal.md`)
Previous: UMIG-003 (COMPLETE, `workflow/archive/umig-003-copy-electron-renderer.md`, source checkpoint `fede1715fadd5900da12fd9630793e3514117caf`)
Next: UMIG-003-V (QUEUED; activate only after ChatGPT reviews independent TEST PASS)

## Primary outcome
Independently build and discover the fixture-only Toolkit renderer under OS.js, run its focused fixture contract, and verify its source/bundle has no Electron or legacy UI runtime dependency. This is a static/build gate ONLY. It does not verify actual rendered behavior, Docker health, or live Memory/Replicator/Diagnostics integration.

## Exact test authority
Use `## JR TEST TASK — CURRENT: UMIG-003-T` in current `handoff.md` for exact ordered commands, expected results, safe disposable setup, evidence, and report-only write permission. Use compatible Node 16; never touch the operator's working Compose stack or persisted data. If the checkout is dirty or prerequisite/task state conflicts, report BLOCKED rather than cleaning or repairing it.

## Current source boundary
The approved donor is pinned at `1c971b9a6e00bafadf329df8821421a40cfc079c`. Renderer implementation was archived source-only at `fede1715fadd5900da12fd9630793e3514117caf`. Since that checkpoint, separately approved UMIG-CF-001/002/003 added unimported `memory-contract.js`, `replicator-contract.js` and `diagnostics-model.js` plus unrun tests. Those three modules are dormant and must not be confused with a wired renderer or silently tested as part of this task. Their mere presence is allowed in a bounded baseline diff, but Electron/legacy imports in any Toolkit runtime source/bundle are not.

## Acceptance and evidence
1. `npm run build:local-packages`, `npm run package:discover`, `npm run build` and `node tests/toolkit-fixtures.test.js` succeed independently; `MCSModbusToolkit` is discovered and local `dist/main.js`/`main.css` exist.
2. Fixture checks cover UNKNOWN/UNAVAILABLE, representative Memory/Replicator shapes and fresh independent snapshots. Inspect Toolkit source/bundle: no Electron or legacy UI runtime import; dormant contract/model modules stay unimported by Toolkit entry points.
3. Report exact HEAD and predecessor ancestry; Node/npm versions; pre/post tracked status; commands, outputs and exit codes; discovery, artifact paths/sizes, explicit grep exit, bundle/import observations and unexpected findings. JR edits only the permitted handoff test report, commits/pushes only `handoff.md`, then STOPS.

## Non-scope
No rendered UI (UMIG-003-V), tests of new dormant modules (independent future tasks), backend requests, privileged endpoints, product fixes, server restarts, operator Docker deployment, ICC, source/workflow changes, or automatic advancement. Original verification dependencies remain intact.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
