# PERSIST-R03 — Startup Restore Wiring

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-R02
Next: PERSIST-UI01

## Primary outcome
Wire the existing startup persistence loader/restore/verification/final-unseal contracts into the real runtime while preserving fail-closed State Sealing behavior.

## Scope
- Add real non-test startup call sites for `LoadPersistenceSnapshots` and `RestorePersistencePlan`.
- Persistence-enabled memory starts sealed, loads validated snapshots, restores through existing Raw Ingest v1, verifies the full required set, and performs the existing final explicit unseal only after success.
- Any loader, Raw Ingest, verification or commit failure remains sealed and preserves the existing deterministic failure classification.

## Non-scope
No new seal flag, no alternate unseal path, no automatic retry/repair, no filesystem-format redesign, no UI, no historian behavior, no ICC edits.

## Acceptance
1. Real startup path stays sealed through all restore writes and unseals only after existing verification succeeds.
2. Missing/corrupt/incompatible snapshot or Raw Ingest/verification/commit failure leaves the runtime sealed with the existing classified reason.
3. Focused runtime tests prove success and at least one fail-closed startup case using the real runtime orchestration path.

## Evidence / handoff
Record exact startup call sites, changed paths, focused tests, bounded regression and delivered source SHA. After genuine delivery, automatically arm PERSIST-UI01 for the next invocation and STOP.

## Dependencies
Requires genuine PERSIST-R02 delivery.

## Sizing
2/1/2/1/1=7.

## CWAL
PERSIST-R03 CODE is complete and no longer ACTIVE; PERSIST-UI01 is armed ACTIVE for the next invocation.

## Coding evidence (PERSIST-R03, OpenHands JR DEV)

- Source checkpoint base: `22de66a8f2702b5b6e3d7e7ffb8c5c163b2c188d` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - A `mma2composer/persistence_startup.go` (blob `cb9d739b52d74c0be467465bfbdf3d75bc8441b9`) — `RestorePersistenceAtStartup`, the real startup orchestration: derives the configured load set from the authoritative rules, loads/validates durable snapshots while sealed (`LoadPersistenceSnapshots`), attaches the authoritative State Sealing flag from configuration (`PersistenceSealingFlagFromExtra`), then restores through the existing Raw Ingest v1 contract (`RestorePersistencePlan`). Full required-area verification and the existing final explicit unseal run only after success. No new seal flag, alternate unseal path or retry.
  - A `simulator/persistence_raw_ingest.go` (blob `685a90ba7396daa41cefcb98831ada33a0c98191`) — `persistenceRawIngestWriterAdapter` implements the existing Raw Ingest v1 write against the MMA2 endpoint (magic RI, version 1, area/unit/address/count, packed payload) and returns the response byte.
  - A `simulator/persistence_startup.go` (blob `c97533891196f9bf194c7504a0f24fb2482f7037`) — `persistenceStartupContext` drives the startup restore for every persistence-enabled memory beneath the appliance data root.
  - M `simulator/apply.go` (blob `c27a09e154ed6430d4d31dcf8ff3a8d977829685`) — `newRuntimeApplyRouter` now has the real non-test startup call site and records the observed outcomes.
  - A `mma2composer/persistence_startup_test.go` (blob `2645364bc6860594a47402b39dc2e05666706066`), A `simulator/persistence_startup_test.go` (blob `4e45e2bad0713bb2f4a4118fd5fa14feee22bd4c`) — self-checks.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRestoreAtStartup -v .` → 3 tests PASS, exit 0; `cd simulator && go test -mod=readonly -run TestPersistenceStartup .` → 4 tests PASS, exit 0 (success: 1 area write + final single-coil unseal; missing-snapshot fail-closed with no write; rejected-unseal fail-closed; unsealed precondition refusal).
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt`/`go vet` clean for changed files (`simulator/apply.go` kept in the repo's existing compact style; it was not gofmt-clean at baseline).
- Acceptance mapping: (1) the real startup path stays sealed through all restore writes and unseals only after existing verification succeeds — the sealing bit is forced sealed during area writes and the final unseal is the last action after `VerifyPersistenceRestore`; (2) any missing/corrupt/incompatible snapshot or Raw Ingest/verification/commit failure leaves the runtime sealed with the existing classified reason — `PersistenceRestoreResult.Failure`/`Sealed` are preserved; (3) focused runtime tests prove success and fail-closed startup cases using the real orchestration path and a real Raw Ingest v1 endpoint.
- Design boundary: startup restore wiring only. Electron settings UI is PERSIST-UI01, status wiring PERSIST-R04, the disposable E2E harness PERSIST-R05. No new seal flag, alternate unseal path, retry/repair loop, format redesign, UI or ICC edit is introduced.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `2236e15277ebb17c60cac206e05d77e17e296a53` on GitHub main.
