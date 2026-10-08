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

OpenHands JR DEV may perform its authorized safe fetch/fast-forward itself. **Recursive JR DEV continuation is enabled for the current promoted persistence CODE chain.** One OPERATION CWAL invocation should continue task-to-task after each genuine delivery, using a context-reset checkpoint between packets. Stop only on failure/blocker, queue end, or before the PERSIST-022 independent-VERIFY identity boundary.

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

**PERSIST-R01 — Filesystem Snapshot Adapter**
CODE COMPLETE at `1ac1477177ec19fa02cb344d6a20a8632fc0c88a`, delivered on main.
- `mma2composer/persistence_filesystem.go` (blob `d99a8f9c8254e64f678133ccf8234289584307e3`) and `mma2composer/persistence_filesystem_test.go` (blob `aae87093f997a2f04e64c0bc1623cde199ab3ecd`).
- `PersistenceFilesystemAdapter` implements both `PersistenceSnapshotStore` and `PersistenceSnapshotSource`, rooted beneath the appliance data root at `persistence/snapshots/port-<P>/unit-<U>/area-<A>/`, reusing the existing raw snapshot format and manifest metadata. Registration rejects non-canonical/duplicate areas; writes splice changed runs and atomically replace `snapshot.bin` + `manifest.json`; malformed/missing content fails closed.
- Self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceFilesystem -v .` → 14 tests PASS, exit 0. Bounded regression: full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt`/`vet` clean.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-R02 — Runtime Save Wiring**
CODE COMPLETE at `407fcc32979b9acd83dd6826e86f4b700da70c1b`, delivered on main.
- `mma2composer/persistence_runtime_save.go` (blob `3dc796dab80259a75b5f75ebbea327a915c89375`), `simulator/persistence_save.go` (blob `a4d8112419dffa7b71544a560093cf7315c7d1c1`), `simulator/persistence_area_reader.go` (blob `a06e45bcaff246ac1132151d133d9b17462180f6`), `simulator/persistence_rbe_subscriber.go` (blob `4bc66ef653ef73f01f3e5fc9d96d010b82d6b9e0`), `simulator/apply.go` (blob `4f93dd73547cbbbab65d3ceff5e7d0beacc19f6e`), `simulator/cmd/modbus-simulator-runtime/main.go` (blob `dc8293978abf48cf4c7ebbe7e857aa8df7366647`), `mma2composer/persistence_filesystem.go` (blob `229cc504f73306bc4c29bef227b8b171b60b208d`), plus tests.
- Persistence-owned RBE events are wired into the existing `PersistenceSnapshotWriter` via the real `SchedulerApplier.PublishPersistenceEvent` call site; the authoritative-state reader performs a real Modbus FC1/FC3 read and never fabricates bytes; only the derived/system-owned projection is subscribed.
- Self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRuntimeSave -v .` → 4 tests PASS; `cd simulator && go test ... 'TestPersistenceRuntimeSave|TestPersistenceRBESubscriber|TestModbusReadArea|TestPersistenceAreaReader'` → 9 tests PASS. Bounded regression: all Go modules and the OSJS node suite pass, exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-R03 — Startup Restore Wiring**
CODE COMPLETE at `2236e15277ebb17c60cac206e05d77e17e296a53`, delivered on main.
- `mma2composer/persistence_startup.go` (blob `cb9d739b52d74c0be467465bfbdf3d75bc8441b9`), `simulator/persistence_raw_ingest.go` (blob `685a90ba7396daa41cefcb98831ada33a0c98191`), `simulator/persistence_startup.go` (blob `c97533891196f9bf194c7504a0f24fb2482f7037`), `simulator/apply.go` (blob `c27a09e154ed6430d4d31dcf8ff3a8d977829685`), plus tests.
- `RestorePersistenceAtStartup` is the real startup orchestration: load/validate durable snapshots while sealed, attach the authoritative State Sealing flag from config, restore through the existing Raw Ingest v1 contract, verify the full required-area set, and perform the existing final explicit unseal only after success. The simulator `newRuntimeApplyRouter` is the real non-test startup call site.
- Self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceRestoreAtStartup -v .` → 3 tests PASS; `cd simulator && go test -mod=readonly -run TestPersistenceStartup .` → 4 tests PASS. Bounded regression: all Go modules pass, exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-UI01 — Electron Persistence Settings**
CODE COMPLETE at `cff34f41ebb350fc1c4380253cc2acc02e4562ca`, delivered on main.
- `electron/renderer/memory-advanced.js` (blob `de068160c35b38ea71cb9e4ba693762485188a5d`), `electron/renderer/app.js` (blob `adb0001e0f52cc00807f2fcf195beb11f61a631f`), plus focused renderer tests.
- Advanced Settings now offers the Persistence tab (after Access Policy) only where supported; it exposes `persistence.enabled` through the same draft, reports RBE mechanism + State Sealing prerequisites (never silently enabling sealing or creating operator-owned RBE rules), and renders derived persisted areas/locked system-owned RBE rows read-only.
- Self-check: `cd electron && node test/memory-advanced.test.js` → 12 PASS; `node test/memory-persistence.test.js` → 4 PASS; all Electron test files pass.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

