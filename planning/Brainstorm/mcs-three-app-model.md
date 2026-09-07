# MCS Three-App MMA2 Usage Model

## Idea

MCS presents three separate OS.js applications/windows that use MMA2 for three different purposes. MMA2 is the common Modbus memory engine; it is not itself the Simulator or Replicator.

## Application Responsibilities

### Simulator

The Simulator opens an MMA2-backed Modbus memory space and loads generated/simulated values into it.

```text
simulated values
      ↓
  Simulator
      ↓
    MMA2
      ↓
 Modbus clients
```

Simulation behavior belongs to the Simulator. MMA2 provides the Modbus memory/runtime used to expose those values.

### Replicator

The Replicator reads another Modbus source/device and pushes the acquired data into an MMA2-backed memory space.

```text
external Modbus
      ↓
  Replicator
      ↓
    MMA2
      ↓
 Modbus clients
```

Acquisition and replication behavior belong to the Replicator. MMA2 remains the destination memory/runtime.

### Memory Appliance

The Memory Appliance opens/loads an MMA2 Modbus memory space for direct third-party use. It does not require simulation or replication behavior.

```text
third-party systems
        ↕
      MMA2
```

In this mode, the Modbus memory service itself is the product exposed by the application.

## Architectural Boundary

```text
Simulator       = generate simulated values → MMA2
Replicator      = read external Modbus → push data → MMA2
Memory Appliance = provide MMA2 memory directly for third-party use

MMA2 = common deterministic Modbus memory/runtime engine
```

The three applications should remain separate OS.js windows/apps because they represent separate user intents and lifecycle responsibilities even though all three rely on the same MMA2 machinery.

MMA2 should remain neutral about where values originate. Simulator, Replicator, and Memory Appliance own their respective intent and lifecycle; MMA2 owns the common Modbus memory behavior and exposure.

## Shared Resource Implication

Because multiple MCS applications may use MMA2 concurrently, shared MMA2 configuration must prevent one application from unintentionally overwriting or deleting resources owned by another. The existing `(port, unit_id)` ownership work is part of preserving this separation.

This brainstorm records the product/application model only. It does not by itself authorize implementation or change current Active Work sequencing.
