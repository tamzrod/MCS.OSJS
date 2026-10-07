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

PERSIST-002 self-check/regression evidence remains JR DEV evidence only; independent TEST/VERIFY is still deferred to PERSIST-011.

## Current task — ACTIVE

**PERSIST-003 — Derived Persistence RBE Generation**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-003-derived-persistence-rbe-generation.md`
Queue: `workflow/active_work/PERSISTENCE_PROMOTION_QUEUE.md`

Execute PERSIST-003 exactly as written:
- generate one system-owned persistence RBE rule for each configured memory area;
- derive area/start/count exclusively from the authoritative memory layout;
- do not introduce independent persistence range fields;
- do not implement UI locking, user-rule collision policy, filesystem persistence, restore, or runtime snapshot behavior;
- run only task-bounded targeted self-checks and bounded in-scope corrective retests;
- commit and non-force push the completed task under standing JR DEV authority;
- record exact changed paths, checks and resulting source checkpoint;
- prepare PERSIST-004 for the next invocation only after genuine PERSIST-003 completion;
- STOP after PERSIST-003.

## Successor routing

PERSIST-004..022 are already human-promoted and remain QUEUED/dependency-gated.

OpenHands JR DEV may autonomously close and deliver CODE/DISCOVERY packets and select the next already-promoted eligible CODE/DISCOVERY packet for the next invocation.

TEST/VERIFY packets switch to independent JR mode and require their exact current packet. JR DEV must not self-certify those gates.

Only BLACK SHEEP WALL edits ICC. No force push, destructive history rewrite, production/operator-data mutation, or scope expansion.
