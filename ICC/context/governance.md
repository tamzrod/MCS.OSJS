# Governance and Operation Discipline

Baseline commit: 500376cfb5c222298aadfcf035aad0af0a635773
Working tree: clean
Source dependencies: AGENTS.md, BLACK_SHEEP_WALL.md, ICC/INDEX.md, operation cwal.md, handoff.md, workflow/active_work/README.md
Parent: L0-project
Zoom In: (none; leaf node(; indexing/active-work program via L0 half; implementation details via active-work node
Zoom Out: L0-project

## Operation Chain (as rooted in AGENTS.md

For every repository-dependent operation:
```text
LOCATE COMMAND
→ READ BLACK_SHEEP_WALL.md
→ READ ICC/INDEX.md FIRST
→ VALIDATE ICC BASELINE AGAINST CURRENT HEAD + WORKING TREE
→ USE ICC IF CURRENT
→ IF STALE, REFRESH ONLY AFFECTED CONTEXT THROUGH BLACK SHEEP WALL
→ LOCATE AUTHORITY
→ ACT WITHIN OPERATION SCOPE
→ VERIFY AGAINST REPOSITORY SOURCE
→ NEVER GUESS
```

## ICC State Model

Repository truth at BASELINE COMMIT + audited uncommitted working-tree overlay. Baseline commit;the HEAD used at audit time (need not equal the commit that stores ICC files(. For uncommitted files, record deterministic fingerprints per audited path. A context is synchronized when:  committed deps represented by current baseline; no changed source dep unincorporated; uncommitted deps match audited overlay fingerprints. Known-stale context is never consumed as authoritative. Unaffected synchronized context is not recomputed.



## BLACK SHEEP WALL (modes

- Direct invocation = repository-wide context compaction:  current HEAD -> audit all relevant context -> discover semantic boundaries -> compact into ICC/context/ -> register in index -> stamp baseline -> record audited uncommitted overlays. Subsequent runs:  do NOT rescan unchanged files; inspect only committed files changed since baseline + uncommitted files differing from audited overlay; refresh only affected context; preserve unaffected. If HEAD unchanged, only new/modified/renamed/deleted uncommitted files differing from audited overlay are inspected.
- Delegated branch refresh = invoked by CWAL/planning/etc with a selected semantic node as navigation ceiling; refresh only that node/required children; never repair stale siblings or widen scope..

## Operation CWAL

Mandatory execution discipline. Sequence: READ AGENTS + BLACK_SHEEP_WALL + ICC/INDEX FIRST, VALIDATE ICC vs HEAD+working tree, USE ICC if current, REFRESH only stale affected, LOCATE AUTHORITY (handoff.md + workflow/active_work/(, SELECT ACTIVE TASK SEMANTIC BRANCH (navigation ceiling(, FOLLOW WORKFLOW, ACT (scoped(, VERIFY vs evidence, NEVER GUESS. Never:  promote/invent work, expand scope, inspect Planning to select/future tasks,, silently resolve architectural questions. If ctx stale/missing/deps missing:  stop, report missing authority, never infer. Active Work consumed only through CWAL; workflow/active_work/ contains only human-promoted tasks; if empty:  stop and wait for human promotion..

## Workflow Boundaries (per AGENTS

- Brainstorm/Planning:  ICC first, refresh affected only through BSW; human-owned, no implementation authority
- Microtask:  ICC first, refresh stale affected; follow planning/microtask/rules.md
- Promotion:  ICC first, refresh stale; promote only human-selected microtask; sync handoff.md
- CWAL:  ICC first for Active Work context; executes only authorized Active Work.



## Key Operational Rules

^- The ICC index is a registry, not a navigation permission menu. Registry presence is not relevance; crossing semantic branches requires the current authorized task to explicitly require that boundary for implementation/understanding/testing/verification.
- Zoom Out must not climb above the task-selected navigation ceiling and re-enter an unrelated sibling branch.
- An operation must not reopen unchanged repository source files when synchronized ICC contains the needed context.
- BLACK SHEEP WALL cannot grant execution authority, promote work, choose future work, let CWAL inspect Planning to select tasks, let brainstorming implement code, widen a bounded operation, or convert unrelated ICC staleness into work for the invoking operation..
