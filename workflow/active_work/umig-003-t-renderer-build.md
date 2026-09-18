# UMIG-003-T — TEST: Copied Renderer Build and Fixtures

Status: QUEUED — human deferred JR testing on 2026-09-19 while OpenHands works on other projects. NOT RUN, NOT PASS. Resume explicitly after the separate code-first preparation task; do not infer verification from later code.
Stage / owner: TEST / OpenHands (JR via Operation CWAL)
Previous: UMIG-003 (COMPLETE, `workflow/archive/umig-003-copy-electron-renderer.md`)
Next: UMIG-003-V (QUEUED; only after independent PASS)

## Primary outcome
Independently prove the frozen-donor Toolkit fixture renderer builds/discovers under OS.js with a passing focused fixture contract and no Electron runtime import. Build/static TEST only, not rendered GUI, functional backend or Docker acceptance.

## Packet preservation and resumption
The complete original JR packet and PENDING report remain in immutable commit `bc8fe7330969259a8e39a0fa4f533d88078b79cd` (`handoff.md`). No packet is CURRENT while this task is QUEUED. On explicit resumption, ChatGPT must ensure this task is the sole ACTIVE, regenerate its packet against then-current HEAD, retain the original code checkpoint `fede1715fadd5900da12fd9630793e3514117caf` and determine whether later dormant files affect checks. Do not give JR an outdated active-task check or infer PASS from code review.

## Acceptance and evidence
1. On a disposable Node 16 checkout, `cd OSJS && npm run build:local-packages`, `npm run package:discover`, and `npm run build` exit 0; Toolkit discovered; its JS/CSS artifacts exist.
2. `cd OSJS && node tests/toolkit-fixtures.test.js` exits 0 for UNKNOWN/UNAVAILABLE, fixture shapes and fresh snapshots, with no Electron/legacy dependency in Toolkit source or bundle.
3. Report HEAD, ancestry, clean tracked pre/post status, commands/output/exits, discovery, artifacts, dependency inspection, and errors; edit only JR's authorized handoff report. No code or workflow edits by JR.

## Non-scope and dependency
No rendered UI (UMIG-003-V), live backend, Docker lifecycle, product fixes or ICC writes. Previous CODE checkpoint and exact newly published JR packet required. Stale ICC is not a prerequisite. User's coding-first priority does not waive this gate.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
