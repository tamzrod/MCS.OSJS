'use strict';

// Toolkit-only Memory editor. Canonical definitions come solely from Simulator
// load/apply, NEVER the fixture preview. No request is sent on mount except load;
// apply requires an explicit Save & Apply click. No legacy UI import.
const FC = [['fc1', 'Coils (FC1)'], ['fc2', 'Discrete Inputs (FC2)'],
  ['fc3', 'Holding Registers (FC3)'], ['fc4', 'Input Registers (FC4)']];
const STATES = new Set(['RUNNING', 'WAITING', 'STOPPED', 'ERROR', 'IDLE']);
const clone = value => JSON.parse(JSON.stringify(value));
const normalizeDocument = value => ({devices: Array.isArray(value && value.devices) ? value.devices : []});
const numberValue = value => value.trim() === '' ? NaN : Number(value);
const integer = (value, min, max) => Number.isSafeInteger(value) && value >= min && value <= max;
const blankDevice = sequence => ({
  name: `Sim-PLC-${sequence}`, enabled: true,
  mma2: {port: 5020, unit_id: 1, fc1: {start: 0, count: 16}, fc2: {start: 0, count: 16},
    fc3: {start: 0, count: 16}, fc4: {start: 0, count: 16}},
  random_runtime: {fc1_interval_ms: 0, fc2_interval_ms: 0, fc3_interval_ms: 0, fc4_interval_ms: 0}
});
const validateDevice = device => {
  if (!device || typeof device.name !== 'string' || !device.name.trim()) return 'Name is required.';
  if (!device.mma2 || !integer(device.mma2.port, 1, 65535)) return 'Listen Port must be between 1 and 65535.';
  if (!integer(device.mma2.unit_id, 0, 255)) return 'Unit ID must be between 0 and 255.';
  for (const [key, label] of FC) {
    const area = device.mma2[key];
    if (!area || !integer(area.start, 0, 65535) || !integer(area.count, 0, 65535) || area.start + area.count > 65536) {
      return `${label}: invalid Start/Count or range outside the 16-bit address space.`;
    }
    if (!device.random_runtime || !integer(device.random_runtime[`${key}_interval_ms`], 0, 4294967295)) {
      return `${label}: interval must be a nonnegative uint32 integer (0 means None).`;
    }
  }
  return null;
};
const validateDocument = document => {
  if (!document || !Array.isArray(document.devices)) return 'Simulator document is unavailable.';
  const names = new Set();
  for (const device of document.devices) {
    const error = validateDevice(device);
    if (error) return error;
    if (names.has(device.name)) return `Duplicate device name: ${device.name}`;
    names.add(device.name);
  }
  return null;
};
const statusWord = (status, name, field, unavailable) => {
  if (unavailable) return 'UNAVAILABLE';
  if (!status || status.name !== name || !STATES.has(status[field])) return 'UNKNOWN';
  return status[field];
};

