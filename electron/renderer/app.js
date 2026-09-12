const tabs = [...document.querySelectorAll('.tab')];
const panels = [...document.querySelectorAll('.panel')];
const log = document.getElementById('log');
const pulseTimers = new Map();
let runtimePaths = null;

const clone = value => JSON.parse(JSON.stringify(value));
const numberValue = value => value === '' ? 0 : Number(value);
const FC_KEYS = ['fc1', 'fc2', 'fc3', 'fc4'];
const FC_LABELS = {fc1: 'Coils (FC1)', fc2: 'Discrete Inputs (FC2)', fc3: 'Holding Registers (FC3)', fc4: 'Input Registers (FC4)'};

const h = (tag, className = '', text) => {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (text !== undefined) node.textContent = text;
  return node;
};

const actionButton = (label, action, className = '') => {
  const node = h('button', `tool-button ${className}`.trim(), label);
  node.type = 'button';
  node.dataset.action = action;
  return node;
};

const field = (label, value, settings, onInput) => {
  const wrapper = h('label', settings.className || 'tool-field');
  wrapper.appendChild(h('span', 'tool-field-label', label));
  const input = document.createElement(settings.multiline ? 'textarea' : 'input');
  if (!settings.multiline) input.type = settings.type || 'number';
  input.value = value ?? '';
  input.readOnly = Boolean(settings.readOnly);
  ['min', 'max', 'step', 'placeholder'].forEach(key => {
    if (settings[key] !== undefined) input[key] = settings[key];
  });
  input.addEventListener('input', event => onInput(event.target.value));
  wrapper.appendChild(input);
  return wrapper;
};

const isEditableControl = target => Boolean(target.closest('input, select, textarea, option, label'));

const checkboxField = (label, checked, onChange) => {
  const wrapper = h('label', 'tool-checkbox');
  const input = document.createElement('input');
  input.type = 'checkbox';
  input.checked = Boolean(checked);
  input.addEventListener('change', event => onChange(event.target.checked));
  wrapper.append(input, h('span', '', label));
  return wrapper;
};

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

const formatTime = value => {
  if (!value) return '-';
  const parsed = new Date(value);
  return Number.isNaN(parsed.getTime()) ? value : parsed.toLocaleString();
};

const addressRange = area => {
  const start = Number(area && area.start) || 0;
  const count = Number(area && area.count) || 0;
  if (!count) return 'Unused';
  const end = start + count - 1;
  return start === end ? String(start) : `${start}-${end}`;
};

const simulatorBlankDevice = sequence => ({
  name: `Sim-PLC-${sequence}`,
  enabled: true,
  mma2: {
    port: 5020,
    unit_id: 1,
    fc1: {start: 0, count: 16},
    fc2: {start: 0, count: 16},
    fc3: {start: 0, count: 16},
    fc4: {start: 0, count: 16}
  },
  random_runtime: {
    fc1_interval_ms: 1000,
    fc2_interval_ms: 1000,
    fc3_interval_ms: 1000,
    fc4_interval_ms: 1000
  }
});

const validateSimulator = device => {
  if (!device.name || !device.name.trim()) return 'Name is required.';
  if (!device.mma2.port || device.mma2.port < 1 || device.mma2.port > 65535) return 'Listen Port must be between 1 and 65535.';
  if (device.mma2.unit_id < 0 || device.mma2.unit_id > 255) return 'Unit ID must be between 0 and 255.';
  for (const fc of FC_KEYS) {
    const area = device.mma2[fc];
    if (area.start < 0 || area.start > 65535 || area.count < 0 || area.count > 65535) return `${FC_LABELS[fc]} values are outside the 16-bit address range.`;
    if (area.start + area.count > 65536) return `${FC_LABELS[fc]} Start + Count exceeds the 16-bit address space.`;
    if (area.count > 0 && device.random_runtime[`${fc}_interval_ms`] < 1) return `${FC_LABELS[fc]} Randomize Every must be greater than zero.`;
  }
  return null;
};

const normalizeSimulatorDocument = value => ({devices: Array.isArray(value && value.devices) ? value.devices : []});

const simulatorState = {
  document: {devices: []},
  persisted: {devices: []},
  selected: null,
  loading: true,
  saving: false,
  message: 'Loading Simulator definitions...',
  error: false,
  runtimeStatus: null
};

const selectedSimulator = () => simulatorState.selected === null ? null : simulatorState.document.devices[simulatorState.selected];

