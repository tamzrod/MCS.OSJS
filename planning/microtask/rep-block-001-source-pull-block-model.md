# REP-BLOCK-001 — Explicit Source Pull Block Model

## Primary Outcome
Replace the ambiguous device-level single FC/start/count source definition with one explicit Replicator source pull block so the configured Modbus read is declared as a block rather than inferred from device fields.

## Scope
- Introduce an explicit pull-block structure containing FC, Start, Count, and Scan Rate.
- First version remains exactly one pull block per Replicator device.
- Map the pull block 1:1 into the existing source read and destination area/range.
- Update persistence/runtime mapping without adding multi-block scheduling yet.
- Update the Replicator UI to present the source range as a clearly declared Pull Block.

## Acceptance Criteria
1. Every enabled Replicator device has exactly one explicit source pull block.
2. FC/start/count/scan rate used by the poll cycle come from that block.
3. Unsupported/invalid block definitions are rejected before runtime apply.
4. Existing FC3/FC4 behavior remains valid after migration.

## Verification
Focused model/runtime mapping tests plus one UI Save & Apply test from Simulator FC3 source to Replicator destination.

## Sizing
Implementation 1; behavior 1; verification 1; migration 1. Total 4 — bounded.
