# Planning

Planning is human-owned. It contains brainstorm material and decomposed microtasks that are not yet authorized for JR execution.

## ICC-First Context Rule

Every repository-dependent Planning operation follows `workflow/ICC_CONSUMER.md`. Resolve the requested planning question and its semantic ceiling, then read `ICC/INDEX.md`, only the relevant node metadata from `ICC/manifest.json`, and the smallest useful context. `ICC/FORMAT.md` defines validity. A historical index baseline, remote snapshot or migrated `unverified` node is not proof of local currency. Compare only that node's source dependencies with the actual checkout and audited overlay when available.

REUSE verified-current context with no relevant delta; do not reopen unchanged source. For a required stale, unverified or missing node, request only the bounded BLACK SHEEP WALL UPDATE/REVEAL authorized by the planning operation; it returns to Planning without choosing work or widening the request. If refresh is unavailable, use only permitted authoritative source and explicitly flag unverified ICC, or report the precise blocker. Do not edit ICC yourself or audit unrelated migrated nodes.

## Flow

```text
PLANNING QUESTION → SELECT BOUNDARY → INDEX + RELEVANT MANIFEST NODE
→ CURRENT? REUSE SMALLEST CONTEXT
→ STALE / UNVERIFIED / MISSING? BOUNDED BLACK SHEEP WALL IF AUTHORIZED;
  OTHERWISE TASK-PERMITTED SOURCE OR PRECISE BLOCKER
→ BRAINSTORM → MICROTASK → HUMAN PROMOTION → workflow/active_work/
```

Planning does not grant execution authority. Architectural questions may remain unresolved here and must not silently become implementation decisions.