const renderSimulator = () => {
  const root = document.getElementById('simulator-root');
  root.replaceChildren();
  const shell = h('div', 'tool-layout');
  const sidebar = h('aside', 'tool-sidebar');
  sidebar.appendChild(h('h2', '', 'Simulator Devices'));
  const listActions = h('div', 'tool-actions');
  listActions.append(actionButton('Add', 'sim-add'), actionButton('Duplicate', 'sim-duplicate'), actionButton('Delete', 'sim-delete', 'tool-danger'));
  sidebar.appendChild(listActions);
  const list = h('div', 'tool-list');
  simulatorState.document.devices.forEach((device, index) => {
    const row = actionButton('', 'sim-select', index === simulatorState.selected ? 'tool-row selected' : 'tool-row');
    row.dataset.index = String(index);
    row.append(h('strong', '', device.name || 'Unnamed device'), h('span', '', `Port ${device.mma2.port} / Unit ${device.mma2.unit_id}${device.enabled ? '' : ' / Disabled'}`));
    list.appendChild(row);
  });
  if (!list.children.length) list.appendChild(h('div', 'tool-empty', simulatorState.loading ? 'Loading...' : 'No devices configured.'));
  sidebar.appendChild(list);
  shell.appendChild(sidebar);

  const editor = h('section', 'tool-editor');
  editor.appendChild(h('h2', '', 'Device Definition'));
  const device = selectedSimulator();
  if (!device) {
    editor.appendChild(h('div', 'tool-empty', 'Select a device or choose Add.'));
  } else {
    const runtime = h('div', 'runtime-row');
    const status = simulatorState.runtimeStatus;
    runtime.append(h('span', '', 'MMA2:'), h('strong', '', status && status.mma2_status || '-'), h('span', '', 'Simulator:'), h('strong', '', status && status.device_status || '-'));
    editor.appendChild(runtime);
    const identity = h('div', 'tool-grid');
    identity.append(field('Name', device.name, {type: 'text'}, value => { device.name = value; }));
    identity.append(checkboxField('Enabled', device.enabled, value => { device.enabled = value; }));
    identity.append(field('Listen Port', device.mma2.port, {min: 1, max: 65535}, value => { device.mma2.port = numberValue(value); }));
    identity.append(field('Unit ID', device.mma2.unit_id, {min: 0, max: 255}, value => { device.mma2.unit_id = numberValue(value); }));
    editor.appendChild(identity);

    const table = h('div', 'fc-table');
    const header = h('div', 'fc-row fc-header');
    ['Function', 'Start', 'Count', 'Randomize Every (ms)', 'Address Range'].forEach(label => header.appendChild(h('span', '', label)));
    table.appendChild(header);
    FC_KEYS.forEach(fc => {
      const row = h('div', 'fc-row');
      const area = device.mma2[fc];
      const range = h('span', 'range-cell', addressRange(area));
      row.appendChild(h('strong', '', FC_LABELS[fc]));
      row.appendChild(field('Start', area.start, {min: 0, max: 65535, className: 'cell-field'}, value => { area.start = numberValue(value); range.textContent = addressRange(area); }));
      row.appendChild(field('Count', area.count, {min: 0, max: 65535, className: 'cell-field'}, value => { area.count = numberValue(value); range.textContent = addressRange(area); }));
      row.appendChild(field('Interval', device.random_runtime[`${fc}_interval_ms`], {min: 1, step: 1, className: 'cell-field'}, value => { device.random_runtime[`${fc}_interval_ms`] = numberValue(value); }));
      row.appendChild(range);
      table.appendChild(row);
    });
    editor.appendChild(table);
    const validation = validateSimulator(device);
    if (validation) editor.appendChild(h('div', 'tool-validation', validation));
    const actions = h('div', 'editor-actions');
    const save = actionButton(simulatorState.saving ? 'Saving...' : 'Save & Apply', 'sim-save', 'tool-primary');
    save.disabled = Boolean(validation) || simulatorState.saving;
    actions.append(save, actionButton('Discard', 'sim-discard'));
    editor.appendChild(actions);
  }
  editor.appendChild(h('div', `tool-status${simulatorState.error ? ' error' : ''}`, simulatorState.message));
  shell.appendChild(editor);
  root.appendChild(shell);
};

