# PERSIST-004 — Locked System RBE Behavior

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-003
Next: PERSIST-005

## Primary outcome
Represent persistence RBE rules as system-owned and non-editable/non-deletable through supported configuration/UI paths.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add new persistence runtime behavior.

## Acceptance
1. System persistence RBE is visibly identified as system-owned/locked.
2. User edit/delete attempts cannot mutate the system-derived range or ownership.
3. Ordinary user RBE behavior is unchanged.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-003. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/1=6; ownership/lock behavior only.

## Coding evidence (PERSIST-004, OpenHands JR DEV)

- Source checkpoint base: `9d546675ecef6edaeefb72549bcd7ad9bdf3c362` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - M `mma2composer/persistence_rbe.go` (blob `0590863fc86f60a66bbf788213619333d428f190`) — `PersistenceRBERule` gains `SystemOwned` lock marker (serialized `system_owned`); `DerivePersistenceRBE` emits rules with `SystemOwned: true`. Adds `PersistenceRBERuleKey` (identity = area, not range) and `ValidatePersistenceRuleMutation(requested, derived)`: fail-closed guard rejecting deletion, range/ownership edits and additions of system-owned persistence rules; read-only.
  - M `mma2composer/persistence_rbe_test.go` (blob `c34cc0469a9f2c2065a001cc501e9e52387a2ea7`) — PERSIST-003 expectations updated for the new marker (in-scope corrective retest).
  - A `mma2composer/persistence_rbe_lock_test.go` (blob `ce68f2bc74df1f7e428e235e72890a796c67bb91`) — lock acceptance/rejection + no-mutation self-check.
- Self-check: `cd mma2composer && go test -run 'TestPersistenceRulesAreSystemOwned|TestPersistenceRuleMutation|TestDerivePersistenceRBE' .` → `ok`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator` suites → `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
  - Note: one run of `simulator/cmd/modbus-simulator-runtime` flaked on its time-based schedule-reload assertion (`main_test.go:47`, 2s deadline / 25ms poll). Re-ran that package 3× in isolation and the full simulator module: all `ok`. It exercises runtime scheduling, not `mma2composer`, and my change cannot affect it; reported as a pre-existing timing flake, not a PERSIST-004 regression.
- Acceptance mapping: (1) derived rules carry an explicit `system_owned` marker; (2) a supported configuration path cannot edit/delete them — `ValidatePersistenceRuleMutation` rejects range edits, ownership stripping, deletion and addition; (3) ordinary user RBE rules are a separate set and are never marked system-owned or processed by the guard, so their behavior is unchanged.
- Deferred by design (out of scope): global RBE ID assignment/collision policy (PERSIST-009), user/persistence overlap (PERSIST-008), UI exposure (PERSIST-010), filesystem persistence, restore.
- Invariant lives on the shared schema owner (`mma2composer`), applying to Simulator and Replicator alike.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `936ab24560d57cbf8aacc8180bc00a2527dcfad5`.

## CWAL
PERSIST-003 is delivered on GitHub main at `4e8f774eb946cc9a946938e4f3a28eaa13de5c9d`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
