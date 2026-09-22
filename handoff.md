# Handoff — MCS.OSJS

## Current queue
APPROVED / ACTIVE QUEUE: `workflow/active_work/FMT_PROMOTION_QUEUE.md` — FMT-001–FMT-008. Human approved promotion and continuous dependency-aware processing. This authorizes selection of eligible packets within this queue only; no new tasks or broader permissions.

## OTR eligibility note
OTR-002 is ineligible (electron/Chrome inventory only, non-product-boundary); defer and archive as-is. Only execute FMT packets from the approved queue. ICC edits require BLACK SHEEP WALL authority.

## Current task
FMT-004 checkpointed as ACTIVE in queue; executing backend boundary inventory per FMT_PROMOTION_QUEUE.md. All FMT prerequisites verified via dependency-aware scanning (FMT-001, FMT-002, FMT-003 complete). Proceed within FMT-004 scope only; upon completion rescan for next eligible packet (FMT-005 dependent on FMT-004, or FMT-007 independent if prerequisites satisfied).

## Execution boundary
One task at a time; continue through eligible independent tasks after each outcome. Enforce prerequisites, role handoffs and exact packet delivery. Do not fabricate PASS, retry FAIL/BLOCKED in the same invocation, or perform unauthorized production actions, merge, force push or main push. FMT-005 requires separately supplied exact test command/environment before execution.
