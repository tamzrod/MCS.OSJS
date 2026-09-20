# Active Work

Baseline commit: c289f2f3872b792c45813ac7a515f8fe90e77671
Working tree: transition overlay contains the OSJT-009 archive, OSJT-010 activation, current JR
packet, and the reviewed active-work queue normalization requested for smaller-model handoff.
Source dependencies: handoff.md, workflow/active_work/README.md, workflow/active_work/*.md,
workflow/archive/electron-001-nsis-nssm-service-installer.md,
workflow/archive/electron-002-compact-industrial-layout.md,
workflow/archive/umig-002-scaffold-single-osjs-toolkit.md,
workflow/archive/umig-002-r-toolkit-discovery-manifest.md
Parent: L0-project
Zoom In: replicator, simulator-device-config
Zoom Out: L0-project

## Current OSJT execution state

OSJT-001 through OSJT-009 are archived. OSJT-010 is the sole ACTIVE TEST task; OSJT-011 onward
remain QUEUED. The current `handoff.md` is the exact source-pinned, chat-only CWAL packet.

OSJT-008 independently PASSed its exact named regression at activation HEAD `644a50c`; the checkout
remained clean and the post-check passed. The coding agent reviewed and accepted that evidence.

OSJT-009 source checkpoint `c289f2f` contains the bounded projection helper and focused regression.
The coding-agent exact prospective check passed; no independent OSJT-010 result exists yet.

`workflow/active_work/OSJT_QUEUE.md` is the compact route for OSJT-010 through OSJT-059. It marks
intentional new artifacts, pairs every CODE deliverable with the following exact TEST gate, and
defines separate CODE/TEST/VERIFY stop and evidence rules. Stale active-work duplicates of archived
OSJT-002 through OSJT-004 are removed. The per-task file and matching handoff remain authoritative.

## Authority and advancement

`workflow/active_work/README.md` now requires the coding agent to archive a task, follow its
explicit `Next`, verify the successor exists as QUEUED with a matching `Previous`, activate only
that successor, synchronize `handoff.md`, and commit/push that state together. OpenHands/JR never
advances.

`handoff.md` agrees with this directory: OSJT-010 is ACTIVE and OSJT-011 is its QUEUED successor.
The coding agent owns CODE and advancement; the independent JR owns TEST/VERIFY execution under
`operation cwal.md` and does not code, fix failures, promote/archive tasks, or write ICC.

## Verification state

- OSJT-010 has no independent result yet. The coding-agent prospective run is preliminary only.
- OSJT-011 through OSJT-059 have no completion evidence and remain non-runnable while QUEUED.

## Scope note

The queue review changes workflow documentation only. It does not execute OSJT-010, authorize a
later task, enable an output, delete a legacy package, or establish any product/runtime PASS.
