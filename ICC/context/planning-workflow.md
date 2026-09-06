# Planning Workflow

Baseline commit: cb869da3caeed040478e67ecb0a01c5f93b3a66b
Working tree: clean
Source dependencies: planning/README.md, planning/Brainstorm/README.md, planning/microtask/README.md, planning/microtask/rules.md

## Model

Planning is human-owned; contains brainstorm material and decomposed microtasks NOT authorized for JR execution. Flow: PLANNING OPERATION -> ICC FIRST -> CURRENT? USE CONTEXT / STALE? BSW REFRESH AFFECTED ONLY -> BRAINSTORM -> MICROTASK -> PROMOTION -> workflow/active_work/. Planning does not grant execution authority; architectural questions may remain unresolved and must not be silently treated as decisions during execution.



## Brainstorm

Human-owned exploration/framing/alternatives/questions/candidates. Not execution authority. ICC-first before repository-dependent brainstorming. BSW may refresh stale/missing affected context only;cannot authorize implementation or widen scope.

 Topic files live under planning/Brainstorm/(e.g., port-config-and-socket-deployment.md, osjs-base-webapp-init.md).



## Microtask Rules

One task = one primary outcome;; independently understandable/implementable/verifiable/completable. Size scoring: 0-3 good JR task, 4-5 review/split if possible, 6-7 split, 8-10 must split. Hard split rules: >3 independent acceptance outcomes, >3 implementation verbs, multiple verification workflows, multiple architectural decisions, or naturally sequential standalone sub-tasks. Task shape: ID/title, primary outcome, scope, explicit non-scope when useful, acceptance criteria, verification method, dependencies, sizing assessment. Promotion boundary: microtask stays planning material until human promotes to workflow/active_work/ + sync handoff.md. Before promotion use synchronized ICC context, refresh stale affected only through BSW. Microtasks directory: sized, independently verifiable work prepared from approved planning.

 ICC-first context rule applies before creating/refining/sizing/splitting microtasks;; use synchronized relevant ICC;; only uncommitted files differing from audited overlay inspected when HEAD unchanged;; BSW does not authorize execution;; follow rules.md after context validation.