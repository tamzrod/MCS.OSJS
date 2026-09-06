# BLACK SHEEP WALL

## Purpose

BLACK SHEEP WALL maintains ICC as a compact, token-efficient cache of repository truth.

Its job is not to repeatedly understand the whole repository. Its job is to avoid repeated repository reading by building a baseline once, then patching only the context invalidated by repository changes.

> The cheapest valid context refresh wins.

BLACK SHEEP WALL is not an authorization mechanism. It cannot grant work, promote tasks, choose future work, or widen an invoking operation.

## Core Invariant

> Full repository audit is allowed only when no valid ICC baseline exists.

Once a valid baseline exists, BLACK SHEEP WALL is strictly incremental.

```text
NO BASELINE
→ bootstrap full audit once
→ build ICC hierarchy
→ establish baseline

BASELINE EXISTS
→ never full-audit again
→ inspect repository delta only
→ patch affected ICC context only
```

Unchanged repository files, unchanged ICC nodes, and unrelated semantic branches should consume no repository-reading context during an incremental run.

## Mode 1 — Bootstrap

Bootstrap is used only when ICC has no valid baseline or the ICC state is genuinely unusable as a baseline.

```text
READ CURRENT HEAD
→ AUDIT REPOSITORY CONTEXT REQUIRED TO ESTABLISH THE INITIAL MODEL
→ DISCOVER SEMANTIC BOUNDARIES
→ BUILD ICC FOLDER / ZOOM HIERARCHY
→ COMPACT REPOSITORY TRUTH INTO ICC/context/
→ UPDATE ICC/INDEX.md
→ STAMP BASELINE WITH CURRENT HEAD
→ RECORD AUDITED UNCOMMITTED OVERLAY IF PRESENT
→ VERIFY HIERARCHY + STATE
→ STOP
```

Bootstrap may inspect broadly because there is no previous state from which to calculate a delta.

After bootstrap succeeds, this broad audit behavior is disabled for normal BLACK SHEEP WALL operation.

## Mode 2 — Incremental Maintenance

Every BLACK SHEEP WALL run after bootstrap begins from the existing ICC baseline.

```text
READ ICC STATE METADATA
→ READ CURRENT HEAD
→ DETERMINE DELTA FROM LAST AUDITED ICC STATE
     │
     ├── committed delta: baseline..HEAD
     └── working-tree delta: new / modified / renamed / deleted files
→ INSPECT ONLY CHANGED PATHS
→ MAP EACH CHANGED PATH TO ITS EXISTING ICC SEMANTIC BRANCH
→ REFRESH DEEPEST AFFECTED NODE FIRST
→ PROPAGATE UPWARD ONLY IF PARENT SUMMARY TRUTH CHANGED
→ CREATE / RESTRUCTURE A BRANCH ONLY IF SEMANTIC STRUCTURE CHANGED
→ PRESERVE ALL UNAFFECTED CONTEXT
→ UPDATE BASELINE / OVERLAY STATE
→ VERIFY AFFECTED PATHS ONLY
→ STOP
```

BLACK SHEEP WALL must not perform a repository-wide audit merely because the human invoked it directly.

Direct invocation after bootstrap means: **maintain the ICC from repository delta**.

## Delta Rules

### HEAD Changed

If current HEAD differs from the ICC baseline:

```text
DIFF baseline..HEAD
→ obtain changed/new/deleted/renamed paths
→ inspect those paths only
→ map them to ICC context
```

Do not reopen unchanged files to confirm that they are still unchanged.

### HEAD Unchanged

If current HEAD equals the ICC baseline:

```text
CHECK WORKING-TREE DELTA
→ compare changed paths against audited overlay fingerprints
→ inspect only paths whose content/state differs from the last audited overlay
```

If there is no committed or working-tree delta, BLACK SHEEP WALL should perform no context rebuild.

```text
NO DELTA
→ ICC remains valid
→ no repository context inspection
→ no ICC regeneration
→ STOP
```

## Token-Efficiency Rule

ICC exists to reduce context use. BLACK SHEEP WALL must therefore optimize for minimum repository reading.

