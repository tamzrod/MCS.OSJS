# OTR-002B — OS.js Toolkit backend baseline

Status: QUEUED
Stage: DISCOVERY
Owner: OpenCode
Previous: OTR-002A
Next: OTR-003A
Size: I0 E1 B1 V1 D0 = 3

## Gate
Do not start unless `workflow/active_work/evidence/otr-002a-report.md` exists and handoff Current task is OTR-002B. If the gate fails, STOP.

## One outcome
Inventory existing OS.js Toolkit backend methods and sockets from source. Do not call them. Do not invent HTTP APIs.

## Allowed files to write
- `workflow/active_work/evidence/otr-002b-report.md`
- `handoff.md` Current task only, after the report exists
- `workflow/active_work/OTR_PROMOTION_QUEUE.md` 002B checkbox only

## Forbidden
Product edits. Service calls. Production YAML. ICC. Archive deletion. Starting OTR-003A in this invocation.

## Exact files to read (only these)
1. `OSJS/src/packages/MCSModbusToolkit/server.js`
2. `OSJS/src/packages/MCSModbusToolkit/memory-transport.js`
3. `OSJS/src/packages/MCSModbusToolkit/memory-contract.js`
4. `OSJS/src/packages/MCSModbusToolkit/replicator-transport.js`
5. `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js`
6. `OSJS/src/packages/MCSModbusToolkit/replicator-adapter.js`
7. `OSJS/src/packages/ModbusSimulator/server.js`
8. `OSJS/src/packages/ModbusReplicator/server.js`

If a file is missing, record MISSING and continue to the next listed file. Do not add unlisted files except one `git ls-files OSJS/src/packages/MCSModbusToolkit` listing if needed to confirm names.

## Numbered steps
1. Confirm branch is `opencode`.
2. Record HEAD.
3. Read the listed files.
4. Write the report with headings:
   - HEAD
   - Method names allowed for Memory/Simulator
   - Method names allowed for Replicator
   - Socket or transport path strings found in source
   - What index/server comments say is absent
   - Operations not found (create/delete/restart/HTTP) marked ABSENT if not in source
   - Verdict: SOURCE INVENTORY COMPLETE or BLOCKED
5. Use exact identifiers from source. Do not write `/OSJS/packages/toolkit/`.
6. Update handoff to OTR-003A only after COMPLETE.
7. STOP.

## Acceptance (max 3)
1. Allowed methods are listed with file anchors.
2. Transport/socket strings are quoted from source or marked ABSENT.
3. Unsupported operations are explicit.
