# Fresh Toolkit backlog — HUMAN REVIEW ONLY

Date: 2026-09-22. Branch: opencode. Status: PLANNING / NOT APPROVED / NOT EXECUTABLE. This file is not an active queue. Root handoff remains NONE. Every task below is PENDING; blocker-free does NOT grant execution authority. On approval, materialize each task as a separate fully scoped packet under planning/microtask, with exact paths, source SHA, allowed commands, evidence and five-dimension sizing, then promote only one explicitly. One task per CWAL invocation; STOP. No ICC edits, production YAML, live service actions or main push.

## Goal
One independently deployed OS.js MCS Modbus Toolkit matching Electron's actual UI, navigation and configuration workflows for Simulator, Replicator and MMA memory, with genuine status/errors; preserve OS.js shell, Start menu, taskbar and one Toolkit desktop icon. Reuse verified implementation; no Electron runtime dependency. Complete only with independent screen and safe end-to-end evidence.

## Observed baseline and trust boundary
- `OSJS/src/packages/MCSModbusToolkit/index.js` creates ONE Toolkit window, uses `toolkit-renderer`, Memory and Replicator contracts/transports/editors, and read-only Diagnostics. This is existing source, not proof of successful build/runtime parity.
- `OSJS/src/packages/MCSModbusToolkit/memory-contract.js` defines load/apply/status; `replicator-contract.js` defines load/apply/status/suggest. Their old comments claim dormant/unwired, while index.js imports them: comments and historical reports are not reliable current-state proof.
- Historical `workflow/active_work/evidence/otr-002a-report.md` invents or inconsistently describes fixture/live paths and reports a clean branch plus detached HEAD. `otr-003a-map.md` contradicts actual contract operations and mislabels its backend mapping as UI parity. Do not treat these as accepted evidence. Previously rejected archived 002B and 003A remain invalid.
- Electron reference exists in `electron/` with main.js, preload and UI-related files; exact screen/asset inventory and rendered parity remain unverified in this review. Linux backend endpoint, restart, save safety and deployment acceptance are NOT verified. No runtime/build/test executed during this planning pass.

## Selection rules
`Blocker Task: NONE` means no task dependencies, not READY/ACTIVE. All blocker IDs must be COMPLETE with verified evidence before READY. Human approval and explicit activation in handoff/queue are additional gates. Select one eligible task; lowest ID breaks ties. Never auto-activate or execute a successor. PENDING dependency wait is not BLOCKED.

## Discovery packets — fixed bounded outcomes

### FMT-001
Task Name: Reconcile existing Toolkit source and invalid OTR claims
Blocker Task: NONE
Status: PENDING | Assigned Agent: OpenCode | Stage: DISCOVERY
Objective: source-anchored inventory of existing Toolkit entry, imports, views and actual code reuse; identify inaccurate OTR claims without rewriting historical evidence.
Read: `OSJS/src/packages/MCSModbusToolkit/`, `workflow/active_work/evidence/otr-002a-report.md`, `workflow/active_work/evidence/otr-003a-map.md`.
Write: one new report under `planning/microtask/evidence/` only, exact filename set on promotion. No source edits.
Acceptance: source path/line citations for current modules; existing versus missing versus unknown distinguished; each disputed old claim checked against source.
Evidence: pinned SHA, file/line matrix and contradiction log. Size target 0–3; split if inventory grows beyond one bounded working set.

### FMT-002
Task Name: Inventory Electron Toolkit screens and assets
Blocker Task: NONE
Status: PENDING | Assigned Agent: OpenCode | Stage: DISCOVERY
Objective: enumerate actual Electron shell, navigation, screen entry points, UI components and reusable styles/assets.
Read: `electron/` UI entry, HTML/CSS/JS and related manifest, discovered from actual tree. Write: one source-anchored inventory report only.
Acceptance: real file paths and view list; UI versus Electron-only dependencies classified; unknown rendered details marked unknown.
Evidence: pinned SHA, paths/lines and screenshot availability (do not invent screenshots). If more than three distinct views require deep investigation, create separate review packets instead of bloating this task.

### FMT-003
Task Name: Inventory existing OS.js Toolkit UI and launcher
Blocker Task: NONE
Status: PENDING | Assigned Agent: OpenCode | Stage: DISCOVERY
Objective: verify shell, editors, diagnostics, launcher and desktop icon implementation and discover actual build/test commands.
Read: Toolkit package and OS.js shell/launcher registration files discovered from source. Write: one report only.
Acceptance: source-linked entry/navigation/view map; actual icon/launcher behavior distinguished from intent; exact existing build/test commands identified without executing live changes.
Evidence: pinned SHA, paths/lines, known/unknown matrix.

