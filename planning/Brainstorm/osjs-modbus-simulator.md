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

## Ownership Boundary

```text
OS.js Simulator App
= simulator configuration / control surface

MMA2 configuration
= Modbus address space, listener and Unit-ID runtime structure

Simulator randomizer
= periodic random-value generation

MMA2 raw ingest
= simulator value injection path

MMA2 Modbus TCP
= external client-facing protocol path
```

## Initial Intent

Keep the first simulator intentionally small:

- configure ranges;
- generate MMA2 config;
- restart MMA2 after structural config changes;
- populate all configured FC1-FC4 ranges with random values every minute through raw ingest;
- allow real external Modbus clients to read the simulated device.

Advanced signal behavior, scripting, ramps, sine waves, per-address manual configuration, and similar simulation features are not part of this initial idea.
