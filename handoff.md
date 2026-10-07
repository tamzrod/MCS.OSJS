# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work on 2026-10-07. OPERATION CWAL is the invocation; no second approval or task-selection prompt is required for the current packet.

OpenHands has two distinct CWAL modes:
- **JR DEV** — CODE/DISCOVERY implementation for an explicitly assigned active task.
- **Independent JR** — separately assigned TEST/VERIFY execution only.

One invocation uses exactly one mode. JR DEV self-checks are not independent TEST/VERIFY evidence.

## Completed predecessor
**PERSIST-001 — Persistence Configuration Schema**
CODE COMPLETE at source checkpoint `faa33929429a0382a64b78bc773f0074e11cb6b1`.
Its task packet records the bounded self-check/regression evidence. This is source/CODE completion only; it is not an independent TEST/VERIFY PASS.

## Current task — ACTIVE
**PERSIST-002 — State Sealing Prerequisite Validation**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-002-state-sealing-prerequisite-validation.md`
Queue: `workflow/active_work/PERSISTENCE_PROMOTION_QUEUE.md`

Execute PERSIST-002 exactly as written:
- reject persistence-enabled memory when State Sealing is absent or disabled;
- allow persistence disabled with sealing either enabled or disabled;
- do not silently enable or mutate State Sealing;
- preserve existing State Sealing semantics and Port → Unit ID → Memory authority;
- do not implement derived persistence RBE, filesystem persistence, restore, Raw Ingest restore, or unseal behavior in this task;
- run only task-bounded targeted self-checks and bounded in-scope corrective retests;
- record exact changed paths, checks, and resulting source checkpoint;
- STOP after PERSIST-002.

## Persistence successor routing
PERSIST-003..022 are already human-promoted in Active Work and remain QUEUED/dependency-gated.

After genuine completion evidence for PERSIST-002 is persisted, a later OPERATION CWAL invocation may select PERSIST-003 according to `PERSISTENCE_PROMOTION_QUEUE.md`. Exactly one task per invocation.

CODE/DISCOVERY successors may be assigned to OpenHands JR DEV. TEST/VERIFY tasks require a **separate independent JR invocation and exact current test packet**; the JR DEV coding run must never certify those gates.

Only BLACK SHEEP WALL edits ICC. No production/customer data, live service actions, destructive cleanup, global restart, or unrelated repository mutation is authorized unless the selected task explicitly grants it.
