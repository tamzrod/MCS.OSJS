# UMIG-007 — VERIFY: Three-tab visual parity to CURRENT Electron model

Status: QUEUED / DEFERRED — UMIG-EM-001 design ACTIVE, model architecture not yet human-approved, successors only PLANNED. No executable JR packet.
Stage / owner VERIFY / OpenHands JR. Previous: UMIG-EM-009-V (PLANNED; MUST be separately human-promoted and completed before this task; historical UMIG-006-V archived PASS remains separate). Next: UMIG-007A (QUEUED; deployment needs separate acceptance).

## Outcome and prerequisites
Compare functionally aligned OS.js Toolkit to latest specifically re-pinned Electron MCS Toolkit at matched window size and synthetic device state, separately for Memory Device Definition/Advanced/RBE/sealing/access, shared MMA modal, Replicator and Diagnostics. Old donor `1c971b9` is historical provenance, not current model. Contract `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md` proposes phases UMIG-EM-002-T/V through EM-009-V but DOES NOT promote them. Their source/code/unit/live gates, plus latest backend regression, must pass before visual parity; evidence from UMIG-006-V is earlier-SHA only.

Obtain human-approved CURRENT Electron donor screenshots for Memory definition/advanced/modal, Replicator and Diagnostics in durable retrievable location with dimensions and comparable safe state, OR approved exact reproducible capture procedure pinned to revision. Prior chat image shows one annotated isolated Memory/RBE view only; `/tmp/jr-shots` not durable. Confirm Electron commit delta at promotion. Include Windows-only ACL/services as explicit platform exceptions; never imitate with fake controls. Inspect visual layout, folder tabs, fields, scrolling, selection, OS.js window/desktop/ShadowRoot CSS leakage separately from runtime/COMMS behavior. Require paired screenshots, dimensions and discrepancy report. Missing baseline BLOCKED, real mismatch FAIL; JR does not fix code.

No production, enabled RBE listener, customer device, protocol change, ICC/CWAL modification, cutover/legacy deletion. QUEUED not actionable. Size 0/1/0/1/1=3.