const loadSimulator = async () => {
  simulatorState.loading = true;
  simulatorState.message = 'Loading Simulator definitions...';
  simulatorState.error = false;
  renderSimulator();
  try {
    const result = await window.mcsDesktop.simulatorCall('load', {});
    const documentValue = normalizeSimulatorDocument(result.document);
    simulatorState.document = clone(documentValue);
    simulatorState.persisted = clone(documentValue);
    simulatorState.selected = documentValue.devices.length ? 0 : null;
    simulatorState.message = documentValue.devices.length ? 'Canonical Simulator definitions loaded.' : 'No devices configured. Choose Add to begin.';
    simulatorState.error = false;
  } catch (error) {
    simulatorState.message = error.message || String(error);
    simulatorState.error = true;
  } finally {
    simulatorState.loading = false;
    renderSimulator();
  }
};

const pollSimulator = async () => {
  const device = selectedSimulator();
  if (!device || simulatorState.saving) return;
  try {
    const result = await window.mcsDesktop.simulatorCall('status', {name: device.name});
    simulatorState.runtimeStatus = result.status || result;
    renderSimulator();
  } catch (_) {}
};

const saveSimulator = async () => {
  const validation = simulatorState.document.devices.map(validateSimulator).find(Boolean);
  if (validation) {
    simulatorState.message = validation;
    simulatorState.error = true;
    renderSimulator();
    return;
  }
  simulatorState.saving = true;
  simulatorState.message = 'Saving Simulator definitions and applying runtime changes...';
  simulatorState.error = false;
  renderSimulator();
  try {
    const result = await window.mcsDesktop.simulatorCall('apply', {document: clone(simulatorState.document)});
    const documentValue = normalizeSimulatorDocument(result.document);
    simulatorState.document = clone(documentValue);
    simulatorState.persisted = clone(documentValue);
    simulatorState.message = `${result.message || 'Simulator settings applied.'} ${result.completed_at ? `Completed at ${formatTime(result.completed_at)}.` : ''}`;
    simulatorState.error = false;
  } catch (error) {
    simulatorState.message = `Save & Apply failed: ${error.message || error}`;
    simulatorState.error = true;
  } finally {
    simulatorState.saving = false;
    renderSimulator();
    pollSimulator();
  }
};

const blankBlock = () => ({function: 3, start: 0, count: 16, scan_rate_ms: 1000});
const blankReplicator = sequence => ({
  name: `Rep-PLC-${sequence}`,
  enabled: true,
  endpoint: '127.0.0.1:5020',
  unit_id: 1,
  pull_blocks: [blankBlock()],
  destination: {port: 5021, unit_id: 1, auto_port: true, auto_unit_id: true, owner: 'replicator', status: 'AVAILABLE'}
});

const normalizeRepDevice = value => {
  const device = clone(value || {});
  if (!Array.isArray(device.pull_blocks) || !device.pull_blocks.length) device.pull_blocks = [blankBlock()];
  if (!device.destination) device.destination = {port: 5021, unit_id: 1, auto_port: true, auto_unit_id: true, owner: 'replicator', status: 'AVAILABLE'};
  return device;
};
const normalizeRepDocument = value => ({devices: Array.isArray(value && value.devices) ? value.devices.map(normalizeRepDevice) : []});

const validateReplicator = device => {
  if (!device.name || !device.name.trim()) return 'Name is required.';
  if (!device.endpoint || !device.endpoint.includes(':')) return 'Endpoint must be host:port.';
  if (device.unit_id < 0 || device.unit_id > 255) return 'Source Unit ID must be between 0 and 255.';
  for (let i = 0; i < device.pull_blocks.length; i += 1) {
    const block = device.pull_blocks[i];
    const label = `Pull Block ${i + 1}`;
    if (![1, 2, 3, 4].includes(Number(block.function))) return `${label} Function must be FC1, FC2, FC3, or FC4.`;
    if (block.start < 0 || block.start > 65535) return `${label} Start must be between 0 and 65535.`;
    if (block.count < 1 || block.start + block.count > 65536) return `${label} Count must be positive and remain inside the 16-bit address space.`;
    if (block.scan_rate_ms < 1) return `${label} Scan Rate must be greater than zero.`;
  }
  if (!device.destination.auto_port && (device.destination.port < 1 || device.destination.port > 65535)) return 'Destination Port must be between 1 and 65535.';
  if (device.destination.unit_id < 0 || device.destination.unit_id > 255) return 'Destination Unit ID must be between 0 and 255.';
  return null;
};

