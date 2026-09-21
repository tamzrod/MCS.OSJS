> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# UMIG-007A — VERIFY: Independent OS.js Deployment

Status: QUEUED — promoted 2026-09-18; wait for UMIG-007 PASS and all functional tab verification.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-007
Next: none — stop at this boundary; UMIG-008 cutover requires separate explicit human approval and promotion.

## Primary outcome
Prove Toolkit builds and runs without Electron installation or artifacts on a separate OS.js test target under Docker Compose.

## Verification action
Deploy only OS.js-owned source/assets and normal OS.js backend services into an isolated target with no Electron binary, node_modules, renderer or build output; run repository-native OS.js package/build/launch workflow using the Docker Compose image where safe. Inspect bundled paths/imports/network and IPC references for dependence on Windows Electron. Do not modify the operator's persistent Docker volume or live config.
Expected: Toolkit builds, launches and uses its own services; Electron source/package unchanged. Evidence: target details, exact commands/exit codes, bundle/import inspection, actual launch and changed-path report. If isolation or runtime unavailable, BLOCKED, never infer PASS.

## Non-scope
No product edits, shared renderer, pipeline change, Windows acceptance, launcher cutover or old app removal.

## Dependencies
UMIG-007 PASS, UMIG-004-V/005-V/006-V PASS and current JR packet. The operator's statement that Docker already works does not by itself establish Electron-free Toolkit integration.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
