const tabs = [...document.querySelectorAll('.tab')];
const panels = [...document.querySelectorAll('.panel')];
const log = document.getElementById('log');
const pulseTimers = new Map();

const setStatuses = value => {
  ['mma2', 'simulator', 'replicator'].forEach(key => {
    const node = document.getElementById(`status-${key}`);
    const status = value && value[key] || 'STOPPED';
    node.textContent = status;
    node.className = status === 'RUNNING' ? 'status-ok' : 'status-stop';
  });
};

const pulseActivity = ({process}) => {
  if (process !== 'mma2' && process !== 'simulator') return;
  const node = document.getElementById(`status-${process}`);
  if (!node || !node.classList.contains('status-ok')) return;

  node.classList.add('status-activity');
  const previous = pulseTimers.get(process);
  if (previous) clearTimeout(previous);
  pulseTimers.set(process, setTimeout(() => {
    node.classList.remove('status-activity');
    pulseTimers.delete(process);
  }, 110));
};

const appendLog = entry => {
  const line = `[${new Date().toLocaleTimeString()}] ${entry.process}/${entry.level}: ${entry.text}`;
  log.textContent += line.endsWith('\n') ? line : `${line}\n`;
  log.scrollTop = log.scrollHeight;
};

tabs.forEach(tab => tab.addEventListener('click', () => {
  tabs.forEach(node => node.classList.toggle('active', node === tab));
  panels.forEach(panel => panel.classList.toggle('active', panel.id === `panel-${tab.dataset.tab}`));
}));

document.getElementById('start-all').addEventListener('click', () => window.mcsDesktop.startAll());
document.getElementById('stop-all').addEventListener('click', () => window.mcsDesktop.stopAll());

window.mcsDesktop.onRuntimeStatus(setStatuses);
window.mcsDesktop.onRuntimeActivity(pulseActivity);
window.mcsDesktop.onRuntimeLog(appendLog);

Promise.all([
  window.mcsDesktop.getRuntimeStatus(),
  window.mcsDesktop.getRuntimePaths()
]).then(([status, paths]) => {
  setStatuses(status);
  document.getElementById('bin-path').textContent = paths.bin;
  document.getElementById('data-path').textContent = paths.data;
  document.getElementById('runtime-mode').textContent = paths.mode === 'windows-service'
    ? 'Windows services (NSSM)'
    : 'Electron child processes';

  if (paths.mode === 'windows-service') {
    const startButton = document.getElementById('start-all');
    const stopButton = document.getElementById('stop-all');
    startButton.disabled = true;
    stopButton.disabled = true;
    startButton.title = 'Installed backend runtimes are managed by Windows Services.';
    stopButton.title = 'Closing Electron does not stop installed Windows services.';
    appendLog({
      process: 'electron',
      level: 'info',
      text: 'Backend runtimes are managed by NSSM Windows services. Use Windows Services for manual start/stop operations.'
    });
  }
}).catch(error => appendLog({process: 'electron', level: 'error', text: error.message || String(error)}));
