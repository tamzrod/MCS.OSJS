# OS.js Modbus Simulator

## Idea

Build a deliberately simple Modbus device simulator as an OS.js application using the existing MMA2 component as the Modbus memory/runtime.

The OS.js application is the simulator configuration and control surface. It does not implement another Modbus server.

## Basic Configuration

A simulated device configuration defines the address ranges and randomization interval required for the four read function-code areas:

- FC1 — coils: start + count + randomize every (ms)
- FC2 — discrete inputs: start + count + randomize every (ms)
- FC3 — holding registers: start + count + randomize every (ms)
- FC4 — input registers: start + count + randomize every (ms)

Device/listener identity required by MMA2, such as listener/port and Unit ID, is also part of the generated runtime configuration where required.

No per-address value configuration is required for the initial simulator.

## Runtime Model

```text
OS.js Modbus Simulator
→ edit simulator/device configuration
→ generate/update MMA2 configuration requirements
→ validate configuration
→ restart MMA2 when structural MMA2 configuration changes
→ MMA2 exposes the configured Modbus address space
```

MMA2 is not restarted for periodic value changes. A randomization-interval-only change should update the simulator scheduler without requiring an MMA2 restart where the runtime design permits it.

## Simulated Values

All configured FC1, FC2, FC3, and FC4 ranges are populated with random values on their independently configured schedules.

```text
FC1 timer ─┐
FC2 timer ─┼→ GENERATE VALUES → MMA2 RAW INGEST
FC3 timer ─┤
FC4 timer ─┘
```

- FC1 produces random boolean values.
- FC2 produces random boolean values.
- FC3 produces random uint16 values.
- FC4 produces random uint16 values.

All simulator-originated writes use MMA2 raw ingest. The simulator must not add a direct-memory bypass or use Modbus writes merely to populate its own simulated values.

## Shared MMA2 Configuration Boundary

MMA2 is a shared engine. The simulator is only one future user; Replicator will use the same MMA2 engine. Therefore the simulator must not become the exclusive owner of MMA2 configuration or overwrite configuration belonging to another producer.

MMA2 remains a configuration consumer and runtime engine. Ownership/conflict policy belongs outside the MMA2 source tree.

```text
Simulator intent ──┐
                   ├── shared configuration authority
Replicator intent ─┘            ↓
                       effective MMA2 config
                                ↓
                              MMA2
```

The effective configuration must be composed/validated before activation so multiple MMA2 users cannot independently claim conflicting listeners, Unit IDs, function-code ranges, or destination addresses.

A hard rule for the eventual shared configuration model is that two producers must not independently own the same MMA2 destination address.

The exact configuration-authority implementation is not decided by this brainstorm.

## Persistent Configuration Location

Persistent runtime configuration must live on the host-mounted MCS.OSJS data/configuration location, not inside the packaged application source tree or MMA2 source directory.

Conceptually:

```text
HOST
└── <MCS.OSJS mounted location>/
    └── config/
        ├── simulator/
        ├── replicator/
        └── mma2/          # effective/generated runtime configuration as appropriate
             │
             └── mounted into the MCS.OSJS container
```

The exact host path and final directory names must follow the appliance mount convention once that convention is established/verified. This brainstorm does not invent a host path.

Rule to cement during implementation:

> No application writes persistent configuration into its packaged application directory. Persistent configuration belongs under the designated host-mounted configuration root.

Simulator and Replicator may own their respective configuration intent, while MMA2 consumes the validated effective runtime configuration generated from those intents.

## Ownership Boundary

```text
OS.js Simulator App
= simulator configuration / control surface

Shared configuration authority
= validates ownership, prevents conflicts, composes effective MMA2 configuration

MMA2 effective configuration
= Modbus address space, listener and Unit-ID runtime structure

Simulator randomizer
= independent FC1-FC4 periodic random-value generation

MMA2 raw ingest
= simulator/producer value injection path

MMA2 Modbus TCP
= external client-facing protocol path
```

## Approved Simulator UI

The approved initial UI is a focused device-definition editor and runtime-status surface. It is not a live SCADA-style value editor.

### Main Window

```text
┌──────────────────────────────────────────────────────────────────────────────────────────────┐
│ Modbus Simulator                                                                        _ □ X │
├──────────────────────────────────────────────────────────────────────────────────────────────┤
│ [ + Add Device ]   [ Duplicate ]   [ Delete ]                 [ Save & Apply ] [ Discard ] │
├──────────────────────┬───────────────────────────────────────────────────────────────────────┤
│ Devices              │ Configuration                                                         │
│                      │                                                                       │
│ [ Search devices... ]│ Device Settings                                                       │
│                      │                                                                       │
│ ● Sim-PLC-1          │ Name        [ Sim-PLC-1                                          ]    │
│   Port: 5020         │                                                                       │
│   Unit: 1            │ Listen Port [ 5020 ]          Unit ID [ 1 ]                           │
│                      │                                                                       │
│ ● Sim-IO-1           │ Enabled     [ ON ]                                                   │
│   Port: 5020         │                                                                       │
│   Unit: 2            ├───────────────────────────────────────────────────────────────────────┤
│                      │ Function Code Ranges                                                   │
│ ● Sim-Relay          │                                                                       │
│   Port: 5021         │ Function              Start    Count    Randomize Every (ms)           │
│   Unit: 1            │ FC1 - Coils           [ 0 ]    [ 64 ]   [ 1000  ]                    │
│                      │ FC2 - Discrete Inputs [ 0 ]    [ 64 ]   [ 5000  ]                    │
│ ○ Sim-Meter          │ FC3 - Holding Regs    [ 0 ]    [100 ]   [ 60000 ]                    │
│   Port: 5022         │ FC4 - Input Regs      [ 0 ]    [100 ]   [ 10000 ]                    │
│   Unit: 10           │                                                                       │
│                      │ Calculated ranges:                                                     │
│                      │ FC1 0–63   FC2 0–63   FC3 0–99   FC4 0–99                            │
│                      │                                                                       │
│                      │ All generated values are written through MMA2 raw ingest.              │
├──────────────────────┼───────────────────────────────────────────────────────────────────────┤
│                      │ Runtime                                                               │
│                      │ MMA2 Status      ● Running                                             │
│                      │ Raw Ingest       ● OK                                                  │
│                      │                                                                       │
│                      │ Last Random Update                                                      │
│                      │ FC1   <timestamp>        next in <time>                                │
│                      │ FC2   <timestamp>        next in <time>                                │
│                      │ FC3   <timestamp>        next in <time>                                │
│                      │ FC4   <timestamp>        next in <time>                                │
│                      │                                                                       │
│                      │ Total Points: <count>                                                  │
├──────────────────────┴───────────────────────────────────────────────────────────────────────┤
│ MMA2: RUNNING     Raw Ingest: OK     Selected Device: Sim-PLC-1                              │
└──────────────────────────────────────────────────────────────────────────────────────────────┘
```

