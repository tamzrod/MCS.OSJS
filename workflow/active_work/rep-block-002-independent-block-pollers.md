# REP-BLOCK-002 — Independent Pull Block Pollers

## Primary Outcome
Allow one Replicator device to contain multiple explicit Pull Blocks, with each block running its own independent poll interval while sharing the same source endpoint and source Unit ID.

## Scope
- Replace the current single `pull_block` field with an ordered collection of Pull Blocks.
- Each Pull Block owns its own FC, Start, Count, and Scan Rate.
- Pull Blocks must support all four basic Modbus read areas: FC1 Coils, FC2 Discrete Inputs, FC3 Holding Registers, and FC4 Input Registers.
- FC1/FC2 blocks replicate bit values into the corresponding destination Coil/Discrete Input memory; FC3/FC4 blocks replicate register values into the corresponding destination Holding/Input Register memory.
- Start one independent polling loop/ticker per enabled Pull Block.
- Each block writes only its corresponding destination area/range mapping.
- Keep device-level Endpoint, Source Unit ID, Enabled state, Name, and destination reservation shared by all blocks.
- Preserve deterministic block ordering and persistence.
- Migrate the current single `pull_block` document into a one-element block collection on load.
- Expose per-block runtime state required by the later tabbed UI task: block identity/index, running state, last poll, source status, and last error.
- The allocated destination memory must remain a normal externally readable Modbus memory. Third-party Modbus clients must be able to read the Replicator destination through the configured listener using the corresponding FC1/FC2/FC3/FC4 range.
- Device operational `Status` is end-to-end health, not MMA2 ownership. `OWNED` is ownership/configuration information only and is not an acceptable operational success status.
- Device `Status = OK` only after all three conditions are true: (1) destination MMA2 memory is allocated/ready and externally serveable over Modbus, (2) configured Pull Block poller(s) are successfully reading the source Modbus, and (3) the values read from the source are successfully written into the allocated MMA2 memory.
- A successful source read without a successful destination write must not report `OK`.
- Allocated/owned destination memory that cannot serve external Modbus reads must not report `OK`.
- Allocated/owned destination memory without successful source polling must not report `OK`.
- Keep `Owner: replicator` as separate ownership information; it must not be used as the device operational status.

## Non-Scope
- No UI tabs in this task.
- No per-block destination port/unit ownership.
- No multiple source endpoints inside one device.
- No retry/backoff redesign.
- No change to shared MMA2 `(port, unit_id)` ownership semantics.

## Acceptance Criteria
1. A Replicator device may persist and reload two or more Pull Blocks.
2. Each block polls at its own configured Scan Rate without changing another block's cadence.
3. FC/start/count used by a poll cycle come from that specific block only.
4. Pull Block FC selection accepts FC1, FC2, FC3, and FC4 and maps each to its matching MMA2 destination area.
5. FC1/FC2 bit blocks replicate and serve correct Coil/Discrete Input values; FC3/FC4 register blocks continue to replicate and serve correct Holding/Input Register values.
6. Legacy/current single-pull-block documents load as exactly one block without losing configuration.
7. Runtime status can report each block independently.
8. The Replicator destination listener serves the mirrored FC1/FC2/FC3/FC4 values to an external Modbus client at the configured `(port, unit_id)`.
9. `Owner: replicator` remains ownership metadata and is not presented as operational success.
10. Device `Status` reports `OK` only when destination memory is allocated/ready and serveable, source polling succeeds, and the polled values are successfully written to MMA2 memory.
11. If any required end-to-end stage fails — memory unavailable/not serveable, source read failure, or destination write failure — device `Status` must not report `OK` and runtime diagnostics must expose the failing stage/error.

## Verification
Focused Go tests using one device with multiple blocks at different scan rates and mixed FC1/FC2/FC3/FC4 types, proving independent cycle counts/status plus single-block migration and persistence round-trip. Verify each FC maps to the correct MMA2 area and the composed Replicator memory includes access policy allowing external FC1/FC2/FC3/FC4 reads. Human/JR end-to-end verification must use a real Modbus client against the Replicator destination `(port, unit_id)` and confirm Coil, Discrete Input, Holding Register, and Input Register values match the source after replication. Add end-to-end status tests proving `OK` only after a successful source-read-to-MMA2-write-and-serve cycle, plus negative cases for unavailable/unserveable memory, failed source read, and failed destination write. Rendered UI verification must confirm `Owner: replicator` remains separate while the displayed device `Status` follows the end-to-end semantics rather than showing `OWNED` as success.

## Dependencies
REP-BLOCK-001 must be complete/accepted first because this task extends its explicit Pull Block model.

## Sizing
Implementation surface 2; environment 0; behavior 3; verification 2; decision/recovery 1. Total 8 — task now includes mixed bit/register Pull Block behavior plus end-to-end serving/status verification and should be split before implementation if further expansion is added.
