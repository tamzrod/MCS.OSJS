# NPE-05A — Refresh Embedded MMA2 Native Persistence Runtime

Status: CODE COMPLETE — delivered on main at `1ad2510312ccba0a09024b70c53a1e9945689957` (OpenHands JR DEV); awaiting independent verification (NPE-06)
Stage: CODE
Owner: OpenHands JR DEV
Previous: NPE-05
Next: NPE-06

## Problem

The repository-root `MMA2/` component is a donor harvest pinned to an old `tamzrod/mma2` commit and does not contain the current native persistence implementation.

The authoritative donor repository `tamzrod/MMA2` main currently contains the native per-memory persistence runtime and verified contract. NPE-06 was blocked because MCS.OSJS is still building/testing the stale embedded MMA2 copy.

This is a source synchronization problem, not a request to invent a second persistence implementation.

## Authoritative donor

Repository: `tamzrod/MMA2`
Pinned donor commit for this task:

```
a38687574645ada94c7a7f844152932afb09f147
```

Native persistence at that checkpoint includes:
- per-memory `persistence` configuration;
- optional directory with YAML-folder default;
- omitted ranges = all allocated areas;
- native disk snapshot runtime;
- startup restore before ingress;
- committed-write observation/runtime save;
- primary + backup recovery;
- independence from State Sealing;
- independence from RBE TCP.

## Scope

Refresh the repository-root `MMA2/` component from the pinned donor commit while preserving the MCS.OSJS component boundary and licensing/provenance requirements.

Import the donor source required to build and run the current MMA2 runtime, including all source/config/runtime packages required by native persistence.

At minimum reconcile:
- `MMA2/cmd/`
- `MMA2/internal/`
- `MMA2/pkg/` if present/required
- `MMA2/go.mod`
- `MMA2/go.sum`
- relevant MMA2 docs needed to describe the runtime/config contract
- `MMA2/LICENSE`

Update `MMA2/README.md` provenance to the exact pinned donor SHA.

Do not modify MMA2 semantics to fit stale MCS assumptions. The donor contract is authoritative for MMA2 runtime behavior.

## Preserve

- MCS-specific integration remains outside MMA2.
- MMA2 must not learn about Electron, OS.js, Simulator, Replicator or appliance workflow.
- No external persistence manager is added.
- State Sealing remains independent.
- RBE remains independent.
- No user snapshot data is deleted.
- Preserve Apache-2.0 licensing/notices.

## Reconciliation

After donor refresh, reconcile only compilation/config-boundary breakage caused by the updated MMA2 public/internal configuration contract.

Do not reintroduce:
- persistence requiring State Sealing;
- persistence-generated RBE;
- Raw Ingest restore orchestration outside MMA2;
- unlock-coil persistence behavior;
- external persistence watchdogs.

## Acceptance

Required self-checks:

1. `MMA2/internal/config/config.go` contains native per-memory persistence configuration.
2. `MMA2/internal/persistence/` exists and builds as part of MMA2.
3. MMA2 builds successfully from repository-owned source.
4. MMA2 native persistence unit/integration tests imported with the donor source, where applicable to the component boundary, pass.
5. MCS composer/simulator/replicator tests affected by the refreshed MMA2 contract pass.
6. Electron focused native-persistence tests remain green.
7. Source search proves native persistence runtime resides inside `MMA2/`, not Electron or Simulator.
8. Provenance documentation records donor `a38687574645ada94c7a7f844152932afb09f147`.

Record exact changed paths, commands, results and delivered SHA.

After genuine delivery, update handoff so NPE-06 becomes current Independent TEST/VERIFY and STOP at the identity boundary.

## Delivery evidence (OpenHands JR DEV)

Delivered SHA: `1ad2510312ccba0a09024b70c53a1e9945689957` on `main` (base `cd6c940`). 94 files changed, +23559/-1365.

Method: cloned the donor `tamzrod/MMA2`, checked out the pinned commit `a38687574645ada94c7a7f844152932afb09f147`, and overlaid its file set (excluding `.git/` and `.github/`) onto the repository-root `MMA2/`, preserving the MCS-added files absent from the donor revision.

Acceptance self-checks:
1. `MMA2/internal/config/config.go` contains native per-memory persistence configuration — `MemoryDefinition.Persistence *MemoryPersistenceConfig` (and a root `Persistence` block explicitly rejected). ✓
2. `MMA2/internal/persistence/` exists and builds — `cd MMA2 && go build ./internal/persistence/` → BUILD OK. ✓
3. MMA2 builds from repository-owned source — `cd MMA2 && go build ./...` → BUILD OK. ✓
4. MMA2 native persistence tests pass — `cd MMA2 && go test -mod=readonly ./...` → all `ok`, including `mma2/internal/persistence` (0.669s), `mma2/internal/config`, `mma2/pkg/configvalidate`. ✓
5. MCS composer/simulator/replicator/mma2raw build and tests pass against the refreshed contract — `mma2composer`, `simulator` (incl. `cmd/modbus-simulator-runtime`), `replicator` (incl. `cmd/modbus-replicator-runtime`), `mma2raw` all `ok`, exit 0; no reconciliation breakage required. ✓
6. Electron focused tests remain green — `cd electron && node --test test/*.test.js` → 60 PASS, 0 FAIL. ✓
7. Source search proves the native persistence runtime resides inside `MMA2/` (`MMA2/internal/persistence/`, `MMA2/internal/config/persistence*.go`); `electron/` retains no persistence runtime (removed in NPE-05); no Electron/OS.js/Simulator/Replicator references inside `MMA2/internal|cmd|pkg`. ✓
8. Provenance updated — `MMA2/README.md` and `THIRD_PARTY_NOTICES.md` record donor `a38687574645ada94c7a7f844152932afb09f147`. ✓

Notes:
- `go.mod`, `go.sum` and `LICENSE` were identical in the pinned donor and are unchanged.
- MCS-local additions preserved: `cmd/mma2-supervisor/`, `internal/restartwatch/`, `pkg/configvalidate/`, `Dockerfile.supervised`, `testdata/smoke-test.yaml`, `docs/MCS_RBE_INTEGRATION.md`, and the MCS component `.gitignore`.
- `gofmt -l` lists 39 donor files as unformatted; this is the authoritative donor state and was intentionally not reformatted (reformatting would deviate from the pinned donor).
- Observation for a later task, not a change in this refresh: the previously-built Simulator/`mma2composer` external persistence orchestration (PERSIST-012..022, R01..R05, e.g. `simulator/persistence_bootstrap.go`, `persistence_startup.go`, `persistence_watchdog.go`, `mma2composer/persistence_*.go`) still exists and runs at simulator apply/startup beside MMA2. Under the current contract ("MMA2 is the sole persistence runtime owner"), removing that stale external orchestration is future work (the NPO-03/legacy-cleanup lineage), not part of this donor refresh.

This is JR DEV self-check evidence only, not an independent TEST/VERIFY PASS.
