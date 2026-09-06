# BLACK SHEEP WALL

## Purpose

BLACK SHEEP WALL is the repository-understanding and context-compaction operation for MCS.OSJS.

It maintains ICC as a compact cache of repository truth. It is not an authorization mechanism.

BLACK SHEEP WALL has two operating modes:

1. **Direct maintenance mode** — explicitly invoked by the human to compact or refresh repository context.
2. **Delegated branch refresh** — invoked by another operation such as CWAL because that operation's selected ICC branch is stale or missing.

These modes must not be confused.

## Direct Invocation

When the user directly says `BLACK SHEEP WALL`, perform repository context compaction.

### First Run / Missing Baseline

```text
READ CURRENT HEAD
→ AUDIT ALL RELEVANT REPOSITORY CONTEXT
→ DISCOVER SEMANTIC BOUNDARIES
→ BUILD A SEMANTIC ZOOM TREE
→ COMPACT INTO ICC/context/
→ UPDATE ICC/INDEX.md
→ STAMP BASELINE WITH CURRENT HEAD COMMIT
→ RECORD ANY AUDITED UNCOMMITTED FILE OVERLAYS
→ VERIFY TREE + CONTEXT
```

A missing ICC baseline is not grounds to do nothing. It requires bootstrap compaction of the repository context that already exists.

### Subsequent Direct Run

Do not rescan unchanged repository files.

```text
READ ICC BASELINE
→ READ CURRENT HEAD
→ IF HEAD CHANGED: DIFF BASELINE..HEAD
→ INSPECT ONLY COMMITTED FILES CHANGED SINCE BASELINE
→ INSPECT CURRENT UNCOMMITTED FILES
→ COMPARE UNCOMMITTED CONTENT WITH LAST AUDITED OVERLAY
→ REFRESH ONLY AFFECTED ICC NODES / BRANCHES
→ RESTRUCTURE ONLY WHERE SEMANTIC BOUNDARIES ACTUALLY CHANGED
→ PRESERVE UNAFFECTED NODES / BRANCHES
→ UPDATE BASELINE / OVERLAY METADATA
→ VERIFY TREE + CONTEXT
```

If current HEAD is unchanged from the ICC baseline, a repeated direct BLACK SHEEP WALL run updates only new, modified, renamed, or deleted uncommitted files whose current state differs from the last audited overlay.

A direct run is repository-wide in scope but incremental in inspection. Repository-wide does not mean reread every repository file or regenerate every context node.

## Semantic Zoom Tree

ICC must be generated as a traversable semantic hierarchy, not a flat bag of summaries.

The broadest project node is the root. More specific semantic boundaries descend beneath it. Detail that belongs to one boundary becomes a child of that boundary rather than a sibling merely because it has its own file.

Example shape only; actual nodes must be discovered from repository truth:

```text
L0 project
├── governance
├── architecture
│   ├── component A
│   │   ├── build/import
│   │   └── runtime
│   └── component B
├── networking
├── planning
│   ├── brainstorm
│   └── microtask
└── provenance
```

Do not manufacture empty branches merely to match this example. Generate only semantic nodes supported by repository content.

### Context Node Contract

Every generated context node must declare enough metadata for deterministic navigation:

- **Semantic Boundary** — what knowledge this node owns.
- **Parent** — the one node reached by Zoom Out; root declares `none`.
- **Zoom In** — only the direct child nodes that may be entered from this node.
- **Source Dependencies** — material repository paths represented by this node.
- **Baseline State** — commit/overlay state against which those dependencies were audited.

A node may also state unresolved questions or established contracts when they belong to that semantic boundary.

The ICC index is a registry and entry map. Registry presence is not navigation permission.

### Zoom In

Zoom In means moving from the current context node to one of its declared direct children because the current authorized operation needs more specific information inside the same semantic branch.

```text
CURRENT NODE
→ choose required declared child
→ ZOOM IN
→ continue inside that child boundary
```

An operation must not skip sideways to a sibling merely because the sibling looks relevant.

### Zoom Out

Zoom Out means returning only to the current node's declared parent.

```text
CHILD
→ ZOOM OUT
→ DECLARED PARENT
```

Zoom Out is not permission to escape the operation's navigation ceiling. If the declared parent lies above the ceiling selected by the invoking operation, stop at the ceiling.

Zoom Out must never be used as a route to climb to a common ancestor and then enter an unrelated sibling branch.

### Cross-Branch Rule

Crossing into another semantic branch requires the current authorized task/outcome to explicitly require that boundary for implementation, understanding, testing, or verification.

Possible relevance, dependency in another context, future usefulness, or registry proximity is insufficient authority.

When a task explicitly crosses boundaries, each required branch is selected deliberately; this does not open other siblings.

## Delegated Branch Refresh

When BLACK SHEEP WALL is invoked by CWAL, planning, brainstorming, promotion, or another bounded operation, the invoking operation must already have selected a semantic branch/node.

