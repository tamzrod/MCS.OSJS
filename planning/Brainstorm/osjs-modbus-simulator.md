# OS.js Modbus Simulator

## Idea

Build a deliberately simple Modbus device simulator as an OS.js application using the existing MMA2 component as the Modbus memory/runtime.

The OS.js application is the simulator configuration and control surface. It does not implement another Modbus server.

## Basic Configuration

A simulated device configuration defines only the address ranges required for the four read function-code areas:

- FC1 — coils: start + count
- FC2 — discrete inputs: start + count
- FC3 — holding registers: start + count
- FC4 — input registers: start + count

Device/listener identity required by MMA2, such as listener/port and Unit ID, is also part of the generated runtime configuration where required.

No per-address value configuration is required for the initial simulator.

## Runtime Model

```text
OS.js Modbus Simulator
→ edit simulator/device configuration
→ generate/update MMA2 configuration
→ validate configuration
→ restart MMA2 with the new configuration
→ MMA2 exposes the configured Modbus address space
```

MMA2 is restarted when the structural simulator configuration changes. It is not restarted for periodic value changes.

## Simulated Values

All configured FC1, FC2, FC3, and FC4 ranges are populated with random values.

Every minute:

```text
GENERATE RANDOM VALUES
→ FC1 boolean values
→ FC2 boolean values
→ FC3 uint16 values
→ FC4 uint16 values
→ WRITE ALL VALUES THROUGH MMA2 RAW INGEST
```

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
= periodic random-value generation

MMA2 raw ingest
= simulator/producer value injection path

MMA2 Modbus TCP
= external client-facing protocol path
```

## Initial Simulator UI Plan

Keep the UI intentionally small. The first version is a device-definition editor and runtime-status surface, not a live SCADA-style value editor.

### Main Window

Use one OS.js application window with a simple two-pane layout:

```text
+---------------------------------------------------------------+
| Modbus Simulator                              [Add Device]     |
+----------------------+----------------------------------------+
| Simulated Devices    | Selected Device                        |
|                      |                                        |
| Device 1   RUNNING   | Name      [____________________]       |
| Device 2   STOPPED   | Port      [_____]                      |
| Device 3   ERROR     | Unit ID   [___]                        |
|                      |                                        |
|                      | FC1 Coils                              |
|                      | Start [_____]   Count [_____]          |
|                      |                                        |
|                      | FC2 Discrete Inputs                    |
|                      | Start [_____]   Count [_____]          |
|                      |                                        |
|                      | FC3 Holding Registers                  |
|                      | Start [_____]   Count [_____]          |
|                      |                                        |
|                      | FC4 Input Registers                    |
|                      | Start [_____]   Count [_____]          |
|                      |                                        |
|                      | Random update: every 60 seconds        |
|                      |                                        |
|                      | [Delete]              [Save / Apply]   |
+----------------------+----------------------------------------+
```

### Device List

The left pane represents simulator-owned device definitions only.

Each row should show enough information to identify runtime state without exposing implementation detail:

- device name;
- listener port;
- Unit ID;
- status: `RUNNING`, `STOPPED`, or `ERROR`.

Selecting a row loads that device into the editor on the right.

`Add Device` creates a new unsaved simulator definition. `Delete` removes only the selected simulator-owned definition and must not remove configuration owned by Replicator or another MMA2 producer.

### Device Editor

The first version needs only:

- Name;
- listener Port;
- Unit ID;
- FC1 Start + Count;
- FC2 Start + Count;
- FC3 Start + Count;
- FC4 Start + Count.

No individual address/value editor is required.

Zero count may represent an unused function-code area if supported by the eventual simulator/config schema; this must be validated rather than assumed during implementation.

### Save / Apply Behavior

`Save / Apply` is one deliberate structural operation:

```text
USER EDITS DEVICE
→ VALIDATE FORM
→ SUBMIT SIMULATOR CONFIGURATION INTENT
→ SHARED MMA2 CONFIG AUTHORITY CHECKS CONFLICTS
→ REJECT WITH EXPLICIT ERROR
   OR
→ GENERATE/UPDATE EFFECTIVE MMA2 CONFIG
→ RESTART MMA2 IF STRUCTURAL CONFIG CHANGED
→ VERIFY RUNTIME
→ UPDATE DEVICE STATUS
```

The browser/UI must not directly overwrite the effective MMA2 configuration file.

Validation should surface at least:

- invalid/missing port;
- invalid Unit ID;
- invalid start/count ranges;
- listener conflict;
- Unit/range ownership conflict with another MMA2 producer;
- MMA2 configuration/start failure.

A rejected save leaves the previously active configuration intact.

### Runtime Information

The first UI does not need a live register table. It only needs enough runtime feedback to prove the simulator is operating:

```text
Status: RUNNING
Port: 1502
Unit ID: 1
Random update: every 60 seconds
Last random update: <timestamp when available>
```

The periodic randomizer remains automatic and fixed at one-minute intervals in the initial version. There is no per-device waveform/randomization editor.

### UI Scope Boundary

Initial UI includes:

- list simulated devices;
- add/select/edit/delete simulator-owned device definitions;
- configure port, Unit ID, and FC1-FC4 start/count;
- save/apply structural configuration;
- show validation/configuration conflicts;
- show basic runtime state and last random update when available.

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

- configure ranges;
- generate/submit its MMA2 configuration requirements without taking exclusive ownership of shared MMA2 configuration;
- restart MMA2 after structural config changes through the eventual configuration/lifecycle authority;
- populate all configured FC1-FC4 ranges with random values every minute through raw ingest;
- allow real external Modbus clients to read the simulated device.

Advanced signal behavior, scripting, ramps, sine waves, per-address manual configuration, and similar simulation features are not part of this initial idea.
