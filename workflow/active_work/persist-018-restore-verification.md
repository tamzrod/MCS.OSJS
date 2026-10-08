# PERSIST-018 — Restore Verification

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-017
Next: PERSIST-019

## Primary outcome
Add the restore-completion gate that requires all configured persisted areas and their Raw Ingest acknowledgements/integrity checks to succeed before commit.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not unseal directly in this task.

## Acceptance
1. Complete required-area set is tracked deterministically.
2. Any failed/missing area prevents restore completion.
3. Success is emitted only after every required area is validated and committed.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-017. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/2=6.

## Coding evidence (PERSIST-018, OpenHands JR DEV)

- Source checkpoint base: `166fb96b2b990f5a1fee7c8e2c00ab877044635f` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - M `mma2composer/persistence_restore.go` (blob `8a3d716e24ed5747499137e1d3a5d8eb42218513`) — `PersistenceRestoreResult` now records `RequiredAreas` (the deterministic configured area set in plan order, via `persistenceRequiredAreas`) and `AcknowledgedAreas` (the subset written and acknowledged with the Raw Ingest success code). `RestorePersistencePlan` no longer sets `Completed` directly; it emits completion only through the new `VerifyPersistenceRestore` gate, which requires a non-empty required set that exactly equals the acknowledged set (no missing, extra or duplicate areas). Any failed or missing area therefore prevents restore completion. No unseal occurs.
  - M `mma2composer/persistence_restore_test.go` (blob `80ed02387f88958e1f49d46fb78cef8d668fbd86`) — PERSIST-018 self-check.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag|TestVerifyPersistenceRestore' -v .` → 18 tests PASS, exit 0.
- Additional bounded real-socket check (temporary, removed before commit): over a real Raw Ingest v1 server, an all-acknowledged restore completed and verified while a server rejecting one area produced no completion with only the accepted area acknowledged; passed with `-race`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- Acceptance mapping: (1) the complete required-area set is tracked deterministically — `RequiredAreas` is the plan's configured areas in plan order; (2) any failed/missing area prevents completion — a non-OK response or transport error returns before completion, and a partial acknowledgement leaves `AcknowledgedAreas` a strict subset, so the gate fails; (3) success is emitted only after every required area is validated and committed — `Completed` is set solely from `VerifyPersistenceRestore`, which requires the two sets to be identical.
- Design boundary: this task adds the completion gate only; the explicit final unseal is PERSIST-019 and no unseal or commit flag is introduced here.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `3eb118f507688d483c5d4985f74fe6a98b531bc3` on GitHub main.

## CWAL
PERSIST-017 is delivered on GitHub main at `fa1276c26cbb639f061436c3f111351232037b57`. PERSIST-018 CODE is now complete and no longer ACTIVE; PERSIST-019 is the next eligible CODE packet. This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
