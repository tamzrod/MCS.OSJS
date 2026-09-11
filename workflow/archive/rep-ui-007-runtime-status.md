# REP-UI-007 — Replicator Runtime Status

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Expose a lightweight truthful Replicator operational status area at the bottom of the application window.

## Scope
- Show Replicator RUNNING/STOPPED/error-relevant state from actual runtime state.
- Show Source OK/ERROR from actual cycle result state.
- Show last poll/last completed cycle timing in a compact form.
- Show last error when source/runtime is in error.
- Reuse existing `RuntimeState` rather than creating a parallel status model.

## Non-Scope
- No history database.
- No metrics dashboard.
- No charts.
- No advanced health/status block from donor Replicator Stack.

## Acceptance Criteria
1. RUNNING/STOPPED reflects actual Replicator runtime state.
2. Source OK/ERROR changes according to actual cycle success/failure.
3. Last poll and last error are based on runtime state and never fabricated by the UI.

## Verification
Run Replicator against a reachable source, then an unreachable source, then stop it; verify the status bar follows the actual runtime transitions.

## Dependencies
REP-UI-006; existing Replicator `RuntimeState`.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 0. Total 3 — good bounded task.
