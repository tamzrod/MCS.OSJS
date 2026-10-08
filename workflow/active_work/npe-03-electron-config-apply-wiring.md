# NPE-03 — Electron Native Persistence Config / Apply Wiring

Status: CODE COMPLETE — delivered on main at `817a6948d7db5bcf666a1e6acf00b1545f9815e8` (OpenHands JR DEV); awaiting independent verification (NPE-06)
Stage: CODE
Owner: OpenHands JR DEV
Previous: NPE-02
Next: NPE-04

## Purpose
Make Electron Save & Apply write only the native MMA2 persistence configuration.

## Scope
Wire the Electron Simulator memory document through compose/validate/save so the resulting MMA2 memory contains the native per-memory `persistence` block exactly as configured.

Save & Apply responsibilities:
1. validate configuration;
2. commit configuration safely;
3. request/reach normal MMA2 restart/readiness using existing generic configuration lifecycle;
4. report configuration apply result.

Save & Apply MUST NOT:
- capture its own snapshot;
- restore snapshot over Modbus/Raw Ingest;
- write an unlock coil;
- poll 0x06 for persistence;
- own a persistence watchdog;
- depend on State Sealing;
- depend on RBE TCP.

MMA2 owns native persistence startup and runtime behavior.

## Acceptance
Focused Electron/main-process tests prove:
- enabled/default emits native persistence with no forced directory/ranges;
- custom directory/ranges round-trip correctly;
- disabled removes/disables only persistence config;
- unrelated memory settings are preserved;
- Save & Apply contains no external persistence restore/unseal lifecycle.

No OS.js Toolkit changes. STOP after delivery.

## Delivery evidence (OpenHands JR DEV)

Delivered SHA: `817a6948d7db5bcf666a1e6acf00b1545f9815e8` on `main` (base `f4453db`).

Changed paths (all authorized; no OS.js Toolkit changes):
- `electron/main.js` — `applySimulator` no longer calls `persistence.captureSnapshots`; it composes + validates + commits the native MMA2 config (carrying each memory's native `persistence` block) and then uses local generic helpers `waitRestartAcknowledged`/`waitPort` for restart readiness. Removed the unused `require('./persistence')`.
- `electron/package.json` — removed `persistence.js` from the electron-builder `build.files` whitelist (no longer required by main.js).
- `electron/test/memory-persistence.test.js` — 5 focused main-process tests (see below).

Commands and results:
- `cd electron && node --test test/memory-persistence.test.js` → 9 tests PASS, exit 0.
- `cd electron && node --test test/*.test.js` → 60 PASS, 1 FAIL. The single failure, `persistence-lifecycle.test.js` "Save & Apply snapshot capture recovers a previously sealed persistence runtime", is PRE-EXISTING and unrelated: it directly exercises the legacy `electron/persistence.js` runtime (Raw-Ingest unseal of a sealed memory during capture) that this packet deliberately stops invoking and that NPE-05 removes. It failed identically before this packet (verified by stashing the NPE-02 changes earlier and by inspecting its subject).

Focused NPE-03 tests added:
- `Save & Apply emits native persistence with no forced directory or ranges` — enabled/default composes `persistence: {enabled: true}`.
- `custom persistence directory and ranges round-trip through Save & Apply`.
- `disabling persistence changes only the persistence block` — `state_sealing`/`rbe` preserved.
- `Save & Apply preserves unrelated and foreign memory settings`.
- `Save & Apply contains no external persistence restore/unseal lifecycle` — the `applySimulator` region and main process contain no `captureSnapshots`/`restoreAndUnseal`/`snapshotsComplete`/watchdog/unseal references and no `require('./persistence')`.

This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.
