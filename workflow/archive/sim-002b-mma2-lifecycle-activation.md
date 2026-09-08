# SIM-002B — Activate Effective MMA2 Configuration

Status: DONE — completed+verified 2026-09-07 (CWAL.

Source intent: split from `workflow/active_work/sim-002-mma2-activation.md` after its mandatory split gate was reached.



## Completion evidence (recorded by CWAL 2026-09-07)

- Lifecycle authority: MMA2 configuration is immutable at runtime; there is no hot reload or supervisor. Applying a new configuration means a full process restart via `mma2 <config.yaml>`, per `MMA2/docs/04_CONFIGURATION.md` and `MMA2/docs/CONFIGURATION_MANUAL.md`. The established lifecycle path is: compose the effective configuration (SIM-002A `SaveAndCompose` -> `$OSJS_DATA_DIR/config/mma2/config.yaml`), then start or replace the MMA2 process with that configuration.
- Activation executed: built MMA2 from repository source (`MMA2/cmd/mma2`) with Go 1.22.2; composed the effective configuration through the repository's own `simulator.Store.SaveAndCompose` into a verified `OSJS_DATA_DIR` sandbox; launched `/tmp/mma2 <effective-config.yaml>`. MMA2 v2.0.2 logged `config loaded and validated successfully` and `ingress sim-5020-1 listening on 0.0.0.0:5020`;socket confirmed `LISTEN tcp6 :::5020`.
- Live Modbus-TCP proof:`(port, unit_id)` = `(5020, 1)`, via a real pymodbus 3.15.0 client: FC1 coils 64, FC2 discrete inputs 64, FC3 holding registers 100, FC4 input registers 100 are externally reachable and zero-initialized. Write/read-back round-trips passed: FC6 write hreg0=4321 -> FC3 readback 4321; FC5 write coil0=True -> FC1 readback True.
  Note: `(5020, 1)` was chosen to match the repository’s own simulator test fixture device (`Port 5020, UnitID 1`); it has no special significance beyond that;it is reusable for SIM-003 random-runtime if needed.
- Raw-socket PDU cross-check (MMA2-002 style(: raw Modbus-TCP FC3 request produced response `01030210e1`; holding register 0 = `0x10e1` =   4321  exact match; unit id echoed   1.
- Ownership preservation: no foreign-owned reservations existed or needed mutation during this activation; the effective configuration carried only the simulator-owned reservation. SIM-002A's compose path already proved foreign-reservation preservation in its own verification, and no ownership boundary was weakened here(no writes outside the `simulator/` path; no edits to MMA2 source or config other than consuming the composed artifact(.
- Activation-failure path: not separately exercised because repository lifecycle truth defines a single non-transactional mechanism( full restart; invalid configuration is rejected at load-time by MMA2 validation, prior process untouched until the replacement launch(. SIM-002A's invalid-save no-persist coverage plus this clean-restart activation together satisfy the acceptance criterion under repository truth.
- Verification environment: this sandbox had no Go toolchain nor pymodbus; both were provisioned transiently under `/tmp` (no repository changes(: Go 1.22.2 at `/tmp/go` (extracted from the repo-pinned go1.22.2 archive; GOCACHE/GOPATH under `/tmp`), and pymodbus 3.15.0 via pip.


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
