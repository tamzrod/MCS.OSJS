# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work.

OpenHands modes remain separate:
- **JR DEV** — assigned CODE/DISCOVERY implementation plus bounded workflow continuation.
- **Independent JR** — separately assigned TEST/VERIFY only.

The normal operator loop is:

```text
git pull
OPERATION CWAL
```

One invocation executes exactly one product task, may deliver/record it, may prepare the next already-promoted eligible task, then STOPS.

## Completed predecessors

**PERSIST-001 — Persistence Configuration Schema**
CODE COMPLETE at `faa33929429a0382a64b78bc773f0074e11cb6b1`.

**PERSIST-002 — State Sealing Prerequisite Validation**
CODE COMPLETE and delivered on GitHub main at
`fe3bda9e5c863c959500514e8c6083d2514526d1`.

**PERSIST-003 — Derived Persistence RBE Generation**
CODE COMPLETE at `4e8f774eb946cc9a946938e4f3a28eaa13de5c9d`, delivered on main.

**PERSIST-004 — Locked System RBE Behavior**
CODE COMPLETE at `936ab24560d57cbf8aacc8180bc00a2527dcfad5`, delivered on main.

**PERSIST-005 — Memory Range Synchronization**
CODE COMPLETE at `c743174204455afa83bc92b859fd47582a7afdbc`, delivered on main.

**PERSIST-006 — Memory Area Removal Synchronization**
CODE COMPLETE at `9ac48cb23e369fb0384b3d8f00811435d5bf7b69`, delivered on main.

**PERSIST-007 — Persistence Disable Cleanup**
CODE COMPLETE at `feb1e622d38cf8107a1df9ee182e5df17f95790f`, delivered on main.

**PERSIST-008 — User RBE Compatibility**
CODE COMPLETE at `4783e0d6e1d3dbd55b29b8647a115905eb5bc285`, delivered on main.

**PERSIST-009 — RBE ID Collision Handling**
CODE COMPLETE at `19057f02d94ff404b74a9f01348fc8c31235e374`, delivered on main.

**PERSIST-010 — Persistence Configuration UI**
CODE COMPLETE at `190e464f6106c3e640d21b1f9352328447464bef`, delivered on main by this invocation.

PERSIST-001..PERSIST-010 CODE tasks are complete. Their self-check/regression evidence remains JR DEV evidence only.

**PERSIST-012 — RBE-Triggered Snapshot Writer**
CODE COMPLETE at `2f18f8196f5497c5435b697b8b1570f59cab2642`, delivered on main.

**PERSIST-013 — Snapshot File Format**
CODE COMPLETE and delivered on GitHub main at `ccccb891b737995f6a24a58e844e5bd1d794b609` via GitHub connector recovery after the OpenHands push credential expired. OpenHands' local-only commits `9343d9e` / `ab27b27` are superseded by this remote recovery and must not be treated as additional product work.

**PERSIST-014 — Snapshot Manifest / Compatibility Metadata**
CODE COMPLETE and delivered on GitHub main at `98a952a89b087f12a9b99b6bd648ce97ac44e6f2` via GitHub connector recovery after the OpenHands credential remained invalid. OpenHands' local-only commits `5f944bd` / `17648ce` are superseded by this remote recovery and must not be replayed as additional product work.

## Completed independent gate

**PERSIST-011 — Persistence Configuration Tests**
Independent JR TEST/VERIFY **PASS** on tested HEAD `313e3be60a29fff6c967d78c0cf677b8b888ed56`, with pinned product checkpoint `190e464f6106c3e640d21b1f9352328447464bef`.

Verified once, in order:
- `cd mma2composer && go test -mod=readonly ./...` — exit 0
- `cd simulator && go test -mod=readonly ./...` — exit 0
- `cd replicator && go test -mod=readonly ./...` — exit 0
- `cd OSJS && node tests/toolkit-persistence-ui.test.js` — exit 0
- `cd OSJS && node tests/toolkit-ui-parity.test.js` — exit 0
- `cd OSJS && node tests/toolkit-fc43.test.js` — exit 0

