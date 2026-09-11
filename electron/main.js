const {app, BrowserWindow, ipcMain} = require('electron');
const {spawn, execFile} = require('child_process');
const fs = require('fs');
const path = require('path');

let mainWindow = null;
let statusTimer = null;
const children = new Map();

const windowsServiceMode = () => app.isPackaged && process.platform === 'win32';

const binRoot = () => app.isPackaged
  ? path.join(process.resourcesPath, 'bin')
  : path.join(__dirname, 'bin');

const dataRoot = () => windowsServiceMode()
  ? path.join(process.env.ProgramData || 'C:\\ProgramData', 'MCS Modbus Toolkit', 'runtime')
  : path.join(app.getPath('userData'), 'runtime');

const executableName = name => process.platform === 'win32' ? `${name}.exe` : name;

const processSpecs = [
  {key: 'mma2', file: executableName('mma2'), service: 'MCS-MMA2'},
  {key: 'simulator', file: executableName('modbus-simulator-runtime'), service: 'MCS-Simulator'},
  {key: 'replicator', file: executableName('modbus-replicator-runtime'), service: 'MCS-Replicator'}
];

const childStatus = key => {
  const child = children.get(key);
  return child && !child.killed ? 'RUNNING' : 'STOPPED';
};

const queryWindowsService = service => new Promise(resolve => {
  execFile('sc.exe', ['query', service], {windowsHide: true}, (error, stdout = '') => {
    if (error) {
      resolve('NOT INSTALLED');
      return;
    }
    resolve(/STATE\s*:\s*\d+\s+RUNNING/i.test(stdout) ? 'RUNNING' : 'STOPPED');
  });
});

const getStatus = async () => {
  if (windowsServiceMode()) {
    const values = await Promise.all(processSpecs.map(async spec => [spec.key, await queryWindowsService(spec.service)]));
    return Object.fromEntries(values);
  }
  return Object.fromEntries(processSpecs.map(spec => [spec.key, childStatus(spec.key)]));
};

const sendStatus = async () => {
  if (!mainWindow || mainWindow.isDestroyed()) return;
  mainWindow.webContents.send('runtime:status', await getStatus());
};

const startProcess = spec => {
  if (windowsServiceMode() || children.has(spec.key)) return;
  const file = path.join(binRoot(), spec.file);
  if (!fs.existsSync(file)) return;

  fs.mkdirSync(dataRoot(), {recursive: true});
  const child = spawn(file, [], {
    cwd: dataRoot(),
    windowsHide: true,
    env: {
      ...process.env,
      MCS_DATA_ROOT: dataRoot()
    },
    stdio: ['ignore', 'pipe', 'pipe']
  });

  children.set(spec.key, child);
  const forward = (stream, level) => stream && stream.on('data', chunk => {
    if (mainWindow && !mainWindow.isDestroyed()) {
      mainWindow.webContents.send('runtime:log', {
        process: spec.key,
        level,
        text: String(chunk)
      });
    }
  });
  forward(child.stdout, 'stdout');
  forward(child.stderr, 'stderr');

  child.on('exit', code => {
    children.delete(spec.key);
    if (mainWindow && !mainWindow.isDestroyed()) {
      mainWindow.webContents.send('runtime:log', {
        process: spec.key,
        level: 'exit',
        text: `${spec.key} exited with code ${code}`
      });
    }
    void sendStatus();
  });
  void sendStatus();
};

const startAll = () => {
  if (windowsServiceMode()) return false;
  processSpecs.forEach(startProcess);
  return true;
};

const stopAll = () => {
  if (windowsServiceMode()) return false;
  for (const [key, child] of children.entries()) {
    try {
      child.kill();
    } catch (_) {}
    children.delete(key);
  }
  return true;
};

const createWindow = () => {
  mainWindow = new BrowserWindow({
    width: 1120,
    height: 760,
    minWidth: 900,
    minHeight: 600,
    title: 'MCS Modbus Toolkit',
    backgroundColor: '#c0c0c0',
    autoHideMenuBar: true,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
      contextIsolation: true,
      nodeIntegration: false
    }
  });
  mainWindow.loadFile(path.join(__dirname, 'renderer', 'index.html'));
  mainWindow.on('closed', () => { mainWindow = null; });
};

ipcMain.handle('runtime:get-status', getStatus);
ipcMain.handle('runtime:get-paths', () => ({
  bin: binRoot(),
  data: dataRoot(),
  mode: windowsServiceMode() ? 'windows-service' : 'child-process'
}));
ipcMain.handle('runtime:start-all', () => startAll());
ipcMain.handle('runtime:stop-all', () => {
  const changed = stopAll();
  void sendStatus();
  return changed;
});

app.whenReady().then(() => {
  createWindow();
  if (windowsServiceMode()) {
    statusTimer = setInterval(() => void sendStatus(), 2000);
    void sendStatus();
  } else {
    startAll();
  }
  app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

app.on('before-quit', () => {
  if (statusTimer) clearInterval(statusTimer);
  if (!windowsServiceMode()) stopAll();
});
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') app.quit();
});
