# PERSIST-021 — Persistence Runtime Status

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-020
Next: PERSIST-022

## Primary outcome
Expose persistence health/status to the OS.js operator surface without changing persistence authority.

## Scope
Add observational status only.

## Non-scope
Do not add historian/telemetry persistence, alternate control authority, or any UI bypass of restore/sealing gates.

## Acceptance
1. Status includes enabled/disabled, snapshot health, last save and last restore outcome when available.
2. Restore failure is clearly visible while sealed state remains authoritative.
3. Status is observational; UI cannot bypass restore/sealing gates.

## Evidence / handoff
Record exact changed paths, source diff, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-020.

## Sizing
2/0/2/1/1=6.

## Coding evidence (PERSIST-021, OpenHands JR DEV)

- Source checkpoint base: `a3d0031feacf34a167fdeaeccab191138bbf3024` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - A `mma2composer/persistence_status.go` (blob `56ceb22b5026fea300c408262e26364554fda5a9`) — new `PersistenceRuntimeStatus` and a pure read-only projection `PersistenceRuntimeStatusFromPlan` from the authoritative restore plan/result. It reports `Configured`, the authoritative `Sealed` state, `Healthy` (only when the restore committed), `SnapshotHealth`, the classified `RestoreOutcome`, and optional `LastSave`/`LastRestore` observations (copied, not aliased). A failed restore stays clearly visible while the sealed state remains authoritative. `PersistenceRuntimeStatusConfigured` covers a memory with no observed restore yet (configured, sealed, health unknown). No control, retry or bypass.
  - M `simulator/apply.go` (blob `c35ae093c9b8580f51f2a0cdcd314bbe07830e25`) — `DeviceRuntimeStatus` gains an optional observational `Persistence` field, populated from the configured persistence enablement only via `persistenceConfigured`.
  - M `OSJS/src/packages/MCSModbusToolkit/diagnostics-model.js` (blob `84970aa4246a8caf8e89996b30e67a4030676c66`) — `mapPersistenceHealth` surfaces the persistence health, failing closed to UNKNOWN for malformed fields and returning null when absent.
  - M `OSJS/src/packages/MCSModbusToolkit/diagnostics-editor.js` (blob `2d46c63609ddd3823db26f0dd958bcf0b952721c`) — renders the persistence health rows when reported; controls remain disabled.
  - A `mma2composer/persistence_status_test.go` (blob `e5796da055c56eb3023537ffebe9aa886b1f6d34`), A `simulator/persistence_runtime_status_test.go` (blob `b93e8b56523c4e04726c2c375887d2ac39aef227`), M `OSJS/tests/toolkit-diagnostics-model.test.js`, M `OSJS/tests/toolkit-diagnostics-editor.test.js` — self-checks.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRuntimeStatus -v .` → 7 tests PASS, exit 0; `cd simulator && go test -mod=readonly -run TestRuntimeStatus .` → ok. JS: `node tests/toolkit-diagnostics-{model,editor,observer}.test.js` → exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0; whole `OSJS` node test suite → all pass, exit 0. New Go files `gofmt`/`go vet` clean; `simulator/apply.go` is intentionally left in the repo's existing compact style (it was not gofmt-clean at baseline).
- Acceptance mapping: (1) status includes enabled/disabled (`Configured`), snapshot health (`SnapshotHealth`), and last save / last restore outcome when available (`LastSave`/`LastRestore`, omitted when absent); (2) restore failure is clearly visible (`RestoreOutcome` classification) while the sealed state remains authoritative (`Sealed` is true whenever the restore did not commit); (3) the status is observational — it is a pure projection with no write path, and the OS.js diagnostics controls remain disabled, so the UI cannot bypass restore or sealing gates.
- Design boundary: observational status only. No historian/telemetry persistence, alternate control authority, or UI bypass of restore/sealing gates is introduced. Independent end-to-end verification is PERSIST-022.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `1310c1367c7f98f084187006084d624d8e20e07b` on GitHub main.

## CWAL
PERSIST-020 is delivered on GitHub main at `1528d3e2451f3588de46f7539127164588878f1a`. PERSIST-021 CODE is now complete and no longer ACTIVE; PERSIST-022 is the next eligible packet (VERIFY — independent JR mode). This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
