# UMIG-007A — VERIFY: Independent OS.js Deployment

Status: PLANNED / BLOCKED — promotion required.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-007
Next: UMIG-008

## Primary outcome
Prove Toolkit builds and runs without Electron installation or artifacts on a separate OS.js test target.

## Verification action
Deploy only OS.js-owned source/assets and normal OS.js backend services into an isolated target with no Electron binary, node_modules, renderer or build output; run repository-native OS.js package/build/launch workflow. Inspect bundled paths/imports/network and IPC references for dependence on Windows Electron.
Expected: Toolkit builds, launches and uses its own services; Electron source/package unchanged. Evidence: target details, commands/exit codes, bundle/import inspection, actual launch and changed-path report. If isolation or runtime unavailable, BLOCKED, never infer PASS.

## Non-scope
No product edits, shared renderer, pipeline change, Windows acceptance, cutover or old app removal.

## Dependencies
UMIG-007 PASS, UMIG-004-V/005-V/006-V PASS, explicit promotion and current JR packet.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