### Device List

The left pane represents simulator-owned device definitions only.

Each row shows:

- device name;
- listener port;
- Unit ID;
- status: `RUNNING`, `STOPPED`, or `ERROR`.

Selecting a row loads that device into the editor on the right.

`Add Device` creates a new unsaved simulator definition. `Duplicate` creates a new simulator-owned definition from the selected device but must still pass conflict validation before activation. `Delete` removes only the selected simulator-owned definition and must not remove configuration owned by Replicator or another MMA2 producer.

### Device Editor

The approved first version exposes:

- Name;
- listener Port;
- Unit ID;
- Enabled;
- FC1 Start + Count + Randomize Every (ms);
- FC2 Start + Count + Randomize Every (ms);
- FC3 Start + Count + Randomize Every (ms);
- FC4 Start + Count + Randomize Every (ms).

`Randomize Every` uses a fixed millisecond unit. There is no unit selector. Each FC has its own independent randomization interval.

No individual address/value editor is required.

Zero count may represent an unused function-code area if supported by the eventual simulator/config schema; this must be validated rather than assumed during implementation.

`Randomize Every (ms)` must be validated as a positive integer for an enabled/configured FC. Exact minimum/maximum scheduler limits remain an implementation decision unless established elsewhere.

### Save / Apply Behavior

`Save / Apply` is one deliberate operation:

```text
USER EDITS DEVICE
→ VALIDATE FORM
→ SUBMIT SIMULATOR CONFIGURATION INTENT
→ SHARED MMA2 CONFIG AUTHORITY CHECKS CONFLICTS
→ REJECT WITH EXPLICIT ERROR
   OR
→ GENERATE/UPDATE EFFECTIVE MMA2 CONFIG IF STRUCTURE CHANGED
→ RESTART MMA2 IF STRUCTURAL MMA2 CONFIG CHANGED
→ UPDATE RANDOMIZER SCHEDULES
→ VERIFY RUNTIME
→ UPDATE DEVICE STATUS
```

The browser/UI must not directly overwrite the effective MMA2 configuration file.

Validation should surface at least:

- invalid/missing port;
- invalid Unit ID;
- invalid start/count ranges;
- invalid randomization interval;
- listener conflict;
- Unit/range ownership conflict with another MMA2 producer;
- MMA2 configuration/start failure.

A rejected save leaves the previously active configuration intact.

Structural changes such as Port, Unit ID, Start, or Count may require MMA2 restart. A change only to `Randomize Every (ms)` belongs to simulator scheduling and should not by itself require an MMA2 restart where the runtime design permits scheduler-only reload.

### Runtime Information

The first UI does not need a live register table. It shows enough feedback to prove the simulator and each FC scheduler are operating:

```text
MMA2 Status: RUNNING
Raw Ingest: OK

FC1 Last Random Update: <timestamp>   Next: <time>
FC2 Last Random Update: <timestamp>   Next: <time>
FC3 Last Random Update: <timestamp>   Next: <time>
FC4 Last Random Update: <timestamp>   Next: <time>

Total Points: <count>
```

Per-FC timing is intentional: FC1, FC2, FC3, and FC4 may all update at different millisecond intervals.

### UI Scope Boundary

Initial UI includes:

- list/search simulated devices;
- add/duplicate/select/edit/delete simulator-owned device definitions;
- configure Name, Port, Unit ID, and Enabled state;
- configure FC1-FC4 start/count;
- configure an independent `Randomize Every (ms)` for every FC;
- save/apply configuration;
- show validation/configuration conflicts;
- show MMA2/raw-ingest runtime state;
- show per-FC last/next random update state.

Initial UI excludes:

- raw YAML editing;
- direct MMA2 configuration editing;
- individual register/coil value editing;
- live memory tables;
- charts;
- waveform/ramp/sine/script configuration;
- Replicator configuration;
- arbitrary MMA2 lifecycle controls unrelated to simulator-owned definitions.

## Initial Intent

Keep the first simulator intentionally small:

- configure ranges and independent per-FC randomization intervals;
- generate/submit its MMA2 configuration requirements without taking exclusive ownership of shared MMA2 configuration;
- restart MMA2 only when structural MMA2 configuration changes through the eventual configuration/lifecycle authority;
- populate configured FC1-FC4 ranges on their own schedules through raw ingest;
- allow real external Modbus clients to read the simulated device.

Advanced signal behavior, scripting, ramps, sine waves, per-address manual configuration, and similar simulation features are not part of this initial idea.
