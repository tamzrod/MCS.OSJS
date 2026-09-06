# SIM-002B — Activate Effective MMA2 Configuration

Status: ACTIVE — human-promoted for JR execution.

Source intent: split from `workflow/active_work/sim-002-mma2-activation.md` after its mandatory split gate was reached.

## Primary Outcome

An already-valid effective MMA2 configuration can be applied through the repository's lifecycle path, and the requested simulator `(port, unit_id)` resource is proven live without weakening ownership protection.

## Scope

- Consume an effective MMA2 configuration that has already passed SIM-002A ownership/collision validation.
- Determine and use the repository-established MMA2 reload/restart mechanism; do not invent a new lifecycle authority if repository truth already defines one.
- Apply the effective configuration only after successful validation/composition.
- Verify MMA2 returns to an operational state after the structural change.
- Verify the accepted simulator listener/Unit/ranges are exposed through Modbus TCP.
- Preserve the previously working runtime when activation fails wherever the existing lifecycle mechanism supports transactional or recoverable application.

## Non-Scope

- No `(port, unit_id)` ownership/composition algorithm; that belongs to SIM-002A.
- No changes to foreign-owned MMA2 reservations.
- No random-runtime scheduler.
- No raw-ingest data generation.
- No OS.js simulator UI.
- No Replicator implementation.

## Acceptance Criteria

1. One SIM-002A-valid effective MMA2 configuration is applied through the established lifecycle mechanism.
2. MMA2 is operational after activation.
3. The requested simulator `(port, unit_id)` and configured ranges are externally reachable and behave according to MMA2's existing runtime semantics.
4. A failed activation does not authorize ownership bypass or mutation of foreign-owned reservations, and the prior working runtime/configuration is preserved or restored according to the repository's supported lifecycle behavior.

## Verification

Apply one valid effective configuration, verify MMA2 runtime health, then connect through Modbus TCP to the requested `(port, unit_id)` and prove the configured address ranges are exposed. Exercise one supported activation-failure path if practical and verify recovery behavior against repository truth.

## Dependencies

- SIM-002A completed and verified.
- Existing MMA2 runtime/lifecycle behavior.

## Sizing

**3 / 10 — Good JR task.** One lifecycle/activation outcome with one runtime verification path; configuration composition and collision ownership are already resolved by SIM-002A.