const replicatorState = {document: {devices: []}, persisted: {devices: []}, selected: null, selectedBlock: 0, loading: true, saving: false, message: 'Loading Replicator definitions...', error: false, runtimeStatus: null};
const selectedReplicator = () => replicatorState.selected === null ? null : replicatorState.document.devices[replicatorState.selected];

const renderReplicator = () => {
  const root = document.getElementById('replicator-root');
  root.replaceChildren();
  const shell = h('div', 'tool-layout');
  const sidebar = h('aside', 'tool-sidebar');
  sidebar.appendChild(h('h2', '', 'Replicator Devices'));
  const listActions = h('div', 'tool-actions');
  listActions.append(actionButton('Add', 'rep-add'), actionButton('Duplicate', 'rep-duplicate'), actionButton('Delete', 'rep-delete', 'tool-danger'));
  sidebar.appendChild(listActions);
  const list = h('div', 'tool-list');
  replicatorState.document.devices.forEach((device, index) => {
    const first = device.pull_blocks[0] || blankBlock();
    const row = actionButton('', 'rep-select', index === replicatorState.selected ? 'tool-row selected' : 'tool-row');
    row.dataset.index = String(index);
    row.append(h('strong', '', device.name || 'Unnamed device'), h('span', '', `${device.endpoint} / FC${first.function}${device.enabled ? '' : ' / Disabled'}`));
    list.appendChild(row);
  });
  if (!list.children.length) list.appendChild(h('div', 'tool-empty', replicatorState.loading ? 'Loading...' : 'No devices configured.'));
  sidebar.appendChild(list);
  shell.appendChild(sidebar);

  const editor = h('section', 'tool-editor');
  editor.appendChild(h('h2', '', 'Device Definition'));
  const device = selectedReplicator();
  if (!device) {
    editor.appendChild(h('div', 'tool-empty', 'Select a device or choose Add.'));
  } else {
    const runtime = h('div', 'runtime-row');
    const status = replicatorState.runtimeStatus;
    runtime.append(h('span', '', 'Replicator:'), h('strong', '', status ? (status.running ? 'RUNNING' : 'STOPPED') : '-'), h('span', '', 'Source:'), h('strong', '', status && status.source_status || '-'), h('span', '', 'Last Poll:'), h('strong', '', formatTime(status && status.last_poll)));
    editor.appendChild(runtime);
    const identity = h('div', 'tool-grid');
    identity.append(field('Name', device.name, {type: 'text'}, value => { device.name = value; }));
    identity.append(checkboxField('Enabled', device.enabled, value => { device.enabled = value; }));
    identity.append(field('Endpoint', device.endpoint, {type: 'text', placeholder: '192.168.1.20:502'}, value => { device.endpoint = value; }));
    identity.append(field('Source Unit ID', device.unit_id, {min: 0, max: 255}, value => { device.unit_id = numberValue(value); }));
    editor.appendChild(identity);

    editor.appendChild(h('h3', '', 'Destination'));
    const destination = h('div', 'tool-grid');
    destination.append(field('Port', device.destination.port, {min: 1, max: 65535, readOnly: device.destination.auto_port}, value => { device.destination.port = numberValue(value); }));
    destination.append(checkboxField('Auto Port', device.destination.auto_port, value => { device.destination.auto_port = value; }));
    destination.append(field('Unit ID', device.destination.unit_id, {min: 0, max: 255, readOnly: device.destination.auto_unit_id}, value => { device.destination.unit_id = numberValue(value); }));
    destination.append(checkboxField('Auto Unit ID', device.destination.auto_unit_id, value => { device.destination.auto_unit_id = value; }));
    editor.appendChild(destination);

    editor.appendChild(h('h3', '', 'Pull Blocks'));
    const blockActions = h('div', 'tool-actions');
    blockActions.append(actionButton('Add Block', 'rep-add-block'), actionButton('Duplicate Block', 'rep-duplicate-block'), actionButton('Delete Block', 'rep-delete-block', 'tool-danger'));
    editor.appendChild(blockActions);
    const table = h('div', 'block-table');
    const header = h('div', 'block-row block-header');
    ['#', 'FC', 'Start', 'Count', 'Scan Rate (ms)'].forEach(label => header.appendChild(h('span', '', label)));
    table.appendChild(header);
    device.pull_blocks.forEach((block, index) => {
      const row = h('div', index === replicatorState.selectedBlock ? 'block-row selected' : 'block-row');
      row.dataset.action = 'rep-select-block';
      row.dataset.index = String(index);
      row.appendChild(h('strong', '', String(index + 1)));
      const select = document.createElement('select');
      [1, 2, 3, 4].forEach(value => {
        const option = document.createElement('option');
        option.value = String(value);
        option.textContent = `FC${value}`;
        option.selected = Number(block.function) === value;
        select.appendChild(option);
      });
      select.addEventListener('pointerdown', event => event.stopPropagation());
      select.addEventListener('mousedown', event => event.stopPropagation());
      select.addEventListener('click', event => event.stopPropagation());
      select.addEventListener('change', event => { block.function = Number(event.target.value); });
      const selectWrap = h('span');
      selectWrap.appendChild(select);
      row.appendChild(selectWrap);
      row.appendChild(field('Start', block.start, {min: 0, max: 65535, className: 'cell-field'}, value => { block.start = numberValue(value); }));
      row.appendChild(field('Count', block.count, {min: 1, max: 65535, className: 'cell-field'}, value => { block.count = numberValue(value); }));
      row.appendChild(field('Scan', block.scan_rate_ms, {min: 1, step: 1, className: 'cell-field'}, value => { block.scan_rate_ms = numberValue(value); }));
      table.appendChild(row);
    });
    editor.appendChild(table);
    const validation = validateReplicator(device);
    if (validation) editor.appendChild(h('div', 'tool-validation', validation));
    const actions = h('div', 'editor-actions');
    const save = actionButton(replicatorState.saving ? 'Saving...' : 'Save & Apply', 'rep-save', 'tool-primary');
    save.disabled = Boolean(validation) || replicatorState.saving;
    actions.append(save, actionButton('Discard', 'rep-discard'));
    editor.appendChild(actions);
  }
  editor.appendChild(h('div', `tool-status${replicatorState.error ? ' error' : ''}`, replicatorState.message));
  shell.appendChild(editor);
  root.appendChild(shell);
};

