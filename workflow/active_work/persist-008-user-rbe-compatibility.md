# PERSIST-008 — User RBE Compatibility

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-007
Next: PERSIST-009

## Primary outcome
Preserve independent user-created RBE rules, including overlap with persistence-owned ranges.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not redefine RBE wire semantics.

## Acceptance
1. User rules may coexist with system persistence rules.
2. Overlapping ranges are accepted when otherwise valid.
3. Persistence lifecycle operations do not mutate user rules.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-007. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5.

## Coding evidence (PERSIST-008, OpenHands JR DEV)

- Source checkpoint base: `01b5e7fe49a8f07ddf0ade4ae2533b1f80a5764b` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - M `mma2composer/persistence_rbe.go` (blob `b6538defa7eaef8cf184ece2518429749fbd10b9`) — adds `RBECoexistence` view + `UserRBERule`, `RBECoexist(memory, userRules)` (pure, copies user rules, conserves overlap), `rangesOverlap`, `PersistenceOverlapsUser` (informational overlap, never rejects), `UserRBEUnchanged` (lifecycle non-mutation check).
  - A `mma2composer/persistence_rbe_coexist_test.go` (blob `1e70626ca51019d9a2ae65062aae6232ba6903f9`) — coexistence/overlap/non-mutation self-check.
- Self-check: `cd mma2composer && go test -run 'TestUserRBE|TestPersistenceOverlap|TestRBECoexist' .` → `ok`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) user rules coexist with system persistence rules in one view; (2) overlapping ranges are accepted and reported, never rejected; (3) persistence lifecycle operations (disable, re-enable, range change) do not mutate user rules — asserted by `UserRBEUnchanged` across all three, and the input slice is copied.
- Non-scope preserved: RBE wire semantics are not redefined; this is a representation/coexistence helper only.
- Invariant lives on the shared schema owner (`mma2composer`), applying to Simulator and Replicator alike.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `4783e0d6e1d3dbd55b29b8647a115905eb5bc285`.

## CWAL
PERSIST-007 is delivered on GitHub main at `feb1e622d38cf8107a1df9ee182e5df17f95790f`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
