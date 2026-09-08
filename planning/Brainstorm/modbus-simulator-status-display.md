# Brainstorm: Modbus Simulator - MMA2 Status + Simulator Status Display

Status: brainstorm material only. Non-authoritative. No implementation authorized until human promotion.

Date: 2026-09-08. Owner: human. JR captures the idea here.



## Intent

Display **MMA2 status** and **Simulator status** inside the Modbus Simulator app, so an operator can see in-window whether MMA2 is reachable and whether the Simulator is producing values through Raw Ingest. Nothing beyond the existing runtime status model.is proposed here: verbatim status truth is already produced by `SchedulerApplier.RuntimeStatus`, and the SIM-018/019/020 pipeline already carries version-1 `load`, `apply`, `status` operations over the authenticated WebSocket -> Unix-domain socket relay. No new transport, no new Go API, and no new HTTP/TCP exposure is suggested.



## What We Are NOT Deciding Here

- **Placement is UNDECIDED**. That is the core open question of this brainstorm.
- Exact visual design/skinning upstream of minimal placement.is undecided.
Whether status belongs ina bottom bar, sidebar footer, editor header, per-device row, a collapsible panel, or another reserved zone.is undecided.



## Candidate Placements (minimalist preferred)

| Candidate | Shape | Minimal points | Open concerns |
| --- | --- | --- | --- |
| A. Bottom status bar (extend existing `.sim-status`) | One compact row of labeled dot/state pairs (MMA2;, Sim;, Ingest;, Last apply) | Very minimal; already the established message strip | Bar is currently a transient message zone; mixing steady-state status may crowd/confuse message text; needs budget for compactness |
| B. Device-list row chips | Per-device dot + state (RUNNING/STOPPED/ERROR( alongside Port/Unit in the sidebar rows | Tight contextual per-selected-device state; no new chrome | MMA2 status is process-wide anyway (one MMA2 per host?); duplicating it in every row may be repetitive; raw-ingest/per-FC detail has no natural per-row home |
| C. Editor-pane status strip (top of device editor) | Small line above the definition form; only for theselected device; shows MMA2 + Simulator + Raw Ingest + FC1-FC4 last/next compact | Per-selected-device FC timing naturally fits beside the form; MMA2 state visible when editing that device | Only visible when a device is selected; (MMA2 global status may warrant visibility before selection) |
| D. Collapsible "Runtime status" section in editor pane | One small heading that expands to show the SIM-021 surface (MMA2, Simulator, Raw Ingest, per-FC Last/Next, total points, last apply) | Detail-rich but hidden until needed (minimalist idle view) | Extra click; where to place the summary dot when collapsed?; may overlap candidate A |
| E. Dedicated compact status header row across the whole window (new top strip) | MMA2 + Simulator (selected device) always visible at a glance; steady-state, not transient | New chrome row costs window height; must stay slim; relationship to message bar must be defined (message strip continues bottom; status strip top?) |

Lean: candidates A, C, and E feel most minimalist; C and D hold the per-FC detail, which has no home in A or B. Key tension: MMA2 status is global (one host/process(, while Simulator state (and especially per-FC timing) is per-device. Mixing both into one place risks confusing global-vs-selected scoping; placement should make that scoping explicit or separate it.



## Status Content (data already available; no new computation proposed)

- MMA2 status: `RUNNING` / `RESTARTING`/`WAITING` / `STOPPED` / `ERROR` (truthful readiness evidence from `RuntimeStatus`; an enabled-but-unready device never displays RUNNING)..
- Simulator device status: `RUNNING` / `STOPPED` / `ERROR` (armed scheduling gate (SIM-015); disabled devices report STOPPED/inactive; raw-ingest faults surface ERROR with the raw error string)..
- Raw Ingest state: `OK` / `WAITING` / `ERROR`, with last error text when present..
- Per-FC (FC1-FC4): configured point counts, `last` successful update timestamp, `next` scheduled timestamp; unused (zero-count) FC displays inactive..
- Total points (exact configured point total); last apply outcome/time (SIM-020 already presents the apply message + time in the bar)...



## Open Questions

1. **Placement**: which candidate (or combination) will be chosen? User preference: minimalism .
2. **Scope**: MMA2 status is global; Simulator/Raw-Ingest/per-FC status is per-device. Should the UI display both scopes at once, e.g. an always-visible global strip plus per-selected-device detail, or hide global status until a device is selected?
3. **Refresh cadence**: 1s polling (as SIM-007 did), event-push, or refresh-on-action (load/save(? Minimal polling suggests 1s client-side, dependent on the existing `status` RPC...
4. **Steady vs transient messaging**: the existing `.sim-status` bar doubles as a transient message zone. Do we split (steady bottom strip + transient toast/message(, or keep one hybrid row)?
5. **Detail depth**: always-visible minimal dots(state words or colored dots( with FC timing only in a collapsed/secondary zone, or everything flat?
6. **Zero-device / degraded states**: what shows when no device exists, when the runtime socket is down, or during MMA2 restart/wait transitions?(SIM-021 acceptance already demands truthful transitions; brainstorm must pick how much to surface for any given state..



## Constraints to Preserve (from established truth)

- MMA2 lifecycle is independent; Simulator never START/STOP/SPAWN/KILL/REPLACE it; only RESTART after a committed valid structural change; values enter via Raw Ingest only..
- No new Simulator HTTP route, TCP listener, or config/control API. Status must flow through the SIM-018/019/020 local boundary (authenticated OS.js session WebSocket -> allowlisted relay -> Unix-domain socket -> runtime `status` RPC)..
- Runtime status gate (SIM-009/SIM-015): an enabled-but-unready device never claims RUNNING; failures remain visible until a later success..
- No register-value viewer, charts, history, or waveform editor in the initial UI (SIM-005 approved UI boundary, still in force until human changes it)...



## Relationship to Retired SIM-021

`workflow/archive/sim-021-restore-truthful-runtime-status.md` was retired (not completed) when the user cleared active work. This brainstorm reopens that question, but deliberately leaves placement undecided and UI-scope open, per user request: minimal display of MMA2 status und Simulator status into the Modbus Simulator app, placement TBD.. Once the human decides placement, decompose microtask(s) and promote normally..
