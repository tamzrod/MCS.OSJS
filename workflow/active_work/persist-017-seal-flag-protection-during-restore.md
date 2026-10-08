# PERSIST-017 — Seal-Flag Protection During Restore

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-016
Next: PERSIST-018

## Primary outcome
Prevent persisted sealing-flag value from causing premature unseal during snapshot restoration.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add a second sealing source of truth or runtime reseal semantics.

## Acceptance
1. The configured State Sealing address remains sealed/zero throughout area restoration.
2. A snapshot containing prior flag=1 cannot unseal during restore.
3. Behavior follows the authoritative State Sealing area/address configuration; no duplicate flag setting is added.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-016. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/2=6.

## Coding evidence (PERSIST-017, OpenHands JR DEV)

- Source checkpoint base: `66553c0fb54b0d25ce04e4c6a75242aff825b15a` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - M `mma2composer/persistence_restore.go` (blob `2f14d6bd4132d15e3607afe9311284c30828d964`) — `RestorePersistencePlan` now forces the authoritative State Sealing bit to sealed (0) within the restored coils payload before writing it, so a snapshot captured while unsealed cannot unseal the memory mid-restore. New `PersistenceSealingFlag` and `PersistenceSealingFlagFromExtra` derive the flag location from the configured `state_sealing` block (area `"coil"`, address), mirroring MMA2's enablement rule (absent block disabled; absent `enabled` defaults enabled; explicit flag wins) and introducing no second sealing source of truth. `forcePersistenceSealingFlag` clears exactly that one bit, preserves every other restored bit, returns a copy (never mutates the source), and fails closed if the flag lies inside the range but beyond the payload length; an out-of-range flag is a no-op.
  - M `mma2composer/persistence_snapshot_loader.go` (blob `1d62df77e3be5d9e26eff0c2fcf1d09636fe76ac`) — `PersistenceRestorePlan` gains an optional `SealingFlag *PersistenceSealingFlag` (nil when state sealing is absent/disabled); the loader still writes nothing and does not unseal.
  - M `mma2composer/persistence_restore_test.go` (blob `03ac862771bf67e80370e81ff2185540977972ec`) — PERSIST-017 self-check.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag' -v .` → 12 tests PASS, exit 0.
- Additional bounded real-socket check (temporary, removed before commit): the coils packet on the wire carried the sealing bit cleared while preserving every other snapshot bit; passed with `-race`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- Acceptance mapping: (1) the configured State Sealing address remains sealed/zero throughout area restoration — the coils payload's flag bit is forced to 0 before the write; (2) a snapshot containing prior flag=1 cannot unseal — the forced bit overrides any restored value; (3) behavior follows the authoritative State Sealing area/address and no duplicate flag setting is added — the location comes only from `PersistenceSealingFlagFromExtra` and only that single bit is written.
- Design boundary: restore-completion verification is PERSIST-018 and unseal is PERSIST-019, so this task only protects the flag during transport and never unseals.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `fa1276c26cbb639f061436c3f111351232037b57` on GitHub main.

## CWAL
PERSIST-016 is delivered on GitHub main at `2fbabc3481b6435ca38596ea7dd060e36ed0655d`. PERSIST-017 CODE is now complete and no longer ACTIVE; PERSIST-018 is the next eligible CODE packet. This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