const loadReplicator = async () => {
  replicatorState.loading = true;
  replicatorState.message = 'Loading Replicator definitions...';
  replicatorState.error = false;
  renderReplicator();
  try {
    const result = await window.mcsDesktop.replicatorCall('load', {});
    const documentValue = normalizeRepDocument(result.document);
    replicatorState.document = clone(documentValue);
    replicatorState.persisted = clone(documentValue);
    replicatorState.selected = documentValue.devices.length ? 0 : null;
    replicatorState.message = documentValue.devices.length ? 'Canonical Replicator definitions loaded.' : 'No devices configured. Choose Add to begin.';
    replicatorState.error = false;
  } catch (error) {
    replicatorState.message = error.message || String(error);
    replicatorState.error = true;
  } finally {
    replicatorState.loading = false;
    renderReplicator();
  }
};

const pollReplicator = async () => {
  const device = selectedReplicator();
  if (!device || replicatorState.saving) return;
  try {
    replicatorState.runtimeStatus = await window.mcsDesktop.replicatorCall('status', {name: device.name});
    renderReplicator();
  } catch (_) {}
};

const saveReplicator = async () => {
  const validation = replicatorState.document.devices.map(validateReplicator).find(Boolean);
  if (validation) {
    replicatorState.message = validation;
    replicatorState.error = true;
    renderReplicator();
    return;
  }
  replicatorState.saving = true;
  replicatorState.message = 'Saving Replicator definitions and applying runtime changes...';
  replicatorState.error = false;
  renderReplicator();
  try {
    const result = await window.mcsDesktop.replicatorCall('apply', {document: clone(replicatorState.document)});
    const documentValue = normalizeRepDocument(result.document);
    replicatorState.document = clone(documentValue);
    replicatorState.persisted = clone(documentValue);
    replicatorState.message = `${result.message || 'Replicator settings applied.'} ${result.completed_at ? `Completed at ${formatTime(result.completed_at)}.` : ''}`;
    replicatorState.error = false;
  } catch (error) {
    replicatorState.message = `Save & Apply failed: ${error.message || error}`;
    replicatorState.error = true;
  } finally {
    replicatorState.saving = false;
    renderReplicator();
    pollReplicator();
  }
};

