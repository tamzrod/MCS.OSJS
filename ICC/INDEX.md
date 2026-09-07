# ICC Index

## Purpose

This directory is the Incremental Context Compaction (ICC) cache for MCS.OSJS. ICC stores compact semantic context derived from repository truth so operations can read context first instead of repeatedly reopening unchanged source files. ICC is also a context-zoom mechanism. An operation selects the semantic boundary required by the current authorized work, then stays inside that branch and zooms deeper only when execution requires more detail.

## Repository Baseline

Every completed BLACK SHEEP WALL audit must record:

- `Baseline Commit`: repository HEAD used as the committed source baseline;
- `Working Tree`: `clean` or `dirty`;
- `Audited Uncommitted Overlay`: changed paths incorporated into ICC, with a deterministic content hash or equivalent fingerprint;
- context registry entries and their source dependencies.

The baseline commit identifies the committed repository state represented by ICC. It does not need to equal the later commit that stores ICC files.

## Registry

Baseline Commit: cfe704c
Working Tree: dirty
Audited Uncommitted Overlay: `simulator/raw_ingest.go` (sha256 ecb30acb1acc090f)`, `simulator/raw_ingest_test.go` (sha256 6c646d7e4c343f65)`

Committed delta: none since this session's start (shallow clone HEAD cfe704c).