```text
UNCHANGED REPOSITORY FILE
→ 0 inspection

UNCHANGED ICC NODE
→ 0 regeneration

UNRELATED SEMANTIC BRANCH
→ 0 inspection

CHANGED FILE
→ inspect only required changed content

DEEPEST AFFECTED ICC NODE
→ patch required summary

PARENT NODE
→ patch only if its zoomed-out truth changed
```

Do not load extra context "for completeness", "to be safe", or because it may become relevant later.

## ICC Semantic Folder Hierarchy

ICC is a semantic zoom tree represented by folders and files.

The filesystem structure itself should guide context navigation.

```text
ICC/
└── context/
    ├── PROJECT.md                 ← broadest project overview
    │
    ├── architecture/
    │   ├── ARCHITECTURE.md        ← zoomed-out architecture overview
    │   ├── mma2/
    │   │   ├── MMA2.md            ← zoomed-out MMA2 overview
    │   │   ├── config/
    │   │   │   ├── CONFIG.md      ← config overview
    │   │   │   └── ...            ← deeper config details
    │   │   └── runtime/
    │   │       ├── RUNTIME.md     ← runtime overview
    │   │       └── ...            ← deeper runtime details
    │   ├── osjs/
    │   │   ├── OSJS.md
    │   │   └── ...
    │   ├── replicator/
    │   │   ├── REPLICATOR.md
    │   │   └── ...
    │   └── orchestrator/
    │       ├── ORCHESTRATOR.md
    │       └── ...
    │
    ├── networking/
    │   ├── NETWORKING.md
    │   └── ...
    │
    ├── workflow/
    │   ├── WORKFLOW.md
    │   └── ...
    │
    └── provenance/
        ├── PROVENANCE.md
        └── ...
```

This is a shape rule, not a fixed required folder list. Actual semantic branches must be derived from repository truth.

### Folder Rule

A semantic folder is a navigation boundary.

The overview file at the root of that folder contains the **Zoom Out view of everything beneath that folder**.

Subfolders contain deeper semantic detail.

```text
FOLDER
= semantic boundary

OVERVIEW FILE AT FOLDER ROOT
= zoomed-out summary of that boundary

SUBFOLDER
= allowed Zoom In boundary

DEEPER FILE
= specific context required only when deeper detail is needed
```

Do not flatten unrelated semantic levels into sibling files when one clearly belongs beneath another.

## Zoom Navigation

A repository-dependent operation begins at the smallest useful semantic boundary for its authorized outcome.

```text
SELECT RELEVANT ICC FOLDER
→ READ ITS OVERVIEW FILE
→ ENOUGH?
     YES → ACT
     NO  → ZOOM IN TO RELEVANT SUBFOLDER
→ READ THAT SUBFOLDER OVERVIEW
→ REPEAT ONLY AS NEEDED
```

### Zoom In

Zoom In means descending into a child folder inside the currently selected semantic branch.

```text
architecture/mma2/MMA2.md
        ↓
need runtime detail
        ↓
architecture/mma2/runtime/RUNTIME.md
        ↓
need raw-ingest detail
        ↓
architecture/mma2/runtime/raw-ingest.md
```

### Zoom Out

Zoom Out means returning to the overview of the current branch's parent.

Zoom Out is for summarization and orientation inside the selected branch. It must not be used to climb above the operation's semantic ceiling and then enter unrelated sibling branches.

### Sideways Navigation

Sibling traversal is prohibited unless the current authorized outcome explicitly requires both boundaries.

Possible relevance is not authority.

## Change-to-Context Mapping

Incremental BLACK SHEEP WALL starts from changed repository paths, not from a tour of ICC.

Example:

```text
changed repository file:
MMA2/runtime/raw_ingest.cpp

        ↓ map

ICC/context/
└── architecture/
    └── mma2/
        └── runtime/
            └── raw-ingest.md
```

Refresh `raw-ingest.md` first.

Then ask whether the changed truth affects the parent summary:

```text
raw-ingest.md changed
        ↓
Does RUNTIME.md summary truth change?
        ├── NO  → leave RUNTIME.md untouched
        └── YES → patch RUNTIME.md
                    ↓
              Does MMA2.md summary truth change?
                    ├── NO  → stop propagation
                    └── YES → patch MMA2.md
```

Do not automatically rewrite ancestors.

## New Files and Structural Changes

