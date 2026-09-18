# Handoff

## Current direction — 2026-09-18

Human priority: begin staged OS.js migration to one `MCS Modbus Toolkit`. Windows LED UI fixes, if needed, are separate work and must not block creating the independent OS.js placeholder package. This direction does NOT mean the entire Windows RLED runtime/installer verification passed.

ACTIVE: UMIG-002 — Scaffold One OS.js Toolkit Package. Only a placeholder source checkpoint is staged. Build/discovery and actual OS.js window launch are NOT VERIFIED; no implementation-complete or advancement claim.

PAUSED QUEUED: RREC-004 — Installed Package Proof; its hash, ProgramData and non-destructive repair gates remain unverified. RLED-003 through RLED-011 remain queued Windows acceptance/feature records; the human reports the LED UI is done, but no missing live-status verification is invented. Never auto-advance the parked Windows chain during OS.js migration.

Other QUEUED work remains unchanged: MEM-004 through MEM-008, REP-BLOCK-002 and REP-BLOCK-003. RREC-001 through RREC-003 are marked COMPLETE in Active Work with focused evidence; they still require separate verified archival reconciliation and are NOT the current task.

UMIG-001 and UMIG-003 through UMIG-009 remain in Planning. UMIG-001 donor snapshot approval is deferred until the independent shell is build/launch verified; Windows LED fixes are not a blanket prerequisite for the placeholder. No other UMIG task was promoted by this kickoff.

## Authorized UMIG-002 boundary

Establish an OS.js-only package with one placeholder window. Do not touch Electron files, copy the renderer, invoke the backends, change old launchers/Start menu, delete old packages, or modify user data. Preserve OS.js desktop, Start menu, taskbar, clock and window management. Future cutover targets one MCS Modbus Toolkit desktop icon; old packages remain fallback until live Memory/Replicator/Diagnostics, visual parity and independent deployment verification pass.

`ModbusSimulator` and `ModbusReplicator` OS.js `server.js` own bridges to live Go runtimes. Move any needed bridge into Toolkit in later UMIG-004/005 before UI retirement. MMA2, Go services, config, memory and user data are not decommissioned.

## Verification and continuation

Required next evidence: `cd OSJS && npm run build:local-packages && npm run package:discover`, independent of Electron, followed by a real OS.js `MCSModbusToolkit` launch and one-window observation. The GitHub-only source checkpoint does not prove any of these. Do not use `electron/build-and-push.ps1` for OS.js build: it performs `git pull` and Windows packaging.

`ICC/INDEX.md` baseline remains stale. A bounded BLACK SHEEP WALL refresh of Planning and Active Work must be completed before JR execution or further promotion. Only BLACK SHEEP WALL may edit ICC; do not fabricate a clean working tree for the Windows checkout. The current handoff has no JR TEST TASK packet; Operation CWAL must report BLOCKED until an explicit packet is authored.

## Historical evidence boundaries

RLED-001/002 are archived for their focused bridge/status gates, not end-to-end installed acceptance. RREC-001/002/003 have recorded focused installer/runtime/status evidence in their task files. The Electron COMMS renderer exists in source; the committed Go status shape still needs review against the renderer's per-layer LED contract. Do not claim the user's installed Windows application was tested by this source-only review.
