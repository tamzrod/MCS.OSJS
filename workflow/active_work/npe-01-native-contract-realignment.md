# NPE-01 — Native Persistence Contract Realignment

Status: CODE COMPLETE — delivered on main at `b868c2d2d8cb132d52498cb47105b1bd8c095551` (OpenHands JR DEV); awaiting independent verification (NPE-06)
Stage: CODE
Owner: OpenHands JR DEV
Previous: NP-01 historical implementation
Next: NPE-02

## Purpose
Realign the MCS-side persistence schema/model with the CURRENT native MMA2 persistence contract before any UI work continues.

## Authoritative contract
Persistence belongs to each MMA2 memory identity.

- `persistence.enabled: true` enables native MMA2 persistence.
- `directory` is OPTIONAL. Omitted means MMA2 uses its native default beside the loaded YAML.
- omitted `ranges` means ALL allocated areas of that memory.
- optional custom ranges must be contained within their allocated area.
- persistence does NOT require State Sealing.
- persistence does NOT require RBE or RBE TCP.
- persistence does NOT require Raw Ingest restore orchestration.
- MMA2 owns disk snapshot, restore, flush, backup and recovery.

## Scope
Update only the shared MCS configuration/model/validation layer needed by Electron to represent the current MMA2 contract correctly.

Correct the stale NP-01 assumptions, especially:
- remove "enabled requires nonempty directory";
- remove persistence-specific State Sealing prerequisite;
- remove persistence-generated/system RBE assumptions from the native model;
- preserve native per-memory enabled/directory/ranges round-trip;
- preserve unrelated memory/RBE/State Sealing configuration.

Do NOT implement UI in this packet.
Do NOT implement persistence runtime outside MMA2.

## Acceptance
Focused schema/model tests prove:
1. enabled + omitted directory is valid;
2. omitted ranges means no duplicate range projection is created by MCS;
3. explicit valid custom ranges round-trip;
4. invalid out-of-area/overlapping ranges reject;
5. persistence and State Sealing validate independently;
6. persistence and RBE validate independently.

Record exact changed paths, commands, outputs and delivered SHA. STOP after delivery.

## Delivery evidence (OpenHands JR DEV)

Delivered SHA: `b868c2d2d8cb132d52498cb47105b1bd8c095551` on `main` (base `bedf33a`).

Changed paths (all authorized):
- `mma2composer/composer.go` — `Persistence` model comment realigned (directory optional; independent of State Sealing/RBE/RBE TCP).
- `mma2composer/persistence_native.go` — `ValidateMemoryPersistence` no longer requires a nonempty directory; enabled + omitted/empty directory is valid.
- `mma2composer/memory_validation.go` — removed the persistence-specific State Sealing prerequisite and the now-unused `memoryStateSealingEnabled` helper; `ValidateMemory` no longer inspects sealing.
- `mma2composer/persistence_native_test.go` — acceptance 1 (`TestNativePersistenceDirectoryOptional`), 2 (`TestNativePersistenceRangesOmittedValid` now asserts no range projection is materialized), 4 (existing range validation), 6 (`TestPersistenceIndependentOfRBE`).
- `mma2composer/persistence_sealing_validation_test.go` — acceptance 5: `TestPersistenceIndependentOfStateSealing` and `TestPersistenceValidationDoesNotMutateSealing` now assert independence instead of a sealing prerequisite.
- `simulator/persistence_config_schema_test.go` — stale PERSIST-002 prerequisite comment corrected to the independence invariant.

Commands and results:
- `cd mma2composer && go test -mod=readonly -run 'TestNativePersistence|TestPersistenceIndependentOfStateSealing|TestPersistenceValidationDoesNotMutateSealing|TestPersistenceIndependentOfRBE|TestCandidateRejectsRootPersistence|TestCommitRejectsInvalidPersistence' -v .` → 11 tests PASS, exit 0.
- `cd mma2composer && go vet -mod=readonly ./...` → exit 0; `gofmt -l` clean for all changed files.
- Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. (The unrelated `TestWatchDocumentReloadsNoneSchedule` timing test flaked on one run and passed on re-run; it exercises None-schedule reload, not persistence validation.)

This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.