### FMT-004
Task Name: Inventory Linux backend operations and safety boundaries
Blocker Task: NONE
Status: PENDING | Assigned Agent: OpenCode | Stage: DISCOVERY
Objective: identify actual per-runtime read, validate, write, restart and observed-status paths without inventing APIs.
Read: Toolkit transports/contracts/server, MMA2, simulator/replicator backend and relevant docs; split per runtime if too large. Write: one bounded report only.
Acceptance: operation-to-source mapping; distinguish existing implementation from merely declared client contract; mark authorization, fixtures, side effects and unknowns.
Evidence: pinned SHA, exact endpoint or IPC paths only when observed, unresolved contract questions. No service/device access.

### FMT-005
Task Name: Verify existing Toolkit build and test baseline
Blocker Task: FMT-001, FMT-003
Status: PENDING | Assigned Agent: OpenCode | Stage: DISCOVERY
Objective: identify and run ONLY explicitly approved non-destructive local static/build checks on pinned source, without code repair.
Read: package scripts/build config and current source. Write: exact report only, no product edits or generated tracked artifacts.
Acceptance: record exact command, environment, exit and output for each approved check; no generic syntax check misrepresented as full build; failures remain failures.
Evidence: raw transcript and git status before/after. If execution environment/commands cannot be authorized, report BLOCKED, do not improvise.

### FMT-006
Task Name: Map Electron-to-OS.js visual and interaction gaps
Blocker Task: FMT-001, FMT-002, FMT-003
Status: PENDING | Assigned Agent: OpenCode | Stage: DESIGN
Objective: map each source-verified Electron view to existing OS.js view and isolate only actual missing visual/interaction behavior.
Read: accepted inventories and actual UI files. Write: one parity matrix report only.
Acceptance: exact source/target paths and copy/adapt/no-change decision; distinguish source comparison from rendered verification; each gap has a proposed <=3-file CODE slice or requires separate discovery.
Evidence: pinned SHA and per-view gap matrix. Do not claim screenshot parity without screenshots.

### FMT-007
Task Name: Map backend contract gaps by runtime
Blocker Task: FMT-001, FMT-004
Status: PENDING | Assigned Agent: OpenCode | Stage: DESIGN
Objective: reconcile UI calls, transport, server and runtime implementation for Simulator, Replicator and MMA; identify unsupported actions.
Read: accepted inventories and actual code. Write: one contract matrix report only.
Acceptance: per-operation implementation/evidence; safe fixture and permissions identified; unresolved restart/status/write paths remain explicit blockers to implementation.
Evidence: pinned SHA, source-linked matrix. Split per runtime if task exceeds size gate.

### FMT-008
Task Name: Generate source-pinned implementation microtasks
Blocker Task: FMT-005, FMT-006, FMT-007
Status: PENDING | Assigned Agent: ChatGPT | Stage: DESIGN
Objective: replace speculative broad OTR parent tasks with only source-proven gap packets; omit already satisfied features.
Read: accepted FMT reports and microtask rules. Write: individual PLANNING task files and a dependency index only; no activation.
Acceptance: each packet has Task Name, Blocker Task, owner, exact <=3 production files plus tests, objective, scope, <=3 acceptance outcomes, evidence, delivery and five-dimension size; no cycles; distinct CODE/TEST/VERIFY gates.
Evidence: cross-referenced packet index and proof every task corresponds to an observed gap. Human reviews before promotion.

## Conditional implementation families — NOT executable tasks or asserted gaps
After FMT-008, generate ONLY required source-pinned children for: shell/navigation/assets; one Simulator view at a time; one Replicator view at a time; one MMA view at a time; per-runtime config read; validation and unknown-field roundtrip; disposable safe save/conflict handling; scoped restart only if supported; independent applied-state/status; one launcher/icon acceptance. For each actual CODE slice create independent TEST and rendered/runtime VERIFY tasks as applicable. Blocker Task points to exact accepted prerequisite IDs, never an entire phase or numeric predecessor. Do not invent tab count, endpoint, test command, permission or production file path before discovery.

## Final independent acceptance candidates (materialize after implementation is known)
- One shell/launcher parity VERIFY preserving OS.js Start menu/taskbar and single icon.
- One rendered parity VERIFY per actual Simulator, Replicator and MMA screen/state.
- One safe disposable end-to-end VERIFY per supported runtime: load/edit/validate/save/(authorized supported restart)/independent applied-state readback; split invalid/discard/conflict/failure scenarios into separate packets.
- Independent OS.js build/deployment verification; Electron remains separately deployable. No production/operator data.

## Review decision requested
Approve or edit FMT-001..008 as a planning backlog; only then create exact standalone packets and activate ONE blocker-free task separately. Until then root handoff NONE and retired OTR queue remain unchanged.
