# Active Work

Baseline commit: ee19b8a
Working tree: dirty; audited Electron overlay recorded in `ICC/INDEX.md`.
Source dependencies: handoff.md, workflow/active_work/README.md, workflow/active_work/rep-block-002-independent-block-pollers.md, workflow/active_work/rep-block-003-tabbed-block-editor.md, workflow/archive/electron-001-nsis-nssm-service-installer.md, workflow/archive/electron-002-compact-industrial-layout.md
Parent: L0-project
Zoom In: replicator
Zoom Out: L0-project

## Current execution state

There are zero ACTIVE tasks. Under the Active Work invariant, execution stops until a human explicitly activates a task.

Two tasks are QUEUED:

- REP-BLOCK-002 — independent Pull Block pollers; implementation authored, JR retest pending.
- REP-BLOCK-003 — Device/Pull Blocks folder tabs; implementation authored, rendered retest pending.

## Authority and advancement

`workflow/active_work/` agrees with `handoff.md`: zero active and two queued. The queued Replicator pair lacks explicit `Previous:` / `Next:` links, so human workflow repair or explicit activation is required before execution resumes.

## Verification state

- ELECTRON-001 and ELECTRON-002: RETIRED by human decision; initial Electron goal accepted as achieved, with remaining verification gates explicitly not claimed complete.
- REP-BLOCK-002/003: last JR retest stopped at a gofmt failure in `replicator/reader_test.go`; later gates were not run in that attempt.

## Scope note

The audited working tree contains parked Electron source/assets. They must be preserved unless later cleanup or continuation is explicitly authorized.
