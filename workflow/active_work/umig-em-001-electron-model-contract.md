# UMIG-EM-001 — DESIGN: Current Electron Model Compatibility Contract

Status: ACTIVE — human promoted 2026-09-19; design/source review checkpoint prepared; HUMAN ARCHITECTURE APPROVAL PENDING. Not CODE, TEST, or a JR packet.
Stage / owner: DESIGN / ChatGPT. Human ratifies or amends the proposed backend contract and promotes successor tasks separately.
Previous: UMIG-006-V (archived COMPLETE/PASS for historical tested SHA).
Next: UMIG-EM-002-T, only AFTER the contract approval gate and explicit human promotion of the ordered successor sequence. Until approval, no advancement.

## One primary outcome
Prepare and obtain human approval of the cross-platform data, authority, failure and telemetry contract for following the current Electron MCS Toolkit in OS.js. Source baseline: `449cfda29d23406295cd0ccbc35fa8f5d6293681`; UI/Go feature commits `ff6846f` and `1ca1740`, later Windows settings-owner additions also accounted for. Deliverable: `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md`. The historical `1c971b9` donor remains provenance only.

## Scope / gates
1. A source-backed field/interaction matrix identifies present, missing, platform-only, unverified and safety-restricted device/global settings, backend responses and four measured COMMS states.
2. An explicit PROPOSED Go-side authenticated shared management contract resolves canonical hydration, authorization, multi-writer serialization/CAS, whole-config validation, failure/recovery, restart acknowledgement, and truthful telemetry freshness. Human must approve unresolved architecture and security decisions before CODE.
3. Independently sized successor microtasks in Planning distinguish existing-upgrade regression TEST/VERIFY from new CODE -> JR TEST -> live VERIFY. Nothing in Planning is executable merely because this DESIGN is ACTIVE.

## Do not do
No source edits to OS.js/Electron/Go/MMA2/Compose, Docker, production data, existing test volumes, ICC, `operation cwal.md`, legacy removal, RBE network exposure or fabricated test acceptance. Only BLACK SHEEP WALL refreshes stale ICC. No JR TEST TASK until a separately promoted TEST/VERIFY stage has an exact packet.

## Completion / handoff
Source readback of the contract and all planned successor files is required. HUMAN APPROVAL of its proposed management/security/recovery design is an additional gate; source-only analysis does not equal such approval. Upon explicit approval, ChatGPT may archive this task, promote only the human-authorized next queued stage, and author its exact JR packet. If further Electron commits land, diff/re-pin before promotion.

Sizing: surface 0 + environment 1 + behavior 1 + verification 1 + recovery 0 = 3.
