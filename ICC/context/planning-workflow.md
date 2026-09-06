# Planning Workflow

Baseline commit: 500376cfb5c222298aadfcf035aad0af0a635773
Working tree: clean
Source dependencies: planning/README.md, planning/Brainstorm/README.md, planning/microtask/README.md, planning/microtask/rules.md
Parent: L0-project
Zoom In: brainstorm-topics
Zoom Out: L0-project

## Model

Planning is human-owned; contains brainstorm material and decomposed microtasks NOT authorized for JR execution. Flow:  PLANNING OPERATION -> ICC FIRST -> CURRENT? USE CONTEXT / STALE? BSW REFRESH AFFECTED ONLY -> BRAINSTORM -> MICROTASK -> PROMOTION -> workflow/active_work/. Planning does not grant execution authority; architectural questions may remain unresolved and must not be silently treated as decisions during execution..

## Brainstorm

Human-owned exploration/framing/alternatives/questions/candidates. Not execution authority. ICC-first before repository-dependent brainstorming. BSW may refresh stale/missing affected context only;cannot authorize implementation or widen scope. Topic files live under planning/Brainstorm/ (current:  mma2-basic-install-test.md, osjs-modbus-simulator.md(; prior topics port-config-and-socket-deployment.md and osjs-base-webapp-init.md were retired when superseded..

## Microtask Rules (per planning/microtask/rules.md

- Core rule:  One task = one primary outcome; independently understandable/implementable/verifiable/completable. A single narrative outcome does not automatically mean one JR-sized task; split across independently failure-prone execution boundaries..
- File boundary:  One microtask = one detailed task file; one task ID + one primary outcome per file; feature may require many files (task-001/002/003 pattern(.
- Size scoring (five dimensions, 0-2 points each(:   implementation surface, env/dependency uncertainty, behavioral surface, verification surface, decision/recovery surface. 0-3 good JR task; 4-5 split unless tightly coupled+one deterministic verification workflow; 6-7 split; 8-10 must split. Do not reduce score because all work serves one feature..
- Mandatory split triggers (>3 independent acceptance outcomes,( >3 implementation verbs,( multiple distinct verification workflows,( multiple architectural decisions,( naturally sequential standalone sub-tasks,( combines donor/import/toolchain establishment with runtime behavioral verification,( combines environment discovery with product implementation unless env already known+established,( early-phase failure would force abandon/reinterpret later-phase work(.
- Donor/Runtime rule:  default decomposition IMPORT/ESTABLISH -> BUILD/STATIC VERIFY, then RUN -> BEHAVIORAL/PROTOCOL VERIFY; keep combined only when import, toolchain, runtime, verification path already proven + no independent investigation expected..
- Context-budget check:  can JR stay inside one semantic branch + one bounded working set? If not, split at strongest independently verifiable boundary..
- Task shape:  one ID+title, primary outcome, scope, non-scope when useful, acceptance criteria, verification method, dependencies, sizing assessment..
- Promotion boundary:  microtask stays planning until human promotes to workflow/active_work/ + sync handoff.md; before promotion use synchronized ICC context for selected task+workflow state; refresh only stale affected through BSW..

## Current Planning Material (non-authoritative for execution

- MMA2 basic install+test brainstorm + microtask (MMA2-001/002( — source of the active MMA2-001/002 active work; MMA2-001 later split in execution into import+build vs runtime smoke-test (active MMA2-002 preserves internal execution order (+completion state( as repository truth(.
- OS.js Modbus Simulator brainstorm (approved simulator UI, two parameter domains, shared MMA2 config boundary, save/apply routing, runtime status( — source of active SIM-001…SIM-007 program. See brainstorm-topics for content summary..
