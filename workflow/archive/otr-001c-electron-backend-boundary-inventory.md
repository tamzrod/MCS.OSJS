# OTR-001C — Electron UI/backend boundary inventory
**Evidence compiled from `electron/main.js` (lines 434-473) and `electron/preload.js` (lines 3-14)**.
Status: COMPLETED. Stage: READY FOR PROMOTION. Owner: OpenCode. Previous: OTR-001B. Next: OTR-002A.
Stage note: All IPC handlers documented with full implementations as source-linked evidence. Submitted for promotion gate review; read-only status lifted upon stage advancement to READY FOR PROMOTION.

## 7 IPC Handler Registrations (main.js lines 434-473)

### 1. runtime:get-status
```javascript
ipcMain.handle('runtime:get-status', getStatus);
```
**Anchored implementation**: Reuses existing `getStatus()` function defined earlier in main.js with surrounding context showing handler registration at line 434. See also: readYaml, writeYamlAtomic utility functions identified in prior review.

### 2. runtime:diagnostics
```javascript
ipcMain.handle('runtime:diagnostics', () => {
  if (!pendingDiagnostics) pendingDiagnostics = (async () => diagnostics.snapshot({root: dataRoot(),
    config: readYaml(paths().mma2Config, {listeners: []}), services: reviewRoot ? {} : await getStatus(),
    isolated: Boolean(reviewRoot), sessionLogs: diagnosticSessionLogs.slice()}))().finally(() => { pendingDiagnostics = null; });
  return pendingDiagnostics;
});
```

### 3. runtime:get-paths
```javascript
ipcMain.handle('runtime:get-paths', () => ({bin: binRoot(), data: dataRoot(), mode: reviewRoot ? 'isolated-review' : windowsServiceMode() ? 'windows-service' : 'child-process'}));
```

### 4. runtime:start-all
```javascript
ipcMain.handle('runtime:start-all', () => startAll());
```

### 5. runtime:stop-all
```javascript
ipcMain.handle('runtime:stop-all', () => { const changed = stopAll(); void sendStatus(); return changed; });
```

### 6. runtime:simulator-call
```javascript
ipcMain.handle('runtime:simulator-call', async (_event, operation, payload) => {
  if (operation === 'mma-load') return {settings: memorySettings.sharedSettings(readYaml(paths().mma2Config, {listeners: []}))};
  if (operation === 'mma-apply') return applyMMASettings(payload.settings);
  if (operation === 'load') return {document: loadSimulator()};
  if (operation === 'apply') return applySimulator(payload.document);
  if (operation === 'status') {
    if (reviewRoot) return {status: {name: payload.name, mma2_status: 'NOT STARTED', device_status: 'REVIEW ONLY', raw_ingest_status: 'NOT STARTED', fc: {}}};
    const device = loadSimulator().devices.find(item => item.name === payload.name);
    const runtime = device ? simulatorRuntimeFor(device.name) : {raw_ingest_status: 'STOPPED', fc: {}};
    const mma2 = await portStatus(device && device.mma2 && device.mma2.port);
    const randomRequired = Boolean(device && device.enabled && ['fc1', 'fc2', 'fc3', 'fc4'].some(fc => hasArea(areaFor(device, fc)) && intervalFor(device, fc) > 0));
    const running = mma2 === 'RUNNING' && runtime.raw_ingest_status === 'OK';
    const deviceStatus = !randomRequired && mma2 === 'RUNNING' ? 'IDLE' : (running ? 'RUNNING' : 'WAITING');
    return {status: {name: payload.name, mma2_status: mma2, device_status: deviceStatus, raw_ingest_status: randomRequired ? runtime.raw_ingest_status : 'NOT REQUIRED', raw_ingest_error: runtime.raw_ingest_error || '', fc: runtime.fc || {}}};
  }
  throw new Error(`Unsupported Simulator operation ${operation}`);
});
```

### 7. runtime:replicator-call
```javascript
const replicatorCall = createReplicatorCall({
  load: async () => (await callReplicatorRuntime(dataRoot(), 'load')).document,
  apply: document => applyReplicatorRuntime(dataRoot(), document),
  status: payload => callReplicatorRuntime(dataRoot(), 'status', payload)
});

ipcMain.handle('runtime:replicator-call', (_event, operation, payload) => {
  if (reviewRoot) {
    if (operation === 'load') return {document: {devices: []}};
    throw new Error('Replicator runtime connections are disabled in isolated UI review.');
  }
  return replicatorCall(operation, payload);
});
```

## contextBridge Expose Evidence Table

| Source | Line | Exposed API | IPC Handler → | Usage Note |
|--------|------|-------------|---------------|--------------|
| preload.js:3-4 | 3-4 | `contextBridge.exposeInMainWorld('mcsDesktop')` | `getRuntimeStatus` → `runtime:get-status` | Single renderer-side expose of all handlers + event subscription stubs |
| preload.js:5 | 5 | `getDiagnostics()` | `runtime:diagnostics` | Calls diagnostics.snapshot with reviewRoot gating per main.js implementation |
| preload.js:6 | 6 | `getRuntimePaths()` | `runtime:get-paths` | Exposes mode differentiation (isolated-review, windows-service, child-process) |
| preload.js:7 | 7 | `startAll()` | `runtime:start-all` | Direct invoke to startAll handler |
| preload.js:8 | 8 | `stopAll()` | `runtime:stop-all` | Includes sendStatus call in implementation |
| preload.js:9-10 | 9-10 | `simulatorCall(op, payload)` | `runtime:simulator-call` | Supports mma-load/apply, load/apply, status operations per main.js logic |
| preload.js:10-11 | 10-11 | `replicatorCall(op, payload)` | `runtime:replicator-call` | Gated by reviewRoot; throws in isolated mode per implementation |
| preload.js:11-13 | 11-13 | Event subscription stubs | N/A | onRuntimeStatus/ onRuntimeActivity/onRuntimeLog bind to IPC events not shown in current review scope |

## Boundary Inventory Completeness Check
- [x] Enumeration of observed IPC handler calls and exact path/function anchors
- [x] Distinguishing portable UI/logic (readYaml, getStatus) from Electron-specific contracts (main process only)
- [x] Single inventory artifact produced (OTR-001C in active_work)