SIM-004 completed+verified 2026-09-07. Added `simulator/raw_ingest.go` (MMA2 raw-ingest v1 client; encodeRawPacket bits-LSB-first/regs big-endian; Send( drives simulator-generated FC1-FC4 values into simulator-owned MMA2 memory exclusively via raw ingest( and `simulator/raw_ingest_test.go` (4 tests: bits/regs encoding frames + real-TCP Send round-trip with OK/reject responses + unconfigured-FC rejection(;gofmt/vet clean; `go test -count=1 ./...` -> ok; SIM-004 marked COMPLETED in handoff; CURRENT advanced to SIM-005;sim-004 task file records verification evidence.. Refreshed `context/active-work.md` (SIM-004 COMPLETED;SIM-005 CURRENT( and `context/simulator-device-config.md`(added SIM-004 raw-ingest established-truth section;( title now SIM-001+SIM-002A+SIM-003+SIM-004(. Unrelated registry rows preserved.

| Context | Parent | Zoom In | Source Dependencies |
| --- | --- | --- | --- |
| `context/L0-project.md` | none | governance, donor-licensing, network-exposure, planning-workflow, osjs-shell, active-work | `README.md`, `PROJECT_IDENTITY.md`, `handoff.md` |
| `context/governance.md` | L0-project | — | `AGENTS.md`, `BLACK_SHEEP_WALL.md`, `ICC/INDEX.md`, `operation cwal.md`, `handoff.md`, `workflow/active_work/README.md` |
| `context/donor-licensing.md` | L0-project | — | `docs/LICENSING.md`, `THIRD_PARTY_NOTICES.md`, `LICENSE` |
| `context/network-exposure.md` | L0-project | — | `docs/NETWORK_EXPOSURE.md`, `deploy/docker-compose.yml`, `OSJS/Dockerfile`, `OSJS/src/server/config.js` |
| `context/planning-workflow.md` | L0-project | brainstorm-topics | `planning/README.md`, `planning/Brainstorm/README.md`, `planning/microtask/README.md`, `planning/microtask/rules.md` |
| `context/brainstorm-topics.md` | planning-workflow | — | `planning/Brainstorm/mma2-basic-install-test.md`, `planning/Brainstorm/osjs-modbus-simulator.md` |
| `context/osjs-shell.md` | L0-project | — | `OSJS/README.md`, `OSJS/package.json`, `OSJS/Dockerfile`, `OSJS/webpack.config.js`, `OSJS/scripts/build-local-packages.js`, `OSJS/src/server/config.js`, `OSJS/src/server/index.js`, `OSJS/src/server/providers/health.js`, `OSJS/src/server/providers/classic-icons.js`, `OSJS/src/client/config.js`, `OSJS/src/client/index.ejs`, `OSJS/src/packages/NamelessClassicIcons/metadata.json`, `OSJS/src/packages/NamelessWorkstationTheme/metadata.json` |
| `context/active-work.md` | L0-project | simulator-device-config | `handoff.md`, `workflow/active_work/*`, `MMA2/testdata/smoke-test.yaml`, `simulator/device.go`, `simulator/store.go`, `simulator/validate.go`, `simulator/scheduler.go`, `simulator/scheduler_test.go`, `simulator/raw_ingest.go`, `simulator/raw_ingest_test.go` |
| `context/simulator-device-config.md` | active-work | — | `simulator/*`, `deploy/docker-compose.yml`, `OSJS/src/server/config.js`, `MMA2/internal/config/validate.go`, `planning/Brainstorm/osjs-modbus-simulator.md`, `workflow/active_work/sim-001-simulator-device-config.md`, `workflow/active_work/sim-002a-mma2-config-ownership.md`, `workflow/active_work/sim-003-random-runtime.md`, `workflow/active_work/sim-004-raw-ingest.md` |

## Access Rule

Repository-dependent operations use ICC first, but ICC-first does not authorize browsing all ICC context.

```text
READ ICC INDEX
→ IDENTIFY THE SEMANTIC BOUNDARY SELECTED BY THE CURRENT OPERATION / AUTHORIZED TASK
→ LOCATE THAT CONTEXT BRANCH
→ VALIDATE BASELINE / OVERLAY FOR THAT BRANCH
→ CURRENT: USE THAT BRANCH
→ NEED MORE DETAIL: ZOOM IN WITHIN THE SAME BRANCH
→ STALE OR MISSING: BLACK SHEEP WALL REFRESHES ONLY THE AFFECTED CONTEXT
→ ACT
```

The current operation or authorized task determines the context boundary. Possible relevance, dependency, or future usefulness does not authorize movement into another semantic boundary. Source files are consulted when the selected ICC branch is absent, stale, insufficiently detailed, or affected by repository changes.

## ICC Navigation Rule

ICC navigation is branch-local.

### Allowed

- Stay at the context selected by the current operation or authorized task.
- Zoom in to child context required to understand, implement, test, or verify that work.
- Return to a parent inside the same selected semantic branch when needed to preserve local context.
- Refresh stale or missing context only inside the affected branch.

### Not Allowed

- Zoom out above the semantic boundary established by the current operation or authorized task merely to search for possibly relevant information.
- Enter sibling semantic contexts because they are registered in ICC, related by dependency, or may matter later.
- Traverse licensing, networking, planning, architecture, deployment, or any other sibling boundary unless the current authorized work explicitly crosses into that boundary.
- Treat the ICC registry as a checklist of contexts to read.

Cross-boundary access is allowed only when the current operation or authorized task explicitly requires that other semantic boundary to complete its stated outcome or verification.

## Zoom Model

```text
L0 — broadest project/system view
 ↓
L1
 ↓
...
 ↓
LX — as deep as required
```

Parent files may use `## Zoom In`; children may use `## Zoom Out`.

`Zoom In` means descend to more specific context inside the selected semantic branch. `Zoom Out` is only for returning within that same branch. It must not be used to climb above the task-selected boundary and then enter a sibling branch. The selected task boundary is the navigation ceiling for that operation unless the authorized task explicitly crosses another semantic boundary.

## Context File Rules

- One context file = one semantic boundary.
- Context files summarize established state, contracts, dependencies, and unresolved questions.
- Context files do not store chat history or duplicate source files verbatim.
- Every context declares its material repository source dependencies.
- Every context records the baseline commit against which its committed dependencies were audited.
- If it incorporates uncommitted dependencies, it records their audited path fingerprints or refers to the index overlay registry.
- Known-stale context is never consumed as authoritative context.
- Unaffected synchronized context is not recomputed.
- Context links must support deliberate zoom navigation; they must not imply permission to traverse sibling boundaries.
- Context files target 100–150 lines, hard max 200; split deeper detail into child context rather than duplicating source files. When a node contains several independently navigable semantic subjects, prefer child nodes over one broad context file.

## Validity

A context is synchronized when its committed dependencies are represented by the current ICC baseline and every relevant working-tree dependency matches the recorded audited overlay. If current HEAD differs from `Baseline Commit`, compare the baseline to HEAD and invalidate only contexts whose declared dependencies intersect the changed committed files. If current HEAD equals `Baseline Commit`, only working-tree changes that differ from the audited overlay can invalidate context. Validity and relevance are separate. A synchronized context may still be outside the semantic boundary of the current operation and therefore must not be consumed.

## BLACK SHEEP WALL Lifecycle

### Bootstrap

When no valid semantic baseline exists:

```text
current HEAD
→ audit all relevant repository context
→ discover semantic boundaries
→ compact into ICC/context/
→ register contexts here
→ stamp baseline commit
→ record audited uncommitted overlay
```

### Incremental Maintenance

When a baseline exists:

```text
baseline commit
→ current HEAD
→ committed diff if HEAD changed
→ current uncommitted changes
→ compare against audited overlay
→ refresh affected context only
→ preserve unaffected context
```

If HEAD has not changed, BLACK SHEEP WALL must inspect only new, modified, renamed, or deleted uncommitted files that differ from the last audited overlay. BLACK SHEEP WALL may maintain multiple semantic branches, but an invoking operation consumes only the branch selected by that operation. Maintaining ICC breadth does not grant operational access to unrelated branches.
