# NPE-05 — Electron Legacy Persistence Cleanup

Status: CODE COMPLETE — delivered on main at `10a4e22aa5350868dac13796ffe78a875fcea641` (OpenHands JR DEV); awaiting independent verification (NPE-06)
Stage: CODE
Owner: OpenHands JR DEV
Previous: NPE-04
Next: NPE-06

## Purpose
Remove the obsolete external persistence implementation from the Electron product after native config wiring is proven.

## Scope
Remove only code proven obsolete because MMA2 now owns persistence runtime, including as applicable:

- Electron snapshot capture/restore helper;
- Electron Raw Ingest persistence restore;
- persistence unlock-coil writes;
- persistence-specific 0x06 watchdog logic;
- persistence-generated/system RBE behavior;
- persistence-specific State Sealing ownership/prerequisites;
- old filesystem snapshot orchestration outside MMA2;
- obsolete tests for those retired behaviors.

Keep:
- ordinary RBE;
- ordinary State Sealing;
- Raw Ingest as an independent feature;
- generic MMA2 restart/apply lifecycle;
- user data/snapshots;
- unrelated diagnostics and settings.

Never delete user snapshot files as part of cleanup.

## Acceptance
Electron focused/full Node tests and relevant Go config tests pass; source search shows no Electron-owned persistence runtime/restore engine remains.

No OS.js Toolkit migration in this packet. STOP after delivery.

## Delivery evidence (OpenHands JR DEV)

Delivered SHA: `10a4e22aa5350868dac13796ffe78a875fcea641` on `main` (base `1427520`).

Changed paths (all authorized; no OS.js Toolkit changes):
- `electron/persistence.js` — deleted (obsolete external persistence runtime: snapshot capture, manifest v1 validation, Raw Ingest restore, State Sealing unlock-coil write, post-unseal Modbus proof, restart/port waiters, snapshot store).
- `electron/test/persistence-lifecycle.test.js` — deleted (focused tests for the retired behaviors).
- `electron/test/package-files.test.js` — added `Electron retains no legacy persistence runtime or restore engine`.

Retained unchanged: ordinary RBE, ordinary State Sealing, Raw Ingest as an independent feature (`main.js` random runtime), the generic MMA2 restart/apply lifecycle, user data/snapshots, unrelated diagnostics/settings. No user snapshot files were deleted.

Commands and results:
- `cd electron && node --test test/*.test.js` → 60 tests PASS, exit 0 (previously 64 with 1 failing legacy test; the failing test was removed with its obsolete subject, and a no-legacy-runtime guard was added).
- `cd mma2composer && go test -mod=readonly ./...` → ok, exit 0.
- `cd simulator && go test -mod=readonly -count=1 ./...` → ok (both packages), exit 0.
- Source search (`grep -rn` over `electron/`) shows no Electron-owned persistence runtime/restore engine remains; only historical workflow-document references to the deleted file remain.

This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.