document.addEventListener('click', event => {
  if (isEditableControl(event.target) && !event.target.closest('button[data-action]')) return;
  const target = event.target.closest('[data-action]');
  if (!target) return;
  const action = target.dataset.action;
  if (action === 'sim-select') simulatorState.selected = Number(target.dataset.index);
  else if (action === 'sim-add') { simulatorState.document.devices.push(simulatorBlankDevice(simulatorState.document.devices.length + 1)); simulatorState.selected = simulatorState.document.devices.length - 1; }
  else if (action === 'sim-duplicate' && selectedSimulator()) { const copy = clone(selectedSimulator()); copy.name = `${copy.name} (copy)`; simulatorState.document.devices.push(copy); simulatorState.selected = simulatorState.document.devices.length - 1; }
  else if (action === 'sim-delete' && selectedSimulator()) { simulatorState.document.devices.splice(simulatorState.selected, 1); simulatorState.selected = simulatorState.document.devices.length ? Math.min(simulatorState.selected, simulatorState.document.devices.length - 1) : null; }
  else if (action === 'sim-discard') { simulatorState.document = clone(simulatorState.persisted); simulatorState.selected = simulatorState.document.devices.length ? Math.min(simulatorState.selected || 0, simulatorState.document.devices.length - 1) : null; }
  else if (action === 'sim-save') { saveSimulator(); return; }
  else if (action === 'rep-select') { replicatorState.selected = Number(target.dataset.index); replicatorState.selectedBlock = 0; }
  else if (action === 'rep-add') { replicatorState.document.devices.push(blankReplicator(replicatorState.document.devices.length + 1)); replicatorState.selected = replicatorState.document.devices.length - 1; replicatorState.selectedBlock = 0; }
  else if (action === 'rep-duplicate' && selectedReplicator()) { const copy = clone(selectedReplicator()); copy.name = `${copy.name} (copy)`; replicatorState.document.devices.push(copy); replicatorState.selected = replicatorState.document.devices.length - 1; replicatorState.selectedBlock = 0; }
  else if (action === 'rep-delete' && selectedReplicator()) { replicatorState.document.devices.splice(replicatorState.selected, 1); replicatorState.selected = replicatorState.document.devices.length ? Math.min(replicatorState.selected, replicatorState.document.devices.length - 1) : null; replicatorState.selectedBlock = 0; }
  else if (action === 'rep-select-block') replicatorState.selectedBlock = Number(target.dataset.index);
  else if (action === 'rep-add-block' && selectedReplicator()) { selectedReplicator().pull_blocks.push(blankBlock()); replicatorState.selectedBlock = selectedReplicator().pull_blocks.length - 1; }
  else if (action === 'rep-duplicate-block' && selectedReplicator()) { const blocks = selectedReplicator().pull_blocks; const copy = clone(blocks[replicatorState.selectedBlock] || blocks[0] || blankBlock()); blocks.splice(replicatorState.selectedBlock + 1, 0, copy); replicatorState.selectedBlock += 1; }
  else if (action === 'rep-delete-block' && selectedReplicator()) { const blocks = selectedReplicator().pull_blocks; if (blocks.length > 1) blocks.splice(replicatorState.selectedBlock, 1); replicatorState.selectedBlock = Math.max(0, Math.min(replicatorState.selectedBlock, blocks.length - 1)); }
  else if (action === 'rep-discard') { replicatorState.document = clone(replicatorState.persisted); replicatorState.selected = replicatorState.document.devices.length ? Math.min(replicatorState.selected || 0, replicatorState.document.devices.length - 1) : null; replicatorState.selectedBlock = 0; }
  else if (action === 'rep-save') { saveReplicator(); return; }
  renderSimulator();
  renderReplicator();
  pollSimulator();
  pollReplicator();
});

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
  runtimePaths = paths;
  document.getElementById('bin-path').textContent = paths.bin;
  document.getElementById('data-path').textContent = paths.data;
  document.getElementById('runtime-mode').textContent = paths.mode === 'windows-service' ? 'Windows services (NSSM)' : 'Electron child processes';
  if (paths.mode === 'windows-service') {
    const startButton = document.getElementById('start-all');
    const stopButton = document.getElementById('stop-all');
    startButton.disabled = true;
    stopButton.disabled = true;
    startButton.title = 'Installed backend runtimes are managed by Windows Services.';
    stopButton.title = 'Closing Electron does not stop installed Windows services.';
    appendLog({process: 'electron', level: 'info', text: 'Backend runtimes are managed by NSSM Windows services. Use Windows Services for manual start/stop operations.'});
  }
}).catch(error => appendLog({process: 'electron', level: 'error', text: error.message || String(error)}));

renderSimulator();
renderReplicator();
loadSimulator();
loadReplicator();
setInterval(pollSimulator, 2000);
setInterval(pollReplicator, 2000);