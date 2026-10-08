# PERSIST-R02 — Runtime Save Wiring

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-R01
Next: PERSIST-R03

## Primary outcome
Wire persistence-owned RBE events in the real runtime to the existing `PersistenceSnapshotWriter` and the PERSIST-R01 filesystem adapter.

## Scope
- Create the non-test runtime call site for `NewPersistenceSnapshotWriter`.
- Subscribe only the system-derived persistence RBE projection and route its one-byte rule events to snapshot writes.
- Preserve change-only behavior: read the authoritative configured area on event and persist only changed bytes/words through existing writer semantics.

## Non-scope
No startup restore, no polling loop, no user-RBE mutation, no alternate ranges, no UI, no historian behavior, no ICC edits.

## Acceptance
1. A real runtime persistence RBE event reaches `PersistenceSnapshotWriter` for the correct memory/area identity.
2. User-owned RBE remains independent and persistence-owned ranges remain derived/system-owned.
3. Focused runtime tests prove changed state saves and unchanged state causes no unnecessary snapshot rewrite.

## Evidence / handoff
Record exact runtime call sites, changed paths, focused tests, bounded regression and delivered source SHA. After genuine delivery, automatically arm PERSIST-R03 for the next invocation and STOP.

## Dependencies
Requires genuine PERSIST-R01 delivery.

## Sizing
2/1/1/1/0=5.

## CWAL
PERSIST-R02 CODE is complete and no longer ACTIVE. Evidence below; PERSIST-R03 is armed ACTIVE for the next invocation.

## Coding evidence (PERSIST-R02, OpenHands JR DEV)

- Source checkpoint base: `6fecbf6a046d1d26ee5dd33d6e0e361d2c473584` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - A `mma2composer/persistence_runtime_save.go` (blob `3dc796dab80259a75b5f75ebbea327a915c89375`) — `PersistenceRuntimeSave` is the concrete runtime save component: it subscribes only the system-derived persistence RBE projection (rejecting any user-owned rule as a subscription), derives change-only writer rules from `PersistenceSnapshotRules`, maps each persistence rule ID to its memory/area, and routes each one-byte event to the existing `PersistenceSnapshotWriter`. `PersistenceSaveStatus` observes events/saves/changed bytes/last error; it grants no control authority.
  - M `mma2composer/persistence_filesystem.go` (blob `229cc504f73306bc4c29bef227b8b171b60b208d`) — `PersistenceSnapshotConfigs` derives the adapter registration from the authoritative rules.
  - A `simulator/persistence_save.go` (blob `a4d8112419dffa7b71544a560093cf7315c7d1c1`) — `PersistenceRuntimeSaveHost` builds and routes the save components beneath the appliance data root; `PersistenceAreaReaderFunc` adapts a function reader.
  - A `simulator/persistence_area_reader.go` (blob `a06e45bcaff246ac1132151d133d9b17462180f6`) — the authoritative current-state reader performs a real Modbus TCP read (FC1/FC3) and returns the snapshot encoding; wrong-length/exception responses fail closed and never fabricate bytes.
  - A `simulator/persistence_rbe_subscriber.go` (blob `4bc66ef653ef73f01f3e5fc9d96d010b82d6b9e0`) — `PersistenceRBESubscriber` forwards non-zero one-byte rule IDs from the MMA2 RBE v1 event stream.
  - M `simulator/apply.go` (blob `4f93dd73547cbbbab65d3ceff5e7d0beacc19f6e`) — the live `SchedulerApplier` gains the real non-test call site `ArmPersistenceSave` and the event router `PublishPersistenceEvent`.
  - M `simulator/cmd/modbus-simulator-runtime/main.go` (blob `dc8293978abf48cf4c7ebbe7e857aa8df7366647`) — the reload watcher keeps the save wiring armed.
  - A `mma2composer/persistence_runtime_save_test.go` (blob `867a5ac9c90f9ed4ebf82b769df098af1ca7d57c`), A `simulator/persistence_runtime_save_test.go` (blob `c5ca18bd3ae1c5d7db3139690774e3fcb52af643`), A `simulator/persistence_area_reader_test.go` (blob `04ac5c8b950b0967020cda6db9063f89037ce5fa`) — self-checks.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRuntimeSave -v .` → 4 tests PASS, exit 0; `cd simulator && go test -mod=readonly -run 'TestPersistenceRuntimeSave|TestPersistenceRBESubscriber|TestModbusReadArea|TestPersistenceAreaReader' -v .` → 9 tests PASS, exit 0 (includes an end-to-end device→Modbus→writer→filesystem-adapter save of changed state, and an unchanged-state no-rewrite case).
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0; whole OSJS node suite → all pass, exit 0. `gofmt`/`go vet` clean for changed files (`simulator/apply.go` is kept in the repo's existing compact style; it was not gofmt-clean at baseline).
- Acceptance mapping: (1) a real runtime persistence RBE event reaches `PersistenceSnapshotWriter` for the correct memory/area identity — `PersistenceRuntimeSave`/`PersistenceRuntimeSaveHost` map persistence rule IDs to the derived rule and route via the real `SchedulerApplier.PublishPersistenceEvent` call site; (2) user-owned RBE remains independent and persistence rules stay derived/system-owned — `NewPersistenceRuntimeSave` rejects a non-system rule and only the derived projection IDs are subscribed; (3) focused runtime tests prove changed state saves and unchanged state causes no rewrite.
- Design boundary: runtime save wiring only. Startup restore wiring is PERSIST-R03, UI PERSIST-UI01, status wiring PERSIST-R04, the disposable E2E harness PERSIST-R05. No startup restore, polling loop, user-RBE mutation, alternate range, UI or ICC edit is introduced.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `407fcc32979b9acd83dd6826e86f4b700da70c1b` on GitHub main.
