# REP-UI-006 — Replicator Runtime Apply Control

Source: `brainstorm/replicator-ui-layout.md`

## Primary Outcome
Make an applied Replicator structural configuration take effect through the established MMA2/Replicator lifecycle without introducing hot reload or a new config API.

## Scope
- Reuse the established Simulator/MMA2 structural restart/reload control path where applicable.
- After a valid Replicator config and destination reservation are committed, ensure MMA2 sees the structural destination configuration.
- Start/restart the Replicator runtime against the persisted configuration only after MMA2 is ready.
- Preserve truthful failure state if restart/apply fails.

## Non-Scope
- No hot reload architecture.
- No new bridge service.
- No new generic config API.
- No replication algorithm changes.

## Acceptance Criteria
1. Save & Apply of a valid structural destination results in MMA2 serving that destination.
2. Replicator runtime uses the persisted applied configuration after lifecycle completion.
3. Failure to apply/restart is reported as failure rather than false RUNNING/success state.

## Verification
Apply a controlled Replicator device configuration, verify MMA2 destination becomes reachable, then verify Replicator runtime starts using that configuration; exercise one controlled apply failure and verify truthful state.

## Dependencies
REP-UI-005; existing MMA2 control/restart mechanism; existing Replicator runtime.

## Sizing
Implementation 1; environment 1; behavior 1; verification 1; decision/recovery 1. Total 5 — tightly coupled lifecycle task with one deterministic apply verification workflow; keep bounded to reuse of the proven control path.
