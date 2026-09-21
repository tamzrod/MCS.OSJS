# OTR remaining atomic microtask decomposition
Status: PLANNING / REVIEWED DECOMPOSITION, not an execution packet. Human-approved roadmap OTR-001..014; OTR-001A is the sole ACTIVE execution task. Parent files are references, not executable replacements for their children. Source of sizing: `planning/microtask/otr-sizing-and-decomposition-review.md` and `planning/microtask/rules.md`.

## Rules for each child before it becomes executable
One ID, one owner and stage, one primary outcome, <=3 production files plus tests, <=3 acceptance outcomes; explicitly pin actual source and target paths, revision, dependencies, Previous/Next, exact tests/actions and expected evidence. CODE owner OpenCode where explicitly assigned; independent TEST/VERIFY owner OpenHands/JR only when explicitly assigned. A child without source-pinned paths or reproducible verification remains a planning candidate, even though the parent roadmap was approved. Only OTR-001A ACTIVE; all others must not execute by inference. Do not change ICC. Do not touch production YAML, live devices or services without specific authority. STOP after each packet.

## Phase 1 — UNDERSTAND (already split into individual planning files)
- OTR-001A shell/navigation inventory (ACTIVE), 001B editor screens/assets inventory, 001C Electron backend boundary inventory.
- OTR-002A OS.js UI/launch baseline, 002B OS.js backend/config/service baseline.
- OTR-003A source-to-target UI parity map, 003B backend/config contract map. Each is read-only, with its own source-anchored artifact and no code edits. Run only one explicitly assigned packet at a time.

## Phase 2 — REPLICATE UI (create each named child as its own detailed file after 003A establishes exact views)
- OTR-004A CODE: unified OS.js Toolkit shell structure and navigation, no editor internals. OTR-004B CODE: shell CSS/assets parity, no behavior change. OTR-004-T independent static/build TEST of committed shell. OTR-004-V independent rendered shell/navigation VERIFY. Dependencies 003A and predecessor child as applicable.
- OTR-005A CODE: Simulator screen container/navigation only. OTR-005B[n] CODE: ONE real Simulator editor view/tab per source-verified view (number and IDs fixed from 001B/003A; do not invent tabs). OTR-005C CODE: Simulator draft/discard state, only if not intrinsic to a bounded view. Each committed CODE slice gets its own 005-[child]-T independent TEST; one actual screen/scenario per 005-[child]-V VERIFY. Dependencies 004 and verified map.
- OTR-006A CODE: Replicator screen container/navigation. OTR-006B[n] CODE: ONE source-verified Replicator view/tab each. OTR-006C CODE: draft/discard if independently verifiable. Separate per-child TEST and per-screen VERIFY. Dependencies 004, 001B, 003A.
- OTR-007A CODE: MMA memory overview/list. OTR-007B[n] CODE: ONE source-verified MMA editor view each. OTR-007C CODE: read-only/draft states only if separate from a view. Separate TEST and VERIFY. Dependencies 004, 001B, 003A.

## Phase 3 — CONNECT BACKEND (runtime-specific; R means Simulator, Replicator or MMA, NEVER an unspecified runtime in an executable packet)
- OTR-008-R-A CODE: read-only backend configuration read adapter for ONE verified runtime and contract; OTR-008-R-B CODE: UI load wiring if independent or >3 files; OTR-008-R-T independent fixture/static TEST. No writes.
- OTR-009-R-A CODE: one runtime's verified schema validator; OTR-009-R-B CODE: unknown-field-preserving roundtrip/serialization if distinct; OTR-009-R-T independent valid/invalid/roundtrip TEST. Never infer schema from UI alone.
- OTR-010-R-A CODE: disposable-target safe write and conflict/failure backend contract; split atomicity/conflict into a further child if independently testable. OTR-010-R-B CODE: UI save wiring and failure display; OTR-010-R-T independent disposable fixture TEST. No production/operator YAML writes.
- OTR-011-R-D DISCOVERY: inspect whether scoped restart exists and its exact authorization/permission boundary, no service action. Only if supported: OTR-011-R-A CODE backend scoped restart adapter; OTR-011-R-B CODE UI action wiring; OTR-011-R-T mock/disposable TEST; OTR-011-R-V authorized runtime VERIFY. Unsupported => document limitation, do not create fictional adapter or PASS.
- OTR-012-R-A CODE: independently observed applied-state/status read adapter; OTR-012-R-B CODE: UI pending/applied/disconnected/error states if separate; OTR-012-R-T independent status-state TEST; OTR-012-R-V independent runtime VERIFY. A save/restart acknowledgement is not applied-state evidence.
For each R, order read -> validate/roundtrip -> safe save -> supported restart -> observed applied status. Pin precise source files, synthetic fixture and tests after 003B. No cross-runtime acceptance.

## Phase 4 — VERIFY (independent JR, one scenario per packet)
- OTR-013-SHELL, OTR-013-SIM, OTR-013-REP, OTR-013-MMA: one rendered UI parity VERIFY each; split further by independent view/dialog if needed. Pin viewport, reference and target screenshots, exact click/input sequence, expected differences and evidence transport. No product edits.
- OTR-014-R-H: one disposable happy-path load/edit/validate/save/(supported and authorized restart)/independent readback VERIFY per runtime R. Separate OTR-014-R-INVALID, -DISCARD, -CONFLICT, -WRITEFAIL, -RESTARTFAIL (only where restart supported), -STALE: each one negative scenario, only if applicable and safe. Pin disposable fixture path, source revision, exact commands/actions, rollback, expected observation and raw report. No production config, live devices or global restart.

## Promotion and STOP gates
These are candidate child IDs and outcomes, not fabricated execution-ready tasks. Materialize a detailed individual file only when the relevant inventory establishes actual paths, view count and contract. CODE, TEST and VERIFY are separate authority and evidence gates. Promotion to `workflow/active_work/` is allowed only after the packet is source-pinned and reviewed; one ACTIVE, others QUEUED, with handoff identifying exactly one task. No automatic successor execution or fabricated PASS. Parent roadmap remains approved without making incomplete children executable.
