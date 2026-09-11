import './index.scss';
import osjs from 'osjs';
import {name as applicationName} from './metadata.json';

const clone = value => JSON.parse(JSON.stringify(value));
const numberValue = value => value === '' ? 0 : Number(value);
const blankBlock = () => ({function: 3, start: 0, count: 16, scan_rate_ms: 1000});

const normalizeDevice = value => {
  const device = clone(value || {});
  if (!Array.isArray(device.pull_blocks) || !device.pull_blocks.length) {
    device.pull_blocks = [device.pull_block ? clone(device.pull_block) : {
      function: Number(device.function || 3),
      start: Number(device.start || 0),
      count: Number(device.count || 16),
      scan_rate_ms: Number(device.scan_rate_ms || 1000)
    }];
  }
  delete device.pull_block;
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
  pull_blocks: [blankBlock()],
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
  if (!device.name || !device.name.trim()) return 'Name is required.';
  if (!device.endpoint || !device.endpoint.includes(':')) return 'Endpoint must be host:port.';
  if (device.unit_id < 0 || device.unit_id > 255) return 'Source Unit ID must be between 0 and 255.';
  if (!Array.isArray(device.pull_blocks) || !device.pull_blocks.length) return 'At least one Pull Block is required.';
  for (let index = 0; index < device.pull_blocks.length; index += 1) {
    const block = device.pull_blocks[index];
    const label = `Pull Block ${index + 1}`;
    if (![1, 2, 3, 4].includes(Number(block.function))) return `${label} Function must be FC1, FC2, FC3, or FC4.`;
    if (block.start < 0 || block.start > 65535) return `${label} Start must be between 0 and 65535.`;
    if (block.count < 1 || block.start + block.count > 65536) return `${label} Count must be positive and remain inside the 16-bit address space.`;
    if (block.scan_rate_ms < 1) return `${label} Scan Rate must be greater than zero.`;
  }
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
    selectedBlock: 0,
    activeTab: 'device',
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

  const patchStatus = () => {
    if (!win.$content) return;
    const runtime = state.runtimeStatus;
    const rep = win.$content.querySelector('[data-runtime="replicator"]');
    const source = win.$content.querySelector('[data-runtime="source"]');
    const poll = win.$content.querySelector('[data-runtime="poll"]');
    const operational = win.$content.querySelector('[data-runtime="operational"]');
    const lastError = win.$content.querySelector('.rep-last-error');
    if (rep) rep.textContent = runtime ? (runtime.running ? 'RUNNING' : 'STOPPED') : '—';
    if (source) source.textContent = runtime ? runtime.source_status || '—' : '—';
    if (poll) poll.textContent = runtime ? formatLastPoll(runtime.last_poll) : '—';
    if (operational) {
      const value = runtime ? runtime.source_status || 'WAITING' : 'WAITING';
      operational.textContent = value;
      operational.className = `rep-operational-status rep-operational-${String(value).toLowerCase()}`;
    }
    if (lastError) {
      lastError.textContent = runtime && runtime.last_error ? runtime.last_error : '';
      lastError.hidden = !(runtime && runtime.last_error);
    }
    const blocks = runtime && Array.isArray(runtime.blocks) ? runtime.blocks : [];
    win.$content.querySelectorAll('[data-block-runtime]').forEach(node => {
      const index = Number(node.dataset.blockRuntime);
      const block = blocks[index];
      const status = node.querySelector('[data-value="status"]');
      const blockPoll = node.querySelector('[data-value="poll"]');
      if (status) status.textContent = block ? block.source_status || '—' : '—';
      if (blockPoll) blockPoll.textContent = block ? formatLastPoll(block.last_poll) : '—';
      node.title = block && block.last_error ? block.last_error : '';
    });
    const bar = win.$content.querySelector('.rep-status');
    if (bar) {
      bar.textContent = state.message;
      bar.classList.toggle('rep-status-error', Boolean(state.error));
    }
  };

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
      state.runtimeStatus = {running: false, source_status: 'ERROR', last_error: error.message || String(error), blocks: []};
    }
    patchStatus();
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

  const renderDeviceTab = (content, device) => {
    const identity = element('div', 'rep-grid');
    identity.appendChild(field('Name', device.name, {type: 'text'}, value => { device.name = value; }));
    identity.appendChild(checkboxField('Enabled', device.enabled, value => { device.enabled = value; }));
    content.appendChild(identity);

    content.appendChild(element('h3', 'rep-section-title', 'Source'));
    const sourceGrid = element('div', 'rep-grid rep-device-source-grid');
    sourceGrid.appendChild(field('Endpoint', device.endpoint, {type: 'text', placeholder: '192.168.1.20:502'}, value => { device.endpoint = value; }));
    sourceGrid.appendChild(field('Unit ID', device.unit_id, {min: 0, max: 255}, value => { device.unit_id = numberValue(value); }));
    content.appendChild(sourceGrid);

    content.appendChild(element('h3', 'rep-section-title', 'Destination'));
    const destGrid = element('div', 'rep-destination-grid');
    destGrid.appendChild(field('Port', device.destination.port, {min: 1, max: 65535, readOnly: device.destination.auto_port}, value => {
      device.destination.port = numberValue(value);
      inspectDestination();
    }));
    destGrid.appendChild(checkboxField('Auto Port', device.destination.auto_port, value => {
      device.destination.auto_port = value;
      if (value) {
        const suggestion = localSuggestion();
        device.destination.port = suggestion.port;
        if (device.destination.auto_unit_id) device.destination.unit_id = suggestion.unit_id;
      }
      render();
      inspectDestination();
    }));
    destGrid.appendChild(field('Unit ID', device.destination.unit_id, {min: 0, max: 255, readOnly: device.destination.auto_unit_id}, value => {
      device.destination.unit_id = numberValue(value);
      inspectDestination();
    }));
    destGrid.appendChild(checkboxField('Auto Unit ID', device.destination.auto_unit_id, value => {
      device.destination.auto_unit_id = value;
      if (value) {
        const suggestion = localSuggestion();
        if (device.destination.auto_port) device.destination.port = suggestion.port;
        device.destination.unit_id = suggestion.unit_id;
      }
      render();
      inspectDestination();
    }));
    const ownership = element('div', 'rep-ownership');
    ownership.append(
      element('span', 'rep-field-label', 'Owner'),
      element('strong', '', device.destination.owner || 'replicator'),
      element('span', 'rep-field-label', 'Status'),
      element('strong', 'rep-operational-status', 'WAITING')
    );
    ownership.lastElementChild.dataset.runtime = 'operational';
    destGrid.appendChild(ownership);
    content.appendChild(destGrid);
    if (String(device.destination.status).toUpperCase() === 'IN USE') {
      content.appendChild(element('div', 'rep-validation', `Destination ${device.destination.port}/${device.destination.unit_id} is owned by ${device.destination.owner || 'another producer'} and cannot be claimed by Replicator.`));
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
    content.appendChild(runtime);
    const lastError = element('div', 'rep-last-error');
    lastError.hidden = true;
    content.appendChild(lastError);
  };

  const tableInput = (value, settings, onChange) => {
    const control = document.createElement('input');
    control.type = 'number';
    control.value = value;
    Object.keys(settings).forEach(key => { control[key] = settings[key]; });
    control.addEventListener('click', event => event.stopPropagation());
    control.addEventListener('input', event => onChange(numberValue(event.target.value)));
    return control;
  };

  const renderBlocksTab = (content, device) => {
    const toolbar = element('div', 'rep-block-toolbar');
    toolbar.append(button('Add Block', 'add-block'), button('Duplicate Block', 'duplicate-block'), button('Delete Block', 'delete-block', 'rep-danger'));
    content.appendChild(toolbar);

    const table = element('table', 'rep-block-table');
    const head = document.createElement('thead');
    const header = document.createElement('tr');
    ['#', 'FC', 'Start', 'Count', 'Scan Rate (ms)', 'Status', 'Last Poll'].forEach(text => header.appendChild(element('th', '', text)));
    head.appendChild(header);
    table.appendChild(head);
    const body = document.createElement('tbody');
    device.pull_blocks.forEach((block, index) => {
      const row = document.createElement('tr');
      row.className = index === state.selectedBlock ? 'rep-block-selected' : '';
      row.dataset.action = 'select-block';
      row.dataset.blockIndex = String(index);
      row.appendChild(element('td', 'rep-block-number', String(index + 1)));

      const fcCell = document.createElement('td');
      const select = document.createElement('select');
      [1, 2, 3, 4].forEach(value => {
        const option = document.createElement('option');
        option.value = String(value);
        option.textContent = `FC${value}`;
        option.selected = Number(block.function) === value;
        select.appendChild(option);
      });
      select.addEventListener('click', event => event.stopPropagation());
      select.addEventListener('change', event => { block.function = Number(event.target.value); });
      fcCell.appendChild(select);
      row.appendChild(fcCell);

      const startCell = document.createElement('td');
      startCell.appendChild(tableInput(block.start, {min: 0, max: 65535}, value => { block.start = value; }));
      row.appendChild(startCell);
      const countCell = document.createElement('td');
      countCell.appendChild(tableInput(block.count, {min: 1, max: 65535}, value => { block.count = value; }));
      row.appendChild(countCell);
      const scanCell = document.createElement('td');
      scanCell.appendChild(tableInput(block.scan_rate_ms, {min: 1, step: 1}, value => { block.scan_rate_ms = value; }));
      row.appendChild(scanCell);

      const statusCell = element('td', 'rep-block-status', '—');
      statusCell.dataset.value = 'status';
      const pollCell = element('td', 'rep-block-poll', '—');
      pollCell.dataset.value = 'poll';
      row.dataset.blockRuntime = String(index);
      row.append(statusCell, pollCell);
      body.appendChild(row);
    });
    table.appendChild(body);
    content.appendChild(table);
  };

  const renderEditor = root => {
    const pane = element('section', 'rep-editor-pane');
    pane.appendChild(element('h2', 'rep-editor-title', 'Device Definition'));
    const device = selectedDevice();
    if (!device) {
      const pendingDelete = state.persisted.devices.length > 0 && state.document.devices.length === 0;
      pane.appendChild(element('div', 'rep-empty', pendingDelete ? 'All devices are marked for deletion. Save & Apply to release Replicator-owned destinations, or Discard to restore them.' : 'Select a device or choose Add.'));
      const actions = element('div', 'rep-editor-actions');
      const saveButton = button(state.saving ? 'Saving...' : 'Save & Apply', 'save', 'rep-primary');
      saveButton.disabled = state.saving || !pendingDelete;
      actions.append(saveButton, button('Discard', 'discard'));
      pane.appendChild(actions);
      root.appendChild(pane);
      return;
    }

    if (state.selectedBlock >= device.pull_blocks.length) state.selectedBlock = Math.max(0, device.pull_blocks.length - 1);
    const tabs = element('div', 'rep-folder-tabs');
    tabs.append(
      button('Device', 'tab-device', `rep-folder-tab${state.activeTab === 'device' ? ' rep-folder-tab-active' : ''}`),
      button('Pull Blocks', 'tab-blocks', `rep-folder-tab${state.activeTab === 'blocks' ? ' rep-folder-tab-active' : ''}`)
    );
    pane.appendChild(tabs);
    const tabContent = element('div', 'rep-folder-content');
    if (state.activeTab === 'blocks') renderBlocksTab(tabContent, device);
    else renderDeviceTab(tabContent, device);
    pane.appendChild(tabContent);

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
    search.addEventListener('input', event => { state.search = event.target.value; render(); });
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
      const firstBlock = device.pull_blocks[0] || blankBlock();
      row.append(
        element('strong', '', device.name || 'Unnamed device'),
        element('span', '', `${device.endpoint} • FC${firstBlock.function}${device.pull_blocks.length > 1 ? ` +${device.pull_blocks.length - 1}` : ''} • ${device.enabled ? 'Enabled' : 'Disabled'}`)
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
    if (error) { setMessage(error, true); render(); return; }
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
    pending.forEach(entry => { clearTimeout(entry.timeout); entry.reject(new Error('Replicator window closed.')); });
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
        state.selected = Number(target.dataset.index); state.selectedBlock = 0; state.runtimeStatus = null;
      } else if (action === 'add') {
        state.document.devices.push(blankDevice(state.document.devices.length + 1, localSuggestion()));
        state.selected = state.document.devices.length - 1; state.selectedBlock = 0; state.runtimeStatus = null;
      } else if (action === 'duplicate' && selectedDevice()) {
        const copy = clone(selectedDevice());
        copy.name = `${copy.name} (copy)`;
        const suggestion = localSuggestion();
        Object.assign(copy.destination, {port: suggestion.port, unit_id: suggestion.unit_id, auto_port: true, auto_unit_id: true, owner: 'replicator', status: 'AVAILABLE'});
        state.document.devices.push(copy); state.selected = state.document.devices.length - 1; state.selectedBlock = 0; state.runtimeStatus = null;
      } else if (action === 'delete' && selectedDevice()) {
        state.document.devices.splice(state.selected, 1);
        state.selected = state.document.devices.length ? Math.min(state.selected, state.document.devices.length - 1) : null;
        state.selectedBlock = 0; state.runtimeStatus = null;
      } else if (action === 'tab-device') state.activeTab = 'device';
      else if (action === 'tab-blocks') state.activeTab = 'blocks';
      else if (action === 'select-block') state.selectedBlock = Number(target.dataset.blockIndex);
      else if (action === 'add-block' && selectedDevice()) {
        selectedDevice().pull_blocks.push(blankBlock()); state.selectedBlock = selectedDevice().pull_blocks.length - 1;
      } else if (action === 'duplicate-block' && selectedDevice()) {
        const blocks = selectedDevice().pull_blocks;
        const source = blocks[state.selectedBlock] || blocks[0];
        if (source) { blocks.splice(state.selectedBlock + 1, 0, clone(source)); state.selectedBlock += 1; }
      } else if (action === 'delete-block' && selectedDevice()) {
        const blocks = selectedDevice().pull_blocks;
        if (blocks.length > 1) { blocks.splice(state.selectedBlock, 1); state.selectedBlock = Math.min(state.selectedBlock, blocks.length - 1); }
        else setMessage('A device must keep at least one Pull Block.', true);
      } else if (action === 'discard') {
        state.document = clone(state.persisted);
        state.selected = state.document.devices.length ? Math.min(state.selected || 0, state.document.devices.length - 1) : null;
        state.selectedBlock = 0; state.runtimeStatus = null;
      } else if (action === 'save') { save(); return; }
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
        state.selectedBlock = 0;
        setMessage(persisted.devices.length ? 'Canonical Replicator definitions loaded.' : 'No devices configured. Choose Add to begin.');
        render();
        pollStatus();
      })
      .catch(error => { setMessage(`Replicator runtime unavailable: ${error.message || error}`, true); render(); });
    statusTimer = setInterval(pollStatus, 1000);
  });

  return proc;
};

osjs.register(applicationName, register);
