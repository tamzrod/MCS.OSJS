# Replicator Authorized Work

Baseline commit: ee19b8a
Working tree: dirty; no Replicator workflow overlay. See `ICC/INDEX.md` for the audited Electron overlay.
Source dependencies: handoff.md, workflow/active_work/rep-block-002-independent-block-pollers.md, workflow/active_work/rep-block-003-tabbed-block-editor.md, workflow/archive/rep-*.md
Zoom In: none
Zoom Out: active-work

## Current boundary

The original REP-002 through REP-006 implementation sequence is archived. REP-007 is no longer active. Current authorized Replicator work remains queued, with no active predecessor:

- REP-BLOCK-002: multi-block persistence, independent pollers, mixed FC1-FC4 replication, external destination serving, and end-to-end operational status. The record says implementation exists but JR retest is pending.
- REP-BLOCK-003: classic Device/Pull Blocks folder tabs and compact spreadsheet rows. It depends on REP-BLOCK-002 and awaits rendered retest.

Neither Replicator record is ACTIVE.

## Workflow integrity observations

- REP-BLOCK-002 declares a sizing dimension value of 3 even though planning rules limit every dimension to 0-2; it totals 8 and is oversized under the current rules.
- REP-BLOCK-003 depends on REP-BLOCK-002, but the pair lacks explicit `Previous` / `Next` advancement links required for an ordered promoted sequence.
