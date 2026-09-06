# Governance and Operation Discipline

Baseline commit: cb869da3caeed040478e67ecb0a01c5f93b3a66b
Working tree: clean
Source dependencies: AGENTS.md, BLACK_SHEEP_WALL.md, ICC/INDEX.md, operation cwal.md, handoff.md, workflow/active_work/README.md

## Context Model (ICC-First

ICC is the first-read context cache of repository truth. Repository-dependent operations: READ ICC INDEX first, LOCATE RELEVANT CONTEXT, VALIDATE baseline/overlay, USE ICC if current; refresh only stale/missing affected context through BLACK SHEEP WALL. An operation must not reopen unchanged source files merely because they exist; consult sources only when ICC missing/stale/insufficient/needs verification.

ICC validity: committed deps represented by current baseline commit + no changed source dep unincorporated + uncommitted deps match audited overlay fingerprints. If HEAD differs from baseline: diff baseline..HEAD, refresh only contexts whose declared dependencies intersect changed files. Do not full-rescan unless baseline missing/unusable/incremental impossible.



## BLACK SHEEP WALL

Maintains ICC as compact cache of repository truth; not an authorization mechanism. Direct invocation = repository-wide context compaction: current HEAD -> audit all relevant context -> discover semantic boundaries -> compact into ICC/context/ -> register contexts in index -> stamp baseline commit -> record audited uncommitted overlay (content hash per audited path). Subsequent runs: do NOT rescan unchanged files; inspect only committed files changed since baseline + uncommitted files differing from last audited overlay; refresh only affected context; preserve unaffected. If HEAD unchanged, only new/modified/renamed/deleted uncommitted files differing from audited overlay are inspected.

 Can't: grant execution authority, promote work, choose future work, let CWAL inspect Planning, let brainstorming implement code, widen a bounded operation.


## Operation CWAL

Mandatory execution discipline. Sequence: READ AGENTS + BLACK_SHEEP_WALL + ICC/INDEX FIRST,VALIDATE ICC vs HEAD+working tree,USE ICC if current,REFRESH only stale affected, LOCATE AUTHORITY(handoff.md + workflow/active_work/), FOLLOW WORKFLOW, ACT(s coped scope), VERIFY vs evidence, NEVER GUESS. Never: promote/invent work, expand scope, silently resolve architectural questions. If ctx stale/missing/deps missing: stop, report missing authority, never infer.

 Active Work consumed only through CWAL. workflow/active_work/ contains only human-promoted tasks; if empty: stop and wait for human promotion. Do not read Planning to choose work.



## Workflow Boundaries (per AGENTS)

- Brainstorm/Planning: ICC first, refresh affected only through BSW; human-owned, no implementation authority
- Microtask: ICC first, refresh stale affected; follow planning/microtask/rules.md
- Promotion: ICC first, refresh stale; promote only human-selected microtask; sync handoff.md
- CWAL: ICC first for Active Work context; executes only authorized Active Work.