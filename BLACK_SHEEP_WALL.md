# BLACK SHEEP WALL

## Purpose

BLACK SHEEP WALL is the repository-understanding and context-compaction operation for MCS.OSJS.

It maintains ICC as a compact cache of repository truth. It is not an authorization mechanism.

## Direct Invocation

When the user says `BLACK SHEEP WALL`, perform repository context compaction.

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

### Subsequent Run

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

If current HEAD is unchanged from the ICC baseline, a repeated BLACK SHEEP WALL run updates only new, modified, renamed, or deleted uncommitted files whose current state differs from the last audited overlay.

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

Every repository-dependent operation must know and use BLACK SHEEP WALL, but source access is ICC-first:

```text
READ ICC/INDEX.md
→ LOCATE RELEVANT CONTEXT
→ CHECK BASELINE + WORKING-TREE OVERLAY
→ CURRENT? USE ICC
→ STALE/MISSING? REFRESH ONLY AFFECTED CONTEXT
→ RETURN TO INVOKING OPERATION
```

An operation must not reopen source files merely because they exist. It should use synchronized ICC context first. Repository source is consulted when ICC is missing, stale, insufficient for the requested detail, or must be verified against changed content.

## Validity Rules

Relevant ICC context is current only when:

1. its baseline commit agrees with the ICC repository baseline for the committed state being represented; and
2. no source dependency has changed since that baseline without being incorporated; and
3. every uncommitted source dependency it claims to represent matches the recorded audited overlay state.

If HEAD differs from the recorded baseline, inspect the committed diff from baseline to HEAD and refresh only contexts affected by those changed files. Do not perform a full repository audit unless the baseline is missing, unusable, or repository history prevents a reliable incremental comparison.

If HEAD is unchanged, inspect only uncommitted changes since the last audit.

## Boundary Rule

BLACK SHEEP WALL cannot:

- grant execution authority;
- promote work;
- choose future work;
- let CWAL inspect Planning to select tasks;
- let brainstorming implement code;
- widen a bounded operation.

The invoking operation retains authority and scope. BLACK SHEEP WALL only establishes synchronized repository understanding.

## Context Size

Context files should stay compact and semantic.

- target: 100–150 lines;
- hard maximum: 200 lines;
- split deeper detail into a child context rather than duplicating source files.
