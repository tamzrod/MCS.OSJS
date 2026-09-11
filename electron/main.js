const {app, BrowserWindow, ipcMain} = require('electron');
const {spawn} = require('child_process');
const fs = require('fs');
const path = require('path');

let mainWindow = null;
const children = new Map();

const binRoot = () => app.isPackaged
  ? path.join(process.resourcesPath, 'bin')
  : path.join(__dirname, 'bin');

const dataRoot = () => path.join(app.getPath('userData'), 'runtime');

const executableName = name => process.platform === 'win32' ? `${name}.exe` : name;

const processSpecs = [
  {key: 'mma2', file: executableName('mma2')},
  {key: 'simulator', file: executableName('modbus-simulator-runtime')},
  {key: 'replicator', file: executableName('modbus-replicator-runtime')}
];

const status = key => {
  const child = children.get(key);
  return child && !child.killed ? 'RUNNING' : 'STOPPED';
};

const sendStatus = () => {
  if (!mainWindow || mainWindow.isDestroyed()) return;
  mainWindow.webContents.send('runtime:status', Object.fromEntries(processSpecs.map(spec => [spec.key, status(spec.key)])));
};

const startProcess = spec => {
  if (children.has(spec.key)) return;
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
    sendStatus();
  });
  sendStatus();
};

const startAll = () => processSpecs.forEach(startProcess);

const stopAll = () => {
  for (const [key, child] of children.entries()) {
    try {
      child.kill();
    } catch (_) {}
    children.delete(key);
  }
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

ipcMain.handle('runtime:get-status', () => Object.fromEntries(processSpecs.map(spec => [spec.key, status(spec.key)])));
ipcMain.handle('runtime:get-paths', () => ({bin: binRoot(), data: dataRoot()}));
ipcMain.handle('runtime:start-all', () => { startAll(); return true; });
ipcMain.handle('runtime:stop-all', () => { stopAll(); sendStatus(); return true; });

app.whenReady().then(() => {
  createWindow();
  startAll();
  app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

app.on('before-quit', stopAll);
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') app.quit();
});
