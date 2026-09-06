# BLACK SHEEP WALL

## Purpose

BLACK SHEEP WALL is the repository-understanding and context-maintenance operation for MCS.OSJS.

It is the mandatory context prelude for every repository-dependent operation. It is not an authorization mechanism.

```text
INVOKING OPERATION
→ READ BLACK_SHEEP_WALL.md
→ READ ICC INDEX
→ CHECK REPOSITORY CHANGES RELEVANT TO INVOKING SCOPE
→ MAP CHANGED DEPENDENCIES
→ INVALIDATE AFFECTED CONTEXT ONLY
→ AUDIT STALE OR MISSING REQUIRED CONTEXT
→ COMPACT
→ VERIFY AGAINST REPOSITORY AUTHORITY
→ RETURN TO INVOKING OPERATION
```

## Mandatory Invocation Rule

Any repository-dependent operation must invoke BLACK SHEEP WALL before it consumes ICC context or acts on repository state.

This includes, at minimum:

- Brainstorm and Planning;
- Microtask creation, sizing, and refinement;
- Promotion;
- Operation CWAL;
- Active Work execution and verification;
- repository-dependent audits and documentation changes.

An operation must not bypass BLACK SHEEP WALL by reading ICC files directly and assuming they are valid.

When the user invokes `BLACK SHEEP WALL` directly, this file defines the operation itself. When another operation invokes it, BLACK SHEEP WALL runs only as that operation's context prelude and then returns control.

## Boundary Rule

BLACK SHEEP WALL must respect the scope of the invoking operation.

It cannot:

- grant execution authority;
- promote work;
- choose future work;
- let CWAL inspect Planning to select tasks;
- let brainstorming implement code;
- widen a bounded operation.

The invoking operation retains authority and scope. BLACK SHEEP WALL only establishes valid repository understanding for that scope.

## ICC Principle

Context validity follows declared dependencies, not repository HEAD alone.

A context remains valid when no material declared dependency has changed since that context was last audited.

When repository changes occur:

```text
inspect relevant diff
→ map changed files to context domains
→ invalidate only affected context
→ refresh affected context
→ preserve unrelated valid context
```

A known-stale context must never be consumed by an operation.

Missing context is created only when the invoking operation requires it and repository source boundaries exist to support it.

## Context Size

Context files should stay compact and semantic.

- target: 100–150 lines;
- hard maximum: 200 lines;
- split deeper detail into a child context rather than deleting useful information.
