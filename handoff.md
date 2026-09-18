# Handoff

## Current direction — 2026-09-18

Human priority: staged OS.js replacement with one `MCS Modbus Toolkit`. Windows LED repairs, if any, are a separate track and do not block initial OS.js package staging. The OS.js desktop, Start menu, taskbar and clock remain; final default desktop has exactly one Toolkit icon. Legacy UI packages are removed only after Toolkit testing and verification. MMA2, Go services, shared memory and user configuration are never decommissioned by this UI migration.

## Roles and separate gates

ChatGPT = coding agent: implements one CODE microtask, reads back exact source change, commits a source-only checkpoint, writes current OpenHands test packet, reviews test evidence and controls task advancement. OpenHands = JR test runner: TEST builds/fixtures and VERIFY actual runtime/UI acceptance from separate tasks using `operation cwal.md`; no source edits, bug fixes, task promotion, archival or ICC edits. A failed/blocked test stops advancement and returns to a separately authorized coding fix. Source presence is not test PASS; unit tests are not rendered runtime PASS.

## Current authorized sequence

ACTIVE: `UMIG-002` CODE — the independent OS.js Toolkit package source is staged; source-only handoff/readback is pending. No build/discovery/UI launch has been performed or claimed.
QUEUED: `UMIG-002-T` OpenHands package build/discovery TEST; `UMIG-002-V` OpenHands one-window launch VERIFY. Ordered links UMIG-002 → UMIG-002-T → UMIG-002-V → STOP. These two test stages must not be executed until each is ACTIVE and has its own current JR TEST TASK packet.
PLANNED: UMIG-001 donor approval, then UMIG-003 through UMIG-009 and their separately planned TEST/VERIFY stages. They require separate human promotion. No Electron donor SHA has been approved.
PAUSED QUEUED: RREC-004, RLED-003 through RLED-011, MEM-004 through MEM-008 and REP-BLOCK-002/003. RREC-001/002/003 remain COMPLETE in Active Work, pending archival reconciliation; they are not currently executable.

## UMIG-002 scope

Only author and inspect the self-owned `OSJS/src/packages/MCSModbusToolkit` placeholder source. Do not change Electron files, backend services, existing launchers, legacy application packages or user data. Existing OS.js Simulator and Replicator server bridges must be migrated before those UI packages are retired in a future approved code task.

## Test packet status and safety

NO CURRENT `JR TEST TASK` yet. The coding agent must finish the UMIG-002 source checkpoint, update Active Work/handoff atomically to UMIG-002-T ACTIVE, then supply its exact command, expected result and evidence. OpenHands must not start a test based on this general handoff paragraph. After TEST PASS, coding agent may activate UMIG-002-V with a separate real GUI observation packet. OpenHands reports raw PASS/FAIL/BLOCKED evidence only and stops. Never run Electron build-and-push or perform destructive config/service changes as a substitute.

`ICC/INDEX.md` and affected Planning/Active Work contexts are stale and need a bounded BLACK SHEEP WALL refresh before OpenHands execution or further promotion. Only BLACK SHEEP WALL edits ICC; no claimed clean Windows working tree or unobserved test result.

## Evidence limits

Prior Windows `RLED` source/UI appearance was accepted by the human, not independently verified live end-to-end here. RREC-004 installer repair/hash/ProgramData evidence remains incomplete. Do not attribute Windows installation/COMMS success to the OS.js migration or these source-only edits.
