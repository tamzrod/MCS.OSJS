# UMIG-009-T — TEST: Post-Retirement Package Build

Status: PLANNED / BLOCKED — retirement approval/promotion required.
Stage / owner: TEST / OpenHands (JR)
Previous: UMIG-009
Next: UMIG-009-V

## Primary outcome
Verify Toolkit builds/discovers after removal of superseded UI packages.

## Instruction / expected / evidence
In disposable checkout, run `cd OSJS && npm run build:local-packages && npm run package:discover` plus exact discovery inspection in the JR packet. Expected: only Toolkit is registered as an MCS Modbus UI; former Simulator/Replicator packages absent; Toolkit's own server bridges included; no Electron build. Return output/exit, discovered registry, bundle paths, changed-file list and HEAD. Failure is FAIL, unavailable environment BLOCKED.

## Non-scope
No runtime acceptance, deletion of user data, source fixes or further removal.

## Dependencies
UMIG-009 code checkpoint, human retirement approval and current JR packet.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
