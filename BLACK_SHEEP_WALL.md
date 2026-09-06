# BLACK SHEEP WALL

## Purpose

BLACK SHEEP WALL is the repository-understanding and context-maintenance operation for MCS.OSJS.

It is a context prelude, not an authorization mechanism.

```text
BLACK SHEEP WALL
→ READ ICC INDEX
→ CHECK REPOSITORY CHANGES
→ MAP CHANGED DEPENDENCIES
→ INVALIDATE AFFECTED CONTEXT ONLY
→ AUDIT STALE OR MISSING CONTEXT
→ COMPACT
→ VERIFY AGAINST REPOSITORY AUTHORITY
```

## Boundary Rule

BLACK SHEEP WALL must respect the scope of the invoking operation.

It cannot:

- grant execution authority;
- promote work;
- choose future work;
- let CWAL inspect Planning to select tasks;
- let brainstorming implement code;
- widen a bounded operation.

## ICC Principle

Context validity follows declared dependencies, not repository HEAD alone.

A context remains valid when no material declared dependency has changed since that context was last audited.

When repository changes occur:

```text
inspect diff
→ map changed files to context domains
→ invalidate only affected context
→ refresh affected context
→ preserve unrelated valid context
```

A known-stale context must never be consumed by an operation.

## Context Size

Context files should stay compact and semantic.

- target: 100–150 lines;
- hard maximum: 200 lines;
- split deeper detail into a child context rather than deleting useful information.
