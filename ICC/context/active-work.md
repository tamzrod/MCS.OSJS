# Active Work

Baseline commit: 7e3e4fd
Working tree: transition overlay contains only the OSJT-007 archive, OSJT-008 activation, current
`handoff.md` JR packet, and this context update.
Source dependencies: handoff.md, workflow/active_work/README.md, workflow/active_work/*.md,
workflow/archive/electron-001-nsis-nssm-service-installer.md,
workflow/archive/electron-002-compact-industrial-layout.md,
workflow/archive/umig-002-scaffold-single-osjs-toolkit.md,
workflow/archive/umig-002-r-toolkit-discovery-manifest.md
Parent: L0-project
Zoom In: replicator, simulator-device-config
Zoom Out: L0-project

## Current OSJT execution state

OSJT-001 through OSJT-007 are archived. OSJT-008 is the sole ACTIVE task; OSJT-009 onward remain
QUEUED. The current `handoff.md` is the exact chat-only JR TEST TASK packet for OSJT-008, pinned
to source checkpoint `7e3e4fd` with a clean local Ubuntu target and remote-main freshness gate.

The repaired fixture supplies the required root RBE output and directly checks omitted, null,
false, empty, extension-projection, and invalid-save preservation cases. Both focused coding-agent
checks passed; those checks remain preliminary until independent OSJT-008 execution.

OSJT-008 remains incomplete until its exact independent check passes and the coding agent reviews
the returned evidence. OSJT-009 must not be activated early.

## Superseded UMIG state

Exactly one task is ACTIVE:

- `UMIG-002-T` — Toolkit build and discovery TEST, owned by OpenHands/JR. Its `JR TEST TASK`
  packet in `handoff.md` is now the **RETEST** packet; the retest has NOT been run or observed,
  and no PASS exists. The original run FAILed at `752a541` (both commands exited 0 and Toolkit
  assets existed, but the package was absent from discovery output, `packages.json`,
  `dist/metadata.json` and `dist/apps/`). Human-authorized bounded repair `UMIG-002-R` added only
  `OSJS/src/packages/MCSModbusToolkit/package.json` at `cd67e15` and is archived as a source-only
  checkpoint in `workflow/archive/umig-002-r-toolkit-discovery-manifest.md`. UMIG-002 itself is
  archived in `workflow/archive/umig-002-scaffold-single-osjs-toolkit.md` as a source-only
  checkpoint, which explicitly does not claim a build or UI pass.

QUEUED, not runnable until promoted or activated:

- `UMIG-002-V` — rendered one-window VERIFY for the Toolkit. Human-authorized but not executed,
  and not authorized by the current retest packet.
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

- UMIG-002-T: the retest has produced no evidence yet. `handoff.md` carries a `PENDING`
  `JR TEST REPORT — UMIG-002-T RETEST` placeholder, not a result. The earlier FAIL remains
  preserved as history at `752a541` and is not reclassified.
- ELECTRON-001 and ELECTRON-002: RETIRED by human decision; the initial Electron goal is
  accepted as achieved, with remaining verification gates explicitly not claimed complete.
- MEM-008, RREC-004 and RLED-003..011: no Windows or installer acceptance is established.

## Scope note

No donor SHA is approved, no Electron renderer has been copied into OS.js, no launcher switch
has happened, and no legacy OS.js application package has been deleted. The bounded ICC
prerequisite that `handoff.md` required before JR execution is satisfied by the `9775593`
refresh of this branch; the previous refresh described this same task state before the FAIL and
repair commits existed.