That selected node is the navigation ceiling and refresh boundary.

```text
RECEIVE SELECTED SEMANTIC NODE
→ READ ICC STATE FOR THAT NODE
→ FOLLOW DECLARED CHILD LINKS ONLY WHEN MORE DETAIL IS REQUIRED
→ CHECK ONLY SOURCE DEPENDENCIES IN THE SELECTED NODE / REQUIRED DESCENDANTS
→ IF HEAD CHANGED: DIFF ONLY TO DETERMINE WHETHER THOSE DEPENDENCIES CHANGED
→ INSPECT ONLY CHANGED DEPENDENCIES INSIDE THAT BRANCH
→ REFRESH ONLY THAT NODE / REQUIRED CHILD CONTEXT
→ DO NOT DISCOVER OR REPAIR STALE SIBLING CONTEXT
→ RETURN TO INVOKING OPERATION
```

### Branch-Local Invariant

> Staleness outside the selected semantic branch is irrelevant to the invoking operation.

A delegated refresh must not:

- inventory all stale ICC contexts;
- repair unrelated stale context;
- open sibling contexts because changed files may affect them;
- update broad project context merely because HEAD advanced;
- turn task execution into repository-maintenance work;
- navigate directly from the ICC registry into arbitrary nodes;
- Zoom Out above the selected navigation ceiling.

If a changed file belongs to several ICC contexts, refresh only the context that lies inside the selected branch. Other affected contexts remain stale until a direct BLACK SHEEP WALL run or an operation that selects those branches needs them.

## ICC State Model

ICC represents:

```text
committed repository truth at BASELINE COMMIT
+
audited uncommitted working-tree overlay
```

The baseline commit is the repository HEAD used as the committed source baseline during the audit. It is not required to equal the commit that later stores ICC files; this avoids a self-referential commit-marker loop.

For uncommitted files, ICC must record enough state to determine whether the current working-tree content is the same content that was audited. Prefer a content hash or equivalent deterministic fingerprint per audited path.

## Operation Prelude

Repository-dependent operations are ICC-first, but **branch selection precedes refresh and navigation proceeds inward**.

```text
READ ICC/INDEX.md
→ LOCATE OPERATION AUTHORITY / REQUESTED OUTCOME
→ SELECT SEMANTIC NODE / BRANCH
→ SET THAT NODE AS NAVIGATION CEILING
→ CHECK BASELINE + WORKING-TREE OVERLAY FOR THAT NODE
→ CURRENT? USE NODE
→ NEED MORE DETAIL? ZOOM IN THROUGH DECLARED CHILD
→ STALE/MISSING? DELEGATED REFRESH INSIDE THAT BRANCH ONLY
→ ACT
```

An operation must not reopen source files merely because they exist. It should use synchronized ICC context first. Repository source is consulted when selected ICC context is missing, stale, insufficient for the requested detail, or must be verified against changed content.

## Validity Rules

Selected ICC context is current only when:

1. its committed source dependencies are represented by the relevant ICC baseline state;
2. no source dependency inside the selected branch has changed without being incorporated; and
3. every uncommitted source dependency it claims to represent matches the recorded audited overlay state.

When HEAD differs from the recorded baseline, do not automatically refresh every context touched by the repository diff. First intersect changed paths with the source dependencies of the **selected branch**. Refresh only when that intersection is non-empty.

If HEAD is unchanged, inspect only uncommitted dependencies inside the selected branch whose current state differs from the recorded overlay.

## Boundary Rule

BLACK SHEEP WALL cannot:

- grant execution authority;
- promote work;
- choose future work;
- let CWAL inspect Planning to select tasks;
- let brainstorming implement code;
- widen a bounded operation;
- raise the selected branch's navigation ceiling;
- convert unrelated ICC staleness into work for the invoking operation;
- treat the ICC index as a flat menu of contexts to browse.

The invoking operation retains authority and scope. BLACK SHEEP WALL only establishes synchronized repository understanding inside that scope.

## Verification

A BLACK SHEEP WALL run is not complete until it verifies:

1. every generated node declares its semantic boundary, parent, direct children, source dependencies, and baseline state;
2. every declared child points to an existing context node;
3. every non-root parent points to an existing context node;
4. parent/child relationships agree in both directions;
5. no node is reachable only by arbitrary registry lookup when it should belong beneath a semantic parent;
6. affected source changes are represented in the appropriate nodes;
7. unaffected nodes were not needlessly regenerated during an incremental run;
8. the ICC index reflects the resulting tree and baseline/overlay state.

## Context Size

Context files should stay compact and semantic.

- target: 100–150 lines;
- hard maximum: 200 lines;
- split deeper detail into a child context rather than duplicating source files;
- when a node contains several independently navigable semantic subjects, prefer child nodes over one broad context file.
