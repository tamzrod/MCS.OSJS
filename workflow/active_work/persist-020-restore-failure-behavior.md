# PERSIST-020 — Restore Failure Behavior

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
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
PERSIST-019 is delivered on GitHub main at `e22459ecdd10433ef21f63df25e3a22a43fb3662`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task: any missing/corrupt/incompatible snapshot, Raw Ingest failure, restore-verification failure, or unseal/commit failure must leave the persistence-enabled memory sealed and expose a deterministic failure reason; add no repair/retry loop, fabricated default, alternate unseal path, or operator-data mutation. Prepare PERSIST-021 for the next invocation only after genuine completion, then STOP.