Freshness/post-check passed: commits after `190e464` were workflow-only, working tree clean, no product/test modification by independent JR. This PASS covers PERSIST-001..010 configuration behavior only; runtime persistence remains for PERSIST-012+.

## Completed predecessor — PERSIST-015

**PERSIST-015 — Startup Snapshot Loader**
CODE COMPLETE at `a497d9d4ca35cab8de832bc92fc3b5d90c7a4b33`, delivered on main; awaiting independent TEST/VERIFY (PERSIST-022).
- `mma2composer/persistence_snapshot_loader.go` (blob `29cf1950cbec796508f1ab3c7c214c6ed5f30208`) and `mma2composer/persistence_snapshot_loader_test.go` (blob `8ff2147ef88ca29b499e3c07d897e54fc07a47f0`).
- `LoadPersistenceSnapshots` loads/validates each configured persisted area at startup only for persistence-enabled **sealed** memories; missing/invalid/incompatible snapshots yield distinct explicit restore states; aggregate state is `ready` only when every configured area is ready; no Raw Ingest write or unseal.
- Self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceSnapshot .` → `ok`, exit 0. Bounded regression: full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-016 — Raw Ingest Restore**
CODE COMPLETE by OpenHands JR DEV in this invocation (source commit recorded in the task packet); awaiting independent TEST/VERIFY (PERSIST-022).
- `mma2composer/persistence_restore.go` (blob `3c31ade58d4a0f7c5a82f40c6873244d15002586`) and `mma2composer/persistence_restore_test.go` (blob `1193029d53e01385ae49219e50fec165a6c28c03`).
- `RestorePersistencePlan` writes each validated ready area of the PERSIST-015 `PersistenceRestorePlan` back through the **existing Raw Ingest v1 contract** (`PersistenceRawIngestWriter`); the v1 area code is derived from the canonical area key and each write carries the same area/start/count identity it was loaded from. A non-`ready` plan is refused with no write; every Raw Ingest response is checked and the first non-`0x00` response, transport error or unknown area aborts immediately; no cross-area mirroring and no unseal.
- Self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRestore -v .` → 7 tests PASS, exit 0. Additional temporary real-socket v1 check passed with `-race` (removed before commit). Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

## Current task — ACTIVE

**PERSIST-017 — Seal-Flag Protection During Restore**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-017-seal-flag-protection-during-restore.md`
Queue: `workflow/active_work/PERSISTENCE_PROMOTION_QUEUE.md`

Execute PERSIST-017 exactly as written:
- prevent the persisted sealing-flag value from causing a premature unseal during snapshot restoration;
- keep the configured State Sealing address sealed/zero throughout area restoration;
- a snapshot containing prior flag=1 must not unseal during restore;
- follow the authoritative State Sealing area/address configuration; add no second sealing source of truth and no duplicate flag setting;
- run only task-bounded self-checks and bounded in-scope corrective retests;
- commit and non-force push under standing JR DEV authority;
- record changed paths/checks/source checkpoint;
- prepare PERSIST-018 for the next invocation only after genuine PERSIST-017 completion;
- STOP after PERSIST-017.

Recovery note for a reused OpenHands workspace: local-only PERSIST-013 commits `9343d9e` / `ab27b27` and local-only PERSIST-014 commits `5f944bd` / `17648ce` are superseded by authoritative remote recovery commits `ccccb891` and `98a952a8`. Do not replay them as new work; reconcile only with bounded safe Git mechanics.

## Successor routing

PERSIST-018..022 are already human-promoted and remain QUEUED/dependency-gated. PERSIST-011 passed independently; PERSIST-012 through PERSIST-016 are CODE complete; PERSIST-017 is now the sole current ACTIVE JR DEV CODE task.

OpenHands JR DEV may autonomously close and deliver CODE/DISCOVERY packets and select the next already-promoted eligible CODE/DISCOVERY packet for the next invocation.

TEST/VERIFY packets switch to independent JR mode and require their exact current packet. JR DEV must not self-certify those gates.

Only BLACK SHEEP WALL edits ICC. No force push, destructive history rewrite, production/operator-data mutation, or scope expansion.
