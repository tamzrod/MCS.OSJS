# PERSIST-014 — Snapshot Manifest / Compatibility Metadata

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-013
Next: PERSIST-015

## Primary outcome
Add snapshot metadata sufficient to reject incompatible or corrupted snapshots before restore.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not restore or unseal memory.

## Acceptance
1. Manifest binds snapshot to Port, Unit ID, area, start, count and format version.
2. Integrity metadata detects corrupted/incomplete area payloads.
3. Changed incompatible memory layout causes explicit rejection rather than silent truncation/remap.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-013. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/2=7; metadata/validation only.

## CWAL
PERSIST-013 is delivered on GitHub main at `ccccb891b737995f6a24a58e844e5bd1d794b609`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. If a reused workspace still has local-only PERSIST-013 commits `9343d9e` / `ab27b27`, treat them as superseded duplicate history and reconcile them with bounded safe Git mechanics; do not replay them as new product work. Execute exactly PERSIST-014, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
