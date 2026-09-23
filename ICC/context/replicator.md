# Replicator — runtime and transport

Parent: [L0-project](L0-project.md)
Zoom Out: [L0-project](L0-project.md)
Zoom In: none
Connectors: [osjs-shell](osjs-shell.md) (Toolkit relay), [simulator-device-config](simulator-device-config.md) (shared MMA2 reservation). These are navigation hints, not automatic imports.
Source dependencies: `replicator/runtime_api.go`, `replicator/cmd/modbus-replicator-runtime/main.go`, `replicator/go.mod`, `OSJS/src/packages/ModbusReplicator/server.js`, `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js`, `OSJS/src/packages/MCSModbusToolkit/replicator-transport.js`, `electron/replicator-runtime.js`.

## Scoped historical facts

The earlier audited Windows runtime transition recorded `replicator/runtime_api.go` returning `\\\\.\\pipe\\mcs-modbus-replicator` and `main.go` listening via `go-winio`. The older OS.js `ModbusReplicator/server.js` still targeted `$OSJS_DATA_DIR/run/modbus-replicator.sock` at that time. This documented a **historical transport mismatch**, not a license to assume it remains unresolved after later Toolkit changes or to copy Windows pipes into Linux Docker. Check the smallest affected source files in the current branch before designing a fix.

Replicator handles data acquisition/replication and shares MMA2 reservation constraints with Simulator. The current OS.js Toolkit source wires Replicator contracts, transport and editor from its single window; this is source wiring, not proof of a live backend connection.

## Task and verification separation

Historical REP-BLOCK and RLED task claims from the old node are not product architecture and are omitted here. Current task status must come from `handoff.md` and canonical workflow packets, summarized only in [active-work](active-work.md). No test, Linux build, Windows installer acceptance or deployment PASS is inferred from this node.

The historical runtime facts above were cached at `ee19b8a`; the Toolkit import/wiring was inspected in remote `main` at `45cd3d2d81831306c6943f844e49b7deccbfc5ad`. Other dependencies and the local uncommitted overlay were **not** re-audited. Delta-check those exact dependencies before consuming detailed transport claims.
