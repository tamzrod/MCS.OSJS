# OTR-003B — Backend contract map

Status: QUEUED
Stage: DESIGN
Owner: OpenCode
Previous: OTR-003A
Next: OTR-004 children (not this invocation)
Size: I0 E0 B1 V1 D1 = 3

## Gate
Need `workflow/active_work/evidence/otr-002b-report.md` and `workflow/active_work/evidence/otr-003a-map.md`. If either is missing, STOP.

Before reading them, confirm branch `opencode`, record HEAD, require empty status, and require both
`origin/opencode` and `git ls-remote origin refs/heads/opencode` to equal HEAD. Otherwise BLOCKED.

## One outcome
A matrix of load / validate / save / restart / status for Memory(Simulator), Replicator, and MMA as found in source. Mark GAP when not found. Do not invent endpoints.

## Allowed writes
- `workflow/active_work/evidence/otr-003b-map.md`
- `handoff.md` Current task after COMPLETE. Set it to `NONE — wait for OTR-004A child packet`.
- queue 003B line only
- this file Status line only

## Read
- `workflow/archive/otr-001c-electron-backend-boundary-inventory.md`
- `workflow/active_work/evidence/otr-002b-report.md`
- `electron/preload.js`
- `OSJS/src/packages/MCSModbusToolkit/server.js`
- `OSJS/src/packages/MCSModbusToolkit/memory-contract.js`
- `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js`

## Matrix columns
Runtime | Operation | Electron symbol | OS.js symbol | Present/GAP | Notes

Runtimes: Memory/Simulator, Replicator, MMA.
Operations: load, validate, save/apply, restart, status.

## STOP
Do not create OTR-004A. Do not edit product files.

On COMPLETE, mark 003B COMPLETE in this file and queue, set handoff to the stated NONE value, verify
only the four allowed paths changed, run diff checks, stage only them, and commit
`OTR-003B: record backend contract map`. Confirm remote `opencode` still equals the pre-task HEAD,
push `git push origin HEAD:opencode`, verify the new remote SHA, report it, and STOP.

A BLOCKED report may be committed/pushed alone as `OTR-003B: record blocked backend contract map`
without status advancement. Failed delivery is INCOMPLETE.
