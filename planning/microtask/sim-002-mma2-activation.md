# SIM-002 — Validate and Activate Simulator MMA2 Parameters

Status: planning material only. Human promotion is required before execution.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Primary Outcome

A valid simulator MMA2 parameter set can be checked against the shared MMA2 namespace and activated without overwriting resources owned by another producer.

## Scope

- Accept only the simulator MMA2 parameter domain from a valid simulator definition.
- Interpret listener, Unit ID, FC, Start, and Count using the actual MMA2 addressing model in repository truth.
- Detect resource collisions before modifying active MMA2 configuration.
- Reject conflicting simulator requests without changing the current active MMA2 state.
- Compose/update only the simulator-owned contribution to effective MMA2 configuration.
- Restart or reload MMA2 only when the accepted structural change requires it.
- Verify MMA2 exposes the accepted simulator listener/Unit/ranges after activation.

## Non-Scope

- No random-runtime scheduler.
- No raw-ingest data generation.
- No OS.js simulator UI.
- No Replicator implementation beyond respecting its ownership boundary.
- No redesign of MMA2 addressing semantics.

## Acceptance Criteria

1. One non-conflicting simulator MMA2 definition is accepted and exposed by MMA2 at the requested listener/Unit/ranges.
2. One deliberate conflicting request is rejected before active MMA2 configuration is replaced.
3. After a rejected or failed activation, the previously working MMA2 configuration remains operational.

## Verification

Activate one valid simulator MMA2 structure and verify it through MMA2 runtime behavior. Submit one deliberate collision and prove it is rejected while the previous structure remains reachable.

## Dependencies

- SIM-001.
- Existing MMA2 config/runtime behavior.
- Approved shared MMA2 configuration-authority boundary.

## Sizing

**4 / 10 — Medium.** One structural activation transaction. If repository inspection reveals composition and lifecycle control are independently unresolved, split before promotion.