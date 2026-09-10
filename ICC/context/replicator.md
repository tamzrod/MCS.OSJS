# Replicator Authorized Sequence

Baseline commit: 366b23e
Working tree: clean
Audited Overlay: none
Source dependencies: handoff.md, workflow/active_work/rep-*.md, mma2composer/*, simulator/mma2_config.go, simulator/compose_test.go, MMA2/pkg/configvalidate/*
Zoom In: none
Zoom Out: active-work

## Current boundary

REP-001 is the sole ACTIVE task. REP-002 through REP-007 are QUEUED in the explicit authorized sequence recorded by `handoff.md` and their `Previous` / `Next` links.

REP-001 is limited to producer-neutral MMA2 effective-config and ownership composition. It must preserve `$OSJS_DATA_DIR/config/mma2/config.yaml` and `owners.yaml`, atomic candidate validation/write behavior, and foreign-owner collision rejection. Simulator must move to the shared implementation using producer identity `simulator`, without intended behavior change.

## Checkpoint at baseline

- `mma2composer/` is a standalone Go module with a verified-good fresh `composer.go` and an explicit producer parameter on `Composer`.
- The shared implementation contains config/ownership loading and saving, atomic writes, reservation drop/add/collision mechanics, authoritative MMA2 YAML validation, and rollback of config bytes if the owners write fails.
- Focused `mma2composer` tests do not yet exist.
- Simulator still carries and uses its private composition implementation in `simulator/mma2_config.go`; migration to the shared module is not yet implemented.
- The full Simulator suite was green at the checkpoint, before migration.

## Completion gate

Add focused shared-package coverage for producer identity, foreign-owner rejection, first-come-first-save persistence, and commit rollback. Migrate Simulator to the shared composer, retain its public Store path shims and ownership semantics, then pass the shared-module native checks and the full Simulator suite before archiving REP-001.

REP-002 and later tasks are outside the current execution boundary until deterministic advancement after REP-001 completes and is pushed.
