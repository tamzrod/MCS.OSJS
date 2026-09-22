# Active work — current reset and future queue behavior

Current state remains RESET: root `handoff.md` is NONE and `OTR_PROMOTION_QUEUE.md` is retired. Historical OTR packets/evidence are not execution authority; preserve them pending separately authorized cleanup. No task may execute from this directory until a human-approved fresh queue is explicitly recorded in handoff and its queue metadata. Merely placing packets here is insufficient.

## Execution after fresh queue activation
Read `AGENTS.md`, `operation cwal.md`, root `handoff.md`, the approved queue and selected packet. CWAL scans nonterminal packet headers, selects any eligible task with verified COMPLETE prerequisites (lowest stable ID by default), activates and processes exactly one at a time, records genuine evidence and result, then rescans. If a task FAILs/BLOCKs, continue another independent eligible task or eligible prerequisite; do not bypass failed blockers or retry indefinitely. Stop when every task is terminal or no eligible work remains; report unresolved PENDING and reasons. An agent cannot impersonate another task's assigned role; checkpoint at required independent TEST/VERIFY handoff. No unauthorized production changes, ICC edits or main push.

Fresh planning must inspect actual product code and validate historical evidence, create small dependency-aware packets and obtain human approval. Do not fabricate PASS or expand task scope.
