# PERSIST-007 — Persistence Disable Cleanup

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-006
Next: PERSIST-008

## Primary outcome
Remove persistence-owned RBE projections when persistence is disabled for a memory.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not delete or rewrite user-owned RBE entries.

## Acceptance
1. Disabling persistence removes all persistence-owned RBE for that memory.
2. User RBE rules survive unchanged.
3. Re-enabling persistence regenerates projections from current memory ranges.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-006. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5.

## Coding evidence (PERSIST-007, OpenHands JR DEV)

- Source checkpoint base: `8c50ba5e10d44b885c4c40a60cf98acaa9ccaf77` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - M `mma2composer/persistence_rbe.go` (blob `9f2e5b0bdc29b37d020d7da7e1494c05e91577fe`) — adds `PersistenceEnabled(memory)` and `PersistenceRBEProjection(memory)`: the persistence-owned RBE projection computed from enablement + authoritative layout (disabled/nil → empty; enabled → layout-derived system-owned rules, same source as re-enable).
  - A `mma2composer/persistence_rbe_disable_test.go` (blob `96c63890fb6767ee02433be1f8a54c7cb79e8038`) — disable/re-enable self-check.
- Self-check: `cd mma2composer && go test -run 'TestPersistenceDisableCleanup|TestPersistenceProjection' .` → `ok`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) disabling persistence removes all persistence-owned RBE for the memory (projection empty); (2) user RBE rules are a separate set and are never represented/deleted here; (3) re-enabling regenerates projections from the current memory ranges (verified after a range change).
- Non-scope preserved: no user-owned RBE entry is deleted or rewritten; no file/runtime operation.
- Invariant lives on the shared schema owner (`mma2composer`), applying to Simulator and Replicator alike.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `feb1e622d38cf8107a1df9ee182e5df17f95790f`.

## CWAL
PERSIST-006 is delivered on GitHub main at `9ac48cb23e369fb0384b3d8f00811435d5bf7b69`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
