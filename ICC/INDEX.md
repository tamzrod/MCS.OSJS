# ICC Index

## Purpose

This directory is the Incremental Context Compaction (ICC) cache for MCS.OSJS.

ICC stores compact semantic context derived from repository truth so operations can read context first instead of repeatedly reopening unchanged source files.

## Repository Baseline

Every completed BLACK SHEEP WALL audit must record:

- `Baseline Commit`: repository HEAD used as the committed source baseline;
- `Working Tree`: `clean` or `dirty`;
- `Audited Uncommitted Overlay`: changed paths incorporated into ICC, with a deterministic content hash or equivalent fingerprint;
- context registry entries and their source dependencies.

The baseline commit identifies the committed repository state represented by ICC. It does not need to equal the later commit that stores ICC files.

## Access Rule

Repository-dependent operations use ICC first.

```text
READ ICC INDEX
→ LOCATE RELEVANT CONTEXT
→ VALIDATE BASELINE / OVERLAY
→ CURRENT: USE ICC
→ STALE OR MISSING: BLACK SHEEP WALL REFRESHES AFFECTED CONTEXT ONLY
```

Source files are consulted when relevant ICC context is absent, stale, insufficiently detailed, or affected by repository changes.

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

## Context File Rules

- One context file = one semantic boundary.
- Context files summarize established state, contracts, dependencies, and unresolved questions.
- Context files do not store chat history or duplicate source files verbatim.
- Every context declares its material repository source dependencies.
- Every context records the baseline commit against which its committed dependencies were audited.
- If it incorporates uncommitted dependencies, it records their audited path fingerprints or refers to the index overlay registry.
- Known-stale context is never consumed as authoritative context.
- Unaffected synchronized context is not recomputed.

## Validity

A context is synchronized when its committed dependencies are represented by the current ICC baseline and every relevant working-tree dependency matches the recorded audited overlay.

If current HEAD differs from `Baseline Commit`, compare the baseline to HEAD and invalidate only contexts whose declared dependencies intersect the changed committed files.

If current HEAD equals `Baseline Commit`, only working-tree changes that differ from the audited overlay can invalidate context.

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

## Registry

No semantic context baseline has been generated yet.

The next direct `BLACK SHEEP WALL` invocation must bootstrap `ICC/context/` from existing repository truth and replace this placeholder with the resulting baseline metadata and context registry.
