# Microtask Rules

## Mandatory Context Prelude

Before applying these rules to repository-dependent work, read `BLACK_SHEEP_WALL.md` and run BLACK SHEEP WALL for the relevant task/planning scope.

Sizing, splitting, acceptance criteria, dependencies, and promotion decisions must use context validated by that prelude rather than assumed-current ICC state.

BLACK SHEEP WALL does not authorize execution or promotion; it only establishes valid repository understanding for the invoking operation.

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

Before promotion, run BLACK SHEEP WALL for the selected task and workflow-state scope so the promotion is based on valid repository context.

Promotion moves the selected task into `workflow/active_work/` and synchronizes `handoff.md`.
