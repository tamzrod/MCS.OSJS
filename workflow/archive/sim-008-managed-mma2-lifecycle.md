# SIM-008 — Activate MMA2 Through a Managed Lifecycle

Status: COMPLETED — verified 2026-09-08 (CWAL).

## Completion Evidence

- Added `simulator/lifecycle.go`: one mutex-serialized MMA2 child-process owner, `MMA2_BINARY` executable contract, effective-config launch, all-listener readiness plus child-stability confirmation, graceful stop/kill fallback, structural replacement, and prior-config restart on failed replacement.
- Structural Save & Apply snapshots effective config and ownership artifacts, composes only through the existing `(port, unit_id)` ownership authority, activates MMA2, and restores both files on activation failure. Simulator document persistence remains after downstream success.
- Existing effective configuration is activated during simbridge startup. Shutdown stops the managed MMA2 child. `SIMULATOR_BRIDGE_ADDR` permits isolated loopback verification while retaining `127.0.0.1:18211` as the default.
- Automated verification: `GOCACHE=/tmp/mcs-osjs-go-cache go test -count=1 ./...` and `go vet ./...` passed. `TestMMA2LifecycleActivationReplacementAndRollback` proves initial listener activation, listener replacement, prior-service restoration, and no superseded listener.
- Real MMA2 workflow used `/tmp/mcs-mma2-sim008`, built from repository `MMA2/cmd/mma2`, with isolated data root `/tmp/mcs-sim008-OxkcB7` and bridge `127.0.0.1:18212`.
- Initial structural PUT returned HTTP 200 and started MMA2 on `*:15031`. Accepted replacement returned HTTP 200, removed `15031`, and left exactly one MMA2 child listening on `*:15032`.
- Deliberately applying port `18212`, already owned by the bridge, returned HTTP 422 with the concrete MMA2 bind error. Effective config, owners, and simulator document remained on port `15032`; `*:15032` was restored with exactly one managed MMA2 child.

## Primary Outcome

The simulator's ownership-validated effective MMA2 configuration is activated by one repository-owned lifecycle manager, leaving MMA2 running and listening on every configured simulator reservation without disrupting reservations owned by other applications.

## Scope

- Establish one managed MMA2 process lifecycle for the appliance runtime.
- Start MMA2 from the composed effective configuration under `$OSJS_DATA_DIR/config/mma2/config.yaml`.
- Route accepted structural Save & Apply changes through the lifecycle manager after shared `(port, unit_id)` ownership validation succeeds.
- Restart or reload MMA2 only when an accepted structural change requires it.
- Preserve the previously active MMA2 process and effective configuration when validation, ownership checks, configuration composition, or replacement startup fails.
- Stop superseded MMA2 processes after the replacement is proven ready so only the intended managed instance remains.
- Expose actionable activation failure details to the existing backend status path.

## Non-Scope

- No scheduler or random-value behavior.
- No raw-ingest data verification.
- No Replicator or Memory Appliance UI.
- No new ownership-key semantics; `(port, unit_id)` remains authoritative.
- No arbitrary MMA2 lifecycle controls in the OS.js window.
- No unrelated MMA2 refactor.

## Acceptance Criteria

1. With a valid simulator-owned reservation, the managed MMA2 instance starts from the effective configuration and binds the configured Modbus listener.
2. An accepted structural change activates the replacement configuration, while a rejected foreign ownership collision leaves the prior configuration and running MMA2 service unchanged.
3. Startup or replacement failure is reported and does not leave an untracked duplicate MMA2 process or falsely report the failed configuration as active.

## Verification Method

Use one lifecycle integration workflow:

```text
PREPARE OWNERSHIP-VALID EFFECTIVE CONFIG
-> START MANAGED MMA2
-> VERIFY PROCESS + LISTENER
-> APPLY ACCEPTED STRUCTURAL CHANGE
-> VERIFY REPLACEMENT LISTENER + SINGLE MANAGED INSTANCE
-> ATTEMPT FOREIGN COLLISION
-> VERIFY REJECTION + PRIOR SERVICE UNCHANGED
-> FORCE REPLACEMENT START FAILURE
-> VERIFY RECOVERY + ACTIONABLE STATUS
```

Record the effective configuration path, process identity, listener evidence, accepted-change evidence, collision rejection, failure/recovery result, and proof that no duplicate managed process remains.

## Dependencies

- Completed SIM-002A ownership-safe effective MMA2 configuration composition.
- Completed SIM-006 Save & Apply classification and rollback behavior.
- Repository-owned MMA2 build and known start/configuration contract.
- Human promotion completed 2026-09-08.

## Sizing Assessment

- Implementation surface: 1 — lifecycle manager plus bounded Save & Apply integration.
- Environment/dependency uncertainty: 1 — MMA2 build and launch contract are proven, but process supervision is new.
- Behavioral surface: 1 — one lifecycle behavior: safely activate the effective configuration.
- Verification surface: 1 — one continuous lifecycle/recovery workflow.
- Decision/recovery surface: 1 — replacement readiness and rollback require bounded recovery logic.
- Total: **5 / 10**. Acceptable as one task because activation, replacement, and rollback are inseparable parts of one managed lifecycle and share one deterministic verification workflow.
