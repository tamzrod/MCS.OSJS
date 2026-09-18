# UMIG-003-T — TEST: Copied Renderer Build and Fixtures

Status: ACTIVE — coding source checkpoint `fede1715fadd5900da12fd9630793e3514117caf` committed and read back; independent TEST not run.
Stage / owner: TEST / OpenHands (JR via Operation CWAL)
Previous: UMIG-003 (COMPLETE, `workflow/archive/umig-003-copy-electron-renderer.md`)
Next: UMIG-003-V (QUEUED; never advance on CODE inspection or an incomplete JR report)

## Primary outcome
Independently prove the frozen-donor Toolkit fixture renderer builds/discovers under OS.js with a passing focused fixture contract and no Electron runtime import. This is a build/static TEST only, not a rendered GUI, functional backend or Docker acceptance.

## Exact packet
Use only the current `JR TEST TASK — UMIG-003-T` in `handoff.md` for safe setup, exact ordered commands, expected observations, evidence and report-write boundary. Run in a disposable checkout with compatible Node 16; do not modify user Docker deployment, persisted config, backend processes or ICC. If commands/required artifacts are missing, return actual FAIL or BLOCKED; do not invent substitutes.

## Acceptance and evidence
1. OS.js local-package build, package discovery and core build exit 0 with Toolkit present and `dist/main.js`/`dist/main.css` artifacts.
2. `node tests/toolkit-fixtures.test.js` exits 0 for UNKNOWN/UNAVAILABLE, fixture shapes and fresh snapshots; no Electron/legacy runtime dependency in Toolkit source or bundled output.
3. Report HEAD, predecessor ancestry, pre/post tracked status, exact command outputs/exit codes, discovered packages/artifacts, import inspection and any errors in only the authorized report section; no source change or workflow advancement.

## Non-scope
No rendered UI assertion (UMIG-003-V), simulator/replicator live calls, Windows Electron, Docker lifecycle, service controls, product fixes or ICC writes.

## Dependencies
UMIG-001 donor approved at `1c971b9`, UMIG-003 archived source-only CODE at `fede171`, current JR packet. ICC cache freshness is not a test prerequisite.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
