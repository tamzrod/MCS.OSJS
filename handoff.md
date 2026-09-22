# Handoff — MCS.OSJS

## Current queue
APPROVED / ACTIVE QUEUE: `workflow/active_work/FMT_PROMOTION_QUEUE.md` — FMT-001–FMT-008. Human approved promotion and continuous dependency-aware processing. This authorizes selection of eligible packets within this queue only; no new tasks or broader permissions.

## Current task
QUEUE SELECTION — no packet currently executing. This is NOT `NONE` and is not a retired queue. On Operation CWAL invocation validate the manifest and all eight packet headers, select one eligible task, checkpoint exactly that packet ACTIVE in packet/queue/handoff, execute within its authority, record outcome, and rescan. Never mark all eight ACTIVE simultaneously. If manifest or packet integrity fails, BLOCKED/STOP without fallback to historical OTR.

## Historical boundary
Old OTR queue, packets, archives and evidence are non-executable and do not prove completion. Preserve evidence; invalid OTR-002B and OTR-003A archives remain invalid. ICC edits require BLACK SHEEP WALL authority.

## Execution boundary
One task at a time; continue through eligible independent tasks after each outcome. Enforce prerequisites, role handoffs and exact packet delivery. Do not fabricate PASS, retry FAIL/BLOCKED in the same invocation, or perform unauthorized production actions, merge, force push or main push. FMT-005 requires separately supplied exact test command/environment before execution.
