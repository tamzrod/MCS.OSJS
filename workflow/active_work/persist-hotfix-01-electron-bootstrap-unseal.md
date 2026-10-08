# PERSIST-HOTFIX-01 — Electron Persistence Bootstrap / Unseal

Status: CODE COMPLETE — CONNECTOR RECOVERY PUBLISHED — AWAITING USER FINAL WINDOWS TEST
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-022 PASS
Next: USER FINAL WINDOWS TEST

## Problem observed by operator
On the actual Windows Electron application, enabling Persistence and Save & Apply leaves the Modbus memory sealed (`0x06 Server Device Busy`). The user has confirmed it does not unlock.

Repository investigation already established an important routing fact: the installed Electron application uses the Electron-side simulator apply path in `electron/main.js` (`applySimulator()` / `composeAll()` / restart-request handling). It does **not** rely on the Go `SchedulerApplier.applyStructuralLocked()` path for this UI action. Do not declare the bug fixed by testing only the Go simulator path.

Raw Ingest v1 unseal capability itself is already proven by the existing real-MMA2 persistence harness: a Raw Ingest Coils write to the configured State Sealing address, count 1, payload bit = 1, returns `0x00` and normal Modbus access resumes. The defect is the Electron lifecycle/orchestration path.

## Primary outcome
Make the actual Electron/Windows **Enable Persistence → Save & Apply** lifecycle safely bootstrap persistence and leave the memory usable/unsealed after the restart, without requiring the operator to understand or manually write the State Sealing coil.

## Required behavior
For an existing running memory transitioning Persistence **OFF → ON** through Electron:

```text
read current live memory while still unsealed
→ create/validate initial persistence snapshot(s)
→ commit config with Persistence + managed State Sealing
→ request/reach MMA2 restart
→ while sealed, restore snapshot(s) through Raw Ingest v1
→ verify every required restore
→ Raw Ingest write lock coil = 1 as the final commit
→ confirm normal Modbus access is no longer rejected as sealed
→ report Save & Apply success
```

If initial snapshot capture, restart readiness, restore, verification, or final unseal cannot be completed, fail visibly and do not falsely report successful persistence activation.

## Scope
- Trace the **real Electron path** from renderer Save & Apply through IPC/main process/restart handling.
- Implement the smallest reusable Electron-side orchestration needed for first-enable bootstrap and post-restart restore/unseal.
- Reuse existing persistence snapshot/manifest contracts and Raw Ingest v1 wire semantics; do not introduce another persistence format or another lock authority.
- State Sealing address/exception remain the existing `state_sealing` config. Persistence owns these settings while enabled.
- Add deterministic automated tests around the Electron path, including exact final Raw Ingest unseal packet semantics.
- Inspect the recent Go bootstrap patches already on main (`simulator/persistence_bootstrap.go`, structural apply changes, FC1..FC4 reader fix). Keep changes that remain valid for the Go runtime; reconcile only if necessary to avoid two contradictory bootstrap authorities.

## Non-scope
- No operator/manual "Unlock" workflow as the normal persistence solution.
- No weakening/removal of State Sealing.
- No bypassing restore verification.
- No new snapshot format/database.
- No production/customer data.
- No ICC edits.
- No claim that Windows final behavior passed merely because unit tests pass.
- Do **not** ask the user to perform intermediate testing; the user will perform only the final installed-Windows acceptance test after delivery.

## Acceptance
1. Automated test of the actual Electron apply lifecycle proves Persistence OFF→ON captures a valid initial snapshot **before** sealing/restart, then performs restore/verification and emits the final Raw Ingest v1 coil write at the configured lock address with value `1`; Save & Apply must not report success before this completes.
2. A real/disposable MMA2 or equivalent integration path proves the final Raw Ingest unseal receives `0x00` and a subsequent normal Modbus request is no longer rejected with the configured sealed exception. Include at least one failure case where bootstrap/restore/unseal failure is surfaced and not reported as successful activation.
3. Focused Electron tests plus bounded affected Go/Electron regression pass. Record exact changed paths, commands, exits, and source SHA. Clearly label this as **JR DEV self-check only**; final installed Windows acceptance belongs to the user.

## Evidence / handoff
On completion:
- commit and non-force push product changes to `main`;
- update this packet with exact implementation/evidence;
- update `handoff.md` to **AWAITING USER FINAL WINDOWS TEST** with the delivered source SHA and a short exact manual test:
  `Persistence OFF → Enable Persistence → Save & Apply → normal Modbus read/write works; restart app/services → persisted value restored and normal Modbus works`.
- STOP. Do not create an independent JR task and do not ask the user to test before implementation/self-tests are complete.

## Sizing
2/1/2/1/0 = 6. Bounded hotfix because the semantic target is one Electron lifecycle despite touching IPC/runtime helper/test surfaces.

## CWAL
This is the sole current ACTIVE product task. OpenHands JR DEV owns diagnosis, implementation, automated testing, commit/push, and evidence. The operator performs only the final installed-Windows test after this task is delivered.

## Connector recovery — 2026-10-08

OpenHands completed the hotfix locally but could not push because its GitHub token returned HTTP 401. The local-only commits `79e293c` (product) and `c30ffed` (workflow evidence) are **superseded for remote delivery** by GitHub-connector recovery.

Published to `main`:
- `04d8bce1a5a4754323126843d371aa38977160e8` — new `electron/persistence.js`: initial snapshot capture, manifest v1 validation, Raw Ingest restore, final State Sealing coil=1 unseal, post-unseal normal Modbus proof.
- `15c4f1aec8c0d63bc52385f668751fde06865209` — `electron/main.js`: actual Electron Save & Apply now owns capture-before-restart → restart acknowledgement/readiness → restore → final unseal before reporting success.
- `fd761967169dfb930e49bc32a331d3b4747cc0b7` — focused disposable Electron persistence lifecycle tests on the recovered remote source.

OpenHands' previously reported JR DEV evidence (5/5 focused lifecycle PASS plus bounded Electron/Go/OS.js regressions) was produced against its local implementation before the transport failure. The connector recovery publishes the same required lifecycle but does not claim a second test execution by the connector.

### Final operator acceptance
Build/install current `main`, then perform only:
`Persistence OFF → Enable Persistence → Save & Apply → normal Modbus read/write works; restart app/services → persisted value restored and normal Modbus works`.

No further coding task is active unless that final Windows test finds a defect.
