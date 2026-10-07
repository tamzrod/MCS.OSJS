# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work on 2026-10-07. OPERATION CWAL is the invocation; no second approval or task-selection prompt is required for the current packet.

OpenHands has two distinct CWAL modes:
- **JR DEV** — CODE/DISCOVERY implementation for an explicitly assigned active task.
- **Independent JR** — separately assigned TEST/VERIFY execution only.

One invocation uses exactly one mode. JR DEV self-checks are not independent TEST/VERIFY evidence.

## Current task — ACTIVE
**PERSIST-001 — Persistence Configuration Schema**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-001-persistence-configuration-schema.md`
Queue: `workflow/active_work/PERSISTENCE_PROMOTION_QUEUE.md`

This human instruction explicitly supersedes the previous OTR-001C current assignment for now. OTR-001C remains promoted work but is no longer the current CWAL task.

Execute PERSIST-001 exactly as written:
- add the persistence configuration schema/round-trip semantics only;
- preserve Port → Unit ID → Memory as authority;
- do not add duplicate persistence-owned start/count/area identity;
- do not implement RBE derivation, snapshot filesystem behavior, restore, Raw Ingest restore, or unseal behavior in this task;
- run only task-bounded targeted self-checks;
- record the exact changed paths, checks and resulting source checkpoint;
- STOP after PERSIST-001.

## Persistence successor routing
PERSIST-002..022 are already human-promoted in Active Work. They are authorized for their written scopes but remain QUEUED.

After genuine completion evidence for the current task is persisted, a later OPERATION CWAL invocation may select the next eligible persistence task according to `PERSISTENCE_PROMOTION_QUEUE.md`. Exactly one task per invocation.

CODE/DISCOVERY successors may be assigned to OpenHands JR DEV. TEST/VERIFY tasks require a **separate independent JR invocation and exact current test packet**; the JR DEV coding run must never certify those gates.

Only BLACK SHEEP WALL edits ICC. No production/customer data, live service actions, destructive cleanup, global restart, or unrelated repository mutation is authorized unless the selected task explicitly grants it.
