import './index.scss';
import osjs from 'osjs';
import {name as applicationName} from './metadata.json';
const {createStatusPoller, createRuntimeMessageController} = require('./runtime-status-poller');
const {appliedPollTargets, pollTargetFor} = require('./poll-target');

const FC_KEYS = ['fc1', 'fc2', 'fc3', 'fc4'];
const FC_LABELS = {fc1: 'Coils (FC1)', fc2: 'Discrete Inputs (FC2)', fc3: 'Holding Registers (FC3)', fc4: 'Input Registers (FC4)'};
const clone = value => JSON.parse(JSON.stringify(value));
const numberValue = value => value === '' ? 0 : Number(value);
const normalizeDocument = value => ({devices: Array.isArray(value && value.devices) ? value.devices : []});

const blankDevice = sequence => ({
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

const addressRange = area => {
  const start = Number(area && area.start) || 0;
  const count = Number(area && area.count) || 0;
  if (!count) return 'Unused';
  const end = start + count - 1;
  return start === end ? String(start) : `${start}-${end}`;
};

const validateDevice = device => {
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

const element = (tag, className, text) => {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (text !== undefined) node.textContent = text;
  return node;
};

const button = (label, action, className = '') => {
  const node = element('button', `sim-button ${className}`.trim(), label);
  node.type = 'button';
  node.dataset.action = action;
  return node;
};

const operatorState = value => ['RUNNING', 'WAITING', 'STOPPED', 'ERROR'].includes(value) ? value : 'UNAVAILABLE';

const statusPair = (label, value) => {
  const pair = element('span', 'sim-runtime-pair');
  const state = value || '—';
  pair.dataset.state = state.toLowerCase();
  const indicator = element('span', 'sim-runtime-indicator', '●');
  indicator.setAttribute('aria-hidden', 'true');
  pair.append(element('strong', '', label), indicator, element('span', 'sim-runtime-word', state));
  return pair;
};

const setStatusPair = (pair, value) => {
  const state = value || '—';
  pair.dataset.state = state.toLowerCase();
  const word = pair.querySelector('.sim-runtime-word');
  if (word) word.textContent = state;
};

const field = (label, value, settings, onChange) => {
  const wrapper = element('label', settings.className || 'sim-field');
  wrapper.appendChild(element('span', 'sim-field-label', label));
  const control = document.createElement('input');
  control.type = settings.type || 'number';
  control.value = value;
  ['min', 'max', 'step'].forEach(key => {
    if (settings[key] !== undefined) control[key] = settings[key];
  });
  control.addEventListener('input', event => onChange(event.target.value));
  wrapper.appendChild(control);
  return wrapper;
};

const register = (core, args, options, metadata) => {
  const proc = core.make('osjs/application', {args, options, metadata});
  const win = proc.createWindow({
    id: 'ModbusSimulatorWindow',
    title: metadata.title && metadata.title.en_EN ? metadata.title.en_EN : 'Modbus Simulator',
    dimension: {width: 920, height: 620},
    position: 'center'
  });
  const state = {document: {devices: []}, persisted: {devices: []}, selected: null, pollTargets: [], runtimeStatus: null, runtimeStatusError: null, search: '', message: 'Connecting to Simulator runtime...', error: false, saving: false};
  const pending = new Map();
  let requestSequence = 0;
  const runtimeCall = (operation, payload = {}) => new Promise((resolve, reject) => {
    const requestId = `${Date.now()}-${proc.pid}-${++requestSequence}`;
    const timeout = setTimeout(() => {
      pending.delete(requestId);
      reject(new Error('Simulator runtime request timed out.'));
    }, 26000);
    pending.set(requestId, {resolve, reject, timeout});
    proc.send({version: 1, request_id: requestId, operation, payload});
  });
  proc.on('ws:message', response => {
    const entry = response && pending.get(response.request_id);
    if (!entry) return;
    clearTimeout(entry.timeout);
    pending.delete(response.request_id);
    if (response.ok) entry.resolve(response.result);
    else entry.reject(new Error(response.error && response.error.message || 'Simulator runtime request failed.'));
  });
  const selectedDevice = () => state.selected === null ? null : state.document.devices[state.selected];
  const appliedPollTarget = () => pollTargetFor(state.pollTargets, state.selected);
  let runtimeMessages;
  const setMessage = (message, error = false) => {
    Object.assign(state, {message, error});
    if (runtimeMessages) runtimeMessages.operationMessage();
  };
  runtimeMessages = createRuntimeMessageController({
    publish: (message, error) => Object.assign(state, {message, error})
  });
  const statusPoller = createStatusPoller({
    requestStatus: name => runtimeCall('status', {name}).then(result => result.status),
    onStatus: (name, status) => {
      if (name !== appliedPollTarget()) return;
      state.runtimeStatus = status;
      state.runtimeStatusError = null;
      runtimeMessages.status(status);
      patchStatusUI();
    },
    onUnavailable: (name, error) => {
      if (name !== appliedPollTarget()) return;
      state.runtimeStatus = null;
      state.runtimeStatusError = error.message || String(error;
      runtimeMessages.unavailable(error);
      patchStatusUI();
    }
  });
  const pollSelectedStatus = (force = false) => {
    if (state.saving) return;
    statusPoller.select(appliedPollTarget(), force);
  };

  const patchStatusUI = () => {
    if (!win.$content) return;
    const row = win.$content.querySelector('.sim-runtime-row');
    if (!row) return;
    const device = selectedDevice();
    const pairs = row.querySelectorAll('.sim-runtime-pair');
    const mma2 = device ? operatorState(state.runtimeStatus && state.runtimeStatus.mma2_status) : '—';
    const sim = device ? operatorState(state.runtimeStatus && state.runtimeStatus.device_status) : '—';
    if (pairs.length > 0) setStatusPair(pairs[0], mma2);
    if (pairs.length > 1) setStatusPair(pairs[1], sim);
    const bar = win.$content.querySelector('.sim-status');
    if (bar) {
      bar.textContent = state.message;
      bar.classList.toggle('sim-status-error', Boolean(state.error));
    }
  };

  const renderEditor = root => {
    const pane = element('section', 'sim-editor-pane');
    const device = selectedDevice();
    pane.appendChild(element('h2', 'sim-editor-title', 'Device Definition'));
    const runtimeRow = element('div', 'sim-runtime-row');
    const status = device && state.runtimeStatus;
    runtimeRow.append(
      statusPair('MMA2', device ? operatorState(status && status.mma2_status) : '—'),
      statusPair('Simulator', device ? operatorState(status && status.device_status) : '—')
    );
    pane.appendChild(runtimeRow);
    if (!device) {
      pane.appendChild(element('div', 'sim-empty', 'Select a device or choose Add.'));
      root.appendChild(pane);
      return;
    }
    const identity = element('div', 'sim-identity-grid');
    identity.appendChild(field('Name', device.name, {type: 'text'}, value => { device.name = value; }));
    const enabled = element('label', 'sim-checkbox');
    const checkbox = document.createElement('input');
    checkbox.type = 'checkbox';
    checkbox.checked = device.enabled;
    checkbox.addEventListener('change', event => { device.enabled = event.target.checked; });
    enabled.append(checkbox, element('span', '', 'Enabled'));
    identity.appendChild(enabled);
    identity.appendChild(field('Listen Port', device.mma2.port, {min: 1, max: 65535}, value => { device.mma2.port = numberValue(value); }));
    identity.appendChild(field('Unit ID', device.mma2.unit_id, {min: 0, max: 255}, value => { device.mma2.unit_id = numberValue(value); }));
    pane.appendChild(identity);

    const table = element('div', 'sim-fc-table');
    const header = element('div', 'sim-fc-row sim-fc-header');
    ['Function', 'Start', 'Count', 'Randomize Every (ms)', 'Address Range'].forEach(label => header.appendChild(element('span', '', label)));
    table.appendChild(header);
    FC_KEYS.forEach(fc => {
      const row = element('div', 'sim-fc-row');
      const area = device.mma2[fc];
      const updateRange = () => {
        const range = row.querySelector('.sim-range');
        if (range) range.textContent = addressRange(area);
      };
      row.appendChild(element('strong', '', FC_LABELS[fc]));
      row.appendChild(field('Start', area.start, {min: 0, max: 65535, className: 'sim-cell-field'}, value => { area.start = numberValue(value); updateRange(); }));
      row.appendChild(field('Count', area.count, {min: 0, max: 65535, className: 'sim-cell-field'}, value => { area.count = numberValue(value); updateRange(); }));
      row.appendChild(field('Interval', device.random_runtime[`${fc}_interval_ms`], {min: 0, step: 1, className: 'sim-cell-field'}, value => { device.random_runtime[`${fc}_interval_ms`] = numberValue(value); }));
      row.appendChild(element('span', 'sim-range', addressRange(area)));
      table.appendChild(row);
    });
    pane.appendChild(table);

    const validation = validateDevice(device);
    if (validation) pane.appendChild(element('div', 'sim-validation', validation));
    const actions = element('div', 'sim-editor-actions');
    const saveButton = button(state.saving ? 'Saving...' : 'Save & Apply', 'save', 'sim-primary');
    saveButton.disabled = Boolean(validation) || state.saving;
    actions.append(saveButton, button('Discard', 'discard'));
    pane.appendChild(actions);
    root.appendChild(pane);
  };

  const render = () => {
    const root = element('div', 'modbus-simulator');
    const sidebar = element('aside', 'sim-sidebar');
    const search = document.createElement('input');
    search.className = 'sim-search';
    search.type = 'search';
    search.placeholder = 'Search devices...';
    search.value = state.search;
    search.addEventListener('input', event => {
      state.search = event.target.value;
      const query = state.search.trim().toLowerCase();
      sidebar.querySelectorAll('.sim-device-row').forEach(row => {
        row.hidden = Boolean(query) && !row.textContent.toLowerCase().includes(query);
      });
    });
    sidebar.appendChild(search);
    const listActions = element('div', 'sim-list-actions');
    listActions.append(button('Add', 'add'), button('Duplicate', 'duplicate'), button('Delete', 'delete', 'sim-danger'));
    sidebar.appendChild(listActions);
    const list = element('div', 'sim-device-list');
    const query = state.search.trim().toLowerCase();
    state.document.devices.forEach((device, index) => {
      if (query && !device.name.toLowerCase().includes(query)) return;
      const row = button('', 'select', index === state.selected ? 'sim-device-selected' : '');
      row.dataset.index = String(index);
      row.classList.add('sim-device-row');
      row.append(element('strong', '', device.name || 'Unnamed device'), element('span', '', `Port ${device.mma2.port} - Unit ${device.mma2.unit_id}${device.enabled ? '' : ' - Disabled'}`));
      list.appendChild(row);
    });
    if (!list.children.length) list.appendChild(element('div', 'sim-empty', 'No matching devices.'));
    sidebar.appendChild(list);
    root.appendChild(sidebar);
    renderEditor(root);
    root.appendChild(element('div', `sim-status${state.error ? ' sim-status-error' : ''}`, state.message));
    win.$content.replaceChildren(root);
  };

  const save = async () => {
    const error = state.document.devices.map(validateDevice).find(Boolean);
    if (error) { setMessage(error, true); render(); return; }
    state.saving = true;
    statusPoller.select(null);  // pause polling: drops in-flight responses (generation bump) and stops new requests while the apply transaction runs.
    setMessage('Saving simulator definitions...');
    render();
    const edited = clone(normalizeDocument(state.document));
    try {
      const result = await runtimeCall('apply', {document: edited});
      const applied = clone(normalizeDocument(result.document));
      state.document = applied;
      state.persisted = clone(applied);
      state.pollTargets = appliedPollTargets(applied.devices;
      if (state.selected !== null && state.selected >= applied.devices.length) state.selected = null;
      setMessage(`${result.message} Applied at ${new Date(result.completed_at).toLocaleString()}.`);
      state.runtimeStatus = null;
      state.runtimeStatusError = null;
      patchStatusUI();
    } catch (error) {
      setMessage(`Save & Apply failed: ${error.message || error}`, true);
    } finally {
      state.saving = false;
      render();
      pollSelectedStatus(true);
    }
  };

  win.on('destroy', () => {
    statusPoller.stop();
    pending.forEach(entry => {
      clearTimeout(entry.timeout);
      entry.reject(new Error('Simulator window closed.'));
    });
    pending.clear();
    proc.destroy();
  });
  win.render($content => {
    win.$content = $content;
    $content.addEventListener('click', event => {
      const target = event.target.closest('[data-action]');
      if (!target) return;
      const action = target.dataset.action;
      if (action === 'select') {
        state.selected = Number(target.dataset.index);
        setMessage('Editing a simulator-owned definition.');
      } else if (action === 'add') {
        state.document.devices.push(blankDevice(state.document.devices.length + 1));
        state.pollTargets.push(null);
        state.selected = state.document.devices.length - 1;
        setMessage('New device added locally. Save & Apply to persist it.');
      } else if (action === 'duplicate' && selectedDevice()) {
        const copy = clone(selectedDevice());
        copy.name = `${copy.name} (copy)`;
        state.document.devices.push(copy);
        state.pollTargets.push(null);
        state.selected = state.document.devices.length - 1;
        setMessage('Device duplicated locally. Save & Apply to persist it.');
      } else if (action === 'delete' && selectedDevice()) {
        state.document.devices.splice(state.selected, 1);
        state.pollTargets.splice(state.selected, 1);
        state.selected = state.document.devices.length ? Math.min(state.selected, state.document.devices.length - 1) : null;
        setMessage('Device deleted locally. Save & Apply to persist it.');
      } else if (action === 'discard') {
        state.document = clone(state.persisted);
        state.pollTargets = appliedPollTargets(state.persisted.devices);
        state.selected = state.document.devices.length ? Math.min(state.selected || 0, state.document.devices.length - 1) : null;
        setMessage('Unapplied changes discarded.');
      } else if (action === 'save') {
        save();
        return;
      }
      state.runtimeStatus = null;
      state.runtimeStatusError = null;
      pollSelectedStatus();
      render();
      patchStatusUI();
    });
    render();
    runtimeCall('load')
      .then(result => {
        const persisted = normalizeDocument(result.document);
        state.document = clone(persisted);
        state.persisted = clone(persisted);
        state.pollTargets = appliedPollTargets(persisted.devices);
        state.selected = persisted.devices.length ? 0 : null;
        setMessage(persisted.devices.length ? 'Canonical Simulator definitions loaded.' : 'No devices configured. Choose Add to begin.');
        render();
        patchStatusUI();
        pollSelectedStatus();
      })
      .catch(error => {
        setMessage(`Simulator runtime unavailable: ${error.message || error}`, true);
        render();
      });
  });
  return proc;
};

osjs.register(applicationName, register);
