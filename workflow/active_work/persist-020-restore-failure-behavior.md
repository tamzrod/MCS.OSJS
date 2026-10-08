# PERSIST-020 — Restore Failure Behavior

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-019
Next: PERSIST-021

## Primary outcome
Keep persistence-enabled memory sealed and expose a deterministic failure state whenever startup persistence restore cannot safely complete.

## Scope
Implement only restore-failure behavior against the current authoritative MCS.OSJS source and existing MMA2 State Sealing/Raw Ingest contracts.

## Non-scope
Do not add repair/retry loops, operator-data mutation, historian behavior, or alternate unseal paths.

## Acceptance
1. Missing, corrupt, or incompatible snapshot keeps the memory sealed.
2. Raw Ingest or restore-verification failure keeps the memory sealed.
3. Failure reason is surfaced without fabricated defaults or automatic unsafe unseal.

## Evidence / handoff
Record exact changed paths, source diff, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-019.

## Sizing
1/0/2/1/2=6.

## Coding evidence (PERSIST-020, OpenHands JR DEV)

- Source checkpoint base: `17d57f055f1a705ccd0f97268c58df7f406b87d0` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - M `mma2composer/persistence_restore.go` (blob `023d89ab1898b4a1414a33929a851459d5f48e35`) — new `PersistenceRestoreFailure` classification with a stable `String()` form and a `persistenceRestoreFailureForArea` mapping, plus a `failedPersistenceRestore` helper. Every non-committing path in `RestorePersistencePlan` now returns through that helper, so the memory is always left sealed and the result carries a deterministic `Failure` and a human-readable `Detail`. Classifications: disabled, unsealed, empty, missing_snapshot, invalid_snapshot, incompatible_snapshot, area_not_ready, unknown_area, raw_ingest_write, raw_ingest_response, incomplete, no_sealing_flag, unseal_write, unseal_response. `PersistenceRestoreResult` gains `Failure` and `Sealed`; a committed restore has `Failure` = `None` and `Sealed` = false. No fabricated defaults, retry loops, alternate unseal paths or operator-data mutation.
  - M `mma2composer/persistence_restore_test.go` (blob `55d9e143b75460a5bd7c28e4944db989750cd4a3`) — PERSIST-020 self-check.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag|TestVerifyPersistenceRestore|TestEncodePersistenceSealingFlag' -v .` → 30 tests PASS, exit 0.
- Additional bounded real-socket check (temporary, removed before commit): over a real Raw Ingest v1 server a rejected area write left the memory sealed and never emitted the unseal packet on the wire; passed with `-race`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- Acceptance mapping: (1) missing/corrupt/incompatible snapshot keeps the memory sealed — the loader yields a non-ready plan and the restore classifies it as `missing_snapshot`/`invalid_snapshot`/`incompatible_snapshot` with `Sealed` true and writes nothing; (2) Raw Ingest or restore-verification failure keeps the memory sealed — `raw_ingest_write`/`raw_ingest_response` and `incomplete` all return before the unseal step, and a rejected unseal is `unseal_write`/`unseal_response`; (3) failure reason surfaced without fabricated defaults or unsafe unseal — every path sets a classified `Failure` plus `Detail`, and the unseal only ever runs on the verified success path.
- Design boundary: this task only adds deterministic failure reporting and the sealed-state guarantee; runtime status surfacing is PERSIST-021. No repair/retry loop, fabricated default, alternate unseal path or operator-data mutation is introduced.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `1528d3e2451f3588de46f7539127164588878f1a` on GitHub main.

## CWAL
PERSIST-019 is delivered on GitHub main at `e22459ecdd10433ef21f63df25e3a22a43fb3662`. PERSIST-020 CODE is now complete and no longer ACTIVE; PERSIST-021 is the next eligible CODE packet. This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
