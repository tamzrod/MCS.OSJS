# BLACK SHEEP WALL

## Purpose

BLACK SHEEP WALL is the sole owner and maintainer of ICC.

Its job is to keep ICC as a compact semantic cache of repository truth while making repository understanding cheap through semantic zoom.

It must answer two questions efficiently:

1. Where am I in the system, and how much context do I need?
2. What is the smallest ICC change required by the repository delta?

> Zoom is for navigation. Delta is for maintenance. The cheapest valid context wins.

BLACK SHEEP WALL does not grant execution authority, choose work, promote tasks, change workflow state, or widen the operation that invoked it.

## Core Model

ICC is a semantic tree with optional cross-tree connectors.

The tree is primary. Connectors are secondary.

```text
ICC/INDEX.md
    ↓ route
semantic node
    ↓
ZOOM IN / ZOOM OUT / PAN / TRACE
```

Each ICC node is both:

- a readable semantic summary;
- a container that may have narrower child nodes.

Parents summarize children. Children hold detail that would make the parent too broad.

The hierarchy follows repository semantics, not necessarily repository folders.

## Two Operating Modes

### 1. Bootstrap

Use bootstrap only when no valid ICC baseline exists or the current ICC is genuinely unusable.

```text
NO VALID BASELINE
→ inspect repository broadly enough to establish the semantic model
→ build the ICC hierarchy
→ create compact context nodes
→ create only established connectors
→ register routing in ICC/INDEX.md
→ establish baseline and audited overlay state
→ verify
→ stop
```

After bootstrap succeeds, repository-wide rebuilding is disabled for normal operation.

### 2. Incremental Maintenance

When a valid baseline exists, BLACK SHEEP WALL is strictly incremental.

```text
READ ICC STATE
→ READ CURRENT HEAD
→ DETERMINE COMMITTED + WORKING-TREE DELTA
→ INSPECT CHANGED PATHS ONLY
→ MAP CHANGES TO THE SMALLEST AFFECTED SEMANTIC NODE
→ PATCH DEEPEST AFFECTED NODE FIRST
→ PROPAGATE ONLY IF SEMANTIC TRUTH CHANGED
→ UPDATE INDEX ONLY IF ROUTING STRUCTURE CHANGED
→ UPDATE BASELINE / OVERLAY STATE
→ VERIFY AFFECTED CONTEXT ONLY
→ STOP
```

A direct BLACK SHEEP WALL invocation after bootstrap means maintain ICC from repository delta. It does not mean perform another full audit.

## Navigation Operations

Semantic navigation is read-only unless a node is stale or missing.

### Zoom In

Move to a child node when the current node is too broad.

Load only the child required by the current question.

### Zoom Out

Move to the parent node when broader ownership, architecture, or boundary context is required.

Do not use Zoom Out to escape an invoking operation's semantic ceiling and enter unrelated work.

### Pan

Move to a sibling only when the current task explicitly requires another concern at the same semantic level.

Possible relevance is not enough.

### Trace

Follow an explicit connector when a dependency, contract, runtime path, data flow, or ownership crossing is required.

Do not recursively load everything connected to the current node.

## Context Gathering Rule

```text
QUESTION / TASK
→ READ ICC/INDEX.md
→ ROUTE TO THE BROADEST USEFUL RELEVANT NODE
→ LOAD THAT NODE
→ ENOUGH?
    YES → STOP GATHERING
    NO  → ZOOM / PAN / TRACE ONLY FOR MISSING DETAIL
→ OPEN REPOSITORY SOURCE ONLY WHEN ICC DOES NOT ESTABLISH THE REQUIRED FACT
  OR WHEN SOURCE VERIFICATION IS REQUIRED
```

The stopping rule is mandatory:

> Stop gathering as soon as the loaded semantic view is sufficient.

Do not load extra context for completeness or future possibility.

## Delta Rules

### HEAD Changed

If current HEAD differs from the ICC baseline:

```text
DIFF baseline..HEAD
→ obtain changed/new/deleted/renamed paths
→ inspect those paths only
→ map them to semantic nodes
```

Do not reopen unchanged files merely to reconfirm them.

### HEAD Unchanged

If HEAD equals the baseline:

```text
CHECK WORKING-TREE DELTA
→ compare changed paths against audited overlay state
→ inspect only paths whose content/state changed since the last audit
```

If there is no committed or working-tree delta:

```text
NO DELTA
→ ICC remains valid
→ no repository inspection
→ no ICC regeneration
→ STOP
```

## Change-to-Context Mapping

Incremental maintenance starts from changed repository paths, not from a tour of ICC.

```text
changed source
    ↓
map to deepest affected semantic node
    ↓
patch node
    ↓
parent summary changed?
    ├── NO  → stop propagation
    └── YES → patch parent
```

Review descendants or connectors only when the changed fact demonstrably affects them.

Do not invalidate an entire branch or tree by default.

## Index Stability Rule

`ICC/INDEX.md` is a stable routing and state map, not a document that must be rewritten after every context refresh.

Update the index only when one of these changes:

- a semantic node is created or deleted;
- a node is renamed or moved;
- parent/child routing changes;
- an explicit connector needed for navigation changes;
- index-level baseline/state metadata actually changes.

A normal source change that only updates facts inside an existing ICC node must not cause unrelated index rewrites.

