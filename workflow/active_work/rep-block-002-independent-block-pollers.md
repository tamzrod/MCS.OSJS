# REP-BLOCK-002 — Independent Pull Block Pollers

## Primary Outcome
Allow one Replicator device to contain multiple explicit Pull Blocks, with each block running its own independent poll interval while sharing the same source endpoint and source Unit ID.

## Scope
- Replace the current single `pull_block` field with an ordered collection of Pull Blocks.
- Each Pull Block owns its own FC, Start, Count, and Scan Rate.
- Start one independent polling loop/ticker per enabled Pull Block.
- Each block writes only its corresponding destination area/range mapping.
- Keep device-level Endpoint, Source Unit ID, Enabled state, Name, and destination reservation shared by all blocks.
- Preserve deterministic block ordering and persistence.
- Migrate the current single `pull_block` document into a one-element block collection on load.
- Expose per-block runtime state required by the later tabbed UI task: block identity/index, running state, last poll, source status, and last error.
- Device operational `Status` is end-to-end health, not MMA2 ownership. `OWNED` is ownership/configuration information only and is not an acceptable operational success status.
- Device `Status = OK` only after all three conditions are true: (1) destination MMA2 memory is allocated/ready, (2) configured Pull Block poller(s) are successfully reading the source Modbus, and (3) the values read from the source are successfully written into the allocated MMA2 memory.
- A successful source read without a successful destination write must not report `OK`.
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
4. Legacy/current single-pull-block documents load as exactly one block without losing configuration.
5. Runtime status can report each block independently.
6. `Owner: replicator` remains ownership metadata and is not presented as operational success.
7. Device `Status` reports `OK` only when destination memory is allocated/ready, source polling succeeds, and the polled values are successfully written to MMA2 memory.
8. If any required end-to-end stage fails — memory unavailable, source read failure, or destination write failure — device `Status` must not report `OK` and runtime diagnostics must expose the failing stage/error.

## Verification
Focused Go tests using one device with at least two blocks at different scan rates, proving independent cycle counts/status plus single-block migration and persistence round-trip. Add end-to-end status tests proving `OK` only after a successful source-read-to-MMA2-write cycle, plus negative cases for unavailable memory, failed source read, and failed destination write. Rendered UI verification must confirm `Owner: replicator` remains separate while the displayed device `Status` follows the end-to-end semantics rather than showing `OWNED` as success.

## Dependencies
REP-BLOCK-001 must be complete/accepted first because this task extends its explicit Pull Block model.

## Sizing
Implementation surface 1; environment 0; behavior 2; verification 1; decision/recovery 1. Total 5 — tightly coupled backend/runtime model with one deterministic test workflow; keep bounded here and separate the UI tab work into REP-BLOCK-003.