const createMemoryEditor = (doc, root, memory) => {
  let documentValue = null;
  let persisted = null;
  let selected = null;
  let loading = true;
  let saving = false;
  let closed = false;
  let message = 'Loading canonical Simulator definitions...';
  let messageError = false;
  let status = null;
  let statusError = null;
  let timer = null;
  let generation = 0;
  const h = (tag, className = '', text) => {
    const el = doc.createElement(tag);
    if (className) el.className = className;
    if (text !== undefined) el.textContent = String(text);
    return el;
  };
  const button = (label, action, disabled = false, className = '') => {
    const el = h('button', `tool-button ${className}`.trim(), label);
    el.type = 'button';
    el.dataset.memoryAction = action;
    el.disabled = disabled;
    return el;
  };
  const selectedDevice = () => documentValue && selected !== null ? documentValue.devices[selected] : null;
  const appliedName = () => {
    const current = selectedDevice();
    const previous = persisted && selected !== null && persisted.devices[selected];
    return current && previous && current.name === previous.name ? previous.name : null;
  };
  const setMessage = (text, error = false) => {
    message = text;
    messageError = error;
    const node = root.querySelector('.tool-status');
    if (node) {
      node.textContent = text;
      node.classList.toggle('error', error);
    }
  };
  const patchStatus = () => {
    const name = appliedName();
    const mma2 = root.querySelector('[data-memory-status="mma2"]');
    const simulation = root.querySelector('[data-memory-status="simulation"]');
    if (mma2) mma2.textContent = name ? statusWord(status, name, 'mma2_status', statusError) : 'UNKNOWN';
    if (simulation) simulation.textContent = name ? statusWord(status, name, 'device_status', statusError) : 'UNKNOWN';
    const diagnostic = root.querySelector('[data-memory-status-error]');
    if (diagnostic) diagnostic.textContent = name && statusError ? `Status unavailable: ${statusError}` : '';
  };
  const stopPolling = () => {
    generation++;
    if (timer !== null) clearTimeout(timer);
    timer = null;
    status = null;
    statusError = null;
    patchStatus();
  };
  const poll = token => {
    const name = appliedName();
    if (closed || saving || token !== generation || !name) return;
    Promise.resolve(memory.status(name)).then(result => {
      if (closed || saving || token !== generation || appliedName() !== name) return;
      const observed = result && result.status;
      status = observed && observed.name === name ? observed : null;
      statusError = status ? null : 'Simulator returned an unmatched device status';
      patchStatus();
    }).catch(error => {
      if (closed || saving || token !== generation || appliedName() !== name) return;
      status = null;
      statusError = error.message || String(error);
      patchStatus();
    }).then(() => {
      if (!closed && !saving && token === generation && appliedName() === name) {
        timer = setTimeout(() => poll(token), 1000);
      }
    });
  };
  const restartPolling = () => {
    stopPolling();
    if (!closed && !loading && !saving && appliedName()) poll(generation);
  };
  const field = (label, value, update, options = {}) => {
    const wrapper = h('label', options.className || 'tool-field');
    wrapper.appendChild(h('span', 'tool-field-label', label));
    const input = h('input');
    input.type = options.type || 'number';
    input.value = value === null || value === undefined ? '' : String(value);
    input.disabled = saving;
    input.setAttribute('aria-label', label);
    if (options.min !== undefined) input.min = options.min;
    if (options.max !== undefined) input.max = options.max;
    input.addEventListener('input', () => {
      update(input.type === 'number' ? numberValue(input.value) : input.value);
      const validation = root.querySelector('[data-memory-validation]');
      const error = validateDocument(documentValue);
      if (validation) validation.textContent = error || '';
      const save = root.querySelector('[data-memory-action="save"]');
      if (save) save.disabled = Boolean(error) || saving;
      patchStatus();
    });
    wrapper.appendChild(input);
    return wrapper;
  };
  const checkbox = (label, value, update) => {
    const wrapper = h('label', 'tool-checkbox');
    const input = h('input');
    input.type = 'checkbox';
    input.checked = Boolean(value);
    input.disabled = saving;
    input.addEventListener('change', () => update(input.checked));
    wrapper.append(input, h('span', '', label));
    return wrapper;
  };
  const render = () => {
    if (closed) return;
    const shell = h('div', 'tool-layout');
    const sidebar = h('aside', 'tool-sidebar');
    sidebar.appendChild(h('h2', '', 'Devices'));
    const actions = h('div', 'tool-actions');
    actions.append(button('Add', 'add', loading || saving), button('Duplicate', 'duplicate', loading || saving || selected === null),
      button('Delete', 'delete', loading || saving || selected === null, 'tool-danger'));
    sidebar.appendChild(actions);
    const list = h('div', 'tool-list');
    if (documentValue) documentValue.devices.forEach((device, index) => {
      const row = button(device.name || 'Unnamed device', 'select', saving, `tool-row${index === selected ? ' selected' : ''}`);
      row.dataset.index = String(index);
      row.appendChild(h('span', '', `Port ${device.mma2 && device.mma2.port} / Unit ${device.mma2 && device.mma2.unit_id}`));
      list.appendChild(row);
    });
    if (!list.children.length) list.appendChild(h('div', 'tool-empty', loading ? 'Loading...' :
      documentValue ? 'No devices configured. Choose Add.' : 'No canonical data available.'));
    sidebar.appendChild(list);
    const editor = h('section', 'tool-editor');
    editor.appendChild(h('h2', '', 'Device Definition'));
    const device = selectedDevice();
    if (device) {
      const runtime = h('div', 'runtime-row');
      runtime.append(h('span', '', 'MMA2:'), h('strong', '', 'UNKNOWN'), h('span', '', 'Simulation:'), h('strong', '', 'UNKNOWN'));
      runtime.children[1].dataset.memoryStatus = 'mma2';
      runtime.children[3].dataset.memoryStatus = 'simulation';
      editor.appendChild(runtime);
      const identity = h('div', 'tool-grid');
      identity.append(field('Name', device.name, value => { device.name = value; }, {type: 'text'}),
        checkbox('Enabled', device.enabled, value => { device.enabled = value; }),
        field('Listen Port', device.mma2.port, value => { device.mma2.port = value; }, {min: 1, max: 65535}),
        field('Unit ID', device.mma2.unit_id, value => { device.mma2.unit_id = value; }, {min: 0, max: 255}));
      editor.appendChild(identity);
      const table = h('div', 'fc-table memory-fc-table');
      const header = h('div', 'fc-row fc-header');
      ['Area', 'Start', 'Count', 'Simulation', 'Interval (ms)', 'Address Range']
        .forEach(label => header.appendChild(h('span', '', label)));
      table.appendChild(header);
      FC.forEach(([key, label]) => {
        const area = device.mma2[key];
        const intervalKey = `${key}_interval_ms`;
        const row = h('div', 'fc-row');
        const range = h('span', 'range-cell');
        const updateRange = () => { range.textContent = integer(area.start, 0, 65535) && integer(area.count, 0, 65535) ?
          (area.count ? `${area.start}-${area.start + area.count - 1}` : 'Unused') : 'Invalid'; };
        const mode = h('span', 'sim-mode-cell');
        const select = h('select');
        select.setAttribute('aria-label', `${label} simulation`);
        select.disabled = saving;
        [['none', 'None'], ['random', 'Random']].forEach(([value, title]) => {
          const option = h('option', '', title); option.value = value; select.appendChild(option);
        });
        select.value = device.random_runtime[intervalKey] > 0 ? 'random' : 'none';
        select.addEventListener('change', () => {
          device.random_runtime[intervalKey] = select.value === 'none' ? 0 : 1000;
          render();
        });
        mode.appendChild(select);
        row.append(h('strong', '', label),
          field('Start', area.start, value => { area.start = value; updateRange(); }, {className: 'cell-field', min: 0, max: 65535}),
          field('Count', area.count, value => { area.count = value; updateRange(); }, {className: 'cell-field', min: 0, max: 65535}),
          mode, field('Interval', device.random_runtime[intervalKey], value => {
            device.random_runtime[intervalKey] = value;
            select.value = value > 0 ? 'random' : 'none';
          }, {className: 'cell-field', min: 0, max: 4294967295}), range);
        updateRange();
        table.appendChild(row);
      });
      editor.appendChild(table);
      const validation = h('div', 'tool-validation');
      validation.dataset.memoryValidation = 'true';
      validation.textContent = validateDocument(documentValue) || '';
      editor.appendChild(validation);
      const controls = h('div', 'editor-actions');
      controls.append(button(saving ? 'Saving...' : 'Save & Apply', 'save', saving || Boolean(validateDocument(documentValue)), 'tool-primary'),
        button('Discard', 'discard', saving));
      editor.appendChild(controls);
    } else {
      editor.appendChild(h('div', 'tool-empty', loading ? 'Loading canonical definitions...' : 'Select a device or choose Add.'));
    }
    const statusMessage = h('div', `tool-status${messageError ? ' error' : ''}`, message);
    editor.appendChild(statusMessage);
    const statusDiagnostic = h('div', 'tool-validation');
    statusDiagnostic.dataset.memoryStatusError = 'true';
    editor.appendChild(statusDiagnostic);
    shell.append(sidebar, editor);
    root.replaceChildren(shell);
    patchStatus();
  };
  const load = async () => {
    try {
      const result = await memory.load();
      if (closed) return;
      persisted = clone(normalizeDocument(result.document));
      documentValue = clone(persisted);
      selected = documentValue.devices.length ? 0 : null;
      message = documentValue.devices.length ? 'Canonical Simulator definitions loaded.' : 'No devices configured. Choose Add to begin.';
      messageError = false;
    } catch (error) {
      if (closed) return;
      persisted = null;
      documentValue = null;
      selected = null;
      message = `Simulator runtime unavailable: ${error.message || error}`;
      messageError = true;
    } finally {
      if (!closed) { loading = false; render(); restartPolling(); }
    }
  };
  const save = async () => {
    if (closed || loading || saving || !documentValue) return;
    const validation = validateDocument(documentValue);
    if (validation) { setMessage(validation, true); return; }
    saving = true;
    stopPolling();
    setMessage('Saving Simulator definitions...');
    render();
    try {
      const result = await memory.apply(clone(documentValue));
      if (closed) return;
      persisted = clone(normalizeDocument(result.document));
      documentValue = clone(persisted);
      if (selected !== null && selected >= documentValue.devices.length) selected = documentValue.devices.length ? 0 : null;
      const stamp = result.completed_at ? ` Applied at ${new Date(result.completed_at).toLocaleString()}.` : '';
      message = `${result.message || 'Simulator definitions applied.'}${stamp}`;
      messageError = false;
    } catch (error) {
      if (closed) return;
      message = `Save & Apply failed: ${error.message || error}`;
      messageError = true; // Keep unsaved document intact on any failed apply.
    } finally {
      if (!closed) { saving = false; render(); restartPolling(); }
    }
  };
  const onClick = event => {
    const target = event.target.closest('[data-memory-action]');
    if (!target || target.disabled || closed || loading || saving || !documentValue) return;
    const action = target.dataset.memoryAction;
    if (action === 'save') { save(); return; }
    if (action === 'select') {
      stopPolling(); selected = Number(target.dataset.index);
      setMessage('Editing canonical Simulator definition. Changes are local until Save & Apply.');
    } else if (action === 'add') {
      stopPolling(); documentValue.devices.push(blankDevice(documentValue.devices.length + 1));
      selected = documentValue.devices.length - 1;
      setMessage('New device added locally. Save & Apply to persist.');
    } else if (action === 'duplicate' && selectedDevice()) {
      stopPolling(); const copy = clone(selectedDevice()); copy.name += ' (copy)';
      documentValue.devices.push(copy); selected = documentValue.devices.length - 1;
      setMessage('Device duplicated locally. Save & Apply to persist.');
    } else if (action === 'delete' && selectedDevice()) {
      stopPolling(); documentValue.devices.splice(selected, 1);
      selected = documentValue.devices.length ? Math.min(selected, documentValue.devices.length - 1) : null;
      setMessage('Device removed locally. Save & Apply to persist.');
    } else if (action === 'discard') {
      stopPolling(); documentValue = clone(persisted);
      selected = documentValue.devices.length ? Math.min(selected === null ? 0 : selected, documentValue.devices.length - 1) : null;
      setMessage('Unapplied changes discarded.');
    } else return;
    render(); restartPolling();
  };
  root.addEventListener('click', onClick);
  render();
  load();
  return {
    destroy: () => { closed = true; stopPolling(); root.removeEventListener('click', onClick); }
  };
};

module.exports = {blankDevice, normalizeDocument, validateDevice, validateDocument, statusWord, createMemoryEditor};
