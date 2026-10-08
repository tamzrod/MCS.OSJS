# PERSIST-UI01 — Electron Persistence Settings

Status: QUEUED — HUMAN PROMOTED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-R03
Next: PERSIST-R04

## Primary outcome
Expose the existing per-memory persistence configuration in the Windows Electron Memory → Advanced Settings editor without creating another configuration authority.

## Scope
- Add Advanced Settings tab order: `RBE Rules | State Sealing | Persistence | Access Policy`.
- Show RBE mechanism availability + State Sealing as persistence prerequisites; never silently enable sealing or create operator-owned RBE rules.
- Expose `persistence.enabled`, plus read-only persisted areas and locked system-owned persistence RBE projection derived from authoritative memory ranges.

## Non-scope
No filesystem/runtime implementation, no manual snapshot/restore buttons, no editable persistence ranges/RBE IDs/pathnames, no alternate config store, no UI bypass of Save & Apply / Discard, no ICC edits.

## Acceptance
1. Persistence tab participates in the existing Electron draft / Save & Apply / Discard flow and round-trips `persistence.enabled`.
2. Missing RBE mechanism or disabled State Sealing is visibly reported and prevents persistence from being presented as valid; navigation may switch tabs only.
3. Derived persisted areas/system RBE rows are read-only/locked while ordinary user RBE remains editable and unchanged.

## Evidence / handoff
Record changed Electron paths, focused renderer tests, bounded Electron regression and delivered source SHA. Do not claim installed-Windows VERIFY. After genuine delivery, automatically arm PERSIST-R04 for the next invocation and STOP.

## Dependencies
Requires genuine PERSIST-R03 delivery. Runtime behavior remains authoritative; UI is configuration only.

## Sizing
2/0/1/1/0=4.

## CWAL
Already human-promoted. Execute only when selected by handoff. After delivery, automatically activate PERSIST-R04 and STOP.
