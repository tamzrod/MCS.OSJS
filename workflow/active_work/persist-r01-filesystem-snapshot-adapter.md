# PERSIST-R01 — Filesystem Snapshot Adapter

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-021
Next: PERSIST-R02

## Primary outcome
Implement the concrete appliance filesystem adapter for persistence snapshot payloads and manifest metadata using the existing persistence snapshot/store/source contracts.

## Scope
- Implement concrete `PersistenceSnapshotStore` / `PersistenceSnapshotSource` behavior needed by the runtime.
- Root persistence data only beneath the authoritative appliance data root using deterministic Port → Unit ID → Area paths.
- Reuse the existing raw snapshot format, manifest metadata, size validation and checksum contracts; do not invent a second format.

## Non-scope
No RBE runtime subscription, startup restore orchestration, runtime status wiring, Electron UI, repair/retry loop, production/customer data, or ICC edits.

## Acceptance
1. Snapshot payload + manifest can be written/read deterministically for one `(Port, UnitID, Area)` without escaping the appliance data root.
2. Existing manifest/length/checksum validation remains authoritative; malformed/missing files surface explicit errors rather than fabricated defaults.
3. Focused tests cover deterministic paths, round-trip, missing/corrupt content and safe replacement behavior without changing existing snapshot semantics.

## Evidence / handoff
Record exact changed paths, focused tests, bounded regression and delivered source SHA. This is JR DEV self-check only. After genuine delivery, automatically arm PERSIST-R02 for the next invocation and STOP.

## Dependencies
PERSIST-001..021 CODE complete. Human has promoted the full PERSIST-R01..R05 + PERSIST-UI01 chain.

## Sizing
2/0/1/1/0=4.

## Coding evidence (PERSIST-R01, OpenHands JR DEV)

- Source checkpoint base: `02feda224dd64e14849bb50926c9bdcb60d539da` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - A `mma2composer/persistence_filesystem.go` (blob `d99a8f9c8254e64f678133ccf8234289584307e3`) — `PersistenceFilesystemAdapter` implements both `PersistenceSnapshotStore` and `PersistenceSnapshotSource`. Rooted beneath the appliance data root at the deterministic `persistence/snapshots/port-<P>/unit-<U>/area-<A>/` layout (`RelPersistenceDir`, `RawSnapshotPath`, `ManifestPath`). Reuses the existing raw snapshot format (`PersistenceSnapshotSize`, `NewPersistenceSnapshotManifest`) and never invents a second format or range. Registration rejects non-canonical area keys, kind mismatches, empty ranges and duplicates. `WritePersistenceBytes` splices changed runs into the area image and atomically replaces `snapshot.bin` + `manifest.json` (recomputing length/SHA-256 over the whole image); a wrong-size existing image, an out-of-range write, a manifest without its payload, or corrupt manifest JSON all fail closed with explicit errors. `ReadPersistenceSnapshot` returns present=false only when no snapshot exists.
  - A `mma2composer/persistence_filesystem_test.go` (blob `aae87093f997a2f04e64c0bc1623cde199ab3ecd`) — focused self-check.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run TestPersistenceFilesystem -v .` → 14 tests PASS, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- Acceptance mapping: (1) snapshot payload + manifest write/read deterministically for one (Port, UnitID, Area) without escaping the appliance data root — the layout is fixed and registration rejects authored/non-canonical area names; (2) existing manifest/length/checksum validation remains authoritative and malformed/missing files surface explicit errors — the adapter performs no fabrication and the loader's `ValidatePersistenceSnapshotManifest` accepts the round-tripped content; (3) focused tests cover deterministic paths, round-trip, missing/corrupt content and safe replacement without changing existing snapshot semantics, including an end-to-end drive of the existing `PersistenceSnapshotWriter` change-only path.
- Design boundary: this task provides the concrete filesystem adapter only. Runtime save wiring is PERSIST-R02, startup restore wiring PERSIST-R03, UI PERSIST-UI01, status wiring PERSIST-R04 and the disposable E2E harness PERSIST-R05. No RBE subscription, startup orchestration, UI, retry loop or ICC edit is introduced.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `1ac1477177ec19fa02cb344d6a20a8632fc0c88a` on GitHub main.

## CWAL
PERSIST-R01 CODE is complete and no longer ACTIVE; PERSIST-R02 is armed ACTIVE for the next invocation. This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
