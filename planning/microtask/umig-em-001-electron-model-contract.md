# UMIG-EM-001 — DESIGN: Reconcile the Electron Toolkit Model with OS.js

Status: PLANNED — human changed migration direction on 2026-09-19; NOT promoted or executable by JR.
Stage / owner: DESIGN / ChatGPT; human approves successor implementation and promotion.
Previous: UMIG-006-V (COMPLETE/PASS at its tested SHA).
Next: none yet — explicit architecture decision gate. Author and link separately approved CODE/TEST/VERIFY successors before promotion.

## Primary outcome
Produce a source-backed cross-platform compatibility contract for following the CURRENT Electron MCS Modbus Toolkit model in Docker-hosted OS.js, keeping Go/MMA2 the backend authority. For this review pin latest assessed model at `1ca174058548996913bdcc13c7e2ef224a768457`: Electron Memory layout at `ff6846f973ad973bd3270b0ff2d707ce5c2616c9` plus the subsequent Go Replicator COMMS telemetry at `1ca1740`. Historical one-time visual donor `1c971b9a6e00bafadf329df8821421a40cfc079c` remains provenance only. Assess/re-pin each later Electron/Go push; never silently synchronize.

## Source findings / differences, NOT runtime verification
- Electron `renderer/app.js`, `memory-advanced.js`, `style.css`: Memory Device Definition / Advanced Settings folder tabs; nested RBE Rules, State Sealing and Access Policy tabs; shared MMA Settings opens as a separate modal inside Advanced. Device and shared drafts/save/discard remain separate. `memory-settings.js` inherits missing device fields and restricts shared root edits to `rbe`, `access_events`, `debug`.
- Electron `main.js` handles `mma-load`/`mma-apply` through local IPC, filesystem and bundled MMA2 `--validate-stdin`. This is not a portable OS.js backend contract. Do not copy preload/IPC/process control, Windows services/binary paths or direct browser-to-YAML writes.
- Go Simulator `device.go`, `mma2_config.go`, `mma2composer` now support per-device policy/sealing/RBE, inheritance, preservation of foreign-owned reservations and full effective-config validation. `MMA2/pkg/configvalidate` checks RBE. Current OS.js `memory-editor.js` has only definition/FC/None-Random editing; Go `simulator/runtime_server.go`, Toolkit `memory-contract.js` and `server.js` permit ONLY `load/apply/status`. No shared `mma-load/mma-apply`, advanced editor or shared modal exists in OS.js. Verify the existing Go load's effective-state inheritance behavior before using it to hydrate UI.
- New `replicator/comms.go`, `manager.go`, `reader.go`, `config_cycle.go` at `1ca1740` expose observed device/per-block network, TCP, Modbus and MMA2 evidence, timestamps, and conservative aggregate `comms` states. Electron `comms-status.js` explains older-backend schema absence and prevents stale/disabled green. The current OS.js Replicator UI intentionally pins four COMMS indicators UNKNOWN; that was correct at its previously tested backend SHA but is now a MODEL GAP, not proof that telemetry has been connected. Only fresh, matching, backend-proven observations may produce an OK state. The commit reports Go and Electron unit tests; no new Docker/Toolkit GUI test was independently reviewed here.
- `MMA2/docs/MCS_RBE_INTEGRATION.md`: global RBE TCP output plus per-memory rules; globally unique rule IDs 1..255; complete config validation and bind exposure constraints. `RBE_V1_DRAFT.md` remains NOT LOCKED: no TCP authentication, replay or proven production latency. Do not automatically enable/publish RBE TCP or claim installed/production safety.

## Exact DESIGN scope (one bounded outcome, three checks)
1. Write a field/interaction/state mapping: Electron per-device policy/RBE/sealing, global MMA RBE/access-events/debug, duplicate/draft/discard/save, and fresh Replicator per-block COMMS -> Go wire/status schema -> Toolkit. Classify implemented, missing, incompatible, Windows-only or uncertain. Separate functionality, visual appearance and actual runtime evidence.
2. Decide an authenticated Go-owned mechanism for shared MMA settings load/validated apply, explicit authorization, request/result/error shape, apply serialization, restart acknowledgement and rollback; preserve foreign owners, listener IDs/metadata and unknown config. Resolve whether Simulator `load` hydrates effective fields. Resolve COMMS freshness/name/version validation and unknown/error fallback; NEVER map aggregate OK alone to four green LEDs.
3. Specify separately promoted, small CODE -> JR UNIT/BUILD TEST -> disposable live VERIFY stages. Sequence: upgraded MMA2/Replicator baseline regression and backend contract, per-device advanced editing/round-trip, shared MMA modal, truthful backend-derived COMMS, then updated visual parity UMIG-007 and independent deployment UMIG-007A. Human cutover/retirement stays separate.

## Acceptance / evidence
A. Versioned compatibility matrix sourced to exact paths and SHAs, with explicit unknowns and distinct device/global scope.
B. Approved backend ownership/validation/failure contract and telemetry semantics before CODE; RBE listener safety explicit.
C. Sized and linked individual CODE/TEST/VERIFY microtask files and safe evidence gates before JR instructions. UMIG-007 needs a retrievable approved CURRENT-model three-tab screenshot baseline or approved reproducible capture with matched state/window dimensions only AFTER functional alignment.

## Non-scope and size
No product, Electron, Go, Compose, ICC, `operation cwal.md`, production, retained-volume or cutover changes. This PLANNED task is not a JR instruction or test PASS. Surface 0, environment 1, behavior 1, verification 1, recovery 0 = 3.
