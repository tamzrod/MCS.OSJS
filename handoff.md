# Handoff — OS.js Toolkit Electron replica

## Autonomous routing authority
Human approved the OTR roadmap. OPERATION CWAL selects exactly ONE task. Read this file, then `workflow/active_work/OTR_PROMOTION_QUEUE.md`, then the packet named below. Execute that packet only. Report evidence. STOP. Do not ask the human what to do. Do not run a second task in the same invocation.

## Current task (ACTIVE)
OTR-002A — OpenCode, read-only DISCOVERY.
Packet: `workflow/active_work/otr-002a-osjs-ui-baseline.md`
Do that packet now.

## Do not treat these as done
`workflow/archive/otr-002b-baseline-inventory.md` and `workflow/archive/otr-003a-ui-parity-map-inventory.md` are INVALID. They invent paths that do not exist (`OSJS/packages/toolkit/app/views/`). Ignore them. Do not archive or delete active_work.

## Allowed writes for OTR-002A only
1. Create `workflow/active_work/evidence/otr-002a-report.md`.
2. Update only this file's Current task section to OTR-002B after a COMPLETE report exists.
3. Update only the OTR-002A/002B lines in `workflow/active_work/OTR_PROMOTION_QUEUE.md`.
4. Change only the Status line in `workflow/active_work/otr-002a-osjs-ui-baseline.md` to COMPLETE.
5. Change only the Status line in `workflow/active_work/otr-002b-osjs-backend-baseline.md` to ACTIVE.

Commit exactly those changed paths, push non-force to `origin/opencode`, verify the remote SHA, then
STOP. Activating OTR-002B does not authorize executing it in this invocation.

## Forbidden
- Do not delete `workflow/active_work/`
- Do not implement MMA2 CRUD or any product code
- Do not edit `electron/`, `simulator/`, `MMA2/`, `OSJS/` product files
- Do not write `simulator/devices.yaml`
- Do not edit `ICC/`
- Do not create new files named `operation c wal.md`
- Do not push anywhere except the packet-required non-force completion push to `origin/opencode`.
- Do not merge, force-push, create a PR, or switch branch.
- Do not start OTR-002B or any later task in this invocation

## After OTR-002A
Next eligible task is OTR-002B. Packet: `workflow/active_work/otr-002b-osjs-backend-baseline.md`. Only a later invocation may run it.

OTR-004 through OTR-014 parent files are NOT executable.
