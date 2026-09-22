# Microtask Rules

## Authority and lifecycle: location is the single source of truth
Read `ICC/INDEX.md` and verify source for planning. Only BLACK SHEEP WALL may edit ICC. Human approval promotes a task into `workflow/active_work/`; planning files alone never authorize execution. Historical OTR packets remain non-executable. A packet has exactly one canonical lifecycle location:
- `planning/microtask/<id>.md`: proposed, not authorized.
- `workflow/active_work/<id>.md`: approved, unfinished work; ready for consideration, not necessarily executable. Dependencies, role, environment and explicit action permissions still apply.
- `workflow/archive/<id>.md`: finished and delivered with verified evidence. Archive is a terminal outcome only when its report states actual COMPLETE with acceptance and delivery proof. Failed/blocked/incomplete tasks remain in active_work and are skipped for the current invocation with a documented reason.

Do not put `Status:`, `ACTIVE`, `PENDING`, `READY`, `COMPLETE`, `FAIL` or `BLOCKED` lifecycle fields inside packets or duplicate them in a queue table or handoff. These words may appear in evidence verdicts and transient reports, not as competing task-state metadata. Never maintain a second executable copy of a promoted packet in planning: retain a non-authoritative pointer or remove the duplicate through separately authorized safe migration. Never archive merely to clear a queue or because a chat claims completion. Move to archive only after verifying exact acceptance evidence, any required independent gates, and delivery. Preserve evidence separately and include its exact reference in the archive record.

## Required packet fields
Task ID: <unique stable ID>
Task Name: <one measurable outcome>
Blocker Task: NONE | <comma-separated exact task IDs>
Assigned Agent: <agent and role>
Stage: DISCOVERY | DESIGN | CODE | TEST | VERIFY

Include objective, exact read/write scope and forbidden actions, bounded execution (exact commands for TEST/VERIFY), at most three acceptance outcomes, exact evidence path and required observations, completion/delivery gates including commit/push authority or its absence, sizing and rationale. No `Status:` field and no per-packet `STOP after one task` instruction that ends the CWAL queue. Packet RETURN/STOP means end that task only; coordinator immediately rescans.

## Blockers and continuous selection
Blocker IDs must resolve to approved queue tasks, with no cycles, duplicates or retired references. A dependency is satisfied only when its canonical packet is in archive AND its evidence/delivery meet its acceptance gates. An active_work packet with unmet blockers, missing authorization, unsupported evidence, wrong agent or unsafe environment remains in active_work; record the precise reason in the invocation report or a single dedicated blocker ledger, not a duplicate status field. Never pretend a blocker was completed or archive a blocked task. Skip it for this invocation and process the next independent eligible packet.

A human-approved queue and root handoff identify the execution boundary, not per-task state. Enumerate canonical active_work packets; select the lowest eligible ID, execute one at a time, record evidence and delivery, archive only after verified completion, then rescan. Do not reattempt a failed/blocked task in the same invocation without new authorization. Continue until no eligible tasks for the actual agent remain, then report every approved task: archived with evidence, or active_work with exact blocker/required actor. If a task is interrupted, preserve its work/evidence and keep it in active_work; resume based on source and evidence, not chat memory.

Global STOP is reserved for explicit human STOP or a genuine repository-wide safety/integrity problem that makes reliable selection impossible. A single missing packet, contradictory historical status, failed task or blocked dependency is a task-level obstacle; isolate it and continue other independently safe work. No autonomous destructive cleanup, merges, force push, main push, ICC edits or unauthorized actions.

## Roles and verification
CODE, independent TEST and live/UI VERIFY are separate by default. Never impersonate another agent or self-certify independent PASS. JR requires an exact authorized test packet, raw evidence, post-check and delivery; a failed gate ends that task, not the queue coordinator. Only an authorized coordinator may archive a task after required independent verification. Handoff to another role when this runner's own eligible work is exhausted.

## Sizing and promotion
Score implementation, environment, behavior, verification and recovery 0–2 each. Prefer 0–3; split 4–5 unless tightly coupled, split 6–7, and always split 8–10. Split when more than three acceptance outcomes, multiple independent workflows, uncertain environment plus implementation, or CODE/TEST/VERIFY combined. Only human-approved packets are promoted. Promotion authorizes selection within exact packet scope, not new tasks or expanded privileges.
