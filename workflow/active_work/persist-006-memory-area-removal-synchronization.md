# PERSIST-006 — Memory Area Removal Synchronization

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-005
Next: PERSIST-007

## Primary outcome
Remove the corresponding derived persistence RBE when an authoritative memory area is removed.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement persistence-disable behavior.

## Acceptance
1. Removing one area removes only its system persistence RBE.
2. Other persisted-area RBE projections remain.
3. User RBE rules remain untouched.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-005. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/1/1/1=4.

## Coding evidence (PERSIST-006, OpenHands JR DEV)

- Source checkpoint base: `da7c0fceb037a0af8892d00a2ac63ee90c2de2f4` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - M `mma2composer/persistence_rbe.go` (blob `c2bdacab42995e543aeda22c99050f75fd21dc0d`) — adds `PruneRemovedPersistenceRBE(projection, memory)` (removal counterpart to `SynchronizePersistenceRBE`) and read-only helper `presentPersistenceAreas`. A rule is dropped iff its area is no longer present (nil or count 0); all other projections are returned byte-for-byte unchanged; the input slice is never mutated.
  - A `mma2composer/persistence_rbe_removal_test.go` (blob `49f6de3c0584f147a7710e32dd0c622a0d3ab24d`) — removal self-check.
- Self-check: `cd mma2composer && go test -run TestPruneRemovedPersistenceRBE .` → `ok`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) removing one area removes only its system persistence RBE; (2) other persisted-area projections remain (asserted unchanged); (3) user RBE rules are a separate set and are untouched (only the persistence projection is passed here).
- Non-scope preserved: persistence-disable behavior is not expressed by this function.
- Invariant lives on the shared schema owner (`mma2composer`), applying to Simulator and Replicator alike.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `9ac48cb23e369fb0384b3d8f00811435d5bf7b69`.

## CWAL
PERSIST-005 is delivered on GitHub main at `c743174204455afa83bc92b859fd47582a7afdbc`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
