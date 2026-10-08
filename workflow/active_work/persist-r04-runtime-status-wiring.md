# PERSIST-R04 — Runtime Status Wiring

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-UI01
Next: PERSIST-R05

## Primary outcome
Feed real runtime save/restore observations into the existing persistence status projection and operator diagnostics.

## Scope
- Add the real non-test runtime call site for `PersistenceRuntimeStatusFromPlan`.
- Carry actual last-save and last-restore observations from PERSIST-R02/R03 into `DeviceRuntimeStatus.Persistence`.
- Keep status observational: restore failure remains visibly sealed and the UI gains no control authority.

## Non-scope
No persistence writes/restores in the status layer, no new retry/control path, no historian metrics, no UI bypass, no ICC edits.

## Acceptance
1. Runtime status reflects actual configured/snapshot/restore/sealed/healthy state after real save/restore activity rather than configured-only placeholders.
2. Last save/restore observations appear when available and remain absent/unknown when genuinely unobserved.
3. Diagnostics remains read-only and fail-closed for malformed/stale observations.

## Evidence / handoff
Record real status call sites, changed paths, focused Go/JS tests, bounded regression and delivered source SHA. After genuine delivery, automatically activate PERSIST-R05. If recursive JR DEV continuation is enabled in handoff, perform the context-reset checkpoint and continue with PERSIST-R05 in the same invocation; otherwise stop.

## Dependencies
Requires genuine PERSIST-UI01 and PERSIST-R03 delivery.

## Sizing
1/0/1/1/0=3.

## CWAL
PERSIST-R04 CODE is complete and no longer ACTIVE; PERSIST-R05 is armed ACTIVE for the next invocation.

## Coding evidence (PERSIST-R04, OpenHands JR DEV)

- Source checkpoint base: `9d954a9f0906eb5ef2adb4e67d3b1b964098c091` (`main`, clean, = `origin/main` at run).
- Changed paths (product):
  - M `mma2composer/persistence_status.go` (blob `6d0b48dd2ff891bd38ebff8a56036bde8a086c0c`) — `PersistenceRuntimeStatusFromObservations` projects the observational status from configured enablement + optional startup restore result + observed save status: real plan/result + last-save/last-restore when observed, conservative configured-only view otherwise.
  - M `mma2composer/persistence_runtime_save.go` (blob `4401194e0b40057d8dfa9ee328a5b3f71764e25b`) — `PersistenceSaveStatus.LastSaveAt` (recorded only on a real save) and `PersistenceRuntimeSave.Key`.
  - M `simulator/persistence_save.go` (blob `f408910f6df9e5d6ebf8cfb73c3781ab8c81086e`) — `PersistenceRuntimeSaveHost.StatusForKey`.
  - M `simulator/apply.go` (blob `7df13e87ce43cfab53374a8c73b93e20706cd7f1`) — `RuntimeStatus` feeds real observations into `DeviceRuntimeStatus.Persistence` instead of a configured-only placeholder.
  - M `mma2composer/persistence_status_test.go`, M `mma2composer/persistence_runtime_save_test.go`, M `simulator/persistence_startup_test.go` (blob `981384cf78365b37b608f6dca0102e6d651d97d9`) — focused tests, including a shared-listener fake MMA2 serving Modbus reads and Raw Ingest writes on one disposable port.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRuntimeStatus|TestPersistenceRuntimeSave' .` → ok; `cd simulator && go test -mod=readonly -run 'TestRuntimeStatusCarriesRealPersistenceObservations|TestPersistenceStartup' .` → ok.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`; OS.js and Electron node suites → all pass, exit 0. `gofmt`/`vet` clean for changed files.
- Acceptance mapping: (1) runtime status reflects actual configured/snapshot/restore/sealed/healthy state after real save/restore activity — `RuntimeStatus` now projects `PersistenceRuntimeStatusFromObservations` from the recorded startup results and the observed save status; (2) last save/restore observations appear when available and remain absent when genuinely unobserved — `LastSaveAt` is recorded only on a real save and `LastRestore` only from a real restore observation; (3) diagnostics remains read-only and fail-closed — the projection is pure and fabricated nothing.
- Design boundary: status wiring only; no persistence writes/restores in the status layer, no retry/control path, no historian metrics, no UI bypass.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `d9251d3c56bf490c21655bdb372b342ee1eec7f8` on GitHub main.
