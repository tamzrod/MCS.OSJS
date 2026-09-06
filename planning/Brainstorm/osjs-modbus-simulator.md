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

## Initial Intent

Keep the first simulator intentionally small:

- configure ranges;
- generate/submit its MMA2 configuration requirements without taking exclusive ownership of shared MMA2 configuration;
- restart MMA2 after structural config changes through the eventual configuration/lifecycle authority;
- populate all configured FC1-FC4 ranges with random values every minute through raw ingest;
- allow real external Modbus clients to read the simulated device.

Advanced signal behavior, scripting, ramps, sine waves, per-address manual configuration, and similar simulation features are not part of this initial idea.
