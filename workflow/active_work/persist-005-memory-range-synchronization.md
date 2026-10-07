# PERSIST-005 — Memory Range Synchronization

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-004
Next: PERSIST-006

## Primary outcome
Keep each persistence-owned RBE projection synchronized when its authoritative memory area's start/count changes.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not alter user-owned RBE rules or snapshot files.

## Acceptance
1. Changing area start updates derived persistence RBE start.
2. Changing area count updates derived persistence RBE count.
3. No second user edit is required.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-004. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5; one synchronization path.

## Coding evidence (PERSIST-005, OpenHands JR DEV)

- Source checkpoint base: `cd7d3a82c426ac7f2a9c7b98890b828b79e2d156` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - M `mma2composer/persistence_rbe.go` (blob `4b6a7ee806b389770edac028e6b75a6abc7c603e`) — adds `SynchronizePersistenceRBE(memory)`: the persistence-owned RBE projection is exactly the layout-derived set, so start/count follow the authoritative area and no rule range is ever hand-authored. Roles mirror MMA2 idioms (e.g. `memory_validation.go`): derive vs apply.
  - A `mma2composer/persistence_rbe_sync_test.go` (blob `7c3079e22dc347a193c970abd263215b1bef7c89`) — synchronization self-check.
- Self-check: `cd mma2composer && go test -run TestSynchronizePersistenceRBE .` → `ok`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) changing area start changes derived start; (2) changing area count changes derived count; (3) no second user edit — synchronization is a pure function of the layout and idempotent once aligned. Newly allocated areas appear and removed areas drop out; every rule stays system-owned.
- Non-scope preserved: user-owned RBE rules and snapshot files are untouched (this function operates only on the persistence projection).
- Invariant lives on the shared schema owner (`mma2composer`), applying to Simulator and Replicator alike.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `c743174204455afa83bc92b859fd47582a7afdbc`.

## CWAL
PERSIST-004 is delivered on GitHub main at `936ab24560d57cbf8aacc8180bc00a2527dcfad5`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
