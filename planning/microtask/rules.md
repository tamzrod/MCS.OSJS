# Microtask Rules

## ICC-First Context Rule

Before applying these rules to repository-dependent work, read `ICC/INDEX.md` first and use synchronized relevant ICC context.

If required context is stale or missing, BLACK SHEEP WALL refreshes only the affected context. Do not recompute unchanged repository knowledge.

Sizing, splitting, acceptance criteria, dependencies, and promotion decisions must be based on synchronized context.

BLACK SHEEP WALL does not authorize execution or promotion; it only establishes current repository understanding.

## Core Rule

> One task = one primary outcome.

A microtask should be independently understandable, implementable, verifiable, and completable.

## Preferred Size

Score tasks across implementation surface, behavior changes, verification, dependencies, and decision points.

- 0–3: good JR task
- 4–5: review and split if possible
- 6–7: split
- 8–10: must split

## Hard Split Rules

Split a task when any of these are true:

1. More than 3 independent acceptance outcomes.
2. More than 3 independent implementation verbs.
3. Multiple distinct verification workflows.
4. Multiple architectural decisions.
5. The task naturally contains sequential sub-tasks that can stand alone.

## Task Shape

Each microtask should contain:

- ID and title;
- primary outcome;
- scope;
- explicit non-scope when useful;
- acceptance criteria;
- verification method;
- dependencies;
- sizing assessment.

## Promotion Boundary

A microtask remains planning material until a human promotes it.

Before promotion, use synchronized ICC context for the selected task and workflow state. Refresh only stale affected context through BLACK SHEEP WALL.

Promotion moves the selected task into `workflow/active_work/` and synchronizes `handoff.md`.
