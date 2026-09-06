# Brainstorm Topics

Baseline commit: 500376cfb5c222298aadfcf035aad0af0a635773
Working tree: clean
Source dependencies: planning/Brainstorm/mma2-basic-install-test.md, planning/Brainstorm/osjs-modbus-simulator.md
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
