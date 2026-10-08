# PERSIST-UI01 — Electron Persistence Settings

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
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
Record changed Electron paths, focused renderer tests, bounded Electron regression and delivered source SHA. Do not claim installed-Windows VERIFY. After genuine delivery, automatically activate PERSIST-R04. If recursive JR DEV continuation is enabled in handoff, perform the context-reset checkpoint and continue with PERSIST-R04 in the same invocation; otherwise stop.

## Dependencies
Requires genuine PERSIST-R03 delivery. Runtime behavior remains authoritative; UI is configuration only.

## Sizing
2/0/1/1/0=4.

## CWAL
PERSIST-UI01 CODE is complete and no longer ACTIVE; PERSIST-R04 is armed ACTIVE for the next invocation.

## Coding evidence (PERSIST-UI01, OpenHands JR DEV)

- Source checkpoint base: `d5730a76f9bad6bb4cc7ab5f9eec9e8ce39c62fa` (`main`, clean, = `origin/main` at run).
- Changed paths (Electron/UI):
  - M `electron/renderer/memory-advanced.js` (blob `de068160c35b38ea71cb9e4ba693762485188a5d`) — the Advanced Settings editor now offers the Persistence tab (after Access Policy) only where the mount declares it supported. It exposes `persistence.enabled` through the same draft, reports the RBE mechanism + State Sealing prerequisites via `persistenceError` (never silently enabling sealing or creating operator-owned RBE rules), and renders the derived persisted areas / locked system-owned persistence RBE projection as read-only rows. Exports `sealingEnabled`/`persistenceEnabled`/`persistenceError`/`derivedPersistenceRules`.
  - M `electron/renderer/app.js` (blob `adb0001e0f52cc00807f2fcf195beb11f61a631f`) — the Memory advanced mount passes `persistenceSupported: true` and the observed `rbeAvailable`.
  - M `electron/test/memory-advanced.test.js` (blob `d55cc9f2f08371d485e0086379381f5657f7fd3b`) and M `electron/test/memory-persistence.test.js` (blob `a64c271ce7bfa73a004869f47a1eb87ddb8d5a23`) — focused renderer tests.
- Targeted self-check: `cd electron && node test/memory-advanced.test.js` → 12 tests PASS, exit 0; `node test/memory-persistence.test.js` → 4 tests PASS, exit 0; every Electron test file passes; `node --check` clean.
- Bounded regression: all Electron test files pass; the OS.js node suite is unaffected (no OS.js change).
- Acceptance mapping: (1) the Persistence tab participates in the existing draft / Save & Apply / Discard flow and round-trips `persistence.enabled` (verified through the same `params` draft and the Electron `inheritSettings`/`applySettings` composition, which already carries unknown keys like `persistence`); (2) missing RBE mechanism or disabled State Sealing is visibly reported (`persistenceError` alert) and prevents persistence from being presented as valid; (3) derived persisted areas/system RBE rows are read-only/locked while ordinary user RBE remains editable and unchanged on the RBE Rules tab.
- Design boundary: configuration UI only. No filesystem/runtime implementation, manual snapshot/restore buttons, editable persistence ranges/RBE IDs/pathnames, alternate config store, or UI bypass of Save & Apply/Discard. This is not an installed-Windows VERIFY.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `cff34f41ebb350fc1c4380253cc2acc02e4562ca` on GitHub main.
