# PERSIST-002 — State Sealing Prerequisite Validation

Status: ACTIVE — CODE COMPLETE LOCALLY / TRANSPORT PENDING 2026-10-07
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-001
Next: PERSIST-003

## Primary outcome
Reject persistence-enabled memory unless State Sealing is present and enabled.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not change State Sealing semantics or implement persistence runtime.

## Acceptance
1. Persistence ON + sealing OFF/absent fails validation with a clear error.
2. Persistence OFF remains valid with sealing either ON or OFF.
3. Validation does not silently enable or mutate State Sealing.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-001. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/1/1/1=4; one validation invariant.

## CWAL
PERSIST-001 has genuine CODE completion evidence at `faa33929429a0382a64b78bc773f0074e11cb6b1`. Human/workflow owner has explicitly selected this packet as the sole current ACTIVE assignment for OpenHands JR DEV. OPERATION CWAL executes this one task, records evidence, and STOPS.


## Transport status

OpenHands JR DEV reported successful implementation and bounded self-checks from base `c3ef3a2f4e81e4e2330a0e8d7c7574b41f0e7ad5`, but the source changes are not yet present on GitHub main.

Reported exact changed paths:
- `mma2composer/memory_validation.go`
- `mma2composer/persistence_sealing_validation_test.go`
- `simulator/persistence_config_schema_test.go`
- this task packet

Reported acceptance:
1. Persistence enabled with absent/disabled State Sealing rejects with `persistence.enabled requires state sealing to be present and enabled`.
2. Persistence disabled remains valid with sealing enabled or disabled.
3. Validation is read-only and does not mutate State Sealing.

Reported self-checks were successful, but remain JR DEV evidence only, not independent TEST/VERIFY PASS.

Next CWAL invocation is transport-only per `handoff.md`: if and only if the exact reported dirty workspace still exists with no unrelated changes, commit and non-force push those four paths, verify remote delivery, report SHA, and STOP. Do not start PERSIST-003.
