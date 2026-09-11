import './index.scss';
import osjs from 'osjs';
import {name as applicationName} from './metadata.json';

const clone = value => JSON.parse(JSON.stringify(value));
const numberValue = value => value === '' ? 0 : Number(value);
const normalizeDevice = value => {
  const device = clone(value || {});
  if (!device.pull_block) {
    device.pull_block = {
      function: Number(device.function || 3),
      start: Number(device.start || 0),
      count: Number(device.count || 16),
      scan_rate_ms: Number(device.scan_rate_ms || 1000)
    };
  }
  delete device.function;
  delete device.start;
  delete device.count;
  delete device.scan_rate_ms;
  return device;
};
const normalizeDocument = value => ({
  devices: Array.isArray(value && value.devices) ? value.devices.map(normalizeDevice) : []
});

const element = (tag, className, text) => {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (text !== undefined) node.textContent = text;
  return node;
};

const button = (label, action, className = '') => {
  const node = element('button', `rep-button ${className}`.trim(), label);
  node.type = 'button';
  node.dataset.action = action;
  return node;
};

const field = (label, value, settings, onChange) => {
  const wrapper = element('label', settings.className || 'rep-field');
  wrapper.appendChild(element('span', 'rep-field-label', label));
  const control = document.createElement('input');
  control.type = settings.type || 'number';
  control.value = value;
  control.readOnly = Boolean(settings.readOnly);
  ['min', 'max', 'step', 'placeholder'].forEach(key => {
    if (settings[key] !== undefined) control[key] = settings[key];
  });
  control.addEventListener('input', event => onChange(event.target.value));
  wrapper.appendChild(control);
  return wrapper;
};

const checkboxField = (label, checked, onChange) => {
  const wrapper = element('label', 'rep-checkbox');
  const control = document.createElement('input');
  control.type = 'checkbox';
  control.checked = checked;
  control.addEventListener('change', event => onChange(event.target.checked));
  wrapper.append(control, element('span', '', label));
  return wrapper;
};

const blankDevice = (sequence, suggestion) => ({
  name: `Rep-PLC-${sequence}`,
  enabled: true,
  endpoint: '127.0.0.1:5020',
  unit_id: 1,
  pull_block: {
    function: 3,
    start: 0,
    count: 16,
    scan_rate_ms: 1000
  },
  destination: {
    port: suggestion && suggestion.port || 5021,
    unit_id: suggestion && suggestion.unit_id || 1,
    auto_port: true,
    auto_unit_id: true,
    owner: 'replicator',
    status: 'AVAILABLE'
  }
});

const validateDevice = device => {
  const block = device.pull_block || {};
  if (!device.name || !device.name.trim()) return 'Name is required.';
  if (!device.endpoint || !device.endpoint.includes(':')) return 'Endpoint must be host:port.';
  if (device.unit_id < 0 || device.unit_id > 255) return 'Source Unit ID must be between 0 and 255.';
  if (![3, 4].includes(Number(block.function))) return 'Pull Block Function must be FC3 or FC4.';
  if (block.start < 0 || block.start > 65535) return 'Pull Block Start must be between 0 and 65535.';
  if (block.count < 1 || block.start + block.count > 65536) return 'Pull Block Count must be positive and remain inside the 16-bit address space.';
  if (block.scan_rate_ms < 1) return 'Pull Block Scan Rate must be greater than zero.';
  if (!device.destination.auto_port && (device.destination.port < 1 || device.destination.port > 65535)) return 'Destination Port must be between 1 and 65535.';
  if (device.destination.unit_id < 0 || device.destination.unit_id > 255) return 'Destination Unit ID must be between 0 and 255.';
  return null;
};

const formatLastPoll = value => {
  if (!value) return '—';
  const parsed = new Date(value);
  return Number.isNaN(parsed.getTime()) ? value : parsed.toLocaleTimeString();
};

