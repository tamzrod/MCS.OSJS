# Planning

Planning is human-owned.

It contains brainstorm material and decomposed microtasks that are not yet authorized for JR execution.

## Mandatory Context Prelude

Every repository-dependent Planning operation must first read `BLACK_SHEEP_WALL.md` and run BLACK SHEEP WALL within the planning scope.

BLACK SHEEP WALL validates or refreshes the minimum required ICC context, then returns control to Planning. It does not authorize implementation or widen planning scope.

Planning must not consume ICC context by assuming it is current without this prelude.

## Flow

```text
PLANNING OPERATION
   ↓
BLACK SHEEP WALL
   ↓
VALID RELEVANT CONTEXT
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
