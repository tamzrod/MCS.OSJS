# MMA2 — Deterministic Modbus Memory Appliance

Repository-root component boundary for MCS.OSJS, sourced from the project-owned donor
`tamzrod/MMA2`. MMA2 is a **deterministic memory appliance** with strict and visible
boundaries: predictability, isolation, failure containment. Config is loaded once at
startup, immutable at runtime, and the sole source of truth. Four conceptual layers:
Configuration, Core Memory, Transport Adapters, Process Runtime. Full donor architecture
documentation is retained in `docs/` (00_OVERVIEW, 02_ARCHITECTURE, 04_CONFIGURATION,
05_MEMORY_MODEL, 06_TRANSPORTS, 07_FAILURE_MODEL, CONFIGURATION_MANUAL, RAW_INGEST,
PERSISTENCE).

## Provenance (donor harvest)

- Source: `tamzrod/MMA2` (main) at commit `a38687574645ada94c7a7f844152932afb09f147`.
- License: Apache-2.0 (see `LICENSE` in this directory); compatible with the MCS.OSJS
  Apache-2.0 project license; retain this notice per the harvest gate.
- Imported material: build/runtime source (`cmd/`, `internal/`, `tools/`, `benchmark/`),
  build scaffolding (`go.mod`, `go.sum`, `Dockerfile`, `.gitignore`), `README.md`
  (rewritten as the MCS component doc; donor architecture content preserved via `docs/`),
  `docs/` (needed for MMA2 configuration/run semantics), and the required `LICENSE` notice.
- MCS-local additions not present in the donor revision: `cmd/mma2-supervisor/`,
  `internal/restartwatch/`, `pkg/configvalidate/` (a composer-facing validation shim),
  `Dockerfile.supervised`, `testdata/smoke-test.yaml` and `docs/MCS_RBE_INTEGRATION.md`;
  these are maintained by MCS.OSJS and preserved across donor refreshes.
- Excluded: prebuilt donor `mma2` binary (build artifact; not source).
- Recorded in `THIRD_PARTY_NOTICES.md` per the MCS.OSJS licensing policy..

## Boundary

MMA2 is an independently bounded component:   it must stay decoupled from OS.js, Orchestrator,
and Replicator implementation. It consumes validated effective runtime configuration composed from
producer intents ((e.g. simulator, later Replicator( by the shared MMA2 configuration authority;
it does not directly overwrite the effective config file. Config is loaded once at startup and immutable at runtime.



## Clean build

Requires Go>=1.25.0 (per `go.mod`):

```bash
cd MMA2
go build -o mma2 ./cmd/mma2
```

The build uses only repository-owned `MMA2/` source+ the pinned module dependency
(`gopkg.in/yaml.v3` in `go.sum`(; it does not depend on the donor checkout at build or runtime..
