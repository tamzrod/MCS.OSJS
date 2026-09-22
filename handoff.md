# Handoff — MCS.OSJS

## Approved execution boundary
Human-approved FMT-001–FMT-008 manifest: `workflow/active_work/FMT_PROMOTION_QUEUE.md`. The manifest is approval and dependency routing only, not a mutable status source. Only the canonical packet location (`workflow/active_work/` or `workflow/archive/`) determines unfinished versus completed lifecycle, with archived acceptance evidence and delivery independently verified. No current-task ID or copied status belongs in handoff.

## Migration caution
Legacy packets may still contain `Status:` fields, planning duplicates and STOP-after-task wording. Treat those as non-authoritative migration debt, not proof of completion. Do not mass-archive based on historical queue claims; reconcile each task's evidence and delivery before moving its canonical packet. A missing or contradictory task is isolated as blocked with a precise report; continue independent eligible tasks if safe. Preserve local unpushed work and do not overwrite it from remote assumptions.

## Boundaries
Execute one eligible authorized packet at a time, record actual evidence, archive only after verified completion, rescan automatically. OTR work retired/non-executable. FMT-005 requires separately authorized exact build instructions; FMT-008 is not authorized for execution by OpenCode. No impersonation, ICC edits outside BLACK SHEEP WALL, destructive operations, unauthorized push, merge, force push or main push. End with coverage report for every approved ID and exact reasons for unresolved work.