const register = (core, args, options, metadata) => {
  const proc = core.make('osjs/application', {args, options, metadata});
  const win = proc.createWindow({
    id: 'ModbusReplicatorWindow',
    title: metadata.title && metadata.title.en_EN ? metadata.title.en_EN : 'Modbus Replicator',
    dimension: {width: 900, height: 580},
    position: 'center'
  });

  const state = {
    document: {devices: []},
    persisted: {devices: []},
    selected: null,
    search: '',
    suggestion: {port: 5021, unit_id: 1, owner: 'replicator', status: 'AVAILABLE'},
    runtimeStatus: null,
    message: 'Connecting to Replicator runtime...',
    error: false,
    saving: false
  };
  const pending = new Map();
  let requestSequence = 0;
  let statusTimer = null;
  let inspectTimer = null;

  const runtimeCall = (operation, payload = {}) => new Promise((resolve, reject) => {
    const requestId = `${Date.now()}-${proc.pid}-${++requestSequence}`;
    const timeout = setTimeout(() => {
      pending.delete(requestId);
      reject(new Error('Replicator runtime request timed out.'));
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
    else entry.reject(new Error(response.error && response.error.message || 'Replicator runtime request failed.'));
  });

  const selectedDevice = () => state.selected === null ? null : state.document.devices[state.selected];
  const selectedIsPersisted = () => {
    const device = selectedDevice();
    return Boolean(device && state.persisted.devices.some(entry => entry.name === device.name));
  };
  const setMessage = (message, error = false) => Object.assign(state, {message, error});

  const pollStatus = async () => {
    const device = selectedDevice();
    if (!device || !selectedIsPersisted() || state.saving) {
      state.runtimeStatus = null;
      patchStatus();
      return;
    }
    try {
      state.runtimeStatus = await runtimeCall('status', {name: device.name});
    } catch (error) {
      state.runtimeStatus = {running: false, source_status: 'ERROR', last_error: error.message || String(error)};
    }
    patchStatus();
  };

  const patchStatus = () => {
    if (!win.$content) return;
    const runtime = state.runtimeStatus;
    const rep = win.$content.querySelector('[data-runtime="replicator"]');
    const source = win.$content.querySelector('[data-runtime="source"]');
    const poll = win.$content.querySelector('[data-runtime="poll"]');
    const lastError = win.$content.querySelector('.rep-last-error');
    if (rep) rep.textContent = runtime ? (runtime.running ? 'RUNNING' : 'STOPPED') : '—';
    if (source) source.textContent = runtime ? runtime.source_status || '—' : '—';
    if (poll) poll.textContent = runtime ? formatLastPoll(runtime.last_poll) : '—';
    if (lastError) {
      lastError.textContent = runtime && runtime.last_error ? runtime.last_error : '';
      lastError.hidden = !(runtime && runtime.last_error);
    }
    const bar = win.$content.querySelector('.rep-status');
    if (bar) {
      bar.textContent = state.message;
      bar.classList.toggle('rep-status-error', Boolean(state.error));
    }
  };

  const inspectDestination = () => {
    clearTimeout(inspectTimer);
    const device = selectedDevice();
    if (!device || device.destination.auto_port || device.destination.auto_unit_id) return;
    inspectTimer = setTimeout(() => {
      runtimeCall('suggest', {inspect: true, port: device.destination.port, unit_id: device.destination.unit_id})
        .then(result => {
          if (selectedDevice() !== device) return;
          Object.assign(device.destination, result);
          render();
        })
        .catch(error => {
          device.destination.status = 'UNAVAILABLE';
          device.destination.owner = error.message || 'unknown';
          render();
        });
    }, 250);
  };

  const localSuggestion = () => {
    let port = Number(state.suggestion.port) || 5021;
    let unit = Number(state.suggestion.unit_id) || 1;
    const used = new Set(state.document.devices.map(device => `${device.destination.port}/${device.destination.unit_id}`));
    while (used.has(`${port}/${unit}`)) {
      unit += 1;
      if (unit > 255) {
        unit = 1;
        port += 1;
      }
    }
    return {port, unit_id: unit, owner: 'replicator', status: 'AVAILABLE'};
  };

  const renderEditor = root => {
    const pane = element('section', 'rep-editor-pane');
    pane.appendChild(element('h2', 'rep-editor-title', 'Device Definition'));
    const device = selectedDevice();
    if (!device) {
      const pendingDelete = state.persisted.devices.length > 0 && state.document.devices.length === 0;
      pane.appendChild(element('div', 'rep-empty', pendingDelete
        ? 'All devices are marked for deletion. Save & Apply to release Replicator-owned destinations, or Discard to restore them.'
        : 'Select a device or choose Add.'));
      const actions = element('div', 'rep-editor-actions');
      const saveButton = button(state.saving ? 'Saving...' : 'Save & Apply', 'save', 'rep-primary');
      saveButton.disabled = state.saving || !pendingDelete;
      actions.append(saveButton, button('Discard', 'discard'));
      pane.appendChild(actions);
      root.appendChild(pane);
      return;
    }

    const identity = element('div', 'rep-grid');
    identity.appendChild(field('Name', device.name, {type: 'text'}, value => { device.name = value; }));
    identity.appendChild(checkboxField('Enabled', device.enabled, value => { device.enabled = value; }));
    pane.appendChild(identity);

    pane.appendChild(element('h3', 'rep-section-title', 'Source'));
    const sourceGrid = element('div', 'rep-grid rep-source-grid');
    sourceGrid.appendChild(field('Endpoint', device.endpoint, {type: 'text', placeholder: '192.168.1.20:502'}, value => { device.endpoint = value; }));
    sourceGrid.appendChild(field('Unit ID', device.unit_id, {min: 0, max: 255}, value => { device.unit_id = numberValue(value); }));
    pane.appendChild(sourceGrid);

    pane.appendChild(element('h3', 'rep-section-title', 'Pull Block'));
    const block = device.pull_block;
    const blockGrid = element('div', 'rep-grid rep-source-grid');
    const fc = element('label', 'rep-field');
    fc.appendChild(element('span', 'rep-field-label', 'FC'));
    const select = document.createElement('select');
    [3, 4].forEach(value => {
      const option = document.createElement('option');
      option.value = String(value);
      option.textContent = `FC${value}`;
      option.selected = Number(block.function) === value;
      select.appendChild(option);
    });
    select.addEventListener('change', event => { block.function = Number(event.target.value); });
    fc.appendChild(select);
    blockGrid.appendChild(fc);
    blockGrid.appendChild(field('Start', block.start, {min: 0, max: 65535}, value => { block.start = numberValue(value); }));
    blockGrid.appendChild(field('Count', block.count, {min: 1, max: 65535}, value => { block.count = numberValue(value); }));
    blockGrid.appendChild(field('Scan Rate (ms)', block.scan_rate_ms, {min: 1, step: 1}, value => { block.scan_rate_ms = numberValue(value); }));
    pane.appendChild(blockGrid);

    pane.appendChild(element('h3', 'rep-section-title', 'Destination'));
    const destGrid = element('div', 'rep-destination-grid');
    const portField = field('Port', device.destination.port, {min: 1, max: 65535, readOnly: device.destination.auto_port}, value => {
      device.destination.port = numberValue(value);
      inspectDestination();
    });
    const unitField = field('Unit ID', device.destination.unit_id, {min: 0, max: 255, readOnly: device.destination.auto_unit_id}, value => {
      device.destination.unit_id = numberValue(value);
      inspectDestination();
    });
    destGrid.appendChild(portField);
    destGrid.appendChild(checkboxField('Auto Port', device.destination.auto_port, value => {
      device.destination.auto_port = value;
      if (value) {
        const suggestion = localSuggestion();
        device.destination.port = suggestion.port;
        if (device.destination.auto_unit_id) device.destination.unit_id = suggestion.unit_id;
        Object.assign(device.destination, {owner: 'replicator', status: 'AVAILABLE'});
      }
      render();
      inspectDestination();
    }));
    destGrid.appendChild(unitField);
    destGrid.appendChild(checkboxField('Auto Unit ID', device.destination.auto_unit_id, value => {
      device.destination.auto_unit_id = value;
      if (value) {
        const suggestion = localSuggestion();
        if (device.destination.auto_port) device.destination.port = suggestion.port;
        device.destination.unit_id = suggestion.unit_id;
        Object.assign(device.destination, {owner: 'replicator', status: 'AVAILABLE'});
      }
      render();
      inspectDestination();
    }));
    const ownership = element('div', 'rep-ownership');
    ownership.append(
      element('span', 'rep-field-label', 'Owner'),
      element('strong', '', device.destination.owner || 'Replicator'),
      element('span', 'rep-field-label', 'Status'),
      element('strong', `rep-destination-status rep-destination-${String(device.destination.status || '').toLowerCase().replace(/\s+/g, '-')}`, device.destination.status || 'AVAILABLE')
    );
    destGrid.appendChild(ownership);
    pane.appendChild(destGrid);
    if (String(device.destination.status).toUpperCase() === 'IN USE') {
      pane.appendChild(element(
        'div',
        'rep-validation',
        `Destination ${device.destination.port}/${device.destination.unit_id} is owned by ${device.destination.owner || 'another producer'} and cannot be claimed by Replicator.`
      ));
    }

    const runtime = element('div', 'rep-runtime-row');
    runtime.append(
      element('span', '', 'Replicator: '), element('strong', '', '—'),
      element('span', '', 'Source: '), element('strong', '', '—'),
      element('span', '', 'Last Poll: '), element('strong', '', '—')
    );
    runtime.querySelectorAll('strong')[0].dataset.runtime = 'replicator';
    runtime.querySelectorAll('strong')[1].dataset.runtime = 'source';
    runtime.querySelectorAll('strong')[2].dataset.runtime = 'poll';
    pane.appendChild(runtime);
    const lastError = element('div', 'rep-last-error');
    lastError.hidden = true;
    pane.appendChild(lastError);

    const validation = validateDevice(device);
    if (validation) pane.appendChild(element('div', 'rep-validation', validation));
    const actions = element('div', 'rep-editor-actions');
    const saveButton = button(state.saving ? 'Saving...' : 'Save & Apply', 'save', 'rep-primary');
    saveButton.disabled = Boolean(validation) || state.saving;
    actions.append(saveButton, button('Discard', 'discard'));
    pane.appendChild(actions);
    root.appendChild(pane);
  };

  const render = () => {
    const root = element('div', 'modbus-replicator');
    const sidebar = element('aside', 'rep-sidebar');
    const search = document.createElement('input');
    search.className = 'rep-search';
    search.type = 'search';
    search.placeholder = 'Search devices...';
    search.value = state.search;
    search.addEventListener('input', event => {
      state.search = event.target.value;
      const query = state.search.trim().toLowerCase();
      sidebar.querySelectorAll('.rep-device-row').forEach(row => {
        row.hidden = Boolean(query) && !row.textContent.toLowerCase().includes(query);
      });
    });
    sidebar.appendChild(search);
    const listActions = element('div', 'rep-list-actions');
    listActions.append(button('Add', 'add'), button('Duplicate', 'duplicate'), button('Delete', 'delete', 'rep-danger'));
    sidebar.appendChild(listActions);
    const list = element('div', 'rep-device-list');
    const query = state.search.trim().toLowerCase();
    state.document.devices.forEach((device, index) => {
      if (query && !device.name.toLowerCase().includes(query)) return;
      const row = button('', 'select', index === state.selected ? 'rep-device-selected' : '');
      row.dataset.index = String(index);
      row.classList.add('rep-device-row');
      row.append(
        element('strong', '', device.name || 'Unnamed device'),
        element('span', '', `${device.endpoint} • FC${device.pull_block.function} • ${device.enabled ? 'Enabled' : 'Disabled'}`)
      );
      list.appendChild(row);
    });
    if (!list.children.length) list.appendChild(element('div', 'rep-empty', 'No matching devices.'));
    sidebar.appendChild(list);
    root.appendChild(sidebar);
    renderEditor(root);
    root.appendChild(element('div', `rep-status${state.error ? ' rep-status-error' : ''}`, state.message));
    win.$content.replaceChildren(root);
    patchStatus();
  };

  const save = async () => {
    const error = state.document.devices.map(validateDevice).find(Boolean);
    if (error) {
      setMessage(error, true);
      render();
      return;
    }
    state.saving = true;
    setMessage('Checking MMA2 ownership, saving shared settings, and restarting MMA2...');
    render();
    try {
      const result = await runtimeCall('apply', {document: clone(normalizeDocument(state.document))});
      const applied = clone(normalizeDocument(result.document));
      state.document = applied;
      state.persisted = clone(applied);
      if (state.selected !== null && state.selected >= applied.devices.length) state.selected = applied.devices.length ? applied.devices.length - 1 : null;
      setMessage(`${result.message} Applied at ${new Date(result.completed_at).toLocaleString()}.`);
      state.runtimeStatus = null;
      runtimeCall('suggest').then(value => { state.suggestion = value; }).catch(() => {});
    } catch (error) {
      setMessage(`Save & Apply failed: ${error.message || error}`, true);
    } finally {
      state.saving = false;
      render();
      pollStatus();
    }
  };

  win.on('destroy', () => {
    clearInterval(statusTimer);
    clearTimeout(inspectTimer);
    pending.forEach(entry => {
      clearTimeout(entry.timeout);
      entry.reject(new Error('Replicator window closed.'));
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
        state.runtimeStatus = null;
        setMessage('Editing a Replicator device definition.');
      } else if (action === 'add') {
        state.document.devices.push(blankDevice(state.document.devices.length + 1, localSuggestion()));
        state.selected = state.document.devices.length - 1;
        state.runtimeStatus = null;
        setMessage('New device added locally. Save & Apply to persist it.');
      } else if (action === 'duplicate' && selectedDevice()) {
        const copy = clone(selectedDevice());
        copy.name = `${copy.name} (copy)`;
        const suggestion = localSuggestion();
        copy.destination.port = suggestion.port;
        copy.destination.unit_id = suggestion.unit_id;
        copy.destination.auto_port = true;
        copy.destination.auto_unit_id = true;
        copy.destination.owner = 'replicator';
        copy.destination.status = 'AVAILABLE';
        state.document.devices.push(copy);
        state.selected = state.document.devices.length - 1;
        state.runtimeStatus = null;
        setMessage('Device duplicated locally with a new automatic destination.');
      } else if (action === 'delete' && selectedDevice()) {
        state.document.devices.splice(state.selected, 1);
        state.selected = state.document.devices.length ? Math.min(state.selected, state.document.devices.length - 1) : null;
        state.runtimeStatus = null;
        setMessage('Device deleted locally. Save & Apply to release its Replicator-owned destination.');
      } else if (action === 'discard') {
        state.document = clone(state.persisted);
        state.selected = state.document.devices.length ? Math.min(state.selected || 0, state.document.devices.length - 1) : null;
        state.runtimeStatus = null;
        setMessage('Unapplied changes discarded.');
      } else if (action === 'save') {
        save();
        return;
      }
      render();
      pollStatus();
    });

    render();
    runtimeCall('load')
      .then(result => {
        const persisted = normalizeDocument(result.document);
        state.document = clone(persisted);
        state.persisted = clone(persisted);
        state.suggestion = result.suggestion || state.suggestion;
        state.selected = persisted.devices.length ? 0 : null;
        setMessage(persisted.devices.length ? 'Canonical Replicator definitions loaded.' : 'No devices configured. Choose Add to begin.');
        render();
        pollStatus();
      })
      .catch(error => {
        setMessage(`Replicator runtime unavailable: ${error.message || error}`, true);
        render();
      });
    statusTimer = setInterval(pollStatus, 1000);
  });

  return proc;
};

osjs.register(applicationName, register);
