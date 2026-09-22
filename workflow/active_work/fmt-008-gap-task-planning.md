# FMT-008 — Gap task planning
Task ID: FMT-008
Task Name: Draft one source-proven shell/navigation implementation microtask
Blocker Task: FMT-005, FMT-006, FMT-007
Status: PENDING
Assigned Agent: ChatGPT — planning only
Stage: DESIGN

## Objective
Draft exactly one independently sized CODE planning packet for the first confirmed shell/navigation gap; if none exists, report NO TASK REQUIRED without inventing work.
## Scope
Read `ICC/INDEX.md`, `planning/microtask/rules.md`, accepted FMT-005/006/007 evidence and source files cited for the selected gap. Write one new task file under `planning/microtask/` only after human approval of this packet, or one no-gap report at `planning/microtask/evidence/fmt-008-no-gap.md`. No product, ICC, workflow or active queue edits.
## Execution
1. Verify blocker evidence and choose one source-proven shell/navigation gap from FMT-006, checking build and backend constraints.
2. If no proven gap, document no-task outcome and STOP; otherwise draft one packet with stable ID, Task Name, Blocker Task, owner, exact paths, bounded actions, <=3 acceptance outcomes, evidence, sizing and explicit TEST/VERIFY follow-up candidates.
3. Leave new packet PENDING for human review; do not activate or draft additional gap packets in this invocation.
## Acceptance Criteria
1. Chosen gap has cited source evidence and is not already implemented; or a supported no-gap report is produced.
2. At most ONE new task file is created, complete against current rules and preferred size, with actual blocker IDs.
3. No product code, ICC, active queue or handoff is changed.
## Evidence
One source-pinned planning packet or `planning/microtask/evidence/fmt-008-no-gap.md`, including accepted evidence refs, current SHA, sizing and selection rationale.
## Completion and delivery
COMPLETE when one valid packet or documented no-gap outcome exists; FAIL for unsupported packet; BLOCKED for invalid prerequisite. No execution or commit/push authority until promotion specifies exact ref and delivery. STOP.
## Sizing
Implementation 0; environment 0; behavior 0; verification 1; decision/recovery 1 = 2/10. One gap packet, not entire implementation roadmap.