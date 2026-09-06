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
→ COMPACT INTO ICC/context/
→ UPDATE ICC/INDEX.md
→ STAMP BASELINE WITH CURRENT HEAD COMMIT
→ RECORD ANY AUDITED UNCOMMITTED FILE OVERLAYS
→ VERIFY
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
→ REFRESH ONLY AFFECTED ICC CONTEXT
→ PRESERVE UNAFFECTED CONTEXT
→ UPDATE BASELINE / OVERLAY METADATA
→ VERIFY
```

If current HEAD is unchanged from the ICC baseline, a repeated direct BLACK SHEEP WALL run updates only new, modified, renamed, or deleted uncommitted files whose current state differs from the last audited overlay.

## Delegated Branch Refresh

When BLACK SHEEP WALL is invoked by CWAL, planning, brainstorming, promotion, or another bounded operation, the invoking operation must already have selected a semantic branch.

That selected branch is the navigation ceiling and refresh boundary.

```text
RECEIVE SELECTED SEMANTIC BRANCH
→ READ ICC STATE FOR THAT BRANCH
→ CHECK ONLY THAT BRANCH'S DECLARED SOURCE DEPENDENCIES
→ IF HEAD CHANGED: DIFF ONLY TO DETERMINE WHETHER THOSE DEPENDENCIES CHANGED
→ INSPECT ONLY CHANGED DEPENDENCIES INSIDE THAT BRANCH
→ REFRESH ONLY THAT BRANCH / REQUIRED CHILD CONTEXT
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
- turn task execution into repository-maintenance work.

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

Repository-dependent operations are ICC-first, but **branch selection precedes refresh**.

```text
READ ICC/INDEX.md
→ LOCATE OPERATION AUTHORITY / REQUESTED OUTCOME
→ SELECT SEMANTIC BRANCH
→ LOCATE CONTEXT INSIDE THAT BRANCH
→ CHECK BASELINE + WORKING-TREE OVERLAY FOR THAT BRANCH
→ CURRENT? USE ICC
→ STALE/MISSING? DELEGATED REFRESH OF THAT BRANCH ONLY
→ RETURN TO INVOKING OPERATION
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
- convert unrelated ICC staleness into work for the invoking operation.

The invoking operation retains authority and scope. BLACK SHEEP WALL only establishes synchronized repository understanding inside that scope.

## Context Size

Context files should stay compact and semantic.

- target: 100–150 lines;
- hard maximum: 200 lines;
- split deeper detail into a child context rather than duplicating source files.
