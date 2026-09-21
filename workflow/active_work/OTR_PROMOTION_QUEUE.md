# OTR autonomous promotion and execution queue

Human-approved roadmap: OTR-001..014. Approval is not evidence of completed work or permission for unspecified operations. OPERATION CWAL must select exactly ONE eligible task without asking the human to pick it. Per-task STOP remains mandatory; the next invocation may select the successor only after predecessor evidence is established. No background loop or multiple tasks per invocation.

## Ordered atomic discovery queue
[x] OTR-001A — COMPLETED; OpenCode, shell/nav inventory; packet `otr-001a-electron-shell-inventory.md`.
[x] OTR-001B — COMPLETED; OpenCode, editor screens inventory; packet `otr-001b-electron-editor-inventory.md`.
1. OTR-001C — QUEUED, OpenCode, discovery; eligible only after genuine 001B source inventory evidence (now ready).
2. OTR-002A — QUEUED, OpenCode, discovery; after 001C.
3. OTR-002B — QUEUED, OpenCode, discovery; after 002A.
4. OTR-003A — QUEUED, design; after 001A/001B/002A evidence (001A/001B complete); require explicit assigned owner and source-linked map.
5. OTR-003B — QUEUED, design; after 001C/002B/003A evidence; require explicit assigned owner and source-linked contract map.

`QUEUED` is prior human approval, not a second approval gate. A packet copied with `Status: PLANNING / UNDER REVIEW` retains an obsolete header: this queue and the handoff resolve its *promotion status*, but do not fill missing owner, source, command, safety or evidence details. Do not execute an incomplete packet. Resolve assignment from handoff first; if no current handoff assignment, choose the first unfinished eligible packet with a documented owner, valid scope and established predecessor evidence; otherwise report the exact missing item once and STOP. Never treat file presence or numbering as completion evidence.

`PLANNING / UNDER REVIEW` packets are parent templates awaiting execution assignment. They contain high-level intent and scope but incomplete STOP details. Treat only queued items with explicit owners and completed predecessors as eligible for execution selection. Do not select `PLANNING` packets unless the human specifically directs work on a particular template or asks to fill missing owner/requisites fields manually.