**PERSIST-R04 — Runtime Status Wiring**
CODE COMPLETE at `d9251d3c56bf490c21655bdb372b342ee1eec7f8`, delivered on main.
- `mma2composer/persistence_status.go` (blob `6d0b48dd2ff891bd38ebff8a56036bde8a086c0c`), `mma2composer/persistence_runtime_save.go` (blob `4401194e0b40057d8dfa9ee328a5b3f71764e25b`), `simulator/persistence_save.go` (blob `f408910f6df9e5d6ebf8cfb73c3781ab8c81086e`), `simulator/apply.go` (blob `7df13e87ce43cfab53374a8c73b93e20706cd7f1`), plus focused tests.
- `RuntimeStatus` now feeds real save/restore observations into `DeviceRuntimeStatus.Persistence` (not a configured-only placeholder); last save/restore appear only when genuinely observed.
- Self-check: `mma2composer` `TestPersistenceRuntimeStatus|TestPersistenceRuntimeSave` → ok; `simulator` `TestRuntimeStatusCarriesRealPersistenceObservations|TestPersistenceStartup` → ok. Bounded regression: all Go modules and OS.js+Electron node suites pass, exit 0.
- This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.

## Current task — none ACTIVE for JR DEV (CODE chain complete)

All CODE packets **PERSIST-001..021 + PERSIST-R01..R05 + PERSIST-UI01 are CODE COMPLETE**. JR DEV has no further CODE task and STOPS at the VERIFY identity boundary.

## Next task — PERSIST-022 independent-JR VERIFY packet (pinned)

**PERSIST-022 — End-to-End Persistence Verification**
Mode / owner: **independent JR** (separate invocation) — JR DEV must not execute or certify this.
Stage: VERIFY. Previous: PERSIST-R05. Next: none.
Packet: `workflow/active_work/persist-022-end-to-end-persistence-verification.md`

### Pinned execution packet (exact)

- Source checkpoint to verify (pinned): `c8f31b2443b69201f36464db72497ed8271aeb7d` (`main`, the PERSIST-R05 delivery commit; tree otherwise workflow-only).
- Disposable target: real MMA2 built from source by the harness; loopback ports and `t.TempDir()` roots only. No production/customer data, no global services, no network output.
- Preflight (independent JR): confirm `git rev-parse HEAD` ancestry includes the pinned checkpoint; working tree clean; Go toolchain available (`go version`); network not required beyond loopback.
- Exact ordered actions/commands:
  1. `cd simulator && go test -mod=readonly -count=1 ./...` — expect `ok`, exit 0 (gated harness SKIP is expected here).
  2. `cd simulator && MCS_RUN_PERSIST_E2E=1 go test -mod=readonly -count=1 -run TestPersistenceDisposableEndToEnd -v .` — expect `--- PASS: TestPersistenceDisposableEndToEnd`, exit 0.
- Expected observations (from the harness output/assertions):
  1. live phase: no-snapshot first boot stays sealed; Modbus read rejected `0x06`; known value `0x1234` written and unsealed reads back `0x1234`; a real RBE save is observed.
  2. restart phase: fresh MMA2 sealed before restore; startup restore → verify → final unseal commits; `Healthy` true, `Sealed` false; Modbus reads restored `0x1234`.
  3. fail-closed phase: corrupted snapshot → restore fails closed; `Healthy` false, `Sealed` true, non-`none` classified reason, Modbus still rejected `0x06`.
- Evidence destination / report permission: this packet is chat-only unless the workflow owner separately pins a report path; independent JR captures raw stdout/stderr + exit codes and returns the verdict. JR changes no product source.
- Cleanup: the harness cleans only its own `t.TempDir()` roots and disposable processes.

### Existing evidence (JR DEV self-check — NOT the independent gate)

- PERSIST-R05 delivered `simulator/persistence_e2e_test.go` (blob `7c833aa20fe3f6ee6005b4f6633b5d4cecc3b8e9`); the exact command was run once by JR DEV and observed `--- PASS` (real MMA2, exit 0). This is context only and does not substitute for independent JR verification.

## Successor routing

The autonomous CODE chain `R01 → R02 → R03 → UI01 → R04 → R05` is complete. Independent TEST/VERIFY for PERSIST-011(PASS) and PERSIST-012..021 + R01..R05 + UI01 remains deferred to PERSIST-022, now pinned above for the next **independent-JR** invocation. JR DEV STOPS here.

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

## Autonomous successor routing — RECURSIVE JR DEV ENABLED

Current continuation:
- UI01 success → activate and immediately execute R04
- R04 success → activate and immediately execute R05
- R05 success → pin and activate PERSIST-022 as **independent JR VERIFY**, then STOP
- Any genuine blocker/failure → fail closed, record it, do not skip ahead

Between CODE tasks, perform a context-reset checkpoint: reread this handoff, read only the newly active packet, and load only its minimal relevant source/context. Prior task implementation detail is evidence, not active reasoning context.

JR DEV never executes PERSIST-022 and never self-certifies the VERIFY gate. Only BLACK SHEEP WALL edits ICC.
