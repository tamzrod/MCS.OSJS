# Handoff

## Current

COMPLETED + ARCHIVED:
- MEM-001 — Rename Memory UI — static verification passed.
- MEM-002 — Simplify Header Status — static verification passed.
- MEM-003 — Simulator Status in Diagnostics — static verification passed.

ACTIVE: MEM-004 — Allow None Simulation.

QUEUED (approved sequence, explicit order):
- MEM-005 — Area Simulation Selector. Previous: MEM-004.
- MEM-006 — Default None, Preserve Legacy Random. Previous: MEM-005.
- MEM-007 — Truthful None-mode Status. Previous: MEM-006.
- MEM-008 — Windows Memory Acceptance. Previous: MEM-007.

Unrelated QUEUED tasks unchanged: REP-BLOCK-002 (JR retest); REP-BLOCK-003 (JR rendered retest).

## Decision / scope

Minimal change to the integrated Electron Simulator editor. Memory presently displays simulator-owned devices only. Keep Simulator as internal runtime/service/API/storage name, and keep MMA2 ownership untouched. None/Random uses zero/positive existing `random_runtime.*_interval_ms`; zero does not deallocate memory. New devices default None; saved positive intervals remain Random. Header shows steady MMA2 + Replicator only; Simulator runtime status moves to Diagnostics.

## Implementation state

The committed source already contains the MEM-004 through MEM-007 implementation surface: zero interval validation support, focused None/mixed scheduler/status regression tests, None/Random renderer selector, new-device None defaults, legacy positive-interval Random interpretation, and truthful all-None IDLE / NOT REQUIRED status.

Formal advancement is intentionally stopped at MEM-004 because its task verification explicitly requires:

```text
cd simulator
go test -count=1 ./...
go vet ./...
```

Those commands have not been executed in this GitHub-only session, so no Go test/vet success is claimed. After that gate passes, MEM-005 through MEM-007 can be advanced and verified in order; MEM-008 remains the final real-Windows acceptance gate.

## Process

All eight MEM tasks are human-approved, but only one is ACTIVE at a time per `workflow/active_work/README.md`. Verify and advance in sequence. Windows-only MEM-008 cannot be accepted without real Windows evidence. Preserve any parked local/uncommitted Electron source/assets when pulling; GitHub main is the committed baseline, not the local overlay.
