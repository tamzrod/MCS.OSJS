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

Scoped simulator projection context: [simulator-projection](context/simulator-projection.md). Handles MMA2 listener matching with device override for state_sealing, RBE, extension hydration.

Scoped Replicator advanced-settings context: [electron-replicator-advanced](context/electron-replicator-advanced.md). Includes shared IP/CIDR input behavior and destination persistence.

Scoped installer settings-owner context: [electron-settings-owner](context/electron-settings-owner.md). Tracks account-specific settings ACLs independently of runtime and workflow state.

Scoped Electron Replicator LED context: [electron-replicator-leds](context/electron-replicator-leds.md). This node tracks the runtime/renderer telemetry contract independently of workflow state.

Scoped Electron Memory layout context: [electron-memory-layout](context/electron-memory-layout.md). This node tracks its own baseline and overlay; it does not refresh the unrelated workflow state below.

Baseline Commit: 71d729c34a9a1dfa8cf62f5ff367a059c7e9c8f1
Working Tree: clean
Audited Uncommitted Overlay: none at audit start. This bounded refresh changes only this index and `context/active-work.md`.

Selected-branch state: `active-work` is refreshed at `71d729c` for the recovered OTR queue. OTR-002A is the sole ACTIVE task; OTR-002B/003A/003B are QUEUED; OTR-004 through OTR-014 are NOT EXECUTABLE parents. The current handoff and CWAL adapter require one handoff-named task, bounded writes, a non-force push to `origin/opencode`, remote verification and STOP. Existing Toolkit source predates the OTR implementation roadmap and already contains live Memory, Replicator and Diagnostics implementation, so the implementation parents require reconciliation and redesign before promotion. The failed `db48c41` run's invalid OTR-002B/003A archives remain explicitly excluded, while its added `simulator/devices.yaml` and typo `operation c wal.md` remain unresolved branch contamination. Other semantic branches retain their own recorded baselines and were not refreshed.

Refreshed at `71d729c` over the `c289f2f..71d729c` delta: active-work plus this index. Unchanged and not consumed as current authority: governance, donor-licensing, network-exposure, planning-workflow, brainstorm-topics, osjs-shell, replicator, simulator-device-config and simulator-memory-none. Product files were inspected only to determine whether the OTR roadmap matches the implemented Toolkit; those semantic branches were not refreshed.

ICC prerequisite status for the ACTIVE task: current for OTR workflow routing at `71d729c`. This refresh does not execute or certify OTR-002A.

| `context/L0-project.md` | none | governance, donor-licensing, network-exposure, planning-workflow, osjs-shell, active-work | `README.md`, `PROJECT_IDENTITY.md`, `handoff.md` |
| `context/governance.md` | L0-project | — | `AGENTS.md`, `BLACK_SHEEP_WALL.md`, `ICC/INDEX.md`, `operation cwal.md`, `handoff.md`, `workflow/active_work/README.md`, `the gathering.md`, `there is no cow level.md` |
| `context/donor-licensing.md` | L0-project | — | `docs/LICENSING.md`, `THIRD_PARTY_NOTICES.md`, `LICENSE` |
| `context/network-exposure.md` | L0-project | — | `docs/NETWORK_EXPOSURE.md`, `deploy/docker-compose.yml`, `OSJS/Dockerfile`, `OSJS/src/server/config.js` |
| `context/planning-workflow.md` | L0-project | brainstorm-topics | `planning/README.md`, `planning/Brainstorm/README.md`, `planning/microtask/README.md`, `planning/microtask/rules.md`, `planning/microtask/umig-*.md`, `workflow/active_work/README.md` |
| `context/brainstorm-topics.md` | planning-workflow | — | `planning/Brainstorm/README.md` |
| `context/osjs-shell.md` | L0-project | — | `OSJS/README.md`, `OSJS/package.json`, `OSJS/Dockerfile`, `OSJS/webpack.config.js`, `OSJS/scripts/build-local-packages.js`, `OSJS/src/server/config.js`, `OSJS/src/server/index.js`, `OSJS/src/server/providers/health.js`, `OSJS/src/server/providers/classic-icons.js`, `OSJS/src/client/config.js`, `OSJS/src/client/index.js`, `OSJS/src/client/index.ejs`, `OSJS/src/client/providers/nameless-app-shortcuts.js`, `OSJS/src/packages/MCSModbusToolkit/*`, `OSJS/src/packages/NamelessClassicIcons/metadata.json`, `OSJS/src/packages/NamelessWorkstationTheme/metadata.json` |
| `context/active-work.md` | L0-project | replicator, simulator-device-config | `handoff.md`, `workflow/active_work/*.md`, `workflow/archive/electron-001-nsis-nssm-service-installer.md`, `workflow/archive/electron-002-compact-industrial-layout.md`, `workflow/archive/umig-002-scaffold-single-osjs-toolkit.md`, `workflow/archive/umig-002-r-toolkit-discovery-manifest.md` |
| `context/simulator-device-config.md` | active-work | simulator-memory-none | `docs/SIMULATOR_RUNTIME_INTEGRATION.md`, `simulator/*`, `deploy/docker-compose.yml`, `OSJS/src/server/*`, `OSJS/src/packages/ModbusSimulator/*`, `MMA2/pkg/configvalidate/validate.go`, `MMA2/internal/restartwatch/*`, `MMA2/cmd/mma2-supervisor/*`, `workflow/archive/rep-001-share-mma2-reservation-composer.md`, `workflow/archive/sim-001-*.md` through `workflow/archive/sim-024-*.md` |
| `context/simulator-memory-none.md` | simulator-device-config | — | `simulator/validate.go`, `simulator/apply.go`, `simulator/store.go`, `simulator/memory_none_test.go`, `simulator/runtime_server.go`, `simulator/runtime_server_test.go`, `simulator/cmd/modbus-simulator-runtime/main.go`, `simulator/cmd/modbus-simulator-runtime/main_test.go`, `simulator/go.mod`, `simulator/go.sum`, `simulator/Dockerfile`, `deploy/docker-compose.yml`, `OSJS/src/packages/ModbusSimulator/server.js` |
| `context/replicator.md` | active-work | — | `handoff.md`, `workflow/active_work/rep-block-002-independent-block-pollers.md`, `workflow/active_work/rep-block-003-tabbed-block-editor.md`, `workflow/archive/rep-*.md`, `replicator/runtime_api.go`, `replicator/cmd/modbus-replicator-runtime/main.go`, `replicator/go.mod` |

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
