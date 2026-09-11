# REP-BLOCK-001 — Explicit Source Pull Block Model

Status: COMPLETED / VERIFIED

## Primary Outcome
Replace the ambiguous device-level single FC/start/count source definition with one explicit Replicator source pull block so the configured Modbus read is declared as a block rather than inferred from device fields.

## Scope
- Introduce an explicit pull-block structure containing FC, Start, Count, and Scan Rate.
- First version remains exactly one pull block per Replicator device.
- Map the pull block 1:1 into the existing source read and destination area/range.
- Update persistence/runtime mapping without adding multi-block scheduling yet.
- Update the Replicator UI to present the source range as a clearly declared Pull Block.

## Verification Result
Backend gates and rendered Replicator Save & Apply verification passed in the corrected pair-scoped ownership test packet recorded in `handoff.md`.
