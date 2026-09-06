# ICC Index

## Purpose

This directory is the Incremental Context Compaction (ICC) cache for MCS.OSJS.

ICC stores compact semantic context derived from repository truth so operations can read context first instead of repeatedly reopening unchanged source files.

ICC is also a context-zoom mechanism. An operation selects the semantic boundary required by the current authorized work, then stays inside that branch and zooms deeper only when execution requires more detail.

## Repository Baseline

Every completed BLACK SHEEP WALL audit must record:

- `Baseline Commit`: repository HEAD used as the committed source baseline;
- `Working Tree`: `clean` or `dirty`;
- `Audited Uncommitted Overlay`: changed paths incorporated into ICC, with a deterministic content hash or equivalent fingerprint;
- context registry entries and their source dependencies.

The baseline commit identifies the committed repository state represented by ICC. It does not need to equal the later commit that stores ICC files.

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

The current operation or authorized task determines the context boundary. Possible relevance, dependency, or future usefulness does not authorize movement into another semantic boundary.

Source files are consulted when the selected ICC branch is absent, stale, insufficiently detailed, or affected by repository changes.

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

Example:

```text
ACTIVE TASK: implement OS.js base desktop

SELECT
→ OS.js implementation context

ALLOWED
→ zoom deeper into OS.js runtime/build/package details

NOT ALLOWED
→ move sideways into donor licensing
→ move sideways into network exposure
→ move sideways into planning

UNLESS
→ the active task explicitly requires one of those boundaries
```

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

If HEAD has not changed, BLACK SHEEP WALL must inspect only new, modified, renamed, or deleted uncommitted files that differ from the last audited overlay.

BLACK SHEEP WALL may maintain multiple semantic branches, but an invoking operation consumes only the branch selected by that operation. Maintaining ICC breadth does not grant operational access to unrelated branches.

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

## Validity

A context is synchronized when its committed dependencies are represented by the current ICC baseline and every relevant working-tree dependency matches the recorded audited overlay.

If current HEAD differs from `Baseline Commit`, compare the baseline to HEAD and invalidate only contexts whose declared dependencies intersect the changed committed files.

If current HEAD equals `Baseline Commit`, only working-tree changes that differ from the audited overlay can invalidate context.

Validity and relevance are separate. A synchronized context may still be outside the semantic boundary of the current operation and therefore must not be consumed.

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

`Zoom In` means descend to more specific context inside the selected semantic branch.

`Zoom Out` is only for returning within that same branch. It must not be used to climb above the task-selected boundary and then enter a sibling branch.

The selected task boundary is the navigation ceiling for that operation unless the authorized task explicitly crosses another semantic boundary.

## Registry

Baseline Commit: b3781e10812299a7a701281e65c064c4da55b68f
Working Tree: clean
Audited Uncommitted Overlay: none

| Context | Source Dependencies |
| --- | --- |
| `context/L0-project.md` | `README.md`, `PROJECT_IDENTITY.md`, `handoff.md` |
| `context/governance.md` | `AGENTS.md`, `BLACK_SHEEP_WALL.md`, `ICC/INDEX.md`, `operation cwal.md`, `handoff.md`, `workflow/active_work/README.md` |
| `context/donor-licensing.md` | `docs/LICENSING.md`, `THIRD_PARTY_NOTICES.md`, `LICENSE` |
| `context/network-exposure.md` | `docs/NETWORK_EXPOSURE.md` |
| `context/planning-workflow.md` | `planning/README.md`, `planning/Brainstorm/README.md`, `planning/microtask/README.md`, `planning/microtask/rules.md` |
| `context/brainstorm-topics.md` | `planning/Brainstorm/port-config-and-socket-deployment.md`, `planning/Brainstorm/osjs-base-webapp-init.md` |

Baseline commit: `b3781e10812299a7a701281e65c064c4da55b68f` - HEAD at audit time; working tree clean, no uncommitted overlay recorded. Context files under `ICC/context/` summarize established repository truth; commit `b3781e1` need not equal the commit that stores them per the ICC state model.
