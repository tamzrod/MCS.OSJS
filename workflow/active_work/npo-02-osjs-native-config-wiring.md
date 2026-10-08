# NPO-02 — OS.js Toolkit Native Persistence Config Wiring

Status: PROMOTED / QUEUED
Stage: CODE
Owner: Codex JR DEV
Previous: NPO-01
Next: NPO-03

## Purpose
Wire OS.js Toolkit persistence settings into the existing authorized server-side MMA configuration path.

## Scope
Persist only native per-memory MMA2 persistence configuration. Preserve ownership, foreign memories, unknown extensions and revision/conflict safety.

Do not add any OS.js persistence runtime engine.

OS.js must not:
- write snapshots;
- restore memory;
- use Raw Ingest for persistence;
- control State Sealing for persistence;
- subscribe to RBE for persistence.

## Acceptance
Focused server/model tests prove correct round-trip, authorization boundaries, conflict handling and preservation of unrelated configuration.

STOP after delivery.
