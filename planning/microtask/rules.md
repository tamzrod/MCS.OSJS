# Microtask Rules

## ICC-First Context Rule

Before applying these rules to repository-dependent work, read `ICC/INDEX.md` first and use synchronized relevant ICC context.

If required context is stale or missing, BLACK SHEEP WALL refreshes only the affected context inside the semantic branch selected by the current planning operation. Do not recompute or refresh unrelated repository knowledge.

Sizing, splitting, acceptance criteria, dependencies, and promotion decisions must be based on synchronized context.

BLACK SHEEP WALL does not authorize execution or promotion; it only establishes current repository understanding.

## Core Rule

> One task = one primary outcome.

A microtask should be independently understandable, implementable, verifiable, and completable.

A single narrative outcome does not automatically mean a single JR-sized task. If the work crosses independently failure-prone execution boundaries, split it even when those boundaries contribute to one larger outcome.

## Preferred Size

Score tasks across five dimensions. Count **0, 1, or 2** points in each dimension:

1. **Implementation surface** — number and spread of files/components that must change.
2. **Environment / dependency uncertainty** — donor import, toolchain setup, external package/runtime requirements, or unknown local prerequisites.
3. **Behavioral surface** — number of independently meaningful behaviors being introduced or changed.
4. **Verification surface** — number of distinct proof workflows needed to establish completion.
5. **Decision / recovery surface** — unresolved choices or independently recoverable failure points likely to require investigation.

Total score:

- 0–3: good JR task;
- 4–5: split unless the work is tightly coupled and has one deterministic verification workflow;
- 6–7: split;
- 8–10: must split.

Do not reduce the score merely because all work contributes to one feature. Size is about execution complexity and context load, not feature count.

## Mandatory Split Triggers

Split a task when any of these are true:

1. More than 3 independent acceptance outcomes.
2. More than 3 independent implementation verbs.
3. Multiple distinct verification workflows.
4. Multiple architectural decisions.
5. The task naturally contains sequential sub-tasks that can stand alone and be verified independently.
6. The task combines **donor/import/toolchain establishment** with **runtime behavioral verification**.
7. The task combines environment discovery with product implementation unless the environment is already known and explicitly established in current repository context.
8. A failure in an early phase would force JR to abandon or substantially reinterpret later-phase work.

### Donor / Runtime Rule

Default decomposition for imported components:

```text
IMPORT / ESTABLISH COMPONENT
→ BUILD / STATIC VERIFY

then

RUN COMPONENT
→ BEHAVIORAL / PROTOCOL VERIFY
```

Keep these in one task only when the imported material, toolchain, runtime, and verification path are already proven in the target repository and no independent investigation is expected.

## Context-Budget Check

Before finalizing a task, ask:

```text
Can JR execute this task while staying inside one semantic branch
and one bounded working set of implementation details?
```

If successful execution predictably requires JR to load donor structure, toolchain setup, runtime configuration, protocol behavior, deployment constraints, and restart/recovery semantics at the same time, the task is oversized even if its numeric score appears acceptable.

When uncertain, split at the strongest independently verifiable boundary.

## Task Shape

Each microtask should contain:

- ID and title;
- primary outcome;
- scope;
- explicit non-scope when useful;
- acceptance criteria;
- verification method;
- dependencies;
- sizing assessment with the five-dimension score or a concise justification.

## Promotion Boundary

A microtask remains planning material until a human promotes it.

Before promotion, use synchronized ICC context for the selected task and workflow state. Refresh only stale affected context inside the selected semantic branch through BLACK SHEEP WALL.

Promotion moves the selected task into `workflow/active_work/` and synchronizes `handoff.md`.
