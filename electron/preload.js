const {contextBridge, ipcRenderer} = require('electron');

contextBridge.exposeInMainWorld('mcsDesktop', {
  getRuntimeStatus: () => ipcRenderer.invoke('runtime:get-status'),
  getRuntimePaths: () => ipcRenderer.invoke('runtime:get-paths'),
  startAll: () => ipcRenderer.invoke('runtime:start-all'),
  stopAll: () => ipcRenderer.invoke('runtime:stop-all'),
  onRuntimeStatus: callback => ipcRenderer.on('runtime:status', (_event, value) => callback(value)),
  onRuntimeLog: callback => ipcRenderer.on('runtime:log', (_event, value) => callback(value))
});
