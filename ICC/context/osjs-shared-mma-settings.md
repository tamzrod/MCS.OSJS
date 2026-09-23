# OS.js — shared MMA settings

Parent: [osjs-shell](osjs-shell.md)
Zoom Out: [osjs-shell](osjs-shell.md)
Zoom In: none
Connector: [osjs-toolkit-advanced-status](osjs-toolkit-advanced-status.md) for editor/status integration; it is not an automatic import.

## Boundary and established scoped facts

Simulator owns authenticated shared MMA settings load/apply through a mutation mutex and writer lock. Allowed patch keys are `rbe`, `access_events` and `debug`; null removes a field while omission preserves it. Compare the on-disk SHA-256 revision, validate the entire candidate, and atomically replace only the effective config. Do not rewrite device definitions or ownership entries. A matching restart acknowledgement is awaited; a post-commit acknowledgement failure remains explicitly committed-but-uncertain rather than silently rolled back. Supervisor acknowledgement is not proof that every listener bound successfully.

The Toolkit relay allowlists shared operations on the Simulator route; the shared dialog retains separate drafts and requires reload after conflict or ambiguous results. Both advanced editors expose the same shared dialog and its last saved RBE port; output requires explicit enable and save. Deploy updated Simulator and Toolkit together. These are the historical scoped findings, not permission for an alternate direct filesystem or Windows service route.

## Source dependencies and freshness

`simulator/shared_settings.go`, `simulator/shared_settings_test.go`, `simulator/runtime_server.go`, `OSJS/src/packages/MCSModbusToolkit/shared-settings.js`, `OSJS/src/packages/MCSModbusToolkit/server.js`, `OSJS/src/packages/MCSModbusToolkit/memory-contract.js`, `OSJS/src/packages/MCSModbusToolkit/memory-transport.js`, `OSJS/tests/toolkit-shared-settings.test.js`.

The previous node's audited checkpoint was `0b7ee17` plus per-file Git blob fingerprints, preserved in Git history. These files were not all freshly audited for remote `main` `45cd3d2d81831306c6943f844e49b7deccbfc5ad`; verify only relevant dependency deltas and local overlay before considering the node current. Previous local tests were not live Ubuntu/operator acceptance, and no new tests were run for this context repair.
