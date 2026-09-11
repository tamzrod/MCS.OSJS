# REP-UI-005 — Replicator Save and Apply Device

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Wire the approved Replicator device definition to persisted Replicator configuration and shared MMA2 destination reservation through Save & Apply.

## Scope
- Save source fields from the UI into the Replicator configuration model.
- Save the resolved destination port/unit ID.
- Validate before persistence/apply.
- Commit only Replicator-owned destination changes through the ownership guard.
- Discard restores the last persisted definition.
- Report validation/collision/apply failures truthfully.

## Non-Scope
- No new replication algorithm.
- No multiple ranges or targets.
- No advanced retry/backoff.
- Runtime restart/reload orchestration is a separate task.

## Acceptance Criteria
1. Valid device definition persists and reloads correctly.
2. Invalid source/destination definition is rejected without corrupting prior persisted state.
3. Foreign ownership collision prevents apply and leaves foreign reservation untouched.

## Verification
Save a valid device, reload it, reject an invalid edit, and reject a controlled foreign-owned destination while confirming prior config/ownership remains intact.

## Dependencies
REP-UI-004; existing Replicator config store.

## Sizing
Implementation 1; environment 0; behavior 1; verification 1; decision/recovery 1. Total 3 — good bounded task.
