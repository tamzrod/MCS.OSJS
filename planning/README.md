# Planning

Planning is human-owned.

It contains brainstorm material and decomposed microtasks that are not yet authorized for JR execution.

## ICC-First Context Rule

Every repository-dependent Planning operation must know `BLACK_SHEEP_WALL.md` and read `ICC/INDEX.md` first.

If the relevant ICC context is synchronized with the current baseline commit and audited working-tree overlay, use ICC directly. If it is stale or missing, BLACK SHEEP WALL refreshes only the affected context before Planning continues.

Planning must not reopen unchanged repository sources when synchronized ICC already contains the required context.

## Flow

```text
PLANNING OPERATION
   ↓
ICC FIRST
   ↓
CURRENT? USE CONTEXT
STALE? BLACK SHEEP WALL REFRESHES AFFECTED ONLY
   ↓
BRAINSTORM
   ↓
MICROTASK
   ↓
PROMOTION
   ↓
workflow/active_work/
```

Planning does not grant execution authority.

Architectural questions may remain unresolved here. They must not be silently treated as decisions during execution.
