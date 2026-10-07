# PERSIST-012 — RBE-Triggered Snapshot Writer

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-011
Next: PERSIST-013

## Primary outcome
Add a persistence runtime that reacts to persistence-owned RBE events and updates the corresponding snapshot without continuous polling.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement startup restore, manifests, unseal, or UI status.

## Acceptance
1. Persistence writer subscribes to the system persistence RBE signal.
2. On an event it obtains the authoritative configured area state and compares against the snapshot image.
3. Only changed bytes/register words are written; unchanged state causes no disk write.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-011. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/1/2/1/1=7; writer only.

## Coding evidence (PERSIST-012, OpenHands JR DEV)

- Source checkpoint base: `08fe5c575045f5facc0ed616a03135f7c5fca980` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - A `mma2composer/persistence_snapshot_writer.go` (blob `e0a6bd866b1f3e2a7b09e6b49840c5fb5877c0f4`) — `PersistenceSnapshotWriter` implementing the one-byte RBE sink contract (`Publish(id uint8)`), plus `OnPersistenceEvent`; on an event it reads the authoritative area state via `PersistenceAreaReader`, compares against the last snapshot image and writes only changed bytes/register words through `PersistenceSnapshotStore`. Helpers: `PersistenceSnapshotRules` (maps PERSIST-003/009 rules + Port→Unit ID→Memory to writer rules), `changedRuns`, `persistenceAreaKind`.
  - A `mma2composer/persistence_snapshot_writer_test.go` (blob `7c960c3faf181f84d4a8bde1f73b4b6335466f5b`) — writer self-check.
- Self-check: `cd mma2composer && go test -run TestPersistenceSnapshotWriter .` → `ok`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) writer subscribes to the persistence RBE signal via the `rbe.Sink` contract — non-persistence IDs are ignored; (2) on an event it obtains the authoritative configured area state and compares to the snapshot image; (3) only changed bytes/register words are written (verified: 1 changed word → 2 bytes at the right offset; two separated words → two runs/4 bytes), and unchanged state causes **no** store (disk) call.
- Design boundary: concrete file encoding is PERSIST-013 and metadata/manifest is PERSIST-014, so this task defines the store as a byte-offset sink interface; restore/unseal/UI status are out of scope.
- Placement: shared schema owner `mma2composer` (consistent with PERSIST-001..010), using minimal reader/store interfaces so it does not couple to MMA2-internal packages.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `2f18f8196f5497c5435b697b8b1570f59cab2642`.

## CWAL
PERSIST-011 independently PASSed on tested HEAD `313e3be60a29fff6c967d78c0cf677b8b888ed56` against pinned product checkpoint `190e464f6106c3e640d21b1f9352328447464bef`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
