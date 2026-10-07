# PERSIST-009 — RBE ID Collision Handling

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-008
Next: PERSIST-010

## Primary outcome
Allocate persistence-owned RBE IDs without collision while preserving the existing RBE v1 ID contract.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not redesign RBE v1 or widen its one-byte ID space.

## Acceptance
1. System and user RBE IDs are globally unique as required by the current contract.
2. Users cannot directly assign/change system-owned IDs.
3. No arbitrary permanent user/system ID partition is introduced unless current implementation requires and documents it.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-008. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/1/2/1/2=7; collision policy only.

## Coding evidence (PERSIST-009, OpenHands JR DEV)

- Source checkpoint base: `b0b4823c920d4be73da85f5df3f3bb84e722ee0a` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - M `mma2composer/persistence_rbe.go` (blob `15744d97e4daff2057f904310c6f3767790e0da5`) — `PersistenceRBERule` gains `ID uint8` (serialized `id`); adds `AllocatePersistenceRBEIDs(rules, usedIDs...)`: deterministic, dynamic, collision-free allocation of globally-unique IDs in 1..255, reserving caller-supplied user IDs; fail-closed with no partial set on exhaustion; input never mutated.
  - A `mma2composer/persistence_rbe_id_test.go` (blob `49cba043b9c459c01273fbaaf56fa783cfb95665`) — ID allocation self-check.
- Self-check: `cd mma2composer && go test -run TestAllocatePersistenceRBEIDs .` → `ok`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) system and user RBE IDs are globally unique — allocator reserves user IDs and never repeats; (2) users cannot assign/change system IDs — IDs are assigned by the system, and PERSIST-004's mutation guard rejects any change to a system-owned rule (including its ID); (3) no permanent user/system partition — allocation is dynamic (draws from free IDs; verified that reserving ID 1 shifts persistence to 2,3).
- RBE v1 one-byte contract preserved: IDs bounded to 1..255; exhaustion is an error, never a widened ID.
- Invariant lives on the shared schema owner (`mma2composer`), applying to Simulator and Replicator alike.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `19057f02d94ff404b74a9f01348fc8c31235e374`.

## CWAL
PERSIST-008 is delivered on GitHub main at `4783e0d6e1d3dbd55b29b8647a115905eb5bc285`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
