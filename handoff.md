# Handoff — MCS.OSJS

## Current queue
FMT-001–FMT-008 — HUMAN APPROVED FOR PROMOTION AND CONTINUOUS DEPENDENCY-AWARE PROCESSING. Queue manifest: `workflow/active_work/FMT_PROMOTION_QUEUE.md`. This authorization does not make all packets simultaneously ACTIVE. The CWAL runner selects one eligible packet at a time, checkpoints it ACTIVE, executes, records its outcome, and rescans until no eligible work remains.

## Current task
NONE — no packet is executing at this checkpoint. On Operation CWAL invocation, use the approved FMT queue selection procedure; NONE here means no in-flight task, not cancellation of the approved queue. If queue manifest or packets are absent/inconsistent, report BLOCKED and STOP without falling back to historical OTR.

## Historical boundary
Old OTR packets, archives and reports are non-executable and not evidence of completion. Preserve them; never select them. Invalid OTR-002B and OTR-003A archives remain invalid. Do not edit ICC except through BLACK SHEEP WALL.

## Execution boundary
One packet at a time; continue through eligible independent packets after a terminal result, respecting blockers and agent roles. Never bypass a prerequisite, fabricate PASS, retry FAIL/BLOCKED in the same invocation, or exceed packet permissions. No merge, force push or main push without separate authority.