```text
ZOOM = cheap reads
DELTA = cheap updates
INDEX = stable routing
```

## Semantic Decomposition

A node should represent one useful semantic boundary.

Split a node when it contains multiple independently understandable concerns or when narrow questions require loading too much unrelated detail.

Semantic breadth is the primary split trigger. Line count is only a guardrail.

Do not eagerly generate every possible leaf.

## Context Inheritance

Children inherit broad meaning from their parent path.

Do not duplicate full parent summaries inside every child.

Repeat only parent facts required to prevent ambiguity at the child boundary.

## Cross-Tree Connectors

Connectors represent relationships that do not belong in the parent-child hierarchy, such as:

- dependency;
- contract;
- runtime/data flow;
- ownership crossing.

They are navigation hints, not automatic imports.

Follow a connector only when the current operation requires that relationship.

## Delegated Branch Refresh

When invoked by CWAL or another bounded operation, the caller supplies the semantic boundary.

That boundary is both the navigation ceiling and refresh ceiling.

```text
RECEIVE SELECTED BRANCH
→ READ ITS ICC CONTEXT
→ CHECK WHETHER REPOSITORY DELTA INTERSECTS ITS SOURCE DEPENDENCIES
→ NO INTERSECTION: RETURN CURRENT CONTEXT
→ INTERSECTION:
     inspect changed dependencies only
     patch deepest affected node
     propagate only within the selected branch when required
→ RETURN
```

Staleness outside the selected branch is irrelevant to the invoking operation.

A delegated refresh must not inventory or repair unrelated ICC branches.

## ICC State Model

ICC represents:

```text
committed repository truth at BASELINE COMMIT
+
audited uncommitted working-tree overlay
```

Record enough deterministic state to determine whether an uncommitted file differs from the version already audited. A content hash or equivalent fingerprint is preferred.

## Context Node Minimums

Each node should contain only what is needed for navigation and reliable work:

- Semantic Boundary;
- what it owns;
- what it does not own where important;
- established facts and decisions;
- unresolved questions that must remain unresolved;
- Parent / Zoom Out target;
- direct Zoom In children;
- relevant connectors;
- Source Dependencies;
- baseline/overlay state needed for maintenance.

Do not duplicate large source content inside ICC.

## Authority

Git remains final repository authority.

ICC is a verified semantic cache, not a replacement for source.

If ICC conflicts with authoritative repository source:

```text
mark affected context stale
→ verify source
→ refresh smallest affected node
→ propagate only actual semantic impact
```

Never silently resolve ambiguity that the repository itself does not resolve.

## Ownership Rule

BLACK SHEEP WALL is the only operation authorized to modify `ICC/`.

Other operations may read ICC and may request a bounded refresh, but they must not create, edit, patch, regenerate, or commit ICC state themselves.

## Boundary Rules

BLACK SHEEP WALL cannot:

- grant execution authority;
- choose or promote work;
- change Active Work state;
- widen an invoking operation;
- perform a full audit when a valid baseline exists;
- inspect unchanged files merely to reconfirm them;
- regenerate unaffected ICC nodes;
- browse unrelated semantic branches;
- raise a delegated navigation ceiling;
- treat `ICC/INDEX.md` as a checklist of context to read.

## Verification

### Bootstrap

Verify semantic hierarchy, routing, source mappings, connectors, and baseline state.

### Incremental

Verify only what the delta could have affected:

1. changed paths map to the correct semantic node;
2. affected context reflects current repository truth;
3. propagation occurred only where summarized truth changed;
4. structural changes are reflected in the index only when required;
5. unrelated branches were untouched;
6. baseline/overlay state represents the audited repository state.

## BLACK SHEEP WALL Command

```text
BLACK SHEEP WALL
→ READ ICC/INDEX.md
→ VALID BASELINE?
    NO  → BOOTSTRAP ONCE
    YES → INCREMENTAL MODE
→ IF CALLED WITH A BRANCH, RESPECT THAT SEMANTIC CEILING
→ ROUTE TO RELEVANT NODE
→ CHECK VALIDITY AGAINST SOURCE DEPENDENCIES + DELTA
→ CURRENT? RETURN MINIMUM VALID CONTEXT
→ STALE/MISSING? INSPECT CHANGED SOURCE ONLY
→ PATCH DEEPEST AFFECTED NODE
→ PROPAGATE ONLY ACTUAL SEMANTIC IMPACT
→ UPDATE INDEX ONLY IF ROUTING/STRUCTURE CHANGED
→ VERIFY AFFECTED CONTEXT
→ RETURN
```

## Design Invariants

1. BLACK SHEEP WALL is the sole ICC writer.
2. Git is final authority.
3. The semantic tree is the primary navigation model.
4. Connectors supplement the tree but never replace it.
5. Zooming is read-only unless stale or missing context requires maintenance.
6. After bootstrap, maintenance is delta-driven and incremental.
7. Unchanged files and unaffected ICC nodes consume no maintenance work.
8. Patch the deepest affected node first.
9. Propagate only demonstrated semantic impact.
10. Keep `ICC/INDEX.md` stable unless routing, structure, connectors, or index-level state changes.
11. Respect the invoking operation's semantic ceiling.
12. Stop gathering as soon as sufficient context exists.
13. ICC must remain cheaper to navigate than reading the repository directly.
