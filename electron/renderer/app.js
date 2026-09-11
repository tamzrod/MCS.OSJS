const tabs = [...document.querySelectorAll('.tab')];
const panels = [...document.querySelectorAll('.panel')];
const log = document.getElementById('log');

const setStatuses = value => {
  ['mma2', 'simulator', 'replicator'].forEach(key => {
    const node = document.getElementById(`status-${key}`);
    const status = value && value[key] || 'STOPPED';
    node.textContent = status;
    node.className = status === 'RUNNING' ? 'status-ok' : 'status-stop';
  });
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
window.mcsDesktop.onRuntimeLog(appendLog);

Promise.all([
  window.mcsDesktop.getRuntimeStatus(),
  window.mcsDesktop.getRuntimePaths()
]).then(([status, paths]) => {
  setStatuses(status);
  document.getElementById('bin-path').textContent = paths.bin;
  document.getElementById('data-path').textContent = paths.data;
}).catch(error => appendLog({process: 'electron', level: 'error', text: error.message || String(error)}));
