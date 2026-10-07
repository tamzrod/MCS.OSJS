# PERSIST-020 — Restore Failure Behavior

Status: QUEUED — HUMAN PROMOTED 2026-10-07
Stage: CODE
Owner: OpenCode
Previous: PERSIST-019
Next: PERSIST-021

## Primary outcome
Keep persistence-enabled memory sealed and expose a deterministic failure state whenever startup persistence restore cannot safely complete.

## Scope
Implement only restore-failure behavior against the current authoritative MCS.OSJS source and existing MMA2 State Sealing/Raw Ingest contracts.

## Non-scope
Do not add repair/retry loops, operator-data mutation, historian behavior, or alternate unseal paths.

## Acceptance
1. Missing, corrupt, or incompatible snapshot keeps the memory sealed.
2. Raw Ingest or restore-verification failure keeps the memory sealed.
3. Failure reason is surfaced without fabricated defaults or automatic unsafe unseal.

## Evidence / handoff
Record exact changed paths, source diff, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-019.

## Sizing
1/0/2/1/2=6.

## CWAL
Human has promoted this packet into Active Work. It is QUEUED, not the repository's sole current ACTIVE assignment. OPERATION CWAL must execute exactly one task selected by `handoff.md`.
