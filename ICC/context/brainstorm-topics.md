# Brainstorm Topics

Baseline commit: a17191c; working tree: dirty (planning/microtask/mma2-basic-install-test.md and planning/microtask/osjs-base-webapp-init.md deleted uncommitted by human, outside this branch.
Audited Overlay: handoff.md; five sim-021 planning microtasks promoted into workflow/active_work as SIM-021A..SIM-021E (ordered; ICC nodes refreshed herein.
Source dependencies: planning/Brainstorm/mma2-basic-install-test.md, planning/Brainstorm/osjs-modbus-simulator.md, planning/Brainstorm/mcs-three-app-model.md, planning/Brainstorm/modbus-simulator-status-display.md
Parent: planning-workflow
Zoom In:(none; leaf node(
Zoom Out: planning-workflow

## Topic: MMA2 Basic Install and Test (brainstorm, non-authoritative

Intent:  bring MMA2 into MCS.OSJS as the next standalone component and prove basic MMA2 runtime works before any Replicator/Orchestrator/OS.js integration. Candidate outcome:  local checkout contains runnable MMA2 component sourced from `tamzrod/mma2`, passing one end-to-end Modbus memory test (BUILD -> START -> VERIFY LISTENER -> WRITE -> READ MATCHES -> RESTART -> LISTENER RETURNS(. Donor facts:  Go project with go.mod, cmd/mma2/main.go, Dockerfile, tests, existing binary; architecture defines config as startup-loaded+immutable, configured ports/Unit IDs/memory sizes feeding core memory+transport/runtime layers. Non-scope:  no Replicator/Orchestrator/OS.js integration, no multi-port/Unit-ID stress, no persistence redesign, no protocol expansion. Open questions (resolved only when converting to microtasks(:  which MCS.OSJS dir owns component; go run vs built binary vs donor Dockerfile; minimal config file; which Modbus client/tool does the write/read. Decomposed into microtask MMA2-001/002; MMA2-001 later split in active execution (import+build vs runtime smoke-test(, with MMA2-002 in active_work preserving internal order+completion state as repository truth..

## Topic: OS.js Modbus Simulator (brainstorm, APPROVED as source of active SIM program

Intent:  build a deliberately simple Modbus device simulator as an OS.js application using existing MMA2 component as Modbus memory/runtime. The OS.js application is the simulator configuration/control surface; it does not implement another Modbus server..

### Two Parameter Domains

Simulator device has two distinct parameter domains (different owners + reload behavior(:

- MMA2 parameters:  Listen Port, Unit ID, FC1-FC4 Start/Count (structural; consumed by shared MMA2 config authority(;
- Random runtime parameters:  FC1-FC4 Randomize Every (ms( (simulator-owned; consumed by simulator scheduler(.

May be persisted together as one simulator-owned device definition, but operationally separate parameter sets with separate consumers. Conceptual shape:  `device:{mma2:{port,unit_id,fc1..fc4:{start,count}}},random_runtime:{fc1_interval_ms,fc2_interval_ms,fc3_interval_ms,fc4_interval_ms}}`. Not a final serialization decision..

### Control / Data Paths

- Configuration/control path:  form fill -> parameter check (invalid->validation error( -> classify change (MMA2 structural vs random-runtime-only( -> validate shared MMA2 namespace / save+update simulator runtime schedule( -> MMA2 restart required only for structural changes; timing-only change never restarts MMA2(.
- Simulation data path:  random runtime -> generate values per FC schedule -> convert to raw ingest -> MMA2 raw ingest -> MMA2 memory. Runtime data path does not modify MMA2 config during normal randomization..

### Value Rules

FC1/FC2 random booleans; FC3/FC4 random uint16. All simulator-originated writes use MMA2 raw ingest ONLY; no direct-memory bypass, no Modbus-write population of own values. Per-FC independent timer. `Randomize Every (ms)` fixed ms unit, no unit selector; positive integer for enabled/configured FC; zero count may represent unused FC area if supported by eventual schema (must be validated, not assumed(..

### Shared MMA2 Configuration Boundary

MMA2 is a shared engine (simulator now, Replicator later(. Simulator must not become exclusive owner of MMA2 configuration nor overwrite another producer's config. Effective configuration must be composed/validated before activation so multiple users cannot independently claim conflicting listeners/Unit IDs/ranges/destination addresses. Hard rule:  two producers must not independently own the same MMA2 destination address. Random-runtime parameters are simulator-owned, not part of Replicator config or shared MMA2 config authority. Exact config-authority implementation not decided by this brainstorm..

### Persistent Configuration Location

Persistent runtime configuration must live under the host-mounted MCS.OSJS data/configuration root, NOT inside packaged application source tree or MMA2 source directory. Exact host path/final directory names must follow the appliance mount convention once established/verified;this brainstorm does not invent a host path. Rule:   no application writes persistent config into its packaged application directory. Simulator and Replicator may own their respective MMA2 config intent; MMA2 consumes validated effective runtime config composed from those intents..

### Approved Simulator UI (SIM-005

A focused device-definition editor + runtime-status surface (NOT a live SCADA-style value editor(:  two-pane layout — device list/search (simulator-owned definitions only; rows show name, listen port, Unit ID, status RUNNING/STOPPED/ERROR( + editor (Name, Listen Port, Unit ID, Enabled, FC1-FC4 rows Start+Count+Randomize Every (ms(, calculated address ranges(; controls Add/Duplicate/Delete/Save & Apply/Discard; runtime status area (MMA2 Status, Raw Ingest, per-FC last/next random update, Total Points(; bottom status bar. `Save & Apply` validates both domains, classifies changed params, routes structural via shared MMA2 config authority (restart/reload only when required( and timing-only to scheduler (no MMA2 restart(; rejected save leaves prior active config intact. UI must not directly overwrite effective MMA2 config file. Initial UI excludes:  raw YAML editing, direct MMA2 config editing, individual register/coil editing, live memory tables, charts, waveform/ramp/sine/script config, Replicator config, arbitrary MMA2 lifecycle controls unrelated to simulator definitions. Zero-count FC area possible unused(function-code unused( if supported by eventual schema. Keep first simulator intentionally small; advanced signal behavior not part of initial idea..

## Topic: MCS Three-App MMA2 Usage Model (brainstorm, non-authoritative)

MCS exposes separate Simulator, Replicator, and Memory Appliance applications over one neutral MMA2 memory/runtime engine. Simulator generates values into MMA2; Replicator acquires external Modbus values and pushes them into MMA2; Memory Appliance exposes MMA2 directly for third-party use. Each application retains its own intent and lifecycle, while shared `(port, unit_id)` ownership prevents cross-application overwrite or deletion. The model does not authorize implementation. The immediate simulator lifecycle gap has been decomposed into planning-only `SIM-008` and `SIM-009` microtasks.

## Topic: Modbus Simulator Status Display (brainstorm, HUMAN-OWNED, placement UNDECIDED

Intent: display **MMA2 status** and **Simulator status** inside themodbus Simulator app, minimalism preferred. Placement inside the app is deliberately undecided;candidate zones are: bottom status bar, device-list row chips, editor-pane status strip, collapsible runtime-status section, or a dedicated top status header row. No decision yet;user will choosethe placement.

Relevant data already exists: `SchedulerApplier.RuntimeStatus` (MMA2 RUNNING/RESTARTING/WAITING/STOPPED/ERROR; Simulator device RUNNING/STOPPED/ERROR; Raw Ingest OK/WAITING/ERROR; per-FC last/next timing, configured point counts, total points, last apply outcome( and the SIM-018/019/020 pipeline already carries a versioned `status` RPC over the authenticated OS.js session WebSocket -> allowlisted relay -> Unix-domain socket. No new transport,Go API,HTTP route,or TCP listener is proposed.

Open questions: placement;(global MMA2 vs per-device Simulator scoping mix;; refresh cadence (1s polling vs event-push vs on-action(;; steady-state vs transient message bar split;; detail depth (dots vs full FC timing(;; zero-device/runtime-down/restart-wait display states.

Non-authoritativebeyond the promoted five. Placement fixed by SIM-021D: one compact row below the Device Definition heading, bound only to SIM-021C status data, RUNNING/WAITING/STOPPED/ERROR + no-device/unavailable states, no color-only cues, bottom .sim-status bar reserved for transient error/messages. SIM-021A..SIM-021E promoted into active work on 2026-09-08.
