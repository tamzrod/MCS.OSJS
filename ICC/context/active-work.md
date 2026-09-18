# Active Work

Baseline commit: ee19b8a
Working tree: clean; the previously audited Electron overlay is now committed, and planning
carries no uncommitted overlay.
Source dependencies: handoff.md, workflow/active_work/README.md, workflow/active_work/*.md,
workflow/archive/electron-001-nsis-nssm-service-installer.md,
workflow/archive/electron-002-compact-industrial-layout.md
Parent: L0-project
Zoom In: replicator, simulator-device-config
Zoom Out: L0-project

## Current execution state

Exactly one task is ACTIVE:

- `UMIG-002-T` — Toolkit build and discovery TEST, owned by OpenHands/JR. Its `JR TEST TASK`
  packet is current in `handoff.md`; the test has NOT been run or observed. UMIG-002 is archived
  in `workflow/archive/umig-002-scaffold-single-osjs-toolkit.md` as a source-only checkpoint,
  which explicitly does not claim a build or UI pass.

QUEUED, not runnable until promoted or activated:

- `UMIG-002-V` — rendered one-window VERIFY for the Toolkit. Human-authorized but not executed.
- `MEM-004` → `MEM-008` — Memory None/Random UI and Windows acceptance chain. MEM-004 is
  paused, not complete.
- `RLED-003` → `RLED-011` — Windows COMMS LED chain, following the archived `RLED-001/002`.
- `RREC-004` — installed-package proof, explicitly PAUSED by the human switch to the OS.js
  Toolkit migration.
- `REP-BLOCK-002`, `REP-BLOCK-003` — Replicator multi-block work; see Zoom In `replicator`.

COMPLETE and still sitting in `workflow/active_work/`:

- `RREC-001`, `RREC-002`, `RREC-003` — ProgramData runtime root, live Replicator transaction,
  truthful service/runtime errors.

## Authority and advancement

`workflow/active_work/README.md` now requires the coding agent to archive a task, follow its
explicit `Next`, verify the successor exists as QUEUED with a matching `Previous`, activate only
that successor, synchronize `handoff.md`, and commit/push that state together. OpenHands/JR never
advances.

`handoff.md` agrees with this directory: `UMIG-002-T` is ACTIVE and `UMIG-002-V` is QUEUED.
ChatGPT owns CODE and advancement; OpenHands is JR for TEST and VERIFY under `operation cwal.md`
and does not code, fix failures, promote/archive tasks, or write ICC.

## Verification state

- UMIG-002-T: no test evidence exists yet. `handoff.md` carries a `PENDING` JR TEST REPORT
  placeholder, not a result.
- ELECTRON-001 and ELECTRON-002: RETIRED by human decision; the initial Electron goal is
  accepted as achieved, with remaining verification gates explicitly not claimed complete.
- MEM-008, RREC-004 and RLED-003..011: no Windows or installer acceptance is established.

## Scope note

No donor SHA is approved, no Electron renderer has been copied into OS.js, no launcher switch
has happened, and no legacy OS.js application package has been deleted. A completed BLACK SHEEP
WALL refresh of this branch, as `handoff.md` required before JR execution, is what produced the
current state above.
