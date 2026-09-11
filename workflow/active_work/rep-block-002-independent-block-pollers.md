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

## Verification
Focused Go tests using one device with at least two blocks at different scan rates, proving independent cycle counts/status plus single-block migration and persistence round-trip.

## Dependencies
REP-BLOCK-001 must be complete/accepted first because this task extends its explicit Pull Block model.

## Sizing
Implementation surface 1; environment 0; behavior 2; verification 1; decision/recovery 1. Total 5 — tightly coupled backend/runtime model with one deterministic test workflow; keep bounded here and separate the UI tab work into REP-BLOCK-003.
