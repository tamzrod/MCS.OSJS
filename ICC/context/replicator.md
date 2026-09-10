# Replicator Authorized Sequence

Baseline commit: 4a137eb
Working tree: clean
Audited Overlay: none
Source dependencies: handoff.md, workflow/active_work/rep-*.md, mma2composer/*, simulator/mma2_config.go, simulator/compose_test.go, MMA2/pkg/configvalidate/*
Zoom In: none
Zoom Out: active-work

## Current boundary

REP-002 is the sole ACTIVE task. REP-003 through REP-007 are QUEUED in the explicit authorized sequence recorded by `handoff.md` and their `Previous` / `Next` links. REP-001 is completed and archived.

REP-002 is limited to producer-neutral Raw Ingest v1 packet encoding and TCP send/ack behavior. It must preserve FC1/FC2 LSB-first bit packing, FC3/FC4 big-endian register encoding, and existing response/error behavior while removing Simulator's private duplicate.

## REP-001 completion at baseline

- `mma2composer/` owns config/ownership loading, atomic writes, collision mechanics, authoritative MMA2 YAML validation, rollback, and producer-neutral reservation operations.
- Simulator delegates composition through producer identity `simulator`; its former private duplicate is gone.
- Focused shared tests and the full Simulator suite and vet gates passed before REP-001 was archived.

## Completion gate

Create a Simulator-independent shared Raw Ingest client, move/adapt byte fixtures for all four areas, migrate Simulator onto it, and pass the shared native gates plus the affected full Simulator suite before archiving REP-002.

REP-003 and later tasks are outside the current execution boundary until deterministic advancement after REP-002 completes and is pushed.
