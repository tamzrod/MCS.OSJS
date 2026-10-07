# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work on 2026-10-07.

OpenHands has two distinct CWAL modes:
- **JR DEV** — CODE/DISCOVERY implementation for an explicitly assigned active task.
- **Independent JR** — separately assigned TEST/VERIFY execution only.

One invocation uses exactly one mode. JR DEV self-checks are not independent TEST/VERIFY evidence.

## Completed predecessor
**PERSIST-001 — Persistence Configuration Schema**
CODE COMPLETE at source checkpoint `faa33929429a0382a64b78bc773f0074e11cb6b1`.

## Current task — ACTIVE / TRANSPORT PENDING
**PERSIST-002 — State Sealing Prerequisite Validation**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-002-state-sealing-prerequisite-validation.md`

OpenHands reported a successful bounded implementation/self-check, but GitHub main does not yet contain the source changes. Therefore PERSIST-002 is NOT yet repository-complete and PERSIST-003 must not start.

### Exact transport authority for the next OPERATION CWAL invocation

Use the existing OpenHands workspace containing the reported PERSIST-002 changes. Do not reimplement if those exact changes are still present.

Required preflight:
1. Confirm current branch is `main`.
2. Confirm working-tree changes are exactly these paths and no others:
   - `mma2composer/memory_validation.go`
   - `mma2composer/persistence_sealing_validation_test.go`
   - `simulator/persistence_config_schema_test.go`
   - `workflow/active_work/persist-002-state-sealing-prerequisite-validation.md`
3. Confirm the implementation base/ancestor includes `c3ef3a2f4e81e4e2330a0e8d7c7574b41f0e7ad5`.
4. Confirm the reported task-bounded self-check evidence is still represented in the PERSIST-002 packet.
5. If the workspace is clean, stale without these changes, contains additional unrelated changes, or has diverged unexpectedly: BLOCKED / STOP. Do not reconstruct, reset, clean, force, or infer missing source during this transport invocation.

If preflight matches exactly, this invocation explicitly authorizes:
- stage ONLY the four paths above;
- commit them as one PERSIST-002 source/workflow checkpoint;
- commit message: `PERSIST-002: require state sealing for persistence`;
- push non-force to `main`;
- verify remote `main` resolves to the new commit;
- verify the pushed commit changed only the four authorized paths;
- report the pushed commit SHA and STOP.

No additional implementation, retest, successor activation, ICC write, cleanup, rebase, force-push, or PERSIST-003 work is authorized in this invocation.

## Successor routing
Only after the PERSIST-002 source checkpoint is confirmed on GitHub may the workflow owner mark PERSIST-002 repository-complete and activate PERSIST-003.

Only BLACK SHEEP WALL edits ICC.
