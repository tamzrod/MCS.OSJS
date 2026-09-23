# OS.js desktop shell and Toolkit wiring

Parent: L0-project
Zoom Out: [L0-project](L0-project.md)
Zoom In: [osjs-toolkit-advanced-status](osjs-toolkit-advanced-status.md), [osjs-shared-mma-settings](osjs-shared-mma-settings.md)
Connectors: [simulator-device-config](simulator-device-config.md) (Memory/Simulator contract); [replicator](replicator.md) (Replicator contract). These are explicit cross-tree routes, not automatic context imports.

## Scope and source dependencies

This node covers the OS.js shell, Toolkit package mounting and routing, not the entire Simulator or Replicator runtime. Source dependencies: `OSJS/src/client/config.js`, `OSJS/src/client/index.js`, `OSJS/src/client/providers/nameless-app-shortcuts.js`, `OSJS/src/packages/MCSModbusToolkit/index.js`, `OSJS/src/packages/MCSModbusToolkit/metadata.json`, `OSJS/src/packages/MCSModbusToolkit/package.json`, `OSJS/src/packages/MCSModbusToolkit/toolkit-renderer.js`, `OSJS/src/packages/MCSModbusToolkit/server.js`, `OSJS/src/server/index.js`, `OSJS/src/server/config.js`, `OSJS/package.json`.

## Verified snapshot and limits

At remote `main` HEAD `45cd3d2d81831306c6943f844e49b7deccbfc5ad`, inspected `MCSModbusToolkit/index.js` imports and mounts a single Toolkit window with Memory and Replicator transports/contracts/editors, Diagnostics editor, shared MMA settings and renderer. Window destroy closes transports and disposes editor components. Its Diagnostics uses existing Memory/Replicator contracts rather than a third service API. This replaces the obsolete claim that the Toolkit is a disconnected placeholder. This review did not execute the app, inspect the local overlay, or certify live Ubuntu/Docker behavior.

The OS.js desktop shell and Windows Electron application are distinct deployments. Do not import Windows named-pipe or filesystem assumptions into the Linux Docker architecture. For advanced UI/status facts, zoom into the Toolkit child; for shared MMA settings transaction boundaries, zoom into the shared-settings child. Refer to actual current sources for exact relay paths and transport compatibility before coding.

## Validity

Source snapshot: remote `main` commit `45cd3d2d81831306c6943f844e49b7deccbfc5ad`; current checkout may differ. Dependency changes or missing coverage require a bounded Black Sheep Wall refresh. Do not copy historical build/acceptance claims into this node without independently verified evidence. Do not refresh unrelated context just to update this shell node.
