# OTR autonomous promotion and execution queue

Human-approved roadmap: OTR-001..014. Approval is not evidence of completed work or permission for unspecified operations. OPERATION CWAL must select exactly ONE eligible task without asking the human to pick it. Per-task STOP remains mandatory; the next invocation may select the successor only after predecessor evidence is established. No background loop or multiple tasks per invocation.

## Ordered atomic discovery queue
1. OTR-001A — ACTIVE; OpenCode; read-only; authoritative packet `otr-001a-electron-shell-inventory.md`.
2. OTR-001B — QUEUED, OpenCode, discovery; eligible only after genuine 001A source inventory evidence.
3. OTR-001C — QUEUED, OpenCode, discovery; after 001B.
4. OTR-002A — QUEUED, OpenCode, discovery; after 001C.
5. OTR-002B — QUEUED, OpenCode, discovery; after 002A.
6. OTR-003A — QUEUED, design; after 001A/001B/002A evidence; require explicit assigned owner and source-linked map.
7. OTR-003B — QUEUED, design; after 001C/002B/003A evidence; require explicit assigned owner and source-linked contract map.

`QUEUED` is prior human approval, not a second approval gate. A packet copied with `Status: PLANNING / UNDER REVIEW` retains an obsolete header: this queue and the handoff resolve its *promotion status*, but do not fill missing owner, source, command, safety or evidence details. Do not execute an incomplete packet. Resolve assignment from handoff first; if no current handoff assignment, choose the first unfinished eligible packet with a documented owner, valid scope and established predecessor evidence; otherwise report the exact missing item once and STOP. Never treat file presence or numbering as completion evidence.

## OTR-004..014 — approved roadmap, BLOCKED as execution packets
The copied OTR-004..014 files are parent templates, NOT executable CODE/TEST/VERIFY microtasks. `planning/microtask/otr-remaining-atomic-decomposition.md` defines candidate children. After the relevant inventories, the workflow owner creates individual source-pinned children with exact paths, owner, commands, acceptance, evidence, safe permissions and Previous/Next; then activates ONE child. Do not invent editor tabs, runtime contracts, production access, restart mechanisms or JR PASS. If an eligible child does not yet exist, report BLOCKED for packet materialization rather than running the parent. Do not ask the human to reapprove the already approved roadmap; ask only for genuinely new unsafe/out-of-scope authority.

Only BLACK SHEEP WALL edits ICC. No production config or live service actions without explicit authorization. No automatic commit/push unless task grants it.
