# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work.

OpenHands modes remain separate:
- **JR DEV** — assigned CODE/DISCOVERY implementation plus bounded workflow continuation.
- **Independent JR** — separately assigned TEST/VERIFY only.

The normal operator loop for this already-promoted autonomous queue is:

```text
OPERATION CWAL
```

OpenHands JR DEV may perform its authorized safe fetch/fast-forward itself. One invocation executes exactly one product task, delivers/records it, activates the named already-promoted successor for the next invocation, then STOPS.

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
CODE COMPLETE at `2fbabc3481b6435ca38596ea7dd060e36ed0655d`, delivered on main; awaiting independent TEST/VERIFY (PERSIST-022).
- `mma2composer/persistence_restore.go` (blob `3c31ade58d4a0f7c5a82f40c6873244d15002586`) and `mma2composer/persistence_restore_test.go` (blob `1193029d53e01385ae49219e50fec165a6c28c03`).
- `RestorePersistencePlan` writes each validated ready area of the PERSIST-015 `PersistenceRestorePlan` back through the **existing Raw Ingest v1 contract** (`PersistenceRawIngestWriter`); the v1 area code is derived from the canonical area key and each write carries the same area/start/count identity it was loaded from. A non-`ready` plan is refused with no write; every Raw Ingest response is checked and the first non-`0x00` response, transport error or unknown area aborts immediately; no cross-area mirroring and no unseal.
- Self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRestore -v .` → 7 tests PASS, exit 0. Additional temporary real-socket v1 check passed with `-race` (removed before commit). Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-017 — Seal-Flag Protection During Restore**
CODE COMPLETE at `fa1276c26cbb639f061436c3f111351232037b57`, delivered on main; awaiting independent TEST/VERIFY (PERSIST-022).
- `mma2composer/persistence_restore.go` (blob `2f14d6bd4132d15e3607afe9311284c30828d964`), `mma2composer/persistence_snapshot_loader.go` (blob `1d62df77e3be5d9e26eff0c2fcf1d09636fe76ac`), `mma2composer/persistence_restore_test.go` (blob `03ac862771bf67e80370e81ff2185540977972ec`).
- `RestorePersistencePlan` forces the authoritative State Sealing bit to sealed (0) within the restored coils payload before writing it, so a snapshot captured while unsealed cannot unseal the memory mid-restore. `PersistenceSealingFlagFromExtra` derives the flag location from the configured `state_sealing` block (area `"coil"`, address) and introduces no second sealing source of truth; `forcePersistenceSealingFlag` clears exactly that one bit, preserves every other restored bit, and fails closed if the flag lies beyond the payload. `PersistenceRestorePlan` gained an optional `SealingFlag`.
- Self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag' -v .` → 12 tests PASS, exit 0. Additional temporary real-socket v1 check confirmed the sealing bit is cleared on the wire (passed with `-race`, removed before commit). Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-018 — Restore Verification**
CODE COMPLETE at `3eb118f507688d483c5d4985f74fe6a98b531bc3`, delivered on main; awaiting independent TEST/VERIFY (PERSIST-022).
- `mma2composer/persistence_restore.go` (blob `8a3d716e24ed5747499137e1d3a5d8eb42218513`) and `mma2composer/persistence_restore_test.go` (blob `80ed02387f88958e1f49d46fb78cef8d668fbd86`).
- `PersistenceRestoreResult` records `RequiredAreas` (deterministic configured area set in plan order) and `AcknowledgedAreas` (subset written and acknowledged with the Raw Ingest success code). `RestorePersistencePlan` emits `Completed` only through `VerifyPersistenceRestore`, which requires a non-empty required set exactly equal to the acknowledged set, so any failed or missing area prevents completion. No unseal.
- Self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag|TestVerifyPersistenceRestore' -v .` → 18 tests PASS, exit 0. Additional temporary real-socket v1 check confirmed the gate (all-ack → completed; one rejected → not completed), passed with `-race`, removed before commit. Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-019 — Atomic Unseal / Commit Step**
CODE COMPLETE at `e22459ecdd10433ef21f63df25e3a22a43fb3662`, delivered on main; awaiting independent TEST/VERIFY (PERSIST-022).
- `mma2composer/persistence_restore.go` (blob `7302908c117091ef3755304e9101f12c9527583e`) and `mma2composer/persistence_restore_test.go` (blob `fb4f9e77b28095e08abaf2f55ca3b5c652d12117`).
- After `VerifyPersistenceRestore` success, `RestorePersistencePlan` performs the explicit final commit: a single-coil Raw Ingest write of the authoritative State Sealing flag to unsealed (1) at the configured address, after every area restore. `PersistenceRestoreResult.Committed` is true only when that write is acknowledged with `0x00`; a rejected unseal leaves the restore completed but not committed. The flag location comes solely from configuration; no runtime reseal or alternate commit flag.
- Self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag|TestVerifyPersistenceRestore|TestEncodePersistenceSealingFlag' -v .` → 23 tests PASS, exit 0. Additional temporary real-socket v1 check confirmed the final wire packet is the unseal at the configured address (passed with `-race`, removed before commit). Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-020 — Restore Failure Behavior**
CODE COMPLETE at `1528d3e2451f3588de46f7539127164588878f1a`, delivered on main; awaiting independent TEST/VERIFY (PERSIST-022).
- `mma2composer/persistence_restore.go` (blob `023d89ab1898b4a1414a33929a851459d5f48e35`) and `mma2composer/persistence_restore_test.go` (blob `55d9e143b75460a5bd7c28e4944db989750cd4a3`).
- New `PersistenceRestoreFailure` classification and `failedPersistenceRestore` helper; every non-committing path in `RestorePersistencePlan` leaves the memory sealed and returns a deterministic `Failure` (disabled, unsealed, empty, missing/invalid/incompatible snapshot, area_not_ready, unknown_area, raw_ingest_write/response, incomplete, no_sealing_flag, unseal_write/response) plus a `Detail`. `PersistenceRestoreResult` gains `Failure` and `Sealed`; a committed restore has `Failure` None and `Sealed` false. No fabricated defaults, retry loops, alternate unseal paths or operator-data mutation.
- Self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag|TestVerifyPersistenceRestore|TestEncodePersistenceSealingFlag' -v .` → 30 tests PASS, exit 0. Additional temporary real-socket v1 check confirmed a rejected area write keeps the memory sealed and never emits the unseal packet (passed with `-race`, removed before commit). Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-021 — Persistence Runtime Status**
CODE COMPLETE at `1310c1367c7f98f084187006084d624d8e20e07b`, delivered on main.
- `mma2composer/persistence_status.go` (blob `56ceb22b5026fea300c408262e26364554fda5a9`), `simulator/apply.go` (blob `c35ae093c9b8580f51f2a0cdcd314bbe07830e25`), `OSJS/src/packages/MCSModbusToolkit/diagnostics-model.js` (blob `84970aa4246a8caf8e89996b30e67a4030676c66`), `OSJS/src/packages/MCSModbusToolkit/diagnostics-editor.js` (blob `2d46c63609ddd3823db26f0dd958bcf0b952721c`), plus tests.
- `PersistenceRuntimeStatus` is a pure read-only projection from the authoritative restore plan/result: `Configured`, authoritative `Sealed`, `Healthy` (only when committed), `SnapshotHealth`, classified `RestoreOutcome`, and optional `LastSave`/`LastRestore`. Surfaced through the simulator `DeviceRuntimeStatus.Persistence` and the OS.js diagnostics model/editor (fail-closed to UNKNOWN; controls remain disabled). Restore failure stays visible while sealed stays authoritative; no bypass of restore/sealing gates.
- Self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRuntimeStatus -v .` → 7 tests PASS, exit 0; simulator persistence status tests PASS; OS.js diagnostics tests exit 0. Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites and the whole OSJS node suite → all pass, exit 0. New Go files `gofmt`/`vet` clean.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

## Current task — ACTIVE

**PERSIST-R01 — Filesystem Snapshot Adapter**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-r01-filesystem-snapshot-adapter.md`

