# PERSIST-016 — Raw Ingest Restore

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-015
Next: PERSIST-017

## Primary outcome
Restore validated snapshot area payloads into the matching MMA2 memory through the existing Raw Ingest v1 contract.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not unseal memory or change Raw Ingest protocol.

## Acceptance
1. Each snapshot area restores to the same area/start/count identity.
2. Every Raw Ingest response is checked; non-zero response aborts restore.
3. No cross-area mirroring is added by persistence.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-015. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/1/2/1/2=8; restore transport only.

## Coding evidence (PERSIST-016, OpenHands JR DEV)

- Source checkpoint base: `a497d9d4ca35cab8de832bc92fc3b5d90c7a4b33` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed/added paths (product):
  - A `mma2composer/persistence_restore.go` (blob `3c31ade58d4a0f7c5a82f40c6873244d15002586`) — `RestorePersistencePlan` consumes the PERSIST-015 `PersistenceRestorePlan` and writes each validated ready area back through the **existing Raw Ingest v1 contract** via the `PersistenceRawIngestWriter` interface. The Raw Ingest area code is derived from the canonical area key (`rawIngestArea`), never authored independently, and each area is submitted to the same `area`/`start`/`count` identity it was loaded from. A non-`ready` plan is refused with no write; every Raw Ingest response is checked and the first non-`PersistenceRawIngestOK` (0x00) response, transport error or unknown area aborts immediately. No cross-area mirroring is added and no unseal occurs.
  - A `mma2composer/persistence_restore_test.go` (blob `1193029d53e01385ae49219e50fec165a6c28c03`) — restore self-check.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRestore -v .` → 7 tests PASS, exit 0 (`WritesEachAreaToItsIdentity`, `AddsNoCrossAreaMirroring`, `AbortsOnNonZeroResponse` for 0x10/0x12/0x14/0x20/0x21/0x30, `RefusesNonReadyPlan`, `SurfacesTransportError`, `RequiresWriter`, `AreaCodes`).
- Additional bounded real-socket check (temporary, removed before commit): the restore drove real Raw Ingest v1 packets over TCP and was verified to carry the exact area/start/count/payload identity, and a real classified response aborted the restore; passed with `-race`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- Acceptance mapping: (1) each snapshot area restores to the same area/start/count identity — `rawIngestArea` maps the canonical area key to the v1 code and the write carries `area.Start`/`area.Count` unchanged; (2) every Raw Ingest response is checked and a non-zero response aborts — the loop returns before any later area is written and records no successful write; (3) no cross-area mirroring is added — only the plan's own areas are written, exactly once each, with their own payloads.
- Design boundary: seal-flag protection is PERSIST-017, restore-completion verification PERSIST-018 and unseal PERSIST-019, so this task performs Raw Ingest transport only and never unseals. Placement is the shared schema owner `mma2composer` (consistent with PERSIST-012..015).
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: recorded in the repository commit for this task.

## CWAL
PERSIST-015 is delivered on GitHub main at `a497d9d4ca35cab8de832bc92fc3b5d90c7a4b33`. PERSIST-016 CODE is now complete and no longer ACTIVE; PERSIST-017 is the next eligible CODE packet. This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
