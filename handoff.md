# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work.

OpenHands modes remain separate:
- **JR DEV** — assigned CODE/DISCOVERY implementation plus bounded workflow continuation.
- **Independent JR** — separately assigned TEST/VERIFY only.

The normal operator loop is:

```text
git pull
OPERATION CWAL
```

One invocation executes exactly one product task, may deliver/record it, may prepare the next already-promoted eligible task, then STOPS.

## Completed predecessors

**PERSIST-001 — Persistence Configuration Schema**
CODE COMPLETE at `faa33929429a0382a64b78bc773f0074e11cb6b1`.

**PERSIST-002 — State Sealing Prerequisite Validation**
CODE COMPLETE and delivered on GitHub main at
`fe3bda9e5c863c959500514e8c6083d2514526d1`.

**PERSIST-003 — Derived Persistence RBE Generation**
CODE COMPLETE at `4e8f774eb946cc9a946938e4f3a28eaa13de5c9d`, delivered on main.

**PERSIST-004 — Locked System RBE Behavior**
CODE COMPLETE at `936ab24560d57cbf8aacc8180bc00a2527dcfad5`, delivered on main.

**PERSIST-005 — Memory Range Synchronization**
CODE COMPLETE at `c743174204455afa83bc92b859fd47582a7afdbc`, delivered on main.

**PERSIST-006 — Memory Area Removal Synchronization**
CODE COMPLETE at `9ac48cb23e369fb0384b3d8f00811435d5bf7b69`, delivered on main by this invocation.

PERSIST-002..PERSIST-006 self-check/regression evidence remains JR DEV evidence only; independent TEST/VERIFY is still deferred to PERSIST-011.

## Current task — ACTIVE

**PERSIST-007 — Persistence Disable Cleanup**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-007-persistence-disable-cleanup.md`
Queue: `workflow/active_work/PERSISTENCE_PROMOTION_QUEUE.md`

Execute PERSIST-007 exactly as written:
- remove persistence-owned RBE projections when persistence is disabled for a memory;
- user RBE rules survive unchanged;
- re-enabling persistence regenerates projections from current memory ranges;
- do not delete or rewrite user-owned RBE entries;
- run only task-bounded targeted self-checks and bounded in-scope corrective retests;
- commit and non-force push the completed task under standing JR DEV authority;
- record exact changed paths, checks and resulting source checkpoint;
- prepare PERSIST-008 for the next invocation only after genuine PERSIST-007 completion;
- STOP after PERSIST-007.

## Successor routing

PERSIST-008..022 are already human-promoted and remain QUEUED/dependency-gated.

OpenHands JR DEV may autonomously close and deliver CODE/DISCOVERY packets and select the next already-promoted eligible CODE/DISCOVERY packet for the next invocation.

TEST/VERIFY packets switch to independent JR mode and require their exact current packet. JR DEV must not self-certify those gates.

Only BLACK SHEEP WALL edits ICC. No force push, destructive history rewrite, production/operator-data mutation, or scope expansion.