Human promoted the full optimized persistence continuation chain on 2026-10-08:

```text
PERSIST-R01  Filesystem Snapshot Adapter
    ↓
PERSIST-R02  Runtime Save Wiring
    ↓
PERSIST-R03  Startup Restore Wiring
    ↓
PERSIST-UI01 Electron Persistence Settings
    ↓
PERSIST-R04  Runtime Status Wiring
    ↓
PERSIST-R05  Disposable Persistence E2E Harness
    ↓
PERSIST-022  Independent End-to-End VERIFY
```

Exactly one product task executes per OPERATION CWAL invocation to keep context bounded. After genuine delivery of a CODE packet, OpenHands JR DEV **must** update completion evidence and arm the named already-promoted successor for the next invocation without asking the human to promote/select it again.

The normal operator action is now simply:

```text
OPERATION CWAL
```

JR DEV may perform its already-authorized safe fetch/fast-forward itself. A reused workspace must still preserve the superseded PERSIST-013/PERSIST-014 local-history recovery notes and must never force/reset destructively.

## Autonomous successor routing

- R01 success → activate R02
- R02 success → activate R03
- R03 success → activate UI01
- UI01 success → activate R04
- R04 success → activate R05
- R05 success → pin and activate PERSIST-022 as **independent JR VERIFY**, then STOP
- Any genuine blocker/failure → fail closed, record it, do not skip ahead

JR DEV never executes PERSIST-022 and never self-certifies the VERIFY gate. Only BLACK SHEEP WALL edits ICC.
