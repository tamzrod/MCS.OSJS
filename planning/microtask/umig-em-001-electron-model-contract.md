# UMIG-EM-001 — DESIGN: Reconcile the Electron Toolkit Model with OS.js

Status: PLANNED — human changed the migration direction on 2026-09-19; NOT promoted, NOT executable by JR.
Stage / owner: DESIGN / ChatGPT; human approves implementation sequence and promotion.
Previous: UMIG-006-V (COMPLETE/PASS; historical live Diagnostics evidence retained).
Next: none yet — explicit architecture/approval gate; author separate CODE, TEST and VERIFY successor files and link them before promotion.

## Primary outcome
Produce one bounded, source-backed compatibility contract for migrating the current Electron MCS Modbus Toolkit model to Docker-hosted OS.js while keeping Go/MMA2 as the backend authority. The migration target for THIS review is Electron at `ff6846f973ad973bd3270b0ff2d707ce5c2616c9`, not the old one-time visual donor `1c971b9a6e00bafadf329df8821421a40cfc079c`. A later Electron push is a new diff to assess and explicitly re-pin; it does not silently authorize changes to OS.js.

## Verified source inventory and gap
- Electron `renderer/app.js`, `memory-advanced.js` and `style.css`: Memory uses Device Definition / Advanced Settings folder tabs; Advanced has RBE Rules, State Sealing and Access Policy folder tabs, with shared MMA Settings in a modal opened there. Device drafts and shared drafts/save/discard are distinct. `memory-settings.js` hydrates/inherits missing advanced device fields from effective config and restricts shared edits to root `rbe`, `access_events`, `debug`.
- Electron `main.js` implements `mma-load` / `mma-apply` using filesystem IPC, whole-config validation via a bundled MMA2 `--validate-stdin`, then a restart request. These are WINDOWS/ELECTRON mechanisms, not an OS.js backend API; do not copy preload, IPC, process control, binary paths or direct renderer-to-YAML writes.
- Go Simulator `device.go` and `mma2_config.go` now contain policy, state_sealing, RBE and inheritance/composition; `mma2composer` validates the full effective configuration and preserves producer-owned/foreign reservations. `MMA2/pkg/configvalidate` checks RBE rules. These are source-level capabilities, not proof that a rebuilt Docker stack and prior runtime verification still pass.
- Current Toolkit `memory-editor.js` has only basic device definition/FC/None-Random editing; `memory-contract.js`, Go `simulator/runtime_server.go`, and Toolkit `server.js` offer ONLY Simulator load/apply/status. There is NO `mma-load` / `mma-apply` on the deployed Unix protocol, no shared MMA modal and no per-device advanced editor. `toolkit-renderer.js` remains a frozen one-time fixture visual adaptation. Existing Memory/Replicator/Diagnostics behavior and their archived PASS evidence are retained at their tested SHA, not extended to the new MMA2 changes.
- `MMA2/docs/MCS_RBE_INTEGRATION.md` specifies global RBE TCP output plus per-memory rules, rule IDs 1..255, full-config validation and source exposure precautions. `RBE_V1_DRAFT.md` remains NOT LOCKED, has no client authentication/replay and no production latency qualification. The new RBE feature cannot be assumed production-safe or enabled by default.

## Exact DESIGN scope
1. Create a field/interaction matrix mapping Electron per-device policy/RBE/sealing, shared global RBE/access-events/debug, draft/duplicate/discard/save, and statuses to present Go, v1 Unix and Toolkit implementation. Mark supported, missing, Windows-only, uncertain and incompatible; separate UI appearance from functional/runtime parity.
2. Decide the ownership of authenticated read/validated apply for shared MMA settings in Docker (producer-neutral or explicitly scoped Go-side service), exact request/response/errors, concurrent apply/restart behavior and preservation of foreign owners, listener metadata and unknown settings. Verify whether Simulator `load` returns inherited effective settings on all paths before choosing UI hydration. Do not design a client-side direct config writer or route global settings through a device `apply` by accident.
3. Specify ordered, separately promoted microtasks with CODE -> independent JR UNIT/BUILD TEST -> disposable live VERIFY boundaries. Cover backend contract and new MMA2 Go regression FIRST, then per-device advanced editing/round-trip, then separate global modal, then updated three-tab visual parity and independent deployment. Explicitly prohibit production activation, exposing an RBE listener without reviewed bind/auth/network controls, legacy retirement and launcher cutover.

## Acceptance / evidence
A. One written compatibility matrix backed by exact current source paths/commit and unresolved differences; no generic screenshot alone is sufficient.
B. One explicit Go-owned management/validation/ownership contract and failure/recovery strategy accepted before backend CODE; preserve disabled or UNKNOWN unsupported controls.
C. Independently sized CODE/TEST/VERIFY task files with predecessor/successor and safe test targets, approved before any JR packet. Only after functional alignment should UMIG-007 compare visuals against a newly approved, reproducible *current-model* three-tab donor baseline.

## Non-scope / authority
No product implementation, Docker run, Go/MMA2 change, Electron edit, ICC edit (only BLACK SHEEP WALL may refresh affected ICC), `operation cwal.md` edit, screenshot fabrication, new test PASS, automatic cutover or promotion by this PLANNED document. This is not a current JR packet.

## Sizing
Surface 0, environment 1 (cross-platform backend decision), behavior 1 (model matrix), verification 1 (contract acceptance), recovery 0 = 3.