When a changed/new path does not map to an existing ICC node:

```text
NEW CHANGED PATH
→ identify its closest existing semantic parent
→ determine whether it belongs in an existing node
→ if necessary create one child node/folder
→ update only the affected parent overview/linkage
```

Do not rediscover the whole ICC tree because one new semantic boundary appeared.

Reorganize an existing branch only when repository changes actually altered that branch's semantic structure.

## Delegated Branch Refresh

When BLACK SHEEP WALL is invoked by CWAL, planning, promotion, execution, or another bounded operation, that operation supplies the semantic boundary.

That boundary is the navigation ceiling and refresh ceiling.

```text
RECEIVE SELECTED BRANCH
→ READ ITS ICC STATE
→ CHECK WHETHER REPOSITORY DELTA INTERSECTS THAT BRANCH'S SOURCES
→ NO INTERSECTION: USE EXISTING ICC AND RETURN
→ INTERSECTION:
     inspect changed dependencies only
     patch deepest affected node
     propagate upward only inside selected branch when needed
→ RETURN
```

Staleness outside the selected branch is irrelevant to the invoking operation.

A delegated refresh must not:

- inventory all stale ICC contexts;
- inspect unrelated repository changes;
- repair sibling branches;
- update broad project context merely because HEAD advanced;
- raise the navigation ceiling;
- turn task execution into repository-maintenance work.

## ICC State Model

ICC represents:

```text
committed repository truth at BASELINE COMMIT
+
audited uncommitted working-tree overlay
```

The baseline commit is the repository HEAD used as the committed source baseline during the last successful audit/refresh.

For uncommitted files, record enough deterministic state to tell whether the working-tree content is identical to the version already audited. A content hash or equivalent fingerprint is preferred.

## Context Node Metadata

Each semantic overview/detail node should retain only metadata needed for deterministic maintenance and navigation:

- Semantic Boundary
- Parent / Zoom Out target
- direct Zoom In children
- Source Dependencies
- Baseline / overlay state relevant to those dependencies

Do not duplicate large source content inside ICC.

## Operation Prelude

Repository-dependent operations are ICC-first and branch-first.

```text
IDENTIFY AUTHORIZED OUTCOME
→ SELECT SMALLEST RELEVANT ICC BRANCH
→ READ THAT BRANCH OVERVIEW
→ CHECK WHETHER ITS SOURCE DEPENDENCIES CHANGED
→ CURRENT? USE ICC
→ STALE? REFRESH CHANGED DEPENDENCIES ONLY
→ NEED MORE DETAIL? ZOOM IN
→ ACT
```

The ICC index is a registry/state map. It is not a checklist of context to read.

## Boundary Rules

BLACK SHEEP WALL cannot:

- grant execution authority;
- promote work;
- choose future work;
- widen a task or planning operation;
- perform a new full audit when a valid baseline exists;
- inspect unchanged files merely to reconfirm them;
- regenerate unaffected ICC nodes;
- browse unrelated semantic branches;
- raise an invoking operation's navigation ceiling;
- treat the ICC registry as a menu to browse.

## Verification

### Bootstrap Verification

Verify the initial hierarchy, source mappings, overview/child relationships, and baseline state.

### Incremental Verification

Verify only what the delta could have affected:

1. each changed repository path was mapped to the correct ICC branch;
2. deepest affected context reflects current repository truth;
3. parent summaries were changed only where their summarized truth changed;
4. new/deleted/renamed semantic nodes are reflected where necessary;
5. unrelated sibling branches were not inspected or regenerated;
6. baseline/overlay metadata now represents the audited state.

Do not perform a full-tree verification after every incremental update unless the change itself altered the tree structure globally.

## Context Size

ICC context should remain compact.

- overview files summarize only their semantic folder;
- deeper detail belongs in subfolders/files;
- avoid duplicating child detail in parent summaries;
- prefer the smallest context necessary for the operation;
- if an overview becomes too large, move detail deeper instead of expanding the overview.

## Mental Model

```text
ICC
= compressed semantic cache

BLACK SHEEP WALL BOOTSTRAP
= build cache once

BLACK SHEEP WALL AFTER BASELINE
= incremental cache invalidation + patch

NOT
= repository re-analysis
```
